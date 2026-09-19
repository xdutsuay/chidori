package tools

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/xdutsuay/lclreason/internal/memory"
)

// RegisterBuiltins adds the default tool set to the registry.
// mem may be nil if memory is disabled.
// shellCfg may be nil for defaults; set shellCfg.Enabled=false to skip shell.
func RegisterBuiltins(r *Registry, mem *memory.VectorStore, shellCfg *ShellConfig) { panic("fake") }

// Models often emit "print" (Python-style); alias to echo so the loop doesn't fail.

func calculatorTool(_ context.Context, params map[string]any) (*ToolResult, error) { panic("fake") }

// Very simple evaluator: supports a op b where op is +, -, *, /.

func echoTool(_ context.Context, params map[string]any) (*ToolResult, error) { panic("fake") }

func printTool(_ context.Context, params map[string]any) (*ToolResult, error) { panic("fake") }

func memorySearchTool(mem *memory.VectorStore) ToolFunc { panic("fake") }

func memoryAddTool(mem *memory.VectorStore) ToolFunc { panic("fake") }
