package agent

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/xdutsuay/lclreason/internal/tools"
)

// LeaseManager provides in-memory file leases so concurrent subagents
// (especially write-capable ones under ADR-0019) cannot propose edits to
// the same path simultaneously. Each path is either unleased or held by
// exactly one ownerID.
type LeaseManager struct {
	mu     sync.Mutex
	leases map[string]string // path → ownerID
}

// NewLeaseManager creates an empty lease manager.
func NewLeaseManager() *LeaseManager { panic("fake") }

// Acquire claims path for ownerID. Returns an error if another owner holds it.
func (lm *LeaseManager) Acquire(path, ownerID string) error { panic("fake") }

// Release drops the lease on path if held by ownerID (or anyone if ownerID is empty).
func (lm *LeaseManager) Release(path, ownerID string) { panic("fake") }

// ReleaseAll drops every lease held by ownerID.
func (lm *LeaseManager) ReleaseAll(ownerID string) { panic("fake") }

// IsHeld reports whether path is currently leased (and by whom).
func (lm *LeaseManager) IsHeld(path string) (holder string, held bool) { panic("fake") }

// Held returns a snapshot of all currently held leases.
func (lm *LeaseManager) Held() map[string]string { panic("fake") }

// writeLeaseRunner acquires a process-wide file lease before write_file /
// apply_patch so concurrent agent runs cannot silently stomp the same path
// (KMA-140). Other tools pass through unchanged.
type writeLeaseRunner struct {
	inner   ToolRunner
	leases  *LeaseManager
	ownerID string
}

func (w writeLeaseRunner) Execute(ctx context.Context, name string, params map[string]any) (*tools.ToolResult, error) {
	panic("fake")
}
