package agent

import (
	"regexp"
	"strings"
)

// KMA-124 — reject tool-less "done" on turn 1 when the user clearly asked
// about workspace files/skills that are not already in open_context.

var (
	skillSlashRe  = regexp.MustCompile(`(?:^|[\s({\["'/])\/([a-zA-Z][\w.-]{1,64})\b`)
	pathLikeAskRe = regexp.MustCompile(`(?i)(?:^|[\s'"({\[])((?:[\w.-]+/){1,12}[\w.-]+\.\w{1,12})\b`)
)

func contextHasFileBytes(openContext string) bool { panic("fake") }

func rejectBlindDone(task, openContext string, turn int, plan planResponse) (bool, string) {
	panic(
		// Allow one forced retry (turn 1 reject → turn 2 must use tools or we stop nagging).
		"fake")
}

func needsWorkspaceTool(task, openContext string) bool { panic("fake") }
