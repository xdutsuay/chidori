package compact

import (
	"strings"
)

// TruncateLines keeps at most maxLines from the end of prefix and the start of suffix.
func TruncateLines(prefix, suffix string, maxLines int) (string, string) { panic("fake") }

// CompactAgentScratch trims tool observation history to fit a token budget.
// Keeps the most recent observations verbatim; older ones collapse to one-line summaries.
func CompactAgentScratch(observations []string, budgetTokens int) string { panic("fake") }

// BuildCodeSystemPrompt assembles a minimal system prompt for code completion.
func BuildCodeSystemPrompt(language, path string) string { panic("fake") }
