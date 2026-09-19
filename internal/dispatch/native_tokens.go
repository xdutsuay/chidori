package dispatch

// chatUsageBlock is the usage object on OpenAI-compatible chat responses.
type chatUsageBlock struct {
	PromptTokens           int `json:"prompt_tokens"`
	CompletionTokens       int `json:"completion_tokens"`
	NativeTokensPrompt     int `json:"native_tokens_prompt"`
	NativeTokensCompletion int `json:"native_tokens_completion"`
	NativeTokensReasoning  int `json:"native_tokens_reasoning"`
	PromptTokensDetails    struct {
		CachedTokens int `json:"cached_tokens"`
	} `json:"prompt_tokens_details"`
}

// preferNativeTokenCounts prefers OpenRouter native_tokens_* when present and
// larger than the generic prompt_tokens/completion_tokens fields (KMA-212).
func preferNativeTokenCounts(u chatUsageBlock) (prompt, completion, reasoning int) { panic("fake") }
