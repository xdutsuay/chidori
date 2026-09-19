package crash

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// MarkerFileName is the on-disk pending crash marker (KMA-234).
const MarkerFileName = "crash-pending.json"

// Pending is the crash-pending.json payload shown on next launch.
type Pending struct {
	Timestamp    time.Time `json:"timestamp"`
	OS           string    `json:"os"`
	Arch         string    `json:"arch"`
	TaskID       string    `json:"task_id,omitempty"`
	FlightOffset int64     `json:"flight_offset,omitempty"` // extension point for KMA-233; not flushed here
	FlightPath   string    `json:"flight_path,omitempty"`
	ArtifactPath string    `json:"artifact_path,omitempty"`
	Reason       string    `json:"reason,omitempty"`
	ExitCode     int       `json:"exit_code,omitempty"`
}

// DefaultDir returns {dataDir}/crashes.
func DefaultDir(dataDir string) string { panic("fake") }

// MarkerPath returns crashDir/crash-pending.json.
func MarkerPath(crashDir string) string { panic("fake") }

// WritePending creates crashDir if needed and writes the marker atomically.
func WritePending(crashDir string, p Pending) error { panic("fake") }

// ReadPending returns (nil, false, nil) when the marker is absent.
func ReadPending(crashDir string) (*Pending, bool, error) { panic("fake") }

// ClearPending removes the marker; missing file is not an error.
func ClearPending(crashDir string) error { panic("fake") }

// RecordExit writes a pending marker iff info.Abnormal is true (KMA-234).
// Normal exits are a no-op so clean shutdown never leaves a false crash dialog.
func RecordExit(crashDir string, info ExitInfo, extra Pending) error { panic("fake") }

// WritePendingFromExit writes a marker from an abnormal child exit.
// Fills os/arch/timestamp when unset on extra; derives reason from signal or
// "abnormal exit" when extra.Reason is empty.
func WritePendingFromExit(crashDir string, info ExitInfo, extra Pending) error { panic("fake") }

// FormatReport builds a clipboard-friendly local crash report (no upload).
func FormatReport(p Pending, flightTail []string) string { panic("fake") }
