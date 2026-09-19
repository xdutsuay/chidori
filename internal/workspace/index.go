package workspace

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"strings"

	"github.com/xdutsuay/lclreason/internal/memory"
)

const chunkLines = 40

// indexSkipDir names are never descended into during workspace indexing.
var indexSkipDir = map[string]bool{
	".git": true, "node_modules": true, "vendor": true, "build": true,
	"dist": true, ".gradle": true, "target": true, "__pycache__": true,
	".lclreason": true, "graphify-out": true,
}

func shouldSkipIndexPath(rel string) bool { panic("fake") }

func fileContentHash(data []byte) string { panic("fake") }

// IndexAll walks the workspace and upserts file chunks into the vector store.
// Unchanged files (matching persisted content_hash) are skipped. Changed files
// replace their old chunks. Deleted files' chunks are removed. Non-workspace
// memory (documents without a path key) is left alone (KMA-63).
func IndexAll(ctx context.Context, ws *FS, mem *memory.VectorStore) { panic("fake") }

func isTextFile(path string) bool { panic("fake") }

// RecallChunks searches the memory store for workspace chunks relevant to query.
func RecallChunks(ctx context.Context, mem *memory.VectorStore, query string, k int) []string {
	panic("fake")
}

// bootstrapDocPaths are read when vector recall is thin — helps "what is this product" without browse_*.
var bootstrapDocPaths = []string{
	"README.md", "markdowns_open/README.md", "PRD.md", "docs/PRD.md",
	"ARCHITECTURE.md", "docs/ARCHITECTURE.md",
}

const bootstrapMaxBytes = 4000

// AgentRecallChunks combines vector recall with README/PRD bootstrap snippets.
func AgentRecallChunks(ctx context.Context, ws *FS, mem *memory.VectorStore, query string, k int) []string {
	panic("fake")
}

// Bootstrap key docs when recall is empty or very thin (common on fresh index).
