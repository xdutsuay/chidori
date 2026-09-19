package coordinator

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/api"
	"github.com/xdutsuay/lclreason/internal/backendprobe"
	"github.com/xdutsuay/lclreason/internal/cache"
	"github.com/xdutsuay/lclreason/internal/chain"
	"github.com/xdutsuay/lclreason/internal/compact"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/memory"
	"github.com/xdutsuay/lclreason/internal/registry"
	"github.com/xdutsuay/lclreason/internal/secrets"
	"github.com/xdutsuay/lclreason/internal/store"
	"github.com/xdutsuay/lclreason/internal/tools"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// Run starts the coordinator HTTP server and blocks until ctx is cancelled.
func Run(ctx context.Context, cfg *config.Config) error { panic("fake") }

// Always register workspace tools against wsFS, even if no root is
// configured yet: wsFS is a long-lived pointer whose root can be set
// later at runtime (Open Folder from the IDE), and the tools resolve
// against it dynamically on every call.

// KMA-120 — thin gh CLI wrappers

// ADR-0006: a workspace already configured+trusted from a previous run
// should get gopls without the user having to re-trigger Open Folder or
// re-toggle trust — WorkspaceSetRoot/WorkspaceTrust cover every
// subsequent change from here.

// ADR-0012: same reasoning — MCP servers configured in a previous run
// should reconnect on startup without needing a Settings round-trip.

// One cadence for both sweep and probe, and it must be shorter than
// offline_after: HealthCheck refreshes LastSeen for reachable passive
// nodes (attached by IP / LAN discovery — they run no heartbeat
// client), so probing at the old offline_after cadence left them
// sitting exactly on Healthy()'s cutoff and randomly excluded from
// dispatch.
//
// Phase 0 energy: when the UI reports idle (POST /api/ui/activity
// stale >45s), back off to 30s instead of cfg.Heartbeat (usually 5s)
// so a hidden laptop app stops hammering Ollama/nodes every few seconds.

// Reflects whether the local node is currently registered+reachable.
// selfRegisterLocal only registers when the backend is live, so a
// missing node means "was down at startup".

// Backend is live. Register it if it wasn't at startup (or
// if it restarted) so a node the user brings up later
// appears without needing an app restart; otherwise just
// refresh its heartbeat.

// Poll every registered node's own address directly. This is
// what keeps "passive" nodes (attached by IP or found via LAN
// discovery — internal/api/nodes_handlers.go) marked
// healthy/unhealthy without requiring them to run any
// lclreason-specific heartbeat client; a bare LM Studio/
// Ollama/vLLM server is enough. A reachable probe also counts
// as a heartbeat (refreshes LastSeen) — see Registry.HealthCheck.

// Probe immediately at startup: self-registration marks the local
// node Healthy without checking, so a dead backend used to show
// "Healthy" in the UI for a full offline_after window after launch.

// bgCtx (not context.Background()): ties the prewarm call's own
// timeout to the coordinator's own shutdown signal (running_issue.md
// Code Critic #4 / H.13). A detached context.Background() let the
// warmup HTTP call keep running for up to its own 60s budget even
// after the coordinator had been asked to shut down — a goroutine
// (and an in-flight request to whatever backend it warms) outliving
// the process' own intended lifetime. bgCtx is the same shutdown
// signal already used by the health-check loop just above.

func selfNodeID() string { panic("fake") }

// LocalBackendReachable reports whether the configured local LLM backend responds.
func LocalBackendReachable(cfg *config.Config) bool { panic("fake") }

func prewarmCoordinator(parent context.Context, cfg *config.Config, pool *dispatch.Pool, reg *registry.Registry) {
	panic("fake")
}

// selfRegisterLocal registers the coordinator's own machine as a local
// inference node — but ONLY if that backend is actually reachable right now.
// The coordinator is a coordinator first; being an inference source is a
// separate, opt-in fact that must be proven by a live probe, not assumed from
// config. Without this gate a defaulted `backend: ollama` (the config default)
// permanently showed a dead "ollama @ 127.0.0.1:11434 — Unreachable" node in
// the UI on any machine that never ran Ollama. Returns true if it registered.
// The probe loop calls this again when the backend later comes up, so a node
// you start after launch still appears without a restart.
func selfRegisterLocal(cfg *config.Config, reg *registry.Registry) bool { panic("fake") }

// Reachability is the ONLY gate. A dead local backend is simply not an
// inference source, regardless of whether the backend was named in config.

// Detect the node's actual loaded model(s) so a request with model="auto"
// resolves to a real id. Without this, an OpenAI-compatible backend (LM
// Studio / AirLLM) registered here got Models=[] → effectiveModel() sent an
// empty "model" → the backend 400s, even though the /v1/models health probe
// (which is served regardless of what's loaded) reports the node Healthy.
// This is exactly the "node shows healthy but every chat fails with
// 'lmstudio failed'" incoherence. The attach-by-IP path already did this
// (internal/api/nodes_handlers.go); the self-registered local node didn't.

type toolAdapter struct {
	reg *tools.Registry
}

func (a *toolAdapter) Execute(ctx context.Context, name string, params map[string]any) (chain.ToolExecResult, error) {
	panic("fake")
}

func (a *toolAdapter) Has(name string) bool { panic("fake") }

type poolInvoker struct {
	pool     *dispatch.Pool
	provider string
}

func (p *poolInvoker) Invoke(ctx context.Context, model, prompt string) (string, error) {
	panic("fake")
}
