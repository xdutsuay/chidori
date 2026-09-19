package dispatch

import "strings"

// Must match agent.IncrementalPromptSplit — duplicated to avoid an import cycle
// (agent -> dispatch). Keep in sync when changing the marker (KMA-163).
const incrementalPromptSplit = "\n<<<CHIDORI_INCREMENTAL>>>\n"

func flattenIncrementalPrompt(prompt string) string { panic("fake") }

// chatMessagesFromIncrementalPrompt turns a joined agent prompt into OpenAI
// chat messages: system=stable, user=volatile when the split marker is present.
func chatMessagesFromIncrementalPrompt(prompt string) (msgs []chatMessage, flat string) {
	panic("fake")
}
