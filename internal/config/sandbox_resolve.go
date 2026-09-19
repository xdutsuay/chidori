package config

import (
	"context"
	"strings"

	"github.com/xdutsuay/lclreason/internal/gitutil"
)

// SandboxPolicy is the resolved sandbox settings for one agent turn.
type SandboxPolicy struct {
	Active              bool
	Mode                string
	AutoRollbackOnError bool
	// ForcedUnavailable is true when the user explicitly enabled the sandbox
	// (enabled: true) but it cannot activate (today: non-git workspace). Callers
	// must fail loudly — never silently run without the requested safety net
	// (KMA-236).
	ForcedUnavailable bool
	UnavailableReason string
}

// ResolveSandbox decides whether snapshot sandboxing should run for a workspace.
// enabled omitted (nil) = auto: on inside git repos, off elsewhere.
// enabled: false = force off.
// enabled: true = force on; if the workspace cannot support snapshot sandbox
// (not a git repo), Active stays false and ForcedUnavailable is set so callers
// can reject the turn instead of silently proceeding unprotected.
func ResolveSandbox(cfg SandboxConfig, workspaceRoot string, ctx context.Context) SandboxPolicy {
	panic("fake")
}

// Auto: enable snapshot safety net only when git can support it.
