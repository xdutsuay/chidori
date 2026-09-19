// Package replay provides read-only helpers for inspecting persisted harness
// traces. It does not re-execute agent runs (see docs/adr/ADR-draft-replay-runs.md).
package replay

import (
	"sort"
	"time"

	"github.com/xdutsuay/lclreason/internal/harnesstrace"
)

// Summary is a compact description of one harness trace for meta-agents and APIs.
type Summary struct {
	ID         string         `json:"id,omitempty"`
	EventCount int            `json:"event_count"`
	KindCounts map[string]int `json:"kind_counts"`
	Tools      []string       `json:"tools,omitempty"`
	MaxTurn    int            `json:"max_turn"`
	TaskID     string         `json:"task_id,omitempty"`
	RunID      string         `json:"run_id,omitempty"`
	StartedAt  time.Time      `json:"started_at,omitempty"`
	EndedAt    time.Time      `json:"ended_at,omitempty"`
	DurationMs int64          `json:"duration_ms"`
}

// Describe computes a summary from decoded harness events.
func Describe(id string, events []harnesstrace.Event) Summary { panic("fake") }

func intFromAny(v any) (int, bool) { panic("fake") }
