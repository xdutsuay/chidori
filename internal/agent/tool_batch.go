package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xdutsuay/lclreason/internal/tools"
)

type planTool struct {
	Name   string
	Params map[string]any
}

type toolStepOutcome struct {
	observations []string
	events       []TurnEvent
	edits        []ProposedEdit
	forceDone    bool
	forceDoneMsg string
}

func planToolsFromResponse(plan planResponse) []planTool { panic("fake") }

// runPlanTools executes plan tools sequentially, or concurrently when every tool
// in the batch is read-only parallel-safe and len>1 (KMA-209).
func (l *Loop) runPlanTools(
	ctx context.Context,
	turn int,
	model, provider string,
	tools []planTool,
	edits []ProposedEdit,
	execRunner ToolRunner,
	toolDedup *toolDedupCache,
	emit func(TurnEvent),
) (observations []string, addedEdits []ProposedEdit, earlyMsg string, early bool) {
	panic("fake")
}

func allToolsParallelSafeFromPlan(tools []planTool) bool { panic("fake") }

func (l *Loop) runPlanToolsParallel(
	ctx context.Context,
	turn int,
	model, provider string,
	tools []planTool,
	edits []ProposedEdit,
	execRunner ToolRunner,
	toolDedup *toolDedupCache,
	emit func(TurnEvent),
) (observations []string, addedEdits []ProposedEdit, earlyMsg string, early bool) {
	panic("fake")
}

func applyToolStepOutcome(out toolStepOutcome, emit func(TurnEvent)) { panic("fake") }

func (l *Loop) runOnePlanTool(
	ctx context.Context,
	turn int,
	model, provider string,
	t planTool,
	edits []ProposedEdit,
	execRunner ToolRunner,
	toolDedup *toolDedupCache,
	emit func(TurnEvent),
) (observations []string, newEdits []ProposedEdit, earlyMsg string, early bool) {
	panic("fake")
}

func cachedToolSkipOutcome(turn int, proposeOnly bool, name string, params map[string]any, cachedOut string, forceDone bool, firstTurn, repeatCount int) toolStepOutcome {
	panic("fake")
}

func (l *Loop) executePlanTool(
	ctx context.Context,
	turn int,
	model, provider string,
	t planTool,
	edits []ProposedEdit,
	execRunner ToolRunner,
	toolDedup *toolDedupCache,
) toolStepOutcome {
	panic("fake")
}

func toolParamsJSON(params map[string]any) string { panic("fake") }
