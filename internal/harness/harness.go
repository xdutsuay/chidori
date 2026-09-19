package harness

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// RunOpts encapsulates parameters passed to an external agent harness execution turn.
type RunOpts struct {
	Task          string
	OpenContext   string
	WorkspaceRoot string
	Workspace     *workspace.FS
	Trusted       bool
	ProposeOnly   bool
	Emit          func(agent.TurnEvent)
	Timeout       time.Duration
	Command       string
	Args          []string
}

// HarnessStats contains execution telemetry for an agent turn.
type HarnessStats struct {
	Duration  time.Duration `json:"duration"`
	ToolCalls int           `json:"tool_calls"`
}

// Result holds the output of an external harness turn execution.
type Result struct {
	Text  string               `json:"text"`
	Edits []agent.ProposedEdit `json:"edits"`
	Error error                `json:"error,omitempty"`
	Stats HarnessStats         `json:"stats"`
}

// ExternalHarness defines the common interface for third-party CLI agents (Grok, Hermes, etc.).
type ExternalHarness interface {
	Name() string
	Available() bool
	Run(ctx context.Context, opts RunOpts) (*Result, error)
}

// Registry manages registered external agent harnesses.
type Registry struct {
	mu        sync.RWMutex
	harnesses map[string]ExternalHarness
}

// NewRegistry creates a new harness registry.
func NewRegistry() *Registry { panic("fake") }

// Register registers an external harness.
func (r *Registry) Register(h ExternalHarness) { panic("fake") }

// Get retrieves a harness by name.
func (r *Registry) Get(name string) (ExternalHarness, bool) { panic("fake") }

// Available returns names of all registered harnesses that are available on the host (e.g. binary on PATH).
func (r *Registry) Available() []string { panic("fake") }

// Run executes a named harness with safety timeout.
func (r *Registry) Run(ctx context.Context, name string, opts RunOpts) (*Result, error) {
	panic("fake")
}
