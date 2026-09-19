package api

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"sync"
)

// runCancelRegistry maps in-flight Ask/Agent/Plan/Debug task IDs to a cancel
// func. POST /api/runs/{id}/cancel uses this so Stop works when WKWebView
// abort does not cancel r.Context() (KMA-6 / KMA-61).
type runCancelRegistry struct {
	mu      sync.Mutex
	cancels map[string]context.CancelFunc
}

func newRunCancelRegistry() *runCancelRegistry { panic("fake") }

func (r *runCancelRegistry) Register(id string, cancel context.CancelFunc) { panic("fake") }

func (r *runCancelRegistry) Unregister(id string) { panic("fake") }

func (r *runCancelRegistry) Has(id string) bool { panic("fake") }

func (r *runCancelRegistry) Cancel(id string) bool { panic("fake") }

// CancelAll cancels every armed run and clears the registry. Returns the
// cancelled task IDs (order is not significant).
func (r *runCancelRegistry) CancelAll() []string { panic("fake") }

// AbortActiveRuns cancels every armed foreground run and emits agent_run_end
// with end=aborted and reason shutdown|aborted so Diagnostics is not silent
// on process teardown (KMA-233). Returns cancelled task IDs. Nil-safe.
func (h *Handlers) AbortActiveRuns(reason string) []string { panic("fake") }

// armRunCancel registers a child cancel under taskID. Caller must defer the
// returned cancel (unregisters + cancels). Safe if h.runCancels is nil.
func (h *Handlers) armRunCancel(parent context.Context, taskID string) (context.Context, context.CancelFunc) {
	panic("fake")
}

// cancelRun stops an in-flight Ask/Agent/Plan/Debug run or background job.
// source is logged on flight (ui, companion, cancel_api, …).
func (h *Handlers) cancelRun(id, source string) bool { panic("fake") }

// RunCancel is POST /api/runs/{id}/cancel — Stop for foreground Ask/Agent
// (and background jobs). Independent of the original HTTP request abort.
func (h *Handlers) RunCancel(w http.ResponseWriter, r *http.Request) { panic("fake") }

// runEndFields classifies how a user run finished for flight/activity.
func runEndFields(taskID string, ctx context.Context, err error) map[string]any { panic("fake") }
