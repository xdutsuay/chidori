package tools

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/executil"
)

// RegisterGitHubTools adds GitHub tools backed by the `gh` CLI (KMA-120).
// Auth is whatever `gh auth` already has — Chidori does not store a second PAT.
// rootFn supplies the workspace cwd so `gh` resolves the repo from git remotes
// the same way a developer shell would.
//
// Mutating tools (create PR/issue, comment, checkout) are registered here but
// gated on workspace trust in the agent guard / advertisableTool policy (same
// class as git_commit). Missing `gh` yields a clear tool error, not a panic.
func RegisterGitHubTools(r *Registry, rootFn func() string) { panic("fake") }

func githubAuthStatusTool() ToolFunc { panic("fake") }

func githubRepoViewTool(rootFn func() string) ToolFunc { panic("fake") }

func githubPRListTool(rootFn func() string) ToolFunc { panic("fake") }

func githubPRViewTool(rootFn func() string) ToolFunc { panic("fake") }

func githubIssueListTool(rootFn func() string) ToolFunc { panic("fake") }

func githubCIStatusTool(rootFn func() string) ToolFunc { panic("fake") }

func githubPRCreateTool(rootFn func() string) ToolFunc { panic("fake") }

func githubIssueCreateTool(rootFn func() string) ToolFunc { panic("fake") }

func githubIssueCommentTool(rootFn func() string) ToolFunc { panic("fake") }

func githubPRCheckoutTool(rootFn func() string) ToolFunc { panic("fake") }

func runGHTool(ctx context.Context, rootFn func() string, args []string) (*ToolResult, error) {
	panic("fake")
}

func runGH(ctx context.Context, dir string, args []string) (string, error) { panic("fake") }

// RunGHInteractive runs gh with optional stdin (for auth logout prompts).
func RunGHInteractive(ctx context.Context, dir, stdin string, args ...string) (string, error) {
	panic("fake")
}

// RunGH executes `gh` with args in dir (empty dir = process cwd). Used by
// HTTP handlers and tests; agent tools go through RegisterGitHubTools.
func RunGH(ctx context.Context, dir string, args ...string) (string, error) { panic("fake") }

func stringParam(params map[string]any, key, def string) string { panic("fake") }

// numberParam accepts string or JSON number for issue/PR identifiers.
func numberParam(params map[string]any, key string) string { panic("fake") }

// labelsParam accepts a comma-separated string or an array of strings.
func labelsParam(params map[string]any, key string) []string { panic("fake") }

func boolParam(params map[string]any, key string, def bool) bool { panic("fake") }

func intParam(params map[string]any, key string, def int) int { panic("fake") }
