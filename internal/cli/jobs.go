package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/config"
)

func jobsHTTPClient(o Options) *http.Client { panic("fake") }

func loadJobsBase(ctx context.Context, o Options) (*config.Config, string, error) { panic("fake") }

type agentJobRow struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Task     string `json:"task"`
	Error    string `json:"error,omitempty"`
	Started  string `json:"started,omitempty"`
	Finished string `json:"finished,omitempty"`
}

func fetchAgentJobs(ctx context.Context, o Options) ([]agentJobRow, error) { panic("fake") }

// JobsList renders background agent jobs from GET /api/agent/jobs.
func JobsList(ctx context.Context, o Options) (string, error) { panic("fake") }

// JobStatus prints one job from GET /api/agent/jobs.
func JobStatus(ctx context.Context, o Options, id string) (string, error) { panic("fake") }

// JobCancel posts POST /api/agent/jobs/{id}/cancel.
func JobCancel(ctx context.Context, o Options, id string) (string, error) { panic("fake") }
