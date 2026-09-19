package dispatch

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
)

// DefaultNVIDIAFreeModels is the fallback pool used when an NVIDIA remote
// returns 429 (KMA-114). Order is preference; the configured model is tried
// first, then these. Only used when isNVIDIARemote(cfg) is true.
// Pool pruned 2026-08-21 against integrate.api.nvidia.com: gemma-2-27b and
// llama-3.1-nemotron-70b returned 404 on free-tier keys; keep models that
// answered 200 in the corrected probe (plus glm which this install uses).
var DefaultNVIDIAFreeModels = []string{
	"meta/llama-3.3-70b-instruct",
	"meta/llama-3.1-8b-instruct",
	"z-ai/glm-5.2",
}

const nvidiaCooldown = 30 * time.Second

func isNVIDIARemote(cfg config.RemoteConfig) bool { panic("fake") }

func isRateLimitedErr(err error) bool { panic("fake") }

// isModelUnavailableErr is true when NIM rejects the model id (404 / not found).
// Treat like a soft miss so the free-tier pool can rotate (KMA-114).
func isModelUnavailableErr(err error) bool { panic("fake") }

func (r *Remote) nvidiaCooldownKey(remoteName, model string) string { panic("fake") }

func (r *Remote) nvidiaModelCooling(remoteName, model string, now time.Time) bool { panic("fake") }

func (r *Remote) markNVIDIACooldown(remoteName, model string, now time.Time) { panic("fake") }

// nvidiaCandidateModels returns preferred model then free-tier pool, skipping
// duplicates and models still in cooldown.
func (r *Remote) nvidiaCandidateModels(remoteName, preferred string, now time.Time) []string {
	panic("fake")
}

// If everything is cooling, still try preferred once (better than fail-closed).

func (r *Remote) invokeNVIDIAWithTools(ctx context.Context, rr resolvedRemote, model, prompt string, tools []ChatTool) (InvokeResult, error) {
	panic("fake")
}

func (r *Remote) invokeNVIDIAStream(ctx context.Context, rr resolvedRemote, model, prompt string, cb func(string)) (InvokeResult, error) {
	panic("fake")
}
