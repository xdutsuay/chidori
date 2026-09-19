package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xdutsuay/lclreason/internal/config"
)

// settingsJSON is the VS Code-ish settings.json shape exported/imported
// by CFG.14. Field names mirror VS Code's own dotted keys.
type settingsJSON struct {
	EditorFontSize              int      `json:"editor.fontSize,omitempty"`
	EditorWordWrap              string   `json:"editor.wordWrap,omitempty"`
	EditorMinimap               *bool    `json:"editor.minimap.enabled,omitempty"`
	EditorLigatures             *bool    `json:"editor.fontLigatures,omitempty"`
	EditorRenderWhitespace      string   `json:"editor.renderWhitespace,omitempty"`
	EditorStickyScroll          *bool    `json:"editor.stickyScroll.enabled,omitempty"`
	EditorMultiCursorModifier   string   `json:"editor.multiCursorModifier,omitempty"`
	EditorFormatOnSave          *bool    `json:"editor.formatOnSave,omitempty"`
	EditorAutoSave              string   `json:"files.autoSave,omitempty"`
	EditorTheme                 string   `json:"workbench.colorTheme,omitempty"`
	EditorTabSize               int      `json:"editor.tabSize,omitempty"`
	EditorKeybindingMode        string   `json:"editor.keybindingMode,omitempty"`
	TerminalScrollback          int      `json:"terminal.integrated.scrollback,omitempty"`
	TerminalShell               string   `json:"terminal.integrated.defaultProfile,omitempty"`
	CompletionDisabledLangs     []string `json:"editor.completionDisabledLanguages,omitempty"`
	UIZoomPercent               int      `json:"window.zoomLevel,omitempty"`
	DraftMode                   *bool    `json:"chidori.draftMode,omitempty"`
	ExperimentalBrowserTools    *bool    `json:"chidori.experimental.browserTools,omitempty"`
	ExperimentalBrowserPanel    *bool    `json:"chidori.experimental.browserPanel,omitempty"`
	ExperimentalWebFetch        *bool    `json:"chidori.experimental.webFetch,omitempty"`
	ExperimentalHarnessVis      *bool    `json:"chidori.experimental.harnessVisualizer,omitempty"`
	ExperimentalExternalHarness *bool    `json:"chidori.experimental.externalHarness,omitempty"`
}

func boolPtr(v bool) *bool { panic("fake") }

func uiConfigToSettingsJSON(u config.UIConfig) settingsJSON { panic("fake") }

func settingsJSONToUIConfig(s settingsJSON, base config.UIConfig) config.UIConfig { panic("fake") }

// GetSettingsJSON exports UIConfig as a VS Code-style settings.json (CFG.14).
func (h *Handlers) GetSettingsJSON(w http.ResponseWriter, r *http.Request) { panic("fake") }

// PutSettingsJSON imports a VS Code-style settings.json into UIConfig (CFG.14).
func (h *Handlers) PutSettingsJSON(w http.ResponseWriter, r *http.Request) { panic("fake") }
