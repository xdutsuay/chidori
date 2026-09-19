package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/harnesstrace"
)

var errNotStreaming = errors.New("upstream refused streaming")

const (
	mcpToolTimeout      = 120 * time.Second
	mcpSubmitChatTimout = 300 * time.Second
	mcpWarmAfter        = 5 * time.Second
)

// runMCP starts an MCP (Model Context Protocol) server over stdio. It speaks
// newline-delimited JSON-RPC 2.0 and exposes the coordinator's REST API as MCP
// tools, so Claude Desktop / Cursor can drive an lclreason cluster.
//
// It is a thin proxy: every tool call maps to an HTTP request against a running
// coordinator (cfg.Coordinator, or 127.0.0.1:<port> by default).
func runMCP(cfg *config.Config) { panic("fake") }

// Env override so the MCP proxy can point at a non-default coordinator
// without editing config.yaml — useful alongside KMA-88's "don't bump the
// port" fix when running multiple coordinator instances.

// keep stdout clean for the JSON-RPC channel

type mcpServer struct {
	base          string
	apiKey        string
	client        *http.Client
	out           io.Writer
	mu            sync.Mutex
	progressToken string
}

type rpcRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type rpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *mcpServer) serve(in io.Reader, out io.Writer) { panic("fake") }

func (s *mcpServer) writeJSON(v any) { panic("fake") }

func (s *mcpServer) handle(req rpcRequest) (rpcResponse, bool) { panic("fake") }

// notification — no reply

// unknown notification — ignore

func mcpProgressToken(raw json.RawMessage) string { panic("fake") }

func (s *mcpServer) emitProgress(progress int, message string) { panic("fake") }

func (s *mcpServer) startWarmWatcher(start time.Time, gotBytes *mcpByteFlag) (stop func()) {
	panic("fake")
}

type mcpByteFlag struct {
	mu sync.Mutex
	v  bool
}

func (f *mcpByteFlag) load() bool { panic("fake") }

func (f *mcpByteFlag) set() { panic("fake") }

func (s *mcpServer) callTool(raw json.RawMessage) (text string, err error) { panic("fake") }

func agentRunBody(args map[string]any, task string) map[string]any { panic("fake") }

func (s *mcpServer) submitChat(msg, model, chain, provider string) (string, int, error) {
	panic("fake")
}

// Fallback only when the coordinator refused SSE (or 4xx on the stream
// attempt). A successful empty stream must not fire a second completion.

func (s *mcpServer) agentRunStream(args map[string]any, task string) (string, int, error) {
	panic("fake")
}

// Non-stream JSON fallback when SSE returned a plain body.

func (s *mcpServer) agentJobStatus(id string) (string, int, error) { panic("fake") }

func (s *mcpServer) subagentStatus(id string) (string, int, error) { panic("fake") }

func (s *mcpServer) enrichChatPlanTrace(chatBody string) string { panic("fake") }

func (s *mcpServer) enrichAgentRunMeta(body string) (string, int, error) { panic("fake") }

func mcpChatTaskID(chatBody string) string { panic("fake") }

func (s *mcpServer) taskPlanTrace(taskID string) (plan, trace any) { panic("fake") }

func (s *mcpServer) harnessTracePath(taskID string) string { panic("fake") }

func (s *mcpServer) get(path string) (string, int, error) { panic("fake") }

func (s *mcpServer) post(path string, body any, timeout time.Duration) (string, int, error) {
	panic("fake")
}

func (s *mcpServer) do(method, path string, body io.Reader, timeout time.Duration) (string, int, error) {
	panic("fake")
}

func (s *mcpServer) postStream(path string, body any, timeout time.Duration, onEvent func(event, data string)) (string, int, error) {
	panic("fake")
}

func (s *mcpServer) postStreamHeaders(path string, body any, timeout time.Duration, onEvent func(event, data string)) (string, int, http.Header, error) {
	panic("fake")
}

func (s *mcpServer) postWithHeaders(path string, body any, timeout time.Duration) (string, int, error) {
	panic("fake")
}

func mcpUnknownJob(idKey, id string) string { panic("fake") }

func mcpChatCompletionJSON(id, model, content string) string { panic("fake") }

func truncate(s string, n int) string { panic("fake") }

// mcpUpstream pulls a worker/cache label from a coordinator JSON body when
// present (submit_chat served_by.primary). Blank for cluster_health etc.
func mcpUpstream(body string) string { panic("fake") }

// mcpTools returns the MCP tool definitions advertised in tools/list.
func mcpTools() []map[string]any { panic("fake") }

// Cursor rejects inputSchema.properties: null; empty object required for zero-arg tools.
