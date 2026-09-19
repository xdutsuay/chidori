package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/xdutsuay/lclreason/internal/tools"
)

// GitHubStatus reports whether `gh` is installed and authenticated (KMA-120).
func (h *Handlers) GitHubStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitHubAuthLogout runs `gh auth logout` for the given hostname (KMA-136).
// gh may prompt on stdin; we answer Y when a TTY is not available.
func (h *Handlers) GitHubAuthLogout(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitHubPRs lists PRs for the workspace repo via `gh pr list` (KMA-120).
func (h *Handlers) GitHubPRs(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitHubIssues lists issues for the workspace repo via `gh issue list` (KMA-120).
func (h *Handlers) GitHubIssues(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitHubPRCreate runs `gh pr create`. Gated on workspace trust — creates a
// remote PR against the authenticated account (same risk class as git_commit).
func (h *Handlers) GitHubPRCreate(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitHubIssueCreate runs `gh issue create`. Gated on workspace trust.
func (h *Handlers) GitHubIssueCreate(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GitHubIssueComment runs `gh issue comment`. Gated on workspace trust.
func (h *Handlers) GitHubIssueComment(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) githubWorkspaceRoot(w http.ResponseWriter) (string, bool) { panic("fake") }

// githubTrustedRoot requires an open workspace that is trusted (Settings),
// matching POST /api/git/commit — mutating GitHub actions touch the remote
// and checkout can run git hooks.
func (h *Handlers) githubTrustedRoot(w http.ResponseWriter) (string, bool) { panic("fake") }

func githubListLimit(r *http.Request, def int) int { panic("fake") }

func parseGitHubLabels(raw json.RawMessage) []string { panic("fake") }

func parseGitHubNumber(raw json.RawMessage) string { panic("fake") }
