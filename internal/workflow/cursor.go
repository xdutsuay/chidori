package workflow

import "errors"

var errWorkflowDone = errors.New("workflow steps exhausted")

// StepCursor locates the next step to execute. Top is the index in wf.Steps;
// Sub, when set, points into a condition/loop branch body.
type StepCursor struct {
	Top int        `json:"top"`
	Sub *SubCursor `json:"sub,omitempty"`
}

type SubCursor struct {
	Branch string `json:"branch"` // then | else | loop
	Index  int    `json:"index"`
}

func startCursor() StepCursor { panic("fake") }

func (c StepCursor) step(wf *Workflow) (*Step, error) { panic("fake") }

func branchSteps(step *Step, branch string) []Step { panic("fake") }

// advanceWithinBranch moves to the next step inside the current branch.
// Returns false when the branch body is exhausted (Sub is unchanged).
func advanceWithinBranch(wf *Workflow, cur StepCursor) (StepCursor, bool) { panic("fake") }

// advancePastComposite clears Sub and moves to the next top-level step.
func advancePastComposite(wf *Workflow, cur StepCursor) (StepCursor, bool) { panic("fake") }

// advanceCursor moves past a completed top-level leaf step.
func advanceCursor(wf *Workflow, cur StepCursor) (StepCursor, bool) { panic("fake") }

// cursorAfterGate advances past an approval_gate at cur without re-running it.
func cursorAfterGate(wf *Workflow, cur StepCursor) (StepCursor, bool) { panic("fake") }

func syncLegacyStepIndex(state *ExecutionState) { panic("fake") }
