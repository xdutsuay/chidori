package dispatch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// verifyTimeout bounds how long a single Verify probe waits for the
// provider before giving up.
//
// Reduced from 20s (running_issue.md P1-1 / H.4): a valid key against a
// slow/cold-starting model (common on NVIDIA NIM's free tier, and
// structurally guaranteed for reasoning-class models that spend time
// thinking before any token comes back) used to time out at 20s and get
// reported as if the key itself were rejected — there was no way to tell
// "the model didn't answer in time" apart from "the key was revoked" from
// the message alone. 8s keeps the UI from feeling frozen while still giving
// a fast provider plenty of room; verifyViaChat below now also words a
// timeout distinctly from a real 401/403 rejection so the two stop being
// conflated. Package-level (not a literal in the function) so tests can
// shrink it to exercise the timeout path without a real multi-second sleep.
var verifyTimeout = 8 * time.Second

// VerifyOpenAICompat checks whether apiKey is actually accepted by the
// OpenAI-compatible provider at baseURL.
//
// When a model is known, the probe is a live 1-token POST /chat/completions —
// the only request class these providers uniformly authenticate. A plain GET
// /models is NOT sufficient: some providers (NVIDIA NIM notably) serve their
// model catalog publicly and return 200 no matter what key you send, which
// made "Verify" report success for keys that had already been revoked.
//
// Without a model (the Add-form "Test" button before a model is picked), it
// falls back to GET /models — but first sends the same request with a
// deliberately bogus key. If the bogus key also gets 200, the endpoint isn't
// authenticating at all and we say so instead of pretending the key checked
// out.
func VerifyOpenAICompat(ctx context.Context, client *http.Client, baseURL, apiKey, model, vendor string) (bool, string) {
	panic("fake")
}

// isTimeoutErr reports whether err represents this probe's own deadline
// firing (context cancellation or an http.Client's Timeout field) rather
// than a real network failure — the distinction verifyViaChat needs to word
// a timeout as "inconclusive", not "rejected".
func isTimeoutErr(ctx context.Context, err error) bool { panic("fake") }

func verifyViaChat(ctx context.Context, client *http.Client, meta OpenAICompatMeta, apiKey, model string) (bool, string) {
	panic("fake")
}

// Auth is checked before the request body on every provider we
// support, so reaching a 400/404 means the key itself authenticated —
// the model name (or a request param) is what got rejected.

func verifyViaModels(ctx context.Context, client *http.Client, meta OpenAICompatMeta, apiKey string) (bool, string) {
	panic("fake")
}

// 200 with the real key. Only meaningful if the endpoint rejects a bogus
// key — otherwise /models is public and proves nothing about our key.

func getModelsStatus(ctx context.Context, client *http.Client, meta OpenAICompatMeta, apiKey string) (int, error) {
	panic("fake")
}

func readErrorSnippet(r io.Reader) string { panic("fake") }
