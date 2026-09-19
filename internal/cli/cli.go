// Package cli implements the chidori user-facing desktop CLI.
// Commands talk to a live coordinator over HTTP (same APIs as the IDE).
package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
)

// Version is stamped at build time via -ldflags "-X …/internal/cli.Version=…".
var Version = "dev"

// Options controls config resolution and live-probe behavior.
type Options struct {
	// ConfigPath is an explicit --config path. Empty → search candidates.
	ConfigPath string
	// BaseURL overrides the live coordinator URL. Empty → http://127.0.0.1:<port>
	// from the resolved config (default port 8080).
	BaseURL string
	// HTTPClient is used for live probes; nil → short-timeout default.
	HTTPClient *http.Client
}

// Resolved describes where config came from and whether the file exists.
type Resolved struct {
	Path   string
	Exists bool
	Source string // "flag", "cwd", "app-support", or "default"
}

// LiveInfo is a subset of GET /api/system/version when the coordinator is up.
type LiveInfo struct {
	OK          bool
	ListenPort  int
	GoVersion   string
	VCSRevision string
	Error       string
}

// DefaultConfigCandidates returns paths to try, in order.
func DefaultConfigCandidates() []string { panic("fake") }

// ResolveConfigPath picks the config file. Explicit path always wins (even if
// missing — Load will apply defaults). Otherwise the first existing candidate,
// else cwd/config.yaml as the default write target.
func ResolveConfigPath(explicit string) (Resolved, error) { panic("fake") }

func (o Options) client() *http.Client { panic("fake") }

func (o Options) baseURL(cfg *config.Config) string { panic("fake") }

// ProbeLive hits GET /api/system/version on the coordinator.
func ProbeLive(ctx context.Context, o Options, cfg *config.Config) LiveInfo { panic("fake") }

// FormatConfigShow renders Phase 0 `chidori config show` output.
func FormatConfigShow(res Resolved, cfg *config.Config, live LiveInfo) string { panic("fake") }

// ConfigShow loads config (file or defaults) and optionally probes live HTTP.
func ConfigShow(ctx context.Context, o Options) (string, error) { panic("fake") }

// Usage is the CLI help text.
func Usage() string { panic("fake") }
