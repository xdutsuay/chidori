package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// ---- "Save as Skill" (borrowed from Hermes Agent, §14.2 #3) ----
//
// After a successful agent/plan/debug run, the user can save it as a named,
// reusable prompt template stored per-workspace under .lclreason/skills/,
// the same way .lclreason/rules already is — plain markdown, versioned with
// the rest of the repo, no separate database. Resurfaced via the command
// palette (public/js/06-palette-keys.js), which lists saved skills and inserts the template into
// the chat input when one is picked.

type skillStep struct {
	Label  string `json:"label"`
	Status string `json:"status"`
}

type saveSkillRequest struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Task        string      `json:"task"`
	Steps       []skillStep `json:"steps"`
	EditedFiles []string    `json:"edited_files"`
	Lessons     string      `json:"lessons,omitempty"` // optional lesson to seed/append under ## Lessons
}

type evolveSkillRequest struct {
	Name   string `json:"name"`   // skill name or slug
	Slug   string `json:"slug"`   // preferred if set
	Lesson string `json:"lesson"` // short lesson string to append
}

var skillSlugRe = regexp.MustCompile(`[^a-z0-9]+`)

func slugifySkillName(s string) string { panic("fake") }

// SaveSkill writes a reusable per-workspace prompt template distilled from
// a completed chat run. This is a direct user action — the user reviews the
// agent's output in chat, then explicitly clicks "Save as Skill" — so,
// consistent with Explorer New File / Save As, it is not trust-gated even
// though it writes a file.
//
// When lessons is set and the skill file already exists, the lesson is
// appended under ## Lessons (Wave F skill evolve) instead of overwriting.
func (h *Handlers) SaveSkill(w http.ResponseWriter, r *http.Request) { panic("fake") }

// Evolve path: existing skill + lessons field → append, don't overwrite.

// EvolveSkill appends a short lesson under ## Lessons on an existing skill
// markdown file (POST /api/skills/evolve). Creates the section if missing.
func (h *Handlers) EvolveSkill(w http.ResponseWriter, r *http.Request) { panic("fake") }

// appendSkillLesson adds a bullet under ## Lessons, creating the section if
// the skill markdown doesn't have one yet.
func appendSkillLesson(content, lesson string) string { panic("fake") }

// Insert after the heading line (and any blank line that follows).

// Skip one extra blank line after the heading if present.

type skillSummary struct {
	Slug        string `json:"slug"`
	Path        string `json:"path"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Template    string `json:"template"` // full markdown body, used to prefill chat input when picked
}

// ListSkills returns every saved skill under .lclreason/skills/*.md, parsed
// enough to show a name/description in the command palette and hand back
// the full body as a ready-to-edit chat prompt when one is picked.
func (h *Handlers) ListSkills(w http.ResponseWriter, r *http.Request) { panic("fake") }

// No .lclreason/skills directory yet is the common (not an error)
// case — nothing has been saved there so far.

// ResolveSkill looks up a skill by name/slug for chat context injection (KMA-124).
// Order: workspace .lclreason/skills → workspace .cursor/skills → $HOME/.cursor/skills
// → $HOME/.agents/skills. Absolute reads are allowlisted to those skill roots only.
func (h *Handlers) ResolveSkill(w http.ResponseWriter, r *http.Request) { panic("fake") }

// parseSkillFrontmatter does a minimal, dependency-free extraction of the
// name/description lines from a "---\nkey: value\n---\n" block — good
// enough for files this same handler generates, without pulling in a full
// YAML parser for two fields.
func parseSkillFrontmatter(content string) (name, description, body string) { panic("fake") }
