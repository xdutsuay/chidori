package mcpclient

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xdutsuay/lclreason/internal/executil"
)

// ToolDef is one tool an MCP server advertises via tools/list.
type ToolDef struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	InputSchema json.RawMessage `json:"inputSchema,omitempty"`
}

// ServerConfig is one configured MCP server (config.yaml's mcp_servers list).
type ServerConfig struct {
	Name    string   `json:"name"` // unique key, used for tool namespacing (see Manager)
	Command string   `json:"command"`
	Args    []string `json:"args,omitempty"`
}

// Client manages one MCP server subprocess over stdio. Mirrors internal/lsp.
// Client's lifecycle discipline exactly (including the running-flag bug
// that ADR-0006 caught and fixed): running only ever flips true *after* the
// initialize/initialized handshake completes, never right after spawn —
// getting this wrong once already cost a real debugging session, so this
// client is written correctly from the start rather than re-discovering it.
type Client struct {
	mu      sync.Mutex
	cmd     *exec.Cmd
	stdin   *bufio.Writer
	running bool
	lastErr error

	nextID  atomic.Int64
	pendMu  sync.Mutex
	pending map[int64]chan *envelope

	readDone chan struct{}
}

func New() *Client { panic("fake") }

func (c *Client) Status() (running bool, lastErr error) { panic("fake") }

// Start spawns the configured command and performs the MCP
// initialize/notifications-initialized handshake. Returns once the server
// is ready to accept tools/list — callers must not call ListTools/CallTool
// before Start returns successfully.
func (c *Client) Start(cfg ServerConfig) error { panic("fake") }

// stays false until the handshake below actually succeeds

func (c *Client) handshake(name string) error { panic("fake") }

// Stop terminates the server process, if running. Idempotent.
func (c *Client) Stop() { panic("fake") }

func (c *Client) waitForExit() { panic("fake") }

func (c *Client) setErr(err error) { panic("fake") }

type logWriter struct{ name string }

func (l logWriter) Write(p []byte) (int, error) { panic("fake") }

// ListTools calls tools/list.
func (c *Client) ListTools(ctx context.Context) ([]ToolDef, error) { panic("fake") }

// CallTool calls tools/call and returns the tool's text output (MCP's
// content blocks, concatenated — v1 only handles text content, which
// covers the overwhelming majority of MCP tools; image/binary content
// blocks are dropped, not a v1 goal).
func (c *Client) CallTool(ctx context.Context, name string, args map[string]any) (string, error) {
	panic("fake")
}

func (c *Client) call(ctx context.Context, method string, params any) (json.RawMessage, error) {
	panic("fake")
}

func (c *Client) notify(method string, params any) error { panic("fake") }

func (c *Client) reply(id *int64, result any) error { panic("fake") }

func (c *Client) write(env envelope) error { panic("fake") }

// readLoop dispatches by envelope shape. Unlike gopls, MCP servers rarely
// send server->client requests (an ID plus a method) — but if one shows up
// (e.g. a sampling/* request from a server that wants the client to run an
// LLM call, which chidori doesn't support in v1), reply with an empty
// result rather than never answering, so a server that expects a reply can
// never hang this client waiting for one.
func (c *Client) readLoop(r *bufio.Reader) { panic("fake") }
