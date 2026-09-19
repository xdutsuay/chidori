package api

import (
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/xdutsuay/lclreason/internal/acpclient"
	"github.com/xdutsuay/lclreason/internal/harness/external"
)

// ExternalHarnessStatus reports install/auth hints for the experimental Grok or Hermes harness.
func (h *Handlers) ExternalHarnessStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }
