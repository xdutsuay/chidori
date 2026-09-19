package api

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/activity"
)

// Stress caps — guard rails so the tool can't DoS the host.
const (
	maxStressCount       = 1000
	maxStressConcurrency = 64
	maxStressTokens      = 100
)

// RunStress fires count random-token inference requests at a target node (or
// the default pool routing) at a bounded concurrency, and returns latency +
// throughput statistics. Requests go straight to the dispatcher/pool, bypassing
// cache/planner/compaction, so they measure raw inference power.
func (h *Handlers) RunStress(w http.ResponseWriter, r *http.Request) { panic("fake") }

// target node ID; empty = default pool routing
// provider hint when no node

// Record one summary event so the run shows on the debug timeline.

func computeStressStats(lat []float64, ok, fail int, totalMs float64) map[string]any { panic("fake") }

func percentile(sorted []float64, p int) float64 { panic("fake") }

const tokenCharset = "abcdefghijklmnopqrstuvwxyz0123456789"

func randomTokens(n int) string { panic("fake") }

func clampInt(v, def, lo, hi int) int { panic("fake") }

func defaultStr(s, def string) string { panic("fake") }

func asFloat(v any) float64 { panic("fake") }
