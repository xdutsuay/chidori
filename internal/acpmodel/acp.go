// Package acpmodel is the ACP-shaped session/event vocabulary the agent
// workspace UI reads (KMA-173).
//
// Every type here mirrors the Agent Client Protocol v1 schema
// (https://agentclientprotocol.com/protocol/v1/schema, consulted 2026-09-07)
// field-for-field, including ACP's camelCase JSON names. That is deliberate:
// an external ACP harness's `session/update` payload must be representable
// here without renaming, so KMA-23 adapters stay a passthrough rather than a
// second translation layer.
//
// Chidori-specific attribution (turn number, run id, harness id, sequence)
// lives under ACP's reserved `_meta` extension point — see chidori.go — never
// as new top-level fields, so a non-chidori host can ignore it per spec.
//
// Nothing in this package has behavior: it is declarations only, so the
// contract can be reviewed and tested as a wire shape. Decoding a
// SessionNotification is a two-pass unmarshal (UpdateEnvelope to read the
// discriminator, then the concrete variant), which is the same pattern
// internal/acpclient/map.go already uses.
package acpmodel

import "encoding/json"

// ACPProtocolVersion is the ACP major version this vocabulary is shaped
// after, and the value internal/acpclient negotiates in `initialize`.
//
// ACP v2 exists but its own migration guide labels the whole v2 surface
// draft, so v1 is the normative baseline here. See chidori.go for how the
// wire format versions independently of ACP itself.
const ACPProtocolVersion = 1

// SessionID identifies one conversation context. ACP: SessionId.
type SessionID string

// MessageID groups streamed chunks into one logical message. Chunks sharing a
// MessageID belong to the same message; a change starts a new message.
// Optional in ACP v1, required in v2 — always set it.
type MessageID string

// ToolCallID identifies one tool call within a session. ACP: ToolCallId.
type ToolCallID string

// TerminalID identifies a terminal created via `terminal/create`.
type TerminalID string

// SessionUpdateKind is the `sessionUpdate` discriminator on a session/update
// notification payload. ACP: SessionUpdate union tag.
type SessionUpdateKind string

// ACP v1 session/update variants.
const (
	UpdateUserMessageChunk  SessionUpdateKind = "user_message_chunk"
	UpdateAgentMessageChunk SessionUpdateKind = "agent_message_chunk"
	UpdateAgentThoughtChunk SessionUpdateKind = "agent_thought_chunk"
	UpdateToolCall          SessionUpdateKind = "tool_call"
	UpdateToolCallUpdate    SessionUpdateKind = "tool_call_update"
	UpdatePlan              SessionUpdateKind = "plan"
	UpdateAvailableCommands SessionUpdateKind = "available_commands_update"
	UpdateCurrentMode       SessionUpdateKind = "current_mode_update"
	UpdateSessionInfo       SessionUpdateKind = "session_info_update"
	UpdateUsage             SessionUpdateKind = "usage_update"
)

// ContentBlockType is the `type` discriminator on a ContentBlock.
type ContentBlockType string

// ACP v1 ContentBlock variants.
const (
	ContentText         ContentBlockType = "text"
	ContentImage        ContentBlockType = "image"
	ContentAudio        ContentBlockType = "audio"
	ContentResourceLink ContentBlockType = "resource_link"
	ContentResource     ContentBlockType = "resource"
)

// ContentBlock is one displayable unit of agent or user content. ACP keeps
// this MCP-compatible so tool output can be forwarded untransformed.
//
// Field validity by Type:
//
//	text          Text
//	image, audio  Data (base64), MimeType
//	resource_link URI, Name, MimeType, Size
//	resource      Resource
type ContentBlock struct {
	Type ContentBlockType `json:"type"`

	Text string `json:"text,omitempty"`

	Data     string `json:"data,omitempty"`
	MimeType string `json:"mimeType,omitempty"`

	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
	Size *int64 `json:"size,omitempty"`

	Resource *EmbeddedResource `json:"resource,omitempty"`

	Annotations json.RawMessage `json:"annotations,omitempty"`
	Meta        *Meta           `json:"_meta,omitempty"`
}

// EmbeddedResource carries resource contents inline: Text for text resources,
// Blob (base64) for binary ones. ACP: TextResourceContents / BlobResourceContents.
type EmbeddedResource struct {
	URI      string `json:"uri"`
	MimeType string `json:"mimeType,omitempty"`
	Text     string `json:"text,omitempty"`
	Blob     string `json:"blob,omitempty"`
}

