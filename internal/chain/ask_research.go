package chain

import (
	"context"
	"fmt"
	"strings"

	"github.com/xdutsuay/lclreason/internal/activity"
)

// askResearchTools is the allowlist for H.5 — read-only / lookup tools only.
// No shell_exec, writes, or memory mutation in Ask research mode.
var askResearchTools = map[string]bool{
	"web_search":           true,
	"fetch_url":            true,
	"calculator":           true,
	"read_file":            true,
	"grep_workspace":       true,
	"list_dir":             true,
	"memory_search":        true,
	"search_chat_messages": true,
}

// HeuristicPlan returns deterministic tool steps for a query (no LLM), or nil.
func HeuristicPlan(query string) []PlanStep { panic("fake") }

func filterResearchPlan(steps []PlanStep) []PlanStep { panic("fake") }

// AugmentAskResearch runs the planner (or heuristics) and executes only
// read-only tools when the user opts into Ask "research" mode (H.5). Default
// Ask stays a one-shot answer with no tool loop.
func (e *Engine) AugmentAskResearch(ctx context.Context, input, model string, em activity.Emitter) (string, []PlanStep) {
	panic("fake")
}
