package workspace

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// LangStat is one language row in a workspace LOC breakdown.
type LangStat struct {
	Language string `json:"language"`
	Files    int    `json:"files"`
	Lines    int    `json:"lines"`
	NonBlank int    `json:"non_blank"`
}

// StatsReport is the payload for GET /api/workspace/stats.
type StatsReport struct {
	Root      string     `json:"root,omitempty"`
	Languages []LangStat `json:"languages"`
	Total     LangStat   `json:"total"`
	Skipped   int        `json:"skipped"` // binary / oversize / unreadable files
}

// Extra directory names skipped by LanguageStats even when not in DenyPaths.
// These inflate "codebase shape" counts without being part of the product
// source the user usually wants to see (Go vendor trees, build outputs, etc.).
var statsSkipDirs = map[string]bool{
	"vendor": true, "node_modules": true, "dist": true, "build": true,
	"bin": true, "coverage": true, "__pycache__": true, ".venv": true,
	"venv": true, "target": true, "graphify-out": true, ".git": true,
	".idea": true, ".vscode": true, ".cursor": true,
}

var extLanguage = map[string]string{
	".go": "Go", ".js": "JavaScript", ".mjs": "JavaScript", ".cjs": "JavaScript",
	".ts": "TypeScript", ".tsx": "TypeScript", ".jsx": "JavaScript",
	".html": "HTML", ".htm": "HTML", ".css": "CSS", ".scss": "CSS",
	".md": "Markdown", ".markdown": "Markdown",
	".py": "Python", ".yaml": "YAML", ".yml": "YAML", ".json": "JSON",
	".sh": "Shell", ".bash": "Shell", ".zsh": "Shell",
	".rs": "Rust", ".java": "Java", ".c": "C", ".h": "C", ".cpp": "C++",
	".cc": "C++", ".hpp": "C++", ".swift": "Swift", ".kt": "Kotlin",
	".toml": "TOML", ".sql": "SQL", ".proto": "Protobuf",
	".txt": "Text", ".xml": "XML",
}

const statsMaxFileBytes = 2 * 1024 * 1024 // hard cap so a huge dump can't stall the UI

// LanguageStats walks the workspace and counts files/lines/non-blank lines
// per language (by extension). Respects DenyPaths via Walk, plus statsSkipDirs
// and a few binary/minified heuristics. On-demand only — not cached.
func (f *FS) LanguageStats() (StatsReport, error) { panic("fake") }

func shouldSkipStatsFile(base string) bool { panic("fake") }

func languageForPath(rel string) string { panic("fake") }

func countFileLines(f *FS, rel string) (lines, nonBlank int, ok bool) { panic("fake") }

// Peek for NUL — treat as binary and skip.

// Allow long lines (generated JSON, etc.) up to 1 MiB.
