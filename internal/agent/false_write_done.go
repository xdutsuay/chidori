package agent

import (
	"regexp"
	"strings"
)

// falseWriteRunState summarizes mutating-tool outcomes for the current agent
// run when plan.Done is being validated (KMA-201).
type falseWriteRunState struct {
	HasSuccessfulWrite bool
	PendingEdits       []ProposedEdit
	ProposeOnly        bool
}

var (
	fileWriteVerbRe = regexp.MustCompile(`(?i)\b(write|append|create|save)\b`)
	fileWritePathRe = regexp.MustCompile(`(?i)(?:^|[\s'"` + "`" + `"({\[])((?:[\w.-]+/){0,12}[\w.-]+\.\w{1,12})\b`)
)

// taskHasFileWriteIntent reports whether the user task asks to create/write/
// append/save a file (heuristic; KMA-201).
func taskHasFileWriteIntent(task string) bool { panic("fake") }

// "write" alone is not enough — need a path-like target or "with <content>".

// e.g. "write 1.txt with hi" — path may be bare basename without slash.

// rejectFalseWriteDone blocks done:true when the task required a file write
// but this run never successfully wrote/proposed one.
func rejectFalseWriteDone(task string, plan planResponse, run falseWriteRunState) (bool, string) {
	panic("fake")
}
