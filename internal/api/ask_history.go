package api

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xdutsuay/lclreason/internal/agent"
	"github.com/xdutsuay/lclreason/internal/session"
)

var (
	askToolCardRe    = regexp.MustCompile(`(?is)<div[^>]*class="[^"]*tool-card[^"]*"[^>]*>(.*?)</div>`)
	askToolEmojiName = regexp.MustCompile(`(?s)^\s*🔧\s*([A-Za-z0-9_.\-]+)`)
	askToolResultPre = regexp.MustCompile(`(?is)<pre[^>]*data-tool-result="([^"]*)"[^>]*>(.*?)</pre>`)
)

// mergeSessionHistoryIntoAsk prepends prior same-session turns into req.Messages
// so Ask/chat completions see earlier agent work (KMA ask parity). Uses
// BuildOpenAIToolHistory for tool-card sequences recovered from stored HTML.
func mergeSessionHistoryIntoAsk(prior []session.Message, current []ChatMessage) []ChatMessage {
	panic("fake")
}

// Drop trailing prior user if it duplicates the incoming latest user content.

func lastUserContentChat(msgs []ChatMessage) string { panic("fake") }

func sessionMessagesToAskChat(msgs []session.Message) []ChatMessage { panic("fake") }

func chatMessageFromHistoryMap(m map[string]any) (ChatMessage, bool) { panic("fake") }

func splitAgentToolHTML(html string) (uses []agent.ToolCallRecord, results []agent.ToolResultRecord, remainder string) {
	panic("fake")
}
