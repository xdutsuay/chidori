package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Path is the file this config was loaded from ("" if none). Not
	// serialized; set by Load and used by Save to write back in place
	// (e.g. persisting the workspace root chosen via Open Folder).
	Path string `yaml:"-"`

	Mode         string            `yaml:"mode"`
	Port         int               `yaml:"port"`
	Coordinator  string            `yaml:"coordinator"`
	Advertise    string            `yaml:"advertise"` // worker: host:port the coordinator should reach it at ("" = hostname:port)
	DB           string            `yaml:"db"`
	Chain        ChainConfig       `yaml:"chain"`
	Cache        CacheConfig       `yaml:"cache"`
	Heartbeat    YAMLDur           `yaml:"heartbeat"`
	OfflineAfter YAMLDur           `yaml:"offline_after"`
	Provider     ProviderConfig    `yaml:"provider"`
	Planner      PlannerConfig     `yaml:"planner"`
	Memory       MemoryConfig      `yaml:"memory"`
	Compaction   CompactionConfig  `yaml:"compaction"`
	Workspace    WorkspaceConfig   `yaml:"workspace"`
	Code         CodeConfig        `yaml:"code"`
	Auth         AuthConfig        `yaml:"auth"`
	Debug        DebugConfig       `yaml:"debug"`
	SecretsFile  string            `yaml:"secrets_file"`
	Profiling    bool              `yaml:"profiling"`
	PrewarmModel bool              `yaml:"prewarm_model"`
	UI           UIConfig          `yaml:"ui"`
	Companion    CompanionConfig   `yaml:"companion"`
	Sandbox      SandboxConfig     `yaml:"sandbox"`
	Agent        AgentConfig       `yaml:"agent,omitempty" json:"agent,omitempty"`
	MCPServers   []MCPServerConfig `yaml:"mcp_servers,omitempty" json:"mcp_servers,omitempty"` // ADR-0012
	Usage        UsageConfig       `yaml:"usage,omitempty" json:"usage,omitempty"`
}

// AgentConfig tunes agent-loop subagent behavior (KMA-107).
type AgentConfig struct {
	// SubagentModels is the allowlist for delegate_subtask params.model.
	// Empty defaults to inherit-only (parent model). Include named model IDs
	// that ResolveSubagentModel can map to a configured inference source.
	SubagentModels []string `yaml:"subagent_models,omitempty" json:"subagent_models,omitempty"`
	// SubagentMaxInflight caps concurrent async delegate_subtask jobs (0 = 2).
	SubagentMaxInflight int `yaml:"subagent_max_inflight,omitempty" json:"subagent_max_inflight,omitempty"`
}

// UsageConfig controls the global token-usage JSONL log (KMA-87,
// Settings → Usage). Retention defaults to 90 days when unset/zero so the
// feature cannot quietly fill the disk.
type UsageConfig struct {
	Dir           string `yaml:"dir,omitempty" json:"dir,omitempty"`
	RetentionDays int    `yaml:"retention_days,omitempty" json:"retention_days,omitempty"`
}

// SandboxConfig tunes workspace sandboxing options.
type SandboxConfig struct {
	// Enabled controls snapshot sandboxing: nil = auto (on in git repos), true = force on, false = force off.
	Enabled             *bool  `yaml:"enabled,omitempty" json:"enabled,omitempty"`
	Mode                string `yaml:"mode" json:"mode"` // snapshot | worktree | seatbelt | disabled
	AutoRollbackOnError bool   `yaml:"auto_rollback_on_error" json:"auto_rollback_on_error"`
}

// CompanionConfig controls the LAN companion API for chidori-nagasa
// (protocol 1.2.0+). Phone traffic uses a dedicated listen port — separate
// from Port (the IDE/coordinator API) — so mDNS and manual pairing always
// target the companion surface, not the Wails UI.
type CompanionConfig struct {
	// Port is the dedicated companion HTTP/WS listen port (default 8027).
	// 0 means "use DefaultCompanionPort".
	Port int `yaml:"port"`
}

