// Package debugger wraps a `dlv dap` subprocess as a Debug Adapter Protocol
// (DAP) client, giving the rest of chidori a small Go-shaped API (Start,
// SetBreakpoints, Continue, StepOver, StepIn, StepOut, StackTrace, Scopes,
// Variables, Terminate) instead of hand-rolling the DAP wire protocol at
// every call site (BE.15 / Run & Debug panel v1). Wire framing and message
// types come from github.com/google/go-dap — the same
// Content-Length-prefixed-JSON framing internal/lsp already implements by
// hand for gopls, here provided by a library since DAP has many more message
// shapes than the handful of LSP calls this project makes.
//
// Scoped deliberately to a single active session at a time (same "one thing
// at a time" precedent as internal/agent's single-threaded loop) and launch
// mode only (`dlv debug <package-dir>`) — attach-to-running-process and
// multi-session debugging are both out of scope for v1.
package debugger

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xdutsuay/lclreason/internal/executil"

	dap "github.com/google/go-dap"
)

// Event is the simplified, UI-facing shape of a DAP event — a small,
// stable set of fields the frontend actually needs, rather than exposing
// every raw DAP event type and body shape to callers outside this package.
type Event struct {
	Type      string `json:"type"` // "stopped" | "continued" | "output" | "terminated" | "exited" | "error"
	Reason    string `json:"reason,omitempty"`
	ThreadID  int    `json:"thread_id,omitempty"`
	Text      string `json:"text,omitempty"`
	ExitCode  int    `json:"exit_code,omitempty"`
	IsErr     bool   `json:"is_error,omitempty"`
	Timestamp int64  `json:"-"`
}

// Breakpoint is the file:line shape used both for the arguments to
// SetBreakpoints and (via Verified) the actual location dlv accepted —
// dlv can move a breakpoint to the nearest executable line, so Verified/Line
// in the result can legitimately differ from what was requested.
type Breakpoint struct {
	Line     int    `json:"line"`
	Verified bool   `json:"verified"`
	Message  string `json:"message,omitempty"`
}

// StackFrame mirrors dap.StackFrame's fields the UI actually needs.
type StackFrame struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	File string `json:"file"`
	Line int    `json:"line"`
}

// Scope mirrors dap.Scope's fields the UI actually needs.
type Scope struct {
	Name               string `json:"name"`
	VariablesReference int    `json:"variables_reference"`
}

// Variable mirrors dap.Variable's fields the UI actually needs.
type Variable struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	Type               string `json:"type,omitempty"`
	VariablesReference int    `json:"variables_reference"`
}

// Session is one running debug session: a `dlv dap` subprocess plus the DAP
// client state (pending request/response correlation, breakpoints,
// current-stopped-thread tracking) needed to drive it.
type Session struct {
	cmd  *exec.Cmd
	conn io.ReadWriteCloser

	seq     int64
	mu      sync.Mutex
	pending map[int]chan dap.Message

	events       chan Event
	eventsMu     sync.Mutex
	eventsClosed bool
	done         chan struct{}
	once         sync.Once

	stateMu         sync.Mutex
	stoppedThreadID int
	isStopped       bool
	terminated      bool

	breakpointsMu sync.Mutex
	breakpoints   map[string][]int // file -> requested lines, resent whenever setBreakpoints is called again
}

// dialTimeout bounds how long Start waits for `dlv dap` to print its
// listen address and accept a connection — dlv itself starts almost
// instantly (it's the later `dlv debug` build step, triggered by the async
// launch request, that can take a few seconds for a real package).
const dialTimeout = 10 * time.Second

// Start launches `dlv dap` as a subprocess, connects to it as a DAP client,
// and asynchronously drives it through initialize -> launch (mode=debug,
// program=programDir) -> initialized event -> setBreakpoints -> configurationDone.
// Returns as soon as the initialize handshake completes; the launch itself
// (which builds the target — can take a few seconds) proceeds in the
// background and its outcome surfaces via Events() (a "stopped" event if
// stopOnEntry fired, an "output"/"error" event on build failure, etc.).
func Start(ctx context.Context, programDir string, breakpoints map[string][]int) (*Session, error) {
	panic("fake")
}

// dlv's own `go build` step runs relative to dlv's process working
// directory, not the launch config's `program`/`cwd` fields — a program
// living in its own module (as a scratch/example package, or any
// workspace outside chidori's own module) fails with "directory ...
// outside main module or its selected dependencies" unless dlv itself is
// started from inside that module. Confirmed by reproducing the failure
// against a real dlv subprocess and fixing it with exactly this.

