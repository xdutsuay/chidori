package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/xdutsuay/lclreason/internal/activity"
	"github.com/xdutsuay/lclreason/internal/compact"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/harnesstrace"
	"github.com/xdutsuay/lclreason/internal/hooks"
	"github.com/xdutsuay/lclreason/internal/sandbox"
	"github.com/xdutsuay/lclreason/internal/tools"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// DefaultStreamTTFT is how long streaming waits for the first content delta
// before cancelling the stream sub-context and falling back to non-stream
// Invoke (RT-THPUT-4 / harness Invariant 2).
const DefaultStreamTTFT = 60 * time.Second

// Invoker runs LLM calls for planning and final answers.
type Invoker interface {
	Invoke(ctx context.Context, model, prompt, provider string) (string, error)
}

// StreamInvoker is an optional Invoker extension. When present, Loop uses it
// for mid-turn token batches (Wave A / RT-THPUT-2) so high tok/s providers
// surface progress before the full JSON plan arrives.
type StreamInvoker interface {
	Invoker
	InvokeStream(ctx context.Context, model, prompt, provider string, onDelta func(string)) (string, error)
}

// ToolRunner executes registered tools.
type ToolRunner interface {
	Execute(ctx context.Context, name string, params map[string]any) (*tools.ToolResult, error)
}

// ProposedEdit is a file change the agent wants to apply.
type ProposedEdit struct {
	ID      string `json:"id,omitempty"`
	Path    string `json:"path"`
	Kind    string `json:"kind"` // write_file | apply_patch
	Summary string `json:"summary,omitempty"`
	Before  string `json:"before,omitempty"`
	After   string `json:"after,omitempty"`
}

// TurnEvent is emitted during agent execution for SSE streaming.
type TurnEvent struct {
	Type     string        `json:"type"` // plan | thought | tool_start | tool_result | edit_proposed | clarify | message | done | error
	Turn     int           `json:"turn,omitempty"`
	Tool     string        `json:"tool,omitempty"`
	Output   string        `json:"output,omitempty"`
	Message  string        `json:"message,omitempty"`
	Provider string        `json:"provider,omitempty"`
	Edit     *ProposedEdit `json:"edit,omitempty"`
	Done     bool          `json:"done,omitempty"`
	// LimitReached marks a "done" event produced by hitting the turn budget
	// (see wrapUpOnLimit) rather than the model genuinely finishing —
	// distinct so the frontend can offer a "Continue" affordance instead of
	// rendering it as an ordinary final answer.
	LimitReached bool `json:"limit_reached,omitempty"`
	// Questions holds the JSON array body from a chidori-questions fence when
	// Type is "clarify" (H.3 — pause loop until the user answers).
	Questions string `json:"questions,omitempty"`
	// PromptTokens/CompletionTokens are set on "done" (and limit-reached done)
	// when the invoker is a *PoolInvoker — cumulative for the whole run.
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	// CacheHitTokens is the cumulative provider cache-read token count
	// (prompt_tokens_details.cached_tokens) surfaced on "done" events
	// when the invoker is a *PoolInvoker (KMA-163 slice 2).
	CacheHitTokens int `json:"cache_hit_tokens,omitempty"`
	// GenerationID is the latest provider completion id when known (KMA-211).
	GenerationID string `json:"generation_id,omitempty"`
	// PromptStableSnapshot / PromptVolatileSnapshot carry the wire prompt parts
	// emitted on inference harness model_request events (KMA-215).
	PromptStableSnapshot   string `json:"prompt_stable_snapshot,omitempty"`
	PromptVolatileSnapshot string `json:"prompt_volatile_snapshot,omitempty"`
}

