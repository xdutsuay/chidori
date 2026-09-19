package external

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/acpclient"
	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/harness"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// GrokHarness implements harness.ExternalHarness for Grok CLI agent integration via ACP.
type GrokHarness struct {
	Command string
	Args    []string
}

// NewGrokHarness creates a new Grok external harness instance.
func NewGrokHarness(command string, args ...string) *GrokHarness { panic("fake") }

// Name returns the identifier of this harness.
func (g *GrokHarness) Name() string { panic("fake") }

// Available returns true if the grok binary is found on PATH.
func (g *GrokHarness) Available() bool { panic("fake") }

// Run executes a Grok agent turn over ACP protocol.
func (g *GrokHarness) Run(ctx context.Context, opts harness.RunOpts) (*harness.Result, error) {
	panic("fake")
}

// GrokOptions configures one external-harness run routed through ACP.
type GrokOptions struct {
	Command     string
	Args        []string
	Task        string
	OpenContext string

	WorkspaceRoot string
	Workspace     *workspace.FS
	Trusted       bool
	ProposeOnly   bool

	Emit func(agent.TurnEvent)
}

// RunGrok executes one ACP-backed external harness session and emits native
// TurnEvents so the rest of the UI/API can stay transport-agnostic.
func RunGrok(ctx context.Context, opts GrokOptions) (string, []agent.ProposedEdit, error) {
	panic("fake")
}

// AG.16 pattern: bound cancel RPC — an unresponsive child is exactly
// why Prompt failed; don't block cleanup on a fresh never-expiring ctx.
