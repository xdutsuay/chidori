package acpclient

import (
	"encoding/json"
	"strings"

	"github.com/xdutsuay/lclreason/internal/agent"
)

// SessionUpdateResult is one parsed session/update notification.
type SessionUpdateResult struct {
	TurnEvent agent.TurnEvent
	HasEvent  bool
	ChunkText string // Grok agent_message_chunk delta
	IsFinal   bool
}

// MapToTurnEvent maps an ACP session/update notification into a native TurnEvent.
// FakeAgent tests use update.kind thought|final; real Grok Build uses
// update.sessionUpdate agent_message_chunk with update.content.text.
func MapToTurnEvent(method string, params json.RawMessage) (agent.TurnEvent, bool) { panic("fake") }

// ParseSessionUpdate decodes session/update params for FakeAgent and Grok shapes.
func ParseSessionUpdate(params json.RawMessage) (SessionUpdateResult, bool) { panic("fake") }

// Grok Build: sessionUpdate + content.text

// FakeAgent baseline: update.kind thought|final
