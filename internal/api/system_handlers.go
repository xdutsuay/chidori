package api

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// netBytesIn/Out are process-wide counters for the in-app resource monitor's
// network readout. Incremented by the logging middleware (bytes in from each
// request's ContentLength, bytes out via countingResponseWriter). The frontend
// polls SystemStats and derives KB/s from the delta between two samples.
var (
	netBytesIn  atomic.Int64
	netBytesOut atomic.Int64

	// lastUIActivityMs is set by POST /api/ui/activity (and optionally other
	// UI-facing routes). The coordinator health probe backs off when stale.
	lastUIActivityMs atomic.Int64

	psCache struct {
		mu       sync.Mutex
		ts       time.Time
		cpu, mem float64
		rssKB    int64
	}
)

// TouchUIActivity records that the desktop UI is interactive. Used to slow
// the coordinator's node/backend health probes while the app is idle/hidden.
func TouchUIActivity() { panic("fake") }

// UIIdleFor reports whether the last UI activity signal is older than d.
// Before any signal (t==0) we treat the UI as active so startup probes stay
// at the configured heartbeat cadence.
func UIIdleFor(d time.Duration) bool { panic("fake") }

// UIActivity accepts an empty POST from the IDE to refresh lastUIActivityMs.
func (h *Handlers) UIActivity(w http.ResponseWriter, r *http.Request) { panic("fake") }

var appFootprintCache struct {
	mu     sync.Mutex
	ts     time.Time
	usedKB int64
}

// SystemStats reports THIS process's resource footprint for the monitor widget
// (not the whole machine). CPU%/RSS come from `ps` (cached ~5s), app storage is
// chidori's own footprint on disk (skipped for ?lite=1), and network comes
// from the app's own served-byte counters.
func (h *Handlers) SystemStats(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Phase 0 lite polls skip the disk walk, but still serve the last
// cached footprint so the title-bar SSD readout doesn't stick on "—"
// forever after the 10s lite cadence took over.

// SystemVersion answers Help ▸ "Coordinator Health & Version" (and the same
// info shown in Settings → Inference Source, next to the live Coordinator
// status). VCS fields come from Go's own automatic build-info stamping
// (runtime/debug.ReadBuildInfo, no ldflags/version-injection setup needed —
// confirmed against a real `go build` of this actual module before relying
// on it: a one-off standalone file showed no VCS fields at all, since Go
// only stamps this for a build run from inside a real VCS checkout) —
// empty strings if built without VCS info available (e.g. from a source
// tarball with no .git directory).
func (h *Handlers) SystemVersion(w http.ResponseWriter, r *http.Request) { panic("fake") }

// processCPUMem / processCPUTimeSecs are OS-specific
// (system_handlers_process_unix.go / system_handlers_process_windows.go).

// processCPUMemCached returns processCPUMem results, reused for up to 5s so
// the title-bar poll does not sample too aggressively on every tick.
func processCPUMemCached() (cpuPct, memPct float64, rssKB int64) { panic("fake") }

// appFootprintKB reports chidori's own on-disk footprint (bundle/binary +
// config/db/log roots), memoized briefly since recursive walks are too
// expensive to run on every 2s UI poll.
func appFootprintKB(h *Handlers) int64 { panic("fake") }

func appFootprintKBNow(h *Handlers) int64 { panic("fake") }

func dedupeRoots(paths []string) []string { panic("fake") }

// If the new path is a parent of an existing one, drop the child.

func pathSizeBytes(path string) int64 { panic("fake") }

// countingResponseWriter tallies response bytes for the network monitor while
// preserving the two things this server relies on the raw writer for:
// streaming (chat SSE, via http.Flusher/ResponseController) and connection
// upgrade (terminal WebSocket, via http.Hijacker). Unwrap() lets
// http.ResponseController reach the real writer; the explicit Flush/Hijack
// methods cover any direct type assertion.
type countingResponseWriter struct {
	http.ResponseWriter
}

func (c *countingResponseWriter) Write(b []byte) (int, error) { panic("fake") }

func (c *countingResponseWriter) Unwrap() http.ResponseWriter { panic("fake") }

func (c *countingResponseWriter) Flush() { panic("fake") }

func (c *countingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) { panic("fake") }
