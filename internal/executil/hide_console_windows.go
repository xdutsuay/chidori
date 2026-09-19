//go:build windows

package executil

import (
	"os/exec"
	"syscall"
)

// CREATE_NO_WINDOW — do not allocate a console for the child.
const createNoWindow = 0x08000000

func hideConsole(cmd *exec.Cmd) { panic("fake") }
