package dispatch

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/compact"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/registry"
)

// InvokeResult carries content plus the routing metadata the API surface
// needs (which node served it, how long it took, whether it was cached).
type InvokeResult struct {
	Content    string  `json:"content"`
	Reasoning  string  `json:"reasoning,omitempty"` // reasoning_content from reasoning-class models, when the answer itself is in Content
	NodeID     string  `json:"node_id"`
	Model      string  `json:"model"`
	DurationMs float64 `json:"duration_ms"`
	Cached     bool    `json:"cached"`
	Provider   string  `json:"provider,omitempty"` // "ollama"|"lmstudio"|<remote name>
	// PromptTokens/CompletionTokens feed the per-node token monitor (Planner
	// panel). Populated on the non-streaming dispatch path only (invokeChatHTTP,
	// invokeOllama) — the packaged app's primary path (see the SSE-in-WKWebView
	// note elsewhere in this file); streaming responses don't reliably carry
	// usage in this codebase's request shape today, so they're left at zero
	// rather than guessed.
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	// UsageEstimated is true when PromptTokens/CompletionTokens were filled
	// by compact.EstimateTokens because the upstream reported all zeros
	// (LM Studio OpenAI-compat, KMA-91). Provider numbers win when present.
	UsageEstimated bool `json:"usage_estimated,omitempty"`
	// CacheRead is prompt_tokens_details.cached_tokens when the provider reports it (KMA-139).
	CacheRead int `json:"cache_read,omitempty"`
	// InputNoCache is prompt tokens that were not cache hits.
	InputNoCache int `json:"input_no_cache,omitempty"`
	// ReasoningTokens is native_tokens_reasoning when the provider reports it (KMA-212).
	ReasoningTokens int `json:"reasoning_tokens,omitempty"`
	// Source classifies which inference backend produced these tokens
	// (KMA-87 Settings → Usage). Nil means unattributed.
	Source *UsageSource `json:"source,omitempty"`
	// ToolCalls are populated when the request included OpenAI-style tools and
	// the model returned message.tool_calls (agent dual protocol).
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	// RequestID correlates this invoke with tasks.id / activity.request_id (KMA-219).
	RequestID string `json:"request_id,omitempty"`
	// Mode is ask|agent|… for usage JSONL joins (KMA-219).
	Mode string `json:"mode,omitempty"`
	// GenerationID is the provider's completion/generation id (OpenAI-compatible
	// top-level "id", e.g. OpenRouter gen-…). KMA-211 first-party RCA.
	GenerationID string `json:"generation_id,omitempty"`
	// FinishReason is choices[0].finish_reason when present (stop, tool_calls, …).
	FinishReason string `json:"finish_reason,omitempty"`
}

// ChatTool is an OpenAI-compatible tools[] entry (type=function).
type ChatTool struct {
	Type     string           `json:"type"` // "function"
	Function ChatToolFunction `json:"function"`
}

// ChatToolFunction is the function schema inside a ChatTool.
type ChatToolFunction struct {
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Parameters  map[string]any `json:"parameters,omitempty"`
}

// ToolCall is one OpenAI-compatible tool_calls[] entry from the assistant.
type ToolCall struct {
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"` // "function"
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"` // JSON object string
	} `json:"function"`
}

// OllamaRequest maps to Ollama's /api/generate endpoint.
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
	// PromptEvalCount/EvalCount are Ollama's own field names for input/output
	// token counts, present on the final (non-streaming, or done:true) response.
	PromptEvalCount int `json:"prompt_eval_count,omitempty"`
	EvalCount       int `json:"eval_count,omitempty"`
}

// chatRequest is the OpenAI /v1/chat/completions request body, used by
// LM Studio and remote.go.
type chatRequest struct {
	Model      string        `json:"model"`
	Messages   []chatMessage `json:"messages"`
	Stream     bool          `json:"stream"`
	MaxTokens  int           `json:"max_tokens,omitempty"` // output reserve; bounds runaway generation and guarantees input+output ≤ window
	Tools      []ChatTool    `json:"tools,omitempty"`
	ToolChoice any           `json:"tool_choice,omitempty"` // "auto" | "none" | object
}

type chatMessage struct {
	Role    string `json:"role"`
	Content any    `json:"content"`
	// ReasoningContent carries the chain-of-thought some reasoning-class
	// models (minimax-m3, deepseek-r1, …) return alongside — or instead of —
	// the final answer. omitempty keeps it off outbound request bodies.
	ReasoningContent string `json:"reasoning_content,omitempty"`
	// ToolCalls carries OpenAI-compatible function calls from the assistant.
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	// ToolCallID / Name are set on role=tool messages (KMA-205/218).
	ToolCallID string `json:"tool_call_id,omitempty"`
	Name       string `json:"name,omitempty"`
}

