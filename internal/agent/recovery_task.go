package agent

import (
	"fmt"
	"strings"
)

// ExpandRecoveryTask rewrites short recovery cues (retry/again/continue) into an
// explicit redo of the last concrete prior user task so the agent does not
// thrash on search_chat_messages / list_dir (KMA-195).
func ExpandRecoveryTask(task string, priorUserTasks []string) string { panic("fake") }

// IsRecoveryCue reports whether task is a short resume/retry cue (exact match
// after trim; case-insensitive).
func IsRecoveryCue(task string) bool { panic("fake") }

func lastConcreteUserTask(prior []string) string { panic("fake") }
