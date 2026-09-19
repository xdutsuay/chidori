package workflow

import (
	"fmt"
	"regexp"
	"strings"
)

// evalCondition parses expressions like:
//   - {{step1.output}} == 0
//   - {{step2.output}} contains "panic"
//   - {{step3.output}} matches /Error: .*/
func evalCondition(expr string, outputs map[string]StepOutput) (bool, error) { panic("fake") }

func resolveTemplateVar(placeholder string, outputs map[string]StepOutput) string { panic("fake") }
