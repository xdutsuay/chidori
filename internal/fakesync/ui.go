package fakesync

import (
	_ "embed"
	"net/http"
)

//go:embed ui.html
var uiHTML []byte

// UIHandler serves the temporary selection UI (GET /).
func UIHandler() http.Handler { panic("fake") }

// MountUI registers UI + plan/sync API on mux.
func MountUI(mux *http.ServeMux, cfg HandlerConfig) { panic("fake") }
