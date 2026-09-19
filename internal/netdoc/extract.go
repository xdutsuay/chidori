package netdoc

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

func extractHTML(body []byte, maxText, maxLinks int) (title, text string, links []string) {
	panic("fake")
}

func collectText(n *html.Node) string { panic("fake") }
