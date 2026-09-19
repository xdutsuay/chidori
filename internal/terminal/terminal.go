package terminal

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os/exec"
	"sync"

	"github.com/coder/websocket"

	wsenv "github.com/xdutsuay/lclreason/internal/workspace"
)

// Session is one PTY-backed shell scoped to a working directory.
type Session struct {
	id  string
	cmd *exec.Cmd // set on unix; may be nil on Windows ConPTY
	pty io.ReadWriteCloser
	// resize updates cols/rows; kill terminates the shell; close releases the PTY.
	resize func(cols, rows int) error
	kill   func() error
}

// Manager tracks active terminal sessions.
type Manager struct {
	mu       sync.Mutex
	sessions map[string]*Session
}

func NewManager() *Manager { panic("fake") }

// Start creates a shell in cwd and returns session id. shellOverride (T.10
// "Terminal profile") takes precedence over the platform default when
// non-empty — a bare command name resolves via PATH.
func (m *Manager) Start(cwd, shellOverride string) (string, error) { panic("fake") }

// Attach bridges a WebSocket to the PTY until the connection closes.
func (m *Manager) Attach(ctx context.Context, id string, conn *websocket.Conn) error { panic("fake") }

// Kill terminates the shell process for id (if still running) and drops the
// session.
func (m *Manager) Kill(id string) error { panic("fake") }

// Resize updates the PTY window size for id.
func (m *Manager) Resize(id string, cols, rows int) error { panic("fake") }

func newID() string { panic("fake") }

// workspaceEnvVars reads .lclreason/env (T.11).
func workspaceEnvVars(root string) []string { panic("fake") }
