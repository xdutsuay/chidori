package api

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const defaultSSEWriteDeadline = 30 * time.Second

// SSEWriter is a thin wrapper around http.ResponseWriter that writes
// Server-Sent Events and flushes after each write so tokens reach the client
// chunk-by-chunk with no full-response buffering.
//
// It uses http.ResponseController (Go 1.20+) rather than a direct
// http.Flusher type assertion, so flushing still works when the writer is
// wrapped by middleware (e.g. the profiler's probeWriter, which exposes
// Unwrap()).
type SSEWriter struct {
	w  http.ResponseWriter
	rc *http.ResponseController

	writeDeadline time.Duration
	onWriteError  func(error)
	writeErrOnce  sync.Once
}

// NewSSEWriter prepares the response for SSE: sets the right headers and a
// flush controller. A probe flush surfaces "not supported" early so callers
// don't silently buffer.
func NewSSEWriter(w http.ResponseWriter) (*SSEWriter, error) { panic("fake") }

// disable buffering on nginx/etc.

// SetWriteDeadline configures the per-write deadline (default 30s when unset).
func (s *SSEWriter) SetWriteDeadline(d time.Duration) { panic("fake") }

// SetOnWriteError registers a callback invoked once on the first write failure.
// Handlers typically cancel the request context so in-flight agent work unwinds.
func (s *SSEWriter) SetOnWriteError(fn func(error)) { panic("fake") }

func (s *SSEWriter) write(doWrite func() error) error { panic("fake") }

// WriteData writes a single SSE event with the default event name.
// payload should already be marshalled (typically JSON).
func (s *SSEWriter) WriteData(payload string) error { panic("fake") }

// WriteEvent writes a named SSE event. Empty event uses the default.
func (s *SSEWriter) WriteEvent(event, data string) error { panic("fake") }

// Done writes the conventional [DONE] sentinel used by OpenAI-style
// streaming clients.
func (s *SSEWriter) Done() error { panic("fake") }
