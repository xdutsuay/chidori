# Demo showcase

Public demo slots for the chidori product page and GitHub README. **Do not commit
video binaries** into `xdutsuay/chidori` — host elsewhere (GitHub Release assets,
CDN, YouTube/Vimeo unlisted, etc.) and paste URLs below.

Product freeze context: public desktop binary is **v0.5.0** through
**15 September 2026**. Site:
[kaustubhtripathi.com/public/lab/lclreason/](https://kaustubhtripathi.com/public/lab/lclreason/).

Local source tapes live on the shared box under `/workspace/chidori-recordings/`
(see [ASSET_MAP.md](../ASSET_MAP.md)). Paths below are **box-local**, not repo paths.

---

## Homepage / README embeds (recommended)

Replace each `[…]` placeholder with a public URL before merging the docs PR.

### 1) Ask mode (~4 min) — primary short demo

| Field | Value |
|---|---|
| Local file | `/workspace/chidori-recordings/ask-mode-20260907-170153.mp4` (~5.7 MB, ~3.9 min) |
| Captured | 2026-09-07 |
| Shows | Ask mode with editor + AI chat side by side; streaming reply; model thinking / task tracking |
| Public URL | `[ASK_MODE_DEMO_URL]` |
| Suggested README caption | Ask mode — editor + chat side by side |

Embed sketch (once hosted):

```html
<!-- optional HTML embed -->
<video controls poster="docs/screenshots/askmode_modelthink_taskid.png"
       src="[ASK_MODE_DEMO_URL]"></video>
```

Markdown link form:

```markdown
[▶ Ask mode demo (~4 min)]([ASK_MODE_DEMO_URL])
```

---

### 2) Highlight reel (~2–3 min) — recommended homepage hero

| Field | Value |
|---|---|
| Source tape | `/workspace/chidori-recordings/chidori-full-ui-test-20260907.mp4` (~36 MB, ~42.5 min) |
| Segments | `/workspace/chidori-recordings/segments/` (43 clips; see `SEGMENT_INDEX.txt`) |
| Status | **Not cut yet** — do **not** link the raw 40+ min tour on the homepage |
| Public URL | `[HIGHLIGHT_REEL_URL]` |
| Suggested cuts | Menus & chrome → Open Folder / explorer → Settings (Inference Source) → Ask stream → Agent / Plan / Debug mode switch → terminal / SCM empty-states honestly → Usage / Diagnostics |

```markdown
[▶ chidori highlight reel (~2–3 min)]([HIGHLIGHT_REEL_URL])
```

---

### 3) Full UI tour (archive only)

| Field | Value |
|---|---|
| Local file | `/workspace/chidori-recordings/chidori-full-ui-test-20260907.mp4` |
| Duration | ~42.5 min |
| Use | Internal diligence / highlight-reel source of truth |
| Public URL | `[FULL_UI_TOUR_URL]` — optional unlisted; **omit from README hero** |

---

### 4) Linux reliability / fixes dogfood (engineering tape)

| Field | Value |
|---|---|
| Local file | `/workspace/chidori-recordings/linux-fix-demo-20260909.mp4` (~45 MB, ~54 min) |
| Captured | 2026-09-08 / 09 (box local) |
| Shows | Linux GUI dogfood: terminal PTY recovery, Ctrl+W tab-close vs quit, companion port fallback, debug/dlv messaging, keymap polish |
| Public URL | `[LINUX_FIXES_DEMO_URL]` — optional; **not** a claim that v0.5.0 ships a public Linux zip |
| Notes | Fixes were local on `/workspace/lclreason` dogfood builds. Public Releases remain macOS arm64 + Windows x64 only. |

---

## Screenshot stills (already in repo)

These PNGs ship under `docs/screenshots/` and should stay linked from README:

| File | Suggested label |
|---|---|
| `docs/screenshots/askmode_modelthink_taskid.png` | Ask mode — thinking + task ID |
| `docs/screenshots/agentmode-inallagentwindow-halfbacked.png` | Agent / All-Agent layout |
| `docs/screenshots/chat.png` | Chat streaming |
| `docs/screenshots/debug.png` | Debug mode |
| `docs/screenshots/dpd.png` | Distributed / inference routing |
| `docs/screenshots/UPD.png` | IDE overview |

---

## Publishing checklist

- [ ] Host Ask-mode mp4 → fill `[ASK_MODE_DEMO_URL]`
- [ ] Cut highlight reel from full tour / segments → fill `[HIGHLIGHT_REEL_URL]`
- [ ] Decide whether full tour / Linux tape are unlisted or stay private
- [ ] Update README “See it in action” table with real links
- [ ] Update product page “See it in action” if desired (cutebot / site owners)
- [ ] Confirm **no** `.mp4` files are added to the GitHub PR
