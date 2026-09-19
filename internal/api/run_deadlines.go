package api

import (
	"context"
	"net/http"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
)

func codeStreamIdle(cfg config.CodeConfig) time.Duration { panic("fake") }

// explicit disable

func (h *Handlers) runDeadlines() dispatch.RunDeadlines { panic("fake") }

// armUserRun applies the shared total-run deadline and Stop cancel registry
// (KMA-61). Caller must defer the returned cancel.
func (h *Handlers) armUserRun(parent context.Context, taskID string) (context.Context, context.CancelFunc) {
	panic("fake")
}

// enforceContextBudget always proceeds. Callers must run fitPromptToBudget
// first so oversized @file / history never hard-400 before chat_begin.
// allowTruncate is retained for API compat.
func (h *Handlers) enforceContextBudget(w http.ResponseWriter, prompt, model string, allowTruncate bool) bool {
	panic("fake")
}

// fitAndEnforceContextBudget clamps prompt to the model budget and returns it.
func (h *Handlers) fitAndEnforceContextBudget(prompt, model string) (fitted string, truncated bool) {
	panic("fake")
}

// fitPromptToBudget is the package-level clamp used by ChatCompletions / Agent
// and locked by run_deadlines_test.go / run_context_budget_test.go.
func fitPromptToBudget(prompt, model string) (fitted string, truncated bool) { panic("fake") }
