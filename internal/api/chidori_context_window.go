package api

import (
	"net/http"
	"strings"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/compact"
	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
)

// UIContextWindowBreakdown reports an estimated breakdown of what's
// currently occupying the active context window — messages, the agent's
// static system prompt, active rules, and tool definitions — against a
// dynamic budget (AI.27). See resolveContextBudget for how model/auto +
// provider resolve to a window.
func (h *Handlers) UIContextWindowBreakdown(w http.ResponseWriter, r *http.Request) { panic("fake") }

// lookupRemoteModel returns the configured model for a saved remote provider
// name (topbar value like "nvidia-7030"), or "" if unknown / no pool.
func (h *Handlers) lookupRemoteModel(providerName string) string { panic("fake") }

// resolveContextBudget picks the token budget shown in the AI.27 gauge.
//
// Order (matches why "nvidia-7030 + auto" used to show 8192: the topbar
// "auto" option has value="" so model was empty and we fell through to
// compaction.context_tokens):
//  1. Explicit model id → dispatch.ContextWindow(model)
//  2. Empty / "auto" model + provider's configured remote Model → that window
//  3. Empty / "auto" model + provider name heuristic (e.g. "nvidia-…") → window
//  4. compaction.context_tokens, else ContextWindow("")
func resolveContextBudget(cfg *config.Config, model, provider, providerModel string) (budget int, source, resolvedModel string) {
	panic("fake")
}
