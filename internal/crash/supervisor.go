package crash

import "context"

// ChildConfig describes the supervised chidori child process (KMA-225).
type ChildConfig struct {
	Executable string
	Args       []string
	CrashDir   string
}

// ExitInfo is the observed child exit for crash capture / pending markers.
type ExitInfo struct {
	Code      int
	Abnormal  bool
	Signal    string
	SessionID string // ide_sessions.id (sess_*)
	TaskID    string // tasks.id / X-Task-ID
	RequestID string // tasks.request_id (activity correlation key)
}

// Supervisor wraps start/wait for the chidori child. Production impl TBD;
// product UX (pending marker + dialog) is shared across platforms.
type Supervisor interface {
	Start(ctx context.Context, cfg ChildConfig) error
	Wait() (ExitInfo, error)
}

// CaptureBackend writes a platform crash artifact (minidump / core / stderr).
// Windows minidump + WER and Unix signal/core backends plug in here later.
type CaptureBackend interface {
	Capture(ctx context.Context, info ExitInfo) (artifactPath string, err error)
}

// StubCapture is the scaffold CaptureBackend: no real dump yet (TODO Windows minidump).
type StubCapture struct{}

func (StubCapture) Capture(context.Context, ExitInfo) (string, error) { panic(
// TODO(KMA-225): Windows minidump + WER; Unix core/stderr capture.
"fake") }
