# Devlog Philosophy

A **one‑page contract** for anyone using Devlog CLI. Stick to these rules; never debate formatting during a stand‑up again.

---

## Core Principles

1. **Log only what ********************\*\*\***********************you**********************\*********************** touched today\*\* — idle issues live in Jira, not here.
2. **Tasks track effort; PRs track code artefacts.** Never mix them.
3. **Inbox is a brain‑dump, not a promise.** Promote items or delete them tomorrow.
4. **One level of indentation.** Anything deeper becomes tomorrow’s top‑level task.
5. **Done work leaves the active file.** The archive keeps history; the daily file stays light.

---

## Daily Template

```md
# YYYY‑MM‑DD

## Tasks

- 👉 Build `/api/users` pagination @wip
  - ✅ Fix 500 on auth callback @done
- 🛑 Investigate race‑condition in `SurveyBuilder`@blocked
- Revise #123 after reviewer feedback @todo

## PRs

- #123 Users pagination endpoint @o
- #124 Hot‑patch auth callback @m
- #125 Typo fix in docs @om

## Inbox

- 🐞 Flaky e2e on Safari (upload spec)
- 💡 Prefetch next page in pagination hook
- ❓ Confirm legal needs for data‑export format
```

### Status Tags for **Tasks**

| Tag                 | Emoji | Meaning             | Removes line when archive runs? |
| ------------------- | ----- | ------------------- | ------------------------------- |
| `@todo` _(default)_ | –     | Planned for today   | **No**                          |
| `@wip`              | 👉    | Actively coding     | **No**                          |
| `@blocked`          | 🛑    | External dependency | **No**                          |
| `@done`             | ✅    | Finished today      | **Yes**                         |

### PR Tags (one per line)

| Tag   | Meaning                      | Appears when     | Vanishes |
| ----- | ---------------------------- | ---------------- | -------- |
| `@o`  | **Opened**                   | PR created today | Next day |
| `@m`  | **Merged**                   | PR merged today  | Next day |
| `@om` | **Opened & Merged same day** | Same day         | Next day |

> **Review changes requested?** Add a new _Task_ — e.g. `- Revise #123 after feedback`.

---

## Subtask Rules

1. One indent level only.
2. Parent is `@done` only when every child is `@done`.
3. Child still open at end of day ⇒ parent line remains.
4. Child spills into tomorrow ⇒ promote to top‑level Task.

---

## Archive Behaviour (MVP)

When you run `devlog archive`:

1. **Snapshot** the full daily file to `history/YYYY‑MM‑DD.md`.
2. Delete every line tagged `@done` from **Tasks**.
3. Wipe the entire **PRs** section (fresh list tomorrow).
4. Leave everything else unchanged.

---

## Cheat‑Sheet

| Action               | What to write                                              |
| -------------------- | ---------------------------------------------------------- |
| Start coding         | Add a line in **Tasks** (`@wip` tag optional)              |
| Finish the line      | Replace tag with `@done`                                   |
| Blocked on something | Replace tag with `@blocked` and jot why                    |
| Open a PR            | Add to **PRs** with `@o`                                   |
| Merge a PR           | Either new line `@m` **or** change existing `@o` to `@m`   |
| Random idea / bug    | Bullet in **Inbox**                                        |
| Next morning         | Promote Inbox bullets or delete them, clear mental clutter |

Follow this and your logs stay _dead simple_ yet endlessly useful.
