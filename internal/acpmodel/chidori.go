package acpmodel

// Chidori's extensions to the ACP v1 vocabulary, plus the session/turn records
// the agent workspace shell (KMA-168) lists and the Changes pane (KMA-174)
// groups by.
//
// Two rules keep this from drifting into a bespoke shape again:
//
//  1. Anything ACP already models is spelled exactly as ACP spells it. Where
//     ACP v2 has stabilized a concept v1 lacks (agent run state), the v2 name
//     and field shape are used verbatim so adopting v2 is a version bump, not
//     a rename.
//  2. Anything ACP genuinely does not model goes under `_meta` (Meta below) or
//     an underscore-prefixed sessionUpdate kind, both of which ACP reserves
//     for extension and instructs other hosts to tolerate.

// WireVersion is the version of chidori's normalized session-update stream.
//
// It is deliberately independent of ACPProtocolVersion: the ACP vocabulary can
// stay at v1 while chidori adds or retires `_meta` fields, and the desktop app
// and its LAN companion can be on different builds. Consumers negotiate it the
// same way the companion negotiates ChidoriProtocolVersion.
//
// Bump the minor for additive `_meta` or new extension kinds; bump the major
// only for a change that removes or re-types an existing field.
const WireVersion = "1.0.0"

// AgentState reports whether foreground work is in progress. Taken verbatim
// from ACP v2's `state_update` variant rather than invented, because chidori
// needs it today and v2 has already settled the shape.
type AgentState string

// Agent states. Background work may keep emitting updates while the agent
// reports StateIdle; those do not change the state.
const (
	// StateRunning means foreground work is in progress.
	StateRunning AgentState = "running"
	// StateIdle means the agent is ready for a new prompt. Carries a
	// StopReason when the transition ended foreground work.
	StateIdle AgentState = "idle"
	// StateRequiresAction means foreground work is blocked on the user —
	// a permission prompt, or a chidori clarification pause.
	StateRequiresAction AgentState = "requires_action"
)

// Extension session/update kinds.
const (
	// UpdateState reports an AgentState transition. This is ACP v2's stable
	// `state_update` variant, emitted under WireVersion while the ACP
	// vocabulary here is still v1.
	UpdateState SessionUpdateKind = "state_update"

	// UpdateChidoriClarification reports that the agent paused to ask the
	// user structured questions (the chidori-questions fence). ACP models
	// this as the `elicitation/create` client request rather than a session
	// update, so there is no v1 or v2 variant to borrow; the underscore
	// prefix marks it as non-core.
	UpdateChidoriClarification SessionUpdateKind = "_chidori.clarification"
)

// StateUpdate reports an AgentState transition. StopReason is set only on the
// transition into StateIdle that ends foreground work.
type StateUpdate struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	State         AgentState        `json:"state"`
	StopReason    StopReason        `json:"stopReason,omitempty"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// ClarificationUpdate carries the questions the agent is blocked on. The
// emitter should pair it with a StateUpdate of StateRequiresAction.
type ClarificationUpdate struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	Questions     []Question        `json:"questions"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// Question is one structured question from a clarification pause.
type Question struct {
	ID      string   `json:"id"`
	Prompt  string   `json:"prompt"`
	Choices []string `json:"choices,omitempty"`
}

// Meta is ACP's reserved `_meta` object. Chidori writes only under the
// Chidori key so a non-chidori host reading the same stream sees an object it
// is required by spec to ignore.
type Meta struct {
	Chidori *ChidoriMeta `json:"chidori,omitempty"`
}

// ChidoriMeta is the attribution ACP notifications do not carry but the
// workspace UI needs: stream ordering, which loop turn produced an update,
// which run and harness it belongs to, and how it was routed.
type ChidoriMeta struct {
	// WireVersion is the emitter's WireVersion. Set it on the first update of
	// a stream so a mismatched consumer can degrade instead of guessing.
	WireVersion string `json:"wireVersion,omitempty"`

	// Seq is a monotonic per-session counter. It is what makes replay,
	// gap detection, and de-duplication across the SSE and polled transports
	// possible; ACP relies on connection ordering alone, which the polled
	// transport does not provide.
	Seq uint64 `json:"seq,omitempty"`

	// TurnID identifies the prompt turn (chidori's task id / run id).
	TurnID string `json:"turnId,omitempty"`

	// Turn is the agent loop's iteration index within the turn. Diagnostic
	// only: it must never be load-bearing for correlating tool calls, which
	// is what ToolCallID is for.
	Turn int `json:"turn,omitempty"`

	// HarnessID names the producer: chidori's own loop or an external harness.
	HarnessID string `json:"harnessId,omitempty"`

	// Provider and Model record the resolved route for this update.
	Provider string `json:"provider,omitempty"`
	Model    string `json:"model,omitempty"`

	// TimestampMs is emitter-side epoch milliseconds.
	TimestampMs int64 `json:"timestampMs,omitempty"`

	// GenerationID is the provider completion id when known (KMA-211).
	GenerationID string `json:"generationId,omitempty"`

	// LimitReached marks an idle transition caused by chidori's own turn
	// budget rather than the model finishing. It is narrower than ACP's
	// StopMaxTurnRequests, which the emitter also sets, and exists so the UI
	// can offer a Continue affordance.
	LimitReached bool `json:"limitReached,omitempty"`

	// Edit is set on a ToolCallContent of type diff.
	Edit *ChidoriEditMeta `json:"edit,omitempty"`
}

