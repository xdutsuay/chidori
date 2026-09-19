package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
	"github.com/xdutsuay/lclreason/internal/coordinator"
	"github.com/xdutsuay/lclreason/internal/devlog"
	"github.com/xdutsuay/lclreason/internal/netutil"
)

// version is stamped at build time via -ldflags "-X main.version=...".
var version = "dev"

func main() { panic("fake") }

// Prefer config/default port; bump if busy so a stray python -m http.server
// on 8080 cannot steal the coordinator.
//
// mcp mode never binds a socket — it speaks JSON-RPC over stdio and dials
// the coordinator instead. Reserving/bumping a port here causes it to
// dial a port nothing is listening on (KMA-88). mcp.go builds its base
// URL from cfg.Coordinator or cfg.Port as-is.

func runWorker(cfg *config.Config) { panic("fake") }

// Start the inference proxy so the coordinator can route requests here.
