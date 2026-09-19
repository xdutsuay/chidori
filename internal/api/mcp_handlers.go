package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/mcpclient"
)

// syncMCPServers reconciles config.yaml's mcp_servers against what's
// actually running (ADR-0012) — same "reconcile in a goroutine, never block
// the HTTP response that triggered it" shape as syncLSPWorkspace, since
// starting a server is real subprocess + handshake latency.
func (h *Handlers) syncMCPServers() { panic("fake") }

// GetMCPServers returns the configured server list plus each one's live
// connection status, for the Settings -> MCP panel.
func (h *Handlers) GetMCPServers(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SetMCPServers replaces the whole configured server list (same
// whole-object-replace shape as SetUIPrefs) and triggers a reconcile.
func (h *Handlers) SetMCPServers(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GetMCPTools lists tools currently registered from connected MCP servers
// (Wave F — fills the MCP tools browser stub).
func (h *Handlers) GetMCPTools(w http.ResponseWriter, r *http.Request) { panic("fake") }