// DefaultCompanionPort is the protocol 1.2.0 default for phone discovery/pairing.
const DefaultCompanionPort = 8027

// ListenPort returns the companion listen port, applying the 8027 default.
func (c CompanionConfig) ListenPort() int { panic("fake") }

// MCPServerConfig is one external MCP server chidori's agent can consume as
// a source of tools (ADR-0012). Deliberately its own type here rather than
// importing internal/mcpclient.ServerConfig — config.go stays a leaf
// package other things import, never the reverse; internal/api converts
// between the two at the one call site that needs mcpclient's actual type.
type MCPServerConfig struct {
	Name    string   `yaml:"name" json:"name"`
	Command string   `yaml:"command" json:"command"`
	Args    []string `yaml:"args,omitempty" json:"args,omitempty"`
}

// UIConfig persists the IDE's editor/UI preferences and default
// provider/model across launches (ADR-0005). Previously these were in-memory
// only in the frontend and reset on every relaunch. Empty/zero values mean
// "use the frontend's own built-in default" (e.g. an empty WordWrap means
// "off", the existing default) rather than an explicit override, so upgrading an
// existing config.yaml with no ui: block changes nothing.
// json tags are required here, not just yaml — this struct round-trips
// through both config.Save (yaml.Marshal) AND the /api/ui/prefs handlers
// (encoding/json). encoding/json ignores yaml struct tags entirely and falls
// back to the bare Go field name, which would silently produce/accept
// "WordWrap" instead of "word_wrap" with every field landing empty — keep
// the two tag sets identical in name so config.yaml and the JSON wire shape
// never disagree.
type UIConfig struct {
	WordWrap               string   `yaml:"word_wrap,omitempty" json:"word_wrap,omitempty"` // "off"|"on"|"bounded"
	Minimap                bool     `yaml:"minimap,omitempty" json:"minimap,omitempty"`
	FontSize               int      `yaml:"font_size,omitempty" json:"font_size,omitempty"`
	AutoSave               bool     `yaml:"auto_save,omitempty" json:"auto_save,omitempty"`
	AutoSaveOnFocusChange  bool     `yaml:"auto_save_on_focus_change,omitempty" json:"auto_save_on_focus_change,omitempty"`
	AutoSaveOnWindowChange bool     `yaml:"auto_save_on_window_change,omitempty" json:"auto_save_on_window_change,omitempty"`
	FormatOnSave           bool     `yaml:"format_on_save,omitempty" json:"format_on_save,omitempty"`
	Ligatures              bool     `yaml:"ligatures,omitempty" json:"ligatures,omitempty"`
	RenderWhitespace       bool     `yaml:"render_whitespace,omitempty" json:"render_whitespace,omitempty"`
	StickyScroll           bool     `yaml:"sticky_scroll,omitempty" json:"sticky_scroll,omitempty"`
	MultiCursorModifier    string   `yaml:"multi_cursor_modifier,omitempty" json:"multi_cursor_modifier,omitempty"` // "alt"|"ctrlCmd"
	CompletionDisabled     []string `yaml:"completion_disabled_langs,omitempty" json:"completion_disabled_langs,omitempty"`
	Theme                  string   `yaml:"theme,omitempty" json:"theme,omitempty"` // "dark"|"light"|"system" (ADR-0009)
	DefaultProvider        string   `yaml:"default_provider,omitempty" json:"default_provider,omitempty"`
	DefaultModel           string   `yaml:"default_model,omitempty" json:"default_model,omitempty"`
	DraftMode              bool     `yaml:"draft_mode,omitempty" json:"draft_mode,omitempty"` // ADR-0008: instant local draft, upgraded from remote on hard prompts
	// FastReading enables bionic-style fixation bolding on chat/plan prose
	// (not code fences). Opt-in; KMA-110.
	FastReading bool `yaml:"fast_reading,omitempty" json:"fast_reading,omitempty"`
	// ShowThinking surfaces model reasoning / thought TurnEvents in the chat UI.
	// Default true (KMA-109); json tag omits omitempty so explicit false round-trips.
	ShowThinking bool `yaml:"show_thinking" json:"show_thinking"`
	// AssistedThinking registers the deep_think agent tool for private scratch
	// reasoning. Opt-in; default false (KMA-109).
	AssistedThinking   bool   `yaml:"assisted_thinking,omitempty" json:"assisted_thinking,omitempty"`
	CenteredLayout     bool   `yaml:"centered_layout,omitempty" json:"centered_layout,omitempty"`
	SidebarHidden      bool   `yaml:"sidebar_hidden,omitempty" json:"sidebar_hidden,omitempty"`
	ChatCollapsed      bool   `yaml:"chat_collapsed,omitempty" json:"chat_collapsed,omitempty"`
	StatusBarHidden    bool   `yaml:"status_bar_hidden,omitempty" json:"status_bar_hidden,omitempty"`
	TerminalScrollback int    `yaml:"terminal_scrollback,omitempty" json:"terminal_scrollback,omitempty"` // T.7; 0 = xterm.js's own default (1000)
	TerminalShell      string `yaml:"terminal_shell,omitempty" json:"terminal_shell,omitempty"`           // T.10; "" = $SHELL / system default
	// Stored inverted (as "disabled") rather than "enabled" like every other
	// boolean pref in this struct: Monaco's own real default for this option
	// is true (unlike renderWhitespace/minimap/etc, which default to false),
	// so with omitempty an "enabled" field would make persisted-false
	// indistinguishable from never-set — silently resurrecting the setting
	// after every restart for anyone who explicitly turned it off. Storing
	// the minority "disabled" case instead means the common case (leave
	// Monaco's own default alone) is correctly represented by the field
	// being absent.
	RenderControlCharsDisabled bool `yaml:"render_control_chars_disabled,omitempty" json:"render_control_chars_disabled,omitempty"`
	// TC.8 — glob patterns (e.g. "*.env, secrets/*") where inline tab
	// completion is skipped entirely, so file contents never get sent to a
	// completion provider for those paths. Separate from CompletionDisabled
	// (TC.7, by language id) — this is by path/privacy, not by language.
	CompletionExcludedPaths []string `yaml:"completion_excluded_paths,omitempty" json:"completion_excluded_paths,omitempty"`
	// 1.4.5 — a separate size control for UI chrome (sidebar, tabs, chat
	// panel, status bar, activity bar) via CSS zoom, independent of
	// FontSize above (which only ever affected the Monaco editor itself).
	// 0 = 100%/default.
	UIZoomPercent int `yaml:"ui_zoom_percent,omitempty" json:"ui_zoom_percent,omitempty"`

	// companion_instance_id is the stable identity this desktop advertises over
	// mDNS (_chidori._tcp) and returns from /pairing/confirm for the Android LAN
	// companion. Stored here (rather than a separate config block) so it
	// persists across restarts alongside other per-install UI identity.
	CompanionInstanceID string `yaml:"companion_instance_id,omitempty" json:"companion_instance_id,omitempty"`

	// ExperimentalBrowserTools enables headless chromedp browse_* agent tools
	// (browse_navigate, browse_snapshot, browse_click). Off by default; trust-
	// gated like shell_exec when on. Settings → Experimental toggles this.
	ExperimentalBrowserTools bool `yaml:"experimental_browser_tools,omitempty" json:"experimental_browser_tools,omitempty"`

	// ExperimentalBrowserPanel shows an in-app iframe browser panel (human
	// browsing). Agent tools remain headless chromedp when ExperimentalBrowserTools is on.
	ExperimentalBrowserPanel bool `yaml:"experimental_browser_panel,omitempty" json:"experimental_browser_panel,omitempty"`

	// ExperimentalWebFetch enables @web mentions and browse_fetch read-only URL
	// extraction (internal/netdoc, SSRF-hardened). Off by default.
	ExperimentalWebFetch bool `yaml:"experimental_web_fetch,omitempty" json:"experimental_web_fetch,omitempty"`

	// ExperimentalExternalHarness enables the Grok external-harness path.
	// Off by default; Settings → Experimental controls it.
	ExperimentalExternalHarness bool `yaml:"experimental_external_harness,omitempty" json:"experimental_external_harness,omitempty"`

	// ExperimentalHarnessVisualizer enables the optional harness trace bus,
	// Diagnostics timeline, and JSONL replay (PRD harness visualizer).
	ExperimentalHarnessVisualizer bool `yaml:"experimental_harness_visualizer,omitempty" json:"experimental_harness_visualizer,omitempty"`
	// ExternalHarnessID is the preferred harness choice for new sessions
	// ("chidori" or "grok"). Empty means the default native harness.
	ExternalHarnessID string `yaml:"external_harness_id,omitempty" json:"external_harness_id,omitempty"`
	// GrokCommand / GrokArgs optionally override the PATH binary/args used for
	// the experimental external harness.
	GrokCommand string   `yaml:"grok_command,omitempty" json:"grok_command,omitempty"`
	GrokArgs    []string `yaml:"grok_args,omitempty" json:"grok_args,omitempty"`
	// HermesCommand / HermesArgs optionally override the PATH binary/args used for
	// the experimental Hermes harness.
	HermesCommand string   `yaml:"hermes_command,omitempty" json:"hermes_command,omitempty"`
	HermesArgs    []string `yaml:"hermes_args,omitempty" json:"hermes_args,omitempty"`

	// SidebarPosition is "left" (default) or "right" — flips the sidebar to
	// the right side of the editor, VS Code-style "Move Primary Side Bar Right".
	SidebarPosition string `yaml:"sidebar_position,omitempty" json:"sidebar_position,omitempty"`

	// SidebarWidthPx / ChatWidthPx are drag-splitter widths for the explorer
	// and chat panes (0 = CSS defaults 240 / 380).
	SidebarWidthPx int `yaml:"sidebar_width_px,omitempty" json:"sidebar_width_px,omitempty"`
	ChatWidthPx    int `yaml:"chat_width_px,omitempty" json:"chat_width_px,omitempty"`

	// All-Agent layout splitters (KMA-123). 0 = CSS defaults.
	AllAgentRailWidthPx    int `yaml:"all_agent_rail_width_px,omitempty" json:"all_agent_rail_width_px,omitempty"`
	AllAgentContextWidthPx int `yaml:"all_agent_context_width_px,omitempty" json:"all_agent_context_width_px,omitempty"`
	AllAgentDashHeightPx   int `yaml:"all_agent_dash_height_px,omitempty" json:"all_agent_dash_height_px,omitempty"`

	// AllAgentMode is AI.24 — chat-only layout (hide explorer/editor/terminal).
	AllAgentMode bool `yaml:"all_agent_mode,omitempty" json:"all_agent_mode,omitempty"`

	// KeybindingMode selects editor keybinding mode: "" (default), "vim", "emacs".
	// ED.20/ED.21 — vendor monaco-vim / monaco-emacs, toggle from Settings / palette.
	KeybindingMode string `yaml:"keybinding_mode,omitempty" json:"keybinding_mode,omitempty"`

	// Keybindings maps command id → serialized chord (e.g. "mod+shift+p").
	// Absent ids use built-in defaults in public/js/06-palette-keys.js (CFG.13).
	Keybindings map[string]string `yaml:"keybindings,omitempty" json:"keybindings,omitempty"`

	// ActivePromptProfile is the selected system-prompt profile id (builtin or .lclreason/prompts).
	// Empty or "default" means no extra profile preamble.
	ActivePromptProfile string `yaml:"active_prompt_profile,omitempty" json:"active_prompt_profile,omitempty"`
	// OpenChatSessionIDs is the set of IDE chat sessions shown in the tab strip
	// (KMA-179). Closing a tab removes an id here; DELETE /api/sessions is only
	// for the Chats manager. Nil/absent means "not yet migrated" (UI seeds from
	// the full session list); an explicit empty slice means all tabs closed.
	OpenChatSessionIDs []string `yaml:"open_chat_session_ids,omitempty" json:"open_chat_session_ids,omitempty"`
	// DevLog enables the rotating developer log without requiring the env var.
	// Shipping default stays off; Settings → Diagnostics can persist this for
	// local development.
	DevLog bool `yaml:"dev_log,omitempty" json:"dev_log,omitempty"`
}

