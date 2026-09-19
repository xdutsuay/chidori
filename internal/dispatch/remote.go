package dispatch

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
)

// RemoteRequestTimeout bounds how long a single chat/agent request may run
// end-to-end. This is the ONE place that number lives — internal/api's chat
// and agent-run handlers (both the streaming and the packaged-app's
// non-streaming path, which is the LIVE path since SSE doesn't survive the
// Wails WKWebView) derive their own request context.WithTimeout from this
// same constant, rather than each hardcoding its own copy of the number.
// Before this, three separate "300 * time.Second" literals (two in
// internal/api/handlers.go, one on Remote's own http.Client below) all had
// to independently agree, and a real bug shipped from that: a hosted
// reasoning model (observed with an NVIDIA NIM model) that legitimately
// took longer than 300s to produce a first token failed with a raw
// "Client.Timeout exceeded while awaiting headers" instead of the intended
// context-deadline behavior, because Remote's http.Client enforced its own
// 300s ceiling independently of whatever the caller's context said — see
// invokeChatHTTP below for the actual fix (relying purely on ctx now, no
// client-level Timeout at all). 10 minutes gives real headroom for a slow
// reasoning model under load without needing another bump the next time
// one is slower still.
const RemoteRequestTimeout = 10 * time.Minute

// Remote is a directory of named OpenAI-compatible providers. Anthropic
// (via its openai-compat endpoint), OpenAI, Groq, Together, DeepInfra, and
// LM Studio all speak the same /chat/completions shape, so they share one
// adapter; the per-provider differences are baked into config (base URL,
// api-key env var, default model).
type Remote struct {
	client       *http.Client
	mu           sync.RWMutex
	def          string // name of default remote
	by           map[string]resolvedRemote
	usageHook    UsageHook
	nvidiaCool   map[string]time.Time // KMA-114: remote\x00model → cooldown until
	visionModel  string               // from ProviderConfig.VisionModel
	visionRemote string               // from ProviderConfig.VisionRemote
}

type resolvedRemote struct {
	cfg    config.RemoteConfig
	apiKey string
}

// NewRemote resolves API keys from env at construction time. Providers
// whose env var is unset are still registered — they just fail loudly at
// invoke time with a clear "missing key" error.
func NewRemote(provider config.ProviderConfig) *Remote { panic("fake") }

// No client-level Timeout — every call site (invokeChatHTTP,
// invokeChatHTTPStream, VerifyOpenAICompat's probe) already derives
// its own context.WithTimeout appropriate to what it's doing (see
// RemoteRequestTimeout above for the chat/agent path, verifyTimeout
// for the short connectivity probe). A second, independent
// client-level cap on top of that was the actual bug: it could fire
// before — or in a way inconsistent with — whatever the caller's
// context said, which is exactly what happened here.

// SetUsageHook installs the KMA-87 persistence callback (remote path has
// no registry node, so this is the only recording site for API keys).
func (r *Remote) SetUsageHook(h UsageHook) { panic("fake") }

func (r *Remote) noteUsage(ctx context.Context, res InvokeResult) { panic("fake") }

// remoteHTTPTransport is the explicit outbound transport for remote providers.
// DialContext gives a connect timeout independent of per-request ctx; response-
// header wait is left to ctx because non-stream LLM responses may not send
// headers until the model finishes (see RemoteRequestTimeout).
func remoteHTTPTransport() *http.Transport { panic("fake") }

// resolveKey prefers a direct key, falling back to the env var.
func resolveKey(rc config.RemoteConfig) string { panic("fake") }

// AddRemote registers (or replaces) a provider at runtime.
func (r *Remote) AddRemote(rc config.RemoteConfig) { panic("fake") }

// RemoveRemote deletes a provider at runtime (External API Keys "Delete").
// If it was the default, the default is cleared — Default() will fall back
// to the first remaining name, same as if none had ever been set explicitly.
func (r *Remote) RemoveRemote(name string) { panic("fake") }

// SetActiveInProvider marks name as the active key within its Provider group,
// deactivating any sibling entries that share the same cfg.Provider value
// ("store many, one active" — see RemoteConfig's doc comment). No-op error
// if name isn't registered.
func (r *Remote) SetActiveInProvider(name string) error { panic("fake") }

// No provider grouping on this entry (e.g. an older manually-added
// remote) — just mark it active without touching unrelated entries.

// SetDefault changes the default remote at runtime.
func (r *Remote) SetDefault(name string) { panic("fake") }

// SetModel changes the model of an existing remote at runtime (the "change
// model on a stored key" flow). The api layer persists the same change to the
// secrets file so it survives a restart.
func (r *Remote) SetModel(name, model string) error { panic("fake") }

// Credentials returns the base URL and resolved API key for a stored remote,
// so the api layer can list the models a *saved* key has access to (for the
// "change model" dropdown) without the key ever crossing to the client. Only
// used server-side.
func (r *Remote) Credentials(name string) (baseURL, apiKey, vendor string, ok bool) { panic("fake") }

