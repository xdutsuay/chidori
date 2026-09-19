package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/xdutsuay/lclreason/internal/compact"
	"github.com/xdutsuay/lclreason/internal/session"
)

var htmlTagRe = regexp.MustCompile(`<[^>]*>`)

// stripHTMLForCompaction turns a stored chat message's rendered HTML back
// into plain text — close enough for token estimation and LLM
// summarization, though not a perfect inverse of marked.parse()/
// escapeHtml(). Session messages are stored as HTML (see
// SaveSessionMessages), never as plain text, so this is the bridge back to
// what internal/compact actually expects as input.
func stripHTMLForCompaction(html string) string { panic("fake") }

// SessionUsage reports an estimated token count for a session's stored
// message history against the coordinator's configured compaction budget
// (AI.18 / Hermes Agent's "/usage", see §14.2 #2 in
// llpcodefeature_Cursorclone.md). This is an estimate only — the same
// ~4-chars-per-token heuristic internal/compact itself uses to decide when
// to compact, not a call to the provider's real tokenizer.
func (h *Handlers) SessionUsage(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SessionCompact runs the real compactor (internal/compact — the same one
// used to shrink the prompt sent to the provider on every chat request)
// against a session's stored history, and returns a summary of the older
// portion plus how many of the most recent messages should stay verbatim
// (AI.18 / Hermes Agent's "/compress", §14.2 #2). It does not mutate the
// stored session — the client decides what to keep and persists it on the
// next save, same as any other in-memory chat edit.
func (h *Handlers) SessionCompact(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SearchSessions searches chat message history across every session for a
// workspace (AI.13; also closes the Hermes Agent design study's §14.2 #6
// "cross-session memory + search" idea, internal/session.Store.Search's own
// doc comment has the full cross-reference). GET /api/sessions/search
// ?workspace=<root>&q=<query>.
//
// Wave F: when a session has a stored LLM Summary that matches the query,
// those hits are returned first (role "summary") and message hits for the
// same session are suppressed.
func (h *Handlers) SearchSessions(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SearchSessionMessages searches one session's stored messages (FTS5).
// GET /api/sessions/{id}/search?q=
func (h *Handlers) SearchSessionMessages(w http.ResponseWriter, r *http.Request) { panic("fake") }

func mergeSearchHits(sumHits, msgHits []session.SearchHit, limit int) []session.SearchHit {
	panic("fake")
}

// SessionSummarize asks the inference pool (or a test stub invoker) to
// summarize a session's messages, stores the result on the session row, and
// returns it (Wave F / H.9). Reuses the same compact-style prompt shape as
// internal/compact's rolling summary.
func (h *Handlers) SessionSummarize(w http.ResponseWriter, r *http.Request) { panic("fake") }

// invokeSessionSummary uses a test stub when set, otherwise the dispatch pool.
func (h *Handlers) invokeSessionSummary(ctx context.Context, model, prompt string) (string, error) {
	panic("fake")
}
