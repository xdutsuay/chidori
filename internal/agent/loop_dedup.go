package agent

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// toolDedupCache tracks repeated tool calls within one agent run so weak models
// cannot burn the whole turn budget re-reading the same path or re-proposing
// writes to the same file with slightly different content each turn.
// Concurrent plan-tool batches share one cache; mu guards the maps.
type toolDedupCache struct {
	mu                sync.Mutex
	entries           map[string]toolDedupEntry
	writes            map[string]toolDedupEntry // canonical path → first successful write_file/apply_patch
	delegateSuccesses int                       // KMA-124: successful delegate_subtask calls this run
}

type toolDedupEntry struct {
	count     int
	output    string
	isErr     bool
	firstTurn int
}

func newToolDedupCache() *toolDedupCache { panic("fake") }

// dedupExemptTools are legitimately repeatable with identical params —
// subagent_result is a poll whose answer changes as the background job
// progresses, so caching/skipping repeats would freeze it at "still running".
var dedupExemptTools = map[string]bool{"subagent_result": true}

func writeMutatingTool(name string) bool { panic("fake") }

// writeToolPathKey returns the canonical workspace path for write tools, or ""
// when the call has no path (caller should not path-throttle those).
func writeToolPathKey(name string, params map[string]any) string { panic("fake") }

func canonicalToolParams(name string, params map[string]any) map[string]any { panic("fake") }

func toolCallFingerprint(name string, params map[string]any) string { panic("fake") }

func idempotentReadTool(name string) bool { panic("fake") }

// before checks whether an identical successful tool call already ran, or a
// second write to the same path is being attempted. When skip is true the
// caller must not execute the tool again; forceDone ends the run after the
// third repeat of a read-only tool, or immediately on a repeat write path.
func (c *toolDedupCache) before(name string, params map[string]any) (skip, forceDone bool, cachedOut string, firstTurn, repeatCount int) {
	panic("fake")
}

// KMA-124: after one successful delegate_subtask, any further call ends the run
// (weak models otherwise re-spawn with rephrased queries forever).

// after records the first successful (or failed) execution of a tool call.
func (c *toolDedupCache) after(name string, params map[string]any, output string, isErr bool, turn int) {
	panic("fake")
}

func duplicateToolFeedback(name string, repeatCount, firstTurn int) string { panic("fake") }

func forceDoneAfterDuplicateRead(output string) string { panic("fake") }

func duplicateWritePathFeedback(path string, repeatCount, firstTurn int, proposeOnly bool) string {
	panic("fake")
}

func forceDoneAfterDuplicateWrite(path, output string, proposeOnly bool) string { panic("fake") }
