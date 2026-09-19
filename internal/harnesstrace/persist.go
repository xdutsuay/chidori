package harnesstrace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// FileStore appends trace events to JSONL files under root/.lclreason/harness-traces/.
type FileStore struct {
	mu   sync.Mutex
	root string
}

func TracesDir(root string) string { panic("fake") }

func NewFileStore(root string) *FileStore { panic("fake") }

func (s *FileStore) SetRoot(root string) { panic("fake") }

func (s *FileStore) Root() string { panic("fake") }

type persistedLine struct {
	Kind      string         `json:"kind"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload,omitempty"`
}

// Append writes one event line to id.jsonl (creates dir as needed).
func (s *FileStore) Append(traceID string, e Event) error { panic("fake") }

// TraceMeta is one persisted trace file.
type TraceMeta struct {
	ID       string `json:"id"`
	Path     string `json:"path"`
	Bytes    int64  `json:"bytes"`
	Modified int64  `json:"modified_ms"`
}

// List traces on disk, newest first.
func (s *FileStore) List() ([]TraceMeta, error) { panic("fake") }

// Read returns decoded events for one trace id.
func (s *FileStore) Read(traceID string) ([]Event, error) { panic("fake") }
