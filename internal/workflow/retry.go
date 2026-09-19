package workflow

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	maxRetryAttemptsHard = 10
	defaultRetryBackoff  = time.Second
	maxRetryBackoff      = 30 * time.Second
)

func normalizeRetry(cfg *RetryConfig) *RetryConfig { panic("fake") }

func retryEnabled(step *Step) bool { panic("fake") }

func retryOn(cfg *RetryConfig, name string) bool { panic("fake") }

func shellFailureRetryable(cfg *RetryConfig, exitCode int, execErr error) bool { panic("fake") }

func llmFailureRetryable(cfg *RetryConfig, err error) bool { panic("fake") }

func parseBackoff(raw string) (time.Duration, error) { panic("fake") }

func retryDelay(cfg *RetryConfig, attempt int) (time.Duration, error) { panic("fake") }

// Up to 25% jitter.

func sleepRetryBackoff(ctx context.Context, cfg *RetryConfig, attempt int) error { panic("fake") }