// DebugConfig controls the developer activity timeline (/debug). Enabled by
// default; can be toggled at runtime via POST /api/debug.
type DebugConfig struct {
	Enabled bool `yaml:"enabled"`
}

// CompactionConfig controls server-side conversation compaction (the hybrid
// strategy: recent window verbatim + rolling summary of older turns + RAG recall
// of the most relevant older turns). When the flattened history fits the budget,
// nothing is changed.
type CompactionConfig struct {
	Enabled       bool   `yaml:"enabled"`
	ContextTokens int    `yaml:"context_tokens"` // model context budget (token estimate)
	ReserveTokens int    `yaml:"reserve_tokens"` // headroom kept for the response
	RecentTurns   int    `yaml:"recent_turns"`   // most-recent messages always kept verbatim
	RAGMessages   int    `yaml:"rag_messages"`   // max older messages pulled back by relevance
	SummaryModel  string `yaml:"summary_model"`  // model used to summarize older turns ("" = chat model)
}

// AuthConfig configures optional API-key auth on the coordinator. Auth is
// disabled when no key resolves (the default). Prefer api_key_env over a
// literal api_key so secrets stay out of the config file.
type AuthConfig struct {
	APIKey    string `yaml:"api_key"`
	APIKeyEnv string `yaml:"api_key_env"`
}

