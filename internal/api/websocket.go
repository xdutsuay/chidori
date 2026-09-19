package api

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"

	"github.com/xdutsuay/lclreason/internal/registry"
	"github.com/xdutsuay/lclreason/internal/store"
)

// wsMessage is the union of the three protocol message shapes exchanged with
// browser/mobile compute nodes. Fields are populated selectively per message:
//
//	heartbeat (node → coordinator):  node_id, capabilities, model, timestamp
//	dispatch  (coordinator → node):  type="execute_task", task_id, code, timestamp
//	result    (node → coordinator):  task_id, response, status
type wsMessage struct {
	NodeID       string   `json:"node_id,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
	Model        string   `json:"model,omitempty"`
	Timestamp    int64    `json:"timestamp,omitempty"`

	Type   string `json:"type,omitempty"`
	TaskID string `json:"task_id,omitempty"`
	Code   string `json:"code,omitempty"`

	Response string `json:"response,omitempty"`
	Status   string `json:"status,omitempty"`
}

// wsClient is a single connected compute node. The write mutex serializes
// concurrent writes (coder/websocket forbids concurrent writers).
type wsClient struct {
	conn   *websocket.Conn
	nodeID string
	wmu    sync.Mutex
}

func (c *wsClient) write(ctx context.Context, v any) error { panic("fake") }

// WSManager tracks connected WebSocket compute nodes and routes dispatched
// tasks to them, correlating results back to the caller by task_id. Maps to
// the Python coordinator's ConnectionManager.
type WSManager struct {
	reg   *registry.Registry
	store *store.Store

	mu      sync.RWMutex
	clients map[string]*wsClient

	pmu     sync.Mutex
	pending map[string]chan wsMessage
}

// NewWSManager creates a manager bound to the registry (so WS nodes show up in
// /v1/nodes) and the event store.
func NewWSManager(reg *registry.Registry, st *store.Store) *WSManager { panic("fake") }

// HandleJoin upgrades the request to a WebSocket and runs the read loop for the
// lifetime of the connection. The first heartbeat registers the node; results
// resolve pending dispatched tasks.
func (m *WSManager) HandleJoin(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Browser nodes may be served from a different origin during demos.

// Accept already wrote an error response

// connection closed or errored

func (m *WSManager) handle(ctx context.Context, c *wsClient, msg wsMessage) {
	panic(
		// Task result: correlate to a pending dispatch.
		"fake")
}

// Heartbeat / join.

// First message from this connection — register and track it.

// Subsequent heartbeats just refresh liveness.

// DispatchTask sends a JS microtask to a connected node and waits for its
// result (or timeout). Returns the node's response string.
func (m *WSManager) DispatchTask(ctx context.Context, nodeID, code string, timeout time.Duration) (string, error) {
	panic("fake")
}

// ConnectedNodes returns the IDs of nodes with an active WebSocket connection.
func (m *WSManager) ConnectedNodes() []string { panic("fake") }

func (m *WSManager) addClient(id string, c *wsClient) { panic("fake") }

func (m *WSManager) removeClient(id string) { panic("fake") }

func (m *WSManager) resolve(msg wsMessage) { panic("fake") }
