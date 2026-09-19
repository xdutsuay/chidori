package lsp

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"

	"github.com/xdutsuay/lclreason/internal/executil"
)

// Diagnostic mirrors just the fields v1 needs from LSP's Diagnostic. Positions
// are 0-based UTF-16, exactly as gopls sends them (LSP-native) — conversion
// to Monaco's 1-based positions happens once, at the frontend boundary, per
// ADR-0006's explicit decision. Do not convert in Go.
type Diagnostic struct {
	Line      int    `json:"line"`
	Character int    `json:"character"`
	EndLine   int    `json:"end_line"`
	EndChar   int    `json:"end_character"`
	Severity  int    `json:"severity"` // 1=Error 2=Warning 3=Info 4=Hint, LSP's own numbering
	Message   string `json:"message"`
}

// Position is a 0-based LSP position, used for both Hover/Definition
// requests and responses — see Diagnostic's doc comment on why no conversion
// happens in Go.
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

// Location is a definition result: which file, and where in it.
type Location struct {
	Path string `json:"path"` // workspace-relative, never a file:// URI
	Position
}

// Symbol is one flattened entry from a textDocument/documentSymbol response.
// gopls returns a hierarchy (struct -> its methods, as nested Children) but
// Go to Symbol in Editor just needs a flat, position-ordered pick-list, so
// DocumentSymbol flattens it here rather than pushing tree-walking onto the
// frontend; ContainerName carries the parent's name for nested entries
// (empty for top-level ones) so the picker can still show "helper (on Foo)".
// Line/Character..EndLine/EndCharacter is the symbol's full extent (e.g. a
// whole function body, LSP's `range`) and SelectionLine/SelectionCharacter is
// the precise jump target (just the name token, LSP's `selectionRange`) —
// Monaco's DocumentSymbol wants both, same split LSP itself makes.
type Symbol struct {
	Name               string `json:"name"`
	Kind               int    `json:"kind"` // LSP SymbolKind numbering (12=Function, 6=Method, ...)
	ContainerName      string `json:"container_name,omitempty"`
	Line               int    `json:"line"`
	Character          int    `json:"character"`
	EndLine            int    `json:"end_line"`
	EndCharacter       int    `json:"end_character"`
	SelectionLine      int    `json:"selection_line"`
	SelectionCharacter int    `json:"selection_character"`
}

// Client manages one gopls subprocess for one workspace root. Not safe to use
// concurrently with Start/Stop/Restart from multiple goroutines racing each
// other (the api layer serializes those — trust/root changes are rare,
// user-driven events, not a hot path).
type Client struct {
	mu      sync.Mutex
	cmd     *exec.Cmd
	stdin   *bufio.Writer // wraps cmd.Stdin's pipe; flushed after every write
	root    string
	running bool
	lastErr error

	nextID  atomic.Int64
	pendMu  sync.Mutex
	pending map[int64]chan *envelope

	ready    chan struct{} // closed once initialize/initialized completes (trap #4)
	readDone chan struct{} // closed when the read loop exits (process gone)

	diagMu sync.RWMutex
	diags  map[string][]Diagnostic // workspace-relative path -> diagnostics

	openVersions map[string]int // path -> last didChange version sent, for didChange's required version field
}

// New returns an idle client. Call Start to actually spawn gopls.
func New() *Client { panic("fake") }

// Status reports whether gopls is currently running, and the last error (from
// a failed Start, or the reason the process most recently exited) — surfaced
// via GET /api/lsp/status for the Settings hint (ADR-0006's graceful
// degradation: "gopls not found" is not a hard failure anywhere else).
func (c *Client) Status() (running bool, lastErr error) { panic("fake") }

// Start spawns `gopls serve` rooted at root and performs the
// initialize/initialized handshake. Returns once gopls is ready to accept
// didOpen — callers must not call DidOpen before Start returns successfully.
func (c *Client) Start(root string) error { panic("fake") }

// gopls's own diagnostic logging — not JSON-RPC, keep off stdout/stdin

