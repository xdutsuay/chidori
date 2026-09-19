package workflow

import (
	"fmt"
	"strings"
	"time"
)

const minCronInterval = time.Minute

// MatchCommitTrigger reports whether wf should run after a git commit.
func MatchCommitTrigger(wf Workflow) bool { panic("fake") }

// MatchChatCommandTrigger reports whether wf is bound to a chat slash command.
func MatchChatCommandTrigger(wf Workflow, command string) bool { panic("fake") }

// CronInterval parses trigger_conf for cron workflows.
// Supports interval_seconds (int/float) or interval (Go duration string, e.g. "5m").
func CronInterval(wf Workflow) (time.Duration, error) { panic("fake") }

func clampCron(d time.Duration) time.Duration { panic("fake") }
