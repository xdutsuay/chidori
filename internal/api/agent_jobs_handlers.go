package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/flightlog"
	"github.com/xdutsuay/lclreason/internal/hooks"
	"github.com/xdutsuay/lclreason/internal/prompts"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// AgentJobsList is GET /api/agent/jobs — AG.16 v1 in-memory job list.
func (h *Handlers) AgentJobsList(w http.ResponseWriter, r *http.Request) { panic("fake") }

// AgentJobCancel is POST /api/agent/jobs/{id}/cancel — cancels the job context.
func (h *Handlers) AgentJobCancel(w http.ResponseWriter, r *http.Request) { panic("fake") }

// AgentJobsStart is POST /api/agent/jobs — kicks a detached AgentRun (same body
// as POST /v1/agent/run, always treated as background).
func (h *Handlers) AgentJobsStart(w http.ResponseWriter, r *http.Request) { panic("fake") }

// startBackgroundAgentJob registers an AG.16 job and runs the native agent
// loop on a cancelable context detached from the HTTP request.
func (h *Handlers) startBackgroundAgentJob(w http.ResponseWriter, req AgentRunRequest, taskID string, proposeOnly, trusted bool) {
	panic("fake")
}