// Config tunes the agent loop.
type Config struct {
	SafetyMaxTurns int
	ContextTokens  int
	ReserveTokens  int
	WorkspaceRoot  string
	ProposeOnly    bool
	Trusted        bool // gates git_commit/shell_exec in guardedRunner — see guard.go
	// ExperimentalBrowserTools advertises browse_* tools in the agent prompt
	// when true (and Trusted). The tools themselves are registered on the
	// shared registry only when ui.experimental_browser_tools is enabled.
	ExperimentalBrowserTools bool
	Files                    FileReader
	// Hooks, when set, reloads .lclreason/hooks.json for each tool call and
	// runs ADR-0013 pre_tool_call / post_tool_call around guardedRunner.Execute.
	// Nil keeps legacy behavior (on_file_save was the only wired event).
	Hooks func() *hooks.Registry

	// ModeInstructions is prepended to the prompt, after the fixed tool-call
	// JSON contract and before the task itself, when the chat UI's mode is
	// Plan or Debug rather than plain Agent (see internal/api/code_handlers.go
	// AgentRun and setMode() in public/js/01-core.js). Empty for ordinary
	// Agent mode — this
	// is deliberately just an extra instruction block, not a different tool
	// contract, so it can't break the existing JSON parsing in parsePlan.
	ModeInstructions string

	// Mode is the chat UI mode ("" agent, plan, debug) stored in checkpoints.
	Mode string

	// OwnedPaths lists workspace-relative paths already claimed by a parent
	// agent's pending ProposedEdit list (ADR-0019). Nested write subagents
	// must not propose for these paths. Nil/empty means no parent claims.
	OwnedPaths map[string]bool

	// Leases is an optional file lease manager (ADR-0019 / AI.26 / KMA-140).
	// When set, write_file/apply_patch acquire a process-wide lease so
	// concurrent runs cannot silently stomp the same path.
	Leases *LeaseManager
	// LeaseOwner identifies this run in Leases (task/job id). Empty uses "agent-run".
	LeaseOwner string
	// PriorScratch is tool-result text persisted from a previous turn in this
	// session (KMA-143). Seeded into observations so follow-ups see prior hits.
	PriorScratch string
	// PersistScratch, when set, is invoked with compacted observations when
	// the loop returns so the next user turn can reload them.
	PersistScratch func(scratch string)

	// Flight is an optional NDJSON recorder (UNIFIED_ROADMAP Wave 0). Nil is fine.
	Flight FlightRecorder
	// SessionID / RequestID are KMA-219 join keys mirrored onto flight turn_event
	// lines (session ↔ task ↔ flight). RequestID equals the coordinator task id.
	SessionID string
	RequestID string

	// StreamTTFT bounds time-to-first-token for streaming invocations (0 =
	// DefaultStreamTTFT).
	StreamTTFT time.Duration
	// StreamIdle bounds silence after the first token (0 = disabled).
	StreamIdle time.Duration

	// ToolProtocol is auto|openai|json (see config.ResolveToolProtocol).
	ToolProtocol string
	// RemoteNames lists configured remote provider names for auto protocol.
	RemoteNames []string
	// ToolSchemas are OpenAI-compatible tool defs from the registry (filtered
	// per turn by trust/browser flags).
	ToolSchemas []dispatch.ChatTool

	// ToolCatalog is the live registry catalog (tools.Registry.Describe()).
	// When set, the JSON-protocol prompt derives its advertised tool list —
	// names AND their param docs — from this single source of truth instead
	// of a hand-maintained string, so registered-but-unadvertised (or
	// advertised-but-missing) drift can't recur. Empty falls back to the
	// legacy hardcoded list.
	ToolCatalog []tools.Tool

	// Sandbox controls workspace snapshot and rollback lifecycle (SBX-003).
	Sandbox sandbox.Sandbox
	// AutoRollbackOnError automatically rolls back sandbox state if turn fails.
	AutoRollbackOnError bool
	// Dispatch manages session concurrency and cancellation (HRD-001).
	Dispatch *SessionDispatch
	// Tracker records file reads and writes per session (HRD-001).
	Tracker *workspace.FileTracker

	// DelegateDepth is how many nested delegate_subtask frames deep this Loop
	// already is (0 = top-level agent). Nested Loops inherit parent+1.
	DelegateDepth int
	// MaxDelegateDepth caps nesting (HARNESS_AUDIT H7). 0 means default of 1 —
	// the parent may spawn one subagent level; that subagent may not re-delegate.
	// Set negative to disable the cap (tests/experiments only).
	MaxDelegateDepth int

	// SubagentModels is the allowlist of delegate_subtask params.model values
	// (KMA-107). Empty or only "inherit" means subagents use the parent's
	// model and provider. Named entries must resolve via ResolveSubagentModel.
	SubagentModels []string
	// SubagentMaxInflight caps concurrent async delegate_subtask goroutines.
	// 0 means default 2.
	SubagentMaxInflight int
	// ResolveSubagentModel maps an allowlisted named model ID to provider+model.
	// Fail closed when ok is false.
	ResolveSubagentModel func(name string) (provider, model string, ok bool)
}

