package workflow

import (
	"fmt"
	"path/filepath"
	"strings"
)

var allowedStepKinds = map[string]struct{}{
	"shell":         {},
	"llm_call":      {},
	"condition":     {},
	"approval_gate": {},
	"loop":          {},
	"sleep":         {},
	"wait_event":    {},
}

// Validate rejects malformed workflow definitions before run or persist.
func Validate(wf *Workflow) error { panic("fake") }

func validateStep(s Step, path string) error { panic("fake") }

func validateRetry(cfg *RetryConfig, path string, shell bool) error { panic("fake") }
