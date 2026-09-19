// Package usage persists per-invocation token counts for Settings → Usage (KMA-87).
//
// This is the global / cumulative / per-source log. It is not the per-session
// compaction estimate served by GET /api/sessions/{id}/usage.
package usage

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
)

const (
	ClassLocal        = "local"
	ClassRemoteKey    = "remote_key"
	ClassUnattributed = "unattributed"

	DefaultRetentionDays = 90
	filePrefix           = "usage-"
	fileSuffix           = ".jsonl"
	dayLayout            = "2006-01-02"
)

// Record is one successful InvokeResult, append-only JSONL.
type Record struct {
	TS               time.Time `json:"ts"`
	SourceClass      string    `json:"source_class"`
	SourceID         string    `json:"source_id,omitempty"`
	Provider         string    `json:"provider,omitempty"`
	Model            string    `json:"model,omitempty"`
	PromptTokens     int       `json:"prompt_tokens"`
	CompletionTokens int       `json:"completion_tokens"`
	RequestID        string    `json:"request_id,omitempty"`
	Mode             string    `json:"mode,omitempty"`
	Host             string    `json:"host,omitempty"`
	KeyFingerprint   string    `json:"key_fingerprint,omitempty"`
	InputNoCache     int       `json:"input_no_cache,omitempty"`
	CacheRead        int       `json:"cache_read,omitempty"`
	CacheWrite       int       `json:"cache_write,omitempty"`
	OutputTokens     int       `json:"output_tokens,omitempty"`
	TotalTokens      int       `json:"total_tokens,omitempty"`
	UsageEstimated   bool      `json:"usage_estimated,omitempty"`
}

// Event is one usage record returned by GET /api/usage/events.
type Event = Record

// SourceRow is one aggregated source in a summary response.
type SourceRow struct {
	SourceClass      string `json:"source_class"`
	SourceID         string `json:"source_id,omitempty"`
	Provider         string `json:"provider,omitempty"`
	Host             string `json:"host,omitempty"`
	KeyFingerprint   string `json:"key_fingerprint,omitempty"`
	Model            string `json:"model,omitempty"`
	PromptTokens     int    `json:"prompt_tokens"`
	CompletionTokens int    `json:"completion_tokens"`
}

// Summary is GET /api/usage/summary.
type Summary struct {
	Range            string      `json:"range"`
	From             time.Time   `json:"from"`
	To               time.Time   `json:"to"`
	PromptTokens     int         `json:"prompt_tokens"`
	CompletionTokens int         `json:"completion_tokens"`
	TotalTokens      int         `json:"total_tokens"`
	Sources          []SourceRow `json:"sources"`
}

// Bucket is one timeseries point (token totals per source class).
type Bucket struct {
	TS           time.Time `json:"ts"`
	Local        int       `json:"local"`
	RemoteKey    int       `json:"remote_key"`
	Unattributed int       `json:"unattributed"`
}

// Timeseries is GET /api/usage/timeseries.
type Timeseries struct {
	Range  string    `json:"range"`
	Bucket string    `json:"bucket"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Series []Bucket  `json:"series"`
}

// Store is an append-only daily-rotated JSONL log.
type Store struct {
	mu     sync.Mutex
	dir    string
	retain time.Duration
	day    string
	f      *os.File
}

// DirFromConfig picks the usage directory: explicit usage.dir, else a
// "usage" folder next to the coordinator DB, else the per-user config dir.
func DirFromConfig(cfg *config.Config) string { panic("fake") }

func defaultUserDir() string { panic("fake") }

func retentionOf(days int) time.Duration { panic("fake") }

// Open creates dir, prunes old files, and returns a ready store.
func Open(dir string, retentionDays int) (*Store, error) { panic("fake") }

// Close the current day's file handle.
func (s *Store) Close() error { panic("fake") }

// Record appends one invocation. Zero-token and nil-store calls are no-ops
// so the dispatch hot path can call this unconditionally.
func (s *Store) Record(rec Record) { panic("fake") }

func (s *Store) append(rec Record) error { panic("fake") }

func (s *Store) ensureFileLocked(day string) error { panic("fake") }

func (s *Store) prune(now time.Time) { panic("fake") }

func (s *Store) pruneLocked(now time.Time) { panic("fake") }

func parseDayFile(name string) (string, bool) { panic("fake") }

// Window is an inclusive-start exclusive-end UTC range.
type Window struct {
	Range string
	From  time.Time
	To    time.Time
}

// ParseRange accepts cycle | 7d | all.
// cycle = current UTC calendar month (local analog of a billing cycle).
func ParseRange(raw string, now time.Time) (Window, error) { panic("fake") }

// ExpandWindowToUTCDayEnd extends w.To to the last instant of its UTC calendar day
// so event export includes the full local day, not just records before time.Now.
func ExpandWindowToUTCDayEnd(w Window) Window { panic("fake") }

// List returns records in w sorted ascending by timestamp. Legacy JSONL lines
// missing cache/token fields get OutputTokens and TotalTokens derived on read.
func (s *Store) List(w Window) []Record { panic("fake") }

func normalizeRecord(rec Record) Record { panic("fake") }

// Summarize aggregates records in w.
func (s *Store) Summarize(w Window) Summary { panic("fake") }

func truncateUTCDay(t time.Time) time.Time { panic("fake") }

// Series returns zero-filled buckets covering w.
// bucket: "hour" (default) or "day" (Settings → Usage chart uses day).
func (s *Store) Series(w Window, bucket string) (Timeseries, error) { panic("fake") }

// "all": start at the oldest file day, or now truncated if empty.

func (s *Store) oldestDay() time.Time { panic("fake") }

func (s *Store) scan(w Window, fn func(Record)) { panic("fake") }

func (s *Store) scanFile(path string, w Window, fn func(Record)) { panic("fake") }

func newID() string { panic("fake") }