// running stays false until the handshake actually succeeds (see below) —
// setting it true here, right after spawning, would let Status() report
// "running" before initialize has even been sent, and let a concurrent
// Stop()/Restart() race the in-flight handshake (caught by this client's
// own tests: a Stop() right after seeing running=true killed the process
// out from under handshake(), surfacing as a confusing "gopls process
// exited before responding to initialize" instead of a clean success).

// handshake sends initialize then initialized, and blocks until gopls's
// initialize response arrives (trap #4: no didOpen before this completes).
func (c *Client) handshake(root string) error { panic("fake") }

// hierarchicalDocumentSymbolSupport: true — without this,
// gopls may fall back to the deprecated flat
// SymbolInformation[] shape ({name, kind, location,
// containerName}) instead of the nested DocumentSymbol[]
// ({name, kind, selectionRange, children}) that
// rawDocSymbol/DocumentSymbol below actually decode.

// Stop terminates the gopls process, if running. Idempotent.
func (c *Client) Stop() { panic("fake") }

func (c *Client) stopLocked(reason error) { panic("fake") }

// Restart stops the current process (if any) and starts a fresh one rooted
// at newRoot — the response to workspace-root changes (SetRoot is
// runtime-mutable; gopls holds per-root state, so it cannot just be
// re-pointed, per the ADR).
func (c *Client) Restart(newRoot string) error { panic("fake") }

// wait for the old process's read loop to fully exit first

func (c *Client) readDoneChan() <-chan struct{} { panic("fake") }

// waitForExit reaps the process and marks the client stopped once it exits —
// covers gopls crashing on its own, not just an explicit Stop().
func (c *Client) waitForExit() { panic("fake") }

// Unblock anyone waiting on a response that will now never arrive.

func (c *Client) setErr(err error) { panic("fake") }

type logWriter struct{}

func (logWriter) Write(p []byte) (int, error) { panic("fake") }

// call sends a request and blocks for its response (or ctx expiring). Every
// outbound request the client itself makes (initialize, hover, definition)
// goes through this.
func (c *Client) call(ctx context.Context, method string, params any) (json.RawMessage, error) {
	panic("fake")
}

// notify sends a one-way notification (no response expected) — didOpen,
// didChange, initialized.
func (c *Client) notify(method string, params any) error { panic("fake") }

// reply answers a server→client request (trap #1) — gopls blocks waiting for
// these, so every kindServerRequest the read loop sees must reach here.
func (c *Client) reply(id *int64, result any) error { panic("fake") }

func (c *Client) write(env envelope) error { panic("fake") }

// readLoop is the single reader of gopls's stdout for this client's lifetime.
// Dispatches by kind() (trap #1 is handled entirely here: server requests get
// an empty/default reply so gopls never blocks waiting on us).
func (c *Client) readLoop(r *bufio.Reader) { panic("fake") }

// handleServerRequest answers gopls's own requests to us. v1 doesn't need to
// honor any of these meaningfully — an empty/null result satisfies gopls for
// workspace/configuration, client/registerCapability, and
// window/workDoneProgress/create, the three named in the ADR's trap #1.
func (c *Client) handleServerRequest(env *envelope) { panic("fake") }

// One empty object per requested config item — gopls asks per-scope.

// handleNotification caches diagnostics; every other notification (log
// messages, progress) is intentionally dropped — v1 doesn't surface them.
func (c *Client) handleNotification(env *envelope) { panic("fake") }

// LSP: severity is optional, absent means Error

// Diagnostics returns a snapshot of every file's current diagnostics,
// keyed by workspace-relative path — polled by GET /api/lsp/diagnostics.
func (c *Client) Diagnostics() map[string][]Diagnostic { panic("fake") }

// waitReady blocks until the current handshake completes, or ctx expires.
// Every request method calls this first — trap #4 in the ADR: reads c.ready
// fresh under lock each call, so a Restart() racing a concurrent HTTP
// request always waits on whichever handshake generation is actually
// current, not a stale channel captured before the restart.
func (c *Client) waitReady(ctx context.Context) error { panic("fake") }

// DidOpen tells gopls a file is open with the given content.
func (c *Client) DidOpen(ctx context.Context, path, text string) error { panic("fake") }

