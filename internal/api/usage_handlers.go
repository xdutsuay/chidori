package api

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/usage"
)

// UsageSummary is GET /api/usage/summary?range=cycle|7d|all (KMA-87).
// Separate from GET /api/sessions/{id}/usage (per-session compaction estimate).
func (h *Handlers) UsageSummary(w http.ResponseWriter, r *http.Request) { panic("fake") }

// UsageTimeseries is GET /api/usage/timeseries?range=…&bucket=hour|day (KMA-87).
func (h *Handlers) UsageTimeseries(w http.ResponseWriter, r *http.Request) { panic("fake") }

// UsageEvents is GET /api/usage/events?range=cycle|7d|all — JSON array of events.
func (h *Handlers) UsageEvents(w http.ResponseWriter, r *http.Request) { panic("fake") }

// UsageExportCSV is GET /api/usage/export.csv?range=cycle|7d|all — Cursor-style CSV.
func (h *Handlers) UsageExportCSV(w http.ResponseWriter, r *http.Request) { panic("fake") }

func kindFromSourceClass(class string) string { panic("fake") }

func formatUsageExportDate(t time.Time) string { panic("fake") }

func (h *Handlers) usageStore() *usage.Store { panic("fake") }

func openUsageStore(cfg *config.Config) *usage.Store { panic("fake") }

// Tests construct api.New with a relative DB in the package directory.
// Don't drop JSONL next to *_test.go unless the test opted in with usage.dir.

func retentionDays(n int) int { panic("fake") }

func attachUsageHooks(st *usage.Store, dis *dispatch.Dispatcher, pool *dispatch.Pool) { panic("fake") }

func recordFromInvoke(res dispatch.InvokeResult) usage.Record { panic("fake") }
