package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// StorageEntry is one row in the storage breakdown drill-down.
type StorageEntry struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	Path       string `json:"path"`
	SizeBytes  int64  `json:"size_bytes"`
	FileCount  int64  `json:"file_count"`
	CreatedMs  int64  `json:"created_ms,omitempty"`
	ModifiedMs int64  `json:"modified_ms,omitempty"`
	Hits       int64  `json:"hits"`
	Deletable  bool   `json:"deletable"`
	Category   string `json:"category"`
	Drillable  bool   `json:"drillable,omitempty"`
	DrillKind  string `json:"drill_kind,omitempty"`
}

type storageScan struct {
	sizeBytes  int64
	fileCount  int64
	created    time.Time
	modified   time.Time
	hits       int64
	hasCreated bool
}

// appStorageRoots lists deduplicated absolute paths whose bytes count toward
// chidori's footprint (shared by stats pill and storage drill-down).
func appStorageRoots(h *Handlers) []string { panic("fake") }

func scanPathTree(path string) storageScan { panic("fake") }

func fileAccessTime(fi os.FileInfo) time.Time { panic("fake") }

func accessHitScore(at, mod time.Time) int64 { panic("fake") }

func storageLabelFor(path string, isDir bool) string { panic("fake") }

func storageCategoryFor(path string) string { panic("fake") }

func isSQLiteSidecarName(name string) bool { panic("fake") }

func sqliteSidecarBytes(dbPath string) (walBytes, shmBytes int64) { panic("fake") }

// sqliteDatabaseDiskBytes is the on-disk footprint of a WAL-mode SQLite file
// (main db + -wal + -shm). The sidecar files are not separate databases.
func sqliteDatabaseDiskBytes(dbPath string) int64 { panic("fake") }

func storageDeletable(path string) bool { panic("fake") }

func collectStorageEntries(h *Handlers) []StorageEntry { panic("fake") }

func sortStorageEntries(entries []StorageEntry, sortBy string) { panic("fake") }

// SystemStorage returns a breakdown of on-disk footprint rows.
func (h *Handlers) SystemStorage(w http.ResponseWriter, r *http.Request) { panic("fake") }

type storageDeleteReq struct {
	Path string `json:"path"`
}

// SystemStorageDelete removes a deletable cache/log path from the breakdown.
func (h *Handlers) SystemStorageDelete(w http.ResponseWriter, r *http.Request) { panic("fake") }
