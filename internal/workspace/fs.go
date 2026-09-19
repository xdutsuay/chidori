// Package workspace provides sandboxed file access for the /code IDE.
// All paths are resolved relative to a configured root; traversal outside
// the root is rejected.
package workspace

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// Config controls workspace file access.
type Config struct {
	Root         string   `yaml:"root"`
	MaxFileBytes int      `yaml:"max_file_bytes"`
	DenyPaths    []string `yaml:"deny_paths"`
}

// Entry is one node in a workspace tree listing.
type Entry struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	IsDir bool   `json:"is_dir"`
}

// FS is a sandboxed view of a directory on disk. Root, maxFileBytes and deny
// can be swapped at runtime via SetRoot (e.g. "Open Folder" from the IDE), so
// every long-lived holder of an *FS (HTTP handlers, registered agent tools)
// picks up the new root without needing to be rewired.
type FS struct {
	mu            sync.RWMutex
	root          string
	contentPrefix string // e.g. "lclreason/" when project lives in root/lclreason/
	maxFileBytes  int
	deny          []string
}

// NewFS creates a workspace FS. An empty root means the workspace is disabled
// until SetRoot is called later (e.g. from the Open Folder picker).
func NewFS(cfg Config) (*FS, error) { panic("fake") }

// SetRoot reconfigures the workspace root in place. Passing an empty
// cfg.Root disables the workspace. All existing references to this *FS
// (tool registrations, HTTP handlers) observe the change immediately.
func (f *FS) SetRoot(cfg Config) error { panic("fake") }

// detectContentPrefix returns a workspace-relative prefix when the user opened
// a parent folder but the project (go.mod) lives in a same-named subdirectory
// (e.g. root=/.../lclreason, project=/.../lclreason/lclreason).
func detectContentPrefix(root string) string { panic("fake") }

// Enabled reports whether a workspace root is configured.
func (f *FS) Enabled() bool { panic("fake") }

// Root returns the absolute workspace root, or "" when disabled.
func (f *FS) Root() string { panic("fake") }

// snapshot returns a consistent (root, contentPrefix, maxFileBytes, deny) quad under lock.
func (f *FS) snapshot() (string, string, int, []string) { panic("fake") }

// Resolve maps a relative path to an absolute path under the workspace root.
func (f *FS) Resolve(rel string) (string, error) { panic("fake") }

// NormalizeRel converts model-supplied paths to workspace-relative form.
func (f *FS) NormalizeRel(rel string) (string, error) { panic("fake") }

// Cursor-style @workspace/path — the @tag is a workspace label, not a folder.

// Model sometimes prefixes paths with the workspace folder name. When the
// project lives directly under root, strip it; when it lives in a nested
// same-named subdir (contentPrefix), strip so Resolve can re-apply the prefix.

// Absolute path under workspace root.

// Model echoed full root path without leading slash.

func deniedIn(deny []string, rel string) bool { panic("fake") }

// Also deny if any path segment starts with denied name as directory prefix.

// Tree lists immediate children of rel (non-recursive). Use recursive client calls or ListTree.
func (f *FS) Tree(rel string) ([]Entry, error) { panic("fake") }

// Read returns file contents for a relative path.
// ModTime returns a file's last-modified time — used to detect external
// changes (BE.12: the file watcher polls this for the active file rather
// than pulling in an fsnotify dependency, matching this codebase's existing
// preference for avoiding new deps where a simple poll suffices).
func (f *FS) ModTime(rel string) (time.Time, error) { panic("fake") }

func (f *FS) Read(rel string) ([]byte, error) { panic("fake") }

// ReadWithMax is Read with an explicit size cap (vision / upload reads).
func (f *FS) ReadWithMax(rel string, maxBytes int) ([]byte, error) { panic("fake") }

// ReadLines reads a line range (1-based inclusive). Empty end means EOF.
func (f *FS) ReadLines(rel string, start, end int) (string, error) { panic("fake") }

// Write replaces a file's contents, creating parent dirs as needed.
func (f *FS) Write(rel string, content []byte) error { panic("fake") }

// WriteWithMax is Write with an explicit size cap (KMA-102 chat image uploads
// may allow more than the general workspace max_file_bytes).
func (f *FS) WriteWithMax(rel string, content []byte, maxBytes int) error { panic("fake") }

func (f *FS) writeWithMax(rel string, content []byte, maxBytes int) error { panic("fake") }

// Mkdir creates a directory (and any missing parents) at rel.
func (f *FS) Mkdir(rel string) error { panic("fake") }

// Create makes a new empty file at rel, creating parent dirs as needed.
// It fails if rel already exists (use Write to overwrite an existing file).
func (f *FS) Create(rel string) error { panic("fake") }

// Rename moves a file or directory from oldRel to newRel, both resolved
// under the workspace root. Fails if newRel already exists.
func (f *FS) Rename(oldRel, newRel string) error { panic("fake") }

// Delete removes a file or directory (recursively) at rel.
func (f *FS) Delete(rel string) error { panic("fake") }

// Walk visits files under rel, skipping denied paths. fn receives rel path.
func (f *FS) Walk(rel string, fn func(rel string, info fs.DirEntry) error) error { panic("fake") }
