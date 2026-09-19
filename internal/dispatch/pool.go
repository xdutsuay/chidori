package dispatch

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/registry"
)

// Pool is the adaptive routing layer between the API handlers and the
// concrete backends (local registry-dispatched nodes vs. configured remote
// providers). It encodes one rule:
//
//   - explicit providerHint (or backend=="remote") → Remote
//   - any healthy local nodes                       → Dispatcher
//   - no nodes but a remote is configured           → Remote (default)
//   - nothing available                             → error
//
// The chain engine doesn't go through Pool — it talks to Dispatcher
// directly via the Invoker interface, which is what keeps the engine
// agnostic to where compute actually happens.
type Pool struct {
	reg     *registry.Registry
	dis     *Dispatcher
	remote  *Remote
	mu      sync.RWMutex
	backend string // mirror of config.Provider.Backend; mutable via SetBackend
}

func NewPool(reg *registry.Registry, dis *Dispatcher, remote *Remote, provider config.ProviderConfig) *Pool {
	panic("fake")
}

// Backend returns the currently active default backend.
func (p *Pool) Backend() string { panic("fake") }

// SetBackend switches the default backend at runtime. Accepts "ollama",
// "lmstudio", or "remote". Unknown values are rejected.
func (p *Pool) SetBackend(backend string) error { panic("fake") }

// Execute runs a single-shot inference. providerHint:
//   - "" → use the configured default (local if available, remote otherwise)
//   - "local" / "ollama" / "lmstudio" → force registry-based local routing
//   - "chidori-nagasa" / "nagasa" / "phone" → phone worker only
//   - any other value → look up a remote by that name
func (p *Pool) Execute(ctx context.Context, model, prompt, providerHint string) (InvokeResult, error) {
	panic("fake")
}

// ExecuteWithTools is Execute with optional OpenAI tools (remote path only).
// When tools is non-empty and the route is local, tools are ignored and the
// plain local path runs (agent dual protocol uses JSON for local).
func (p *Pool) ExecuteWithTools(ctx context.Context, model, prompt, providerHint string, tools []ChatTool) (InvokeResult, error) {
	panic("fake")
}

// ExecuteWithToolMessages is ExecuteWithTools for an explicit messages[] history.
func (p *Pool) ExecuteWithToolMessages(ctx context.Context, model string, msgs []map[string]any, providerHint string, tools []ChatTool) (InvokeResult, error) {
	panic("fake")
}

// Local path has no native tools[] history; flatten last user content.

// RemoteNames returns configured remote provider names (for tool-protocol auto).
func (p *Pool) RemoteNames() []string { panic("fake") }

// PreferredToolsRemote returns the best remote name for OpenAI tools[] calls
// when the route would otherwise be hybrid/local (KMA-203). Prefers OpenRouter
// by name/URL; otherwise the default remote; empty if none configured.
func (p *Pool) PreferredToolsRemote() string { panic("fake") }

// LegResult is one leg of a hybrid execution (local or remote).
type LegResult struct {
	Provider   string  `json:"provider"`
	NodeID     string  `json:"node_id,omitempty"`
	Model      string  `json:"model,omitempty"`
	DurationMs float64 `json:"duration_ms"`
	OK         bool    `json:"ok"`
	Error      string  `json:"error,omitempty"`
}

// ExecuteHybrid runs the request on the local cluster AND the default remote
// provider concurrently (redundant execution), returning the FIRST successful
// result. The slower leg is cancelled and abandoned — it used to be awaited
// so both timings could be compared, but that meant the user waited for the
// slowest leg: a dead-slow reasoning model on one leg stalled an answer the
// other leg had produced in seconds (and the stall could outlive the HTTP
// connection entirely, so nothing ever reached the UI). If one leg has
// nothing to run (no nodes / no remote) it simply fails and the other wins.
// The returned legs contain only what actually resolved before the winner.
//
// Difficulty-aware routing (§13.2 #1 in llpcodefeature_Cursorclone.md): an
// "easy" prompt with local capacity available skips the remote leg entirely
// via tryLocalFirst below — racing both for "hi" or "what's 2+2" is a pure
// wasted remote call. Anything classified "hard", or with no local capacity,
// falls through to the original always-race-both behavior unchanged.
func (p *Pool) ExecuteHybrid(ctx context.Context, model, prompt string) (InvokeResult, []LegResult, error) {
	panic("fake")
}

// Local leg failed despite looking healthy (e.g. model unloaded
// mid-request) — fall back to the normal race-both path below rather
// than surfacing an error the remote leg could have avoided.

// aborts the losing leg's in-flight HTTP call

// Empty model → the remote uses its own configured model (e.g.
// NVIDIA's), rather than inheriting the local model name.

// Both legs failed. Include both legs' actual errors rather than a bare
// "all legs failed" — that message alone gave no way to tell "both
// backends are genuinely down" apart from "one backend is misconfigured"
// (e.g. a remote key saved with no model — see requireModel() in
// internal/dispatch/remote.go) without digging into server logs.

// tryLocalFirst runs only the local leg, for prompts classifyDifficulty has
// already judged easy. ok is false on failure, signaling ExecuteHybrid to
// fall back to its normal race-both path rather than surfacing this error
// directly — an easy-looking prompt is still entitled to the same reliability
// as any other request.
func (p *Pool) tryLocalFirst(ctx context.Context, model, prompt string) (InvokeResult, []LegResult, bool) {
	panic("fake")
}

// ExecuteStream is the streaming variant of Execute.
func (p *Pool) ExecuteStream(ctx context.Context, model, prompt, providerHint string, cb func(string)) (InvokeResult, error) {
	panic("fake")
}

// Remote exposes the underlying Remote for handlers that need to list
// providers or check configuration.
func (p *Pool) Remote() *Remote { panic("fake") }

func (p *Pool) shouldUseRemote(hint string) bool { panic("fake") }

// No explicit hint: honour configured backend first.

// Prefer local if any healthy nodes exist.

// No nodes — fall back to remote if configured.

func remoteName(hint string) string { panic("fake") }

// resolve() will pick the default

func isLocalHint(hint string) bool { panic("fake") }

// "local (auto)" is the topbar dropdown's DISPLAY LABEL for its local
// option (whose real <option value=""> is empty — see loadProviders in
// public/js/09-agent-ask.js), not a hint any caller should be
// constructing on purpose.
// Recognized here defensively anyway (running_issue.md P0-1 / H.2): a
// caller that accidentally sends the label instead of the value —
// exactly what ADR-0008's draft leg did — used to fail routing with
// "unknown remote provider \"local (auto)\"" instead of just routing
// local, which is unambiguously what was meant. The real fix is at the
// call site (send "local", not the label); this is the safety net so
// the same mistake can't silently break routing again.

// localBackendFilter maps an explicit local hint to a registry Backend.
// Empty means any healthy local node (Ollama and phone compete on load).
func localBackendFilter(hint string) string { panic("fake") }

// ErrNoBackend is returned by Execute when neither a healthy node nor a
// remote provider is available.
func ErrNoBackend() error { panic("fake") }

// InvokeVision runs multimodal vision for Ask/Agent image attachments (KMA-117).
// When the configured backend is local and healthy nodes exist, local OpenAI-
// compatible backends are tried first; otherwise a provider-appropriate remote
// is used (NVIDIA preferred when configured).
func (p *Pool) InvokeVision(ctx context.Context, model, prompt string, images []ImagePart) (InvokeResult, error) {
	panic("fake")
}
