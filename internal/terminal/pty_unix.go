//go:build !windows

package terminal

import (
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func startSession(cwd, shellOverride string) (*Session, func(), error) { panic("fake") }
