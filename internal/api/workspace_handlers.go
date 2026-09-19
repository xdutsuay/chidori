package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/hooks"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// WorkspaceInfo is returned by GET /api/workspace/info.
type WorkspaceInfo struct {
	Enabled           bool     `json:"enabled"`
	Root              string   `json:"root,omitempty"`
	Trusted           bool     `json:"trusted"`
	Recent            []string `json:"recent,omitempty"`
	HasRules          bool     `json:"has_rules"`
	GlobalRulesActive bool     `json:"global_rules_active"` // ADR-0011
	ProjectRulesFile  string   `json:"project_rules_file,omitempty"`
}

func (h *Handlers) workspaceDisabled(w http.ResponseWriter) bool { panic("fake") }

// WorkspaceInfoHandler returns workspace status.
func (h *Handlers) WorkspaceInfoHandler(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceSetRoot opens a new workspace root (e.g. picked via the native
// "Open Folder" dialog), persists it to config.yaml, and — for long-lived
// consumers of the *workspace.FS pointer (registered agent tools, other HTTP
// handlers) — makes the switch take effect immediately, no restart needed.
func (h *Handlers) WorkspaceSetRoot(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ADR-0006 — new root may be Go + (now) trusted, or may not be

// WorkspaceClose clears the current workspace root (File ▸ Close Folder),
// returning to the "no folder open" state — the same state a fresh launch
// with no configured root starts in. Unlike WorkspaceSetRoot, an empty root
// is exactly the point here rather than a rejected input: workspace.FS.
// SetRoot already treats an empty Config.Root as "disable the workspace"
// (see its own doc comment), so this is a thin wrapper around that plus the
// matching config.yaml + LSP-bridge updates WorkspaceSetRoot makes on every
// root change. Recent/trusted lists are left untouched — closing a folder
// isn't forgetting it existed, Open Recent should still offer it.
func (h *Handlers) WorkspaceClose(w http.ResponseWriter, r *http.Request) { panic("fake") }

// empty Root disables it — never errors for an empty root

// no root anymore — stops gopls, same as any other root change

// WorkspaceTrust marks a workspace root trusted or untrusted. A trusted
// workspace is one the user has confirmed they trust the contents of —
// mirroring VS Code's workspace trust prompt. Frontends should gate agent
// auto-apply / write tools on this flag for a freshly opened, unfamiliar
// folder.
func (h *Handlers) WorkspaceTrust(w http.ResponseWriter, r *http.Request) { panic("fake") }

// ADR-0006 — trust is half of gopls's start condition

// WorkspaceRecent lists recently opened workspace roots and which are trusted.
func (h *Handlers) WorkspaceRecent(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) isTrusted(path string) bool { panic("fake") }

func (h *Handlers) recentList() []string { panic("fake") }

// addRecentPath prepends path (deduped) to cfg.Workspace.Recent, capped at 10.
// Caller must hold h.cfgMu.
func addRecentPath(cfg *config.Config, path string) { panic("fake") }

// addTrustedPath adds path to cfg.Workspace.Trusted if not already present.
// Caller must hold h.cfgMu.
func addTrustedPath(cfg *config.Config, path string) { panic("fake") }

// removeTrustedPath drops path from cfg.Workspace.Trusted. Caller must hold h.cfgMu.
func removeTrustedPath(cfg *config.Config, path string) { panic("fake") }

// WorkspaceTree lists directory entries under ?path=.
func (h *Handlers) WorkspaceTree(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceReadFile returns file contents for ?path=.
func (h *Handlers) WorkspaceReadFile(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspacePreview serves binary workspace files (images, etc.) with a real
// Content-Type — WorkspaceReadFile UTF-8 stringifies bytes and breaks PNG preview.
func (h *Handlers) WorkspacePreview(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceFileModTime reports a file's last-modified time (BE.12: the
// frontend polls this for the active file to detect changes made outside
// the IDE — another process, git checkout, etc. — and offer to reload).
func (h *Handlers) WorkspaceFileModTime(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceWriteFile saves file contents from JSON body { path, content }.
// Optional encoding "base64" writes decoded bytes (KMA-102 image attach).
func (h *Handlers) WorkspaceWriteFile(w http.ResponseWriter, r *http.Request) { panic("fake") }

// "" | "base64"

// Chat image attach lands under .lclreason/uploads/; allow up to 8 MiB even
// when workspace.max_file_bytes is still the old 512 KiB default (KMA-102).

// ADR-0013
// ADR-0015

// WorkspaceVision runs a vision model over a workspace image (KMA-117).
// Body: { path, prompt? }. Returns { text, model }. Used so Ask/Agent can
// reason about pasted screenshots even when the selected chat model is text-only.
func (h *Handlers) WorkspaceVision(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceCreate makes a new file or directory from JSON body { path, is_dir }.
// Like WorkspaceWriteFile/Save, this is a direct user action (Explorer context
// menu "New File…"/"New Folder…") and is not gated on workspace trust.
func (h *Handlers) WorkspaceCreate(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceRename renames/moves a file or directory from JSON body { from, to }.
func (h *Handlers) WorkspaceRename(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceDelete removes a file or directory from JSON body { path }.
func (h *Handlers) WorkspaceDelete(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceStats returns an on-demand LOC breakdown by language for the
// open workspace (Settings → Workspace). Not gated on trust — read-only
// metadata, same class as tree/search listings.
func (h *Handlers) WorkspaceStats(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceIndex rebuilds the workspace vector index for agent RAG.
func (h *Handlers) WorkspaceIndex(w http.ResponseWriter, r *http.Request) { panic("fake") }

// rulesMaxBytes caps how much of a .lclreason/rules file gets injected into
// every prompt — these are meant to be a short project brief, not a second
// context window.
const rulesMaxBytes = 6000

// workspaceRulesText reads the project's custom instructions file, if any —
// checked in order: .lclreason/rules.md, .lclreason/rules, .cursorrules (for
// drop-in compatibility with existing Cursor projects). Returns "" if none
// exist, the workspace is disabled, or the file can't be read. Truncated to
// rulesMaxBytes so a huge accidental file doesn't blow the model's context
// budget.
func (h *Handlers) workspaceRulesText() string { panic("fake") }

// workspaceRulesFile reports which rules candidate is actually active for
// the current workspace (empty if none) — same search order as
// workspaceRulesText, kept as its own small function rather than changing
// that one's signature, since most of its callers only want the text.
func (h *Handlers) workspaceRulesFile() string { panic("fake") }

// userGlobalRulesPath returns where the cross-project rules file lives: the
// same directory h.secretsPath resolves to (the app's private per-user data
// dir in the packaged app, the cwd in dev/CLI mode) — deliberately reusing
// that already-resolved path instead of threading a new constructor
// parameter through every api.New() call site just for this.
func (h *Handlers) userGlobalRulesPath() string { panic("fake") }

// userGlobalRulesText (ADR-0011) reads the cross-project rules file — unlike
// workspaceRulesText, this applies to every workspace, not just the current
// one. Same silent-absence contract: no file, empty file, or unreadable all
// just mean "no global rules," never an error.
func (h *Handlers) userGlobalRulesText() string { panic("fake") }

// fullRulesText (ADR-0011) combines global and project rules for prompt
// injection — global first, so a project's rules can extend or override it
// in the model's reading order, without either ever being silently dropped
// when only one of the two exists.
//
// Invariant (do not relax): this text only ever reaches the prompt as
// ordinary instruction content. It must never be re-read anywhere to
// influence isTrusted()/apply-mode decisions — a rules file saying
// "always auto-apply" must have exactly as much effect as a chat message
// saying the same thing, which is to say none, since that's not something
// either surface is allowed to change.
func (h *Handlers) fullRulesText() string { panic("fake") }

// GetGlobalRules returns the cross-project rules file's raw content (ADR-0011)
// for the Settings -> Rules editor. Unlike project rules, this file lives
// outside workspace.FS's sandbox (it's not workspace content), so it needs
// its own small pair of handlers rather than going through
// /api/workspace/file.
func (h *Handlers) GetGlobalRules(w http.ResponseWriter, r *http.Request) { panic("fake") }

// SetGlobalRules writes the cross-project rules file. No trust gating here —
// same direct-user-action precedent as Save/Save As/Explorer CRUD (rule #3
// in HANDOFF.md): a human explicitly editing their own global rules in
// Settings is not an agent-triggered write.
func (h *Handlers) SetGlobalRules(w http.ResponseWriter, r *http.Request) { panic("fake") }

// fireHook loads .lclreason/hooks.json fresh (v1 doesn't cache — call
// frequency here, on_file_save, is low enough that this is simplest, not a
// hot path like pre/post_tool_call would be if this pass wired those in
// too) and fires event. Fire-and-log: never blocks or fails the caller's own
// response on a hook's outcome, so a slow/broken hook can't break saving.
func (h *Handlers) fireHook(event string, env map[string]string) { panic("fake") }

// ServeCodeUI serves the code IDE page.
func (h *Handlers) ServeCodeUI(w http.ResponseWriter, r *http.Request) { panic("fake") }
