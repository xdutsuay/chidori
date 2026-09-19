package workspace

import (
	"os"
	"path/filepath"
	"strings"
)

// EnvVarsFromRoot reads KEY=VALUE lines from .lclreason/env under root (T.11).
// Used by the UI terminal PTY and agent shell_exec so builds share the same env.
func EnvVarsFromRoot(root string) []string { panic("fake") }
