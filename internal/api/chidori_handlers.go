package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/xdutsuay/lclreason/internal/activity"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/registry"
)

// ---- protocol / headers ----

func (h *Handlers) chidoriProtocolVersion() string { panic("fake") }

func (h *Handlers) chidoriEchoVersionHeader(w http.ResponseWriter, r *http.Request) {
	panic(
		// Per WIRE_CONTRACT.md: client sends X-Chidori-Protocol-Version, server echoes
		// the version it will actually use in the same header.
		"fake")
}

// A real, common early-integration failure mode: the phone and
// desktop built against different protocol revisions. Worth
// surfacing even though it isn't fatal by itself (the phone is
// expected to fall back to its own "update required" UI).

func (h *Handlers) chidoriRequireBearer(w http.ResponseWriter, r *http.Request) bool { panic("fake") }

// ---- GET /version ----

func (h *Handlers) ChidoriVersion(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- POST /pairing/begin ----

func (h *Handlers) ChidoriPairingBegin(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Body is {} (ignored).

// ---- POST /pairing/confirm ----

func (h *Handlers) ChidoriPairingConfirm(w http.ResponseWriter, r *http.Request) { panic("fake") }

// auth_token is mandatory per contract.

// ---- DELETE /pairing/{instance_id} ----

func (h *Handlers) ChidoriPairingRevoke(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- GET /coordinator/status ----

func (h *Handlers) ChidoriCoordinatorStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- GET /runs?limit= ----

func (h *Handlers) ChidoriRunsList(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- GET /runs/{run_id} ----

func (h *Handlers) ChidoriRunsGet(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- POST /runs/{run_id}/cancel ----

func (h *Handlers) ChidoriRunsCancel(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- POST /runs/{run_id}/messages ----

func (h *Handlers) ChidoriRunsInjectMessage(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- WS /chat/stream ----

func (h *Handlers) ChidoriChatStream(w http.ResponseWriter, r *http.Request) { panic("fake") }

// The phone runs in a different origin.

// Echo the user message first (phone does not locally echo).

// Stream the model's reply as one or more from_user=false frames.

// Use the same pool streaming path as the desktop chat surface.

// Best-effort: if streaming fails, send a single error-like assistant frame.

// The single most common "chat doesn't work" cause in practice —
// no local/remote inference source currently attached — surfaces
// here as a real error from the pool, not a silent no-op.

func (h *Handlers) chidoriDefaultModelAndProvider() (model string, provider string) {
	panic(
		// Keep this minimal: match the desktop's current default routing as best we
		// can without introducing new config. Prefer explicitly configured model,
		// otherwise pick the first healthy node model.
		"fake")
}

// Fall back to empty model (pool/dispatch will use backend defaults where possible).

// ---- UI Endpoints for Companion ----

func (h *Handlers) UICompanionStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Phone-driven POST /pairing/begin generates a code the contract says is
// "shown on desktop". Settings only used to display codes from the
// Generate Pairing Code button — if the phone initiated begin, the
// panel stayed blank and the user had nothing to type. Surface the
// active code (while unexpired) so the panel poll can show it.

func (h *Handlers) UICompanionPair(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ---- Node mode (NODE_MODE_SPIKE.md / protocol §2.5) ----

type chidoriNodeRegisterRequest struct {
	NodeID             string   `json:"node_id"`
	DisplayName        string   `json:"display_name"`
	APIBase            string   `json:"api_base"`
	Models             []string `json:"models"`
	ContextLength      *int     `json:"context_length"`
	ApproxTokensPerSec *float64 `json:"approx_tokens_per_sec"`
	BatteryPct         *int     `json:"battery_pct"`
	Charging           *bool    `json:"charging"`
	Available          *bool    `json:"available"`
	DataPlaneToken     string   `json:"data_plane_token"`
}

// ChidoriNodeRegister upserts a phone worker into the inference registry.
func (h *Handlers) ChidoriNodeRegister(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ChidoriNodeHeartbeat(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Get returns a snapshot; re-upsert

func (h *Handlers) ChidoriNodeUnregister(w http.ResponseWriter, r *http.Request) { panic("fake") }

func parseLooseURL(raw string) (*url.URL, error) { panic("fake") }

// UICompanionLog is a desktop-local (not phone-facing, no bearer auth)
// endpoint for the Settings -> Companion App panel's log viewer — see
// ChidoriCompanion.Log's doc comment for why this exists: debugging a real
// phone that "isn't working" previously meant reading server stdout, which
// isn't visible from the desktop UI at all.
func (h *Handlers) UICompanionLog(w http.ResponseWriter, r *http.Request) { panic("fake") }
