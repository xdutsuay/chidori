package crash

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/google/uuid"
)

// ArtifactManifest is the on-disk shape of a crash artifact manifest.json.
type ArtifactManifest struct {
	Timestamp time.Time `json:"timestamp"`
	OS        string    `json:"os"`
	Arch      string    `json:"arch"`
	SessionID string    `json:"session_id"`
	TaskID    string    `json:"task_id"`
	RequestID string    `json:"request_id"`
	ExitCode  int       `json:"exit_code"`
	Kind      string    `json:"kind"` // "minidump", "stderr", or "panic"
	Signal    string    `json:"signal,omitempty"`
	DumpPath  string    `json:"dump_path"`
}

// WindowsCapture is a CaptureBackend that writes a platform crash artifact for
// abnormal child exits. On Windows it attempts a real minidump; on all other
// platforms it falls back to a stderr log so the type compiles and tests run
// everywhere.
type WindowsCapture struct {
	DataDir string // root data directory; artifacts live under {DataDir}/crashes
}

// Capture implements CaptureBackend. Clean exits return ("", nil). Abnormal
// exits create a per-crash directory under {DataDir}/crashes containing
// manifest.json and a capture payload (minidump.dmp on Windows, stderr.log
// elsewhere). The returned path is the per-crash artifact directory.
func (w WindowsCapture) Capture(ctx context.Context, info ExitInfo) (string, error) { panic("fake") }
