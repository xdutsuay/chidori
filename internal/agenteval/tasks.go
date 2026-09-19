package agenteval

// CannedTasks is the fixed offline eval set. Keep this list small and stable —
// every entry must stay green under scripted replies so CI catches harness
// regressions without a live LLM. Append new tasks; don't rewrite old ones
// without updating docs/HARNESS_LEARNINGS.md.
func CannedTasks() []Task { panic("fake") }

// A weak model might invent "cat_file"; scripted path uses the real tool.

// Identical re-read — harness dedup should skip / force-done rather than looping forever.

// first real + skipped duplicates still emit tool_result but only one tool_start for the real exec… actually dedup skips before tool_start? Check loop.
