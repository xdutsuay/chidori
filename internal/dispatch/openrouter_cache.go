package dispatch

import (
	"encoding/json"
	"strings"
)

// chatContentPart is one OpenAI/Anthropic-style content block with optional cache_control.
type chatContentPart struct {
	Type         string `json:"type"`
	Text         string `json:"text,omitempty"`
	CacheControl *struct {
		Type string `json:"type"`
	} `json:"cache_control,omitempty"`
}

func messageContentString(m chatMessage) string { panic("fake") }

func contentAsString(c any) string { panic("fake") }

// applyOpenRouterSystemCache attaches ephemeral cache_control to a stable system
// message when the remote is OpenRouter and the first message is a system message
// with non-empty string content, regardless of total message count (KMA-213,
// KMA-163 slice 2). Assistant/tool messages keep their original shape and are not
// rewritten into content-part arrays.
func applyOpenRouterSystemCache(msgs []chatMessage, meta OpenAICompatMeta) []chatMessage {
	panic("fake")
}

// marshalChatMessages JSON-encodes chat messages for request bodies.
func marshalChatMessages(msgs []chatMessage, meta OpenAICompatMeta) ([]byte, error) { panic("fake") }

func buildChatRequestBody(model string, msgs []chatMessage, meta OpenAICompatMeta, stream bool, tools []ChatTool) ([]byte, error) {
	panic("fake")
}
