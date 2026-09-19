package workspace

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

// SearchMatch is one workspace search hit.
type SearchMatch struct {
	Path    string `json:"path"`
	Line    int    `json:"line"`
	Content string `json:"content"`
}

// Search scans file contents under the workspace for a regex pattern.
// include/exclude are optional comma-separated glob lists (e.g.
// "*.go,*.ts"); a file must match at least one include glob (if any are
// given) and none of the exclude globs. caseSensitive=false compiles the
// pattern with Go regexp's inline (?i) case-insensitive flag rather than
// doing any manual case-folding, so existing regex-metacharacter behavior
// (anchors, groups, etc.) is unaffected either way.
func (f *FS) Search(pattern, include, exclude string, caseSensitive bool, limit int) ([]SearchMatch, error) {
	panic("fake")
}

// ReplaceResult is one file touched by a replace-in-files operation.
type ReplaceResult struct {
	Path  string `json:"path"`
	Count int    `json:"count"`
}

// ReplaceInFiles finds every file matching pattern (optionally filtered by
// glob) and, when apply is true, rewrites each with all matches substituted
// for replacement (regexp.ReplaceAllString semantics — $1 etc. work as
// capture-group references). When apply is false this is a dry run: it
// reports which files would be touched and how many matches each has,
// without writing anything, so the caller can show a preview before
// committing (mirrors the agent's propose/accept pattern, but for a
// user-initiated bulk edit rather than an AI-proposed one).
func (f *FS) ReplaceInFiles(pattern, replacement, include, exclude string, caseSensitive bool, limit int, apply bool) ([]ReplaceResult, error) {
	panic("fake")
}

// ListFiles returns all file paths under the workspace (for @file picker).
func (f *FS) ListFiles(limit int) ([]string, error) { panic("fake") }

func matchGlobSearch(glob, path string) bool { panic("fake") }

// matchAnyGlob reports whether path matches any glob in a comma-separated
// list (e.g. "*.go, *.ts"); empty/blank entries are skipped.
func matchAnyGlob(globs, path string) bool { panic("fake") }
