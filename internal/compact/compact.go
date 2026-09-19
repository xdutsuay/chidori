// Package compact implements server-side conversation compaction so the prompt
// sent to the model doesn't grow without bound as a chat gets longer.
//
// Strategy (hybrid):
//   - system messages are always kept;
//   - the most recent N messages are kept verbatim (the "recent window");
//   - older messages are condensed into a rolling LLM summary; and
//   - the most relevant older messages are pulled back via embedding similarity
//     (RAG recall) so specific facts survive summarization.
//
// When the flattened history already fits the token budget, nothing changes.
package compact

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/xdutsuay/lclreason/internal/config"
)

// Message is a role-tagged chat message (mirrors api.ChatMessage without the
// dependency).
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Report describes one compaction pass — what went in, what comes out, and the
// exact prompt that will be sent to the inference provider. It is surfaced
// per-user on the dashboard.
type Report struct {
	Strategy        string   `json:"strategy"` // "none" | "hybrid"
	Model           string   `json:"model"`
	BudgetTokens    int      `json:"budget_tokens"`
	OriginalTokens  int      `json:"original_tokens"`
	CompactedTokens int      `json:"compacted_tokens"`
	OriginalMsgs    int      `json:"original_messages"`
	KeptMsgs        int      `json:"kept_messages"`
	RecentKept      int      `json:"recent_kept"`
	Summarized      int      `json:"summarized_messages"`
	Summary         string   `json:"summary,omitempty"`
	Recalled        []string `json:"recalled,omitempty"`
	SentPrompt      string   `json:"sent_prompt"`
}

// Invoker runs an LLM call (used for the rolling summary).
type Invoker interface {
	Invoke(ctx context.Context, model, prompt string) (string, error)
}

// Embedder produces an embedding for text (used for RAG recall).
type Embedder interface {
	Embed(ctx context.Context, text string) ([]float32, error)
}

// Compactor applies the hybrid strategy. Summaries and embeddings are cached so
// a growing conversation doesn't redo work every turn.
type Compactor struct {
	cfg      config.CompactionConfig
	invoker  Invoker
	embedder Embedder

	mu        sync.Mutex
	summaries map[string]string
	embeds    map[string][]float32
}

func New(cfg config.CompactionConfig, invoker Invoker, embedder Embedder) *Compactor { panic("fake") }

// Enabled reports whether compaction is active.
func (c *Compactor) Enabled() bool { panic("fake") }

// Compact returns the compacted message list plus a Report. The original slice
// is never mutated.
func (c *Compactor) Compact(ctx context.Context, msgs []Message, model string) ([]Message, Report) {
	panic("fake")
}

// System messages are always retained; the rest is the conversation.

func (c *Compactor) summarize(ctx context.Context, older []Message, model string) string {
	panic("fake")
}

func (c *Compactor) ragRecall(ctx context.Context, older []Message, query string) []Message {
	panic("fake")
}

// relevance floor

func (c *Compactor) embed(ctx context.Context, text string) ([]float32, error) { panic("fake") }

// ---- helpers ----

// Flatten renders messages as "role: content\n", matching the coordinator's
// prompt format.
func Flatten(msgs []Message) string { panic("fake") }

// EstimateTokens is a cheap heuristic (~4 chars/token).
func EstimateTokens(s string) int { panic("fake") }

func tokensOf(msgs []Message) int { panic("fake") }

// per-message role/formatting overhead

func lastUserContent(msgs []Message) string { panic("fake") }

func cosine(a, b []float32) float64 { panic("fake") }

func hash(s string) string { panic("fake") }

func truncate(s string, n int) string { panic("fake") }
