package tools

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// maxFetchBytes caps how much of a page we read into memory.
const maxFetchBytes = 1 << 20 // 1 MiB

// maxFetchText caps the cleaned text returned to the model so a huge page
// doesn't blow the context window.
const maxFetchText = 8000

var scriptStyleRe = regexp.MustCompile(`(?is)<(script|style|noscript)[^>]*>.*?</(script|style|noscript)>`)
var whitespaceRe = regexp.MustCompile(`[ \t\r\f\v]+`)
var blankLinesRe = regexp.MustCompile(`\n{3,}`)

// FetchURLTool fetches a single web page and returns its readable text.
// It complements web_search: search to find URLs, then fetch to read one.
// Params: url (string, required). Only http/https schemes are allowed.
func FetchURLTool() ToolFunc { panic("fake") }

// extractReadableText strips scripts/styles and tags, leaving collapsed text.
func extractReadableText(html string) string { panic("fake") }

// shared with web_search.go

// Restore paragraph breaks roughly, then collapse runs of blank lines.