// Key returns the resolved API key, or "" when auth is disabled.
func (a AuthConfig) Key() string { panic("fake") }

type ChainConfig struct {
	Dir     string `yaml:"dir"`
	Default string `yaml:"default"`
}

type CacheConfig struct {
	Enabled           bool    `yaml:"enabled"`
	TTL               YAMLDur `yaml:"ttl"`
	Semantic          bool    `yaml:"semantic"`           // enable embedding-similarity cache tier
	SemanticThreshold float64 `yaml:"semantic_threshold"` // min cosine similarity for a hit (default 0.92)
}

// ProviderConfig selects the active LLM backend.
// Backend may be "ollama", "lmstudio", or "remote".
// When backend is "remote" (or a chat request specifies provider=<name>),
// the matching entry in Remotes is used.
type ProviderConfig struct {
	Backend       string         `yaml:"backend"`
	OllamaURL     string         `yaml:"ollama_url"`
	LMStudioURL   string         `yaml:"lmstudio_url"`
	AirLLMURL     string         `yaml:"airllm_url"`
	Model         string         `yaml:"model"`
	RemoteDefault string         `yaml:"remote_default"`
	Remotes       []RemoteConfig `yaml:"remotes"`
	// VisionModel is the multimodal model id for the image→text bridge (KMA-117).
	// Required for remote vision when local backends are unavailable or lack vision.
	// Example: google/gemini-2.0-flash-001 on OpenRouter, or an NVIDIA NIM vision id.
	VisionModel string `yaml:"vision_model,omitempty"`
	// VisionRemote names which configured remote runs vision (empty → remote_default).
	VisionRemote string `yaml:"vision_remote,omitempty"`
}

