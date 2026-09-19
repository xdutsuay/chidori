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

// WebSearchTool searches the web using DuckDuckGo's HTML interface.
// No API key required — it scrapes the lite HTML results page.
func WebSearchTool() ToolFunc { panic("fake") }

type searchResult struct {
	title   string
	url     string
	snippet string
}

func duckDuckGoSearch(ctx context.Context, client *http.Client, query string, maxResults int) ([]searchResult, error) {
	panic(
		// Use DuckDuckGo HTML lite version — no JS needed.
		"fake")
}

// parseDDGResults extracts titles, URLs, and snippets from DuckDuckGo HTML lite.
func parseDDGResults(html string, max int) []searchResult { panic("fake") }

// Pattern: result links are in <a rel="nofollow" class="result__a" href="...">title</a>
// Snippets are in <a class="result__snippet" ...>snippet</a>

// DuckDuckGo wraps URLs in a redirect; extract the real URL.

var htmlTagRe = regexp.MustCompile(`<[^>]*>`)

func stripHTMLTags(s string) string { panic("fake") }
