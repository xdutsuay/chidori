package netdoc

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultMaxBody      = 512 * 1024
	defaultMaxRedirects = 5
	defaultFetchTimeout = 20 * time.Second
	defaultMaxTextLen   = 32 * 1024
	defaultMaxLinkCount = 50
)

// Config tunes fetch limits for netdoc.
type Config struct {
	MaxBody      int64
	MaxRedirects int
	Timeout      time.Duration
	MaxTextLen   int
	MaxLinks     int
	// TestAllowLoopback disables SSRF private-space blocking (httptest only).
	TestAllowLoopback bool
}

func (c *Config) applyDefaults() { panic("fake") }

// Client fetches and extracts public HTML with SSRF guards.
type Client struct {
	cfg      Config
	http     *http.Client
	resolver *net.Resolver
}

func NewClient(cfg Config) *Client { panic("fake") }

func validateRequestURL(u *url.URL) error { panic("fake") }

func (c *Client) resolveHost(ctx context.Context, host string) error { panic("fake") }

func (c *Client) resolveHostStrict(ctx context.Context, host string) error { panic("fake") }

func (c *Client) dialContext(ctx context.Context, network, address string) (net.Conn, error) {
	panic("fake")
}

// Fetch downloads url and returns extracted text (SSRF-hardened).
func (c *Client) Fetch(ctx context.Context, rawURL string) (*Result, error) { panic("fake") }

// SetResolver replaces DNS resolution (tests).
func (c *Client) SetResolver(r *net.Resolver) { panic("fake") }
