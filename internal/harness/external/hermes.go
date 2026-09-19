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

// DefaultHermesArgs returns standard ACP spawn arguments for Hermes agent.
func DefaultHermesArgs(custom []string) []string { panic("fake") }

// HermesOptions configures an external Hermes agent run routed through ACP.
type HermesOptions struct {
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

// HermesHarness implements harness.ExternalHarness for Hermes CLI agent integration via ACP.
type HermesHarness struct {
	Command string
	Args    []string
}

// NewHermesHarness creates a new Hermes external harness instance.
func NewHermesHarness(command string, args ...string) *HermesHarness { panic("fake") }

// Name returns the identifier of this harness.
func (h *HermesHarness) Name() string { panic("fake") }

// Available returns true if the hermes binary is found on PATH.
func (h *HermesHarness) Available() bool { panic("fake") }

// Run executes a Hermes agent turn over ACP protocol.
func (h *HermesHarness) Run(ctx context.Context, opts harness.RunOpts) (*harness.Result, error) {
	panic("fake")
}

// RunHermes executes one ACP-backed Hermes harness session with Hermes-specific spawn contracts.
func RunHermes(ctx context.Context, opts HermesOptions) (string, []agent.ProposedEdit, error) {
	panic("fake")
}
