// Package lsp bridges chidori to a Go language server (gopls) over stdio,
// per docs/adr/ADR-0006-gopls-lsp-bridge.md: diagnostics, hover, and
// go-to-definition for Go only, v1. Hand-rolled JSON-RPC framing rather than
// a new dependency, matching this project's "single binary, minimal deps"
// convention (see HANDOFF.md rule 5).
package lsp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// envelope is the on-the-wire shape for every LSP/JSON-RPC message. A
// pointer ID (rather than a bare int64) is what lets the read loop tell a
// notification (no ID at all) apart from a request/response whose ID happens
// to be the zero value.
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

// kind classifies a decoded envelope for the read loop's dispatch switch —
// see client.go's readLoop, and trap #1 in the ADR: gopls sends server→client
// *requests* (kindServerRequest) that MUST be answered or it stalls.
type kind int

const (
	kindResponse kind = iota
	kindServerRequest
	kindNotification
)

func (e *envelope) kind() kind { panic("fake") }

// writeMessage frames v as `Content-Length: N\r\n\r\n<json>` — the LSP wire
// format (a restricted form of the JSON-RPC "base protocol"). No trailing
// newline after the body; Content-Length is exact.
func writeMessage(w io.Writer, v any) error { panic("fake") }

// readMessage reads one framed message: headers (only Content-Length is
// meaningful; others, e.g. Content-Type, are skipped) terminated by a blank
// line, then exactly Content-Length bytes of JSON body.
func readMessage(r *bufio.Reader) (*envelope, error) { panic("fake") }

// blank line ends the header block

// int64Ptr is a small helper since Go has no literal syntax for "pointer to
// this int64 value".
func int64Ptr(v int64) *int64 { panic("fake") }

// isEOF reports whether err is (or wraps) io.EOF/io.ErrUnexpectedEOF — the
// read loop uses this to distinguish "gopls process exited" (expected on
// Stop/Restart) from a genuine decode error worth logging.
func isEOF(err error) bool { panic("fake") }
