package sandbox

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/xdutsuay/lclreason/internal/executil"
)

// SnapshotSandbox implements Phase 1 git stash/snapshot based sandboxing.
type SnapshotSandbox struct {
	mu                   sync.Mutex
	workspaceRoot        string
	snapshotRef          string
	preExistingUntracked map[string]bool
	active               bool
}

// NewSnapshotSandbox creates a new SnapshotSandbox instance.
func NewSnapshotSandbox(workspaceRoot string) (*SnapshotSandbox, error) { panic("fake") }

// getUntrackedFiles lists untracked files relative to workspace root.
func (s *SnapshotSandbox) getUntrackedFiles(ctx context.Context) (map[string]bool, error) {
	panic("fake")
}

// Enter captures a git snapshot of the current workspace state.
func (s *SnapshotSandbox) Enter(ctx context.Context) error { panic("fake") }

// Inventory pre-existing untracked files

// Run git stash create -u to get a commit hash including untracked changes

// Commit accepts the current workspace changes and drops the snapshot.
func (s *SnapshotSandbox) Commit(ctx context.Context) error { panic("fake") }

// Rollback restores the workspace back to the pre-turn snapshot while preserving pre-existing untracked files.
func (s *SnapshotSandbox) Rollback(ctx context.Context) error { panic("fake") }

// 1. Reset tracked files

// 2. Clean only untracked files created during the turn (preserve pre-existing untracked)

// 3. Restore pre-turn dirty state if a snapshot ref was created

// Close releases sandbox resources.
func (s *SnapshotSandbox) Close() error { panic("fake") }
