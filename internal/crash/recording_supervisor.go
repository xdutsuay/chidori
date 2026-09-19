package crash

import (
	"context"
	"strings"
)

// RecordingSupervisor wraps a real Supervisor and records crash artifacts for
// abnormal child exits. It stamps transcript-join IDs onto the ExitInfo before
// handing it to a CaptureBackend and writes a crash-pending.json marker.
type RecordingSupervisor struct {
	Inner     Supervisor     // real child supervisor
	Capture   CaptureBackend // platform capture backend; nil is allowed
	CrashDir  string         // pending marker target ({dataDir}/crashes)
	SessionID string         // ide_sessions.id
	TaskID    string         // tasks.id / X-Task-ID
	RequestID string         // tasks.request_id
}

// Start delegates to the inner supervisor.
func (s *RecordingSupervisor) Start(ctx context.Context, cfg ChildConfig) error { panic("fake") }

// Wait delegates to the inner supervisor. On abnormal exit it stamps the join
// IDs, asks the capture backend for an artifact, then writes a pending marker
// with the artifact path. Clean exits write nothing.
func (s *RecordingSupervisor) Wait() (ExitInfo, error) { panic("fake") }

// Stamp transcript-join IDs so the manifest carries them.
