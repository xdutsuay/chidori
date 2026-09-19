package dispatch

import (
	"context"
	"time"
)

// DefaultStreamIdle is how long a stream may sit with no new tokens after the
// first token before the server cancels it (KMA-61). Distinct from TTFT
// (time to first token) and from RemoteRequestTimeout (total wall clock).
// YAML stream_idle: 0 uses this default; a negative duration disables idle.
const DefaultStreamIdle = 3 * time.Minute

// RunDeadlines is the shared server-owned lifecycle for one user run.
type RunDeadlines struct {
	Total time.Duration // wall-clock cap; 0 = RemoteRequestTimeout
	TTFT  time.Duration // time-to-first-token; 0 = agent default
	Idle  time.Duration // post-first-token stream silence; 0 = disabled
}

// Normalize fills zero Total with RemoteRequestTimeout.
func (d RunDeadlines) Normalize() RunDeadlines { panic("fake") }

// WithTotalTimeout wraps parent with the run's wall-clock cap.
func WithTotalTimeout(parent context.Context, d RunDeadlines) (context.Context, context.CancelFunc) {
	panic("fake")
}

// InputBudget is the prompt-side token budget for model (window minus output reserve).
func InputBudget(model string) int { panic("fake") }

// EstimatePromptTokens uses the same ~4 chars/token heuristic as compact.EstimateTokens.
func EstimatePromptTokens(prompt string) int { panic("fake") }

// ExceedsContext reports whether prompt is larger than the model's input budget.
func ExceedsContext(prompt, model string) (exceeds bool, estTokens, budget int) { panic("fake") }