// RemoteConfig describes one OpenAI-compatible hosted provider.
// The API key comes from APIKey (direct, e.g. from the secrets file) if set,
// otherwise from the APIKeyEnv environment variable.
//
// Provider/KeyLabel/AddedAt/ExpiresAt/Active support the External API Keys
// settings panel: several RemoteConfig entries can share the same Provider
// (vendor label, e.g. "openai") — each is still dispatched by its own unique
// Name, but Active marks which one of a Provider's keys is the one actually
// meant to be used ("store many, one active" — see
// llpcodefeature_Cursorclone.md §CFG.20-23). ExpiresAt is optional; an
// expired key is excluded from live dispatch (internal/dispatch.Remote) but
// stays listed/editable in the UI rather than silently disappearing.
type RemoteConfig struct {
	Name      string `yaml:"name"`
	BaseURL   string `yaml:"base_url"`
	APIKeyEnv string `yaml:"api_key_env"`
	APIKey    string `yaml:"api_key"` // direct key (secrets file / runtime add)
	Model     string `yaml:"model"`
	Type      string `yaml:"type"` // "openai_compat" for now; reserved for future native adapters

	Provider    string     `yaml:"provider,omitempty"`     // vendor label for grouping, e.g. "openai", "openrouter", "groq"
	VisionModel string     `yaml:"vision_model,omitempty"` // per-key vision override (optional)
	KeyLabel    string     `yaml:"key_label,omitempty"`    // random 4-digit id shown in the UI instead of the raw key, e.g. "4821"
	AddedAt     time.Time  `yaml:"added_at,omitempty"`     // when this key was added
	ExpiresAt   *time.Time `yaml:"expires_at,omitempty"`   // optional; nil = never expires
	Active      bool       `yaml:"active,omitempty"`       // the one key in its Provider group meant to be used
	Tier        string     `yaml:"tier,omitempty"`         // user-declared cost tier: "free", "paid", "capped", or "" (unspecified)
}

