package tools

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

// ToolFunc is the signature for an executable tool.
type ToolFunc func(ctx context.Context, params map[string]any) (*ToolResult, error)

// ProposedChange describes a file edit preview (propose mode).
type ProposedChange struct {
	Path   string `json:"path"`
	Kind   string `json:"kind"` // write_file | apply_patch
	Before string `json:"before"`
	After  string `json:"after"`
}

// ToolResult carries the output of a tool execution.
type ToolResult struct {
	Success  bool            `json:"success"`
	Output   string          `json:"output"`
	Error    string          `json:"error,omitempty"`
	Proposed *ProposedChange `json:"proposed,omitempty"`
}

// Tool describes a registered tool.
type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	fn          ToolFunc
}

// Registry holds named tools that the planner can reference in its steps.
type Registry struct {
	mu    sync.RWMutex
	tools map[string]Tool
	calls map[string]*atomic.Int64
}

// NewRegistry creates a registry with no built-in tools.
// Call Register to add tools.
func NewRegistry() *Registry { panic("fake") }

// Register adds a tool. Overwrites if the name already exists.
func (r *Registry) Register(name, description string, fn ToolFunc) { panic("fake") }

// Unregister removes a tool by name. No-op when absent.
func (r *Registry) Unregister(name string) { panic("fake") }

// requiredParams lists the params each builtin tool cannot run without —
// checked centrally in Execute so a malformed call fails fast with a
// corrective message (naming the missing param and the tool's real usage)
// instead of bubbling whatever confusing error the tool body produces.
// Tools absent from this table (or with alternative-param shapes, like
// browse_click's selector-OR-text) validate themselves.
// Also feeds OpenAIParameters so native tool_calls advertise real required keys
// (empty schemas caused models to call with {}).
var requiredParams = map[string][]string{
	"read_file":       {"path"},
	"write_file":      {"path", "content"},
	"apply_patch":     {"path", "old"},
	"grep_workspace":  {"pattern"},
	"git_commit":      {"message"},
	"shell_exec":      {"command"},
	"web_search":      {"query"},
	"fetch_url":       {"url"},
	"browse_navigate": {"url"},
}

// schemaRequiredExtra keys are advertised as required in OpenAI tools[] but
// validated inside the tool body (keep Execute error strings stable).
var schemaRequiredExtra = map[string][]string{
	"search_chat_messages": {"query"},
	"get_chat_attachment":  {"path"},
}

// optionalParams are advertised in OpenAI tools[] properties but not enforced
// by missingRequiredParams (path on list_dir, line ranges, etc.).
var optionalParams = map[string][]string{
	"read_file":            {"start_line", "end_line"},
	"write_file":           {"append"},
	"apply_patch":          {"new"},
	"grep_workspace":       {"glob"},
	"list_dir":             {"path"},
	"web_search":           {"max_results"},
	"browse_snapshot":      {"selector"},
	"browse_click":         {"selector", "text"},
	"search_chat_messages": {"scope", "limit"},
}

// RequiredParams returns a copy of the centrally required keys for name, or nil.
func RequiredParams(name string) []string { panic("fake") }

func advertisedRequired(name string) []string { panic("fake") }

// OpenAIParameters builds a JSON-Schema-ish parameters object for tools[] ads.
// Required keys come from requiredParams (+ schemaRequiredExtra); optional
// known fields from optionalParams. additionalProperties stays true.
func OpenAIParameters(name string) map[string]any { panic("fake") }

// missingRequiredParams returns the required params absent (or blank strings)
// in the call. Non-string values count as present.
func missingRequiredParams(name string, params map[string]any) []string { panic("fake") }

// Execute looks up a tool by name, validates its required params, and runs it.
func (r *Registry) Execute(ctx context.Context, name string, params map[string]any) (*ToolResult, error) {
	panic("fake")
}

// Include the live tool list so a model that hallucinated a name can
// self-correct on its next turn instead of retrying blind.

// List returns all registered tool names, sorted.
func (r *Registry) List() []string { panic("fake") }

// Describe returns all tools with their descriptions, sorted by name.
func (r *Registry) Describe() []Tool { panic("fake") }

// Stats returns execution counts per tool.
func (r *Registry) Stats() map[string]int64 { panic("fake") }

// Has reports whether a named tool is registered.
func (r *Registry) Has(name string) bool { panic("fake") }
