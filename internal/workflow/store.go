package workflow

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func Dir(root string) string { panic("fake") }

func TrustPath(root string) string { panic("fake") }

func RunsDir(root string) string { panic("fake") }

func List(root string) ([]Workflow, error) { panic("fake") }

func Get(root, id string) (*Workflow, error) { panic("fake") }

// never from YAML

// ReadRaw returns the raw YAML content of a workflow file (WF.5 — YAML editor).
func ReadRaw(root, id string) ([]byte, error) { panic("fake") }

// SaveRaw writes raw YAML content to a workflow file (WF.5 — YAML editor).
// It validates the YAML parses as a Workflow before writing.
func SaveRaw(root, id string, data []byte) error { panic("fake") }

func Save(root string, wf *Workflow) error { panic("fake") }

// Marshal without trust_granted — it lives only in workflow_trust.json.

func SetTrusted(root, id string, trusted bool) error { panic("fake") }

const maxRunRetention = 100

func SaveRun(root string, state *ExecutionState) error { panic("fake") }

func LoadRun(root, runID string) (*ExecutionState, error) { panic("fake") }

func normalizeRunState(st *ExecutionState) { panic("fake") }

func ListRuns(root string, max int) ([]ExecutionState, error) { panic("fake") }

func loadTrust(root string) (map[string]bool, error) { panic("fake") }

func loadTrustList(root string) ([]string, error) { panic("fake") }

func workflowPath(root, id string) string { panic("fake") }

func pruneRuns(root string, keep int) error { panic("fake") }
