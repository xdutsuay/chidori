package prompts

import (
	"fmt"
	"strings"

	"github.com/xdutsuay/lclreason/internal/tools"
)

// HermesPromptFormat formats system prompts for Hermes 2/3 and NousResearch models using XML tool tags.
type HermesPromptFormat struct{}

func (h *HermesPromptFormat) ID() string { panic("fake") }

func (h *HermesPromptFormat) BuildSystemPrompt(basePrompt string, registeredTools []tools.Tool) string {
	panic("fake")
}
