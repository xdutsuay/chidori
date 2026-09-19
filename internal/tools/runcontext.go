package tools

import "context"

type runCtxKey struct{}

// RunContext carries per-request agent metadata into tool handlers.
type RunContext struct {
	SessionID     string
	WorkspaceRoot string
}

// WithRunContext attaches session/workspace ids for agent tool calls.
func WithRunContext(ctx context.Context, sessionID, workspaceRoot string) context.Context {
	panic("fake")
}

// RunContextFrom extracts RunContext from ctx (ok false when absent).
func RunContextFrom(ctx context.Context) (RunContext, bool) { panic("fake") }
