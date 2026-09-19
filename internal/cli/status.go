package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/xdutsuay/lclreason/internal/config"
)

// WorkspaceStatus prints workspace root/trust (live GET /api/workspace/info, else config file).
func WorkspaceStatus(ctx context.Context, o Options) (string, error) { panic("fake") }

// CompanionStatus prints companion listen state (live GET /api/companion only).
// Read-only — does not call pair/scan. File fallback reports configured port only.
func CompanionStatus(ctx context.Context, o Options) (string, error) { panic("fake") }

func getJSON(ctx context.Context, o Options, cfg *config.Config, path string) (map[string]any, error) {
	panic("fake")
}

func writeMapLines(b *strings.Builder, raw map[string]any, keys []string) { panic("fake") }