// FlightRecorder is the subset of flightlog.Logger the agent loop needs.
type FlightRecorder interface {
	Info(comp, op string, fields map[string]any)
	Error(comp, op string, fields map[string]any)
	Debug(comp, op string, fields map[string]any)
}

// Loop runs a capped multi-turn agent with JSON tool plans.
type Loop struct {
	cfg       Config
	invoker   Invoker
	tools     ToolRunner
	memRecall func(ctx context.Context, query string) []string
}

// NewLoop creates an agent loop.
func NewLoop(cfg Config, invoker Invoker, toolReg ToolRunner, memRecall func(ctx context.Context, query string) []string) *Loop {
	panic("fake")
}

type planResponse struct {
	// Thought is the model's own brief reasoning for this turn — required by
	// the prompt (AI.25, "industry-standard system prompt & CoT", matching
	// Aider/Cline convention: think before acting, not after). Optional here
	// (omitempty / no `json:"...,required"` — encoding/json has no such tag)
	// so a model that ignores the instruction and omits it doesn't fail
	// parsePlan outright; Run just has nothing to surface as a "thought"
	// event that turn instead of erroring.
	Thought string `json:"thought,omitempty"`
	Done    bool   `json:"done"`
	Summary string `json:"summary,omitempty"`
	Tools   []struct {
		Name   string         `json:"name"`
		Params map[string]any `json:"params"`
	} `json:"tools"`
}

// busActivityFor maps a TurnEvent to its activity-bus representation.
// ok=false means the event type has no bus mirror ("message" — it is always
// immediately followed by a "done" carrying the same content, so mirroring
// both would double-record it).
//
// This mapping exists because of running_issue.md P0-3 (ADR-0017): the SSE
// emit callback used to be the ONLY carrier for tool_result/error events, so
// any client that couldn't stream (the packaged app's WKWebView) structurally
// could not see tool failures — they were discarded server-side. The activity
// bus is the transport-independent source of truth; every TurnEvent goes
// through here so no future event type can accidentally become SSE-only.
// The "planning"/"tool_call" labels reproduce exactly what the loop used to
// emit inline, so the Planner panel's existing feed is unchanged.
func busActivityFor(ev TurnEvent) (activity.Activity, bool) { panic("fake") }

// AI.25: the model's own CoT reasoning for the turn, surfaced in the
// same "planning" feed as the "plan" event just above (reusing the
// existing, already-wired kind rather than inventing a new one the
// frontend/activity-bus consumers would need to special-case).

// Emitted immediately before invoker.Invoke so Diagnostics can show
// "waiting on model" between planning and either tools or error
// (API failures otherwise looked like "planned once and died").
// Provider/Output carry resolved route (CFG.28): hint→local|remote|hybrid.

// Mid-turn streamed content batches (Wave A). Packaged poll maps these
// back to token TurnEvents without inventing a second progress UI.

// BusActivityFor exposes the transport-parity event-bus mapping for other
// harnesses (e.g. external ACP runners) that need to emit the same
// activity-bus semantics as the native loop.
func BusActivityFor(ev TurnEvent) (activity.Activity, bool) { panic("fake") }

// Run executes the agent loop, calling emit for each event. When the model
// emits a chidori-questions fence, clarify is non-nil and the loop pauses
// without a terminal "done" event (H.3).
func (l *Loop) Run(ctx context.Context, task, openContext, model, provider string, em activity.Emitter, emit func(TurnEvent)) (string, []ProposedEdit, *ClarifyPause, error) {
	panic("fake")
}

// RunResume continues a paused run after the user submits clarification answers.
func (l *Loop) RunResume(ctx context.Context, cp Checkpoint, answers string, em activity.Emitter, emit func(TurnEvent)) (string, []ProposedEdit, *ClarifyPause, error) {
	panic("fake")
}

type loopState struct {
	observations   []string
	edits          []ProposedEdit
	turn           int
	parseFailCount int
}

