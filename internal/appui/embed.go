package appui

import (
	"embed"
	"io/fs"
	"strings"
)

//go:embed index.html
var IndexHTML []byte

//go:embed public/*
var Public embed.FS

//go:embed index.html public
var Root embed.FS

// WailsFS adapts the embedded assets to the flat "index.html" + "assets/..."
// layout the frontend actually requests (see internal/api/ui.go, which
// serves the same tree at GET / and GET /assets/* for the coordinator's own
// HTTP server) into a plain fs.FS suitable for assetserver.Options.Assets in
// the Wails desktop shell (cmd/lclreason-app).
//
// This exists because the desktop shell previously served every request —
// including index.html itself — through a raw reverse-proxy http.Handler.
// Wails only injects its runtime/IPC bridge (window.runtime, window.go) into
// HTML it serves from its own Assets filesystem; a bare Handler bypasses
// that injection entirely, silently leaving window.runtime/window.go
// undefined in the packaged app. Passing WailsFS as Assets lets Wails serve
// (and inject into) index.html and static files itself, while the reverse
// proxy Handler remains the fallback for everything Assets doesn't have —
// i.e. all /api/*, /v1/*, /ws/* requests, which is exactly the set that
// needs to reach the embedded coordinator.
type WailsFS struct{}

// Open implements fs.FS.
func (WailsFS) Open(name string) (fs.File, error) { panic("fake") }
