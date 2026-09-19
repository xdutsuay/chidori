package session

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/store"
)

// Session is a persisted IDE chat/editor state.
type Session struct {
	ID            string    `json:"id"`
	WorkspaceRoot string    `json:"workspace_root"`
	Title         string    `json:"title"`
	Mode          string    `json:"mode"`
	Provider      string    `json:"provider"`
	Model         string    `json:"model"`
	ActivePath    string    `json:"active_path"`
	OpenTabs      []string  `json:"open_tabs"`
	Summary       string    `json:"summary,omitempty"` // LLM session summary (Wave F / H.9)
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Message is one chat message in a session.
type Message struct {
	ID          int64     `json:"id"`
	SessionID   string    `json:"session_id"`
	Role        string    `json:"role"`
	HTML        string    `json:"html"`
	ContentJSON string    `json:"content_json,omitempty"`
	RequestID   string    `json:"request_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// Store persists IDE sessions in SQLite.
type Store struct {
	db *sql.DB
}

// New attaches to an existing store DB and ensures session tables exist.
func New(st *store.Store) (*Store, error) { panic("fake") }

func (s *Store) migrate() error { panic("fake") }

// stripHTML turns a stored chat message's rendered HTML back into plain text
// for indexing — a duplicate of internal/api's stripHTMLForCompaction (same
// small regex approach) rather than a cross-package dependency, since session
// (storage) shouldn't import api (handlers) just for a 5-line helper.
var htmlTagRe = regexp.MustCompile(`<[^>]*>`)
var modelTextCommentRe = regexp.MustCompile(`<!--\s*model-text:([A-Za-z0-9+/=]+)\s*-->`)
var uploadPathRe = regexp.MustCompile(`\.lclreason/uploads/[^\s"'<>\\]+`)

// PlainText returns searchable plain text for a stored chat HTML blob
// (prefers embedded model-text comments when present).
func PlainText(html string) string { panic("fake") }

func stripHTML(html string) string { panic("fake") }

func decodeModelTextComment(b64 string) (string, error) { panic("fake") }

// List returns recent sessions, optionally filtered by workspace root.
func (s *Store) List(ctx context.Context, workspaceRoot string, limit int) ([]Session, error) {
	panic("fake")
}

// Get loads one session by id.
func (s *Store) Get(ctx context.Context, id string) (*Session, error) { panic("fake") }

// Create inserts a new session.
func (s *Store) Create(ctx context.Context, sess Session) error { panic("fake") }

// Update patches session metadata.
func (s *Store) Update(ctx context.Context, sess Session) error { panic("fake") }

// Delete removes a session, its messages, and the standalone FTS index rows.
// FTS5 is not a child of ide_sessions, so ON DELETE CASCADE cannot clean it.
func (s *Store) Delete(ctx context.Context, id string) error { panic("fake") }

// SetSummary stores an LLM-generated session summary (Wave F / H.9).
func (s *Store) SetSummary(ctx context.Context, id, summary string) error { panic("fake") }

// GetMessages returns all messages for a session in order.
func (s *Store) GetMessages(ctx context.Context, sessionID string) ([]Message, error) { panic("fake") }

// AppendMessage adds one message to a session, and indexes it for search
// (AI.13 / Hermes Agent §14.2 #6 — cross-session chat history search).
func (s *Store) AppendMessage(ctx context.Context, sessionID, role, html string) (int64, error) {
	panic("fake")
}

// AppendMessageStructured stores content_json as source of truth and derives html (KMA-217).
func (s *Store) AppendMessageStructured(ctx context.Context, sessionID, role string, blocks []ContentBlock) (int64, error) {
	panic("fake")
}

func (s *Store) appendMessageRow(ctx context.Context, sessionID, role, html, contentJSON, requestID string) (int64, error) {
	panic("fake")
}

func nullIfEmpty(s string) any { panic("fake") }

// ReplaceMessages clears and rewrites all messages (bulk save) and rebuilds
// that session's search index entries to match.
func (s *Store) ReplaceMessages(ctx context.Context, sessionID string, msgs []Message) error {
	panic("fake")
}

// SearchHit is one message matching a cross-session search query.
type SearchHit struct {
	SessionID    string    `json:"session_id"`
	SessionTitle string    `json:"session_title"`
	Role         string    `json:"role"`
	Snippet      string    `json:"snippet"` // plain text, match wrapped in **bold** (FTS5 snippet())
	CreatedAt    time.Time `json:"created_at,omitempty"`
	HTML         string    `json:"-"`
	Attachments  []string  `json:"attachments,omitempty"`
	VisionText   string    `json:"vision_text,omitempty"`
}

// sanitizeFTSQuery quotes each word of a raw search box string as its own FTS5
// phrase, joined with FTS5's default implicit AND. This is what makes free-
// typed search text ("fix auth bug") behave like a normal "contains all these
// words" search instead of being parsed as an FTS5 query expression — a bare
// hyphen, asterisk, or unbalanced quote in user input would otherwise be a
// syntax error (FTS5 MATCH) rather than a search with no results.
func sanitizeFTSQuery(raw string) string { panic("fake") }

// Search finds messages matching query across every session for
// workspaceRoot, most-recently-updated session first. Scoped to one
// workspace root so search never surfaces messages from a different project
// the user happened to open previously (AI.13; also closes the Hermes Agent
// design study's §14.2 #6 "cross-session memory + search" idea — same
// feature from that doc's perspective).
func (s *Store) Search(ctx context.Context, workspaceRoot, query string, limit int) ([]SearchHit, error) {
	panic("fake")
}

// SearchInSession finds messages matching query within one chat session (FTS5).
func (s *Store) SearchInSession(ctx context.Context, sessionID, query string, limit int) ([]SearchHit, error) {
	panic("fake")
}

// SearchWorkspace finds messages (and summary hits) across every session for a workspace.
func (s *Store) SearchWorkspace(ctx context.Context, workspaceRoot, query string, limit int) ([]SearchHit, error) {
	panic("fake")
}

// SearchSummaries finds sessions whose stored LLM summary matches query
// (case-insensitive substring of every word). Used by SearchSessions to
// prefer summary hits when present (Wave F).
func (s *Store) SearchSummaries(ctx context.Context, workspaceRoot, query string, limit int) ([]SearchHit, error) {
	panic("fake")
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanSession(r rowScanner) (Session, error) { panic("fake") }

// NewID generates a simple session id.
func NewID() string { panic("fake") }

func (s *Store) LastToolScratch(ctx context.Context, sessionID string) string { panic("fake") }

func (s *Store) SetLastToolScratch(ctx context.Context, sessionID, scratch string) error {
	panic("fake")
}

func extractUploadPaths(html string) []string { panic("fake") }

func extractVisionText(html string) string { panic("fake") }

func isRecursiveSearchNoise(snippet, html string) bool { panic("fake") }

func (s *Store) enrichHits(ctx context.Context, hits []SearchHit) []SearchHit { panic("fake") }

// Attachment hits first so a prior search dump cannot outrank the real photo.

// FindAttachment locates a chat message that references an uploads path.
func (s *Store) FindAttachment(ctx context.Context, sessionID, path string) (SearchHit, error) {
	panic("fake")
}
