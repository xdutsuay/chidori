package tools

import (
	"context"
	"fmt"
	"io/fs"
	"regexp"
	"strconv"
	"strings"
	"time"

	ws "github.com/xdutsuay/lclreason/internal/workspace"
)

// RegisterWorkspaceTools adds coding agent tools backed by a workspace FS.
// fs may not have a root configured yet — its root can be set later at
// runtime (e.g. via the "Open Folder" picker), and these tools resolve
// against it dynamically on every call, so they're safe to register early.
func RegisterWorkspaceTools(r *Registry, fs *ws.FS) { panic("fake") }

func readFileTool(fs *ws.FS) ToolFunc { panic("fake") }

func writeFileTool(fs *ws.FS) ToolFunc { panic("fake") }

func applyPatchTool(fs *ws.FS) ToolFunc { panic("fake") }

func grepWorkspaceTool(wfs *ws.FS) ToolFunc { panic("fake") }

func listDirTool(fs *ws.FS) ToolFunc { panic("fake") }

func matchGlob(glob, path string) bool { panic("fake") }

func globToRegex(glob string) string { panic("fake") }

// ParseLineParam helper for tests.
func ParseLineParam(params map[string]any, key string) int { panic("fake") }
