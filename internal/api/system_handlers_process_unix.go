//go:build !windows

package api

import (
	"os"
	"strconv"
	"strings"

	"github.com/xdutsuay/lclreason/internal/executil"
)

// processCPUMem shells out to ps for this PID's CPU% (of one core), memory% (of
// physical RAM), and RSS in KB.
func processCPUMem() (cpuPct, memPct float64, rssKB int64) { panic("fake") }

// processCPUTimeSecs shells out to ps for this PID's cumulative CPU time
// (user+system, since process start — not a rate) for the Planner panel's
// "CPU hours" readout. ps -o time= reports it as [[dd-]hh:]mm:ss; parsed into
// seconds.
func processCPUTimeSecs() float64 { panic("fake") }

// processCPUReady reports whether a CPU% sample is meaningful (ps always is).
func processCPUReady() bool { panic("fake") }
