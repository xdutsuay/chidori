// Package executil is the standard way to spawn subprocesses from chidori.
//
// Use Command / CommandContext instead of os/exec so Windows GUI builds do not
// flash a console when spawning git, gh, netsh, cmd.exe, and similar tools
// (KMA-158). On other OSes the helpers are identical to os/exec.
package executil

import (
	"context"
	"os/exec"
)

// Command wraps os/exec.Command and applies platform spawn policy.
func Command(name string, arg ...string) *exec.Cmd { panic("fake") }

// CommandContext wraps os/exec.CommandContext and applies platform spawn policy.
func CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd { panic("fake") }

// HideConsole applies Windows CREATE_NO_WINDOW policy to an existing Cmd.
// Prefer Command / CommandContext for new code; this exists for rare cases
// that must build the Cmd elsewhere first.
func HideConsole(cmd *exec.Cmd) { panic("fake") }
