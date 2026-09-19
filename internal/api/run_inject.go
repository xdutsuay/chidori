package api

import (
	"context"
	"html"
	"strings"
	"sync"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/flightlog"
	"github.com/xdutsuay/lclreason/internal/hooks"
	"github.com/xdutsuay/lclreason/internal/prompts"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// runInjectRegistry holds companion-injected follow-up prompts keyed by run_id.
// Active runs queue here until the run finishes; idle runs drain immediately.
type runInjectRegistry struct {
	mu    sync.Mutex
	queue map[string][]string
}

func newRunInjectRegistry() *runInjectRegistry { panic("fake") }

func (r *runInjectRegistry) Push(runID, text string) { panic("fake") }

func (r *runInjectRegistry) Pop(runID string) (string, bool) { panic("fake") }

func (r *runInjectRegistry) Len(runID string) int { panic("fake") }

func (h *Handlers) endCompanionRun(taskID string, ok bool, errMsg string) { panic("fake") }

func (h *Handlers) isRunActive(id string) bool { panic("fake") }

func (h *Handlers) runKnownToCompanion(runID string) bool { panic("fake") }

func (h *Handlers) appendCompanionInjectMessage(ctx context.Context, runID, text string) {
	panic("fake")
}

func (h *Handlers) enqueueCompanionInject(ctx context.Context, runID, text string) string {
	panic("fake")
}

func (h *Handlers) drainCompanionInject(runID string) { panic("fake") }

func (h *Handlers) tryStartCompanionInjectRun(runID string) { panic("fake") }

func (h *Handlers) runCompanionInjectJob(req AgentRunRequest, taskID string) { panic("fake") }
