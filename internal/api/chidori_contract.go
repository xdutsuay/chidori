package api

// Wire-contract types for the chidori-nagasa LAN companion API.
//
// Source of truth: chidori-nagasa/WIRE_CONTRACT.md (protocol_version 1.1.0).

type ChidoriRunMode string

const (
	ChidoriRunAsk   ChidoriRunMode = "ask"
	ChidoriRunAgent ChidoriRunMode = "agent"
	ChidoriRunPlan  ChidoriRunMode = "plan"
	ChidoriRunDebug ChidoriRunMode = "debug"
)

type ChidoriRunState string

const (
	ChidoriRunRunning   ChidoriRunState = "running"
	ChidoriRunCompleted ChidoriRunState = "completed"
	ChidoriRunFailed    ChidoriRunState = "failed"
)

type ChidoriRunSummary struct {
	RunID       string          `json:"run_id"`
	Mode        ChidoriRunMode  `json:"mode"`
	StartedAt   int64           `json:"started_at"` // epoch millis
	State       ChidoriRunState `json:"state"`
	CurrentStep string          `json:"current_step,omitempty"`
}

type ChidoriRunDetail struct {
	Summary     ChidoriRunSummary `json:"summary"`
	CurrentStep *string           `json:"current_step,omitempty"`
	LogTail     []string          `json:"log_tail"`
}

// RunTracker is the minimal surface the LAN API needs for /coordinator/status and /runs.
type RunTracker interface {
	Begin(runID string, mode ChidoriRunMode)
	Step(runID string, step string)
	Log(runID string, line string)
	End(runID string, ok bool, errMsg string)

	Status() (status string, errMsg *string)
	List(limit int) []ChidoriRunSummary
	Detail(runID string) (ChidoriRunDetail, bool)
}
