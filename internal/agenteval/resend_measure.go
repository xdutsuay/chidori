package agenteval

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strconv"

	"github.com/xdutsuay/lclreason/internal/agent"
)

// ResendTurn is one inference turn's prompt-split measurements.
type ResendTurn struct {
	Turn           int    `json:"turn"`
	StableBytes    int    `json:"stable_bytes"`
	VolatileBytes  int    `json:"volatile_bytes"`
	StableReuse    int    `json:"stable_reuse"`
	StableHash     string `json:"stable_hash,omitempty"`
	CacheHitTokens int    `json:"cache_hit_tokens,omitempty"`
}

// ResendReport is the KMA-163 slice-1 baseline.
type ResendReport struct {
	TaskID          string       `json:"task_id"`
	Turns           []ResendTurn `json:"turns"`
	CumulativeBytes int          `json:"cumulative_bytes"`
	ResendBytes     int          `json:"resend_bytes"`
	CacheHitTokens  int          `json:"cache_hit_tokens,omitempty"`

	stableSnapshots []string
}

// StableSnapshot returns the captured stable prompt prefix for turn i.
func (r ResendReport) StableSnapshot(i int) string { panic("fake") }

var resendOutputRe = regexp.MustCompile(`\bstable_reuse=(\d+)\b`)

// cacheReadUsageSeam is the optional invoker seam for provider cache-read
// token accounting. When a caller-supplied invoker implements it,
// RunOfflineMeasuredWithInvoker surfaces its cumulative value on the report.
type cacheReadUsageSeam interface {
	CacheReadUsage() int
}

// RunOfflineMeasured runs Task through the real agent Loop with scripted model
// replies (like RunOffline) and additionally captures every "inference"
// TurnEvent to build a ResendReport.
func RunOfflineMeasured(task Task) (ResendReport, []Check) { panic("fake") }

// RunOfflineMeasuredWithInvoker runs Task through the real agent Loop with a
// caller-supplied invoker and builds the same ResendReport as
// RunOfflineMeasured. When inv implements cacheReadUsageSeam, the cumulative
// CacheReadUsage() value is copied onto report.CacheHitTokens.
func RunOfflineMeasuredWithInvoker(task Task, inv agent.Invoker) (ResendReport, []Check) {
	panic("fake")
}
