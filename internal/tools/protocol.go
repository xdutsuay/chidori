package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// ToolCall represents a requested tool invocation parsed from an LLM response.
type ToolCall struct {
	ID     string         `json:"id,omitempty"`
	Name   string         `json:"name"`
	Params map[string]any `json:"params"`
}

// ToolProtocol abstracts formatting tools for LLMs and parsing tool call responses.
type ToolProtocol interface {
	FormatTools(tools []Tool) any
	ParseToolCalls(response string) ([]ToolCall, error)
}

var (
	jsonBlockRegex = regexp.MustCompile("(?s)```(?:json)?\\s*(\\{.*?\\})\\s*```")
	xmlBlockRegex  = regexp.MustCompile("(?s)<tool_call>\\s*(.*?)\\s*</tool_call>")
)

// OpenAIToolProtocol handles native function/tool calling for OpenAI-compatible models.
type OpenAIToolProtocol struct{}

func (p *OpenAIToolProtocol) FormatTools(tools []Tool) any { panic("fake") }

func (p *OpenAIToolProtocol) ParseToolCalls(response string) ([]ToolCall, error) {
	panic(
		// For native OpenAI mode, JSON parsing fallback if raw text was returned instead of structured calls
		"fake")
}

// LocalModelToolProtocol handles robust tool call parsing for local models (Ollama, vLLM, LM Studio).
type LocalModelToolProtocol struct{}

func (p *LocalModelToolProtocol) FormatTools(tools []Tool) any { panic("fake") }

func (p *LocalModelToolProtocol) ParseToolCalls(response string) ([]ToolCall, error) {
	panic(
		// Try parsing XML <tool_call> tags (Hermes format)
		"fake")
}

// Fallback to JSON plan extraction

// parseJSONPlan extracts tool calls from structured JSON plan text or markdown fences.
func parseJSONPlan(text string) ([]ToolCall, error) { panic("fake") }

// Try parsing single tool call JSON directly

// GetToolProtocol returns the appropriate protocol for a provider or model type.
func GetToolProtocol(providerType string) ToolProtocol { panic("fake") }
