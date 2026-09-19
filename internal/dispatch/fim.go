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

	"github.com/xdutsuay/lclreason/internal/registry"
)

// ollamaFIMRequest is Ollama /api/generate with optional suffix (FIM).
type ollamaFIMRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Suffix string `json:"suffix,omitempty"`
	Stream bool   `json:"stream"`
}

// InvokeFIM attempts fill-in-the-middle via Ollama suffix support.
// Returns an error if the backend rejects FIM so callers can fall back.
func (d *Dispatcher) InvokeFIM(ctx context.Context, model, prefix, suffix string) (InvokeResult, error) {
	panic("fake")
}

func (d *Dispatcher) invokeOllamaFIM(ctx context.Context, node *registry.Node, model, prefix, suffix string) (InvokeResult, error) {
	panic("fake")
}

// InvokePrefixCompletion uses a minimal chat prompt for next-token completion.
func (d *Dispatcher) InvokePrefixCompletion(ctx context.Context, model, system, prefix string) (InvokeResult, error) {
	panic("fake")
}

// ExecuteCodeCompletion tries FIM when suffix is set, else prefix-only.
func (p *Pool) ExecuteCodeCompletion(ctx context.Context, model, system, prefix, suffix, providerHint string) (InvokeResult, string, error) {
	panic("fake")
}

// CleanCompletion strips markdown fences and leading/trailing whitespace from model output.
func CleanCompletion(s string) string { panic("fake") }
