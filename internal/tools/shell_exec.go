package tools

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/executil"
	ws "github.com/xdutsuay/lclreason/internal/workspace"
)

// ShellConfig controls shell tool behaviour and security.
type ShellConfig struct {
	Enabled    bool          `yaml:"enabled"`
	Timeout    time.Duration `yaml:"timeout"`     // max execution time (default 30s)
	MaxOutput  int           `yaml:"max_output"`  // max bytes of stdout+stderr to return (default 64KB)
	AllowList  []string      `yaml:"allow_list"`  // if non-empty, only these commands are allowed
	DenyList   []string      `yaml:"deny_list"`   // commands that are always blocked
	WorkingDir string        `yaml:"working_dir"` // cwd for commands (default: /tmp)
}

func defaultShellConfig() ShellConfig { panic("fake") }

// cappedWriter accumulates up to max bytes then discards the rest.
type cappedWriter struct {
	buf       bytes.Buffer
	max       int
	truncated bool
}

func (w *cappedWriter) Write(p []byte) (int, error) { panic("fake") }

func (w *cappedWriter) String() string { panic("fake") }

// ShellExecTool returns a tool that executes shell commands in a subprocess.
func ShellExecTool(cfg *ShellConfig) ToolFunc { panic("fake") }

// Security: check deny list

// Security: check allow list (if configured)

// Execute with timeout

// KMA-131: Windows has no sh/tmp by default — use cmd.exe like ConPTY.

func mergeShellOutput(stdout, stderr *cappedWriter, max int) string { panic("fake") }
