package prompts

import (
	"strings"

	"github.com/xdutsuay/lclreason/internal/tools"
)

// PromptFormat defines an interface for model-family-specific prompt formatting.
type PromptFormat interface {
	ID() string
	BuildSystemPrompt(basePrompt string, registeredTools []tools.Tool) string
}

// DefaultPromptFormat provides standard chidori JSON plan formatting.
type DefaultPromptFormat struct{}

func (d *DefaultPromptFormat) ID() string { panic("fake") }

func (d *DefaultPromptFormat) BuildSystemPrompt(basePrompt string, registeredTools []tools.Tool) string {
	panic("fake")
}

// GetFormat returns the prompt formatter for the given format ID.
func GetFormat(formatID string) PromptFormat { panic("fake") }
