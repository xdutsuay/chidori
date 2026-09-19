//go:build !windows

package executil

import "os/exec"

func hideConsole(cmd *exec.Cmd) { panic("fake") }
