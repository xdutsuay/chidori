//go:build !windows

package crash

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// platformCapture writes a stderr fallback log on non-Windows platforms. The
// real minidump syscalls are gated to windows_capture_windows.go so this type
// compiles on Linux CI and other non-Windows targets.
func platformCapture(ctx context.Context, artifactDir string, info ExitInfo) (string, string, error) {
	panic("fake")
}
