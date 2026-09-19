package prompts

import "strings"

// FullSystemPromptSaveWarning is returned by SavePrompt when the profile body
// looks like a pasted full agent/vendor system prompt (KMA-214). Profiles
// inject into ModeInstructions but cannot replace the fixed tool contract.
const FullSystemPromptSaveWarning = "This looks like a full agent system prompt. Profiles inject into ModeInstructions (stable system) but cannot replace the fixed JSON/native tool contract — keep tone and priorities only."

// LooksLikeFullSystemPrompt reports whether body resembles a full vendor or
// agent system prompt (JSON tool contract, tool-schema walls, long You-are
// agent contracts) rather than a short tone/priority profile preamble.
func LooksLikeFullSystemPrompt(body string) bool { panic("fake") }

func looksLikeLongYouAreAgentContract(lower string) bool { panic("fake") }