// DidChange sends the file's full new content (full-text sync — see the
// ADR's rationale: incremental diffs are the classic bug farm and gopls
// handles full sync fine).
func (c *Client) DidChange(ctx context.Context, path, text string) error { panic("fake") }

// SyncFile brings gopls's in-memory view of path in line with text, whether
// or not gopls has seen this path before — didOpen if this is the first
// time, didChange (bumping the tracked version) if it's already open.
// Needed after writes that land on disk outside the normal editor open/
// change flow (e.g. Rename's apply path writing every affected file
// directly, or replace-in-files' apply path — see syncLSPAfterWrite):
// without this, gopls keeps serving stale in-memory content for any of
// those paths it already had open, and a subsequent request against the
// same path (another rename, a references lookup, diagnostics) silently
// computes against the wrong text instead of what's actually on disk now.
func (c *Client) SyncFile(ctx context.Context, path, text string) error { panic("fake") }

// Hover returns gopls's hover contents (already markdown/plaintext) for a
// 0-based LSP position, or "" if there's nothing to show there.
func (c *Client) Hover(ctx context.Context, path string, pos Position) (string, error) { panic("fake") }

// Definition resolves the symbol at pos to its defining location, or nil if
// none. Only the first result is used — v1 doesn't support multi-target
// go-to-definition (rare for Go; a follow-up if it matters).
func (c *Client) Definition(ctx context.Context, path string, pos Position) (*Location, error) {
	panic("fake")
}

// Declaration runs textDocument/declaration ("Go to Declaration" — distinct
// from Definition for languages with separate forward-declarations, e.g. a
// C header).
//
// Confirmed against a real gopls instance (not assumed from the LSP spec):
// gopls does NOT implement this method at all — it replies with a -32601
// "method not found" RPC error, unconditionally (see TestDeclarationResolves,
// which caught this). Go has no separate declaration-vs-definition concept
// the way a C header does, so the Go team never prioritized wiring it up.
// Rather than expose a "Go to Declaration" command that always errors, this
// falls back to Definition on that specific error — for Go, definition IS
// the meaningful target a "declaration" command would want anyway.
func (c *Client) Declaration(ctx context.Context, path string, pos Position) (*Location, error) {
	panic("fake")
}

// TypeDefinition runs textDocument/typeDefinition ("Go to Type Definition" —
// jumps to a variable's TYPE's definition, not the variable's own
// declaration; e.g. from a `var x Foo` usage straight to `type Foo struct`).
func (c *Client) TypeDefinition(ctx context.Context, path string, pos Position) (*Location, error) {
	panic("fake")
}

// singleLocationRequest is the shared shape behind Definition/Declaration/
// TypeDefinition — all three LSP methods return the identical
// Location[]-or-null response (gopls answers with an array even though at
// most one entry is meaningful for these three; only References needs every
// entry). Factored out because it's the same response-parsing logic
// verified against real gopls three times over, not three independent
// implementations that could quietly drift from each other.
func (c *Client) singleLocationRequest(ctx context.Context, method, path string, pos Position) (*Location, error) {
	panic("fake")
}

//nolint:nilerr // "no result found" is not a failure

// References runs textDocument/references, returning every usage of the
// symbol at pos — unlike Definition/Declaration/TypeDefinition, ALL entries
// matter here (not just the first), and results can span multiple files.
// includeDeclaration controls whether the symbol's own declaration site is
// included alongside its call sites (LSP's own ReferenceContext field).
func (c *Client) References(ctx context.Context, path string, pos Position, includeDeclaration bool) ([]Location, error) {
	panic("fake")
}

// a reference outside the workspace root (e.g. stdlib) — not actionable in the editor, skip rather than leak a raw URI

// Stable, deterministic order for the frontend picker: by file, then by
// position within the file — gopls doesn't guarantee an order on its own.

// Implementation runs textDocument/implementation, returning every
// implementation of the interface/method at pos. Same response shape as
// References (Location[]), same multi-result relevance — all entries matter.
func (c *Client) Implementation(ctx context.Context, path string, pos Position) ([]Location, error) {
	panic("fake")
}

// WorkspaceSymbol is a single result from workspace/symbol.
type WorkspaceSymbol struct {
	Name          string `json:"name"`
	Kind          int    `json:"kind"`
	ContainerName string `json:"container_name,omitempty"`
	Path          string `json:"path"`
	Line          int    `json:"line"`
	Character     int    `json:"character"`
}