func (l *Loop) run(ctx context.Context, task, openContext, model, provider string, em activity.Emitter, emit func(TurnEvent), st *loopState) (summary string, edits []ProposedEdit, clarify *ClarifyPause, err error) {
	panic("fake")
}

// KMA-140: drop this run's path leases when the loop returns so the
// next concurrent job/session can claim them. Keep leases during an
// H.3 clarify pause — RunResume continues under the same LeaseOwner.

// Every event goes to BOTH carriers: the per-transport emit callback
// (SSE frames when streaming, the trace collector when not — see
// AgentRun in internal/api/code_handlers.go) AND the activity bus via
// em (nil-safe). Wrapping emit here, at the single chokepoint, is what
// makes it impossible for an individual emit site below to reach one
// carrier and not the other — the exact drift that caused P0-3.
// Wave 0 also mirrors a compact subset onto the flight recorder.

// parseFailCount tracks consecutive parsePlan failures across turns

// KMA-163 incremental transport: cache the stable prompt prefix within this Run.

// KMA-205/218: native OpenAI path accumulates assistant.tool_calls + role:tool.

// Graceful wrap-up instead of a hard stop (owner-specified UX
// improvement, 2026-07-09 — the OpenCode pattern: force one
// last text-only response summarizing progress instead of just
// truncating a long task with an unhelpful "Agent stopped
// (safety limit)."). LimitReached lets the frontend offer a
// "Continue" affordance rather than treating this like a
// normal finished answer.

// Resolve route before building the prompt so openai vs json shapes match the backend.

// KMA-163: reuse a byte-identical stable prefix across turns when the
// route/tool shape has not changed (enables OpenAI-compat system-message
// prompt cache). Volatile still carries Turn N + scratch.

// OpenAI path: no JSON self-correction loop — surface cleanly.

// Incomplete local-model plans: thought-only / empty {} / truncated JSON
// recovered as thought — do not fake "Task completed" (KMA local recovery).

// subagentAllowedTools is the tool set permitted inside a delegated subtask
// (see runDelegatedSubtask). Read/search/calc tools plus propose-only
// write_file/apply_patch (ADR-0019). Still excludes shell_exec/git_commit/
// browse_*/memory_add — a subagent must never execute code or auto-apply.
var subagentAllowedTools = map[string]bool{
	"read_file": true, "grep_workspace": true, "list_dir": true,
	"memory_search": true, "web_search": true, "fetch_url": true,
	"calculator": true, "echo": true,
	"write_file": true, "apply_patch": true,
	"search_chat_messages": true,
	"get_chat_attachment":  true,
}

// subagentReadOnlyTools is the legacy read-only subset retained for tests that
// assert the pre-ADR-0019 boundary. Prefer subagentAllowedTools for new code.
var subagentReadOnlyTools = map[string]bool{
	"read_file": true, "grep_workspace": true, "list_dir": true,
	"memory_search": true, "web_search": true, "fetch_url": true,
	"calculator": true, "echo": true,
}

// readOnlyToolRunner rejects any tool call outside subagentReadOnlyTools
// before it ever reaches the real registry — kept for unit tests of the
// hard read-only boundary.
type readOnlyToolRunner struct{ inner ToolRunner }

func (r readOnlyToolRunner) Execute(ctx context.Context, name string, params map[string]any) (*tools.ToolResult, error) {
	panic("fake")
}

func (r writeProposeToolRunner) Execute(ctx context.Context, name string, params map[string]any) (*tools.ToolResult, error) {
	panic("fake")
}

// writeProposeToolRunner is the ADR-0019 subagent allowlist: read tools plus
// write_file/apply_patch. Writes still hit the nested Loop's ProposeOnly
// guardedRunner (never the live FS apply path). Code-execution tools stay blocked.
type writeProposeToolRunner struct{ inner ToolRunner }

// ownedPathsFromEdits builds the ADR-0019 parent ownership map from pending proposals.
func ownedPathsFromEdits(edits []ProposedEdit) map[string]bool { panic("fake") }

// asyncSubResult stores results from async subagent runs.
type asyncSubResult struct {
	Summary string
	Edits   []ProposedEdit
	Err     error
	Done    bool
	Write   bool
	// Claimed is set by ClaimAsyncJob after the first done-poll hands the
	// edits to a parent loop, so a repeated subagent_result call can't merge
	// the same proposals twice.
	Claimed bool
}

