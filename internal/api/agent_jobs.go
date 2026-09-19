package api

import (
	"context"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/agentjobs"
)

// AgentJobStatus is the coarse lifecycle of an AG.16 background agent job.
type AgentJobStatus string

const (
	AgentJobRunning   AgentJobStatus = "running"
	AgentJobCompleted AgentJobStatus = "completed"
	AgentJobFailed    AgentJobStatus = "failed"
	AgentJobCancelled AgentJobStatus = "cancelled"
)

// AgentJob is the wire shape for GET /api/agent/jobs (AG.16 v1).
type AgentJob struct {
	ID          string         `json:"id"`
	Status      AgentJobStatus `json:"status"`
	Task        string         `json:"task"`
	Started     time.Time      `json:"started"`
	Finished    time.Time      `json:"finished,omitempty"`
	Error       string         `json:"error,omitempty"`
	LeasedPaths []string       `json:"leased_paths,omitempty"`
}

type agentJobEntry struct {
	AgentJob
	cancel func()
}

// agentJobRegistry tracks detached AgentRun jobs in memory and on disk.
type agentJobRegistry struct {
	mu          sync.Mutex
	jobs        map[string]*agentJobEntry
	persistRoot string
}

func newAgentJobRegistry() *agentJobRegistry { panic("fake") }

func (r *agentJobRegistry) SetPersistRoot(root string) { panic("fake") }

func (r *agentJobRegistry) loadPersisted() { panic("fake") }

func (r *agentJobRegistry) persistLocked(e *agentJobEntry) { panic("fake") }

func (r *agentJobRegistry) Start(id, task string, cancel func()) { panic("fake") }

func (r *agentJobRegistry) Finish(id string, status AgentJobStatus, errMsg string) { panic("fake") }

func (r *agentJobRegistry) IsRunning(id string) bool { panic("fake") }

func (r *agentJobRegistry) Cancel(id string) bool { panic("fake") }

func (r *agentJobRegistry) List() []AgentJob { panic("fake") }

func (h *Handlers) sharedLeases() *agent.LeaseManager { panic("fake") }

func (h *Handlers) sessionScratch(sessionID string) string { panic("fake") }

func (h *Handlers) persistSessionScratch(sessionID string) func(string) { panic("fake") }
