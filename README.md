# chidori

![macOS](https://img.shields.io/badge/macOS-Apple%20Silicon-000000?logo=apple)
![Windows](https://img.shields.io/badge/Windows-x64-0078D4?logo=windows)
![Release](https://img.shields.io/github/v/release/xdutsuay/chidori?label=release)
![v0.5.0](https://img.shields.io/badge/latest-v0.5.0-informational)

> **This repository is a public release placeholder only.** It hosts pre-built
> downloads, screenshots, and product documentation. **There is no source code here**
> — do not expect to clone and build from this repo.

**chidori** is a native desktop IDE with a built-in reasoning engine: Ask, Agent, Plan,
and Debug modes, real language intelligence, an integrated terminal, and a Git panel —
on a local-first inference stack that routes between your own machines and hosted
providers.

One native app per platform. No Python, no Ray, no Electron.

**Latest public binary: [v0.5.0](https://github.com/xdutsuay/chidori/releases/tag/v0.5.0).**
Code freeze through **15 September 2026**; next build expected early October.

[**Download v0.5.0 →**](https://github.com/xdutsuay/chidori/releases/tag/v0.5.0)
· [Product page](https://kaustubhtripathi.com/public/lab/lclreason/)
· [Feature inventory](docs/FEATURES.md)
· [Demo videos](docs/DEMO.md)

---

## Why this exists

Most AI coding tools assume you're either fully local (no hosted-model access) or fully
cloud (every token leaves your machine). chidori treats "where does this request run" as a
routing decision, not an architectural commitment — a laptop with no GPU and zero
attached nodes still gets fast answers by routing to a hosted key, and a LAN with a
couple of Ollama boxes gets used automatically once nodes are attached, without changing
how you interact with the IDE.

---

## See it in action

Short demos (link placeholders — replace with hosted URLs before merge; **do not** commit
video binaries to this repo). Full showcase layout: [docs/DEMO.md](docs/DEMO.md).

| Demo | Length | Status |
|---|---|---|
| **Ask mode** — editor + chat side by side, streaming answer | ~4 min | `[Ask mode demo URL]` |
| **Highlight reel** — curated clips from the full UI tour | ~2–3 min (recommended) | `[Highlight reel URL]` |
| Full UI tour (archive / internal) | ~40+ min | Too long for the homepage — use highlight reel |
| Linux reliability / fixes dogfood | long-form | Engineering tape — not a public ship claim |

> Public desktop packages today are **macOS arm64** and **Windows x64**. Linux GUI is
> exercised in CI / dogfood; a public Linux zip is **not** part of v0.5.0.

---

## Download v0.5.0

Pre-built binaries ship from
[**Releases**](https://github.com/xdutsuay/chidori/releases). Pick your platform:

| Platform | Artifact |
|---|---|
| macOS (Apple Silicon) | [`chidori-macos-arm64-v0.5.0.zip`](https://github.com/xdutsuay/chidori/releases/download/v0.5.0/chidori-macos-arm64-v0.5.0.zip) — unzip, then open `chidori.app` |
| Windows (x64) | [`chidori-windows-amd64-v0.5.0.zip`](https://github.com/xdutsuay/chidori/releases/download/v0.5.0/chidori-windows-amd64-v0.5.0.zip) — unzip, then run `chidori.exe` |
| Android companion | APK from [`xdutsuay/chidori-nagasa`](https://github.com/xdutsuay/chidori-nagasa/releases) (nagasa) |

```bash
# macOS
curl -LO https://github.com/xdutsuay/chidori/releases/download/v0.5.0/chidori-macos-arm64-v0.5.0.zip
unzip chidori-macos-arm64-v0.5.0.zip

# Windows (PowerShell)
curl.exe -LO https://github.com/xdutsuay/chidori/releases/download/v0.5.0/chidori-windows-amd64-v0.5.0.zip
```

### macOS Gatekeeper (honest note)

chidori is **not code-signed or notarized** in v0.5.0. On first launch, macOS may block
the app. That is Gatekeeper's normal response to unsigned apps — **not** a malware scan
result.

**Fix:** right-click `chidori.app` → **Open** → confirm **Open**. After that, double-click
works normally. Signing / notarization are on the public roadmap (no date yet).

### System requirements

- **macOS** — 10.10+, Apple Silicon only for the published zip
- **Windows** — 64-bit Windows 7+
- **Android companion** — phone/tablet that can sideload APKs
- **Optional** — [Ollama](https://ollama.com), [LM Studio](https://lmstudio.ai), or another OpenAI-compatible server for local inference

---

## What's new in v0.5.0

High-reliability desktop checkpoint (macOS arm64 + Windows x64), plus the Android
companion APK. Aligned with the
[product page](https://kaustubhtripathi.com/public/lab/lclreason/) story:

- **Stop actually stops** — shared Ask/Agent/Plan/Debug deadlines (TTFT, stream-idle,
  total-run) and packaged Stop cancel; runs no longer hang after you hit Stop.
- **Oversize prompts fail loudly** — fail-closed context budget: too-large prompts
  return HTTP 400 unless you opt into truncate, instead of silently dropping code.
- **Faster re-index after edits** — incremental workspace RAG skips unchanged files,
  replaces edits, and drops deletes.
- **Broken tools error out** — Grok ACP rejects terminal RPCs instead of stub-succeeding;
  MCP is fail-closed. Zen / secondary sidebar landed too.
- **Android companion APK** — pair over LAN, monitor runs, remote chat, optional
  phone-as-node ([chidori-nagasa](https://github.com/xdutsuay/chidori-nagasa)).

### What's next (not in this binary)

| Status | Item |
|---|---|
| Shipped | Android companion APK |
| Next | CLI inside the packaged app (`chidori ask` on PATH from the `.app`) |
| Later | Signed / notarized macOS builds (Gatekeeper no longer needs right-click Open) |

---

## Features (product inventory)

A product-level inventory of what chidori ships today (as of **v0.5.0**). Not every row
is VS Code–parity deep; this is what the packaged IDE actually offers. Expanded detail:
[docs/FEATURES.md](docs/FEATURES.md).

### Agent & chat

- **Modes:** Ask, Agent, Plan, and Debug — plus **All-Agent**, a chat-first layout with
  task rail, build status, context rail, and local/cloud build toggle.
- **Tool-calling agent loop** with turn budget, Continue-from-summary, per-edit diffs
  (Accept / Reject / Accept All), and Plan mode that proposes instead of auto-applying.
- **Chain-of-thought / prompt profiles** — selectable system-prompt library; mandatory
  reasoning before tools; read-before-write and verify-your-work rules.
- **Context mentions** resolved at send time: `@file`, `@folder:`, `@symbol:`,
  `@codebase`, plus `@web` / `@docs` where configured.
- **Subagent delegation** — nested research loops (read-focused) and async-write paths
  with file leases so parallel work does not clobber the same paths.
- **Sessions** — multi-turn SQLite persistence, multiple chat threads, history search,
  transcript export, compact/summarize long threads.
- **Chat UX** — streaming replies, copy / insert / apply code blocks, live step checklist,
  context-window usage breakdown, stop generation.
- **Rules, slash commands, and Skills** — project rules (`.lclreason/rules.md` /
  `.cursorrules`), Settings → Commands, Save as Skill templates.
- **MCP client** — external MCP servers as tools alongside the built-in tool registry
  (fail-closed in v0.5.0).

### Editor (Monaco)

- Syntax highlighting, multi-cursor, folding, bracket matching, sticky scroll.
- Per-path model cache (tab switches keep undo history).
- Inline ghost-text / FIM completion, format-on-save, organize imports.
- Inline diagnostics, hover, peek definition, rename, quick fixes (lightbulb).
- Side-by-side / inline diff for agent edits; split editor and editor groups.
- **Vim** and **Emacs** keybinding modes.
- Tabs: dirty indicator, pin, close, split; breadcrumb path; status bar (line/col,
  language, branch, problems count).

### Language intelligence (Go via gopls)

- Diagnostics, hover, go-to-definition, find references, implementations.
- Multi-file rename (preview → apply), quick fixes, Organize Imports, Format Document.
- Go to Symbol in file and across the workspace.

### Workspace & navigation

- File tree with lazy folders, icons, filter, Open Editors, drag-and-drop move,
  New/Rename/Delete, copy path, Reveal in Finder, drag-to-chat `@mention`.
- Respects `.gitignore` / deny paths; Collapse All.
- Full-text search and replace-in-files (preview, case, include/exclude globs).
- Command palette, fuzzy Quick Open, Find / Replace in editor and across files.

### Terminal, Git, debug

- Integrated multi-tab terminal (xterm).
- Source Control: status, diff, stage/unstage, commit, branch switch/create; All-Agent
  working bar with diff stats.
- Run & Debug panel (scoped v1 via Delve DAP for Go).

### Settings, workflows, diagnostics

- Full Settings UI: workspace root / trust, provider & model persistence, apply mode,
  turn safety caps, memory/index, compaction budgets, keybindings editor.
- **settings.json** import/export; secrets and hosted API keys (multi-key, expiry,
  verify, cost tier, model picker from key capabilities).
- Attach / discover LAN worker nodes; node dashboard and developer metrics.
- Diagnostics panel (planner + hang routing clocks + codebase LOC stats).
- **Workflows** — YAML under `.lclreason/workflows/`, headless engine (shell, LLM,
  condition, approval, loop), triggers (save / cron / commit / chat command), REST API,
  Settings panel, visual canvas.
- Optional **harness visualizer** — SSE event stream, Diagnostics trace, JSONL replay.
- **Companion** settings — desktop listener (default port `8027`) for the Android APK.

### Inference & dispatch

- Local: Ollama, LM Studio, AirLLM (this machine or LAN).
- Hosted: any OpenAI-compatible provider (Anthropic, OpenAI, Groq, NVIDIA NIM,
  OpenRouter, …).
- **Hybrid** routing races local vs remote and keeps the first answer.
- Vector / BM25 memory for `@codebase`-style retrieval; Hermes-style tool protocol
  support in the agent loop.
- Incremental workspace RAG re-index after edits (v0.5.0).

### App shell & packaging

- Native desktop app (Wails) for **macOS (Apple Silicon)** and **Windows (x64)**;
  Linux GUI build exercised in CI / dogfood (no public Linux zip in v0.5.0).
- Application menus (File / Edit / Selection / View / Go / Run / Terminal / Help),
  New Window, Open Folder, deep link `lclreason://`.
- Themes: dark, light, system follow; editor + UI font size; ligatures.
- Offline Monaco (no CDN); binary obfuscation (garble) on release builds.
- First-run config into App Support paths; coordinator can run embedded or headless.

### Explicitly not in this release

- Extensions marketplace
- Code signing / notarization (macOS Gatekeeper workaround still required)
- Auto-update (Sparkle / equivalent)
- Public Linux desktop package

---

## AI reasoning modes

| Mode | What it's for |
|---|---|
| **Ask** | Direct Q&A with codebase context; streaming answers beside the editor |
| **Agent** | Autonomous coding tasks with tool loop, diffs, Accept / Reject |
| **Plan** | Architecture & investigation — proposes instead of auto-applying |
| **Debug** | Error diagnosis & fixes (pairs with Run & Debug / Delve for Go) |

All four modes share the same inference coordinator and local ↔ hybrid ↔ remote routing.

---

## How routing works

```
chidori (native desktop IDE)
  │
  ├── Monaco editor  ──┬── gopls (Go intelligence)
  │                    └── Integrated terminal
  │
  └── Agent chat  ──►  Tool-calling agent loop
                          │
                          ▼
                       Dispatch layer
                          │  local ↔ hybrid ↔ remote routing
                    ┌─────┴─────┐
              Local nodes    Hosted providers
          (Ollama/LM Studio/   (Anthropic/OpenAI/Groq/
           AirLLM, LAN or       NVIDIA NIM/OpenRouter/…)
           this machine)
```

**Local** sends requests to a node on your machine or LAN. **Hybrid** races a local leg
against a hosted leg and keeps whichever answers first. **Remote** uses a hosted provider
directly — useful when you have no local GPU.

Configure providers and nodes in **Settings → Inference Source**. The agent picks routing
based on your settings and prompt difficulty.

### Supported providers

| Type | Examples |
|---|---|
| Local | Ollama, LM Studio, AirLLM |
| Hosted (OpenAI-compatible) | Anthropic, OpenAI, Groq, NVIDIA NIM, OpenRouter, and others |

Add multiple hosted keys in Settings; each can have an optional expiry and cost tier
(free / capped / paid) that influences default routing.

---

## Screenshots

Real screenshots from a live coordinator session (also on the
[product page](https://kaustubhtripathi.com/public/lab/lclreason/)).

### Agent mode

![Agent mode — full IDE with agent chat](docs/screenshots/agentmode-inallagentwindow-halfbacked.png)

### Ask mode (model thinking + task tracking)

![Ask mode with model thinking and task ID](docs/screenshots/askmode_modelthink_taskid.png)

### Chat

![Chat streaming in the IDE](docs/screenshots/chat.png)

### Debug mode

![Debug mode activity view](docs/screenshots/debug.png)

### Distributed inference

![Distributed inference panel](docs/screenshots/dpd.png)

### IDE overview

![chidori IDE overview](docs/screenshots/UPD.png)

---

## Companion mobile app

Pair an Android phone as companion or optional inference node.

- **APK:** debug-signed sideload from
  [`xdutsuay/chidori-nagasa` Releases](https://github.com/xdutsuay/chidori-nagasa/releases)
  (source is public / MIT).
- **Desktop:** Settings → Companion must be listening (default port **8027**).
- **Phone:** Settings → Chidori Desktop — discover via mDNS or type `host:8027`.
- **Capabilities:** monitor runs, remote chat through desktop models, optional
  phone-as-node (on-device GGUF).

Desktop IDE source stays private — binaries only via this repo.

---

## Getting started

1. **Download** [v0.5.0](https://github.com/xdutsuay/chidori/releases/tag/v0.5.0) for
   your platform (links above).
2. **Launch** chidori and use **File → Open Folder…** to pick a project workspace.
3. **Configure inference** — open **Settings → Inference Source**:
   - **Local:** run [Ollama](https://ollama.com) or [LM Studio](https://lmstudio.ai) on
     this machine or another machine on your LAN, then use **Scan LAN** or attach by IP.
   - **Hosted:** add an API key for any OpenAI-compatible provider (Anthropic, OpenAI,
     Groq, NVIDIA NIM, OpenRouter, …). A hosted key works with zero local nodes attached.
4. **Start chatting** — pick a mode (Ask, Agent, Plan, or Debug) and send a message.
   Use `@file`, `@folder`, or `@codebase` to pull workspace context into the prompt.

### Optional: Go language features

For full Go LSP support inside the IDE, install `gopls` once:

```bash
go install golang.org/x/tools/gopls@latest
```

Restart chidori after installing. Other languages use Monaco's built-in editing; Go gets
diagnostics, go-to-definition, rename, and quick fixes through gopls.

### Project rules & skills

- Drop a `.lclreason/rules.md` or `.cursorrules` file in your project root — chidori
  injects it into every agent request automatically.
- Define reusable slash commands from **Settings → Commands**.
- After a useful agent run, use **Save as Skill** to turn it into a per-workspace prompt
  template you can invoke again.

---

## Feedback

Found a bug or want a feature? Open an issue on this repo:
[github.com/xdutsuay/chidori/issues](https://github.com/xdutsuay/chidori/issues)

Product site & changelog:
[kaustubhtripathi.com/public/lab/lclreason/](https://kaustubhtripathi.com/public/lab/lclreason/)

Watch [GitHub Releases](https://github.com/xdutsuay/chidori/releases) for new binaries
(no mailing list yet).

---

## License

Documentation in this repository is MIT-licensed (see [LICENSE](LICENSE)).

chidori is distributed as pre-built binaries only. **Source code is not published.**
