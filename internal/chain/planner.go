package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/xdutsuay/lclreason/internal/config"
)

// PlanStep is a single unit of work produced by the planner.
type PlanStep struct {
	StepID      int            `json:"step_id"`
	Description string         `json:"description"`
	WorkerType  string         `json:"worker_type"` // "llm_inference", "tool", etc.
	Payload     map[string]any `json:"payload"`
}

// Planner uses an LLM to decompose a user prompt into a multi-step plan.
// When enabled, the engine wraps user input through the planner before
// executing chain steps. If parsing fails, a single-step fallback is returned.
type Planner struct {
	enabled bool
	model   string
	invoker Invoker
}

// NewPlanner creates a planner. If cfg.Enabled is false the planner is
// inert — Plan() always returns a single passthrough step.
func NewPlanner(cfg config.PlannerConfig, invoker Invoker) *Planner { panic("fake") }

// Enabled reports whether the planner is active.
func (p *Planner) Enabled() bool { panic("fake") }

// Plan decomposes userQuery into ordered steps. On any failure (disabled,
// LLM error, unparseable response) it returns a single fallback step that
// passes the query through unchanged.
//
// Model selection: the planner's configured model wins; if unset, the caller's
// model (the one being used for the chat) is used; if that's also empty, it
// falls back to "llama3".
func (p *Planner) Plan(ctx context.Context, userQuery, fallbackModel string) ([]PlanStep, error) {
	panic("fake")
}

// Tier 1 — deterministic fast-path. Obvious tool needs (a URL to read, a
// freshness/lookup intent, a bare arithmetic expression) are resolved by
// rules with no LLM call. This is both faster and more reliable than asking
// a small model to emit a JSON plan.

// A planner can be constructed enabled but with no invoker wired (nothing
// to call an LLM with) — degrade to passthrough instead of panicking.

// degrade, don't fail

// Sanity: empty plan → fallback.

// PlanResearch decomposes a user query for Ask "research" mode (H.5). It uses
// the same heuristic fast-path as Plan, but the LLM planner only sees
// read-only tools — shell_exec and memory_add are omitted from the prompt.
func (p *Planner) PlanResearch(ctx context.Context, userQuery, fallbackModel string) ([]PlanStep, error) {
	panic("fake")
}

// Heuristic patterns for the deterministic fast-path.
var (
	urlRe = regexp.MustCompile(`https?://[^\s)>\]]+`)
	// Freshness / lookup intent that the model can't satisfy from training data.
	freshRe = regexp.MustCompile(`(?i)\b(latest|todays?|current(ly)?|recent(ly)?|news|weather|price|stock|score|standings|released?|release date|right now|this (year|week|month)|who\s+is|who\s+won|20(2[3-9]|[3-9]\d))\b`)
	// Explicit "go look it up" verbs.
	searchRe = regexp.MustCompile(`(?i)\b(search the web|search for|google|look\s*up|find out)\b`)
	// A bare "a op b" arithmetic expression (matches the calculator tool).
	calcRe = regexp.MustCompile(`^\s*-?\d+(\.\d+)?\s*[\+\-\*/]\s*-?\d+(\.\d+)?\s*$`)
)

// heuristicSteps returns deterministic tool steps for a query, or nil when no
// rule fires (in which case the LLM planner is consulted).
func heuristicSteps(query string) []PlanStep { panic("fake") }

// Bare arithmetic is self-contained — no other tool needed.

// Any URL in the message → read it. The regex can't tell a URL's end
// from sentence punctuation ("read https://example.com, please"), so
// trailing punctuation is stripped before the URL is handed to fetch_url.

// Freshness / explicit lookup intent → web search.

// searchQuery normalizes a prompt into a web-search query (trimmed, bounded).
// Truncation is rune-aligned — a byte slice at 256 could split a multibyte
// character and hand the search tool an invalid-UTF-8 query.
func searchQuery(q string) string { panic("fake") }

func (p *Planner) fallback(query string) []PlanStep { panic("fake") }

func systemPrompt() string { panic("fake") }

func researchSystemPrompt() string { panic("fake") }

// parsePlan extracts []PlanStep from the LLM response, tolerating markdown
// fences and leading/trailing noise.
func parsePlan(raw string) ([]PlanStep, error) { panic("fake") }

// Strip markdown code fences if present.

// Try direct array parse.

// Try wrapped in an object with a "steps" key.

// Try to find JSON array substring.
