//go:build windows

package terminal

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/UserExistsError/conpty"
)

func startSession(cwd, shellOverride string) (*Session, func(), error) { panic("fake") }

// Best-effort: enter workspace and apply .lclreason/env (ConPTY Start has no cwd arg).

// Upstream conpty does not export Resize; initial size is set at Start.

func overlayShell(override string) string { panic("fake") }

func windowsShell(override string) (string, []string) { panic("fake") }

func buildCmdLine(shell string, args []string) string { panic("fake") }

func quoteArgs(args []string) []string { panic("fake") }

func quoteWinPath(p string) string { panic("fake") }
