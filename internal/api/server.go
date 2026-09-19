package api

import (
	"context"
	"crypto/subtle"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/grandcat/zeroconf"

	"github.com/xdutsuay/lclreason/internal/activity"
	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/cache"
	"github.com/xdutsuay/lclreason/internal/chain"
	"github.com/xdutsuay/lclreason/internal/compact"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/crash"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/executil"
	"github.com/xdutsuay/lclreason/internal/flightlog"
	"github.com/xdutsuay/lclreason/internal/harness"
	"github.com/xdutsuay/lclreason/internal/harness/external"
	"github.com/xdutsuay/lclreason/internal/harnesstrace"
	"github.com/xdutsuay/lclreason/internal/lsp"
	"github.com/xdutsuay/lclreason/internal/mcpclient"
	"github.com/xdutsuay/lclreason/internal/memory"
	"github.com/xdutsuay/lclreason/internal/netutil"
	"github.com/xdutsuay/lclreason/internal/profiler"
	"github.com/xdutsuay/lclreason/internal/registry"
	"github.com/xdutsuay/lclreason/internal/secrets"
	"github.com/xdutsuay/lclreason/internal/session"
	"github.com/xdutsuay/lclreason/internal/store"
	"github.com/xdutsuay/lclreason/internal/terminal"
	"github.com/xdutsuay/lclreason/internal/tools"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// ChidoriProtocolVersion is the wire-contract version advertised over mDNS
// and returned from GET /version (chidori-nagasa WIRE_CONTRACT / protocol 1.2.0).
const ChidoriProtocolVersion = "1.2.0"

type Server struct {
	http     *http.Server
	handlers *Handlers

	// companionHTTP is the dedicated LAN companion listener (default :8027).
	// Phone discovery/pairing/status/chat go here; the IDE stays on http.
	companionHTTP *http.Server
	companionPort int

	mdns *zeroconf.Server

	// companionMu guards companion start/retry against Shutdown.
	companionMu     sync.Mutex
	companionCancel context.CancelFunc
	companionReady  bool // true once Serve is running on an owned listener
}

func New(port int,
	reg *registry.Registry,
	dis *dispatch.Dispatcher,
	pool *dispatch.Pool,
	eng *chain.Engine,
	cch *cache.Cache,
	st *store.Store,
	mem *memory.VectorStore,
	toolReg *tools.Registry,
	compactor *compact.Compactor,
	authKey string,
	profiling bool,
	debugEnabled bool,
	secretsPath string,
	wsFS *workspace.FS,
	codeCfg config.CodeConfig,
	cfg *config.Config,
	runs RunTracker,
) *Server {
	panic("fake")
}

// KMA-234: pending crash markers live next to the DB under {dataDir}/crashes.

// Stable instance id for LAN discovery + pairing (persisted in config.yaml).
// Best-effort: if config can't be saved (read-only path), keep it in memory
// for this run, but still satisfy the "stable across restarts" requirement
// whenever config persists.

// Core API (IDE / coordinator). Phone-facing companion routes live on the
// dedicated companion listener below — not here — so mDNS port 8027 and
// the IDE port stay distinct (protocol 1.2.0).

// Extended API

// Prometheus metrics

// In-app resource monitor widget (CPU/RAM/disk/network of this process)

// Per-user context / compaction inspector

// Developer activity / debug timeline

// Stress / load test

// Memo (shared clipboard)

// Memory (RAG)

// P3 netdoc — experimental @web / browse_fetch (SSRF-hardened).

// AG.15b — shared browser state (agent chromedp ↔ in-app panel sync).

// Workspace (IDE)

// Workflows (ADR-0015 Phase 1) — register /runs before /{id}

// Persisted editor/UI preferences (ADR-0005) — config.yaml `ui:` section.

// CFG.14 settings.json import/export (VS Code-ish shape).

// Cross-project rules file (ADR-0011) — outside workspace.FS's sandbox,
// so it gets its own small pair of handlers rather than /api/workspace/file.

// gopls LSP bridge (ADR-0006) — Go-only v1: diagnostics, hover, definition,
// document symbols (Go to Symbol in Editor).

// MCP client (ADR-0012) — chidori's agent consuming external MCP servers.

// Source control (git status/diff/stage/commit)

// Agent edits

// AG.16 background jobs (list / start / cancel)

// Harness visualizer (experimental, off by default — HV.1/HV.3)

// Skills (Save as Skill, borrowed from Hermes Agent — §14.2 #3)

// IDE sessions

// Terminal + cluster status

// Run & Debug panel (BE.15) — scoped v1 via Delve DAP (internal/debugger)

// WebSocket compute nodes (browser / mobile)

// App UI (embedded frontend)

// Legacy UIs (dev only)

// Middleware stack (IDE server)

// Dedicated companion listener (protocol 1.2.0 default :8027). Phone
// discovery/pairing/status/runs/chat only — no IDE routes, no coordinator
// API-key gate (companion uses its own bearer after pairing).

// No WriteTimeout: chat streams can run for many minutes.

func (s *Server) Start() error { panic("fake") }

// Always publish manual-pairing address candidates up front — even when
// the companion port is currently held by another process — so Settings
// isn't blank while the user frees the port / waits for retry.

// A transient squatter (stale chidori, test harness, etc.) used to
// brick pairing for the whole IDE lifetime. Retry in the background
// so freeing the port is enough — no forced relaunch.

// tryStartCompanion binds the companion port with tcp4, then starts mDNS and
// Serve. Returns false when the port is unavailable (no mDNS advertised).
//
// Bind-before-advertise: ListenAndServe in a fire-and-forget goroutine used
// to hide bind failures in stdout only, while mDNS still pointed phones at a
// port this process didn't own — discovery "worked", pairing got connection
// refused (or hit a stale owner). "tcp4" deliberately, not "tcp": phones only
// dial IPv4 A records from RegisterProxy, and on macOS a dual-stack IPv6 bind
// can succeed while IPv4 traffic still goes to an AF_INET squatter.
func (s *Server) tryStartCompanion() bool { panic("fake") }

// Firewall helper can shell out — run after releasing companionMu.

func (s *Server) retryCompanionListen(ctx context.Context) { panic("fake") }

func (s *Server) publishCompanionAddresses() { panic("fake") }

// ensureWindowsCompanionFirewall best-effort adds an inbound allow rule for
// the companion port. Windows Defender Firewall commonly blocks inbound TCP
// to unsigned/local apps with no prompt when the phone is on another subnet
// hop — which looks exactly like "manual IP:8027 pairing is broken" while
// localhost curls still succeed. Failure is non-fatal (often needs elevation).
func ensureWindowsCompanionFirewall(port int, h *Handlers) { panic("fake") }

func (s *Server) Shutdown(ctx context.Context) error {
	panic(
		// KMA-233: emit agent_run_end + cancel armed runs before HTTP teardown so
		// Diagnostics is not silent when SIGTERM / desktop OnShutdown aborts mid-run.
		"fake")
}

func (s *Server) startMDNS() { panic("fake") }

// startMDNSLocked registers _chidori._tcp. Caller must hold companionMu
// when invoking from tryStartCompanion; the public startMDNS wrapper locks.
func (s *Server) startMDNSLocked() { panic("fake") }

// Advertise the companion listen port (8027 by default), not the IDE port.

// Prefer RegisterProxy with a sanitized single-label HostName and the
// real LAN IPs. plain Register calls os.Hostname(), and on macOS that
// returns "Name.local"; grandcat/zeroconf then appends another ".local."
// because HasSuffix(..., "local.") is false — yielding HostName
// "Name.local.local.", which Android NsdManager fails to resolve
// (info.host == null → discovered instance silently dropped).

// mdnsHostLabelFromOS returns a single DNS label suitable for zeroconf's
// HostName field (RegisterProxy will append ".local."). See startMDNS.
func mdnsHostLabelFromOS() string { panic("fake") }

// localIPv4Addresses returns every non-loopback IPv4 address currently
// assigned to this machine, across all interfaces — a Mac routinely has
// more than one (Wi-Fi, Ethernet, a VPN's virtual adapter), and only one is
// typically on the same LAN as the phone. Listing all of them turns
// "which IP do I type in" from a guess into a short, explicit list to try.
func localIPv4Addresses() []string { panic("fake") }

func (s *Server) stopMDNS() { panic("fake") }

// SyncLSPForWorkspace checks the current workspace root/trust state and
// starts, restarts, or stops the gopls bridge to match (ADR-0006: "start only
// when root contains go.mod AND the workspace is trusted"). Exported so
// internal/coordinator can call it once at startup for a workspace that was
// already configured+trusted from a previous run — WorkspaceSetRoot and
// WorkspaceTrust call the same underlying logic on every subsequent change.
func (s *Server) SyncLSPForWorkspace() { panic("fake") }

// SyncMCPServers reconciles the configured MCP servers (config.yaml's
// mcp_servers) against what's actually running (ADR-0012). Exported so
// internal/coordinator can call it once at startup, same reasoning as
// SyncLSPForWorkspace above — servers configured in a previous run should
// come up without the user needing to re-trigger anything from Settings.
func (s *Server) SyncMCPServers() { panic("fake") }

// SyncWorkflowCron starts or restarts cron workflow schedules for the current workspace.
func (s *Server) SyncWorkflowCron() { panic("fake") }

// SyncNetdocTools registers browse_fetch when experimental web fetch is enabled.
func (s *Server) SyncNetdocTools() { panic("fake") }

// requireAuth enforces a bearer key on state-changing requests. Read-only
// requests (GET/HEAD) and WebSocket upgrades stay open so dashboards and
// browser compute nodes work without credentials; everything that mutates
// state (chat submission, model swap, restart, dispatch, cache clear, …)
// requires Authorization: Bearer <key>.
//
// Companion pairing and desktop Companion Settings actions are exempt:
// the phone has no coordinator API key (it uses the post-pair companion
// bearer instead), and the Wails UI never attaches Authorization on
// /api/companion/*. Without this, enabling cfg.auth.api_key silently
// breaks all phone pairing and the Settings companion buttons.
func requireAuth(key string, next http.Handler) http.Handler { panic("fake") }

func companionAuthExempt(path string) bool { panic("fake") }

func logging(next http.Handler) http.Handler { panic("fake") }

// Wrap so response bytes are counted for the resource monitor's network
// readout. The wrapper preserves Flush/Hijack (see countingResponseWriter).
