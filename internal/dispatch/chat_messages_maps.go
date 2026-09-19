package dispatch

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ChatMessagesFromMaps converts OpenAI-style message maps (from agent
// BuildOpenAIToolHistory) into dispatch chatMessage values (KMA-205/218).
func ChatMessagesFromMaps(msgs []map[string]any) []chatMessage { panic("fake") }

func stringField(m map[string]any, key string) string { panic("fake") }

func toolCallsFromAny(raw any) []ToolCall { panic("fake") }
