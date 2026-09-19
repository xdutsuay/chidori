package fakesync

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"
)

// HandlerConfig holds shared settings for selection UI HTTP handlers.
type HandlerConfig struct {
	SourceRoot string
	DestRoot   string   // PlanHandler default dest for mapping dst paths
	Denylist   []string // empty uses DefaultDenylist()
}

type apiMapping struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
	Rel string `json:"rel"`
}

type planAPIResponse struct {
	Mappings []apiMapping `json:"mappings"`
	Count    int          `json:"count"`
}

type syncAPIRequest struct {
	Dest     string   `json:"dest"`
	Selected []string `json:"selected"`
	DryRun   bool     `json:"dry_run"`
}

type syncAPIResponse struct {
	Mappings []apiMapping `json:"mappings"`
	Count    int          `json:"count"`
	Written  int          `json:"written"`
}

func (cfg HandlerConfig) denylist() []string { panic("fake") }

func resultToAPI(sourceRoot string, res Result) ([]apiMapping, error) { panic("fake") }

// PlanHandler serves GET /api/plan -> JSON {mappings:[{src,dst,rel}], count}.
func PlanHandler(cfg HandlerConfig) http.Handler { panic("fake") }

// SyncHandler serves POST /api/sync -> JSON {mappings, count, written}.
// Request body: {dest, selected, dry_run}.
func SyncHandler(cfg HandlerConfig) http.Handler { panic("fake") }

func writeJSON(w http.ResponseWriter, v any) { panic("fake") }