// ProviderInfo describes a configured remote for the dashboard (no secrets).
type ProviderInfo struct {
	Name      string     `json:"name"`
	Provider  string     `json:"provider,omitempty"`
	KeyLabel  string     `json:"key_label,omitempty"`
	Model     string     `json:"model"`
	HasKey    bool       `json:"has_key"`
	Default   bool       `json:"default"`
	Active    bool       `json:"active"`
	AddedAt   time.Time  `json:"added_at,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Expired   bool       `json:"expired"`
	Tier      string     `json:"tier,omitempty"` // user-declared: "free" | "paid" | "capped" | ""
}

// Describe returns per-provider info (whether a key resolved, the model, and
// which is default) without exposing the key itself. Default reflects the
// EFFECTIVE default — the key dispatch will actually use — not just an
// explicitly pinned one, so the UI's star matches real routing even when the
// default comes from the tier-based fallback.
func (r *Remote) Describe() []ProviderInfo { panic("fake") }

// Names returns the configured provider names, sorted.
func (r *Remote) Names() []string { panic("fake") }

// Has reports whether a named provider is configured.
func (r *Remote) Has(name string) bool { panic("fake") }

// HasAny reports whether at least one remote is configured.
func (r *Remote) HasAny() bool { panic("fake") }

// Default returns the name of the default remote, or empty if none. An
// explicitly-set default always wins; with no (valid) explicit default the
// fallback is cheapest-first by the user-declared key tier — free > capped >
// paid > unspecified ("let the planner spend the free quota before the paid
// key") — with name order breaking ties so the choice stays deterministic.
// Expired keys are skipped in the fallback (resolve() would reject them
// anyway).
func (r *Remote) Default() string { panic("fake") }

func (r *Remote) defaultUnlocked() string { panic("fake") }

// tierRank orders user-declared key tiers cheapest-first for Default()'s
// fallback; unspecified/unknown tiers sort last.
func tierRank(t string) int { panic("fake") }

// Invoke calls the named remote (or the default if name is empty) with a
// non-streaming OpenAI-compatible chat request.
func (r *Remote) Invoke(ctx context.Context, name, model, prompt string) (InvokeResult, error) {
	panic("fake")
}

// InvokeWithTools is Invoke plus optional OpenAI tools[] / tool_calls parsing.
func (r *Remote) InvokeWithTools(ctx context.Context, name, model, prompt string, tools []ChatTool) (InvokeResult, error) {
	panic("fake")
}

// InvokeWithToolMessages sends an explicit messages[] history with tools (KMA-205/218).
// NVIDIA remotes fall back to flattening the last user content into InvokeWithTools.
func (r *Remote) InvokeWithToolMessages(ctx context.Context, name, model string, msgs []map[string]any, tools []ChatTool) (InvokeResult, error) {
	panic("fake")
}

// NVIDIA path does not yet accept multi-turn tool messages; flatten.

// InvokeStream calls the named remote with stream=true and feeds deltas
// through cb.
func (r *Remote) InvokeStream(ctx context.Context, name, model, prompt string, cb func(string)) (InvokeResult, error) {
	panic("fake")
}

// VerifyByName performs a live authenticated probe using the stored key for
// name, without ever exposing the key itself to the caller (only ok/detail
// cross the internal/api boundary). Deliberately bypasses resolve()'s expiry
// gate — checking whether a key that's expired *per this app's own records*
// is still accepted by the provider is exactly the kind of thing this is for
// (e.g. confirming before extending its expiry). The probe itself is a
// 1-token chat completion when a model is configured — see VerifyOpenAICompat
// for why GET /models alone can't be trusted (public on NVIDIA NIM).
func (r *Remote) VerifyByName(ctx context.Context, name string) (bool, string) { panic("fake") }

func (r *Remote) resolve(name string) (resolvedRemote, error) { panic("fake") }

// requireModel returns an error if the model that's actually about to be
// sent on the wire is empty. Checked in Invoke/InvokeStream (after the
// caller-supplied model, if any, has already been allowed to override
// rr.cfg.Model) rather than in resolve() itself, since resolve() doesn't
// know whether the caller passed an explicit model — resolve() rejecting
// unconditionally on a blank rr.cfg.Model would misfire for calls that
// supply their own model and don't need the config default at all.
//
// This exists because sending an empty "model" field to an OpenAI-compatible
// /chat/completions endpoint gets rejected by every provider we support (400
// Bad Request), which previously surfaced as an opaque "hybrid: all legs
// failed" with no indication of why the remote leg specifically lost —
// caught via a user report where "Test" (a plain GET /models reachability +
// key check, see VerifyByName) reported success but real inference still
// failed, because /models doesn't require a model param and so never
// exercises this path.
func requireModel(name, model string) error { panic("fake") }