// ToolKind categorizes a tool call so clients can pick icons and UI
// treatment. ACP: ToolKind.
type ToolKind string

// ACP v1 tool kinds. ToolKindOther is the documented default.
const (
	ToolKindRead       ToolKind = "read"
	ToolKindEdit       ToolKind = "edit"
	ToolKindDelete     ToolKind = "delete"
	ToolKindMove       ToolKind = "move"
	ToolKindSearch     ToolKind = "search"
	ToolKindExecute    ToolKind = "execute"
	ToolKindThink      ToolKind = "think"
	ToolKindFetch      ToolKind = "fetch"
	ToolKindSwitchMode ToolKind = "switch_mode"
	ToolKindOther      ToolKind = "other"
)

// ToolCallStatus is a tool call's lifecycle state. ACP: ToolCallStatus.
type ToolCallStatus string

// ACP v1 tool call statuses. ToolCallPending is the documented default and
// covers both "input still streaming" and "awaiting approval".
const (
	ToolCallPending    ToolCallStatus = "pending"
	ToolCallInProgress ToolCallStatus = "in_progress"
	ToolCallCompleted  ToolCallStatus = "completed"
	ToolCallFailed     ToolCallStatus = "failed"
)

// ToolCallContentType is the `type` discriminator on ToolCallContent.
type ToolCallContentType string

// ACP v1 tool call content variants.
const (
	ToolContentContent  ToolCallContentType = "content"
	ToolContentDiff     ToolCallContentType = "diff"
	ToolContentTerminal ToolCallContentType = "terminal"
)

// ToolCallContent is output produced by a tool call: a content block, a file
// diff, or a reference to a live terminal.
//
// The diff variant is the Changes pane's source of truth (KMA-174): Path is
// absolute, OldText is nil for a newly created file, and NewText is the full
// post-edit content. ACP carries whole-file before/after rather than a
// unified patch, so the client owns diff rendering.
type ToolCallContent struct {
	Type ToolCallContentType `json:"type"`

	Content *ContentBlock `json:"content,omitempty"`

	Path    string  `json:"path,omitempty"`
	OldText *string `json:"oldText,omitempty"`
	NewText string  `json:"newText,omitempty"`

	TerminalID TerminalID `json:"terminalId,omitempty"`

	Meta *Meta `json:"_meta,omitempty"`
}

// ToolCallLocation is a file (and optional line) a tool call touches, so the
// client can implement follow-along. ACP: ToolCallLocation.
type ToolCallLocation struct {
	Path string `json:"path"`
	Line *int   `json:"line,omitempty"`
	Meta *Meta  `json:"_meta,omitempty"`
}

// PlanEntryStatus is a plan entry's execution state. ACP: PlanEntry.status.
type PlanEntryStatus string

// ACP v1 plan entry statuses.
const (
	PlanEntryPending    PlanEntryStatus = "pending"
	PlanEntryInProgress PlanEntryStatus = "in_progress"
	PlanEntryCompleted  PlanEntryStatus = "completed"
)

// PlanEntryPriority is a plan entry's relative importance. ACP: PlanEntry.priority.
type PlanEntryPriority string

// ACP v1 plan entry priorities.
const (
	PlanEntryHigh   PlanEntryPriority = "high"
	PlanEntryMedium PlanEntryPriority = "medium"
	PlanEntryLow    PlanEntryPriority = "low"
)

// PlanEntry is one task in the agent's execution plan. All three fields are
// required by ACP.
type PlanEntry struct {
	Content  string            `json:"content"`
	Priority PlanEntryPriority `json:"priority"`
	Status   PlanEntryStatus   `json:"status"`
	Meta     *Meta             `json:"_meta,omitempty"`
}

// StopReason says why the agent stopped a prompt turn. ACP: StopReason,
// returned in the session/prompt response.
type StopReason string

// ACP v1 stop reasons.
const (
	StopEndTurn         StopReason = "end_turn"
	StopMaxTokens       StopReason = "max_tokens"
	StopMaxTurnRequests StopReason = "max_turn_requests"
	StopRefusal         StopReason = "refusal"
	StopCancelled       StopReason = "cancelled"
)

// SessionNotification is the session/update envelope. Update holds one
// session/update variant; read UpdateEnvelope from it first to learn which.
type SessionNotification struct {
	SessionID SessionID       `json:"sessionId"`
	Update    json.RawMessage `json:"update"`
	Meta      *Meta           `json:"_meta,omitempty"`
}

