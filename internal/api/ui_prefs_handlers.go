package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/tools"
)

// GetUIPrefs returns the persisted editor/UI preferences (ADR-0005). Never
// unavailable — an unconfigured coordinator just returns the zero UIConfig,
// which the frontend already treats as "use built-in defaults".
func (h *Handlers) GetUIPrefs(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SetUIPrefs persists editor/UI preferences to config.yaml (ADR-0005) — same
// lock-mutate-Save-unlock shape as WorkspaceSetRoot/WorkspaceTrust in
// workspace_handlers.go, so this stays consistent with every other runtime
// config write in the app.
func (h *Handlers) SetUIPrefs(w http.ResponseWriter, r *http.Request) { panic("fake") }

// syncBrowserTools registers or removes experimental browse_* tools on the
// shared agent registry to match ui.experimental_browser_tools.
func (h *Handlers) syncBrowserTools() { panic("fake") }

// syncDeepThinkTools registers or removes deep_think on the shared agent
// registry to match ui.assisted_thinking (KMA-109).
func (h *Handlers) syncDeepThinkTools() { panic("fake") }
