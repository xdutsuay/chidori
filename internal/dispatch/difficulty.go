package dispatch

import "strings"

// Difficulty-aware hybrid routing — borrowed from RouteLLM/Hybrid-LLM
// (llpcodefeature_Cursorclone.md §13.2 #1, "highest leverage" idea from the
// distributed-inference design study). ExecuteHybrid used to always race
// local AND remote for every request; that's correct for hard prompts but
// wasteful for easy ones — the local leg would have answered fine, and the
// remote call is pure cost (an API charge, or just cycles on a free key)
// for no quality gain.
//
// This is a cheap rule-based classifier, not RouteLLM's trained win-predictor
// (an embedding/learned model is real future work, not something to bring in
// as a new dependency for a first cut) — but the research is explicit that
// router cost is a non-issue at any tier (<1ms rule-based, ~5ms embedding, vs
// 500-2000ms of model time), so even this simple version is free to run on
// every hybrid call. It deliberately errs toward "hard" (race both, the
// original behavior) whenever a prompt shows any sign of real complexity —
// wrongly racing an easy prompt costs a redundant call; wrongly skipping the
// remote leg on a hard prompt costs answer quality, which is the worse trade.
const (
	difficultyEasy = "easy"
	difficultyHard = "hard"
)

// easyPromptCharLimit: prompts longer than this are treated as hard even with
// no other signal — length alone correlates with multi-part/detailed asks.
const easyPromptCharLimit = 160

// hardSignals: any of these substrings (case-insensitive) flips the
// classification to hard regardless of length. Short but names a
// code/debug/design task ("why is this crashing") is still hard.
var hardSignals = []string{
	"```", "def ", "func ", "class ", "```go", "```py", "```js", // code content
	"debug", "refactor", "architect", "design", "optimize", "explain in detail",
	"why does", "why is", "how does", "step by step", "step-by-step",
	"analyze", "prove", "algorithm", "compare", "trade-off", "tradeoff",
	"write a", "implement", "generate a",
}

// classifyDifficulty returns "easy" or "hard" for a prompt using cheap,
// deterministic heuristics — no model call, no network, negligible cost.
//
// Ported verbatim to JS as classifyDifficultyJS in
// internal/appui/public/js/08-debug-search-diff.js (ADR-0008, "draft &
// upgrade" mode) — keep both copies in sync if this changes.
func classifyDifficulty(prompt string) string { panic("fake") }

// multiple questions in one message
