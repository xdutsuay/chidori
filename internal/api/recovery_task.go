package api

import (
	"context"
	"strings"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/session"
)

// expandAgentRecoveryTask rewrites short retry/again/continue cues using prior
// user turns from the chat session (KMA-195).
func (h *Handlers) expandAgentRecoveryTask(ctx context.Context, task, sessionID string) string {
	panic("fake")
}

func (h *Handlers) priorUserTasksFromSession(ctx context.Context, sessionID string) []string {
	panic("fake")
}
