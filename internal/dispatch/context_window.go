package dispatch

import (
	"log"
	"strings"
	"unicode/utf8"
)

// Context-window handling for the hybrid flow.
//
// The problem this solves: OpenAI-compatible providers (NVIDIA NIM included)
// do NOT expose a model's context window at runtime — GET /v1/models has no
// max_model_len field and there's no metadata endpoint (confirmed against
// NVIDIA's docs). So if the assembled prompt (history + agent scratch + RAG +
// file context) is larger than the model can accept, the request either fails
// with "maximum context length is N … your request has M tokens" or — worse —
// the prefill is so large the provider never returns headers before our 300s
// client timeout ("context deadline exceeded"). Both were seen in real use.
//
// Since we can't discover the limit, we (1) keep a best-effort map of known
// model families → context window, and (2) hard-clamp the outbound prompt to
// `window − outputReserve` at the single chokepoint every chat call passes
// through. This is a safety net, not a replacement for compaction: normal
// prompts are already small; this only trims the pathological oversized ones.

// estCharsPerToken mirrors compact.EstimateTokens' ~4 chars/token heuristic so
// token budgets translate to a character budget without a real tokenizer.
const estCharsPerToken = 4

// ContextWindow returns a best-effort maximum context (in tokens) for a model
// id, matched by substring against known families. Unknown models get a
// conservative modern default. Deliberately errs toward *smaller* windows for
// ambiguous matches so the clamp never over-estimates and lets an overflow
// through.
func ContextWindow(model string) int { panic("fake") }

// unknown/auto — stay conservative

// Nemotron-3-Ultra: up to 1M, 256k typical deployment default

// some variants advertise far more; stay safe

// original Llama 3 (non-.1) is 8k

// NVIDIA NIM hosted keys (e.g. topbar "nvidia-7030" with model left
// on auto) — typical deploy window is 128k; prefer this over the
// compaction default of 8192 when the UI asks for a provider-aware budget.

// reasonable modern floor

// outputReserve returns how many tokens to hold back from the input budget for
// the model's own response, scaled so tiny windows still leave *some* input
// room. Capped at 4096 for large windows — OpenRouter credit checks reject
// requests when max_tokens exceeds what the account can afford (often well
// below 8192 on capped keys), which aborted agent turns before any tool call.
func outputReserve(window int) int { panic("fake") }

// alignRuneBoundary returns the nearest byte index <= idx that falls on a
// UTF-8 rune boundary (running_issue.md P2-4 / H.7). Go string slicing
// operates on raw bytes, so `prompt[:idx]` or `prompt[idx:]` at an
// arbitrary index can land inside a multi-byte rune's continuation bytes
// (0x80-0xBF, which utf8.RuneStart reports false for) — invalid UTF-8 that
// later gets mangled to U+FFFD on any JSON round-trip. Only matters for
// prompts big enough to actually hit the clamp below; normal-sized prompts
// never reach this function's truncation branches at all.
func alignRuneBoundary(s string, idx int) int { panic("fake") }

// clampPromptToModel trims prompt to fit model's context window (minus an
// output reserve), keeping the head (system rules / instructions) and the tail
// (the actual user question) and dropping the middle bulk. Returns the possibly
// trimmed prompt and whether it was trimmed. The chokepoint callers use this so
// no oversized prompt ever reaches the provider.
func clampPromptToModel(prompt, model string) (string, bool) { panic("fake") }

// keep more of the head (instructions/context)

// FitPromptToBudget clamps prompt to the model's input window (head+tail).
// Ask/Agent preflight uses this so oversized @file context never hard-400s.
func FitPromptToBudget(prompt, model string) (fitted string, truncated bool) { panic("fake") }
