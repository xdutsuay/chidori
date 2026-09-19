package crash

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/google/uuid"
)

// WritePanicDump creates a per-panic artifact under {dataDir}/crashes containing
// panic.log (panic value + debug.Stack) and manifest.json with kind "panic".
// Returns the artifact directory path. Used by in-process recover hooks when no
// supervised child exists (KMA-225 slice 2 — no re-exec launcher).
func WritePanicDump(dataDir string, recovered any) (string, error) { panic("fake") }

// RecordPanic writes a panic self-dump under dataDir and leaves crash-pending.json
// with ArtifactPath set so Diagnostics / next-launch dialog can open the dump.
func RecordPanic(dataDir string, recovered any) error { panic("fake") }
