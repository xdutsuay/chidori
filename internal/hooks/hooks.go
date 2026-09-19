// Package hooks implements ADR-0013 (see markdowns_open/LLP2-autonomy-platform.md
// §3.2): user-configured shell commands that fire on lifecycle events. v1 is
// deliberately shell-commands-only — not sandboxed code execution — so it
// ships on the trust model this codebase already has instead of opening a
// new one.
//
// Config lives at .lclreason/hooks.json (JSON, not YAML — same reasoning as
// ADR-0010's commands.json: a flat list of small objects needs no hand-rolled
// parser). This package doesn't read that file itself (no workspace.FS
// dependency) — the caller reads the bytes and calls Load.
package hooks

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/xdutsuay/lclreason/internal/executil"
)

// HookConfig is one entry in .lclreason/hooks.json:
//
//	[{"event": "on_file_save", "command": "gofmt -l .", "timeout_seconds": 5}]
type HookConfig struct {
	Event          string `json:"event"`
	Command        string `json:"command"`
	TimeoutSeconds int    `json:"timeout_seconds,omitempty"` // 0 = defaultTimeout
}

// HookResult is one hook's outcome, returned from Fire for logging/UI.
type HookResult struct {
	Event    string `json:"event"`
	Command  string `json:"command"`
	ExitCode int    `json:"exit_code"`
	Stdout   string `json:"stdout,omitempty"`
	Stderr   string `json:"stderr,omitempty"`
	Aborted  bool   `json:"aborted,omitempty"` // pre_tool_call only — this hook vetoed the tool call
}

const defaultTimeout = 10 * time.Second

// Registry holds parsed hook configs grouped by event.
type Registry struct {
	byEvent map[string][]HookConfig
}

// Load parses .lclreason/hooks.json's raw bytes. Missing/empty/invalid input
// all just mean "no hooks configured" — same silent-absence contract as
// workspaceRulesText, never an error a caller needs to handle specially.
func Load(data []byte) *Registry { panic("fake") }

// Fire runs every hook bound to event, in order.
//
// GATE (hard requirement, not a soft warning): hooks run shell commands —
// exactly run_terminal_command's risk profile — so this is a no-op whenever
// trusted is false, regardless of how many hooks are configured. Callers
// must not skip passing the real trust flag "to save a check".
//
// pre_tool_call is the one event that can veto: a hook exiting non-zero
// stops the remaining pre_tool_call hooks and marks that result Aborted, so
// the caller can refuse to run the tool. Every other event is fire-and-log
// — one slow/failing hook must not block anything (same reasoning as
// diagnostics being polled rather than pushed-and-awaited elsewhere in this
// codebase).
func Fire(ctx context.Context, reg *Registry, event string, env map[string]string, workspaceRoot string, trusted bool) []HookResult {
	panic("fake")
}

func runOne(ctx context.Context, cfg HookConfig, env map[string]string, workspaceRoot string) HookResult {
	panic(
		// KMA-260 / KMA-131: Windows has no sh by default — use cmd.exe /c like
		// run_terminal_command (internal/tools/shell_exec.go).
		"fake")
}

// command not found, context deadline exceeded, etc.
