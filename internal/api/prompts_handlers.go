package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/xdutsuay/lclreason/internal/prompts"
)

var promptSlugRe = regexp.MustCompile(`[^a-z0-9]+`)

func slugifyPromptID(s string) string { panic("fake") }

type savePromptRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Modes       string `json:"modes"` // comma-separated, e.g. agent,ask
	Body        string `json:"body"`
}

// ListPrompts returns bundled + workspace prompt profiles (metadata).
func (h *Handlers) ListPrompts(w http.ResponseWriter, r *http.Request) { panic("fake") }

// GetPrompt returns one profile body by id (builtin or workspace slug).
func (h *Handlers) GetPrompt(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) mergePromptProfiles(includeBody bool) ([]prompts.Profile, error) { panic("fake") }

// workspace overrides builtin id

// activePromptProfileText returns the body of the user's selected profile.
func (h *Handlers) activePromptProfileText() string { panic("fake") }

// withPromptProfileInstructions appends the active Settings prompt profile to
// the agent ModeInstructions (stable system channel). Profiles used to be
// prepended to OpenContext, which buildAgentPrompt labels ambient — models
// are told to ignore ambient unless the Task references it (KMA-214).
func (h *Handlers) withPromptProfileInstructions(base string) string { panic("fake") }

// SavePrompt writes a workspace prompt profile to .lclreason/prompts/{id}.md.
// Builtin ids cannot be overwritten — save under a new id instead.
func (h *Handlers) SavePrompt(w http.ResponseWriter, r *http.Request) { panic("fake") }
