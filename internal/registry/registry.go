package registry

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/backendprobe"
)

type Node struct {
	ID           string    `json:"id"`
	Address      string    `json:"address"`
	Models       []string  `json:"models"`
	Healthy      bool      `json:"healthy"`
	LastSeen     time.Time `json:"last_seen"`
	Load         float64   `json:"load"` // 0.0-1.0
	Capabilities []string  `json:"capabilities,omitempty"`
	ClientIP     string    `json:"client_ip,omitempty"`
	CurrentTask  string    `json:"current_task,omitempty"`
	APIBase      string    `json:"api_base,omitempty"`
	Backend      string    `json:"backend,omitempty"`
	// APIKey is the OpenAI-compat data-plane bearer for this node (KMA-65).
	// Never serialized on /v1/nodes.
	APIKey string `json:"-"`
	// DisplayName is a human label (e.g. phone model name from chidori-nagasa
	// node mode). Empty for ordinary attached workers.
	DisplayName string `json:"display_name,omitempty"`
	BytesSent   int64  `json:"bytes_sent,omitempty"`
	// TokensIn/TokensOut are cumulative prompt/completion tokens this node has
	// served since the coordinator started (Planner panel's per-node token
	// monitor). Only incremented on the non-streaming dispatch path — see
	// InvokeResult's own doc comment in internal/dispatch/dispatcher.go.
	TokensIn  int64 `json:"tokens_in,omitempty"`
	TokensOut int64 `json:"tokens_out,omitempty"`

	// pendingCmd is a one-shot control command (e.g. "restart") delivered to
	// the node on its next heartbeat. Unexported so it never leaks into the
	// /v1/nodes JSON listing.
	pendingCmd string
}

// HasCapability reports whether the node advertises cap.
func (n *Node) HasCapability(cap string) bool { panic("fake") }

// HasModel reports whether the node has loaded the given model.
func (n *Node) HasModel(model string) bool { panic("fake") }

// HasReachableEndpoint reports whether the node has an address or api_base
// the dispatcher can actually dial. Nodes registered with neither (KMA-92)
// must not enter pickRanked even if they were marked Healthy at insert time.
func (n *Node) HasReachableEndpoint() bool { panic("fake") }

type Registry struct {
	mu           sync.RWMutex
	nodes        map[string]*Node
	offlineAfter time.Duration
}

func New(offlineAfter time.Duration) *Registry { panic("fake") }

func (r *Registry) Register(n *Node) { panic("fake") }

// Heartbeat updates load, optional current_task, and bytes_sent counters.
// Pass empty strings / zero to leave a field unchanged.
func (r *Registry) Heartbeat(id string, load float64) error { panic("fake") }

// HeartbeatFull updates extended heartbeat fields. Backwards-compatible:
// callers using Heartbeat() still work.
func (r *Registry) HeartbeatFull(id string, load float64, currentTask string, bytesSent int64) error {
	panic("fake")
}

func (r *Registry) Deregister(id string) { panic("fake") }

// deepCopyNode returns a heap-allocated copy of n, including its slice
// fields (Models, Capabilities) copied to fresh backing arrays — never
// shared with the live registry entry's slices.
//
// running_issue.md P0-2 / H.3: Healthy()/All()/Get()/GetByCapability() used
// to return the *live* `*Node` stored in r.nodes, so a caller (dispatch's
// pickRanked, invokeNode, etc.) read `.Load`/`.Healthy`/`.Models` on that
// same pointer with no lock held, while Heartbeat/HealthCheck/Sweep mutated
// those exact fields under r.mu from the coordinator's background loop — a
// confirmed data race (caught with a targeted `-race` probe; the ordinary
// test suite never exercised concurrent read-during-write and so never
// caught it). Every reader below now gets its own private snapshot instead;
// nothing outside this file ever holds a pointer into the live map again.
func deepCopyNode(n *Node) *Node { panic("fake") }

// Healthy returns all nodes that are marked healthy AND have been seen
// recently (heartbeat or successful HealthCheck probe — both refresh
// LastSeen). Checking only LastSeen let freshly-registered nodes whose
// backend the prober had already found dead (Healthy=false, LastSeen=just
// registered) into the dispatch pool.
func (r *Registry) Healthy() []*Node { panic("fake") }

func (r *Registry) All() []*Node { panic("fake") }

// Get returns a single node by id, or nil if not present. The returned
// pointer is a private snapshot (see deepCopyNode) — mutating it has no
// effect on the registry; use the Heartbeat/UpdateModel/etc. setters instead.
func (r *Registry) Get(id string) *Node { panic("fake") }

// AddTokens accumulates a successful invocation's token usage onto a node's
// running totals (Planner panel's per-node token monitor). No-ops silently
// for an unknown id or zero counts — called unconditionally from the
// dispatch hot path, so it must never be the reason a request fails.
func (r *Registry) AddTokens(id string, promptTokens, completionTokens int) { panic("fake") }

// Sweep marks stale nodes offline; called periodically.
func (r *Registry) Sweep() { panic("fake") }

// UpdateModel swaps the active model on a node. If the model is not yet in
// Models it is prepended so it shows up as the primary model.
func (r *Registry) UpdateModel(id, model string) error { panic("fake") }

// Move to front to mark as primary.

// RequestRestart queues a one-shot "restart" command for a node. The command
// is delivered the next time the node heartbeats (see ConsumeCommand). Returns
// an error if the node isn't registered.
func (r *Registry) RequestRestart(id string) error { panic("fake") }

// ConsumeCommand returns and clears any pending control command for a node.
// Returns "" when there is nothing queued. Called from the heartbeat handler so
// the worker learns about the command on its next beat.
func (r *Registry) ConsumeCommand(id string) string { panic("fake") }

// GetByCapability returns healthy nodes that advertise a given capability.
func (r *Registry) GetByCapability(cap string) []*Node { panic("fake") }

// Same liveness rule as Healthy(): marked healthy AND seen recently.

// HealthCheck probes each node's address with a short HTTP GET and marks
// unreachable nodes Unhealthy. Reachable means "any HTTP response received" —
// status code is not inspected so the check works against Ollama, LM Studio,
// and the worker's own /health endpoint without special-casing.
//
// A successful probe also refreshes LastSeen — a direct "are you there" answer
// counts the same as a heartbeat. This is what keeps passive nodes (attached
// by IP or found via LAN discovery, which never run a heartbeat client) inside
// Healthy()'s LastSeen window and out of Sweep()'s cutoff. Before this,
// HealthCheck set only the Healthy flag: Sweep would flip a reachable passive
// node back to unhealthy within seconds (LastSeen never advanced past attach
// time), and Healthy() — what the dispatcher routes on — excluded it entirely,
// so "attach a worker by IP" silently stopped dispatching after offline_after.
//
// Probes run in parallel: with the sequential loop, N slow/dead nodes took
// N×3s, longer than the probe cadence.
func (r *Registry) HealthCheck(ctx context.Context) error { panic("fake") }
