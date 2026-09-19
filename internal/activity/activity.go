// Package activity is a tiny in-process event bus for request-lifecycle
// activity (planning, tool calls, inference, node selection). It backs two
// views: the chat UI's live status pill (a filtered subset streamed over SSE)
// and the developer /debug timeline (the full recent history).
//
// See docs/adr/ADR-0001-activity-events.md.
package activity

import (
	"sync"
	"sync/atomic"
	"time"
)

// Activity is one lifecycle event for a request turn.
type Activity struct {
	RequestID  string `json:"request_id"`
	ClientID   string `json:"client_id,omitempty"`
	Kind       string `json:"kind"`  // planning | tool_call | tool_result | edit_proposed | inference | node_selected | cache_hit | done | error
	Label      string `json:"label"` // human text for the status pill
	Detail     string `json:"detail,omitempty"`
	Node       string `json:"node,omitempty"`
	DurationMs int64  `json:"duration_ms,omitempty"`
	Timestamp  int64  `json:"timestamp"`
	// GenerationID is the provider completion id when known (KMA-211).
	GenerationID string `json:"generation_id,omitempty"`
	// PromptTokens/CompletionTokens are set on terminal done rows when known.
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
}

// Emitter receives activity events for one request.
type Emitter interface {
	Emit(Activity)
}

// Emit is a nil-safe helper: it stamps the time and forwards to em (or no-ops
// when em is nil), so emit sites don't need nil checks.
func Emit(em Emitter, a Activity) { panic("fake") }

// Bus is a bounded ring buffer of recent activities for the debug view. It can
// be toggled off at runtime — when disabled, Record is a no-op (the chat status
// pill is independent and keeps working regardless).
type Bus struct {
	mu      sync.Mutex
	buf     []Activity
	max     int
	enabled atomic.Bool
}

// NewBus creates a bus retaining the most recent max events (default 500),
// enabled by default.
func NewBus(max int) *Bus { panic("fake") }

// SetEnabled toggles recording. Disabling does not clear existing history.
func (b *Bus) SetEnabled(on bool) { panic("fake") }

// Enabled reports whether recording is on.
func (b *Bus) Enabled() bool { panic("fake") }

// Record appends an activity, trimming to the retention bound. No-op when the
// bus is disabled.
func (b *Bus) Record(a Activity) { panic("fake") }

// Recent returns up to limit activities, newest first, optionally filtered by
// clientID and/or requestID (empty string = no filter).
func (b *Bus) Recent(limit int, clientID, requestID string) []Activity { panic("fake") }

// LabelForTool maps a tool name to a friendly status label for the chat pill.
func LabelForTool(tool string) string { panic("fake") }