// UpdateEnvelope is the discriminator-and-metadata prefix every session/update
// variant shares. Unmarshal a SessionNotification.Update into this to route,
// then unmarshal the same bytes into the concrete variant below.
type UpdateEnvelope struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// UserMessageChunk appends one content item to the user's message.
// SessionUpdate is UpdateUserMessageChunk.
type UserMessageChunk struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	MessageID     MessageID         `json:"messageId,omitempty"`
	Content       ContentBlock      `json:"content"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// AgentMessageChunk appends one content item to the agent's reply.
// SessionUpdate is UpdateAgentMessageChunk.
type AgentMessageChunk struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	MessageID     MessageID         `json:"messageId,omitempty"`
	Content       ContentBlock      `json:"content"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// AgentThoughtChunk appends one content item to the agent's reasoning stream.
// SessionUpdate is UpdateAgentThoughtChunk.
type AgentThoughtChunk struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	MessageID     MessageID         `json:"messageId,omitempty"`
	Content       ContentBlock      `json:"content"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// ToolCallStart announces a new tool call. SessionUpdate is UpdateToolCall.
//
// ACP v2 folds this into ToolCallPatch (first patch for an unseen ToolCallID
// creates the call), so emitters should keep the field set here identical to
// ToolCallPatch's to make that migration a rename.
type ToolCallStart struct {
	SessionUpdate SessionUpdateKind  `json:"sessionUpdate"`
	ToolCallID    ToolCallID         `json:"toolCallId"`
	Title         string             `json:"title,omitempty"`
	Kind          ToolKind           `json:"kind,omitempty"`
	Status        ToolCallStatus     `json:"status,omitempty"`
	Content       []ToolCallContent  `json:"content,omitempty"`
	Locations     []ToolCallLocation `json:"locations,omitempty"`
	RawInput      json.RawMessage    `json:"rawInput,omitempty"`
	RawOutput     json.RawMessage    `json:"rawOutput,omitempty"`
	Meta          *Meta              `json:"_meta,omitempty"`
}

// ToolCallPatch updates an existing tool call. SessionUpdate is
// UpdateToolCallUpdate.
//
// Pointer fields carry ACP's three-state patch semantics: absent leaves the
// previous value alone, JSON null clears it, a value replaces it. Array fields
// are replaced wholesale, never appended to. Only ToolCallID is required.
type ToolCallPatch struct {
	SessionUpdate SessionUpdateKind   `json:"sessionUpdate"`
	ToolCallID    ToolCallID          `json:"toolCallId"`
	Title         *string             `json:"title,omitempty"`
	Kind          *ToolKind           `json:"kind,omitempty"`
	Status        *ToolCallStatus     `json:"status,omitempty"`
	Content       *[]ToolCallContent  `json:"content,omitempty"`
	Locations     *[]ToolCallLocation `json:"locations,omitempty"`
	RawInput      json.RawMessage     `json:"rawInput,omitempty"`
	RawOutput     json.RawMessage     `json:"rawOutput,omitempty"`
	Meta          *Meta               `json:"_meta,omitempty"`
}

// PlanUpdate replaces the session's whole plan. SessionUpdate is UpdatePlan.
// ACP requires the complete entry list on every update; the client replaces
// rather than merges.
type PlanUpdate struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	Entries       []PlanEntry       `json:"entries"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// CurrentModeUpdate reports that the session's mode changed (chidori's
// agent/plan/debug/ask). SessionUpdate is UpdateCurrentMode.
type CurrentModeUpdate struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	CurrentModeID string            `json:"currentModeId"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// SessionInfoUpdate reports changed session metadata. Title and LastActiveAt
// are patch fields: absent leaves them alone, null clears them.
// SessionUpdate is UpdateSessionInfo.
type SessionInfoUpdate struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	Title         *string           `json:"title,omitempty"`
	LastActiveAt  *string           `json:"lastActiveAt,omitempty"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// UsageUpdate reports context-window consumption and cumulative cost.
// Used and Size are required token counts. SessionUpdate is UpdateUsage.
type UsageUpdate struct {
	SessionUpdate SessionUpdateKind `json:"sessionUpdate"`
	Used          int64             `json:"used"`
	Size          int64             `json:"size"`
	Cost          *Cost             `json:"cost,omitempty"`
	Meta          *Meta             `json:"_meta,omitempty"`
}

// Cost is a cumulative session cost. Currency is an ISO 4217 code.
type Cost struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
