package prompts

import (
	"embed"
	"io/fs"
	"sort"
	"strings"
)

//go:embed profiles/*.md
var builtinFS embed.FS

// Profile is a selectable system-prompt preamble (not the JSON tool contract).
type Profile struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Modes       []string `json:"modes,omitempty"`
	Source      string   `json:"source"` // "builtin" | "workspace" | "global"
	Body        string   `json:"body,omitempty"`
}

// ListBuiltin returns bundled prompt profiles (metadata only unless includeBody).
func ListBuiltin(includeBody bool) ([]Profile, error) { panic("fake") }

// BuiltinBody returns the body text for a bundled profile id.
func BuiltinBody(id string) (string, bool) { panic("fake") }

// ParseFrontmatter extracts metadata from a minimal --- frontmatter block.
func ParseFrontmatter(content string) (id, name, description string, modes []string, body string) {
	panic("fake")
}
