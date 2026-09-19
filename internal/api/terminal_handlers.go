package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/coder/websocket"
)

// TerminalStart creates a new PTY session. An optional {"shell": "..."} body
// (T.10 "Terminal profile") overrides $SHELL for this one session — a bare
// POST with no body is unchanged, existing behavior, so Decode's error on an
// empty body is deliberately ignored rather than treated as a bad request.
func (h *Handlers) TerminalStart(w http.ResponseWriter, r *http.Request) { panic("fake") }

// TerminalKill terminates a terminal session (T.6 "Kill terminal"). Killing
// an already-exited or unknown id is treated as success — the client's goal
// (no more session with this id) is already satisfied.
func (h *Handlers) TerminalKill(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) TerminalResize(w http.ResponseWriter, r *http.Request) { panic("fake") }

// wailsOriginPatterns authorizes the packaged Wails webview when it dials the
// coordinator's real loopback http.Server for the terminal WebSocket
// (terminalWsUrl in public/js/05-inference-keys.js). That path is
// cross-origin by design: the page
// Origin is wails://… / wails.localhost while Host is 127.0.0.1:<port>, so
// coder/websocket's same-origin default rejects the handshake and the UI
// shows "[disconnected]" immediately. Request Host is always authorized by
// the library (dev browser / same-origin stays fine). Evil third-party sites
// are still rejected.
var wailsOriginPatterns = []string{
	"wails", // Origin: wails://wails
	"wails.localhost",
	"http://wails.localhost",
	"https://wails.localhost",
	"wails://*",
	"*.wails.localhost",
}

func (h *Handlers) TerminalWS(w http.ResponseWriter, r *http.Request) { panic("fake") }

// OriginPatterns (not InsecureSkipVerify): allow the Wails→loopback hop
// while still rejecting forged Origin: https://evil.example.com (CSWSH).
