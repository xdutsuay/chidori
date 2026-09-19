package api

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/store"
)

// StorageSidecar is a SQLite -wal / -shm sibling file.
type StorageSidecar struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	SizeBytes int64  `json:"size_bytes"`
}

// StorageTableStat is one SQLite table inside the session database.
type StorageTableStat struct {
	Name     string `json:"name"`
	RowCount int64  `json:"row_count"`
	EstBytes int64  `json:"est_bytes"`
	Kind     string `json:"kind,omitempty"`
}

// StorageSessionRow is one chat session ranked by on-disk message bulk.
type StorageSessionRow struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	WorkspaceRoot string `json:"workspace_root"`
	MessageCount  int64  `json:"message_count"`
	BytesHTML     int64  `json:"bytes_html"`
	UpdatedMs     int64  `json:"updated_ms"`
	Deletable     bool   `json:"deletable"`
}

// StorageDeepResponse is the drill-down payload for lclreason.db.
type StorageDeepResponse struct {
	Kind          string              `json:"kind"`
	Path          string              `json:"path"`
	FileBytes     int64               `json:"file_bytes"`
	PageSize      int64               `json:"page_size"`
	PageCount     int64               `json:"page_count"`
	FreelistPages int64               `json:"freelist_pages"`
	WALMode       string              `json:"wal_mode"`
	Sidecars      []StorageSidecar    `json:"sidecars"`
	Tables        []StorageTableStat  `json:"tables"`
	Sessions      []StorageSessionRow `json:"sessions"`
	Sort          string              `json:"sort"`
	GeneratedMs   int64               `json:"generated_ms"`
	SessionCount  int64               `json:"session_count"`
	MessageCount  int64               `json:"message_count"`
}

func configuredDBPath(h *Handlers) string { panic("fake") }

func authorizedDBPath(h *Handlers, raw string) (string, bool) { panic("fake") }

func handlersDB(h *Handlers) *sql.DB { panic("fake") }

func sidecarFiles(dbPath string) []StorageSidecar { panic("fake") }

func queryTableStat(db *sql.DB, name, kind, query string) StorageTableStat { panic("fake") }

func collectSessionDBDeep(h *Handlers, dbPath, sortBy string) (*StorageDeepResponse, error) {
	panic("fake")
}

func sortStorageSessions(rows []StorageSessionRow, sortBy string) { panic("fake") }

// SystemStorageDeep returns a session-database breakdown (tables + per-session bulk).
func (h *Handlers) SystemStorageDeep(w http.ResponseWriter, r *http.Request) { panic("fake") }

func pruneAgeFromRequest(r *http.Request) (age time.Duration, days int) { panic("fake") }

// SystemStoragePrunePreview estimates reclaim from deleting finished tasks older than the cutoff.
func (h *Handlers) SystemStoragePrunePreview(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SystemStorageVacuum runs SQLite VACUUM on the configured session database.
func (h *Handlers) SystemStorageVacuum(w http.ResponseWriter, r *http.Request) { panic("fake") }
