### 30-Day / 30-Hour Go + CLI ramp-up

_(≈ 1 focused hour per day, matched to the Devlog CLI milestones you already sketched)_

| Day    | Focus (learn → apply)                                                                            | Mini-deliverable                                                                                 |
| ------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| **01** | **Project boot** – install Go 1.24, `go env`, `go mod init`, `git init`, commit empty `main.go`. | Repo pushed, tooling verified.                                                                   |
| **02** | **Std library tour #1** – packages, modules, workspace layout.                                   | Add `internal/` & `cmd/` folders.                                                                |
| **03** | **flag pkg** – parsing `--help`, string / bool flags.                                            | Hard-code `--version` flag (prints v0.0.0).                                                      |
| **04** | **os + fmt** – basic `os.Args`, exit codes, `fmt.Fprintf(os.Stderr, ...)`.                       | Replace yesterday’s dummy flag logic with manual parsing once to feel the pain -> keep flag pkg. |
| **05** | **Shell basics** – `$PATH`, shebangs, chmod +x, symlinks.                                        | Add `make install` that symlinks the binary into `~/bin`.                                        |
| **06** | **File I/O intro** – `os.Open`, `os.Create`, `defer f.Close()`.                                  | Write “Hello Devlog” into `.devlog/daily.md`.                                                    |
| **07** | **path/filepath** – absolute vs relative, `Dir`, `Join`, `Base`.                                 | `devlog init` detects repo root (`git rev-parse --show-toplevel` fallback to cwd).               |
| **08** | **Permissions** – Unix file modes, `os.MkdirAll(..., 0o755)`, `umask`.                           | Ensure `.devlog` and `history` dirs created with sane perms.                                     |
| **09** | **Text scanning** – `bufio.Scanner`, line loops, error handling.                                 | Read `daily.md`, print the first 10 lines.                                                       |
| **10** | **Strings and bytes** – UTF-8 basics, `strings.TrimSpace`, `bytes.Buffer`.                       | Auto-insert today’s `# YYYY-MM-DD` header if missing.                                            |
| **11** | **Time & dates** – `time.Now()`, formatting, parsing.                                            | Generate correct ISO date header programmatically.                                               |
| **12** | **Structuring commands** – tiny `switch cmd := os.Args[1]`.                                      | Split `init` and `today` sub-commands (no third-party CLI lib).                                  |
| **13** | **Read / Write JSON** – `encoding/json`.                                                         | Prototype global `config.json` with project path registration.                                   |
| **14** | **Error wrapping** – `fmt.Errorf(\"…: %w\", err)`, `log.Fatal`.                                  | Replace all naked panics with wrapped errors.                                                    |
| **15** | **Directory walking** – `filepath.WalkDir`.                                                      | Print all files under `.devlog/history`.                                                         |
| **16** | **Markdown parsing (light)** – regex on `^## ` headings via `regexp` pkg.                        | Extract a single day block from `daily.md`.                                                      |
| **17** | **File rewriting** – read-modify-write pattern without data loss.                                | Implement “archive today”: cut block, append to `history/YYYY-MM-DD.md`.                         |
| **18** | **Environment variables** – `os.Getenv`, defaults.                                               | Allow `$DEVLOG_DATACENTER` override for history path.                                            |
| **19** | **Unit testing intro** – `*_test.go`, table-driven tests.                                        | Test the `extractDay()` function.                                                                |
| **20** | **Testing file ops** – `t.TempDir()`, golden files.                                              | Test “archive today” end-to-end in temp folder.                                                  |
| **21** | **ANSI basics** – colouring output with escape codes.                                            | Green “✔ Archived” message after running `archive`.                                              |
| **22** | **Terminal paging** – spawn `less` if `$PAGER` set (`os/exec`).                                  | Implement `devlog view` piping aggregated history into pager.                                    |
| **23** | **Search** – naive `strings.Contains` over loaded entries.                                       | Add `devlog view --grep foo`.                                                                    |
| **24** | **Cross-platform build** – `GOOS`, `GOARCH`, `go build -o`.                                      | Produce `devlog-linux-amd64` and `devlog-darwin-arm64`.                                          |
| **25** | **Go tests race detector & vet** – `go test -race`, `go vet`.                                    | Integrate into `Makefile` or simple `./run_tests.sh`.                                            |
| **26** | **Signals & exit codes** – handle Ctrl-C (`os/signal`, `syscall`).                               | Graceful cancel when viewing large history.                                                      |
| **27** | **Benchmark basics** – `testing.B` to profile parsing.                                           | Benchmark `extractDay`; identify bottleneck if any.                                              |
| **28** | **Simple pre-commit hook** – create `.git/hooks/pre-commit` via Go.                              | Warn if `@todo` tasks still open.                                                                |
| **29** | **Packaging docs** – fill in `README`, update `PHILOSOPHY.md` links, add usage examples.         | README shows real command outputs.                                                               |
| **30** | **Tag v0.1.0** – `git tag`, GitHub release, attach binaries.                                     | Self-install devlog via release asset, dog-food on another repo.                                 |

### How to use the schedule

1. **Clock in** for one focused hour. No internet rabbit holes until the micro-goal is done.
2. Open a new `## YYYY-MM-DD` section in your own `daily.md`, jot what you’ll tackle.
3. When the hour’s over, record ✅, commit, push.
4. Ask me for “Day N lesson” tomorrow; I’ll drill into the concepts and sample code.

This way you’ll finish month #1 with:

- Solid comfort in the Go std library relevant to CLI + filesystem.
- A usable v0.1 Devlog CLI already powering your own workflow.
