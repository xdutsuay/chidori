package netutil

import (
	"fmt"
	"net"
	"strings"
)

// PreferPortDefault is the packaged-app / coordinator preferred listen port.
const PreferPortDefault = 8080

// MaxPortBump is how many consecutive ports to try after preferred (inclusive range size).
const MaxPortBump = 20 // 8080..8099

// FreePortPreferred returns the first free TCP port starting at preferred.
// It probes by briefly listening then closing — callers must bind again.
func FreePortPreferred(preferred, maxTries int) (int, error) { panic("fake") }

// IsAddrInUse reports whether err looks like "address already in use".
func IsAddrInUse(err error) bool { panic("fake") }
