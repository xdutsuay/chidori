package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/hooks"
	"github.com/xdutsuay/lclreason/internal/prompts"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// AgentResumeRequest continues a paused agent run after chidori-questions (H.3).
type AgentResumeRequest struct {
	TaskID   string `json:"task_id"`
	Answers  string `json:"answers"`
	Stream   bool   `json:"stream"`
	ClientID string `json:"client_id,omitempty"`
}

// AgentResume executes the next leg of a paused agent run.
func (h *Handlers) AgentResume(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) agentResumeExecution(w http.ResponseWriter, r *http.Request, ctx context.Context, cancel context.CancelFunc, req AgentResumeRequest, taskID string, cp agent.Checkpoint) {
	panic("fake")
}
