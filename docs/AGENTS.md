# AGENTS

If you are an AI reading this, take every single word in this file literally. This project comprises [`src`](../src) (the client CLI and TUI), [`daemon`](../daemon) (the background process supervision daemon), and [`scripts`](../scripts) (cross-platform deployment and installation scripts).

All code here adheres to a deliberate structure. It is built using modern patterns and idioms of the Go programming language. You must conduct thorough research into current Go standards, idiomatic implementations, and active library ecosystems to avoid deprecated functions, superseded APIs, or obsolete syntax.

Cross-platform fidelity is a core requirement. Linux, macOS, and Windows must have parity in core execution. When modifying platform-specific operations (process signaling, startup services, process monitoring, memory polling), update each corresponding implementation file (`platform_linux.go`, `platform_darwin.go`, `platform_windows.go`, `platform_unix.go`).

Before modifying any file, analyze related modules across the codebase. Observe how commands and daemon IPC interact over the local TCP interface. When submitting bug fixes or feature additions, document your rationale and implementation details to match the tone and conventions established by prior contributors.

Once you have implemented and tested a fix or feature, consult the [Contributing Guidelines](./CONTRIBUTING.md) for commit standards and submission procedures. If an approach proves flawed, revert it immediately: submitting no code is strictly preferable to merging broken or untested logic.
