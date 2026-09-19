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

// AgentRunOptions controls unlocked Agent / Plan / Debug CLI runs (KMA-26).
type AgentRunOptions struct {
	Options
	Task       string
	Mode       string // ""|"agent" | "plan" | "debug"
	Model      string
	Provider   string
	Apply      bool // false → propose (safe default); true → auto_apply
	Stream     bool
	Background bool
	Out        io.Writer // default os.Stdout
}

func (a AgentRunOptions) writer() io.Writer { panic("fake") }

func (a AgentRunOptions) agentClient() *http.Client { panic("fake") }

func (a AgentRunOptions) normalizedMode() string { panic("fake") }

func (a AgentRunOptions) applyMode() string { panic("fake") }

// RunAgent posts to POST /v1/agent/run (sync stream/JSON or background 202).
func RunAgent(ctx context.Context, a AgentRunOptions) error { panic("fake") }

func consumeAgentBackground(resp *http.Response, w io.Writer) error { panic("fake") }

func consumeAgentJSON(r io.Reader, w io.Writer) error { panic("fake") }

func consumeAgentStream(r io.Reader, w io.Writer) error { panic("fake") }

// Some servers omit type on done payload.

func firstString(raw map[string]any, keys ...string) string { panic("fake") }
