package api

import (
	"net/http"
	"strings"
	"sync"

	"github.com/xdutsuay/lclreason/internal/netdoc"
	"github.com/xdutsuay/lclreason/internal/tools"
)

var (
	netdocMu     sync.Mutex
	netdocCache  *netdoc.Cache
	netdocClient *netdoc.Client
)

func netdocCacheForHandlers() (*netdoc.Cache, error) { panic("fake") }

// NetdocFetch returns extracted page text for @web mentions (experimental, SSRF-hardened).
func (h *Handlers) NetdocFetch(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) syncNetdocTools() { panic("fake") }
