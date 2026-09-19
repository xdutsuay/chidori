package api

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/activity"
	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/browser"
	"github.com/xdutsuay/lclreason/internal/cache"
	"github.com/xdutsuay/lclreason/internal/chain"
	"github.com/xdutsuay/lclreason/internal/compact"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/crash"
	"github.com/xdutsuay/lclreason/internal/debugger"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/flightlog"
	"github.com/xdutsuay/lclreason/internal/harness"
	"github.com/xdutsuay/lclreason/internal/harnesstrace"
	"github.com/xdutsuay/lclreason/internal/lsp"
	"github.com/xdutsuay/lclreason/internal/mcpclient"
	"github.com/xdutsuay/lclreason/internal/memory"
	"github.com/xdutsuay/lclreason/internal/registry"
	"github.com/xdutsuay/lclreason/internal/sandbox"
	"github.com/xdutsuay/lclreason/internal/secrets"
	"github.com/xdutsuay/lclreason/internal/session"
	"github.com/xdutsuay/lclreason/internal/store"
	"github.com/xdutsuay/lclreason/internal/terminal"
	"github.com/xdutsuay/lclreason/internal/tools"
	"github.com/xdutsuay/lclreason/internal/usage"
	"github.com/xdutsuay/lclreason/internal/workflow"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

type Handlers struct {
	reg   *registry.Registry
	dis   *dispatch.Dispatcher
	pool  *dispatch.Pool
	eng   *chain.Engine
	cache *cache.Cache
	store *store.Store
	mem   *memory.VectorStore
	tools *tools.Registry
	ws    *WSManager

	compactor     *compact.Compactor
	compactionLog *compactionLog
	bus           *activity.Bus
	secretsPath   string
	secretStore   secrets.Store // KMA-256: seals provider api_key at rest (nil = plaintext YAML legacy)

	wsFS            *workspace.FS
	codeCfg         config.CodeConfig
	sessions        *session.Store
	sessionDispatch *agent.SessionDispatch // HRD-001 session concurrency controller
	harnessReg      *harness.Registry      // HRD-001 external harness registry
	term            *terminal.Manager
	lspMu           sync.Mutex  // serializes syncLSPWorkspace calls (root/trust changes are rare, user-driven)
	lspGopls        *lsp.Client // ADR-0006 — always non-nil, idle until Start()

	mcpMgr *mcpclient.Manager // ADR-0012 — always non-nil, no servers configured until Reconcile()

	// summarizeInvoker, when non-nil, is used by SessionSummarize instead of
	// the pool (tests inject a fixed-response stub; production leaves it nil).
	summarizeInvoker compact.Invoker

	// usage is the KMA-87 global token log (nil-safe).
	usage *usage.Store

	cfg   *config.Config
	cfgMu sync.Mutex // guards cfg mutation + Save (config.Config itself isn't otherwise concurrency-safe)

	// chidori LAN companion API (chidori-nagasa).
	companion     *registry.ChidoriCompanion
	companionPort int // dedicated listen port (protocol 1.2.0 default 8027)
	runs          RunTracker

	// Run & Debug panel (BE.15) — a single active Delve DAP session at a
	// time, same "one thing at a time" precedent as the agent loop and the
	// companion pairing state.
	debugMu      sync.Mutex
	debugSession *debugger.Session
	debugProgram string // workspace-relative program dir of the active session, for GET /api/debug/status

	// Flight is the Wave 0 NDJSON recorder (nil-safe). Path exposed via GET /api/flightlog/path.
	flight *flightlog.Logger

	// crashDir holds {dataDir}/crashes for KMA-234 pending markers (nil-safe empty).
	crashDir string
	// crashReporter is an optional local sink (file/no-op); never network-upload by default.
	crashReporter crash.Reporter

	// agentCheckpoints holds paused H.3 clarify-gate state keyed by task_id.
	agentCheckpoints *agentCheckpointStore

	// agentJobs is the AG.16 background AgentRun registry (nil-safe).
	agentJobs *agentJobRegistry
	// leases is the process-wide file lease map (KMA-140). Shared by every
	// AgentRun / background job so two sessions cannot silently edit one path.
	leases *agent.LeaseManager
	// runCancels lets Stop cancel a foreground Ask/Agent run without relying
	// on the original HTTP abort (KMA-61 / KMA-6).
	runCancels *runCancelRegistry

	// runInjects queues companion POST /runs/{id}/messages while a run is active.
	runInjects *runInjectRegistry

	// harnessStore persists optional harness trace JSONL (HV.3).
	harnessStore *harnesstrace.FileStore

	// wfCron schedules trusted cron workflows (ADR-0015 Phase 2).
	wfCron *workflow.CronScheduler

	// wfRuntime is the shared workspace-scoped workflow executor (Ch12 P1).
	wfRuntimeMu   sync.Mutex
	wfRuntimeRoot string
	wfRuntime     *workflow.Runtime

	// wfNotifyMu / wfNotifies — recent auto-triggered workflow starts for UI toasts.
	wfNotifyMu sync.Mutex
	wfNotifies []workflowNotify
}

