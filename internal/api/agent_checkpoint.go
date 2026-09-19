package api

import (
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/agent"
)

const agentCheckpointTTL = 2 * time.Hour

type agentCheckpointEntry struct {
	cp      agent.Checkpoint
	expires time.Time
}

type agentCheckpointStore struct {
	mu sync.Mutex
	m  map[string]agentCheckpointEntry
}

func newAgentCheckpointStore() *agentCheckpointStore { panic("fake") }

func (s *agentCheckpointStore) Put(taskID string, cp agent.Checkpoint) { panic("fake") }

func (s *agentCheckpointStore) Take(taskID string) (agent.Checkpoint, bool) { panic("fake") }
