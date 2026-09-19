package mcpclient

import (
	"context"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/tools"
)

// callTimeout bounds every tools/call — these back agent tool calls, which
// already run inside the agent loop's own turn budget; fail fast rather
// than hold a turn open indefinitely on a stuck MCP server.
const callTimeout = 30 * time.Second

// ServerStatus is what the Settings -> MCP panel shows per configured server.
type ServerStatus struct {
	Name    string `json:"name"`
	Running bool   `json:"running"`
	Error   string `json:"error,omitempty"`
	Tools   int    `json:"tools"`
}

// Manager owns one Client per configured MCP server and keeps a
// *tools.Registry's set of MCP-sourced tools in sync with them — same
// "reconcile configured vs running" shape as syncLSPWorkspace() in
// internal/api/lsp_handlers.go.
type Manager struct {
	mu         sync.Mutex
	clients    map[string]*Client // keyed by ServerConfig.Name
	toolReg    *tools.Registry
	registered map[string][]string // server name -> tool names it registered, so Reconcile can cleanly deregister on removal
}

func NewManager(toolReg *tools.Registry) *Manager { panic("fake") }

// Reconcile starts clients for newly configured servers, stops+removes ones
// no longer configured, and refreshes each running server's tool list into
// toolReg. Runs in a goroutine by the caller (server startup is real
// subprocess + handshake latency per server) — never block an HTTP response
// on this.
func (m *Manager) Reconcile(ctx context.Context, cfgs []ServerConfig) { panic("fake") }

// Stop and deregister servers no longer configured.

// Start (or refresh tools for) each configured server.

// drop this server's old tool set before re-registering the current one

// namespaced — two servers can both expose e.g. "search"

// GATE: MCP tool calls go through the exact same agent-tool
// trust gate as any built-in tool (run_terminal_command,
// git_commit, ...) — the caller (agent loop) applies that
// gate uniformly to every registry entry, MCP-sourced or
// not. No special-casing here, deliberately.

// deregisterLocked removes a server's previously-registered tools from
// toolReg. Caller must hold m.mu. Tools are namespaced as "server/tool"
// so Unregister cannot collide with builtins or another server's tools.
func (m *Manager) deregisterLocked(server string) { panic("fake") }

// Statuses reports each configured server's current state, for the Settings
// -> MCP panel.
func (m *Manager) Statuses() []ServerStatus { panic("fake") }

// StopAll stops every managed server — called on coordinator shutdown.
func (m *Manager) StopAll() { panic("fake") }

// ListedTool is one MCP-sourced tool name for the Settings tools browser.
type ListedTool struct {
	Server      string `json:"server"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description,omitempty"`
}

// ListTools returns tools currently registered from connected MCP servers.
func (m *Manager) ListTools() []ListedTool { panic("fake") }