// KeyTiers are the accepted RemoteConfig.Tier values. The tier is declared by
// the user when adding a key — nothing is inferred or enforced, since there is
// no uniform way to read a provider's real plan/quota — and drives
// cheapest-first routing when no explicit default remote is set:
// free > capped > paid > unspecified. See dispatch.Remote.Default.
var KeyTiers = []string{"free", "paid", "capped"}

// ValidTier reports whether t is an accepted Tier value ("" = unspecified is
// allowed).
func ValidTier(t string) bool { panic("fake") }

// Expired reports whether this key's optional expiry date has passed.
func (rc RemoteConfig) Expired() bool { panic("fake") }

// Secrets holds runtime-added provider credentials, persisted separately from
// the main config so keys stay out of config.yaml.
type Secrets struct {
	Remotes []RemoteConfig `yaml:"remotes"`
	Default string         `yaml:"default,omitempty"` // overrides provider.remote_default
}

// LoadSecrets reads the secrets file; a missing file yields an empty set.
func LoadSecrets(path string) (*Secrets, error) { panic("fake") }

// SaveSecrets writes the secrets file with 0600 perms (owner-only).
func SaveSecrets(path string, s *Secrets) error { panic("fake") }

type PlannerConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Model    string `yaml:"model"`
	Provider string `yaml:"provider"` // route planning LLM calls here ("" = local); e.g. "nvidia" for speed
}

type MemoryConfig struct {
	Enabled bool   `yaml:"enabled"`
	Dir     string `yaml:"dir"`
	// Backend selects the retrieval index: "bm25" (default, no Ollama) or
	// "ollama" (embedding cosine search via Provider.OllamaURL).
	Backend string `yaml:"backend"`
}

// WorkspaceConfig controls server-side file access for the /code IDE.
type WorkspaceConfig struct {
	Root         string   `yaml:"root"`
	MaxFileBytes int      `yaml:"max_file_bytes"`
	DenyPaths    []string `yaml:"deny_paths"`
	// Recent holds the most recently opened workspace roots, newest first.
	Recent []string `yaml:"recent,omitempty"`
	// Trusted holds workspace roots the user has explicitly trusted (agent
	// auto-apply / write tools are only enabled for a trusted root), mirroring
	// VS Code-style workspace trust.
	Trusted []string `yaml:"trusted,omitempty"`
}

