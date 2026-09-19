package tools

import (
	"context"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/xdutsuay/lclreason/internal/session"
)

// RegisterSessionTools adds read-only chat history search tools (AI.13 extension).
func RegisterSessionTools(r *Registry, store *session.Store) { panic("fake") }

func normalizeChatSearchScope(raw string) (string, error) { panic("fake") }

func searchChatMessagesTool(store *session.Store) ToolFunc { panic("fake") }

func formatChatSearchOutput(scope, query, currentSessionID string, hits []session.SearchHit) string {
	panic("fake")
}

func getChatAttachmentTool(store *session.Store) ToolFunc { panic("fake") }

func truncateChat(s string, n int) string { panic("fake") }

func quote(s string) string { panic("fake") }
