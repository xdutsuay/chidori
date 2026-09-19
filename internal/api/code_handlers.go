package api

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/activity"
	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/compact"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/flightlog"
	"github.com/xdutsuay/lclreason/internal/harness"
	"github.com/xdutsuay/lclreason/internal/hooks"
	"github.com/xdutsuay/lclreason/internal/prompts"
	"github.com/xdutsuay/lclreason/internal/prompts/modules"
	"github.com/xdutsuay/lclreason/internal/sandbox"
	"github.com/xdutsuay/lclreason/internal/store"
	"github.com/xdutsuay/lclreason/internal/tools"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// CodeCompletionRequest is the body for POST /v1/code/completions.
type CodeCompletionRequest struct {
	Path      string `json:"path"`
	Language  string `json:"language"`
	Prefix    string `json:"prefix"`
	Suffix    string `json:"suffix"`
	MaxTokens int    `json:"max_tokens"`
	Provider  string `json:"provider"`
	Model     string `json:"model"`
}

// CodeCompletionResponse is returned for tab completions.
type CodeCompletionResponse struct {
	Completion string  `json:"completion"`
	Mode       string  `json:"mode"` // fim | prefix
	Cached     bool    `json:"cached,omitempty"`
	DurationMs float64 `json:"duration_ms,omitempty"`
}

// codeStreamTTFT returns the configured stream time-to-first-token budget.
func codeStreamTTFT(cfg config.CodeConfig) time.Duration { panic("fake") }

// codeCompletionTimeout bounds tab-completion requests. Shorter than
// RemoteRequestTimeout (10 min): a hung completion must not block every
// keystroke's goroutine until the editor disconnects (H4).
const codeCompletionTimeout = 60 * time.Second

// CodeCompletions handles inline code completion (FIM with prefix fallback).
func (h *Handlers) CodeCompletions(w http.ResponseWriter, r *http.Request) { panic("fake") }

func completionCacheKey(path, model, prefix, suffix string) string { panic("fake") }

func truncateCompletion(s string, maxTokens int) string {
	panic(
		// Rough char budget from token estimate.
		"fake")
}

// Cut at newline boundary if possible.

