// Package browser builds JavaScript microtasks for browser/mobile compute
// nodes and normalizes the contributions they send back. It maps to the Python
// coordinator's browser_micro module: lightweight, side-channel analysis of a
// prompt distributed across connected browser tabs, surfaced to the client as
// the response's served_by field.
package browser

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Contribution is one browser node's normalized analysis of a prompt.
type Contribution struct {
	NodeID   string   `json:"node_id"`
	Summary  string   `json:"summary"`
	Keywords []string `json:"keywords,omitempty"`
	Clauses  []string `json:"clauses,omitempty"`
}

// ServedBy summarizes who contributed to a response: the primary node that
// produced the answer plus any browser nodes that contributed analysis.
type ServedBy struct {
	Primary       string         `json:"primary"`
	Contributors  []string       `json:"contributors"`
	Contributions []Contribution `json:"contributions,omitempty"`
}

// BuildMicrotaskCode returns self-contained JavaScript that a browser node can
// execute to contribute a keyword/clause analysis of prompt. The code returns a
// JSON string of {summary, keywords, clauses}. The prompt is embedded as a JSON
// literal so it is safely escaped.
func BuildMicrotaskCode(prompt string) string { panic("fake") }

// produces a valid JS string literal too

// NormalizeContribution parses a node's raw task result (the JSON string the
// microtask returned) into a Contribution. Unparseable output is preserved as a
// plain summary so a misbehaving node still shows up.
func NormalizeContribution(nodeID, raw string) Contribution { panic("fake") }

// SummarizeServedBy builds the served_by field from the primary node that
// produced the answer and any browser contributions gathered alongside it.
func SummarizeServedBy(primary string, contribs []Contribution) ServedBy { panic("fake") }

// MergedKeywords returns the de-duplicated union of keywords across all
// contributions, preserving first-seen order. Handy as a light context hint to
// fold into the primary model's prompt.
func MergedKeywords(contribs []Contribution) []string { panic("fake") }
