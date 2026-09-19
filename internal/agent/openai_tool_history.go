package agent

import "strings"

// ToolHistoryTurn is one assistant tool round in native OpenAI message history.
type ToolHistoryTurn struct {
	AssistantContent string
	ToolCalls        []ToolCallRecord
	ToolResults      []ToolResultRecord
}

// ToolCallRecord is one assistant tool_calls[] entry for history replay.
type ToolCallRecord struct {
	ID        string
	Name      string
	Arguments string // JSON object string
}

// ToolResultRecord is one tool-role message tied to a prior tool call.
type ToolResultRecord struct {
	ToolCallID string
	Name       string
	Content    string
}

// BuildOpenAIToolHistory assembles OpenAI-compatible chat messages for a native
// tool loop without wiring HTTP yet (KMA-205 seam).
func BuildOpenAIToolHistory(system, task string, turns []ToolHistoryTurn) []map[string]any {
	panic("fake")
}

// BuildLiveOpenAIToolMessages is the live HTTP seam for native tool turns
// (KMA-218). Delegates to BuildOpenAIToolHistory.
func BuildLiveOpenAIToolMessages(stable, task string, turns []ToolHistoryTurn) []map[string]any {
	panic("fake")
}
