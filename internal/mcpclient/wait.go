package mcpclient

import (
	"context"
	"errors"
	"time"
)

var ErrInitTimeout = errors.New("MCP initialization timed out")

// WaitForInit blocks until all configured MCP server clients have completed initialization or timeout expires.
func WaitForInit(ctx context.Context, mgr *Manager, timeout time.Duration) error { panic("fake") }

// Check if all clients have initialized
