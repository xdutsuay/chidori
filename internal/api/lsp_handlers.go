package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/xdutsuay/lclreason/internal/lsp"
	"github.com/xdutsuay/lclreason/internal/workspace"
)

// syncLSPWorkspace starts, restarts, or stops the gopls bridge to match the
// current workspace root/trust state (ADR-0006: "start only when root
// contains go.mod AND the workspace is trusted"). Called after every root or
// trust change (WorkspaceSetRoot, WorkspaceTrust) and once at coordinator
// startup (Server.SyncLSPForWorkspace) for a workspace already configured
// from a previous run. gopls startup is real subprocess + handshake latency,
// so this runs in a goroutine — never block the HTTP response that triggered
// it on gopls being ready.
func (h *Handlers) syncLSPWorkspace() { panic("fake") }

// syncLSPAfterWrite tells gopls about every file a bulk out-of-band write
// (currently: replace-in-files apply) just touched on disk, so a buffer gopls
// already has open for one of those paths doesn't go stale relative to what's
// actually on disk — see Client.SyncFile's doc comment for the failure mode
// this prevents. Best-effort: a sync failure here doesn't undo the write,
// which is already the source of truth, so errors are swallowed the same way
// LSPOpen/LSPChange already treat gopls being unavailable as non-fatal.
func (h *Handlers) syncLSPAfterWrite(ctx context.Context, touched []workspace.ReplaceResult) {
	panic("fake")
}

func hasGoMod(root string) bool { panic("fake") }

// LSPStatus reports whether gopls is running and its last error, for the
// Settings hint (ADR-0006's graceful degradation: "gopls not found" is
// visible here, not a hard failure anywhere else).
func (h *Handlers) LSPStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }

// lspRequestTimeout bounds every per-request LSP call (open/change/hover/
// definition) — these are notifications or quick queries, not long-running
// work; if gopls hasn't finished its handshake or is otherwise stuck, fail
// fast rather than hold the HTTP request open.
const lspRequestTimeout = 5 * time.Second

