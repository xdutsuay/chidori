# chidori

![macOS](https://img.shields.io/badge/macOS-Apple%20Silicon-000000?logo=apple)
![Windows](https://img.shields.io/badge/Windows-x64-0078D4?logo=windows)
![Releases](https://img.shields.io/github/v/release/xdutsuay/chidori?label=release)

**chidori** is a cross-platform native desktop IDE with agent chat, real language
intelligence, an integrated terminal, and a Git panel — built on a local-first
inference engine that routes between your own machines and hosted providers.

One native app per platform. No Python, no Ray, no Electron.

[**Download the latest release →**](https://github.com/xdutsuay/chidori/releases/latest)

## Why this exists

Most AI coding tools assume you're either fully local (no hosted-model access) or fully
cloud (every token leaves your machine). chidori treats "where does this request run" as a
routing decision, not an architectural commitment — a laptop with no GPU and zero
attached nodes still gets fast answers by routing to a hosted key, and a LAN with a
couple of Ollama boxes gets used automatically once nodes are attached, without changing
how you interact with the IDE.

## What's in the IDE

- **Agent chat — Ask / Agent / Plan / Debug modes.** The agent runs a capped tool-calling
  loop with mandatory chain-of-thought reasoning before every tool call and explicit
  read-before-write / verify-your-fix / self-correct-on-error rules. Diffs are shown
  per-edit (Monaco side-by-side view) with Accept/Reject or Accept All; Plan mode always
  proposes rather than auto-applies. On hitting the turn budget the agent produces one
  last plain-prose progress summary instead of just stopping, with a **Continue** button
  to pick up from there.
- **Scoped subagent delegation.** The agent can spin off a self-contained research
  question into an isolated, read-only nested sub-loop — a way to investigate something
  without spending the parent's own turn budget on exploratory back-and-forth.
- **Context mentions.** `@file`, `@folder:path/`, `@symbol:name`, and `@codebase` resolve
  to real file listings and workspace content at send time — not just text hints the model
  has to guess what to do with.
- **Real Go language intelligence via gopls** — diagnostics, hover, go-to-definition,
  find-all-references, multi-file rename (preview → apply), quick fixes, Organize
  Imports, Format Document, and Go to Symbol (file + workspace).
- **Monaco editor** with a per-path model cache (tab switches don't destroy undo history),
  inline diagnostics, sticky scroll, multi-cursor, format-on-save, and a full
  Selection/View/Go menu surface.
- **Integrated terminal** — multi-tab shell sessions built into the IDE.
- **Source Control panel** — status, diff, stage/unstage, commit, branch switch and create.
- **Workspace search** — full-text search and replace-in-files (preview then apply, case
  toggle, include/exclude globs), a command palette with fuzzy Quick Open, and
  cross-session chat history search.
- **MCP client** — the agent can consume external MCP servers as tools, alongside its own
  built-in registry.
- **Rules, Commands, and Skills** — project-level rules injected into every request,
  user-definable slash commands, and "Save as Skill" to turn a finished agent run into a
  reusable per-workspace prompt template.
- **Multi-provider dispatch** — Ollama, LM Studio, AirLLM, or any OpenAI-compatible
  hosted provider (Anthropic, OpenAI, Groq, NVIDIA NIM, OpenRouter, …), with hybrid mode
  racing a local and a remote leg and keeping whichever answers first.

## Screenshots

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

## Download & install

Pre-built binaries are published on the [**Releases**](https://github.com/xdutsuay/chidori/releases)
page. Pick the archive for your platform:

| Platform | Artifact |
|---|---|
| macOS (Apple Silicon) | `chidori-macos-arm64.zip` — unzip, then open `chidori.app` |
| Windows (x64) | `chidori-windows-amd64.zip` — unzip, then run `chidori.exe` |

> **macOS Gatekeeper:** chidori is not code-signed yet. On first launch, macOS may block
> the app. Right-click `chidori.app` → **Open** → confirm **Open** in the dialog. After
> that, double-click works normally.

## Quick start

1. **Download** the latest release for your platform (link above).
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

## Supported providers

| Type | Examples |
|---|---|
| Local | Ollama, LM Studio, AirLLM |
| Hosted (OpenAI-compatible) | Anthropic, OpenAI, Groq, NVIDIA NIM, OpenRouter, and others |

Add multiple hosted keys in Settings; each can have an optional expiry and cost tier
(free / capped / paid) that influences default routing.

## Project rules & skills

- Drop a `.lclreason/rules.md` or `.cursorrules` file in your project root — chidori
  injects it into every agent request automatically.
- Define reusable slash commands from **Settings → Commands**.
- After a useful agent run, use **Save as Skill** to turn it into a per-workspace prompt
  template you can invoke again.

## Feedback

Found a bug or want a feature? Open an issue on this repo:
[github.com/xdutsuay/chidori/issues](https://github.com/xdutsuay/chidori/issues)

## License

Documentation in this repository is MIT-licensed (see [LICENSE](LICENSE)).

chidori is distributed as pre-built binaries only. **Source code is not published.**
