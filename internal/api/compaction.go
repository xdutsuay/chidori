package api

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/compact"
)

// userCompaction is the most recent compaction pass for one user (client_id).
type userCompaction struct {
	ClientID  string         `json:"client_id"`
	UpdatedAt time.Time      `json:"updated_at"`
	Report    compact.Report `json:"report"`
}

// compactionLog keeps the latest compaction Report per user so the dashboard
// can show what is actually being sent to the inference provider, per user.
type compactionLog struct {
	mu   sync.Mutex
	byID map[string]userCompaction
}

func newCompactionLog() *compactionLog { panic("fake") }

func (l *compactionLog) record(clientID string, rep compact.Report) { panic("fake") }

func (l *compactionLog) snapshot() []userCompaction { panic("fake") }

// ListCompaction returns the most recent compaction report for each user.
func (h *Handlers) ListCompaction(w http.ResponseWriter, r *http.Request) { panic("fake") }

// compactPrompt runs the conversation through the compactor, records the
// per-user report, and returns the flattened prompt to send to the provider.
// Falls back to a plain flatten when compaction isn't configured, the request
// ctx is already canceled, or Compact returns an empty/partial prompt.
func (h *Handlers) compactPrompt(r *http.Request, req ChatRequest, model string) string {
	panic("fake")
}

// toCompactMessages converts API chat messages to the compact package's type.
func toCompactMessages(msgs []ChatMessage) []compact.Message { panic("fake") }
