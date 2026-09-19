package workflow

import (
	"context"
	"log"
	"sync"
	"time"
)

// CronRunner executes a workflow by id (async, trust-checked by caller).
type CronRunner interface {
	RunCronWorkflow(ctx context.Context, wfID string)
}

// CronScheduler runs trusted cron workflows while the app is open.
type CronScheduler struct {
	mu     sync.Mutex
	cancel context.CancelFunc
	runner CronRunner
	listFn func() ([]Workflow, error)
}

func NewCronScheduler(runner CronRunner, listFn func() ([]Workflow, error)) *CronScheduler {
	panic("fake")
}

func (s *CronScheduler) Restart(ctx context.Context) { panic("fake") }

func (s *CronScheduler) Stop() { panic("fake") }
