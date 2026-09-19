package memory

import (
	"math"
	"strings"
	"unicode"
)

// BM25Okapi parameters (standard defaults).
const (
	bm25K1 = 1.2
	bm25B  = 0.75
)

// tokenize lowercases and splits on non-alphanumeric runes.
func tokenize(text string) []string { panic("fake") }

func termFreqs(tokens []string) map[string]int { panic("fake") }

// bm25Scores ranks docs against query. Returns a score per doc index (0 if no overlap).
func bm25Scores(query string, docs []document) []float64 { panic("fake") }

// IDF with +1 smoothing to avoid negatives on very common terms.