// WorkspaceSymbols runs workspace/symbol with the given query string, returning
// matching symbols across the entire workspace (Go to Symbol in Workspace).
func (c *Client) WorkspaceSymbols(ctx context.Context, query string) ([]WorkspaceSymbol, error) {
	panic("fake")
}

// FileEdit is one file's share of a workspace-wide edit (Rename) — unlike
// OrganizeImports/FormatDocument, a rename can touch more than one file, so
// the single-file []TextEdit shape those two use isn't enough here.
type FileEdit struct {
	Path  string     `json:"path"`
	Edits []TextEdit `json:"edits"`
}

// Rename runs textDocument/rename, renaming the symbol at pos to newName
// everywhere it's used across the workspace. Reuses the same
// WorkspaceEdit.documentChanges decoding OrganizeImports already verified
// against real gopls — rename replies with the identical shape — but unlike
// OrganizeImports (deliberately single-file), groups edits by file here
// since a rename routinely spans several.
func (c *Client) Rename(ctx context.Context, path string, pos Position, newName string) ([]FileEdit, error) {
	panic("fake")
}

// CodeAction is one gopls quickfix suggestion. Filtered server-side to kind
// "quickfix" only (via context.only in CodeActions below) — confirmed
// against real gopls that this excludes its many other "source.*" code
// actions (browse docs/assembly/gc details/split package/etc.), which are
// command-based (an arbitrary gopls.* command needing textDocument/
// executeCommand) rather than a direct edit, and aren't fixes at all. Genuine
// quickfixes (e.g. "Add import: \"fmt\"" for an undefined identifier) DO
// carry a direct WorkspaceEdit, same documentChanges shape as Rename/
// OrganizeImports.
type CodeAction struct {
	Title string     `json:"title"`
	Edits []FileEdit `json:"edits"`
}

// CodeActions runs textDocument/codeAction over the given range, restricted
// to quickfix-kind actions, using the diagnostics the caller already knows
// about for that range (Monaco's own markers, via the API handler — no
// separate diagnostics round-trip here, matching how a real LSP client
// always drives this request). Actions with no edit at all (gopls's
// command-only "source.*" actions) are silently dropped: this returns a
// list of one-click fixes, not a command palette.
func (c *Client) CodeActions(ctx context.Context, path string, startLine, startChar, endLine, endChar int, diagnostics []Diagnostic) ([]CodeAction, error) {
	panic("fake")
}

// splitReceiver parses gopls's own naming convention for Go methods: rather
// than nesting a method under its receiver type via LSP's `children` field,
// gopls names it "(Receiver).Method" or "(*Receiver).Method" and returns it
// as a top-level sibling of the type. Found by a real integration test
// against actual gopls, not documented LSP behavior to assume — an earlier
// version of DocumentSymbol assumed methods would arrive as Children of
// their struct and got a flat "(Greeter).Hello" name instead. Splits that
// into Name="Method" + ContainerName="Receiver" (pointer marker stripped —
// ContainerName is a grouping label, not a receiver-kind signal) so a symbol
// picker can show "Method (on Receiver)" instead of the raw gopls string.
func splitReceiver(name string) (plainName, container string) { panic("fake") }

// rawDocSymbol mirrors LSP's hierarchical DocumentSymbol shape (nested
// Children referencing the same type is valid Go — a slice element, not an
// embedded value, so it doesn't recurse infinitely at the type level).
type rawDocSymbol struct {
	Name  string `json:"name"`
	Kind  int    `json:"kind"`
	Range struct {
		Start Position `json:"start"`
		End   Position `json:"end"`
	} `json:"range"`
	SelectionRange struct {
		Start Position `json:"start"`
	} `json:"selectionRange"`
	Children []rawDocSymbol `json:"children"`
}

// DocumentSymbol returns every symbol gopls finds in path, flattened and
// sorted by position (top-to-bottom, matching how Go to Symbol in Editor
// wants to present them) rather than gopls's nested hierarchy.
func (c *Client) DocumentSymbol(ctx context.Context, path string) ([]Symbol, error) { panic("fake") }

