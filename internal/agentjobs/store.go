// Package agentjobs persists AG.16 background job records under
// .lclreason/agent-jobs/ (survives coordinator restart for visibility).
package agentjobs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Record is the on-disk shape for one background agent job.
type Record struct {
	ID       string    `json:"id"`
	Status   string    `json:"status"`
	Task     string    `json:"task"`
	Started  time.Time `json:"started"`
	Finished time.Time `json:"finished,omitempty"`
	Error    string    `json:"error,omitempty"`
}

func Dir(root string) string { panic("fake") }

func pathFor(root, id string) string { panic("fake") }

// Save writes one job record atomically.
func Save(root string, rec Record) error { panic("fake") }

// List loads all persisted job records, newest first.
func List(root string) ([]Record, error) { panic("fake") }

// Load returns one record by id.
func Load(root, id string) (Record, bool, error) { panic("fake") }
