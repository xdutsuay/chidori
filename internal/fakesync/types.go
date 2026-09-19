package fakesync

// Options configures a fake-sync run from SourceRoot into DestRoot.
type Options struct {
	SourceRoot string   // lclreason root
	DestRoot   string   // output root (e.g. temp or chidori checkout)
	Allowlist  []string // relative dirs; empty/nil = entire SourceRoot tree
	Denylist   []string // path substrings; prefer DefaultDenylist()
	Selected   []string // relative slash paths or prefixes; non-empty limits mapped files
	DryRun     bool
}

// Mapping records one source file and its destination path.
type Mapping struct {
	Src string
	Dst string
}

// Result summarizes a Sync run.
type Result struct {
	Mappings []Mapping
	Written  int // 0 if DryRun
}
