package tools

import (
	"context"
	"strings"

	"github.com/xdutsuay/lclreason/internal/gitutil"
)

// RegisterGitTools adds git-mutating tools scoped to the live workspace root.
// rootFn returns the CURRENT root on every call (the root is runtime-mutable
// via workspace.FS.SetRoot — Open Folder can point it elsewhere while the
// coordinator keeps running — so a *gitutil.Repo can't be captured once at
// startup; same pattern as RegisterWorkspaceTools taking wsFS itself).
//
// Trust gating happens one layer up, in the agent loop's guardedRunner
// (internal/agent/guard.go) — the same place write_file/apply_patch are
// gated — not here: this tool is reachable from both the agent loop and the
// planner's tool executor, and only the former has a trust concept to gate
// on. Mechanism lives with the mutating action; policy lives with the caller
// that knows about trust.
func RegisterGitTools(r *Registry, rootFn func() string) { panic("fake") }

func gitCommitTool(rootFn func() string) ToolFunc { panic("fake") }
