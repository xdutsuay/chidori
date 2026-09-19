//go:build windows

package api

import (
	"sync"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// KMA-131 — Win32 process CPU/RSS (no ps/wmic/powershell on the UI poll).

var (
	modpsapi                 = windows.NewLazySystemDLL("psapi.dll")
	procGetProcessMemoryInfo = modpsapi.NewProc("GetProcessMemoryInfo")

	winCPUMu       sync.Mutex
	winCPULastWall time.Time
	winCPULastProc int64 // 100ns units (FILETIME)
)

type processMemoryCounters struct {
	CB                         uint32
	PageFaultCount             uint32
	PeakWorkingSetSize         uintptr
	WorkingSetSize             uintptr
	QuotaPeakPagedPoolUsage    uintptr
	QuotaPagedPoolUsage        uintptr
	QuotaPeakNonPagedPoolUsage uintptr
	QuotaNonPagedPoolUsage     uintptr
	PagefileUsage              uintptr
	PeakPagefileUsage          uintptr
}

func filetimeToInt64(ft windows.Filetime) int64 { panic("fake") }

func processCPUMem() (cpuPct, memPct float64, rssKB int64) { panic("fake") }

// FILETIME → seconds

// Match macOS ps %cpu scale (can exceed 100 on multi-core).

// processCPUReady is false until a second sample can compute a delta (KMA-131).
func processCPUReady() bool { panic("fake") }

func processCPUTimeSecs() float64 { panic("fake") }
