package workflow

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var ErrRunCapacity = errors.New("workflow run capacity reached")

const DefaultMaxConcurrentRuns = 4

// Runtime is a workspace-scoped workflow executor with shared engine state and
// a run admission cap across every trigger path.
type Runtime struct {
	mu          sync.Mutex
	root        string
	invoker     Invoker
	engine      *Engine
	slots       chan struct{}
	maxRuns     int
	inFlight    int32
	recoverOnce sync.Once
	sleepTimers map[string]*time.Timer
}

func NewRuntime(root string, inv Invoker, maxConcurrent int) *Runtime { panic("fake") }

func (r *Runtime) Root() string { panic("fake") }

// RecoverOnBoot re-attaches paused runs and re-arms sleep timers / event waits.
func (r *Runtime) RecoverOnBoot(ctx context.Context) { panic("fake") }

func (r *Runtime) Run(ctx context.Context, wf *Workflow) (*ExecutionState, error) { panic("fake") }

func (r *Runtime) Resume(ctx context.Context, runID string, wf *Workflow, resumeToken string) (*ExecutionState, error) {
	panic("fake")
}

// Signal resumes a wait_event pause when event matches WaitingEvent.
func (r *Runtime) Signal(ctx context.Context, runID, event, resumeToken string) (*ExecutionState, error) {
	panic("fake")
}

// TryRunAsync acquires a slot without blocking; on success it runs fn in a new goroutine.
func (r *Runtime) TryRunAsync(ctx context.Context, fn func(context.Context)) bool { panic("fake") }

func (r *Runtime) armIfPaused(st *ExecutionState) { panic("fake") }

// Event waits stay tracked until Signal; engine.track already holds them.

func (r *Runtime) scheduleSleep(runID string, resumeAt time.Time) { panic("fake") }

func (r *Runtime) fireSleep(runID string) { panic("fake") }

func (r *Runtime) acquire(ctx context.Context) error { panic("fake") }

func (r *Runtime) release() { panic("fake") }

func (r *Runtime) InFlight() int { panic("fake") }

func (r *Runtime) MaxConcurrentRuns() int { panic("fake") }

func (r *Runtime) Engine() *Engine { panic("fake") }

// FormatCapacityErr wraps ErrRunCapacity with the configured limit.
func FormatCapacityErr(max int) error { panic("fake") }
