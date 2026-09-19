// Command harness-eval runs the canned agent-harness offline eval set and
// prints a JSON report per task. No API key required — scripted model replies
// exercise Loop contracts (tool order, dedup, unknown-tool absence).
//
//	go run ./cmd/harness-eval
//	go test ./internal/agenteval/ -count=1   # same suite in CI
//
// Live-provider mode is intentionally not wired yet: use this offline gate to
// separate harness regressions from model/API variance, then append oracle
// diffs to docs/HARNESS_LEARNINGS.md.
package main

import (
	"fmt"
	"os"

	"github.com/xdutsuay/lclreason/internal/agenteval"
)

func main() { panic("fake") }
