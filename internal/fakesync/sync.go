package fakesync

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// DefaultDenylist returns path substrings always skipped during sync/plan.
func DefaultDenylist() []string { panic("fake") }

// Plan is dry-run sync: same mapping rules as Sync with DryRun forced; Written is always 0.
func Plan(opts Options) (Result, error) { panic("fake") }

// Sync walks SourceRoot (or Allowlist dirs), skips Denylist matches and non-selected
// paths, and emits stubbed Go files into DestRoot (or only a mapping when DryRun).
// Empty Allowlist means the entire SourceRoot tree. Empty Selected means all
// non-denied candidates. Non-empty Selected keeps only matching relative paths
// or path prefixes (slash-separated).
func Sync(opts Options) (Result, error) { panic("fake") }

// Skip denylisted directory names early (vendor, .git, …).

func denied(slashPath string, denylist []string) bool { panic("fake") }

func selected(slashRel string, selectedPaths []string) bool { panic("fake") }

func stubGoFile(path string) ([]byte, error) { panic("fake") }

// RelSlash returns the slash-separated path of m.Src relative to sourceRoot.
func RelSlash(sourceRoot, src string) (string, error) { panic("fake") }
