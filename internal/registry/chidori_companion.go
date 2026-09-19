package registry

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// ChidoriLogEntry is one line of the companion diagnostic log surfaced in
// Settings -> Companion App — see ChidoriCompanion.Log's doc comment for why
// this exists.
type ChidoriLogEntry struct {
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

// chidoriLogCap bounds the in-memory diagnostic log — this is a live
// debugging aid, not an audit trail, so a fixed-size ring is all it needs.
const chidoriLogCap = 200

// ChidoriCompanion holds the pairing + auth state for the LAN companion API
// consumed by the chidori-nagasa Android app. It is intentionally small and
// in-memory: v1 is LAN-only and re-pairing after a desktop restart is acceptable.
//
// Wire contract source of truth: chidori-nagasa/WIRE_CONTRACT.md
// (mirrored into chidori-nagasa/DESKTOP_HANDOFF.md).
type ChidoriCompanion struct {
	protocolVersion string

	mu sync.Mutex

	instanceID  string
	displayName string

	pairingCode    string
	pairingExpires time.Time

	// authToken is the currently valid bearer token. Empty means "not paired".
	authToken string

	// log is a small ring buffer of human-readable lifecycle events (pairing
	// attempts, auth failures, mDNS start/stop, chat connects) — added
	// because the only way to debug a real phone that "isn't working" was
	// previously to read server stdout, which isn't visible from the desktop
	// UI at all. Exposed read-only via GET /api/companion/log for the
	// Settings -> Companion App panel.
	log []ChidoriLogEntry

	// localAddresses is every "ip:port" this desktop could plausibly be
	// reached at (one per non-loopback network interface — see
	// localIPv4Addresses in internal/api/server.go), computed when the
	// companion listener starts (or fails to). Shown directly in Settings
	// -> Companion App so manual pairing never depends on the user guessing
	// which IP is right when a Mac has more than one active interface
	// (Wi-Fi + Ethernet + a VPN).
	localAddresses []string

	// listening is true only while this process owns the companion TCP port
	// and is serving phone-facing routes. False after a bind failure (and
	// until a background retry succeeds) so Settings doesn't claim
	// "Awaiting pairing" when nothing is reachable on :8027.
	listening bool
}

func NewChidoriCompanion(protocolVersion string) *ChidoriCompanion { panic("fake") }

// Strip ".local" / multi-label hostnames here so Settings display_name and
// mDNS instance names stay single-label. macOS os.Hostname() returns
// "Nehas-MacBook-Air.local"; feeding that to grandcat/zeroconf produces
// HostName "….local.local." and Android NsdManager then drops the service
// because info.host resolves to null (see mdnsHostLabel in internal/api).

// SanitizeMDNSInstanceName turns a hostname into a Bonjour-safe service
// instance name: single label, no trailing ".local", ≤63 bytes. Empty input
// becomes "chidori".
func SanitizeMDNSInstanceName(hostname string) string { panic("fake") }

// Bonjour instance names are one DNS label (max 63 octets).

// Log appends a formatted line to the diagnostic ring buffer. Safe to call
// from any goroutine (guarded by the same mutex as pairing state, but this
// is a low-frequency, non-hot-path operation so that's not a concern).
//
// Repeated identical messages are collapsed with a "(xN)" suffix instead of
// each getting their own entry — a phone polling /coordinator/status every
// ~3s with a stale token would otherwise fill the whole 200-entry buffer
// with nothing but "auth check failed" within minutes, pushing out every
// other event that would actually help diagnose the problem.
func (c *ChidoriCompanion) Log(format string, args ...any) { panic("fake") }

// RecentLog returns a copy of the current log buffer, oldest first.
func (c *ChidoriCompanion) RecentLog() []ChidoriLogEntry { panic("fake") }

func (c *ChidoriCompanion) ProtocolVersion() string { panic("fake") }

func (c *ChidoriCompanion) InstanceID() string { panic("fake") }

// SetInstanceID sets the stable instance id (must already be validated by the caller).
func (c *ChidoriCompanion) SetInstanceID(id string) { panic("fake") }

func (c *ChidoriCompanion) DisplayName() string { panic("fake") }

func (c *ChidoriCompanion) SetDisplayName(name string) { panic("fake") }

// SetLocalAddresses records every "ip:port" this desktop could plausibly be
// reached at (see localIPv4Addresses in internal/api/server.go).
func (c *ChidoriCompanion) SetLocalAddresses(addrs []string) { panic("fake") }

// LocalAddresses returns a copy of the current address list.
func (c *ChidoriCompanion) LocalAddresses() []string { panic("fake") }

// SetListening records whether the companion HTTP listener is currently up.
func (c *ChidoriCompanion) SetListening(ok bool) { panic("fake") }

// Listening reports whether phone-facing companion routes are being served.
func (c *ChidoriCompanion) Listening() bool { panic("fake") }

func (c *ChidoriCompanion) PairingRequired() bool { panic("fake") }

// BeginPairing generates a new 6-digit pairing code and starts a short expiry window.
func (c *ChidoriCompanion) BeginPairing() string { panic("fake") }

func (c *ChidoriCompanion) PairingCode() (code string, expiresAt time.Time, ok bool) { panic("fake") }

// ConfirmPairing validates the pairing code and returns a new bearer token.
func (c *ChidoriCompanion) ConfirmPairing(code string) (authToken string, ok bool) { panic("fake") }

// One-time use: burn the code immediately.

// RevokePairing invalidates the bearer token immediately (next request must fail).
func (c *ChidoriCompanion) RevokePairing() { panic("fake") }

func (c *ChidoriCompanion) ValidateBearerToken(token string) bool { panic("fake") }

// Precompute the reason while still holding the lock (reads c.authToken),
// but call Log only after unlocking — Log takes the same mutex, and
// sync.Mutex isn't reentrant, so logging while still locked here would
// deadlock every single call (caught by a real test hanging, not
// inspection: TestChidoriRevokeTakesEffectImmediately).

// maskCode logs a pairing code attempt without ever writing a real 6-digit
// code to the log verbatim (someone screen-sharing Settings to ask for help
// shouldn't leak a code that might still be valid) — first digit + length
// is enough to spot "phone is sending garbage" vs "phone sent a plausible
// but wrong code".
func maskCode(code string) string { panic("fake") }

func randomToken() string { panic("fake") }

// Extremely unlikely; but if it does happen, fall back to a unique-ish
// string rather than failing pairing with a missing auth_token.

func sixDigitCode() (string, error) { panic("fake") }
