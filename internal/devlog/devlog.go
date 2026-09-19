// Package devlog tees the standard library log package to a rotating file
// for local development. Off by default — shipping builds stay quiet unless
// the developer sets CHIDORI_DEV_LOG=1.
//
// This does not auto-instrument every Go function; it captures every
// log.Printf / log.Print / log.Fatal / log.Panic (and anything else that
// writes through log.Default()), plus flightlog lines mirrored when enabled.
//
// Usage:
//
//	CHIDORI_DEV_LOG=1 ./cmd/lclreason-app/build/bin/chidori.app/Contents/MacOS/chidori
//	tail -f ~/Library/Application\ Support/lclreason/logs/chidori-dev.log
package devlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	envEnable = "CHIDORI_DEV_LOG"
	envPath   = "CHIDORI_DEV_LOG_PATH"
)

var (
	mu      sync.Mutex
	enabled bool
	path    string
	orig    io.Writer
)

// Enabled reports whether Start successfully turned on the rotating sink.
func Enabled() bool { panic("fake") }

// Path returns the active log file path, or "" if disabled.
func Path() string { panic("fake") }

// Wanted reports whether the env var asks for dev logging (even before Start).
func Wanted() bool { panic("fake") }

func wanted(force bool) bool { panic("fake") }

// DefaultPath picks the rotating log location.
// Priority: CHIDORI_DEV_LOG_PATH → <dataDir>/logs/chidori-dev.log → ./logs/chidori-dev.log.
func DefaultPath(dataDir string) string { panic("fake") }

// Start enables the rotating sink when CHIDORI_DEV_LOG is on or force=true.
// Safe to call multiple times.
// dataDir should be the Application Support / config data directory when known.
// Returns the path used, or "" when disabled / on failure.
func Start(dataDir string, force ...bool) (string, error) { panic("fake") }

// MiB

// days

// Log after SetOutput so this line lands in the file too.

// Mirror writes a one-line summary into the standard logger when enabled.
// Used by flightlog so agent/chat events appear in the same tailed file.
func Mirror(lvl, comp, op string, fields map[string]any) { panic("fake") }
