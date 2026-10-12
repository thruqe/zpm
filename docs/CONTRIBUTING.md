# Contributing Guidelines

Welcome to ZPM. These guidelines should be reviewed before submitting issues or pull requests.

## 1. Development Management

All development workflows are managed through standard Go tools:

| Command | Description |
| :------ | :---------- |
| `go fmt ./...` | Format all Go source files. |
| `go vet ./...` | Vet and run static analysis on source code. |
| `go test ./...` | Run the test suite across `src` and `daemon`. |
| `go build -o daemon/zpmd ./daemon` | Build the daemon binary. |
| `go build -o src/zpm ./src` | Build the CLI binary. |

## 2. Architecture & Guidelines

- **Architecture Guidelines**: If you are an AI assistant, please read [AGENTS](./AGENTS.md).
- **Code Quality**: Ensure code is formatted and vetted before opening a pull request.
- **Cross-Platform Compatibility**: ZPM supports Linux, macOS, and Windows. When modifying OS-specific logic, ensure builds pass on all platforms (`linux`, `darwin`, `windows`).
- **Testing**: Existing tests must pass, and unit tests should be added for new logic using `go test ./...`.
- **Commits**: Conventional commit conventions must be followed (`feat:`, `fix:`, `refactor:`, `docs:`, `test:`).

## 3. Pull Request Workflow

1. Create a feature branch from `main`.
2. Implement changes following the architecture guidelines.
3. Validate locally with `go test ./...` and verify cross-compilation (`GOOS=windows`, `GOOS=darwin`, `GOOS=linux`).
4. Submit a Pull Request targeting the `main` branch.
