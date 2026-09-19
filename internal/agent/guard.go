package agent

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/xdutsuay/lclreason/internal/hooks"
	"github.com/xdutsuay/lclreason/internal/tools"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// FileReader reads workspace files for edit previews.
type FileReader interface {
	Read(rel string) ([]byte, error)
}

type guardOpts struct {
	proposeOnly   bool
	files         FileReader
	trusted       bool   // workspace trust — gates tools that execute arbitrary code (git_commit, shell_exec)
	workspaceRoot string // live root, injected as shell_exec's cwd so commands run in the workspace, not /tmp
	// hooks loads .lclreason/hooks.json for this call (ADR-0013). Nil means
	// hooks are not wired. Fire itself still no-ops when trusted is false.
	hooks func() *hooks.Registry
	// ownedPaths are paths claimed by a parent agent (ADR-0019). Write/patch
	// proposals for these paths are rejected in nested write subagents.
	ownedPaths map[string]bool
	// tracker records file reads/writes per session (HRD-001).
	tracker *workspace.FileTracker
}

// guardedRunner enforces agent write policies on top of a tool registry.
// reads is guarded by mu because KMA-209 runs parallel read-only tools on the
// same runner (concurrent map write is process-fatal — KMA-238).
type guardedRunner struct {
	inner ToolRunner
	mu    sync.Mutex
	reads map[string]bool
	opts  guardOpts
}

func newGuardedRunner(inner ToolRunner, opts guardOpts) *guardedRunner { panic("fake") }

func (g *guardedRunner) hasRead(key string) bool { panic("fake") }

func (g *guardedRunner) markRead(key string) { panic("fake") }

// codeExecutionTools can run arbitrary code (a git hook, a shell command) or
// drive a headless browser (navigation, clicks) — the same class of risk as
// agent auto-apply, gated the same way: untrusted workspaces always get
// rejected here, regardless of apply_mode, mirroring the UI/API's own git
// commit and switch-branch handlers.
var codeExecutionTools = map[string]bool{
	"git_commit": true, "shell_exec": true,
	"browse_navigate": true, "browse_snapshot": true, "browse_click": true,
	// KMA-120 — mutate remote GitHub state or checkout (git hooks).
	"github_pr_create": true, "github_issue_create": true,
	"github_issue_comment": true, "github_pr_checkout": true,
}

func (g *guardedRunner) Execute(ctx context.Context, name string, params map[string]any) (*tools.ToolResult, error) {
	panic("fake")
}

// Auto-read from disk when a workspace reader is available — models
// often skip read_file before write_file; failing the turn is worse
// than grounding the write in the file's current on-disk content.
// Missing path = create: allow write without a prior successful read
// (KMA-193). Without a files reader we cannot tell create vs edit.

// preToolHooks runs ADR-0013 pre_tool_call hooks. A veto (Aborted) blocks the
// tool; Fire already hard-gates on trust.
func (g *guardedRunner) preToolHooks(ctx context.Context, name string, params map[string]any) (blocked bool, reason string) {
	panic("fake")
}

// postToolHooks runs ADR-0013 post_tool_call hooks (fire-and-log; never blocks).
func (g *guardedRunner) postToolHooks(ctx context.Context, name string, params map[string]any, res *tools.ToolResult) {
	panic("fake")
}

func (g *guardedRunner) fireHooks(ctx context.Context, event string, env map[string]string) []hooks.HookResult {
	panic("fake")
}

func hookEnv(name string, params map[string]any, res *tools.ToolResult) map[string]string {
	panic("fake")
}

func (g *guardedRunner) readContent(path string) string { panic("fake") }

func (g *guardedRunner) proposeWrite(path, after, before string) (*tools.ToolResult, error) {
	panic("fake")
}

func (g *guardedRunner) proposePatch(path, old, newS string) (*tools.ToolResult, error) {
	panic("fake")
}

func normalizePath(p string) string { panic("fake") }