// TextEdit is a single LSP-native (0-based) text replacement.
type TextEdit struct {
	StartLine      int    `json:"start_line"`
	StartCharacter int    `json:"start_character"`
	EndLine        int    `json:"end_line"`
	EndCharacter   int    `json:"end_character"`
	NewText        string `json:"new_text"`
}

// positionToByteOffset converts an LSP Position (0-based line, 0-based
// UTF-16 CODE-UNIT character — the LSP spec's own choice, inherited from
// JavaScript/TypeScript's native string representation) to a byte offset
// into src. A direct byte-offset use of Position.Character would corrupt any
// line containing a multi-byte UTF-8 rune before the target column (and an
// astral-plane rune, e.g. most emoji, occupies 2 UTF-16 units but only 1 Go
// rune) — real Go source can contain non-ASCII bytes in comments, string
// literals, or identifiers, so this has to be correct, not just
// ASCII-typical.
func positionToByteOffset(src string, pos Position) int { panic("fake") }

// astral-plane rune = a UTF-16 surrogate pair

// ApplyTextEdits applies a batch of LSP TextEdits to src, in reverse
// document order (last edit first) so an earlier edit's offsets aren't
// invalidated by a later one shifting the text before it — the standard way
// to materialize a batch of LSP edits into real file content. Used by the
// Rename handler (internal/api/lsp_handlers.go) to turn gopls's edit
// descriptions into the bytes actually written to each affected file.
func ApplyTextEdits(src string, edits []TextEdit) string { panic("fake") }

// OrganizeImports runs gopls's "source.organizeImports" code action (ED.10)
// for path and returns the edits to apply. This package never writes the
// file itself — the frontend applies these directly to its own Monaco model
// via executeEdits and the user's normal Save/Auto-Save flow persists it,
// keeping file-write authority in exactly one place, same as every other
// editor-mutation path in this app (agent edits, Format Document, etc.).
func (c *Client) OrganizeImports(ctx context.Context, path string) ([]TextEdit, error) { panic("fake") }

// organizeImports is a whole-file action — gopls doesn't use the
// range to scope it, but codeAction's params require one, so this
// is just a valid placeholder, not a meaningful selection.

// gopls replies with WorkspaceEdit.documentChanges (a list of
// TextDocumentEdit, each carrying its own versioned textDocument), not
// the older flat WorkspaceEdit.changes map — confirmed by dumping the
// raw response against a real gopls instance rather than assumed from
// the LSP spec's simpler example shape; an earlier version of this
// method decoded `changes` and silently got zero edits every time
// despite gopls actually returning a correct fix.

// only this file's edits — organize imports is single-file

// FormatDocument runs gopls's textDocument/formatting (real gofmt-equivalent
// formatting for Go — ED.9's "no language server" limitation, since
// Monaco's own built-in formatter has no Go support at all) and returns the
// edits to apply. Unlike OrganizeImports's codeAction response (a nested
// WorkspaceEdit), formatting replies with a flat TextEdit[] directly —
// confirmed against real gopls rather than assumed, since the two LSP
// methods' response shapes aren't symmetric.
func (c *Client) FormatDocument(ctx context.Context, path string) ([]TextEdit, error) { panic("fake") }

// gofmt is always tabs; gopls ignores insertSpaces for Go anyway

// pathToFileURI builds an RFC 8089 file URI. On Windows abs paths must be
// file:///C:/... (three slashes) — file://C:/... puts the drive in Host and
// breaks gopls (KMA-131).
func pathToFileURI(abs string) string { panic("fake") }

// Ensure leading slash before drive letter.

func fileURIToAbs(uri string) (string, bool) { panic("fake") }

// url.Parse("file:///C:/Users/...") → Path="/C:/Users/..."

// Also handle file://C:/... (legacy wrong form) via Host.

// relToURI / uriToRel are the ONLY place file:// URIs are constructed or
// parsed (trap #3 — never let a file:// URI reach the frontend, and never
// send a workspace-relative path to gopls).
func relToURI(rel, root string) string { panic("fake") }

func uriToRel(uri, root string) (string, bool) { panic("fake") }
