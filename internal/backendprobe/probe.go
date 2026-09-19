// Package backendprobe builds HTTP GET URLs for LLM backend reachability checks
// without hitting invalid endpoints (e.g. bare GET /v1 on LM Studio).
package backendprobe

import "strings"

// ReachabilityURL returns a GET endpoint suitable for probing whether a backend
// is up. OpenAI-compatible servers (LM Studio, AirLLM) reject bare GET /v1;
// use GET /v1/models instead. Ollama uses GET /api/tags.
func ReachabilityURL(baseURL, backend string) string { panic("fake") }

// ollama and unknown

// BackendReachabilityURL picks the configured backend URL from provider config
// and returns the probe endpoint for reachability checks.
func BackendReachabilityURL(backend, ollamaURL, lmStudioURL, airLLMURL string) string { panic("fake") }

// NodeReachabilityURL returns a probe URL for a registered worker node.
func NodeReachabilityURL(apiBase, address, backend string) string { panic("fake") }
