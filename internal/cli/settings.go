package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/xdutsuay/lclreason/internal/config"
)

// settingSpec maps a CLI settings key to UIConfig fields.
type settingSpec struct {
	Key  string
	Help string
	Get  func(u config.UIConfig) string
	Set  func(u *config.UIConfig, val string) error
}

func parseBool(val string) (bool, error) { panic("fake") }

func parseInt(val string) (int, error) { panic("fake") }

// KnownSettings is the Phase 1 allowlist (UI prefs only — no Agent/Ask).
func KnownSettings() []settingSpec { panic("fake") }

func lookupSetting(key string) (settingSpec, bool) { panic("fake") }

func loadUI(ctx context.Context, o Options) (config.UIConfig, Resolved, bool, error) { panic("fake") }

func saveUI(ctx context.Context, o Options, ui config.UIConfig) (Resolved, bool, error) {
	panic("fake")
}

func fetchUIPrefs(ctx context.Context, o Options, cfg *config.Config) (config.UIConfig, error) {
	panic("fake")
}

func putUIPrefs(ctx context.Context, o Options, cfg *config.Config, ui config.UIConfig) error {
	panic("fake")
}

// SettingsList prints known keys and current values.
func SettingsList(ctx context.Context, o Options) (string, error) { panic("fake") }

// SettingsGet returns one key's value.
func SettingsGet(ctx context.Context, o Options, key string) (string, error) { panic("fake") }

// SettingsSet updates one key (live prefs API if coordinator up, else config.yaml).
func SettingsSet(ctx context.Context, o Options, key, value string) (string, error) { panic("fake") }
