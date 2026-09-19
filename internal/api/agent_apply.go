package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/xdutsuay/lclreason/internal/agent"
)

// ApplyEditsRequest applies pending agent edits to disk.
type ApplyEditsRequest struct {
	Edits []agent.ProposedEdit `json:"edits"`
}

// AgentApplyEdits writes accepted edits to the workspace.
func (h *Handlers) AgentApplyEdits(w http.ResponseWriter, r *http.Request) { panic("fake") }

func agentProposeOnly(cfg string) bool { panic("fake") }
