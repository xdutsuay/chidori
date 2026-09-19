package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xdutsuay/lclreason/internal/appui"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/coordinator"
	"github.com/xdutsuay/lclreason/internal/crash"
	"github.com/xdutsuay/lclreason/internal/devlog"
	"github.com/xdutsuay/lclreason/internal/executil"
	"github.com/xdutsuay/lclreason/internal/netutil"
)

func main() { panic("fake") }

// KMA-234 / KMA-225 slice 2: unrecovered panics write a self-dump artifact
// plus crash-pending.json (ArtifactPath) for Diagnostics / next-launch dialog.
// Coordinator panics are recovered in runCoordinatorOnce and do not hit this path.
// No re-exec launcher here — supervisor spawn is a later ticket.

// KMA-236 / KMA-167 floor (640x520)

// Assets serves index.html and /assets/* directly (and is what
// lets Wails inject window.runtime/window.go into the page —
// see appui.WailsFS for why that matters). Handler is the
// fallback for anything Assets doesn't recognize: every
// /api/*, /v1/*, /ws/* request, proxied to the embedded
// coordinator.

// Hidden (not HiddenInset) so the app's custom HTML title bar
// (.app-title-bar in index.html — app name + live resource monitor)
// occupies that strip instead of a native title. HiddenInset enables a
// toolbar (UseToolbar:true) that vertically insets the traffic lights
// downward, which left them sitting ~6px below the "chidori" title text;
// TitleBarHidden has no toolbar, so the close/min/max buttons sit at the
// standard top position and line up with the title bar's centered text.
// The HTML bar's left padding still clears them and it's --wails-draggable.

type App struct {
	ctx        context.Context
	cancel     context.CancelFunc
	configPath string
	dataDir    string
	port       int // actual listen port (may bump if preferred is busy)
	preferred  int // PreferPortDefault (8080) unless overridden
	proxy      *httputil.ReverseProxy

	// Coordinator supervision state (Inference Source settings panel, "start
	// coordinator on local if it errors/crashes"). Guarded by coordMu since
	// it's read from a Wails-bound method (CoordinatorStatus, called from the
	// frontend) and written from the supervisor goroutine.
	coordMu       sync.Mutex
	coordUp       bool
	coordRestarts int
	coordLastErr  string
	coordSince    time.Time

	fullscreen bool // tracked here since Wails' runtime has no WindowIsFullscreen query
}

func NewApp() *App { panic("fake") }

func (a *App) httpHandler() http.Handler { panic("fake") }

func (a *App) startup(ctx context.Context) { panic("fake") }

// config.yaml's default secrets_file (".secrets.yaml") is a relative
// path, meant for the dev/CLI case where the process runs from the repo
// root. Inside the packaged .app the working directory is wherever
// macOS launched the bundle from — often read-only, or at least not
// guaranteed writable — so a relative path here fails with "read-only
// file system" the first time the user adds an API key (never hit in
// dev because `go run` from the repo root has a writable cwd). Anchor
// it to the same writable per-user Application Support dir a.dataDir
// already points at, same as cfg.DB above, unless the user explicitly
// configured an absolute path of their own.

// PL.12: check os.Args and CHIDORI_OPEN_URL for deep link URLs.

func (a *App) domReady(ctx context.Context) { panic("fake") }

// superviseCoordinator runs coordinator.Run in a loop: if it ever returns
// (a genuine error, e.g. the port being taken) or panics, and the app itself
// isn't shutting down, it relaunches the coordinator with a fresh context
// after a backoff, up to 30s between attempts. This is the "start
// coordinator on local if any error / if process crashed" behavior from the
// Inference Source settings panel — since the coordinator is an in-process
// goroutine (not a separate OS process, see AssetServer.Options.Handler
// above), "restart" means relaunching that goroutine, not respawning a
// process. An unrecovered panic elsewhere in the app would still take down
// the whole process; this only guards against the coordinator's own
// goroutine erroring out or panicking.
func (a *App) superviseCoordinator(runCtx context.Context, cfg *config.Config) { panic("fake") }

// Up flips true only once waitForCoordinatorHealth actually
// confirms GET /health responds — not eagerly here, since the
// coordinator hasn't started serving yet at this point in the loop.

// App is shutting down — this is not a crash.

// runCoordinatorOnce runs coordinator.Run and converts a panic into a
// returned error so the supervisor loop can log it and relaunch instead of
// the panic propagating up and killing the whole desktop app.
func (a *App) runCoordinatorOnce(runCtx context.Context, cfg *config.Config) (err error) {
	panic("fake")
}

// waitForCoordinatorHealth polls GET /health until the coordinator answers
// (or gives up after ~8s) and updates the tracked status accordingly. Called
// once synchronously at startup (blocking, same as before this feature) and
// again in the background after every automatic restart.
func (a *App) waitForCoordinatorHealth() { panic("fake") }

func (a *App) setCoordinatorStatus(up bool, errMsg string) { panic("fake") }

