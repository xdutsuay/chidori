package agenteval

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Task is one canned agent-mode scenario. Offline mode uses ScriptedReplies
// (JSON plan turns) so the harness can be scored without a live LLM. Live mode
// (future) ignores ScriptedReplies and posts Task to a real provider.
type Task struct {
	ID              string   `json:"id"`
	Prompt          string   `json:"prompt"`
	ScriptedReplies []string `json:"scripted_replies"`
	// ExpectToolsOrdered is the subsequence of tool names that must appear in
	// order among successful tool_start events (extras are allowed unless
	// ForbidExtraTools is set).
	ExpectToolsOrdered []string `json:"expect_tools_ordered"`
	// ForbidTools must never appear in the run's tool_start list.
	ForbidTools []string `json:"forbid_tools,omitempty"`
	// MaxToolCalls caps total tool_start events (0 = no cap). Catches
	// redundant-call loops without requiring an exact call count.
	MaxToolCalls int `json:"max_tool_calls,omitempty"`
	// RequireDone requires a terminal done event (not limit_reached).
	RequireDone bool `json:"require_done"`
	// ForbidUnknownTool fails if any error output mentions "unknown tool".
	ForbidUnknownTool bool `json:"forbid_unknown_tool"`
	// SummaryContains substrings that must appear in the final done message.
	SummaryContains []string `json:"summary_contains,omitempty"`
}

// Trace is a compact recording of one Loop.Run for scoring.
type Trace struct {
	TaskID     string   `json:"task_id"`
	Summary    string   `json:"summary"`
	ToolStarts []string `json:"tool_starts"`
	Errors     []string `json:"errors"`
	Done       bool     `json:"done"`
	LimitHit   bool     `json:"limit_hit"`
	RunErr     string   `json:"run_err,omitempty"`
}

// Check is one scored assertion against a Trace.
type Check struct {
	Name    string `json:"name"`
	Pass    bool   `json:"pass"`
	Detail  string `json:"detail,omitempty"`
	Harness bool   `json:"harness"` // true = harness contract (not model quality)
}

// Score evaluates Trace against Task expectations.
func Score(task Task, tr Trace) []Check { panic("fake") }

// may fail on a weak live model even when harness is fine

// ReportJSON pretty-prints a score report for the learnings log / CI artifact.
func ReportJSON(task Task, tr Trace, checks []Check) string { panic("fake") }

func orderedSubsequence(got, want []string) (bool, string) { panic("fake") }

func truncate(s string, n int) string { panic("fake") }