// CodeConfig tunes tab completions and agent mode for small local contexts.
type CodeConfig struct {
	CompletionContextLines int    `yaml:"completion_context_lines"`
	AgentMaxTurns          int    `yaml:"agent_max_turns"`        // 0 = no fixed cap (safety cap only)
	AgentApplyMode         string `yaml:"agent_apply_mode"`       // auto_apply | propose
	AgentSafetyMaxTurns    int    `yaml:"agent_safety_max_turns"` // infinite-loop guard
	IndexOnStartup         bool   `yaml:"index_on_startup"`
	ContextTokens          int    `yaml:"context_tokens"`
	ReserveTokens          int    `yaml:"reserve_tokens"`
	// StreamTTFT is the time-to-first-token watchdog for streaming inference
	// (agent loop and Ask stream). 0 = 60s default.
	StreamTTFT YAMLDur `yaml:"stream_ttft"`
	// StreamIdle is post-first-token silence before the server cancels a
	// stream (KMA-61). 0 = dispatch.DefaultStreamIdle (3m). Negative duration
	// disables the idle watchdog.
	StreamIdle YAMLDur `yaml:"stream_idle"`
	// RequestTimeout is the wall-clock cap for one Ask/Agent/Plan/Debug run.
	// 0 = dispatch.RemoteRequestTimeout (10m).
	RequestTimeout YAMLDur `yaml:"request_timeout"`
	// ToolProtocol selects how the agent asks models to call tools:
	// "auto" (default) = openai for remote providers, json otherwise;
	// "openai" = native tools/tool_calls; "json" = custom JSON plan in content.
	ToolProtocol string `yaml:"tool_protocol"`
}

// ToolProtocolOpenAI / ToolProtocolJSON / ToolProtocolAuto are CodeConfig.ToolProtocol values.
const (
	ToolProtocolAuto   = "auto"
	ToolProtocolOpenAI = "openai"
	ToolProtocolJSON   = "json"
)

// ResolveToolProtocol returns the effective protocol for an agent turn.
// provider is the resolved route name (remote name or "" / "local" / "hybrid").
// remotes is the configured remote provider name set used when protocol is auto.
func ResolveToolProtocol(protocol, provider string, remotes []string) string { panic("fake") }

// auto — OpenAI native tools for remotes (incl. OpenRouter) and for
// hybrid/auto/local/empty when an OpenRouter remote is configured
// (KMA-203). Pure local stacks without OpenRouter stay on JSON.

// remotesIncludeOpenRouterName reports whether any configured remote name is
// an OpenRouter vendor id (openrouter / openrouter-*). Kept in config to avoid
// an import cycle with dispatch.
func remotesIncludeOpenRouterName(names []string) bool { panic("fake") }

func providerLooksLikeOpenRouter(provider string) bool { panic("fake") }

// YAMLDur wraps time.Duration for YAML unmarshalling.
type YAMLDur struct{ time.Duration }

func (d *YAMLDur) UnmarshalYAML(v *yaml.Node) error { panic("fake") }

// yaml.Marshal on embedded time.Duration emits {duration: ...}; accept both.

func (d YAMLDur) MarshalYAML() (interface{}, error) { panic("fake") }

func Load(path string) (*Config, error) { panic("fake") }

// Save writes cfg back to its Path (or to path, if given). Used to persist
// runtime changes made from the IDE — e.g. the workspace root chosen via
// Open Folder, or a newly trusted workspace.
func Save(path string, cfg *Config) error { panic("fake") }

func defaults() *Config { panic("fake") }

// 8 MiB so chat image paste/attach (KMA-102) works for typical screenshots.
