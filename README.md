# Devlog CLI

_A zero‑friction command‑line tool for keeping a private developer journal right inside every project._
Written in **Go 1.24**, using only the standard library — no third‑party dependencies, no external services.

---

## Why Devlog?

| Pain                                 | Devlog cure                                                  |
| ------------------------------------ | ------------------------------------------------------------ |
| “I forget what I did yesterday.”     | One Markdown file per project, auto‑archived daily.          |
| “I don’t want yet another SaaS.”     | Local files, git‑friendly, offline‑first.                    |
| “Tools add prompts and boilerplate.” | Single `devlog today` loop; CLI disappears 95 % of the time. |

---

## Key Features (MVP)

- **`devlog init`** — bootstrap hidden `.devlog/` folder and global config.
- **`devlog today`** — open today’s entry in `$EDITOR`, auto‑inserts date header.
- **`devlog archive`** — snapshot the day, delete `@done` tasks, wipe today’s PR list.
- **`devlog view`** — aggregate history across projects (`--since 7d`, `--grep bug`).
- **Inbox + Tasks workflow** — see the [Philosophy](PHILOSOPHY.md) for exact rules.

No databases, no binaries to install beyond the CLI itself.

---

## 60‑Second Quick Start

```bash
# Install (needs Go in your PATH)
go install github.com/mohammed/devlog@latest

# Inside a project
cd ~/code/sada-platform

devlog init --org sada --project platform   # one‑time setup

devlog today                                # write Tasks/In­box, open/merge PRs
# … work …

devlog archive                              # end of day snapshot
```

Run `devlog view --since 7d` the next morning to skim last week’s progress.

---

## File & Folder Layout

```
project/
└─ .devlog/
   ├─ daily.md           # the only file you touch by hand
   └─ history/
      └─ 2025-05-18.md   # auto‑archived snapshots
```

Global data lives in `~/.devlog/` (config + cross‑project history). Everything is plain text; back it up with git, rsync, or Dropbox.

---

## Design Goals

1. **Human‑first** — Markdown is the API; type at the speed of thought.
2. **Standard‑library‑only** — easier audits, faster builds, fewer CVEs.
3. **Zero lock‑in** — migrate by copying a folder.
4. **Visible when summoned, invisible the rest of the time.**

---

## Roadmap (post‑MVP)

- Shell‑completion scripts (`bash`, `zsh`, `fish`).
- Optional pre‑commit hook to remind you to archive.
- Simple `devlog web` static viewer (no JS frameworks).
- Plugin system for exporting JSONL or SQLite indexes.

---

## Contributing

PRs welcome — but keep it dependency‑free. Run `go test ./...` and `go vet` before pushing.

---

## License

MIT. Your logs, your control.
