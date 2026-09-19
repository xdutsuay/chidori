package dispatch

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/registry"
)

// ErrVisionModelNotConfigured is returned when remote vision is requested but
// no model id is configured (provider.vision_model or remote.vision_model).
var ErrVisionModelNotConfigured = fmt.Errorf("vision model not configured — set provider.vision_model in config.yaml or Settings → Inference Source")

// ResolveVisionModel picks the vision model id from config only — never a
// hardcoded vendor-specific default.
func ResolveVisionModel(globalModel string, rc config.RemoteConfig) (string, error) { panic("fake") }

// PickVisionRemote selects the remote name + configured vision model for inference.
func (r *Remote) PickVisionRemote() (name, model string, ok bool) { panic("fake") }

// SetVisionConfig updates global vision routing (call when provider config reloads).
func (r *Remote) SetVisionConfig(model, remote string) { panic("fake") }

// ImagePart is one image attached to a multimodal chat turn.
type ImagePart struct {
	MIME string // e.g. image/png
	Data []byte // raw bytes
}

// InvokeVision sends text + images as OpenAI multimodal content parts.
func (r *Remote) InvokeVision(ctx context.Context, name, model, prompt string, images []ImagePart) (InvokeResult, error) {
	panic("fake")
}

func invokeChatHTTPVision(ctx context.Context, client *http.Client, baseURL, authHeader, model, prompt string, images []ImagePart, start time.Time, nodeID, providerLabel, vendor string) (InvokeResult, error) {
	panic("fake")
}

// invokeNodeVision sends multimodal content to a local OpenAI-compatible backend.
func invokeNodeVision(ctx context.Context, client *http.Client, node *registry.Node, model, prompt string, images []ImagePart) (InvokeResult, error) {
	panic("fake")
}

func isVisionUnsupportedErr(err error) bool { panic("fake") }
