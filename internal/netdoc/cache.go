package netdoc

import (
	"context"
	"strings"
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"
)

const defaultCacheEntries = 32

// Cache wraps Client with a bounded in-memory LRU (M2).
type Cache struct {
	client *Client
	mu     sync.Mutex
	lru    *lru.Cache[string, *Result]
}

func NewCache(client *Client, maxEntries int) (*Cache, error) { panic("fake") }

func (c *Cache) Fetch(ctx context.Context, rawURL string) (*Result, error) { panic("fake") }
