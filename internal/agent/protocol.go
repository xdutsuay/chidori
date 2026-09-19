package agent

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/tools"
)

// DecodeTurn normalizes a model turn into planResponse.
// protocol is config.ToolProtocolJSON or config.ToolProtocolOpenAI.
// For openai, prefer structured toolCalls from the dispatch result; raw may
// still hold assistant content used as Summary when Done.
func DecodeTurn(protocol, raw string, toolCalls []dispatch.ToolCall) (planResponse, error) {
	panic("fake")
}

func decodeOpenAITurn(raw string, toolCalls []dispatch.ToolCall) (planResponse, error) { panic("fake") }

// Some remotes ignore tools[] and still emit our JSON plan in content —
// accept that as a soft fallback so the execute loop still runs.

// ToolSchemasFromRegistry builds OpenAI tools[] from registry descriptions.
// Parameters come from tools.OpenAIParameters so required keys are advertised
// (empty object schemas caused remotes to call with {} and fail validation).
func ToolSchemasFromRegistry(describe []tools.Tool) []dispatch.ChatTool { panic("fake") }

// promptHiddenTools are registered but deliberately not advertised: echo and
// print are model-mistake aliases kept so a stray call doesn't fail the turn,
// and advertising them would invite non-coding chatter turns.
var promptHiddenTools = map[string]bool{"echo": true, "print": true}

// advertisableTool is the single advertisement policy shared by the OpenAI
// tools[] path (filterToolSchemas) and the JSON prompt path
// (promptToolCatalog): everything actually registered is advertised, minus
// hidden aliases, minus the code-execution tools guardedRunner would reject
// for this trust/browser configuration — advertising a tool whose every call
// fails is exactly the "tool exists but is useless" loop this avoids.
func advertisableTool(name string, trusted, browserTools bool) bool { panic("fake") }

// filterToolSchemas applies the advertisement policy to registry-derived
// OpenAI schemas, then appends the loop-internal tools (delegate_subtask,
// subagent_result) that live in the agent loop rather than the tool registry
// — without that, the OpenAI native path never offers them even though the
// prompt text mentions them.
func filterToolSchemas(all []dispatch.ChatTool, trusted, browserTools bool) []dispatch.ChatTool {
	panic("fake")
}

// deferredToolThreshold: advertise full catalog when len < this; at/above,
// keep core tools + tool_search (KMA-86 / KMA-100). Was <= which left a live
// catalog of exactly 16 never deferred.
const deferredToolThreshold = 16

var deferredCoreTools = map[string]bool{
	"read_file": true, "write_file": true, "apply_patch": true,
	"grep_workspace": true, "list_dir": true,
	"shell_exec": true, "git_commit": true,
	"delegate_subtask": true, "subagent_result": true,
	"tool_search":          true,
	"search_chat_messages": true, "get_chat_attachment": true,
	"web_search": true, "fetch_url": true,
	"browse_navigate": true, "browse_snapshot": true, "browse_click": true,
}

func deferToolSchemasIfBloated(all []dispatch.ChatTool) []dispatch.ChatTool { panic("fake") }

// promptToolCatalog builds the advertised tool list for the JSON prompt path
// from the live registry catalog: same policy as filterToolSchemas, sorted by
// name (stable prompt prefix), with the loop-internal tools appended. Returns
// nil when the caller has no catalog so buildAgentPrompt can fall back to its
// legacy hardcoded list (tests, older callers).
func promptToolCatalog(catalog []tools.Tool, trusted, browserTools bool) []tools.Tool { panic("fake") }

// loopInternalToolSchemas describes the tools dispatched inside Loop.run
// itself (never via the registry), so the OpenAI tools[] advertisement
// matches what the loop actually accepts.
func loopInternalToolSchemas() []dispatch.ChatTool { panic("fake") }

func openAILoopToolParameters(required string, properties map[string]any) map[string]any {
	panic("fake")
}