// asyncJobs stores results of async delegate_subtask goroutines, keyed by job id.
var asyncJobs sync.Map // map[string]*asyncSubResult

// asyncClaimMu serializes ClaimAsyncJob's read-modify-write so two concurrent
// polls of the same job id can't both claim its edits.
var asyncClaimMu sync.Mutex

// AsyncJobExists reports whether job_id was returned by delegate_subtask with async:true.
func AsyncJobExists(id string) bool { panic("fake") }

// LookupAsyncJob retrieves an async subagent result by job id (read-only —
// the subagent_result tool path uses ClaimAsyncJob so edits merge exactly once).
func LookupAsyncJob(id string) (summary string, edits []ProposedEdit, done bool, write bool) {
	panic("fake")
}

// ClaimAsyncJob retrieves an async subagent result and, when the job is done,
// atomically claims its proposed edits: the first claimer gets them, every
// later call gets the summary with no edits. found=false means no such job id.
func ClaimAsyncJob(id string) (summary string, edits []ProposedEdit, done, alreadyClaimed, found bool) {
	panic("fake")
}

// pollAsyncSubagent implements the subagent_result tool: it reports an async
// delegate_subtask job's status and, on the first done-poll, returns its
// proposed edits so the parent loop merges them (exactly once).
func pollAsyncSubagent(params map[string]any) (*tools.ToolResult, []ProposedEdit) { panic("fake") }

// searchTools implements tool_search (KMA-86 deferred discovery): BM25 + lexical
// hybrid over the full ToolCatalog / ToolSchemas (name, description, param docs),
// not only the short advertised set. API contract unchanged for callers.
// tests: KMA-100 searchTools BM25 (delegate)
func (l *Loop) searchTools(params map[string]any) *tools.ToolResult { panic("fake") }

// leaseCheckedToolRunner wraps a ToolRunner to acquire leases before write_file/apply_patch.
type leaseCheckedToolRunner struct {
	inner   ToolRunner
	leases  *LeaseManager
	ownerID string
}

func (lc leaseCheckedToolRunner) Execute(ctx context.Context, name string, params map[string]any) (*tools.ToolResult, error) {
	panic("fake")
}

// runDelegatedSubtask spawns a nested Loop for a bounded sub-question.
// Supports optional params:
//   - "write": true — use lease-checked write-capable tool runner (ADR-0019 AI.26)
//   - "async": true — run in a goroutine, return job id immediately (AI.26)
//
// Sync mode (default): runs inline, merges proposals into parent.
// Async mode: fires goroutine, returns job id; parent polls via subagent_result.
// Write tools are always propose-only; proposals merge into parent's pending list.
// Trusted is always false — no shell_exec/git_commit regardless of parent trust.
//
// Nesting is capped by Config.MaxDelegateDepth (default 1): a subagent may not
// call delegate_subtask again (HARNESS_AUDIT H7). The switch that dispatches
// delegate_subtask lives in Loop.run itself, so the ToolRunner allowlist alone
// cannot stop recursion — this gate is the hard stop.
func (l *Loop) runDelegatedSubtask(ctx context.Context, model, provider string, params map[string]any, parentEdits []ProposedEdit) (*tools.ToolResult, []ProposedEdit, error) {
	panic("fake")
}

// wrapUpOnLimit runs once, when the turn budget is exhausted, instead of
// just returning a flat "stopped" message. It makes ONE additional invoker
// call with no tool schema at all — the model is asked directly for a plain
// prose progress report (what's done, what's left, recommended next steps)
// instead of another {"done":...} tool plan. If that call can't run (the
// context is already done) or fails/returns nothing, falls back to a plain
// listing of the raw tool observations so progress narration never comes
// back empty just because the wrap-up call itself hit an error.
func (l *Loop) wrapUpOnLimit(ctx context.Context, task, openContext, model, provider string, observations []string) string {
	panic("fake")
}

// buildWrapUpPrompt asks for plain prose, not the usual JSON tool-plan
// schema — deliberately a different contract from buildAgentPrompt so the
// model can't slip back into proposing more tool calls at the exact moment
// it's being told to stop.
func buildWrapUpPrompt(task, openContext, scratch string, maxTurns int) string { panic("fake") }

