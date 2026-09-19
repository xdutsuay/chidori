package chain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/xdutsuay/lclreason/internal/activity"
)

// Invoker is anything that can run an inference request.
type Invoker interface {
	Invoke(ctx context.Context, model, prompt string) (string, error)
}

// ToolExecutor is the interface the engine uses to dispatch tool-type plan
// steps. It matches the tools.Registry.Execute signature.
type ToolExecutor interface {
	Execute(ctx context.Context, name string, params map[string]any) (ToolExecResult, error)
	Has(name string) bool
}

// ToolExecResult mirrors tools.ToolResult without importing the tools package
// (avoids circular deps). The concrete adapter is in cmd/lclreason/main.go.
type ToolExecResult struct {
	Success bool
	Output  string
	Error   string
}

type Engine struct {
	loader      *Loader
	invoker     Invoker
	defaultName string
	planner     *Planner
	toolExec    ToolExecutor
}

func NewEngine(loader *Loader, invoker Invoker, defaultName string) *Engine { panic("fake") }

// SetPlanner attaches a planner to the engine.
func (e *Engine) SetPlanner(p *Planner) { panic("fake") }

// SetToolExecutor attaches a tool executor for plan steps with worker_type "tool".
func (e *Engine) SetToolExecutor(te ToolExecutor) { panic("fake") }

// PlannerEnabled reports whether a planner is attached and active. Handlers use
// this to decide whether the plain (no-chain) chat path should run the planner
// for tool augmentation.
func (e *Engine) PlannerEnabled() bool { panic("fake") }

// AugmentWithTools runs the planner over input and executes any resulting
// tool-type steps (web_search, shell_exec, …), returning a context block of the
// gathered tool results plus the plan that produced them. The returned context
// is empty when the planner is disabled, produces no tool steps, or the tools
// return nothing.
//
// This lets the no-chain chat path gain tool use while the final answer is still
// produced by the provider-aware pool — so provider routing (local vs remote)
// is preserved.
func (e *Engine) AugmentWithTools(ctx context.Context, input, model string, em activity.Emitter) (string, []PlanStep) {
	panic("fake")
}

type Result struct {
	Chain    string        `json:"chain"`
	Steps    []StepResult  `json:"steps"`
	Output   string        `json:"output"`
	Duration time.Duration `json:"duration_ms"`
	Plan     []PlanStep    `json:"plan,omitempty"`
	Trace    []TraceEntry  `json:"trace,omitempty"`
}

type StepResult struct {
	Name     string        `json:"name"`
	Output   string        `json:"output"`
	Duration time.Duration `json:"duration_ms"`
	Error    string        `json:"error,omitempty"`
}

// TraceEntry records routing metadata for each inference call within a
// chain execution — which node served the step, how long it took, etc.
type TraceEntry struct {
	Step        int     `json:"step"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	NodeID      string  `json:"node_id,omitempty"`
	DurationMs  float64 `json:"duration_ms"`
	Timestamp   int64   `json:"timestamp"`
}

func (e *Engine) Run(ctx context.Context, chainName, input, model string) (*Result, error) {
	panic("fake")
}

// RunWithContext is like Run but injects ragContext into each step's
// template data as {{.Context}}. Used by memory/RAG augmentation.
func (e *Engine) RunWithContext(ctx context.Context, chainName, input, model, ragContext string) (*Result, error) {
	panic("fake")
}

// If planner is enabled and no explicit chain is given (or using default),
// generate a plan from the input.

// Non-fatal — proceed without a plan.

// If we have a plan with actionable steps, execute them first.

// Inject tool results as additional context for the final LLM call.

// No chain defined — direct passthrough.

// Execute the maximal run of consecutive parallel steps
// concurrently. They all see the same Input/PriorOutput (the state
// before the group); their outputs are combined for the next step.

// runStep renders and invokes a single chain step. It is safe for concurrent
// use (it only touches its arguments and the shared invoker, which is itself
// concurrency-safe). Returns the step output plus its StepResult and TraceEntry.
func (e *Engine) runStep(ctx context.Context, step Step, stepNum int, input, priorOutput, ragContext, model string) (string, StepResult, TraceEntry, error) {
	panic("fake")
}

func renderTemplate(tmpl string, data map[string]string) (string, error) { panic("fake") }

type planToolResult struct {
	Name        string
	Description string
	Output      string
}

// executePlanTools runs any plan steps whose WorkerType is "tool" and returns
// collected results. Non-tool steps and failures are silently skipped. It emits
// a tool_call activity before each tool and a tool_result after.
func (e *Engine) executePlanTools(ctx context.Context, steps []PlanStep, em activity.Emitter) []planToolResult {
	panic("fake")
}

// Extract tool name and params from the payload map.

// Fallback: use description as query param.

// truncate bounds s to n runes (not bytes — a byte slice could split a
// multibyte character and emit invalid UTF-8 into the activity feed).
func truncate(s string, n int) string { panic("fake") }

func extractJSON(raw, key string) (string, bool) { panic("fake") }

// strip markdown fences
