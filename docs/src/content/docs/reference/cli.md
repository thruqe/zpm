---
title: CLI Reference
description: Complete command-line argument and flag reference.
---

Complete reference for all `zpm` CLI subcommands and parameters.

## Core Commands

### `zpm start`
Starts managing an application command.

```bash
zpm start <command> [options]
```

| Flag | Description | Default |
| :--- | :---------- | :------ |
| `--name <string>` | Unique name identifier | Derived from command |
| `--cwd <path>` | Working directory | `.` |
| `--instances <n>` | Cluster worker count | `1` |
| `--follow` | Stream logs directly to stdout | `false` |
| `--restart-delay <ms>` | Restart delay backoff | `1000` |
| `--max-restarts <n>` | Max restart retries (-1 for infinite) | `-1` |
| `--health <command>` | Health check shell probe | `""` |
| `--log-max-size <size>` | Size threshold before log rotation | `10mb` |
| `--log-backups <n>` | Historical log files to preserve | `5` |
| `--env <KEY=VAL>` | Set environment variable | None |

---

### `zpm list` / `zpm status`
Lists all registered processes, their current operational status, uptime, PID, and health metrics.

---

### `zpm tui` / `zpm ui` / `zpm tux`
Launches the full-screen animated terminal user interface.

---

### `zpm stop <name>`
Stops a running process or entire cluster group.

---

### `zpm restart <name>`
Restarts a process or entire cluster group.

---

### `zpm purge <name>`
Stops the process, deletes its registry record, and removes its log files.

---

### `zpm logs <name>`
Prints the output logs for the specified process.

---

## Daemon Commands

### `zpm daemon <action>`
Controls the background supervisor:
- `start`: Starts `zpmd` if not currently running.
- `stop`: Gracefully terminates `zpmd`.
- `reload`: Restarts `zpmd`.

---

## System Integration

### `zpm startup <action>`
- `install`: Registers auto-start on logon (systemd, LaunchAgent, or Task Scheduler).
- `uninstall`: Unregisters and removes auto-start service configurations.

---

## Ecosystem Management

### `zpm ecosystem start <file> [--env <profile>]`
Loads and executes declarative multi-process definitions from a `zpm.config.json` file.
