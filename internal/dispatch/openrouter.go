package dispatch

import (
	"net/url"
	"strings"
)

// OpenRouterAutoModel is OpenRouter's Auto Router slug. Sending this as
// "model" lets OpenRouter classify the prompt and pick an upstream model
// (docs: https://openrouter.ai/docs/guides/routing/routers/auto-router).
// The response's "model" field reports which model actually ran.
const OpenRouterAutoModel = "openrouter/auto"

// IsOpenRouterRemote reports whether a remote is OpenRouter by vendor label
// or API host. Used to apply OpenRouter-only defaults (auto routing) without
// changing other OpenAI-compatible providers.
func IsOpenRouterRemote(provider, baseURL string) bool { panic("fake") }

// ResolveRemoteModel fills an empty model for providers that support routing.
// OpenRouter → openrouter/auto. Other vendors leave model unchanged (blank
// still fails requireModel — they need an explicit model).
func ResolveRemoteModel(provider, baseURL, model string) string { panic("fake") }
