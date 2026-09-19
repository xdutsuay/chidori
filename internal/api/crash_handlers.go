package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/xdutsuay/lclreason/internal/crash"
)

// RecordAbnormalExit writes crash-pending.json via h.crashDir when info is abnormal.
// Nil-safe; empty crashDir is a no-op. Production callers: supervisor / panic hooks.
func (h *Handlers) RecordAbnormalExit(info crash.ExitInfo) error { panic("fake") }

// RecordCrashExit captures an artifact via backend (when non-nil) and writes the
// pending marker into h.crashDir with ArtifactPath and join ids. It is nil-safe:
// empty crashDir or nil backend are no-op / marker-only. Clean exits write nothing.
func (h *Handlers) RecordCrashExit(info crash.ExitInfo, backend crash.CaptureBackend) error {
	panic("fake")
}

// CrashPending is GET /api/crash/pending — next-launch detection for KMA-234.
func (h *Handlers) CrashPending(w http.ResponseWriter, r *http.Request) { panic("fake") }

// CrashDismiss is POST /api/crash/dismiss — clears the pending marker.
func (h *Handlers) CrashDismiss(w http.ResponseWriter, r *http.Request) { panic("fake") }

// CrashReport is POST /api/crash/report — builds a local clipboard report.
// uploaded is always false in this scaffold (no network sink).
func (h *Handlers) CrashReport(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Local file sink when configured; never network-upload in scaffold.

// CrashArtifact is GET /api/crash/artifact — Diagnostics path stub (KMA-225).
func (h *Handlers) CrashArtifact(w http.ResponseWriter, r *http.Request) { panic("fake") }
