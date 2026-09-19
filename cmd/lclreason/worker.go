package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/dispatch"
	"github.com/xdutsuay/lclreason/internal/registry"
)

type worker struct {
	cfg    *config.Config
	client *http.Client
	id     string
	models []string
}

// detectModels queries the local Ollama instance for available models.
// Falls back to a sensible default if Ollama is unreachable.
func (w *worker) detectModels() { panic("fake") }

func (w *worker) register() error { panic("fake") }

// Detect models before registering so the coordinator knows what
// this worker can serve.

// Address the coordinator will dispatch inference to. Defaults to
// hostname:port (works on a LAN if the hostname resolves via mDNS), but can
// be overridden with `advertise: <ip>:<port>` for a guaranteed-reachable IP.

func (w *worker) heartbeat() error { panic("fake") }

// A 200 with a body carries a control command (e.g. restart); 204 means
// nothing to do.

// prewarm sends a tiny generation request to the local Ollama so the model is
// loaded into memory before the first real request (honors prewarm_model).
func (w *worker) prewarm() { panic("fake") }

func (w *worker) deregister() error { panic("fake") }

func (w *worker) post(path string, body any) error { panic("fake") }

// serveInference starts an HTTP server on the worker that proxies
// inference requests to the local Ollama instance. This lets the
// coordinator route requests to this worker's address. Without this
// the coordinator can't reach a worker that's behind NAT — the worker
// exposes its own /api/generate endpoint.
func (w *worker) serveInference() *http.Server { panic("fake") }

// Health endpoint for the coordinator's HealthCheck probe.

// Proxy /api/generate to local Ollama.

// Stream the response through so SSE works.

// Proxy /api/tags so coordinator can query models.
