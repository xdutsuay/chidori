package api

import (
	"strings"

	"github.com/xdutsuay/lclreason/internal/config"
)

// applyAgentSubagentConfig copies KMA-107 allowlist / inflight settings from
// config.yaml agent: into an agent.Config (and installs a fail-closed resolver).
func (h *Handlers) agentSubagentModels() []string { panic("fake") }

func (h *Handlers) agentSubagentMaxInflight() int { panic("fake") }

// resolveConfiguredSubagentModel maps an allowlisted slug to a configured
// inference source (remote name/model or local provider.model / ui default).
func (h *Handlers) resolveConfiguredSubagentModel(name string) (provider, model string, ok bool) {
	panic("fake")
}

func resolveSubagentModelFromConfig(cfg *config.Config, name string) (provider, model string, ok bool) {
	panic("fake")
}
