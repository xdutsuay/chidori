package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/gitutil"
)

// gitTimeout bounds every git subprocess call — plenty for status/diff/add/
// commit on a local repo, short enough that a wedged git process (e.g. an
// interactive credential prompt) doesn't hang the request indefinitely.
const gitTimeout = 15 * time.Second

func timeoutCtx(r *http.Request) (context.Context, context.CancelFunc) { panic("fake") }

// gitRepo returns a gitutil.Repo rooted at the current workspace, or nil if
// no workspace is open.
func (h *Handlers) gitRepo() *gitutil.Repo { panic("fake") }

// gitValidatePath resolves a workspace-relative path through the sandboxed
// *workspace.FS (rejecting traversal outside the workspace root) and returns
// it in the workspace-relative form git expects for its "-- path" args.
func (h *Handlers) gitValidatePath(path string) (string, error) { panic("fake") }

// GitStatus returns branch/ahead-behind + changed files (staged/unstaged/
// untracked). 200 with in_repo:false if the workspace isn't a git repo at
// all, rather than an error — that's a normal, expected state.
func (h *Handlers) GitStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitDiffStats returns aggregate +/- line counts vs HEAD for the All-Agent working bar.
func (h *Handlers) GitDiffStats(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitDiff returns the diff for ?path= (?staged=1 for the staged/cached diff,
// otherwise the working-tree-vs-index diff).
func (h *Handlers) GitDiff(w http.ResponseWriter, r *http.Request) { panic("fake") }

type gitPathsRequest struct {
	Paths []string `json:"paths"`
}

// GitStage runs `git add` for the given paths. Staging doesn't run git
// hooks, so — unlike Commit — this isn't gated on workspace trust.
func (h *Handlers) GitStage(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitUnstage runs `git restore --staged` for the given paths.
func (h *Handlers) GitUnstage(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitCommit commits whatever is currently staged. Gated on workspace trust:
// commit can run pre-commit/commit-msg/post-commit hooks, which is arbitrary
// code execution — the same reason agent auto-apply is gated (see
// code_handlers.go's AgentRun). An untrusted folder can still be inspected
// (status/diff) and staged, just not committed.
func (h *Handlers) GitCommit(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ADR-0013
// ADR-0015 Phase 2

// GitBranches lists local branches.
func (h *Handlers) GitBranches(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitSwitch checks out an existing local branch. Gated on workspace trust —
// checkout can run post-checkout hooks, the same arbitrary-code-execution
// concern as GitCommit above.
func (h *Handlers) GitSwitch(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitCreateBranch creates a new local branch off HEAD and checks it out.
// Gated on workspace trust — like GitSwitch, `checkout -b` runs the
// post-checkout hook and rewrites the working tree.
func (h *Handlers) GitCreateBranch(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) decodeAndValidatePaths(r *http.Request) ([]string, error) { panic("fake") }
