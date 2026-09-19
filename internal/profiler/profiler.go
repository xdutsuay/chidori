package profiler

import (
	"log"
	"net/http"
	"time"
)

// Middleware wraps an http.Handler and logs per-request timing.
// Requests slower than threshold are logged at a higher severity.
func Middleware(threshold time.Duration) func(http.Handler) http.Handler { panic("fake") }

// ProfileFunc times a named function call and logs the result.
// Useful for non-HTTP profiling (planner, chain steps, etc.).
func ProfileFunc(name string, fn func() error) (time.Duration, error) { panic("fake") }

// probeWriter captures the status code without interfering with the
// response body. Used by the middleware to log status alongside timing.
type probeWriter struct {
	http.ResponseWriter
	status int
	wrote  bool
}

func (pw *probeWriter) WriteHeader(code int) { panic("fake") }

func (pw *probeWriter) Write(b []byte) (int, error) { panic("fake") }

// Unwrap lets http.ResponseController reach the underlying writer so SSE
// flushing and WebSocket hijacking keep working when profiling is enabled.
func (pw *probeWriter) Unwrap() http.ResponseWriter { panic("fake") }
