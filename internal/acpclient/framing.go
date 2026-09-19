package acpclient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

// rpcEnvelope is the minimal on-the-wire JSON-RPC 2.0 shape used by FakeAgent
// and the initial ACP client implementation.
type rpcEnvelope struct {
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

func writeNDJSON(w io.Writer, v any) error { panic("fake") }

func readNDJSONLine(r *bufio.Reader) (json.RawMessage, error) { panic("fake") }

// Trim trailing newline; ReadBytes includes it.
