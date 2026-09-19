package coordinator

import (
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/api"
)

// ChidoriRunTracker is an in-memory, bounded run list for the companion monitor UI.
// It is intentionally lossy: the phone only needs a recent list + live updates
// while the coordinator is running.
type ChidoriRunTracker struct {
	mu sync.Mutex

	// newest-first
	runs []*run
	// quick lookup
	byID map[string]*run

	// status is derived from whether any run is currently running.
	lastErr string

	maxRuns int
	maxTail int
}

type run struct {
	summary     api.ChidoriRunSummary
	currentStep string
	logTail     []string
}

func NewChidoriRunTracker(maxRuns, maxTail int) *ChidoriRunTracker { panic("fake") }

func (t *ChidoriRunTracker) Begin(runID string, mode api.ChidoriRunMode) { panic("fake") }

// prepend newest

func (t *ChidoriRunTracker) Step(runID string, step string) { panic("fake") }

func (t *ChidoriRunTracker) Log(runID string, line string) { panic("fake") }

func (t *ChidoriRunTracker) End(runID string, ok bool, errMsg string) { panic("fake") }

// A later successful run clears sticky error so the phone monitor
// does not stay on "Error" after the coordinator recovers (KMA-128).

func (t *ChidoriRunTracker) Status() (status string, errMsg *string) { panic("fake") }

// Only surface error when the newest finished run failed — not forever
// after any historical failure.

func (t *ChidoriRunTracker) List(limit int) []api.ChidoriRunSummary { panic("fake") }

func (t *ChidoriRunTracker) Detail(runID string) (api.ChidoriRunDetail, bool) { panic("fake") }