type chatResponse struct {
	ID      string `json:"id"`
	Choices []struct {
		Message      chatMessage `json:"message"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage chatUsageBlock `json:"usage"`
}

// chatStreamChunk is one SSE delta from an OpenAI-compatible streaming endpoint.
type chatStreamChunk struct {
	Choices []struct {
		Delta struct {
			Content          string `json:"content"`
			ReasoningContent string `json:"reasoning_content"`
		} `json:"delta"`
	} `json:"choices"`
	// Some providers emit usage on the final stream chunk (or a trailing chunk).
	Usage chatUsageBlock `json:"usage"`
}

// UsageHook is called after a successful invoke so the usage log can
type UsageHook func(InvokeResult)

// UsageSource is the KMA-87 classification of one InvokeResult.
type UsageSource struct {
	Class          string `json:"class"`                     // local | remote_key | unattributed
	ID             string `json:"id,omitempty"`              // node id or remote name
	Provider       string `json:"provider,omitempty"`        // ollama | lmstudio | vendor label
	Host           string `json:"host,omitempty"`            // host:port or api_base for local
	KeyFingerprint string `json:"key_fingerprint,omitempty"` // last 4 of remote API key; never the secret
}

type Dispatcher struct {
	reg       *registry.Registry
	client    *http.Client
	provider  config.ProviderConfig
	usageHook UsageHook
}

func New(reg *registry.Registry, provider config.ProviderConfig) *Dispatcher {
	panic(
		// Tuned transport with connection pooling. Go's default transport caps
		// MaxIdleConnsPerHost at 2, which serializes connection reuse against a
		// single backend (the common case here: one Ollama/LM Studio host). Raising
		// the per-host idle pool lets concurrent requests reuse warm keep-alive
		// connections instead of paying a fresh TCP+TLS handshake each time.
		"fake")
}

// unlimited in-flight; only idle reuse is pooled

// Match RemoteRequestTimeout: local models (esp. LM Studio) often
// buffer the full completion before sending response headers. A 120s
// header timeout cut off slow Ask completions while the model still
// finished — UI aborted, LM Studio log still showed success.

// ctx alone; see invokeChatHTTPOnce

// SetUsageHook installs the KMA-87 persistence callback. Called once at
// coordinator start; nil disables recording.
func (d *Dispatcher) SetUsageHook(h UsageHook) { panic("fake") }

func (d *Dispatcher) noteUsage(ctx context.Context, nodeID string, res InvokeResult) { panic("fake") }

// LocalSource classifies a registry node as a free local worker.
func LocalSource(node *registry.Node, provider string) *UsageSource { panic("fake") }

// RemoteSource classifies a configured API key. Fingerprint is last-4 only.
func RemoteSource(name, provider, apiKey string) *UsageSource { panic("fake") }

// MaskKeyFingerprint returns the last 4 characters of a key, or **** when
// the secret is too short to truncate safely.
func MaskKeyFingerprint(key string) string { panic("fake") }

// Invoke satisfies chain.Invoker. Returns only the content string so the
// chain engine doesn't need to know about routing metadata.
func (d *Dispatcher) Invoke(ctx context.Context, model, prompt string) (string, error) { panic("fake") }

// InvokeRich picks the best registered node and routes by its Backend, failing
// over to the next-best healthy node if one errors (e.g. its backend is down
// and returns 400). Returns the first success, or the last error if all fail.
func (d *Dispatcher) InvokeRich(ctx context.Context, model, prompt string) (InvokeResult, error) {
	panic("fake")
}

// InvokeRichOnBackend is InvokeRich restricted to nodes whose Backend matches
// backend (e.g. "chidori-nagasa"). Empty backend keeps the old all-local pool.
func (d *Dispatcher) InvokeRichOnBackend(ctx context.Context, model, prompt, backend string) (InvokeResult, error) {
	panic("fake")
}

// caller cancelled / timed out — stop retrying

// InvokeVisionLocal routes multimodal vision through healthy local nodes
// (LM Studio / Ollama OpenAI-compat / AirLLM) using the same payload as remote
// InvokeVision.
func (d *Dispatcher) InvokeVisionLocal(ctx context.Context, model, prompt string, images []ImagePart) (InvokeResult, error) {
	panic("fake")
}

// InvokeNodeByID routes directly to a specific registered node, bypassing
// load-based selection. Used by the stress harness to target one node.
func (d *Dispatcher) InvokeNodeByID(ctx context.Context, nodeID, model, prompt string) (InvokeResult, error) {
	panic("fake")
}

// InvokeStream streams tokens through cb, with failover *before the first
// token* — once bytes are on the wire we can't switch nodes, so a mid-stream
// error is returned as-is.
func (d *Dispatcher) InvokeStream(ctx context.Context, model, prompt string, cb func(string)) (InvokeResult, error) {
	panic("fake")
}

// InvokeStreamOnBackend is InvokeStream restricted to a single local backend.
func (d *Dispatcher) InvokeStreamOnBackend(ctx context.Context, model, prompt string, cb func(string), backend string) (InvokeResult, error) {
	panic("fake")
}

// already streaming, or cancelled — can't fail over

func (d *Dispatcher) invokeNode(ctx context.Context, node *registry.Node, model, prompt string) (InvokeResult, error) {
	panic("fake")
}

// requireLoadedModel guards the OpenAI-compatible local backends (LM Studio,
// AirLLM), whose /v1/models reachability probe reports Healthy even when no
// model is actually loaded. Sending an empty "model" to /chat/completions
// then 400s with an opaque message; this replaces it with an actionable one.
func requireLoadedModel(node *registry.Node, model, label string) error { panic("fake") }

func (d *Dispatcher) invokeNodeStream(ctx context.Context, node *registry.Node, model, prompt string, cb func(string)) (InvokeResult, error) {
	panic("fake")
}

// invokeOllama posts to Ollama's /api/generate with stream=false.
func (d *Dispatcher) invokeOllama(ctx context.Context, node *registry.Node, model, prompt string, start time.Time) (InvokeResult, error) {
	panic("fake")
}

// invokeOllamaStream calls Ollama with stream=true and decodes the
// newline-delimited JSON stream.
func (d *Dispatcher) invokeOllamaStream(ctx context.Context, node *registry.Node, model, prompt string, cb func(string), start time.Time) (InvokeResult, error) {
	panic("fake")
}

// skip malformed line

// invokeChat posts a non-streaming OpenAI /v1/chat/completions request.
// Shared by LM Studio nodes and the remote backend.
func (d *Dispatcher) invokeChat(ctx context.Context, node *registry.Node, baseURL, model, prompt string, start time.Time, providerLabel string) (InvokeResult, error) {
	panic("fake")
}

// invokeChatStream posts a streaming chat request and feeds deltas to cb.
func (d *Dispatcher) invokeChatStream(ctx context.Context, node *registry.Node, baseURL, model, prompt string, cb func(string), start time.Time, providerLabel string) (InvokeResult, error) {
	panic("fake")
}

// invokeChatHTTP is a backend-agnostic OpenAI-compatible chat call used by
// both LM Studio (no auth) and Remote (Bearer auth). authHeader == "" skips auth.
//
// UNIFIED Wave 1 / H.2: transient upstream failures (504/502/503, connection
// reset, unexpected EOF) are retried a few times with short backoff. Auth and
// client errors (4xx other than 429) are not retried. Context cancellation
// aborts between attempts immediately.
func invokeChatHTTP(ctx context.Context, client *http.Client, baseURL, authHeader, model, prompt string, start time.Time, nodeID, providerLabel, vendor string) (InvokeResult, error) {
	panic("fake")
}

// invokeChatHTTPTools is like invokeChatHTTP but may attach OpenAI tools[].
func invokeChatHTTPTools(ctx context.Context, client *http.Client, baseURL, authHeader, model, prompt string, tools []ChatTool, start time.Time, nodeID, providerLabel, vendor string) (InvokeResult, error) {
	panic("fake")
}

// invokeChatHTTPMessages posts an explicit messages[] history (KMA-205/218).
func invokeChatHTTPMessages(ctx context.Context, client *http.Client, baseURL, authHeader, model string, msgs []chatMessage, flat string, tools []ChatTool, start time.Time, nodeID, providerLabel, vendor string) (InvokeResult, error) {
	panic("fake")
}

// 100ms, 200ms

func invokeChatHTTPOnce(ctx context.Context, client *http.Client, url, authHeader string, body []byte, start time.Time, nodeID, providerLabel, prompt, model string, meta OpenAICompatMeta) (InvokeResult, error) {
	panic("fake")
}

// Zero out whatever client-level Timeout the caller's *http.Client
// happens to carry (a shallow copy — the Transport, and its connection
// pool, is a pointer and stays shared) and rely purely on ctx instead.
// Same reasoning as invokeChatHTTPStream just below: a client-level
// Timeout is a second, independent cap that can fire out of step with
// whatever deadline the caller's context actually established — that
// mismatch is exactly what caused a real bug (a slow hosted reasoning
// model failing with "Client.Timeout exceeded" well before it should
// have, because Remote's client used to hardcode its own 300s
// regardless of the caller's own — also 300s, but enforced twice,
// independently, is still the wrong shape). The local Dispatcher's own
// client (120s) gets the same treatment here so local non-streaming
// calls behave consistently with local streaming, which already relied
// on ctx alone (see below).

// A reasoning-class model that put its entire answer in reasoning_content
// and left content empty would otherwise render as a blank "(no
// response)". Promote the reasoning to be the answer in that case; when
// content is present, keep reasoning separate so the UI can show it as a
// collapsible "thinking" section rather than mixing it into the answer.
// Do not promote when tool_calls are present — content may be empty while
// the model is requesting tools.

// applyUsageEstimate prefers provider-reported token counts. When both are
// zero (LM Studio OpenAI-compat, KMA-91), fall back to compact.EstimateTokens
// (~4 chars/token), the same heuristic GET /api/sessions/{id}/usage uses.
func applyUsageEstimate(prompt, content string, promptTokens, completionTokens int) (int, int, bool) {
	panic("fake")
}

func splitCacheRead(promptTokens, cachedTokens int) (cacheRead, inputNoCache int) { panic("fake") }

// isRetryableUpstream reports whether an invokeChatHTTP error is worth
// another attempt (UNIFIED H.2). Matches gateway timeouts, overload, and
// common transport resets — not auth/validation failures.
func isRetryableUpstream(err error) bool { panic("fake") }

// Don't retry local "connection refused" forever via this path —
// local dispatcher already fails over nodes. Remotes briefly refusing
// (proxy bounce) are still covered by "connection reset"/EOF below.

func ctxDone(err error) bool { panic("fake") }

// invokeChatHTTPStream is the SSE variant of invokeChatHTTP.
func invokeChatHTTPStream(ctx context.Context, client *http.Client, baseURL, authHeader, model, prompt string, cb func(string), start time.Time, nodeID, providerLabel, vendor string) (InvokeResult, error) {
	panic("fake")
}

// Streaming responses may keep the connection open longer than the
// default client timeout; use a copy with no timeout for the request,
// the context cancels it instead.

// Buffer reasoning deltas but don't stream them into the answer;
// they're only used as a fallback below if no content arrives.

// Same blank-answer guard as the non-streaming path: if the model only
// ever emitted reasoning deltas, surface those as the answer rather than
// returning an empty string. Replay them through cb so a streaming client
// (dev browser) actually sees the text.

// KMA-124: stream path previously left tokens at 0 forever (KMA-91 only
// covered non-stream). Prefer provider usage from a stream chunk; else estimate.

// pickRanked returns the healthy candidate nodes for a model, ordered by load
// (lowest first). The caller tries them in order, enabling failover.
func (d *Dispatcher) pickRanked(model string) ([]*registry.Node, error) { panic("fake") }

func (d *Dispatcher) pickRankedBackend(model, backend string) ([]*registry.Node, error) {
	panic("fake")
}

// Address-empty / api_base-empty nodes (KMA-92) are never candidates —
// they can sit in Healthy() after a naive Register but cannot be dialed.

// Filter by model if specified; fall back to reachable nodes already
// constrained by backend so a pinned phone route cannot leak to Ollama.

// DetectModels queries an Ollama server's /api/tags endpoint and returns
// the list of locally available model names. Used by the worker at startup
// and by /v1/nodes/{id}/models.
func DetectModels(ctx context.Context, ollamaURL string) ([]string, error) { panic("fake") }

// DetectOpenAIModels queries an OpenAI-compatible server's GET /v1/models
// endpoint (LM Studio, vLLM, llama.cpp server, AirLLM, ...) and returns the
// model IDs it reports. Used by the "attach node by IP" and "discover on
// LAN" inference-source flows (internal/api/nodes_handlers.go) to identify
// what's running at a bare endpoint without requiring any lclreason-specific
// software there.
func DetectOpenAIModels(ctx context.Context, base string) ([]string, error) { panic("fake") }

// effectiveModel resolves an empty ("auto") model to the node's own loaded
// model, so a request without an explicit model still names one for backends
// that require it (Ollama). LM Studio / AirLLM tolerate an empty model and use
// whatever is loaded.
func effectiveModel(n *registry.Node, model string) string { panic("fake") }

func backendOf(n *registry.Node) string { panic("fake") }

func lmStudioURL(n *registry.Node) string { panic("fake") }

// Default LM Studio listens on 1234.

// openAICompatURL returns the OpenAI-compatible base URL for a node (AirLLM and
// other /v1 servers). Honors an explicit APIBase; otherwise assumes /v1 on the
// node's address.
func openAICompatURL(n *registry.Node) string { panic("fake") }

func idOf(n *registry.Node) string { panic("fake") }

func nodeBearerAuth(n *registry.Node) string { panic("fake") }