// EditKind is how a proposed change was produced.
type EditKind string

// Edit kinds, matching the tools that produce them.
const (
	EditKindWriteFile  EditKind = "write_file"
	EditKindApplyPatch EditKind = "apply_patch"
)

// EditApplyState is where a change sits in the propose/apply flow.
type EditApplyState string

// Edit apply states.
const (
	// EditProposed means the change is awaiting user review and is not on
	// disk. Untrusted workspaces and plan mode force every edit here.
	EditProposed EditApplyState = "proposed"
	// EditApplied means the change is written to the workspace.
	EditApplied EditApplyState = "applied"
	// EditRejected means the user declined the change.
	EditRejected EditApplyState = "rejected"
)

// ChidoriEditMeta is the propose/apply axis ACP's diff content block has no
// field for. ACP diffs are purely descriptive: they say a file changed, not
// whether the change is on disk or awaiting review, which is the distinction
// the Changes pane is built around.
type ChidoriEditMeta struct {
	// EditID is stable across the propose -> apply transition so the UI can
	// update a row rather than append a duplicate.
	EditID string `json:"editId"`

	Kind       EditKind       `json:"kind,omitempty"`
	ApplyState EditApplyState `json:"applyState,omitempty"`
	Summary    string         `json:"summary,omitempty"`

	// RelPath is the workspace-relative path. ACP's ToolCallContent.Path is
	// required to be absolute, but every chidori write path and the UI's file
	// tree are workspace-relative, so both are carried rather than making
	// consumers re-derive one from the workspace root.
	RelPath string `json:"relPath,omitempty"`
}

// Session is one conversation context as the workspace shell lists it.
// Mirrors the ide_sessions row plus the live state the shell needs.
type Session struct {
	SessionID     SessionID `json:"sessionId"`
	Title         string    `json:"title,omitempty"`
	WorkspaceRoot string    `json:"workspaceRoot"`

	// ModeID is the chidori mode this session runs in (agent, plan, debug,
	// ask), reported over the wire as ACP's currentModeId.
	ModeID string `json:"currentModeId,omitempty"`

	Provider string `json:"provider,omitempty"`
	Model    string `json:"model,omitempty"`

	// HarnessID names which agent implementation serves this session.
	HarnessID string `json:"harnessId,omitempty"`

	// State is the session's live agent state.
	State AgentState `json:"state,omitempty"`

	// ActiveTurnID is the in-flight turn, if any. Chidori serializes one
	// active turn per session, so this is a single value rather than a list;
	// see the KMA-173 spec on lifting that constraint.
	ActiveTurnID string `json:"activeTurnId,omitempty"`

	CreatedAtMs    int64 `json:"createdAtMs,omitempty"`
	UpdatedAtMs    int64 `json:"updatedAtMs,omitempty"`
	LastActiveAtMs int64 `json:"lastActiveAtMs,omitempty"`

	Meta *Meta `json:"_meta,omitempty"`
}

// Turn is one prompt turn: everything between a user prompt and the idle
// StateUpdate that ends it. ACP has no turn object on the wire — the turn is
// implicit in the session/prompt request and its response — so this is the
// record chidori persists and the shell groups a transcript by.
type Turn struct {
	TurnID    string    `json:"turnId"`
	SessionID SessionID `json:"sessionId"`

	// Prompt is the user content that opened the turn.
	Prompt []ContentBlock `json:"prompt,omitempty"`

	State AgentState `json:"state"`

	// StopReason is set once State is StateIdle.
	StopReason StopReason `json:"stopReason,omitempty"`

	HarnessID string `json:"harnessId,omitempty"`
	Provider  string `json:"provider,omitempty"`
	Model     string `json:"model,omitempty"`

	StartedAtMs int64 `json:"startedAtMs,omitempty"`
	EndedAtMs   int64 `json:"endedAtMs,omitempty"`

	// Usage is the turn's context and cost accounting, same shape as the
	// usage_update variant carries.
	Usage *UsageUpdate `json:"usage,omitempty"`

	// Error is the human-readable failure text when the turn failed. ACP has
	// no failure StopReason — an agent crash surfaces as a JSON-RPC error
	// response to session/prompt, which is not part of the update stream — so
	// the text lands on the turn record next to an idle StateUpdate.
	Error string `json:"error,omitempty"`

	Meta *Meta `json:"_meta,omitempty"`
}