// ---- node management ----

func (h *Handlers) RegisterNode(w http.ResponseWriter, r *http.Request) {
	panic(
		// url is accepted as an alias for api_base: the 2026-08-07 repro (KMA-92)
		// posted {"id":"manual-lmstudio","url":"http://127.0.0.1:1234/v1",...}
		// which decoded into Address="" / APIBase="" and was stored as healthy.
		"fake")
}

func (h *Handlers) Heartbeat(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Deliver any queued control command (e.g. a restart requested via
// POST /v1/nodes/{id}/restart) on this heartbeat.

func (h *Handlers) DeregisterNode(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ListNodes(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- OpenAI-compatible chat ----

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Chain    string        `json:"chain"`    // lclreason extension
	Provider string        `json:"provider"` // lclreason extension: "" | "local" | "auto" | "hybrid" | "ollama" | "lmstudio" | <remote name>
	Stream   bool          `json:"stream"`
	Research bool          `json:"research"` // H.5: opt-in read-only tools for Ask (default off)
	ClientID string        `json:"client_id,omitempty"`
	// TaskID is optional. When set, activity/flight log keys under it so the
	// packaged Ask UI can correlate Diagnostics with a slow local completion.
	// Empty → server generates one and returns it as X-Task-ID (same as AgentRun).
	TaskID string `json:"task_id,omitempty"`
	// SessionID binds Ask into the shared IDE session so prior turns (incl.
	// agent tool work) are merged into messages before completion.
	SessionID string `json:"session_id,omitempty"`
	// AllowTruncate opts into dispatcher clamp when the prompt exceeds the
	// model window (KMA-5). Default is fail-closed (HTTP 400).
	AllowTruncate bool `json:"allow_truncate,omitempty"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	// ReasoningContent surfaces a reasoning-class model's chain-of-thought to
	// the UI (rendered as a collapsible "thinking" section) when it's distinct
	// from the answer. Empty for ordinary models and for cached answers.
	ReasoningContent string `json:"reasoning_content,omitempty"`
}

type ChatResponse struct {
	ID       string            `json:"id"`
	Object   string            `json:"object"`
	Created  int64             `json:"created"`
	Model    string            `json:"model"`
	Choices  []Choice          `json:"choices"`
	Usage    Usage             `json:"usage"`
	ServedBy *browser.ServedBy `json:"served_by,omitempty"` // lclreason extension: node attribution
	Notice   string            `json:"notice,omitempty"`    // lclreason extension: trust/verification warning
	Grounded bool              `json:"grounded"`            // true if a tool grounded the answer
}

type Choice struct {
	Index   int         `json:"index"`
	Message ChatMessage `json:"message"`
}

type Usage struct {
	PromptTokens     int  `json:"prompt_tokens"`
	CompletionTokens int  `json:"completion_tokens"`
	TotalTokens      int  `json:"total_tokens"`
	Estimated        bool `json:"estimated,omitempty"` // true when filled by compact.EstimateTokens (KMA-91)
}

func (h *Handlers) ChatCompletions(w http.ResponseWriter, r *http.Request) { panic("fake") }

// KMA-90: empty provider+model silently routed to a random healthy node
// and returned confident hallucinations. Apply Settings defaults first;
// if still neither field is set, fail closed.

// Leave model empty when the client says "auto" — each backend then uses its
// natural default (the local node's loaded model, or a remote's configured
// model). Forcing "llama3" here used to leak that name to remote providers
// like NVIDIA, which rejected it.

// Shared-session Ask history: prepend prior IDE turns (agent tools via
// BuildOpenAIToolHistory) before rules so compactPrompt sees the full thread.

// Custom instructions — global (ADR-0011) + project (.lclreason/rules,
// .cursorrules): injected as a leading system message so they survive
// compaction like any other message, rather than being bolted on after
// the fact.

// Compact the conversation before building the prompt: keep recent turns
// verbatim, summarize older turns, and RAG-recall the most relevant ones.
// The per-user report drives the dashboard's context inspector.

// Smart route: resolve "auto" once before hybrid/Execute branches.

// Cache check. Cache key still includes only model+prompt — the same
// answer is valid regardless of which backend produced it.

// dispatch.RemoteRequestTimeout (not 120s) so a slow reasoning model on
// the remote leg isn't cut off — the packaged desktop app uses this
// non-streaming path for every chat message (SSE doesn't survive the
// Wails WKWebView), so this is the live path, not just a fallback.
// Matches streamChat's budget — see RemoteRequestTimeout's doc comment
// for why this used to be a separate, independently-hardcoded 300s that
// caused a real bug.

// Semantic cache: near-duplicate prompt (different wording) → reuse answer.

// If a chain is specified, route through the engine (engine uses
// the dispatcher's Invoker, so it stays local-only for now — per-step
// remote routing is a follow-up). Otherwise call the pool directly so
// the provider hint takes effect.

// Persist plan and trace if present.

// No explicit chain. When the planner is enabled, run it to decompose
// the request and execute any tool steps (e.g. web_search), then feed
// the gathered results to the model as context. The final answer still
// goes through the pool, so provider routing is preserved.

// H.5: read-only research tools only when the user opts in — default
// Ask stays a one-shot answer with no tool loop.

// Distributed browser/mobile compute nodes contribute keyword analysis,
// which we inject so they actually influence the answer (not just appear
// in served_by).

// Only cache a real answer. Caching an empty/blank output (e.g. a provider
// hiccup that still returned 200) would pin "(no response)" for every
// identical prompt until TTL — a confusing "it's stuck/cached wrong" bug.

// Attribute the browser nodes that contributed to this answer.

// a tool plan ran → grounded

// gatherBrowserContributions dispatches a keyword/clause microtask to every
// connected WebSocket compute node, in parallel, and collects their normalized
// results. Best-effort: returns nil when no browser nodes are connected, and
// silently drops nodes that error or time out.
func (h *Handlers) gatherBrowserContributions(ctx context.Context, prompt string) []browser.Contribution {
	panic("fake")
}

// injectBrowserKeywords prepends the merged keyword analysis from browser nodes
// to the prompt so their contribution actually informs the answer. Returns the
// prompt unchanged when there are no contributions.
func injectBrowserKeywords(prompt string, contribs []browser.Contribution) string { panic("fake") }

// runHybrid runs the prompt on local + the default remote concurrently,
// emits each resolved leg to the activity timeline, and returns the first
// successful answer plus a worker label (e.g. "hybrid:nvidia").
//
// H.5 / running_issue.md P1-2 decision (2026-07-09, Opus 4.8): this used to
// also run a "models disagree" lexical cross-check (bowCosine) whenever both
// legs succeeded. But ExecuteHybrid cancels the losing leg the instant the
// winner resolves (see its own doc comment) — answering fast beats a
// hallucination hint — so by the time this ran, two successful Contents
// essentially never coexisted; the check was unreachable in ordinary
// operation (confirmed: ExecuteHybrid's own returned `legs` only ever
// contains two OK entries when the first-arriving leg happened to fail).
// The alternative was adding an opt-in "await both legs" mode so the
// cross-check could actually fire — rejected here because nothing in this
// codebase currently wants to pay that latency cost; no caller, UI toggle,
// or config flag asks for a "verify both answers" mode today, so building
// it now would be speculative. Removed the dead check instead, along with
// its only consumers (bowCosine, wordFreq, LegResult.Content) — dead code
// that LOOKS like a working feature is worse than no code at all. The
// `notice` return value stays (always "" here now) since ChatResponse.Notice
// is a general-purpose extension field other callers may still populate.
func (h *Handlers) runHybrid(ctx context.Context, model, prompt string, em activity.Emitter) (content, reasoning, worker, notice string, err error) {
	panic("fake")
}

// ListProviders returns the names of configured remote providers plus
// "local" (always available — whether or not nodes are healthy is shown
// separately via /v1/nodes). Used by the chat UI provider dropdown.
func (h *Handlers) ListProviders(w http.ResponseWriter, r *http.Request) { panic("fake") }

// name, model, has_key, default

func (h *Handlers) ListTasks(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) GetCacheStats(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ClearCache(w http.ResponseWriter, r *http.Request) { panic("fake") }

// CleanupCache evicts expired entries on demand and reports how many were
// removed. (Clear wipes everything; this only drops entries past their TTL.)
func (h *Handlers) CleanupCache(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Metrics exposes cluster state in Prometheus text exposition format so the
// coordinator can be scraped by Prometheus / viewed in Grafana.
func (h *Handlers) Metrics(w http.ResponseWriter, r *http.Request) { panic("fake") }

// stable output for scrapers

// SetProvider switches the default backend at runtime. Per-request provider
// hints in chat completions still override.
func (h *Handlers) SetProvider(w http.ResponseWriter, r *http.Request) { panic("fake") }

// streamChat handles a streaming chat request. It emits OpenAI-format
// `chat.completion.chunk` events as tokens arrive, then `[DONE]`.
// Chains are run non-streaming under the hood (the chain engine doesn't
// stream per-step yet) and emitted as a single final chunk.
func (h *Handlers) streamChat(w http.ResponseWriter, r *http.Request, req ChatRequest, prompt, model, taskID string) {
	panic("fake")
}

// Cache hit short-circuit — emit the whole answer as one chunk. The
// activity event matters for observability: a cached answer used to leave
// the developer timeline completely empty, which reads as "my message
// never reached the backend" when it actually did.

// Semantic cache short-circuit — emit the matched answer as one chunk.

// Chains run non-streaming for now; emit the final output as one chunk.

// No explicit chain. Run the planner (if enabled) to gather tool
// results first, then stream the final answer through the pool so
// provider routing is preserved.

// Distributed browser/mobile compute nodes contribute keyword
// analysis, injected so they influence the streamed answer.

// Hybrid races local + remote; emit the winner as one chunk.

// Token-by-token streaming via the pool, coalesced so high tok/s
// paths do not Flush a full OpenAI envelope on every upstream piece
// (Wave A / RT-THPUT-1).

// Only cache a real answer. Caching an empty/blank output (e.g. a provider
// hiccup that still returned 200) would pin "(no response)" for every
// identical prompt until TTL — a confusing "it's stuck/cached wrong" bug.

// lastUserMessage returns the content of the most recent user-role message,
// falling back to fallback when there is none (e.g. a client that sent only a
// system message). This is what the planner should reason about — the
// compacted prompt wraps it in rules/history framing.
func lastUserMessage(msgs []ChatMessage, fallback string) string { panic("fake") }

// streamChunkJSON builds one OpenAI `chat.completion.chunk` payload.
// finishReason is non-empty only on the last chunk ("stop"); intermediate
// chunks set it to null and carry a content delta.
func streamChunkJSON(id, model, contentDelta, finishReason string) string { panic("fake") }

// streamReasoningJSON emits a chunk carrying only a reasoning_content delta,
// which the chat UI renders as a collapsible "thinking" section. Kept separate
// from streamChunkJSON so the answer-content callsites stay untouched.
func streamReasoningJSON(id, model, reasoningDelta string) string { panic("fake") }

// ServeChatUI serves the chat HTML page.
func (h *Handlers) ServeChatUI(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ListModels(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- events / observability ----

func (h *Handlers) ListEvents(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- health ----

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- memo (shared clipboard) ----

func (h *Handlers) GetMemos(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) AddMemo(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) DeleteMemo(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ServeMemoUI(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- memory (RAG) ----

func (h *Handlers) AddMemory(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) SearchMemory(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- tools ----

func (h *Handlers) ListTools(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ExecuteTool(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- node management (extended) ----

func (h *Handlers) ListNodeModels(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) SwapNodeModel(w http.ResponseWriter, r *http.Request) { panic("fake") }

// RestartNode queues a restart command for a worker. The command is delivered
// on the node's next heartbeat; the worker then re-detects its models and
// re-registers with the coordinator.
func (h *Handlers) RestartNode(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DispatchNode sends a JavaScript microtask to a connected WebSocket compute
// node and returns its result.
func (h *Handlers) DispatchNode(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ListWSNodes returns the IDs of nodes currently connected over WebSocket.
func (h *Handlers) ListWSNodes(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ServeNodeUI serves the browser compute-node client page.
func (h *Handlers) ServeNodeUI(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- dashboard (embedded HTML) ----

func (h *Handlers) Dashboard(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- helpers ----

func flattenMessages(msgs []ChatMessage) string { panic("fake") }

func buildResponse(id, model, content, reasoning string) ChatResponse { panic("fake") }

func newTaskID() string { panic("fake") }

func providerHintLabel(p string) string { panic("fake") }

func writeJSON(w http.ResponseWriter, code int, v any) { panic("fake") }

func (h *Handlers) agentSandbox(ctx context.Context, wsRoot string) (sandbox.Sandbox, bool, error) {
	panic("fake")
}

func (h *Handlers) waitForMCPInit(ctx context.Context, taskID string) { panic("fake") }
