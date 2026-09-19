package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/workflow"
)

type workflowSummary struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TriggerKind  string `json:"trigger_kind"`
	TrustGranted bool   `json:"trust_granted"`
	StepCount    int    `json:"step_count"`
}

type workflowCommand struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
}

type workflowNotify struct {
	At         time.Time `json:"at"`
	WorkflowID string    `json:"workflow_id"`
	Trigger    string    `json:"trigger"`
}

func (h *Handlers) workflowRuntime() *workflow.Runtime { panic("fake") }

func (h *Handlers) workflowEngine() *workflow.Engine { panic("fake") }

// workflowInvoker adapts dispatch.Pool for llm_call steps. Returns nil when
// no pool is configured — shell, condition, and approval_gate steps still work.
func (h *Handlers) workflowInvoker() workflow.Invoker { panic("fake") }

type poolWorkflowInvoker struct {
	pi *agent.PoolInvoker
}

func (a poolWorkflowInvoker) Invoke(ctx context.Context, model, prompt string) (string, error) {
	panic("fake")
}

func (h *Handlers) ListWorkflows(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) GetWorkflow(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GetWorkflowRaw returns the raw YAML of a workflow file (WF.5 — YAML editor).
func (h *Handlers) GetWorkflowRaw(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SaveWorkflowRaw writes raw YAML content back to a workflow file (WF.5).
func (h *Handlers) SaveWorkflowRaw(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) SetWorkflowTrust(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) RunWorkflow(w http.ResponseWriter, r *http.Request) { panic("fake") }

// CreateSampleWorkflow installs docs/examples/workflows/hello-manual.yaml into the workspace.
func (h *Handlers) CreateSampleWorkflow(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) noteWorkflowTrigger(wfID, trigger string) { panic("fake") }

// ListWorkflowNotifications returns trigger events since ?since=unix (UTC seconds).
func (h *Handlers) ListWorkflowNotifications(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ApproveWorkflowRun(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) SignalWorkflowRun(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ListWorkflowRuns(w http.ResponseWriter, r *http.Request) { panic("fake") }

// fireWorkflowsOnSave runs trusted on_file_save workflows matching path (async).
func (h *Handlers) fireWorkflowsOnSave(path string) { panic("fake") }

// fireWorkflowsOnCommit runs trusted on_commit workflows (async).
func (h *Handlers) fireWorkflowsOnCommit() { panic("fake") }

// ListWorkflowCommands returns trusted chat_command triggers for slash routing.
func (h *Handlers) ListWorkflowCommands(w http.ResponseWriter, r *http.Request) { panic("fake") }

// FireWorkflowCommand runs trusted workflows bound to a chat slash command.
func (h *Handlers) FireWorkflowCommand(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) runWorkflowAsync(wfID, trigger string, vars map[string]any) { panic("fake") }

var _ workflow.CronRunner = (*Handlers)(nil)

// RunCronWorkflow implements workflow.CronRunner.
func (h *Handlers) RunCronWorkflow(ctx context.Context, wfID string) { panic("fake") }

func (h *Handlers) restartWorkflowCron() { panic("fake") }

func (h *Handlers) stopWorkflowCron() { panic("fake") }
