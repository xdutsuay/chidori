// Command fake-sync runs KMA-258 structure-only sync (local / manual).
//
//	go run ./cmd/fake-sync -source . -dest /path/to/chidori -dry-run
//	go run ./cmd/fake-sync -ui -source . -dest E:\codes\chidori
//
// -ui serves a temporary selection page (default :8765) to pick packages/files.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/xdutsuay/lclreason/internal/fakesync"
)

func main() { panic("fake") }

func splitCSV(s string) []string { panic("fake") }

func fatal(err error) { panic("fake") }
