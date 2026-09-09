# Demo showcase

Beautiful, short clips for the chidori homepage and GitHub README.

**Rule:** host video elsewhere — **never** commit multi‑MB `.mp4` files into `xdutsuay/chidori`.
Use GitHub Release assets, a CDN, or an unlisted YouTube/Vimeo link, then paste the URL below.

| | |
|:--|:--|
| **Public binary** | [v0.5.0](https://github.com/xdutsuay/chidori/releases/tag/v0.5.0) |
| **Freeze** | through **15 September 2026** |
| **Product page** | [kaustubhtripathi.com/public/lab/lclreason/](https://kaustubhtripathi.com/public/lab/lclreason/) |

Source tapes live on the shared box under `/workspace/chidori-recordings/`
(see [ASSET_MAP.md](../ASSET_MAP.md) for maintainers). Paths below are **box-local**, not repo paths.

---

## Homepage slots

Replace each placeholder with a public URL before merge.

### 1 · Ask mode — primary short demo

~4 minutes · editor + chat side by side · streaming reply · model thinking / task tracking

| | |
|:--|:--|
| **Staged clip (ready)** | `/workspace/chidori-recordings/public-clips/ask-demo.mp4` |
| **Source tape** | `/workspace/chidori-recordings/ask-mode-20260907-170153.mp4` (~5.7 MB, ~3.9 min) |
| **Captured** | 2026-09-07 |
| **Poster still** | `docs/screenshots/askmode_modelthink_taskid.png` |
| **Public URL** | `https://github.com/xdutsuay/chidori/releases/download/v0.5.0/ask-demo.mp4` — *not uploaded yet* |
| **README caption** | Ask mode — editor + chat side by side |

Once hosted:

```markdown
[▶ Ask mode demo (~4 min)](https://github.com/xdutsuay/chidori/releases/download/v0.5.0/ask-demo.mp4)
```

```html
<video controls poster="docs/screenshots/askmode_modelthink_taskid.png"
       src="https://github.com/xdutsuay/chidori/releases/download/v0.5.0/ask-demo.mp4"></video>
```

---

### 2 · Highlight reel — recommended homepage hero

~2–3 minutes · curated cuts from the full UI tour

| | |
|:--|:--|
| **Source tape** | `/workspace/chidori-recordings/chidori-full-ui-test-20260907.mp4` (~36 MB, ~42.5 min) |
| **Segments** | `/workspace/chidori-recordings/segments/` (43 clips; see `SEGMENT_INDEX.txt`) |
| **Staged clip (ready)** | `/workspace/chidori-recordings/public-clips/ui-highlight.mp4` (**90s**) |
| **Public URL** | `https://github.com/xdutsuay/chidori/releases/download/v0.5.0/ui-highlight.mp4` |
| **Suggested cuts** | Menus & chrome → Open Folder → Settings (Inference Source) → Ask stream → Agent / Plan / Debug switch → terminal / SCM → Usage / Diagnostics |

```markdown
[▶ chidori highlight reel (~2–3 min)](https://github.com/xdutsuay/chidori/releases/download/v0.5.0/ui-highlight.mp4)
```

---

### 3 · Full UI tour — archive only

| | |
|:--|:--|
| **Local file** | `/workspace/chidori-recordings/chidori-full-ui-test-20260907.mp4` |
| **Duration** | ~42.5 min |
| **Use** | Internal diligence / highlight-reel source of truth |
| **Public URL** | `[FULL_UI_TOUR_URL]` — optional unlisted; **omit from README hero** |

---

### 4 · Linux reliability tape — engineering only

| | |
|:--|:--|
| **Local file** | `/workspace/chidori-recordings/linux-fix-demo-20260909.mp4` (~45 MB, ~54 min) |
| **Staged clip** | `/workspace/chidori-recordings/public-clips/linux-pass.mp4` (short pass) |
| **Shows** | Linux GUI dogfood: PTY recovery, Ctrl+W tab-close vs quit, companion port fallback, debug/dlv messaging, keymap polish |
| **Public URL** | `https://github.com/xdutsuay/chidori/releases/download/v0.5.0/linux-pass.mp4` — optional |
| **Honesty** | **Not** a claim that v0.5.0 ships a public Linux zip. Public Releases remain **macOS arm64 + Windows x64** only. |

---

## Screenshot stills (in repo)

Already under `docs/screenshots/` — linked from the README gallery:

| File | Label |
|:-----|:------|
| `UPD.png` | IDE overview |
| `askmode_modelthink_taskid.png` | Ask — thinking + task ID |
| `agentmode-inallagentwindow-halfbacked.png` | Agent / All-Agent layout |
| `chat.png` | Chat streaming |
| `debug.png` | Debug mode |
| `dpd.png` | Distributed / inference routing |

---

## Clip hosting

| Do | Don't |
|:---|:------|
| Upload to a Release asset, CDN, or unlisted video host | Commit `.mp4` into the docs PR |
| Paste the resulting HTTPS URL into README + this file | Link the raw 40+ min full tour on the homepage |
| Keep Ask demo + highlight reel as the two public slots | Market the Linux tape as a shipped Linux package |

---

## Publishing checklist

- [ ] Host `public-clips/ask-demo.mp4` → fill `https://github.com/xdutsuay/chidori/releases/download/v0.5.0/ask-demo.mp4`
- [ ] Host `public-clips/ui-highlight.mp4` → fill `https://github.com/xdutsuay/chidori/releases/download/v0.5.0/ui-highlight.mp4`
- [ ] Decide whether full tour / Linux tape are unlisted or stay private
- [ ] Update README “See it in action” with real links
- [ ] Update product page embeds if desired (cutebot / site owners)
- [ ] Confirm **no** `.mp4` files are added to the GitHub PR
