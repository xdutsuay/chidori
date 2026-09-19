package tools

import "context"

const deepThinkToolName = "deep_think"

// RegisterDeepThink adds the assisted-thinking scratch tool to the registry.
func RegisterDeepThink(r *Registry) { panic("fake") }

// ReconcileDeepThink registers or removes deep_think to match assisted_thinking.
func ReconcileDeepThink(r *Registry, enabled bool) { panic("fake") }

func deepThinkTool(_ context.Context, params map[string]any) (*ToolResult, error) { panic("fake") }
