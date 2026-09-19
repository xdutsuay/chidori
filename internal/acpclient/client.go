package acpclient

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/executil"
)

// Client speaks ACP (JSON-RPC over stdio NDJSON).
//
// v1 rule (mirrors internal/lsp): running only flips true after initialize
// completes successfully.
type Client struct {
	mu      sync.Mutex
	running bool
	lastErr error

	transport *stdioTransport
	stdin     io.Writer
	stdout    *bufio.Reader

	nextID atomic.Int64
	pendMu sync.Mutex
	pend   map[int64]chan *rpcEnvelope

	readDone chan struct{}

	// one active prompt at a time in v1
	onEventMu sync.Mutex
	onEvent   EventHandler

	// cancel the active prompt by sending session/cancel (adapter-defined).
	activeSessionIDMu sync.Mutex
	activeSessionID   string

	closed chan struct{}

	workDir string

	writeMu sync.Mutex

	// v1 callbacks for answering ACP fs/* requests.
	fsReadMu    sync.Mutex
	fsReadFunc  func(ctx context.Context, path string) (string, error)
	fsWriteMu   sync.Mutex
	fsWriteFunc func(ctx context.Context, path string, content string) error

	promptStreamMu sync.Mutex
	promptStream   *promptStreamState
}

func NewClient() *Client { panic("fake") }

// SetFSHandlers registers handlers for ACP filesystem requests:
// - fs/read_text_file
// - fs/write_text_file
func (c *Client) SetFSHandlers(
	read func(ctx context.Context, path string) (string, error),
	write func(ctx context.Context, path string, content string) error,
) {
	panic("fake")
}

func (c *Client) Status() (running bool, lastErr error) { panic("fake") }

// Start spawns the external binary (or uses TestTransport), and performs the
// ACP initialize handshake. Returns once initialize/response succeeds.
func (c *Client) Start(cfg SpawnConfig) error { panic("fake") }

// keep stderr off the ACP stream

// wire transport

func (c *Client) initialize(ctx context.Context) (json.RawMessage, error) {
	panic(
		// Grok Build expects clientCapabilities; capabilities kept for FakeAgent compat.
		"fake")
}

func (c *Client) authenticate(ctx context.Context, initResult json.RawMessage) error { panic("fake") }

// Close terminates the client transport and cancels any in-flight calls.
func (c *Client) Close() error { panic("fake") }

// LookGrok resolves the grok binary on PATH (or validates a configured path)
// and probes its version as best-effort.
func LookGrok(cfg LookGrokConfig) (BinaryInfo, error) { panic("fake") }

// newPromptSession creates a new ACP session (session/new).
func (c *Client) NewSession(ctx context.Context) (string, error) { panic("fake") }

// Prompt runs session/prompt, maps streamed notifications to TurnEvents, waits
// for Grok-style chunk streams to settle, then emits a synthetic done if needed.
func (c *Client) Prompt(ctx context.Context, sessionID string, task string, onEvent EventHandler) error {
	panic("fake")
}

// Cancel attempts to cancel the currently-active session by sending session/cancel.
func (c *Client) Cancel(ctx context.Context) error { panic("fake") }

// Cancellation may be implemented as request or notification by the agent.
// For v1, send as request and ignore failures.

func (c *Client) call(ctx context.Context, method string, params map[string]any) (json.RawMessage, error) {
	panic("fake")
}

func mustMarshal(v any) json.RawMessage { panic("fake") }

func (c *Client) decodeLoop() { panic("fake") }

// server request (agent asks the client for some capability)

// unknown server request: ignore for now

// response to our call

// notification

func (c *Client) respondRPC(id int64, result json.RawMessage, rpcErr *rpcError) error { panic("fake") }
