package workspace

import (
	"sort"
	"sync"
	"time"
)

// FileTracker records all file read and write operations per session for safety and auditability.
type FileTracker struct {
	mu        sync.RWMutex
	reads     map[string]time.Time
	writes    map[string]time.Time
	sessionID string
}

// NewFileTracker creates a new file tracking service instance for a session.
func NewFileTracker(sessionID string) *FileTracker { panic("fake") }

// TrackRead records that a relative file path was read.
func (t *FileTracker) TrackRead(path string) { panic("fake") }

// TrackWrite records that a relative file path was written or modified.
func (t *FileTracker) TrackWrite(path string) { panic("fake") }

// WasRead reports whether the path was read during the current session.
func (t *FileTracker) WasRead(path string) bool { panic("fake") }

// WasWritten reports whether the path was written during the current session.
func (t *FileTracker) WasWritten(path string) bool { panic("fake") }

// ModifiedFiles returns a sorted list of relative paths written during the session.
func (t *FileTracker) ModifiedFiles() []string { panic("fake") }

// TouchedFiles returns a sorted list of all unique relative paths read or written during the session.
func (t *FileTracker) TouchedFiles() []string { panic("fake") }

// Reset clears all tracked reads and writes.
func (t *FileTracker) Reset() { panic("fake") }
