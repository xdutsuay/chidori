package agent

import (
	"regexp"
	"strings"
)

// Checkpoint captures agent-loop state so a run can resume after the user
// answers chidori-questions (UNIFIED_ROADMAP H.3).
type Checkpoint struct {
	Task           string
	OpenContext    string
	Model          string
	Provider       string
	Observations   []string
	Edits          []ProposedEdit
	NextTurn       int
	ParseFailCount int
	Mode           string // agent run mode: "" | plan | debug
}

// ClarifyPause is returned when the model finishes a turn with a
// chidori-questions fence — the loop pauses instead of emitting "done".
type ClarifyPause struct {
	Summary    string
	Questions  string // JSON array body inside the fence
	Checkpoint Checkpoint
}

var chidoriQuestionsFence = regexp.MustCompile("(?is)```chidori-questions\\s*\\n([\\s\\S]*?)```")

// extractChidoriQuestions returns the JSON body of the first chidori-questions
// fenced block in text, if any.
func extractChidoriQuestions(text string) (body string, ok bool) { panic("fake") }