// fallbackWrapUp covers the wrap-up call itself failing (or being skipped
// because the context is already done) — still better than silently losing
// all progress narration and reverting to the old flat "stopped" message.
func fallbackWrapUp(maxTurns int, scratch string) string { panic("fake") }

// EstimateSystemPromptTokens estimates the token cost of the agent's static
// system prompt — the fixed rules/tool-list text buildAgentPrompt always
// emits regardless of task content — for context-window breakdown reporting
// (AI.27). Built by calling the real prompt-building function with empty
// task-specific fields, rather than duplicating its text as a second copy
// that could drift out of sync with what's actually sent to the model.
func EstimateSystemPromptTokens(trusted, browserTools bool) int { panic("fake") }

// agentJSONPlanSchemaExample is the one-line JSON contract shown in prompts and
// parse-failure feedback. Uses list_dir "." so models do not copy a fake path
// like web/code.html from an unrelated repo layout.
const agentJSONPlanSchemaExample = `{"thought":"brief reasoning about what to do next","done":false,"summary":"","tools":[{"name":"list_dir","params":{"path":"."}}]}`

func buildAgentPrompt(workspaceRoot, modeInstructions, task, openContext, rag, scratch string, turn int, trusted, browserTools, nativeTools bool, catalog []tools.Tool) string {
	panic("fake")
}

// promptProfileOpenContextPrefix is the marker code_handlers used to prepend
// Settings profiles into OpenContext before KMA-214. buildAgentPromptParts
// peels a leading block into ModeInstructions (stable) so residual callers
// still land the profile in the system channel, not ambient.
const promptProfileOpenContextPrefix = "Prompt profile:\n\n"

// peelPromptProfileFromOpenContext lifts a leading "Prompt profile:" block out
// of openContext into a ModeInstructions fragment (KMA-214).
func peelPromptProfileFromOpenContext(openContext string) (profileBlock, rest string) { panic("fake") }

func buildAgentPromptParts(workspaceRoot, modeInstructions, task, openContext, rag, scratch string, turn int, trusted, browserTools, nativeTools bool, catalog []tools.Tool) agentPromptParts {
	panic(
		// KMA-214: profiles that still arrive via OpenContext (pre-fix callers)
		// must land in the stable ModeInstructions channel, not ambient.
		"fake")
}

// advertised is the registry-derived tool list (single source of truth);
// nil means no catalog was provided and the legacy hardcoded list is used.

// One line per tool, name + description (descriptions carry the
// param docs) — local/JSON-protocol models get the same tool
// knowledge native function-calling models get from tools[].

// git_commit/shell_exec run arbitrary code (a git hook, a shell command),
// so they're only advertised once the workspace is trusted — the same
// gate guardedRunner enforces server-side (guard.go). Not listing them
// otherwise avoids a "tool exists but every call fails" loop.

// Volatile: turn counter + growing scratch (+ open/rag/task). Keeping Turn N
// out of the stable prefix is what makes provider prompt-cache hits possible
// across agent turns (KMA-163 / AGENT_LOOP_EFFICIENCY gap 1).

func parsePlan(raw string) (planResponse, error) { panic("fake") }

// Support XML <tool_call> blocks (Hermes format)

// Handle NDJSON (newline-delimited JSON) from streaming planner responses.
// Each line is a plan update; the last one (with "done":true) is the final state.
// If there are multiple JSON objects, take the last one that parses successfully.

// Fallback: try to extract JSON from the raw text (handle wrapped/formatted responses).

// Local/weak models often truncate mid-object after streaming a useful
// "thought". Recover that so the loop can nudge instead of hard-failing.

// recoverPartialPlan extracts a usable thought (and optional done/summary/tools
// fragments) from truncated or almost-JSON planner output.
func recoverPartialPlan(raw string) (planResponse, bool) { panic("fake") }

func extractJSONStringField(raw, key string) (string, bool) {
	panic(
		// Closed string: "key":"value"
		"fake")
}

// Truncated open string: "key":"value…EOF

func extractJSONBoolField(raw, key string) (bool, bool) { panic("fake") }

