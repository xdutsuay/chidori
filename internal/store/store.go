package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Event struct {
	ID         int64
	Kind       string
	NodeID     string
	Model      string
	Prompt     string
	Response   string
	DurationMs int64
	CreatedAt  time.Time
}

type Task struct {
	ID         string    `json:"id"`
	ClientID   string    `json:"client_id"`
	SessionID  string    `json:"session_id,omitempty"`
	RequestID  string    `json:"request_id,omitempty"`
	Prompt     string    `json:"prompt"`
	Status     string    `json:"status"`
	Model      string    `json:"model"`
	Chain      string    `json:"chain"`
	Worker     string    `json:"worker"`
	Response   string    `json:"response"`
	Error      string    `json:"error,omitempty"`
	DurationMs int64     `json:"duration_ms"`
	PlanJSON   string    `json:"plan_json,omitempty"`
	TraceJSON  string    `json:"trace_json,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Store struct {
	db     *sql.DB
	dbPath string
}

func Open(path string) (*Store, error) { panic("fake") }

// WAL mode + a single serialized connection (running_issue.md Code
// Critic #2 / H.12): database/sql's default pool can open multiple
// concurrent connections to the same SQLite file, and every activity/
// task/event write from a busy coordinator goes through this store —
// under concurrent load that surfaced as "database is locked". WAL lets
// readers and a writer coexist without blocking each other on ordinary
// SELECTs; capping the pool at one connection is the strict
// single-writer pattern that removes the remaining lock contention
// outright rather than just making it rarer. Verified this driver
// (modernc.org/sqlite) applies both PRAGMAs from one multi-statement
// Exec rather than silently only running the first.

func (s *Store) Close() error { panic("fake") }

// DB exposes the underlying database for extensions (sessions, etc.).
func (s *Store) DB() *sql.DB { panic("fake") }

func (s *Store) migrate() error { panic("fake") }

// migrateTaskJoinColumns adds KMA-219 session_id/request_id on legacy DBs.
func (s *Store) migrateTaskJoinColumns() error { panic("fake") }

func (s *Store) Log(ctx context.Context, e Event) error { panic("fake") }

func (s *Store) Recent(ctx context.Context, limit int) ([]Event, error) { panic("fake") }

// RecentByKind returns events filtered by kind, newest first.
func (s *Store) RecentByKind(ctx context.Context, kind string, limit int) ([]Event, error) {
	panic("fake")
}

// MaxTaskFieldBytes caps coordinator task blobs so a failed-request flood
// cannot grow lclreason.db by gigabytes (KMA-135). 8 KiB keeps a useful
// prompt prefix for the dashboard without retaining full context dumps.
const MaxTaskFieldBytes = 8192

// DefaultTaskPruneAge is how long finished coordinator tasks are kept.
const DefaultTaskPruneAge = 7 * 24 * time.Hour

func capTaskField(s string) string { panic("fake") }

const finishedTaskStatuses = `lower(status) IN ('failed','success','cancelled','canceled','completed','done','error')`

func pruneCutoff(olderThan time.Duration) string { panic("fake") }

// PruneEstimate is the approximate reclaim from deleting finished tasks.
type PruneEstimate struct {
	Tasks int64
	Bytes int64
}

// EstimatePruneFinishedTasks reports how many finished tasks (and payload bytes)
// would be deleted for olderThan without modifying the database.
func (s *Store) EstimatePruneFinishedTasks(ctx context.Context, olderThan time.Duration) (PruneEstimate, error) {
	panic("fake")
}

// PruneFinishedTasks deletes finished coordinator tasks older than olderThan.
// Running/pending rows are never removed. Returns the number of deleted rows.
func (s *Store) PruneFinishedTasks(ctx context.Context, olderThan time.Duration) (int64, error) {
	panic("fake")
}

// UpsertTask inserts a task or updates it if the id already exists.
func (s *Store) UpsertTask(ctx context.Context, t Task) error { panic("fake") }

// GetTasks returns recent tasks ordered by created_at DESC.
func (s *Store) GetTasks(ctx context.Context, limit int) ([]Task, error) { panic("fake") }

// GetTask returns a single task by id, or sql.ErrNoRows if not found.
func (s *Store) GetTask(ctx context.Context, id string) (*Task, error) { panic("fake") }

type scanner interface {
	Scan(dest ...any) error
}

func scanTask(r scanner) (Task, error) { panic("fake") }

// ---- memos (shared clipboard) ----

type Memo struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	Author    string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// AddMemo inserts a new memo.
func (s *Store) AddMemo(ctx context.Context, text, author string) error { panic("fake") }

// GetMemos returns recent memos, newest first.
func (s *Store) GetMemos(ctx context.Context, limit int) ([]Memo, error) { panic("fake") }

// DeleteMemo removes a memo by id.
func (s *Store) DeleteMemo(ctx context.Context, id int64) error { panic("fake") }
