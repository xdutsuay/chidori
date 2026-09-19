package dispatch

import "context"

// Subtask is one independently dispatchable unit of a decomposed request.
// ContextRefs are opaque pointers (file paths, section ids) copied into the
// subtask prompt — not a live shared scratchpad across nodes.
type Subtask struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Prompt      string   `json:"prompt"`
	ContextRefs []string `json:"context_refs,omitempty"`
}

// DecomposeRequest carries the user-facing prompt to split. MaxSubtasks caps
// planner output (healthy-node count applied by the caller).
type DecomposeRequest struct {
	Prompt      string
	Model       string
	MaxSubtasks int
}

// Decomposer splits a request into subtasks for asymmetric hybrid dispatch.
// See docs/adr/ADR-0024-task-decomposition-hybrid.md.
type Decomposer interface {
	Decompose(ctx context.Context, req DecomposeRequest) ([]Subtask, error)
}

// NoOpDecomposer returns a single subtask wrapping the original prompt.
// v1 stub until the LLM planner lands (ADR-0024 step 4).
type NoOpDecomposer struct{}

func (NoOpDecomposer) Decompose(_ context.Context, req DecomposeRequest) ([]Subtask, error) {
	panic("fake")
}
