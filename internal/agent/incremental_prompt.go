package agent

import "strings"

// IncrementalPromptSplit marks the boundary between a stable (cacheable) system
// prefix and the per-turn volatile user suffix inside a single prompt string.
// OpenAI-compat chat paths split on this marker into system + user messages
// (KMA-163). Local backends flatten it back to plain text.
const IncrementalPromptSplit = "\n<<<CHIDORI_INCREMENTAL>>>\n"

// JoinIncrementalPrompt concatenates stable + volatile for Invoker transport.
func JoinIncrementalPrompt(stable, volatile string) string { panic("fake") }

// SplitIncrementalPrompt returns the stable/volatile halves when the marker is present.
func SplitIncrementalPrompt(prompt string) (stable, volatile string, ok bool) { panic("fake") }

// FlattenIncrementalPrompt removes the wire marker for backends that take a
// single text blob (Ollama, etc.).
func FlattenIncrementalPrompt(prompt string) string { panic("fake") }

// agentPromptParts is the H1 incremental-transport split (KMA-163):
// Stable is byte-identical across turns when tools/workspace/task inputs match;
// Volatile carries turn number + scratch (and other per-turn context).
type agentPromptParts struct {
	Stable   string
	Volatile string
}

// join returns the wire form used by Invoker.
func (p agentPromptParts) join() string { panic("fake") }
