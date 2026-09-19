package workflow

import "time"

// PauseKind records why a run is paused (approval | sleep | event).
type PauseKind string

const (
	PauseApproval PauseKind = "approval"
	PauseSleep    PauseKind = "sleep"
	PauseEvent    PauseKind = "event"
)

// Workflow YAML lives under .lclreason/workflows/<id>.yaml.
type Workflow struct {
	ID           string         `json:"id" yaml:"id"`
	Name         string         `json:"name" yaml:"name"`
	Description  string         `json:"description" yaml:"description"`
	TriggerKind  string         `json:"trigger_kind" yaml:"trigger_kind"` // manual | on_file_save | on_commit | cron | chat_command
	TriggerConf  map[string]any `json:"trigger_conf" yaml:"trigger_conf"`
	Steps        []Step         `json:"steps" yaml:"steps"`
	TrustGranted bool           `json:"trust_granted"` // set at load from trust file, NEVER from YAML
}

// RetryConfig is optional YAML on leaf steps (P2). Absent = single attempt (legacy).
type RetryConfig struct {
	MaxAttempts int      `json:"max_attempts,omitempty" yaml:"max_attempts,omitempty"`
	Backoff     string   `json:"backoff,omitempty" yaml:"backoff,omitempty"`
	Jitter      bool     `json:"jitter,omitempty" yaml:"jitter,omitempty"`
	On          []string `json:"on,omitempty" yaml:"on,omitempty"` // nonzero_exit | error
}

type Step struct {
	ID             string       `json:"id" yaml:"id"`
	Name           string       `json:"name" yaml:"name"`
	Kind           string       `json:"kind" yaml:"kind"` // shell | llm_call | condition | approval_gate | loop | sleep | wait_event
	Command        string       `json:"command,omitempty" yaml:"command,omitempty"`
	Model          string       `json:"model,omitempty" yaml:"model,omitempty"`
	Template       string       `json:"template,omitempty" yaml:"template,omitempty"`
	GateMessage    string       `json:"gate_message,omitempty" yaml:"gate_message,omitempty"`
	ConditionExpr  string       `json:"condition_expr,omitempty" yaml:"condition_expr,omitempty"`
	OverExpr       string       `json:"over_expr,omitempty" yaml:"over_expr,omitempty"`
	ThenSteps      []Step       `json:"then_steps,omitempty" yaml:"then_steps,omitempty"`
	ElseSteps      []Step       `json:"else_steps,omitempty" yaml:"else_steps,omitempty"`
	LoopBodySteps  []Step       `json:"loop_body_steps,omitempty" yaml:"loop_body_steps,omitempty"`
	TimeoutSeconds int          `json:"timeout_seconds,omitempty" yaml:"timeout_seconds,omitempty"`
	SleepSeconds   int          `json:"sleep_seconds,omitempty" yaml:"sleep_seconds,omitempty"` // sleep step duration
	WaitEvent      string       `json:"wait_event,omitempty" yaml:"wait_event,omitempty"`       // wait_event step event name
	Retry          *RetryConfig `json:"retry,omitempty" yaml:"retry,omitempty"`
}

type StepOutput struct {
	Output     string `json:"output"`
	ExitCode   int    `json:"exit_code,omitempty"`
	DurationMs int64  `json:"duration_ms"`
	Attempt    int    `json:"attempt,omitempty"`
	LastError  string `json:"last_error,omitempty"`
}

type ExecutionState struct {
	RunID        string                `json:"run_id"`
	WorkflowID   string                `json:"workflow_id"`
	CurrentStep  int                   `json:"current_step"` // legacy mirror of Cursor.Top
	Cursor       StepCursor            `json:"cursor"`
	Outputs      map[string]StepOutput `json:"outputs"`
	Variables    map[string]any        `json:"variables"`
	Status       string                `json:"status"` // running|paused|completed|failed|needs_trust
	PauseKind    PauseKind             `json:"pause_kind,omitempty"`
	ResumeAt     *time.Time            `json:"resume_at,omitempty"`
	WaitingEvent string                `json:"waiting_event,omitempty"`
	ResumeToken  string                `json:"resume_token,omitempty"`
	Message      string                `json:"message,omitempty"`
	UpdatedAt    time.Time             `json:"updated_at"`
	Log          []string              `json:"log,omitempty"`
}