// CoordinatorStatus reports the in-process coordinator's health for the
// Inference Source settings panel. This is a Wails binding (IPC), not an
// HTTP call to the coordinator itself — deliberately, since the whole point
// is to report status *while the coordinator's own HTTP server is down*.
type CoordinatorStatus struct {
	Up        bool   `json:"up"`
	Restarts  int    `json:"restarts"`
	LastError string `json:"last_error"`
	SinceUnix int64  `json:"since_unix"`
}

func (a *App) CoordinatorStatus() CoordinatorStatus { panic("fake") }

func (a *App) shutdown(ctx context.Context) { panic("fake") }

func (a *App) ensurePaths() error { panic("fake") }

func appSupportDir() (string, error) { panic("fake") }

func copyFile(src, dst string) error { panic("fake") }

func (a *App) GetAppPaths() map[string]string { panic("fake") }

// PickFolder shows the native macOS "Open Folder" dialog and returns the
// chosen absolute path, or "" if the user cancelled. The frontend is
// responsible for calling POST /api/workspace/root with the result (which
// persists it to config.yaml and — if the user confirms trust — enables
// agent auto-apply for that folder).
func (a *App) PickFolder() string { panic("fake") }

// PickFile shows the native "Open File" dialog and returns the chosen
// absolute path, or "" if cancelled. The frontend soft-opens the parent
// directory as the workspace root (trust flow) then opens the file in the
// focused editor pane.
func (a *App) PickFile() string { panic("fake") }

// RevealInFinder opens the OS file manager with absPath selected. Used by
// the Explorer context menu's "Reveal in Finder" item; in a plain browser
// tab (LCLREASON_DEV) this binding doesn't exist, so the frontend falls
// back to copying the path to the clipboard instead.
func (a *App) RevealInFinder(absPath string) string { panic("fake") }

// explorer.exe often exits 1 even when the folder opened (KMA-131).

// ToggleFullscreen flips the window's fullscreen state and returns the new
// state. Wails' runtime only exposes WindowFullscreen/WindowUnfullscreen (no
// query), so the bool is tracked here — this binding is the single place
// that mutates it, called from both the native View ▸ Full Screen menu item
// and its Ctrl+Cmd+F equivalent in the frontend.
func (a *App) ToggleFullscreen() bool { panic("fake") }

// CloseWindow (File ▸ Close Window) hides the window rather than quitting —
// chidori has no multi-window support, so unlike VS Code (where Close Window
// closes one window among several) this is the single-window equivalent:
// the coordinator/agent state and any running terminals keep going in the
// background, and clicking the Dock icon (standard macOS app-activation
// behavior, not anything this binding has to implement) brings the window
// back exactly as it was. Distinct from Quit, which actually terminates the
// process — see the Exit/Quit row in the feature doc.
func (a *App) CloseWindow() { panic("fake") }

// OpenExternalURL opens url in the user's default system browser rather than
// the app's own webview (wailsruntime.BrowserOpenURL) — used for Help ▸ Show
// Release Notes. Only ever called with a hardcoded URL from our own bundled
// frontend code, never user-supplied input, so no allowlist is needed here.
func (a *App) OpenExternalURL(url string) { panic("fake") }

// NewWindow spawns a second chidori process using the same executable, with
// an optional workspace path. The new window is a fully independent OS
// process with its own coordinator, not a shared-state second window.
func (a *App) NewWindow(workspacePath string) string { panic("fake") }

// NewWindowWithProfile spawns a new window with both workspace and profile env vars.
func (a *App) NewWindowWithProfile(workspacePath, profile string) string { panic("fake") }

// PickWorkspaceFile shows a native file picker filtered for *.code-workspace files.
func (a *App) PickWorkspaceFile() string { panic("fake") }

// SaveWorkspaceFile shows a native save dialog for *.code-workspace.
func (a *App) SaveWorkspaceFile(suggestedName string) string { panic("fake") }

// HandleDeepLink processes a chidori:// or lclreason:// URL (PL.12). Wails v2
// has no options.Mac.URLHandlers — on macOS the OS delivers URLs via Apple
// Events or, on first launch, via os.Args. This method is called from startup
// for argv URLs and can be called from the frontend via Wails binding.
// The frontend receives a "deep-link" event with the raw URL string and
// decides what to do (open folder, open file, etc.).
func (a *App) HandleDeepLink(rawURL string) { panic("fake") }

// checkArgsForDeepLink scans os.Args for chidori:// or lclreason:// URLs
// and emits them as deep-link events.
func (a *App) checkArgsForDeepLink() { panic("fake") }

// OpenDevTools emits a menu event so the frontend can show instructions for
// accessing the WebView inspector (Wails v2 has no runtime API for this; the
// env var CHIDORI_OPEN_INSPECTOR=1 opens it on startup via options.Debug).
func (a *App) OpenDevTools() { panic("fake") }

// SetSystemTheme mirrors the in-app theme (ADR-0009) to the native window
// chrome so the title bar matches. dark==false means "light" — there's no
// separate "system" case here because setTheme() in
// public/js/06-palette-keys.js already resolves
// system/light/dark down to a boolean before calling this.
func (a *App) SetSystemTheme(dark bool) { panic("fake") }
