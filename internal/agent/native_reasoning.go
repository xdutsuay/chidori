package agent

import (
	"strings"

	"github.com/xdutsuay/lclreason/internal/dispatch"
)

// thoughtFromNativeInvokeResult extracts provider chain-of-thought for the
// Agent UI thought card (KMA-204). Ask path already surfaces InvokeResult.Reasoning;
// the native tools path previously dropped it.
func thoughtFromNativeInvokeResult(res dispatch.InvokeResult) string { panic("fake") }
