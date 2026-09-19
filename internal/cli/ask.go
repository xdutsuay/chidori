package cli

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
)

// AskOptions controls Phase 3 `chidori ask` (Ask chat only — no Agent tools).
type AskOptions struct {
	Options
	Prompt   string
	Model    string
	Provider string
	Stream   bool
	Out      io.Writer // default os.Stdout
}

func (a AskOptions) writer() io.Writer { panic("fake") }

func (a AskOptions) askClient() *http.Client { panic("fake") }

// Ask runs a single Ask-mode chat completion against a live coordinator.
func Ask(ctx context.Context, a AskOptions) error { panic("fake") }

func consumeAskJSON(r io.Reader, w io.Writer) error { panic("fake") }

func consumeAskStream(r io.Reader, w io.Writer) error { panic("fake") }

// Model chunks can be large; raise buffer.

func extractMessageContent(raw map[string]any) string { panic("fake") }

func extractDeltaContent(raw map[string]any) string { panic("fake") }
