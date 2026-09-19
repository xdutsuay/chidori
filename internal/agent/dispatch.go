package agent

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrSessionBusy = errors.New("session is currently busy executing another agent turn")
)

// SessionDispatch manages per-session synchronization and concurrency for agent execution.
type SessionDispatch struct {
	mu      sync.Map // sessionID -> *sync.Mutex
	active  sync.Map // sessionID -> bool
	cancels sync.Map // sessionID -> context.CancelFunc
}

// NewSessionDispatch creates a new SessionDispatch controller.
func NewSessionDispatch() *SessionDispatch { panic("fake") }

// sessionMu gets or creates the per-session mutex.
func (d *SessionDispatch) sessionMu(sessionID string) *sync.Mutex { panic("fake") }

// IsActive returns true if the session currently has an active agent turn.
func (d *SessionDispatch) IsActive(sessionID string) bool { panic("fake") }

// Dispatch executes fn under the session lock, ensuring serial execution per session.
// If a cancel function is returned by fn, it is registered for cancellation support.
func (d *SessionDispatch) Dispatch(ctx context.Context, sessionID string, fn func(ctx context.Context) error) error {
	panic("fake")
}

// Cancel cancels the active run context for the specified session if active.
func (d *SessionDispatch) Cancel(sessionID string) { panic("fake") }
