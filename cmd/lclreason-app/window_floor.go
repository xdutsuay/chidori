package main

// Window floor from KMA-167 chrome density tokens (--chrome-window-min-*).
// Enforced via Wails MinWidth/MinHeight so the editor pane cannot collapse to
// zero when the user shrinks the window (KMA-236 item 2).
const (
	windowMinWidth  = 640
	windowMinHeight = 520
)
