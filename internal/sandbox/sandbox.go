package sandbox

import (
	"context"
	"fmt"

	"github.com/xdutsuay/lclreason/internal/workspace"
)

// Mode defines the sandboxing isolation mechanism.
type Mode string

const (
	ModeDisabled Mode = "disabled"
	ModeSnapshot Mode = "snapshot" // Phase 1: Git stash-based safety net
	ModeWorktree Mode = "worktree" // Phase 2: Git worktree isolation (unsupported in this release)
	ModeSeatbelt Mode = "seatbelt" // Phase 3: macOS sandbox-exec process isolation (unsupported in this release)
)

// Sandbox defines the lifecycle interface for a workspace sandbox.
type Sandbox interface {
	Enter(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Close() error
}

// Config specifies configuration for the sandbox subsystem.
type Config struct {
	Enabled             bool   `json:"enabled" yaml:"enabled"`
	Mode                Mode   `json:"mode" yaml:"mode"`
	AutoRollbackOnError bool   `json:"auto_rollback_on_error" yaml:"auto_rollback_on_error"`
	WorkspaceRoot       string `json:"workspace_root" yaml:"workspace_root"`
}

// New creates an appropriate Sandbox instance based on the configuration.
func New(cfg Config, ws *workspace.FS) (Sandbox, error) { panic("fake") }

type noopSandbox struct{}

func (n *noopSandbox) Enter(ctx context.Context) error    { panic("fake") }
func (n *noopSandbox) Commit(ctx context.Context) error   { panic("fake") }
func (n *noopSandbox) Rollback(ctx context.Context) error { panic("fake") }
func (n *noopSandbox) Close() error                       { panic("fake") }
