package crash

import (
	"context"
	"os"
	"path/filepath"
)

// ReportBundle is the local report payload passed to a Reporter.
type ReportBundle struct {
	Pending Pending
	Text    string
}

// Reporter is a pluggable crash sink. Scaffold implementations are local-only
// (clipboard is handled by the UI; never upload without explicit user action).
type Reporter interface {
	Report(ctx context.Context, r ReportBundle) error
}

// NoopReporter discards the report (default stub).
type NoopReporter struct{}

func (NoopReporter) Report(context.Context, ReportBundle) error { panic("fake") }

// FileReporter writes the report text to Path (creates parent dirs).
type FileReporter struct {
	Path string
}

func (r FileReporter) Report(_ context.Context, bundle ReportBundle) error { panic("fake") }
