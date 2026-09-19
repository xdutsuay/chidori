//go:build windows

package crash

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

// platformCapture writes a real Windows minidump for the current process.
func platformCapture(ctx context.Context, artifactDir string, info ExitInfo) (string, string, error) {
	panic("fake")
}

var (
	modDbghelp            = windows.NewLazySystemDLL("dbghelp.dll")
	procMiniDumpWriteDump = modDbghelp.NewProc("MiniDumpWriteDump")
)

const miniDumpNormal = 0x00000001
