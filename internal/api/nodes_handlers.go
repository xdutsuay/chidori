package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/registry"
	"github.com/xdutsuay/lclreason/internal/store"
)

// ---- Inference Source: attach-by-IP + LAN discovery ----
//
// These endpoints let a bare OpenAI-compatible or Ollama server (LM Studio,
// vLLM, llama.cpp server, Ollama, AirLLM — anything answering GET /v1/models
// or GET /api/tags) join the worker pool as a "passive" node: the coordinator
// polls it directly (see registry.Registry.HealthCheck, wired into the sweep
// loop in internal/coordinator/run.go) instead of requiring the endpoint to
// push its own heartbeats like a real lclreason worker does. No extra
// software is required on the target machine.

// discoverPorts lists the ports + probe styles checked during LAN discovery.
// "ollama" probes GET /api/tags on the bare host:port; anything else probes
// GET /v1/models (OpenAI-compatible).
var discoverPorts = []struct {
	port    int
	backend string
}{
	{11434, "ollama"},
	{1234, "lmstudio"}, // LM Studio default
	{8000, "airllm"},   // AirLLM / common vLLM default
	{8080, "lmstudio"}, // generic OpenAI-compatible server (llama.cpp server, etc.)
}

type attachNodeRequest struct {
	Address string `json:"address"`            // "host:port" or a full URL
	Backend string `json:"backend,omitempty"`  // optional hint: "ollama" | "lmstudio" | "airllm"
	APIBase string `json:"api_base,omitempty"` // optional explicit override, e.g. "http://host:1234/v1"
}

// AttachNode probes a user-supplied address, detects what's running there
// (or trusts an explicit backend/api_base hint), and registers it as a
// worker-pool node. This is a direct user action (typed into Settings), not
// something an agent can trigger on its own, so — consistent with git
// stage/unstage and the existing bare RegisterNode endpoint — it is not
// trust-gated.
func (h *Handlers) AttachNode(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Reject attaching chidori's own coordinator as if it were an inference
// server (see isOwnCoordinator). This is the classic "I attached
// 192.168.1.4:8080 and it shows Healthy with no models" trap — 8080 is the
// coordinator's own port, and LM Studio's default is 1234.

// One machine, one inference source: reject attaching a second node at
// an IP that's already registered under a different node, regardless of
// port. On cheap hardware a single box can't usefully run two separate
// model-serving processes anyway, and showing both as independent
// capacity in the dashboard is actively misleading (see
// llpcodefeature_Cursorclone.md's node-dedupe fix). Re-attaching the
// *same* host:port (same id) is still allowed — that's just a refresh.

type discoverCandidate struct {
	Address string   `json:"address"`
	Backend string   `json:"backend"`
	APIBase string   `json:"api_base"`
	Models  []string `json:"models"`
}

// DiscoverNodes scans the local /24 subnet for known LLM ports and reports
// what it finds; it does not register anything — the user reviews the list
// in Settings and attaches the ones they want via AttachNode. Read-only
// network probing, so — like AttachNode — no trust gate.
func (h *Handlers) DiscoverNodes(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Skip hosts that already have a registered node (any port) — one
// machine, one inference source (see the matching check in AttachNode).
// Filtering here means an already-attached machine doesn't even show up
// as a "discovered" candidate a second time.

// Probe this machine itself on loopback first. LM Studio and Ollama very
// often bind only 127.0.0.1 (LM Studio's "Serve on Local Network" is off by
// default), so a pure subnet scan of 192.168.x.y never sees them — the
// #1 reported "it can't find my local LM Studio" case. isOwnCoordinator
// filtering downstream keeps chidori's own :8080 out of the results.

// Don't surface chidori's own coordinator (its :8080 answers
// /v1/models) as a discoverable "worker node" — attaching it
// would loop back into this app.

// probeOne checks a single host:port. If backendHint is set, only that
// probe style is tried (used by discovery, which already knows which style
// goes with which well-known port); an empty hint tries Ollama first, then
// OpenAI-compatible (used by AttachNode when the user doesn't specify one).
func probeOne(ctx context.Context, host string, port int, backendHint string) (backend, apiBase string, models []string, ok bool) {
	panic("fake")
}

// splitHostPort accepts "host:port" or a full URL and returns the bare host,
// numeric port, and normalized "host:port" form.
func splitHostPort(addr string) (host string, port int, hostPort string, err error) { panic("fake") }

// hostOnly strips a port off "host:port" (or a bare "http://host:port" URL,
// via the same prefix trimming splitHostPort uses); returns addr unchanged
// if it has no parseable port (e.g. already just a bare host).
func hostOnly(addr string) string { panic("fake") }

// isOwnCoordinator reports whether host:port is chidori's own coordinator API,
// not a separate inference server. This matters because the coordinator's
// default port (8080) is ALSO in the attach/discovery probe list as a generic
// OpenAI-compatible endpoint (llama.cpp etc.), and the coordinator answers
// GET /v1/models on it — so probing its own address returns 200 and the node
// gets marked "Healthy" while actually being itself. Dispatching to it would
// loop straight back into the coordinator. host matches when it's loopback,
// "localhost", or any of this machine's own interface IPs (e.g. its LAN
// 192.168.x.y), since all of those route back here.
func (h *Handlers) isOwnCoordinator(host string, port int) bool { panic("fake") }

// nodeAtHost returns an already-registered node whose address resolves to
// the same host as host, other than the node identified by selfID (so
// re-attaching/refreshing the same manual entry isn't treated as a
// conflict with itself). Returns nil if no other node is at that host.
func (h *Handlers) nodeAtHost(host, selfID string) *registry.Node { panic("fake") }

// localSubnet returns this machine's own LAN /24 prefix (e.g. "192.168.1")
// by opening a UDP "connection" to a public IP — no packets are actually
// sent, this just asks the OS which local interface/address would be used,
// which is a standard trick for finding the outbound-facing local IP.
func localSubnet() string { panic("fake") }