// LSPOpen is textDocument/didOpen (POST {path, text}). Gopls not running is
// not an error here — ADR-0006's graceful degradation means a `.go` file in
// an untrusted or non-Go-module workspace just quietly gets no LSP features,
// the frontend doesn't need to treat that as a failure.
func (h *Handlers) LSPOpen(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPChange is textDocument/didChange, full-text sync (POST {path, text}) —
// see the ADR's rationale for full over incremental sync.
func (h *Handlers) LSPChange(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPDiagnostics returns the current diagnostics cache, keyed by
// workspace-relative path — polled every 2s by the frontend while a .go file
// is active (see the ADR's transport rationale: REST+poll, not SSE/WS, since
// the packaged app's WKWebView can't receive SSE through Wails' pipeline).
func (h *Handlers) LSPDiagnostics(w http.ResponseWriter, r *http.Request) { panic("fake") }

type lspPositionRequest struct {
	Path      string `json:"path"`
	Line      int    `json:"line"`      // 0-based, LSP-native — see the ADR
	Character int    `json:"character"` // ditto
}

// LSPHover is textDocument/hover (POST {path, line, character}) → {contents}
// (markdown/plaintext, already rendered by gopls — no client-side markdown
// re-parsing needed beyond what marked.js already does for chat). Empty
// contents (not an error) means "nothing to show at this position", which
// Monaco's hover provider treats as "no hover".
func (h *Handlers) LSPHover(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPDefinition is textDocument/definition (POST {path, line, character}) →
// {path, line, character} of the definition site, or {} if none was found —
// the frontend treats a bodyless-but-200 response as "no definition here",
// not a failure.
func (h *Handlers) LSPDefinition(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPDeclaration is textDocument/declaration (POST {path, line, character}).
// Same request/response shape as LSPDefinition — see Client.Declaration's
// own doc comment for why: gopls doesn't implement this LSP method at all,
// so the client falls back to Definition internally, transparently to
// callers here.
func (h *Handlers) LSPDeclaration(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPTypeDefinition is textDocument/typeDefinition (POST {path, line,
// character}) — jumps from a variable's usage to its TYPE's definition
// (e.g. `var f Foo` usage → `type Foo struct` site), a different target
// from LSPDefinition/LSPDeclaration.
func (h *Handlers) LSPTypeDefinition(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPReferences is textDocument/references (POST {path, line, character,
// include_declaration}) → {references: [{path, line, character}, ...]}.
// Unlike Definition/Declaration/TypeDefinition, every result matters here
// (not just the first), so the response is always an array, empty rather
// than omitted when there are no references.
func (h *Handlers) LSPReferences(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPImplementation runs textDocument/implementation (Go to Implementations).
func (h *Handlers) LSPImplementation(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPWorkspaceSymbols runs workspace/symbol (Go to Symbol in Workspace).
func (h *Handlers) LSPWorkspaceSymbols(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPRenamePreview runs textDocument/rename and returns the edits WITHOUT
// writing anything — POST {path, line, character, new_name} →
// {files: [{path, edits: [...]}]}. Same "preview first" shape as
// WorkspaceReplacePreview/WorkspaceReplaceApply (internal/api/ide_handlers.go)
// since a rename, like replace-in-files, can touch multiple files at once —
// a genuinely different shape from OrganizeImports/FormatDocument, which are
// deliberately single-file and apply directly to the open Monaco model.
func (h *Handlers) LSPRenamePreview(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPRenameApply re-runs the rename (never trusts edits computed at preview
// time — the buffer may have changed since) and writes every affected file
// directly to disk via the workspace FS. Not trust-gated: same class of
// direct, explicit user action as WorkspaceReplaceApply and Cmd+S save, not
// an AI-proposed edit being auto-applied.
func (h *Handlers) LSPRenameApply(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) lspRename(w http.ResponseWriter, r *http.Request, apply bool) { panic("fake") }

// Keep gopls's in-memory view in sync with what was just written —
// otherwise a later rename/references request against this same path
// silently computes against stale content (see SyncFile's doc
// comment). Best-effort: the disk write above already succeeded and
// is the source of truth, so a sync failure here shouldn't fail the
// whole rename.

// LSPCodeActions is textDocument/codeAction restricted to quickfix-kind
// actions (ED.15) — POST {path, start_line, start_character, end_line,
// end_character, diagnostics: [{line, character, end_line, end_character,
// severity, message}]} → {actions: [{title, edits: [{path, edits}]}]}.
// The caller (Monaco, via registerCodeActionProvider in
// public/js/09-agent-ask.js) supplies the
// diagnostics it already has as markers for that range — no separate
// diagnostics fetch here, matching how every real LSP client drives this
// request. Unlike Rename, this is NOT a preview-then-apply flow: quickfixes
// are meant to be an instant, single-click fix like Organize Imports/Format
// Document, applied directly to the Monaco model client-side — this handler
// only returns the edits, it never writes to disk itself.
func (h *Handlers) LSPCodeActions(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPDocumentSymbol is textDocument/documentSymbol (GET ?path=) → {symbols}
// for Go to Symbol in Editor. GET+query rather than Hover/Definition's
// POST+body since it needs only a path, no cursor position — same shape as
// LSPDiagnostics above.
func (h *Handlers) LSPDocumentSymbol(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPOrganizeImports is gopls's "source.organizeImports" code action
// (ED.10), POST {path} → {edits}. Returns the edits for the frontend to
// apply itself (executeEdits + the user's normal save flow) — this handler
// never writes the file, same as every other editor-mutation path.
func (h *Handlers) LSPOrganizeImports(w http.ResponseWriter, r *http.Request) { panic("fake") }

// LSPFormatDocument is gopls's textDocument/formatting (ED.9's gofmt-backed
// real Go formatting), POST {path} → {edits}. Same "return edits, never
// write the file" shape as LSPOrganizeImports above.
func (h *Handlers) LSPFormatDocument(w http.ResponseWriter, r *http.Request) { panic("fake") }
