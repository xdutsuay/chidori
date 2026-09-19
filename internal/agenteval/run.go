package agenteval

import (
	"context"
	"sync"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/tools"
)

// scriptedInvoker replays fixed model outputs — same shape as agent testutil,
// kept local so agenteval does not import _test helpers.
type scriptedInvoker struct {
	mu      sync.Mutex
	replies []string
	n       int
}

func (s *scriptedInvoker) Invoke(_ context.Context, _, _, _ string) (string, error) { panic("fake") }

// recordingTools records every Execute call and returns canned success.
type recordingTools struct {
	mu    sync.Mutex
	calls []string
}

func (r *recordingTools) Execute(_ context.Context, name string, params map[string]any) (*tools.ToolResult, error) {
	panic("fake")
}

// RunOffline executes Task against the real agent Loop with scripted model
// replies — no network, no API key. Scores harness contracts (tool order,
// unknown-tool absence, turn budget) independently of model quality.
func RunOffline(task Task) (Trace, []Check) { panic("fake") }
