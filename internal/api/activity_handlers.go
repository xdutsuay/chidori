package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/xdutsuay/lclreason/internal/activity"
)

// statusEmitter records activity to the bus (for /debug) and, when streaming,
// forwards a friendly subset to the chat UI as `event: status` SSE frames.
type statusEmitter struct {
	sw        *SSEWriter // nil for non-streaming requests
	bus       *activity.Bus
	requestID string
	clientID  string

	runs RunTracker
}

func (e *statusEmitter) Emit(a activity.Activity) { panic("fake") }

// A friendly, low-cardinality "current step" plus a small log tail for
// the chidori-nagasa monitor (WIRE_CONTRACT.md).

// isStatusKind selects the events worth showing in the chat status pill.
func isStatusKind(kind string) bool { panic("fake") }

// emitRunTerminal records the activity-bus end of an Ask (and similar) run so
// Diagnostics does not keep a live "still waiting" clock after Stop or failure.
func emitRunTerminal(em activity.Emitter, err error) { panic("fake") }

// Generic abort — Stop, companion, timeout, or disconnect.
// Chat UI labels intentional Stop as "Stopped by you"; do not claim that here.

// newEmitter builds a per-request emitter. sw may be nil (non-streaming).
func (h *Handlers) newEmitter(requestID, clientID string, sw *SSEWriter) *statusEmitter {
	panic("fake")
}

// ListActivity returns recent activity events for the developer debug view,
// optionally filtered by ?client_id= and ?request_id= (alias ?task=, ADR-0017
// Decision 3), limited by ?limit=.
func (h *Handlers) ListActivity(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ServeDebugUI serves the developer activity timeline page.
func (h *Handlers) ServeDebugUI(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DebugStatus reports whether activity recording is on.
func (h *Handlers) DebugStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SetDebug toggles activity recording at runtime. The chat status pill is
// unaffected (it's product UX, not developer telemetry).
func (h *Handlers) SetDebug(w http.ResponseWriter, r *http.Request) { panic("fake") }

// IngestActivity records a client-reported UI action (e.g. "deleted
// conversation") onto the timeline. Server-stamps the time; caps text sizes.
func (h *Handlers) IngestActivity(w http.ResponseWriter, r *http.Request) { panic("fake") }

func clip(s string, n int) string { panic("fake") }
