// Package mcpclient lets chidori's agent consume external MCP (Model
// Context Protocol) servers as tools, per docs/adr/... ADR-0012 (see
// markdowns_open/LLP2-autonomy-platform.md §3.1). This is the reverse
// direction of cmd/lclreason/mcp.go, which exposes chidori itself as an MCP
// server — unrelated code, not touched by this package.
//
// v1 scope: stdio transport only (the overwhelmingly common case for local
// MCP servers, e.g. `npx @modelcontextprotocol/server-filesystem`). SSE
// transport is real but deliberately deferred — land stdio first.
package mcpclient

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// envelope is the on-the-wire JSON-RPC 2.0 shape. Structurally identical to
// internal/lsp's own envelope/framing (both are the same base protocol) —
// deliberately duplicated rather than factored into a shared package for
// now, to avoid touching internal/lsp's already-shipped, already-tested
// code for the sake of DRY. A reasonable extraction candidate later if a
// third consumer shows up.
type envelope struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      *int64          `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *rpcError) Error() string { panic("fake") }

func (e *envelope) isResponse() bool { panic("fake") }

// writeMessage frames v as `Content-Length: N\r\n\r\n<json>` — the same
// wire format LSP uses, since both are the JSON-RPC 2.0 base protocol over
// stdio.
func writeMessage(w io.Writer, v any) error { panic("fake") }

// readMessage reads one framed message: headers terminated by a blank line
// (only Content-Length matters), then exactly that many bytes of JSON body.
func readMessage(r *bufio.Reader) (*envelope, error) { panic("fake") }

func int64Ptr(v int64) *int64 { panic("fake") }

// isEOF reports whether err means "the server process exited" (expected on
// Close) rather than a genuine decode error worth surfacing.
func isEOF(err error) bool { panic("fake") }
