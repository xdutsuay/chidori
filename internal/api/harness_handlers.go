package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/xdutsuay/lclreason/internal/harnesstrace"
	"github.com/xdutsuay/lclreason/internal/replay"
)

// HarnessEvents streams live harness trace events over SSE when the experimental
// visualizer flag is on (HV.1). Returns 404 when disabled.
func (h *Handlers) HarnessEvents(w http.ResponseWriter, r *http.Request) { panic("fake") }

// HarnessTracesList returns persisted JSONL trace files (HV.3).
func (h *Handlers) HarnessTracesList(w http.ResponseWriter, r *http.Request) { panic("fake") }

// HarnessTracesGet returns decoded events for one trace id (HV.3 replay).
func (h *Handlers) HarnessTracesGet(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) harnessVisualizerEnabled() bool { panic("fake") }

func (h *Handlers) harnessTraceStore() *harnesstrace.FileStore { panic("fake") }

func (h *Handlers) appendHarnessTrace(ev harnesstrace.Event) { panic("fake") }

func traceIDFromPayload(p map[string]any) string { panic("fake") }

func (h *Handlers) syncHarnessTrace() { panic("fake") }

var (
	harnessPersistMu   sync.Mutex
	harnessPersistStop func()
)

func (h *Handlers) syncHarnessTraceSubscriber(on bool) { panic("fake") }

func (h *Handlers) syncAgentJobsPersist() { panic("fake") }