func extractJSONToolsArray(raw string) ([]struct {
	Name   string         `json:"name"`
	Params map[string]any `json:"params"`
}, bool) {
	panic("fake")
}

func unescapeJSONString(s string) string { panic("fake") }

func stripMarkdownJSONFence(s string) string { panic("fake") }

// drop language tag on first line

// extractBalancedJSONObject returns the first top-level {...} span with
// balanced braces (string-aware enough for typical plan JSON).
func extractBalancedJSONObject(s string) (string, bool) { panic("fake") }

func editFromResult(res *tools.ToolResult) *ProposedEdit { panic("fake") }

// openContextAmbientHeader labels auto-injected editor excerpts so models do
// not treat them as part of the user Task (KMA-202).
const openContextAmbientHeader = "Current editor (ambient — use only if relevant to the Task):\n"

func truncate(s string, n int) string { panic("fake") }

// PoolInvoker adapts dispatch.Pool to agent.Invoker / StreamInvoker.
// It accumulates PromptTokens/CompletionTokens across Invoke/InvokeStream
// calls for the lifetime of the invoker (typically one agent run). Usage()
// returns the running totals and does not reset them.
type PoolInvoker struct {
	Pool *dispatch.Pool

	mu               sync.Mutex
	promptTokens     int
	completionTokens int
	cacheRead        int
	lastGenerationID string
}

// Usage returns cumulative prompt and completion tokens from all successful
// Invoke/InvokeStream results so far. Totals are never zeroed by this method.
func (p *PoolInvoker) Usage() (prompt, completion int) { panic("fake") }

// LastGenerationID returns the most recent provider generation id observed
// on a successful invoke (KMA-211).
func (p *PoolInvoker) LastGenerationID() string { panic("fake") }

func (p *PoolInvoker) addUsage(res dispatch.InvokeResult) { panic("fake") }

// CacheReadUsage returns the cumulative provider cache-read token count
// (prompt_tokens_details.cached_tokens) observed across Invoke/InvokeStream
// results so far. Totals are never zeroed by this method.
func (p *PoolInvoker) CacheReadUsage() int { panic("fake") }

func (p *PoolInvoker) Invoke(ctx context.Context, model, prompt, provider string) (string, error) {
	panic("fake")
}

// InvokeWithTools runs a non-streaming remote (or local) call with OpenAI tools.
func (p *PoolInvoker) InvokeWithTools(ctx context.Context, model, prompt, provider string, tools []dispatch.ChatTool) (dispatch.InvokeResult, error) {
	panic("fake")
}

// KMA-203: hybrid/local race or local-only routes cannot attach tools[].
// When the agent chose the OpenAI protocol, pin to a tools-capable remote
// (prefer OpenRouter) instead of dropping tools for a content-only race.

// Hybrid race does not attach tools; fall back to content-only.

// InvokeWithToolMessages posts structured assistant/tool history (KMA-205/218).
func (p *PoolInvoker) InvokeWithToolMessages(ctx context.Context, model string, msgs []map[string]any, provider string, tools []dispatch.ChatTool) (dispatch.InvokeResult, error) {
	panic("fake")
}

// InvokeStream streams provider deltas when the resolved route supports it.
// Hybrid still races and returns the winner as one blob (stream winner-only
// mid-race is deferred — Wave A keeps race semantics intact).
//
// Non-streaming OpenAI-compat upstreams (common in tests and some gateways)
// return empty content under stream=true — fall back to Invoke so the agent
// JSON plan path stays reliable. TTFT watchdog lives in WrapStreamWithTTFT
// (invokeForTurn), not here.
func (p *PoolInvoker) InvokeStream(ctx context.Context, model, prompt, provider string, onDelta func(string)) (string, error) {
	panic("fake")
}

// StreamLimits is the TTFT + optional post-first-token idle watchdog (KMA-61).
type StreamLimits struct {
	TTFT time.Duration
	Idle time.Duration
}

// WrapStreamWithTTFT runs stream under a sub-context with a time-to-first-token
// watchdog. On TTFT timeout it falls back to fallback(parentCtx). Empty or
// errored stream results without a TTFT cancel are returned as-is so callers
// (e.g. PoolInvoker) can apply their own empty-content fallback.
func WrapStreamWithTTFT(
	parentCtx context.Context,
	ttft time.Duration,
	stream func(streamCtx context.Context, onDelta func(string)) (string, error),
	fallback func(context.Context) (string, error),
) (string, error) {
	panic("fake")
}

