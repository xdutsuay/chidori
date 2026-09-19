package acpclient

import (
	"context"
	"strings"
	"sync"
	"time"
)

const (
	streamSettleInterval = 150 * time.Millisecond
	streamSettleChecks   = 2
)

type promptStreamState struct {
	mu           sync.Mutex
	acc          strings.Builder
	lastActivity time.Time
	sawDone      bool
}

func (ps *promptStreamState) noteChunk(chunk string) { panic("fake") }

func (ps *promptStreamState) noteDone() { panic("fake") }

func (ps *promptStreamState) snapshot() (text string, sawDone bool, lastActivity time.Time) {
	panic("fake")
}

func isCancelStopReason(reason string) bool { panic("fake") }

func waitStreamSettle(ctx context.Context, ps *promptStreamState) error { panic("fake") }
