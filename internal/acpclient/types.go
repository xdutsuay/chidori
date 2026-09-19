package acpclient

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/executil"
)

// ACPInitTimeout bounds spawn + initialize + authenticate for Grok Build stdio.
// Prompt/session work uses the caller's ctx (RemoteRequestTimeout on Agent runs).
const ACPInitTimeout = 60 * time.Second

// DefaultGrokArgs returns argv for Grok Build ACP mode when config does not override.
func DefaultGrokArgs(override []string) []string { panic("fake") }

// BinaryInfo captures the resolved external harness binary.
type BinaryInfo struct {
	Path    string
	Version string
}

// LookGrokConfig tunes how we locate and identify the Grok binary.
type LookGrokConfig struct {
	// Command name or absolute/relative path.
	// Defaults to "grok".
	Command string
	// VersionArgs override the probe args (default: ["--version"]).
	VersionArgs []string
}

// SpawnConfig creates one ACP client transport to a spawned process or,
// in tests, to an in-memory transport.
type SpawnConfig struct {
	Command string
	Args    []string
	Env     []string
	// WorkDir is forwarded to session/new as cwd (workspace root for Grok).
	WorkDir string

	// TestTransport, when non-nil, bypasses exec and uses injected stdio.
	TestTransport *stdioTransport
}

type stdioTransport struct {
	// Client writes requests to In; client reads notifications/responses from Out.
	// (Names match the *client's* direction: client.Out is what it reads.)
	In  io.Writer
	Out io.Reader

	// Close is best-effort; for pipes it may be a no-op.
	Close func() error
}

// EventHandler consumes mapped TurnEvents from ACP notifications.
type EventHandler func(agent.TurnEvent)

// Session models an ACP session created by session/new.
type Session interface {
	// Prompt starts an interactive turn; onEvent receives streamed TurnEvents.
	Prompt(ctx context.Context, task string, onEvent EventHandler) error
}

// CommandProbe runs an external binary probe (helper for LookGrok).
func CommandProbe(cmd string, args []string) (string, error) { panic("fake") }
