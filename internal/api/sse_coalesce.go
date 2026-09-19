package api

import (
	"strings"
	"sync"
	"time"
)

// ChunkCoalescer batches Ask SSE content deltas so high token/s streams
// (local cluster ~250 tok/s, fast remotes much higher) do not Flush a full
// OpenAI chunk envelope on every upstream piece. Wave A / RT-THPUT-1.
type ChunkCoalescer struct {
	sw    *SSEWriter
	id    string
	model string

	mu       sync.Mutex
	buf      strings.Builder
	timer    *time.Timer
	maxWait  time.Duration
	maxBytes int
	flushN   int // successful content flushes (for tests)
	deltaN   int // WriteDelta calls
	closed   bool
}

// NewChunkCoalescer wraps sw. maxWait is the max time to hold a partial batch
// (≤0 → 12ms). maxBytes forces a flush when the pending delta exceeds it (≤0 → 512).
func NewChunkCoalescer(sw *SSEWriter, id, model string, maxWait time.Duration, maxBytes int) *ChunkCoalescer {
	panic("fake")
}

// WriteDelta appends a content delta and may flush immediately or after maxWait.
func (c *ChunkCoalescer) WriteDelta(chunk string) error { panic("fake") }

// Flush writes any pending delta. Call before stop/[DONE].
func (c *ChunkCoalescer) Flush() error { panic("fake") }

// Stats returns (deltaCalls, contentFlushes) for tests.
func (c *ChunkCoalescer) Stats() (deltas, flushes int) { panic("fake") }

func (c *ChunkCoalescer) flushLocked() error { panic("fake") }
