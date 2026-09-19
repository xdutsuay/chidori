package main

import (
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// buildMenu wires a native macOS menu bar to the existing frontend commands
// in public/js/*. Each item just emits a "menu:<action>" event; the frontend
// listens for these via window.runtime.EventsOn and calls the same function
// the equivalent keyboard shortcut or button already calls — the menu bar
// isn't a second implementation of these actions, just another way to reach
// the one that already exists.
//
// Menu accelerators (keys.CmdOrCtrl etc.) are handled by the OS menu system
// before the webview's own keydown listeners see them, so this doesn't
// double-fire against the in-page shortcuts already wired in the frontend — those
// remain the only path when running in a plain browser tab (LCLREASON_DEV),
// where there is no native menu at all.
func buildMenu(app *App) *menu.Menu { panic("fake") }

// Standard "lclreason" app menu (About/Services/Hide/Quit) — must be
// appended immediately after NewMenu() on macOS.

// Standard Edit menu (Undo/Redo/Cut/Copy/Paste/Select All) — gives
// native Cmd+C/V/Z inside text inputs (chat box, commit message,
// search fields), not just inside Monaco.

// VS Code ships neither with a default keybinding either.

// VS Code's own default is the Ctrl+K Ctrl+Q chord; chidori has no
// chord-sequence keyboard infrastructure (same precedent as Zen Mode /
// Centered Layout), so this is menu/palette only, no accelerator.

// No accelerator: Monaco's own built-in quickOutline action already owns
// Cmd/Ctrl+Shift+O internally (its precondition is just "a
// DocumentSymbolProvider is registered" — see registerDocumentSymbolProvider
// in public/js/06-palette-keys.js). Binding it here too would make Wails' OS-level menu
// accelerator intercept the keypress before the webview ever sees it,
// silently breaking Monaco's own handling in the packaged app for no
// reason — the menu item still gives mouse discoverability without
// claiming the combo.
