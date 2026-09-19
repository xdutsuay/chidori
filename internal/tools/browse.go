package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

const maxBrowseOutput = 8000

const (
	browseToolNavigate = "browse_navigate"
	browseToolSnapshot = "browse_snapshot"
	browseToolClick    = "browse_click"
)

var browseSession struct {
	sync.Mutex
	ready       bool
	allocCtx    context.Context
	allocCancel context.CancelFunc
	tabCtx      context.Context
	tabCancel   context.CancelFunc
	currentURL  string
	lastTitle   string
}

// BrowseState returns the shared browser session's current URL and title.
// Both the agent chromedp tools and the in-app browser panel see this state.
func BrowseState() (url, title string) { panic("fake") }

// SetBrowseURL updates the shared URL from the in-app panel (so the agent's
// session and the user's panel stay in sync — AG.15b shared session).
func SetBrowseURL(url string) { panic("fake") }

// RegisterBrowserTools adds experimental headless-browser agent tools.
func RegisterBrowserTools(r *Registry) { panic("fake") }

// ReconcileBrowserTools registers or removes browse_* tools to match enabled.
// When disabled, any running headless session is closed.
func ReconcileBrowserTools(r *Registry, enabled bool) { panic("fake") }

// CloseBrowseSession shuts down the lazy headless browser, if running.
func CloseBrowseSession() { panic("fake") }

func closeBrowseSessionLocked() { panic("fake") }

func browseSessionCtx(parent context.Context) (context.Context, error) { panic("fake") }

func withBrowse(parent context.Context, fn func(ctx context.Context) error) error { panic("fake") }

func browseNavigateTool(ctx context.Context, params map[string]any) (*ToolResult, error) {
	panic("fake")
}

func browseSnapshotTool(ctx context.Context, params map[string]any) (*ToolResult, error) {
	panic("fake")
}

func browseClickTool(ctx context.Context, params map[string]any) (*ToolResult, error) { panic("fake") }

func snapshotJS(selector string) string { panic("fake") }

func truncateBrowseOutput(s string) string { panic("fake") }
