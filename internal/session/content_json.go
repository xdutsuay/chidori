package session

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strings"
)

// Content block types for ide_session_messages.content_json (KMA-217).
const (
	ContentBlockTypeText       = "text"
	ContentBlockTypeToolUse    = "tool_use"
	ContentBlockTypeToolResult = "tool_result"
)

// ContentBlock is one OpenAI/Cursor-like content unit.
type ContentBlock struct {
	Type       string           `json:"type"`
	Text       string           `json:"text,omitempty"`
	ToolUse    *ToolUseBlock    `json:"tool_use,omitempty"`
	ToolResult *ToolResultBlock `json:"tool_result,omitempty"`
}

// ToolUseBlock is an assistant tool_use content block.
type ToolUseBlock struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// ToolResultBlock is a tool role tool_result content block.
type ToolResultBlock struct {
	ToolCallID string `json:"tool_call_id"`
	Name       string `json:"name,omitempty"`
	Content    string `json:"content"`
}

// NormalizeMessageRole maps UI roles onto the wire/store set.
func NormalizeMessageRole(role string) string { panic("fake") }

// RenderStructuredMessageHTML derives display HTML from content_json blocks.
func RenderStructuredMessageHTML(blocks []ContentBlock) string { panic("fake") }

var toolCardNameRe = regexp.MustCompile(`(?i)<div\s+class="tool-card">\s*([^<]+?)\s*</div>`)

// BackfillContentJSONFromHTML best-effort converts legacy html into blocks.
// Tool call IDs are not recoverable from HTML alone.
func BackfillContentJSONFromHTML(rawHTML string) ([]ContentBlock, error) { panic("fake") }

func encodeContentJSON(blocks []ContentBlock) (string, error) { panic("fake") }