// building+starting a real Go program can legitimately take a while

// Events returns the channel of simplified events (stopped/continued/output/
// terminated/exited/error) for the API layer to fan out to a WebSocket.
// Closed once the session terminates.
func (s *Session) Events() <-chan Event { panic("fake") }

// emit sends e on the events channel, guarded by eventsMu so it can never
// race with markDone's close(s.events) — a bare `select { case s.events <-
// e: default: }` is not safe here: a send on an already-closed channel
// panics regardless of the surrounding select/default, and the async launch
// goroutine (which can still be waiting on its own response, and therefore
// still emitting an "error" event on failure) can genuinely still be running
// after readLoop has exited and closed the channel — reproduced by a real
// "send on closed channel" panic under `go test ./...` before this guard.
func (s *Session) emit(e Event) { panic("fake") }

// A slow/absent consumer shouldn't block the debugger's own read
// loop — dropping an event here is better than deadlocking the
// whole session over a UI that isn't listening.

func (s *Session) newRequest(command string) dap.Request { panic("fake") }

// sendRequest writes req and blocks for its matching response (a 10s
// default timeout — every request in this package except launch is
// expected to be near-instant).
func (s *Session) sendRequest(req dap.Message) (dap.Message, error) { panic("fake") }

func (s *Session) sendRequestTimeout(req dap.Message, timeout time.Duration) (dap.Message, error) {
	panic("fake")
}

// readLoop is the single reader of the DAP connection: dispatches Responses
// to whichever sendRequest call is waiting on that Seq, and turns Events
// into this package's simplified Event shape (plus the "initialized" event's
// special handling — the trigger to actually send breakpoints and
// configurationDone, per the DAP handshake).
func (s *Session) readLoop() { panic("fake") }

// EOF/closed connection — session over, one way or another

// deliver hands a response to whichever sendRequest call registered seq.
func (s *Session) deliver(seq int, msg dap.Message) { panic("fake") }

// onInitialized runs the DAP configuration sequence once the adapter
// signals it's ready: set every requested breakpoint, then configurationDone
// (which is what actually lets the launched program start running).
func (s *Session) onInitialized() { panic("fake") }

func (s *Session) markDone() { panic("fake") }

// SetBreakpoints sets (replacing any previous set for that file) the
// breakpoints dlv should stop at in file, and returns what dlv actually
// verified (a requested line can be moved to the nearest executable one, or
// rejected entirely if the file isn't part of the compiled binary).
func (s *Session) SetBreakpoints(file string, lines []int) ([]Breakpoint, error) { panic("fake") }

func (s *Session) setBreakpointsRequest(file string, lines []int) ([]Breakpoint, error) {
	panic("fake")
}

// currentThread returns the thread ID execution is currently stopped on —
// every step/continue call operates on this thread. Errors if nothing is
// stopped (there's nothing to step from).
func (s *Session) currentThread() (int, error) { panic("fake") }

// Continue resumes execution of the currently stopped thread.
func (s *Session) Continue() error { panic("fake") }

// StepOver executes one line, stepping over any function calls on it.
func (s *Session) StepOver() error { panic("fake") }

// StepIn steps into a function call on the current line, if any.
func (s *Session) StepIn() error { panic("fake") }

// StepOut resumes execution until the current function returns.
func (s *Session) StepOut() error { panic("fake") }

// StackTrace returns the call stack for the currently stopped thread.
func (s *Session) StackTrace() ([]StackFrame, error) { panic("fake") }

// Scopes returns the variable scopes (locals, arguments, etc.) for a stack
// frame previously returned by StackTrace.
func (s *Session) Scopes(frameID int) ([]Scope, error) { panic("fake") }

// Variables returns the child variables for a scope or a structured
// variable's VariablesReference (0 means "no children" and shouldn't be
// called).
func (s *Session) Variables(variablesRef int) ([]Variable, error) { panic("fake") }

// Terminate asks dlv to disconnect (which also kills the debuggee, since it
// was started via launch) and force-kills the dlv subprocess itself as a
// backstop if it doesn't exit promptly.
func (s *Session) Terminate() { panic("fake") }

// readListenAddr reads dlv dap's stdout until it prints its listen address
// ("DAP server listening at: 127.0.0.1:PORT"), which is how dlv reports the
// actual ephemeral port chosen for --listen=127.0.0.1:0.
func readListenAddr(stdout io.Reader) (string, error) { panic("fake") }
