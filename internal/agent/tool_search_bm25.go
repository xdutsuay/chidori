package agent

import (
	"math"
	"sort"
	"strings"
	"unicode"
)

// BM25Okapi defaults (same as internal/memory) plus hybrid lexical boosts so
// exact tool-name queries still dominate soft description matches.
const (
	toolSearchBM25K1   = 1.2
	toolSearchBM25B    = 0.75
	toolSearchMaxHits  = 12
	toolSearchLexExact = 100.0
	toolSearchLexSub   = 50.0
	toolSearchLexName  = 10.0
	toolSearchLexDesc  = 3.0
)

// toolSearchDoc is one catalog entry indexed for tool_search.
type toolSearchDoc struct {
	Name string
	Desc string
	Body string // searchable corpus: name, description, param docs
}

// toolSearchHit is a ranked match returned to searchTools formatting.
type toolSearchHit struct {
	Name  string
	Desc  string
	Score float64
}

// rankToolSearch ranks docs with BM25 + lexical hybrid. Empty query or empty
// docs yield nil. Scores are comparable only within a single call.
func rankToolSearch(query string, docs []toolSearchDoc) []toolSearchHit { panic("fake") }

func lexicalToolBoost(qLower string, d toolSearchDoc) float64 { panic("fake") }

func bm25ScoresOverBodies(query string, docs []toolSearchDoc) []float64 { panic("fake") }

func tokenizeToolSearch(text string) []string { panic("fake") }

func termFreqsToolSearch(tokens []string) map[string]int { panic("fake") }

// toolSearchBody builds the BM25 corpus for one tool. Name is repeated so
// identifier tokens outweigh long descriptions when both match.
func toolSearchBody(name, desc, paramsText string) string { panic("fake") }

// flattenSchemaText extracts property names and description strings from an
// OpenAI-style JSON Schema parameters object (and nested maps/slices).
func flattenSchemaText(v any) string { panic("fake") }
