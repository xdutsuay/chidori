package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/xdutsuay/lclreason/internal/debugger"
)

// intQueryParam parses a required integer query parameter, with an error
// message that names the parameter — used by the scopes/variables debug
// endpoints, which both take a single int ref.
func intQueryParam(r *http.Request, name string) (int, error) { panic("fake") }

// Run & Debug panel v1 (BE.15): a thin HTTP/WS layer over internal/debugger's
// DAP client. Not trust-gated — same precedent as the integrated terminal
// (internal/api/terminal_handlers.go): a human clicking "Start Debugging" on
// their own open workspace is a direct user action, not an agent-triggered
// one, and the terminal already lets that same human run arbitrary shell
// commands with no trust gate, so gating "build and run this Go package"
// behind trust would be a narrower, inconsistent restriction on a strictly
// less powerful capability.

// resolveWorkspacePath resolves a workspace-relative path (program dir or
// breakpoint file) to an absolute path via the sandboxed wsFS.Resolve, then
// evaluates symlinks — dlv/DAP report source paths post-symlink-resolution
// (confirmed against a real dlv subprocess: macOS's /tmp being a symlink to
// /private/tmp was enough to break naive path comparison), so every path
// handed to or compared against the debugger needs the same treatment.
func (h *Handlers) resolveWorkspacePath(rel string) (string, error) { panic("fake") }

// relWorkspacePath is resolveWorkspacePath's inverse, best-effort — used to
// report stack frame file paths back to the frontend as workspace-relative
// (what every other panel in this app already expects), falling back to the
// absolute path if it isn't actually under the (symlink-resolved) root for
// any reason rather than erroring the whole response over a display detail.
func (h *Handlers) relWorkspacePath(abs string) string { panic("fake") }

// DebugStart starts a new debug session (POST /api/debug/start), building
// and launching the Go package at body.Program (a workspace-relative
// directory) under `dlv debug`, with an initial set of breakpoints. Rejects
// if a session is already active — v1 is deliberately single-session, same
// "one thing at a time" precedent as the agent loop.
func (h *Handlers) DebugStart(w http.ResponseWriter, r *http.Request) { panic("fake") }

// activeDebugSession returns the current session, or writes a 404 and
// returns nil if none is active — every control/inspection endpoint below
// starts with this same check.
func (h *Handlers) activeDebugSession(w http.ResponseWriter) *debugger.Session { panic("fake") }

// DebugSetBreakpoints updates breakpoints for one file mid-session (POST
// /api/debug/breakpoints), returning what dlv actually verified.
func (h *Handlers) DebugSetBreakpoints(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugContinue resumes execution (POST /api/debug/continue).
func (h *Handlers) DebugContinue(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugStepOver steps over the current line (POST /api/debug/step-over).
func (h *Handlers) DebugStepOver(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugStepIn steps into the current line's function call (POST /api/debug/step-in).
func (h *Handlers) DebugStepIn(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugStepOut runs until the current function returns (POST /api/debug/step-out).
func (h *Handlers) DebugStepOut(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugStop terminates the active session (POST /api/debug/stop). Clearing
// the session is unconditional even if Terminate itself hit an error — a
// wedged dlv subprocess shouldn't leave the UI permanently stuck believing a
// session is still active with no way to start a new one.
func (h *Handlers) DebugStop(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugSessionStatus reports whether a debug session is active (GET
// /api/debug/status) — named to avoid colliding with the pre-existing
// DebugStatus (activity_handlers.go), which reports the unrelated
// activity-bus developer-telemetry toggle.
func (h *Handlers) DebugSessionStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugStackTrace returns the call stack for the currently stopped thread
// (GET /api/debug/stack), with file paths translated back to
// workspace-relative form.
func (h *Handlers) DebugStackTrace(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugScopes returns the variable scopes for a stack frame (GET
// /api/debug/scopes?frame=<id>).
func (h *Handlers) DebugScopes(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugVariables returns the child variables of a scope or structured
// variable (GET /api/debug/variables?ref=<variablesReference>).
func (h *Handlers) DebugVariables(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugEvents streams the active session's events (stopped/continued/output/
// terminated/exited/error) over a WebSocket (GET /api/debug/events). Same
// no-InsecureSkipVerify-origin-check posture as the terminal WS
// (terminal_handlers.go) — this is a local desktop endpoint exposing live
// debuggee state, not something a same-browser different-origin page should
// be able to attach to.
//
// v1 scoping note: exactly one WS connection can meaningfully drain a
// session's event channel (ranging over it consumes items) — matches this
// package's single-session model, but a page reload while a session is
// running will miss events that arrive before the new connection attaches.
// Acceptable for a v1 aimed at one desktop window watching one session; a
// future pass could add a small ring buffer + fan-out if that turns out to
// matter in practice.
func (h *Handlers) DebugEvents(w http.ResponseWriter, r *http.Request) { panic("fake") }
