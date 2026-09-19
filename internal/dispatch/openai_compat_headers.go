package dispatch

import (
	"net/http"
	"net/url"
	"strings"
)

// OpenAICompatMeta identifies an OpenAI-compatible provider for optional
// per-vendor HTTP headers (OpenRouter attribution, future adapters).
type OpenAICompatMeta struct {
	BaseURL string // e.g. https://openrouter.ai/api/v1
	Vendor  string // RemoteConfig.Provider label, e.g. "openrouter"
}

// openAICompatProfile holds optional extra request headers beyond Bearer auth.
type openAICompatProfile struct {
	ByVendor map[string]map[string]string
	ByHost   map[string]map[string]string
}

// chidoriAppReferer is sent as HTTP-Referer for providers that require app
// attribution (OpenRouter free/auto routes reject requests without it).
const chidoriAppReferer = "https://github.com/xdutsuay/lclreason"

var openAICompatProfiles = openAICompatProfile{
	ByVendor: map[string]map[string]string{
		"openrouter": {
			"HTTP-Referer": chidoriAppReferer,
			"X-Title":      "chidori",
		},
	},
	ByHost: map[string]map[string]string{
		"openrouter.ai": {
			"HTTP-Referer": chidoriAppReferer,
			"X-Title":      "chidori",
		},
	},
}

// OpenAICompatExtraHeaders returns optional HTTP headers for a provider.
// Vendor label wins over base-URL host matching so custom base URLs still work
// when the user picks a preset.
func OpenAICompatExtraHeaders(baseURL, vendor string) map[string]string { panic("fake") }

// Subdomain fallback (api.openrouter.ai → openrouter.ai).

func hostFromBaseURL(baseURL string) string { panic("fake") }

func cloneHeaderMap(in map[string]string) map[string]string { panic("fake") }

// SetOpenAICompatRequestHeaders sets Authorization (Bearer apiKey) and any
// provider-specific headers on req.
func SetOpenAICompatRequestHeaders(req *http.Request, meta OpenAICompatMeta, apiKey string) {
	panic("fake")
}

// SetOpenAICompatAuthHeader sets Authorization from a pre-built header value
// (e.g. "Bearer sk-…") plus provider-specific extras.
func SetOpenAICompatAuthHeader(req *http.Request, meta OpenAICompatMeta, authHeader string) {
	panic("fake")
}
