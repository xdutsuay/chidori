package agent

import (
	"fmt"
	"strings"
)

const defaultSubagentMaxInflight = 2

// resolveSubagentModelChoice picks provider and model for a nested subagent
// from delegate_subtask params.model (KMA-107). Empty choice or "inherit"
// returns the parent's model and provider. Named choices must appear in
// cfg.SubagentModels and resolve via cfg.ResolveSubagentModel.
func resolveSubagentModelChoice(cfg Config, parentModel, parentProvider, choice string) (model, provider string, err error) {
	panic("fake")
}

func subagentModelAllowlisted(allow []string, choice string) bool { panic("fake") }

func effectiveSubagentMaxInflight(n int) int { panic("fake") }

func countAsyncSubagentInflight() int { panic("fake") }
