package workflow

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/xdutsuay/lclreason/internal/executil"
)

var ErrPaused = errors.New("execution paused")

const (
	defaultShellTimeout = 30 * time.Second
	maxOutputBytes      = 32 * 1024
	maxLoopItems        = 50
)

// Invoker performs a one-shot LLM completion for llm_call steps.
type Invoker interface {
	Invoke(ctx context.Context, model, prompt string) (string, error)
}

type Engine struct {
	workRoot string
	invoker  Invoker
	mu       sync.RWMutex
	active   map[string]*ExecutionState
}

func NewEngine(root string, inv Invoker) *Engine { panic("fake") }

func (e *Engine) Run(ctx context.Context, wf *Workflow) (*ExecutionState, error) { panic("fake") }

// Resume continues a paused approval_gate run. resumeToken must match state.ResumeToken when set.
// Sleep and event pauses must use ResumeAfterSleep / Signal — not this path.
func (e *Engine) Resume(ctx context.Context, runID string, wf *Workflow, resumeToken string) (*ExecutionState, error) {
	panic("fake")
}

// ResumeAfterSleep continues a run paused on a sleep step (after ResumeAt).
func (e *Engine) ResumeAfterSleep(ctx context.Context, runID string, wf *Workflow) (*ExecutionState, error) {
	panic("fake")
}

// Signal continues a run paused on wait_event when event matches WaitingEvent.
func (e *Engine) Signal(ctx context.Context, runID string, wf *Workflow, event, resumeToken string) (*ExecutionState, error) {
	panic("fake")
}

// continueAfterPause advances the cursor past the pause step and continues the run loop.
func (e *Engine) continueAfterPause(ctx context.Context, state *ExecutionState, wf *Workflow, logMsg string) (*ExecutionState, error) {
	panic("fake")
}

func (e *Engine) runLoop(ctx context.Context, wf *Workflow, state *ExecutionState) error {
	panic("fake")
}

func (e *Engine) runCondition(ctx context.Context, wf *Workflow, step *Step, state *ExecutionState) error {
	panic("fake")
}

func (e *Engine) runLoopStep(ctx context.Context, wf *Workflow, step *Step, state *ExecutionState) error {
	panic("fake")
}

func (e *Engine) runBranch(ctx context.Context, wf *Workflow, state *ExecutionState) error {
	panic("fake")
}

func (e *Engine) executeLeafStep(ctx context.Context, step *Step, state *ExecutionState) error {
	panic("fake")
}

func (e *Engine) recordStepAttempt(state *ExecutionState, stepID string, attempt int, msg string) {
	panic("fake")
}

func (e *Engine) executeLeafStepOnce(ctx context.Context, step *Step, state *ExecutionState) error {
	panic("fake")
}

func (e *Engine) runShell(ctx context.Context, step *Step, state *ExecutionState) error {
	panic("fake")
}

func (e *Engine) runLLMCall(ctx context.Context, step *Step, state *ExecutionState) error {
	panic("fake")
}

func renderLLMTemplate(tmpl string, state *ExecutionState) (string, error) { panic("fake") }

func (e *Engine) resolveLoopItems(expr string, state *ExecutionState) ([]string, error) {
	panic("fake")
}

func (e *Engine) track(state *ExecutionState) { panic("fake") }

func newRunID() (string, error) { panic("fake") }

func newResumeToken() string { panic("fake") }

func truncateOutput(s string) string { panic("fake") }

// MatchSaveTrigger reports whether path matches an on_file_save workflow trigger.
func MatchSaveTrigger(wf Workflow, path string) bool { panic("fake") }

// EnsureRunsDir creates the runs directory if missing (used by tests).
func EnsureRunsDir(root string) error { panic("fake") }
