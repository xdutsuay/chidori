package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/xdutsuay/lclreason/internal/session"
)

func (h *Handlers) ListSessions(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) CreateSession(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) GetSession(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) UpdateSession(w http.ResponseWriter, r *http.Request) { panic("fake") }

// saveSession PATCHes mode/provider/tabs without title — keep existing
// auto-name / rename instead of wiping to "" (reverts tabs to "Chat").

func (h *Handlers) DeleteSession(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) SaveSessionMessages(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) WorkspaceSearch(w http.ResponseWriter, r *http.Request) { panic("fake") }

// kept as the include param name for backward compat with existing callers

// WorkspaceReplacePreview reports which files a replace-in-files would touch
// and how many matches each has, without writing anything.
func (h *Handlers) WorkspaceReplacePreview(w http.ResponseWriter, r *http.Request) { panic("fake") }

// WorkspaceReplaceApply performs the replace across every matching file.
// Unlike agent auto-apply, this isn't gated on workspace trust: it's a
// direct, explicit user action (same class as Cmd+S save), not an
// AI-proposed edit being silently auto-applied.
func (h *Handlers) WorkspaceReplaceApply(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) workspaceReplace(w http.ResponseWriter, r *http.Request, apply bool) {
	panic("fake")
}

// include globs, kept as "glob" for backward compat

func (h *Handlers) WorkspaceFiles(w http.ResponseWriter, r *http.Request) { panic("fake") }

func (h *Handlers) ClusterStatus(w http.ResponseWriter, r *http.Request) { panic("fake") }
