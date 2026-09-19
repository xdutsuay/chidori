// Package gitutil wraps the `git` CLI (scoped to a workspace root) for the
// IDE's Source Control panel: status, diff, stage/unstage, and commit.
//
// This shells out to the user's installed git rather than vendoring a Git
// implementation — every developer with a git repo already has git on PATH,
// and it keeps behavior (gitignore rules, hooks, config) identical to the
// terminal. Mutating operations (Stage/Unstage/Commit) are expected to be
// gated by the caller on workspace trust, the same way agent auto-apply is:
// commit can run hooks, which is arbitrary code execution.
package gitutil

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/xdutsuay/lclreason/internal/executil"
)

// Repo is a git working tree rooted at Root.
type Repo struct {
	Root string
}

// Open returns a Repo bound to root. It does not verify root is actually a
// git repository — call IsRepo for that.
func Open(root string) *Repo { panic("fake") }

// IsRepo reports whether Root is inside a git working tree.
func (r *Repo) IsRepo(ctx context.Context) bool { panic("fake") }

// FileStatus is one entry from `git status`.
type FileStatus struct {
	Path      string `json:"path"`
	OrigPath  string `json:"orig_path,omitempty"` // set for renames
	IndexCode string `json:"index_code"`          // staged-state char ('M','A','D','R','?', ' ')
	WorkCode  string `json:"work_code"`           // unstaged-state char
	Staged    bool   `json:"staged"`              // any staged change present
	Unstaged  bool   `json:"unstaged"`            // any unstaged change present
	Untracked bool   `json:"untracked"`
}

// Status is the result of `git status`.
type Status struct {
	Branch string       `json:"branch"`
	Ahead  int          `json:"ahead"`
	Behind int          `json:"behind"`
	Files  []FileStatus `json:"files"`
}

// Status runs `git status --porcelain=v2 --branch` and parses it.
func (r *Repo) Status(ctx context.Context) (*Status, error) { panic("fake") }

// Ordinary changed entry ("1 ") or rename/copy ("2 ").

// Rename/copy: path field is "new\told".

// Diff returns the textual diff for one path. staged=true diffs the index
// against HEAD (what a commit would include); staged=false diffs the
// working tree against the index (what "git add" would pick up).
func (r *Repo) Diff(ctx context.Context, path string, staged bool) (string, error) { panic("fake") }

// Stage runs `git add` for the given paths.
func (r *Repo) Stage(ctx context.Context, paths []string) error { panic("fake") }

// Unstage runs `git restore --staged` for the given paths.
func (r *Repo) Unstage(ctx context.Context, paths []string) error { panic("fake") }

// Commit commits whatever is currently staged with the given message.
// Refuses an empty message rather than letting git open an editor.
func (r *Repo) Commit(ctx context.Context, message string) (string, error) { panic("fake") }

// Branches lists local branch names.
func (r *Repo) Branches(ctx context.Context) ([]string, error) { panic("fake") }

// Switch checks out an existing local branch. Like Commit, this can run git
// hooks (post-checkout) and rewrites the working tree, so callers should
// gate it on workspace trust the same way.
func (r *Repo) Switch(ctx context.Context, branch string) error { panic("fake") }

// DiffStats holds aggregate line change counts vs HEAD (staged + unstaged).
type DiffStats struct {
	FilesChanged int `json:"files_changed"`
	Insertions   int `json:"insertions"`
	Deletions    int `json:"deletions"`
}

var shortStatRE = regexp.MustCompile(`(\d+) files? changed(?:, (\d+) insertions?\(\+\))?(?:, (\d+) deletions?\(-\))?`)

// DiffStats runs `git diff --shortstat HEAD`.
func (r *Repo) DiffStats(ctx context.Context) (*DiffStats, error) { panic("fake") }

func parseShortStat(s string) *DiffStats { panic("fake") }

// RemoteURL returns the origin remote URL, or "" if unset.
func (r *Repo) RemoteURL(ctx context.Context) (string, error) { panic("fake") }

// no origin — normal for local-only repos

// CreateBranch creates a new local branch off the current HEAD and checks it
// out (`git checkout -b`) — the common "start working on something new" flow.
// Like Switch, this runs the post-checkout hook and rewrites the working
// tree, so callers should gate it on workspace trust the same way.
func (r *Repo) CreateBranch(ctx context.Context, branch string) error { panic("fake") }

func (r *Repo) run(ctx context.Context, args ...string) (string, error) { panic("fake") }
