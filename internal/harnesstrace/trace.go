// Package harnesstrace is an optional, off-by-default pub/sub trace bus for
// agent harness debugging (PRD harness visualizer). Dependency direction is
// one-way: callers import harnesstrace; this package imports nothing from
// internal/api, internal/tools, or internal/agent.
package harnesstrace

import (
	"sync"
	"sync/atomic"
	"time"
)

const defaultRingCap = 500

// Event is one harness trace point (tool call, model call, turn boundary).
type Event struct {
	Kind      string         `json:"kind"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload,omitempty"`
}

var enabled atomic.Bool

type subscriber struct {
	ch chan Event
}

var (
	mu          sync.RWMutex
	subscribers []subscriber
	ring        []Event
	ringCap     = defaultRingCap
)

// SetEnabled flips the kill switch (config / Settings toggle).
func SetEnabled(v bool) { panic("fake") }

// Enabled reports whether trace events are collected and emitted.
func Enabled() bool { panic("fake") }

// Emit records an event when enabled. Safe to call unconditionally — cost is
// one atomic load when disabled.
func Emit(e Event) { panic("fake") }

// Subscribe returns a channel of live events plus an unsubscribe func.
// The ring buffer snapshot is replayed first (best-effort, non-blocking).
func Subscribe() (<-chan Event, func()) { panic("fake") }

// Recent returns up to limit most recent events (newest last).
func Recent(limit int) []Event { panic("fake") }

const maxPayloadStr = 512

func truncatePayload(p map[string]any) map[string]any { panic("fake") }
