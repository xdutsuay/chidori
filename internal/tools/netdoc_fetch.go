package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/xdutsuay/lclreason/internal/netdoc"
)

const browseFetchTool = "browse_fetch"

// RegisterNetdocTools adds experimental read-only HTTP fetch tools (P3).
func RegisterNetdocTools(r *Registry, cache *netdoc.Cache) { panic("fake") }

// ReconcileNetdocTools registers or removes browse_fetch to match enabled.
func ReconcileNetdocTools(r *Registry, enabled bool, cache *netdoc.Cache) { panic("fake") }

func netdocFetchTool(cache *netdoc.Cache) ToolFunc { panic("fake") }