// AgentRunRequest is the body for POST /v1/agent/run.
type AgentRunRequest struct {
	Task        string `json:"task"`
	OpenContext string `json:"open_context"`
	Model       string `json:"model"`
	Provider    string `json:"provider"`
	Harness     string `json:"harness,omitempty"` // ""|"chidori" (default) | "grok" experimental
	Stream      bool   `json:"stream"`
	ClientID    string `json:"client_id,omitempty"`
	// TaskID is optional. When set (packaged-app live progress, ADR-0017
	// Decision 3), activity bus events are keyed under this id so the client
	// can poll GET /api/activity?request_id=<id> (or ?task=) mid-run. When
	// empty, the server generates one and returns it as X-Task-ID.
	TaskID    string `json:"task_id,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	ApplyMode string `json:"apply_mode,omitempty"` // auto_apply | propose
	Mode      string `json:"mode,omitempty"`       // "" (agent, default) | "plan" | "debug" — see modeInstructionsFor
	// Background (AG.16) detaches the run from the HTTP request: returns 202
	// with a job id immediately and runs under a cancelable registry context.
	Background bool `json:"background,omitempty"`
	// AllowTruncate opts into dispatcher clamp when the assembled agent
	// prompt exceeds the model window (KMA-5). Default is fail-closed.
	AllowTruncate bool `json:"allow_truncate,omitempty"`
}

// modeInstructionsFor returns the extra prompt block for Plan/Debug/Agent
// modes. Agent (mode "") gets a short clarification-fence nudge; Plan/Debug
// get full mode instructions plus structured fence preferences for the UI.
func modeInstructionsFor(mode string) string { panic("fake") }

// Ordinary Agent mode uses Mode "" — nudge clarification cards when blocking on the user.

// AgentRun executes the multi-turn coding agent (SSE when stream=true).
func (h *Handlers) AgentRun(w http.ResponseWriter, r *http.Request) { panic("fake") }

// AgentRunWithBody is the shared entry for POST /v1/agent/run and POST /api/agent/jobs.
func (h *Handlers) AgentRunWithBody(w http.ResponseWriter, r *http.Request, req AgentRunRequest) {
	panic("fake")
}

// Custom instructions — global (ADR-0011) + project (.lclreason/rules,
// .cursorrules) — folded into the agent's open_context alongside
// whatever the editor attached.

// Prompt profiles inject via ModeInstructions (stable system), not OpenContext
// (ambient) — see withPromptProfileInstructions / KMA-214.

// Match Ask's budget (UNIFIED H.2): Agent/Plan/Debug used bare r.Context()
// so a hung remote could outlive any sensible desktop wait unless the user
// hit Stop. RemoteRequestTimeout is the shared ceiling for slow hosted
// reasoning models; client Abort still cancels sooner via r.Context().
// POST /api/runs/{id}/cancel covers packaged WKWebView where abort may
// not cancel r.Context() (KMA-6).

// Clamp @file / open_context dumps before the agent loop — never 400.

// Surface early so packaged clients can poll the activity bus mid-run
// before the JSON body arrives (ADR-0017 Decision 3).

// Workspace trust gate: regardless of the requested apply mode, an
// untrusted workspace (one the user hasn't confirmed via the Open
// Folder trust prompt) never gets auto-applied writes — edits are
// always proposed for review. Enforced server-side so a stale/buggy
// client can't bypass it. Also threaded through as Trusted below: it
// separately gates the agent's git_commit/shell_exec tools (guard.go),
// which is a stricter, non-overridable "no" rather than propose-vs-apply.

// Plan mode always proposes rather than auto-applies, regardless of the
// workspace's normal apply mode or trust status — "read-only until you
// approve" is the whole point of the mode, not just its default.

func (h *Handlers) runExternalHarnessViaRegistry(w http.ResponseWriter, r *http.Request, ctx context.Context, cancel context.CancelFunc, req AgentRunRequest, taskID string, em activity.Emitter, proposeOnly, trusted bool, extHarness harness.ExternalHarness) {
	panic("fake")
}

func (h *Handlers) agentRunNative(w http.ResponseWriter, r *http.Request, ctx context.Context, cancel context.CancelFunc, req AgentRunRequest, taskID string, em activity.Emitter, proposeOnly, trusted bool, wsRoot string) {
	panic("fake")
}

// ADR-0013: reload hooks.json each tool call so edits take effect
// mid-run without a coordinator restart (same fresh-read pattern as
// fireHook for on_file_save).

// Collect every TurnEvent into a trace and return it (ADR-0017
// Decision 2, running_issue.md P0-3): this used to be a no-op
// callback — `func(agent.TurnEvent) {}` — which discarded
// tool_result/error events server-side, so the packaged app (the
// only client of this branch; WKWebView can't stream SSE through
// Wails) rendered a run's final message with no way to show which
// tools ran or failed. Size is bounded: SafetyMaxTurns × tools-per-
// turn × 2000-char truncated outputs. The frontend renders `trace`
// through the same renderTurnEvent the SSE path uses (one renderer,
// two feeders), so the two transports can't drift in what they show.

func withSessionID(fields map[string]any, sessionID string) map[string]any { panic("fake") }

func marshalAgentTraceJSON(trace []agent.TurnEvent) string { panic("fake") }

func (h *Handlers) persistAgentTask(ctx context.Context, t store.Task) { panic("fake") }

// poolUsageJSON returns {"prompt_tokens":N,"completion_tokens":M} from a
// PoolInvoker after an agent run (Wave D).
func poolUsageJSON(inv *agent.PoolInvoker) map[string]int { panic("fake") }

func (h *Handlers) poolRemoteNames() []string { panic("fake") }

func (h *Handlers) agentToolSchemas() []dispatch.ChatTool { panic("fake") }

func (h *Handlers) agentToolsList() []tools.Tool { panic("fake") }