// WrapStreamWithLimits applies TTFT until the first token, then optional idle
// silence after that. Idle 0 disables the post-first-token watchdog.
func WrapStreamWithLimits(
	parentCtx context.Context,
	limits StreamLimits,
	stream func(streamCtx context.Context, onDelta func(string)) (string, error),
	fallback func(context.Context) (string, error),
) (string, error) {
	panic("fake")
}

// doneEvent builds a terminal "done" TurnEvent, copying PoolInvoker Usage()
// onto the event when the loop's invoker is a *PoolInvoker.
func (l *Loop) doneEvent(msg string, limitReached bool) TurnEvent { panic("fake") }

// invokeForTurn prefers StreamInvoker so mid-turn token batches reach the bus
// and SSE before the JSON plan is fully available (RT-THPUT-2).
// When nativeTools is true, uses non-streaming ExecuteWithTools (OpenAI path).
// KMA-207: no StreamWithTools API yet — native tool turns stay on InvokeWithTools
// (full response) rather than streaming tool-call deltas; revisit when dispatch
// exposes a streaming tools surface.
// For the JSON/Hermes stream path, cancels the provider stream as soon as a
// complete tool plan parses (KMA-86 mid-stream tool dispatch slice) so tool
// execution can start without waiting for trailing tokens.
func (l *Loop) invokeForTurn(ctx context.Context, model, prompt, provider string, turn int, emit func(TurnEvent), nativeTools bool, toolHistory []ToolHistoryTurn, stable, task string) (string, []dispatch.ToolCall, error) {
	panic("fake")
}

// Non-pool invokers (tests): fall back to plain Invoke + JSON decode path.

// Prefer the accumulator (complete plan); ignore cancel-induced stream errors.

// buildToolHistoryTurn maps one native OpenAI tool round into ToolHistoryTurn.
func buildToolHistoryTurn(assistantContent string, toolCalls []dispatch.ToolCall, observations []string) ToolHistoryTurn {
	panic("fake")
}

// tryEarlyToolPlan reports whether raw already contains a complete, executable
// tool plan (JSON or Hermes). Used to cut the provider stream short (KMA-86).
func tryEarlyToolPlan(raw string) bool { panic("fake") }

// resolveRouteDetail maps the UI provider hint to a concrete route string for
// Diagnostics / status pill (CFG.28). Returns (resolvedHint, humanDetail).
func (l *Loop) resolveRouteDetail(hint, prompt, model string) (string, string) { panic("fake") }

func emptyAs(s, fallback string) string { panic("fake") }

func emitHarnessTrace(ev TurnEvent) { panic("fake") }

// modelRequestContentMax caps wire prompt snapshot content in harness model_request (KMA-215).
const modelRequestContentMax = 8000

// Redact key=value style secrets inside prompt bodies (flightlog-compatible).
var harnessPromptKeyLike = regexp.MustCompile(`(?i)(api[_-]?key|token|secret|password|authorization)\s*[:=]\s*\S+`)

// Redact bare provider tokens embedded mid-prompt (sk-/nvapi-/Bearer …).
var harnessPromptBareSecret = regexp.MustCompile(`(?i)\b(?:sk-[A-Za-z0-9_-]{8,}|nvapi-[A-Za-z0-9_-]{8,}|Bearer\s+\S+)`)

// redactHarnessPromptContent strips credential-looking substrings before snapshot
// truncation so a cut cannot leave a partial secret unredacted (KMA-215).
func redactHarnessPromptContent(s string) string {
	panic(
		// Bearer first so "Authorization: Bearer <tok>" becomes "Authorization: [REDACTED]"
		// before keyLike collapses the authorization= form.
		"fake")
}

func modelRequestWireMessagesFromSnapshots(stable, volatile string) []map[string]string {
	panic("fake")
}

// protocolFromInferenceDetail extracts protocol=… from route detail strings
// (provider-agnostic; used for harness model_request richness).
func protocolFromInferenceDetail(detail string) string { panic("fake") }

func toolStartArgs(ev TurnEvent) string { panic("fake") }
