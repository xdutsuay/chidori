# Features — chidori v0.5.0

Detailed product inventory for the public desktop IDE. Evidence base:

- Public README historically published on `xdutsuay/chidori` (feature depth)
- Product page [kaustubhtripathi.com/public/lab/lclreason/](https://kaustubhtripathi.com/public/lab/lclreason/) (v0.5.0 story + freeze)
- GitHub Release notes for [v0.5.0](https://github.com/xdutsuay/chidori/releases/tag/v0.5.0) / [v0.4.0](https://github.com/xdutsuay/chidori/releases/tag/v0.4.0)

This is what the **packaged** IDE offers. Not every surface is VS Code–parity deep.
Do **not** treat Linux dogfood notes as public ship claims — Releases today are
**macOS arm64** and **Windows x64** only.

---

## AI modes

| Mode | Intent | Behavior highlights |
|---|---|---|
| **Ask** | Direct Q&A | Streams answers beside the editor; `@file` / `@folder` / `@codebase` context; Research toggle for deeper retrieval |
| **Agent** | Autonomous coding | Tool-calling loop, turn budget, per-edit diffs (Accept / Reject / Accept All) |
| **Plan** | Architecture & investigation | Proposes changes instead of auto-applying |
| **Debug** | Error diagnosis | Mode + Run & Debug panel; Go DAP via Delve when `dlv` is installed |
| **All-Agent** | Chat-first layout | Task rail, build status, context rail, local/cloud build toggle |

Shared across modes (v0.5.0 reliability):

- Packaged **Stop** cancel actually stops runs
- Shared deadlines: TTFT, stream-idle, total-run
- Fail-closed oversize context (HTTP 400 unless truncate opted in)

---

## Agent loop & context

- Tool-calling agent loop with Continue-from-summary
- Chain-of-thought / prompt profiles (system-prompt library; reason-before-tools;
  read-before-write; verify-your-work)
- Context mentions at send time: `@file`, `@folder:`, `@symbol:`, `@codebase`,
  plus `@web` / `@docs` where configured
- Subagent delegation — nested research (read-focused) and async-write paths with
  file leases
- Sessions — SQLite multi-turn persistence, multiple threads, history search,
  transcript export, compact/summarize
- Chat UX — streaming, copy / insert / apply code blocks, live step checklist,
  context-window usage breakdown, stop generation
- Project rules — `.lclreason/rules.md` or `.cursorrules`
- Slash commands — Settings → Commands
- Skills — Save as Skill templates per workspace
- MCP client — external servers as tools; **fail-closed** in v0.5.0
- Grok ACP — rejects unsupported terminal RPCs instead of stub-succeeding (v0.5.0)

---

## Editor (Monaco)

- Syntax highlighting, multi-cursor, folding, bracket matching, sticky scroll
- Per-path model cache (undo history survives tab switches)
- Inline ghost-text / FIM completion
- Format-on-save, organize imports
- Inline diagnostics, hover, peek definition, rename, quick fixes (lightbulb)
- Side-by-side / inline diff for agent edits
- Split editor and editor groups
- Vim and Emacs keybinding modes
- Tabs: dirty indicator, pin, close, split
- Breadcrumb path; status bar (line/col, language, branch, problems count)
- Zen / secondary sidebar (landed with the reliability checkpoint)

---

## Language intelligence

**Go (via gopls)** — first-class:

- Diagnostics, hover, go-to-definition, find references, implementations
- Multi-file rename (preview → apply)
- Quick fixes, Organize Imports, Format Document
- Go to Symbol (file + workspace)

Other languages use Monaco's built-in editing without a dedicated LSP package in
the public story.

Install once:

```bash
go install golang.org/x/tools/gopls@latest
```

---

## Workspace & navigation

- File tree: lazy folders, icons, filter, Open Editors, drag-and-drop move,
  New / Rename / Delete, copy path, Reveal in Finder, drag-to-chat `@mention`
- Respects `.gitignore` / deny paths; Collapse All
- Full-text search and replace-in-files (preview, case, include/exclude globs)
- Command palette, fuzzy Quick Open
- Find / Replace in editor and across files
- Workspace trust / untrusted-folder banner

---

## Terminal, Git, debug

- Integrated multi-tab terminal (xterm)
- Source Control: status, diff, stage/unstage, commit, branch switch/create
- All-Agent working bar with diff stats
- Run & Debug panel — scoped v1 via Delve DAP for Go (`dlv` on PATH)

---

## Inference & hybrid routing

```
Local nodes (Ollama / LM Studio / AirLLM, this machine or LAN)
        ↕ hybrid race
Hosted OpenAI-compatible (Anthropic, OpenAI, Groq, NVIDIA NIM, OpenRouter, …)
```

- **Local** — attach by Scan LAN or IP; auto-detect localhost Ollama when present
- **Remote** — hosted key with zero local nodes
- **Hybrid** — race local vs remote; keep first answer
- Multi-key hosted secrets: expiry, verify, cost tier (free / capped / paid),
  model picker from key capabilities
- Vector / BM25 memory for `@codebase`-style retrieval
- Hermes-style tool protocol support in the agent loop
- **Incremental workspace RAG** (v0.5.0) — skips unchanged files, replaces edits,
  drops deletes

Configure in **Settings → Inference Source**.

---

## Settings, workflows, diagnostics, companion

Settings surface includes (non-exhaustive): General, Workspace, Rules, Prompts,
Integrations, MCP, Workflows, Inference Source, Usage, Agent, Diagnostics,
Companion App, Editor, Terminal, Context Window, Experimental.

Also:

- `settings.json` import/export
- Attach / discover LAN worker nodes; node dashboard; developer metrics
- Diagnostics panel (planner + hang routing clocks + codebase LOC stats)
- Usage surfaces for operator visibility
- **Workflows** — YAML under `.lclreason/workflows/`; headless engine (shell, LLM,
  condition, approval, loop); triggers (save / cron / commit / chat command);
  REST API; Settings panel; visual canvas
- Optional harness visualizer — SSE event stream, Diagnostics trace, JSONL replay
- **Companion** — desktop listener (default **8027**) for Android APK pairing

---

## App shell & packaging

- Native desktop via **Wails** (Go backend + frontend) — not Electron
- Published zips: **macOS Apple Silicon**, **Windows x64**
- Linux GUI exercised in CI / internal dogfood — **no public Linux zip in v0.5.0**
- Application menus: File / Edit / Selection / View / Go / Run / Terminal / Help
- New Window, Open Folder, deep link `lclreason://`
- Themes: dark, light, system follow; editor + UI font size; ligatures
- Offline Monaco (no CDN)
- Binary obfuscation (garble) on release builds
- First-run config into App Support paths
- Coordinator can run embedded or headless

Android companion APK ships from
[`xdutsuay/chidori-nagasa`](https://github.com/xdutsuay/chidori-nagasa)
(LAN pairing, monitor, remote chat, optional phone-as-node).

---

## Explicitly not in v0.5.0

| Missing | Notes |
|---|---|
| Extensions marketplace | Not planned in this checkpoint |
| Code signing / notarization | Gatekeeper right-click Open still required on macOS |
| Auto-update | No Sparkle / equivalent yet |
| Public Linux desktop package | Dogfood / CI only |
| Packaged CLI on PATH | Listed as **next** on the product page (`chidori ask`) |

---

## What's new specifically in v0.5.0

Mirrored from the public product page / release story:

1. Stop actually stops (shared deadlines + packaged Stop cancel)
2. Oversize prompts fail loudly (fail-closed context budget)
3. Faster re-index after edits (incremental RAG)
4. Broken tools error out (Grok ACP + MCP fail-closed; Zen / secondary sidebar)
5. Android companion APK available

Freeze: through **15 September 2026**; next public build expected early October.
