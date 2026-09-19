package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/xdutsuay/lclreason/internal/tools"
)

// BrowserState returns the shared headless browser session's current URL
// and title (AG.15b — agent chromedp ↔ in-app panel share one session).
func (h *Handlers) BrowserState(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SetBrowserState updates the shared browser URL from the in-app panel,
// so the agent's session sees the same page the user navigated to.
func (h *Handlers) SetBrowserState(w http.ResponseWriter, r *http.Request) { panic("fake") }
