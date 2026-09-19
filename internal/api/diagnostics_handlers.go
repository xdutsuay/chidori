package api

import (
	"net/http"
	"strconv"
)

// DiagnosticsRunsList is the desktop Settings → Diagnostics feed of recent
// Ask/Agent/Plan/Debug runs (ChidoriRunTracker). Same payload shape as the
// companion GET /runs, but without bearer auth — localhost Settings, same
// class as GET /api/activity.
func (h *Handlers) DiagnosticsRunsList(w http.ResponseWriter, r *http.Request) { panic("fake") }

// DiagnosticsRunsGet returns detail (current_step + log_tail) for one run id.
func (h *Handlers) DiagnosticsRunsGet(w http.ResponseWriter, r *http.Request) { panic("fake") }
