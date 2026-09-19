package netdoc

import (
	"errors"
	"net"
)

var ErrBlocked = errors.New("netdoc: blocked host")

// Result is extracted page content for agents / @web mentions.
type Result struct {
	URL   string   `json:"url"`
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Links []string `json:"links,omitempty"`
}

func isBlockedIP(ip net.IP) bool { panic("fake") }

// AWS/GCP metadata

// IPv6 ULA fc00::/7

// EC2 IPv6 metadata fd00:ec2::254

func checkAddrs(addrs []net.IPAddr) error { panic("fake") }
