---
title: Quick Start
description: Get up and running with ZPM in 5 minutes.
---

This guide gets you from zero to managing running services in minutes.

## 1. Start the Daemon

ZPM operations communicate with a background supervisor daemon (`zpmd`):

```bash
zpm daemon start
```

You can verify the daemon status anytime:
```bash
zpm daemon status
```

## 2. Start Your First App

Start managing any Node.js, Bun, Python, Go, or shell process:

```bash
zpm start "bun index.ts" --name api --follow
```

Flags explained:
- `--name api`: Assigns a readable identifier for referencing the process.
- `--follow`: Streams standard output and error directly to your terminal.

## 3. Inspect Running Processes

Check status in a clean tabular view:

```bash
zpm list
```

Output includes process ID, uptime, restart count, RSS memory footprint, and health status.

## 4. Open the Live Terminal Dashboard (TUI)

Launch the interactive curses interface:

```bash
zpm tui
```

- Navigate through processes with arrow keys or `j`/`k`.
- View live uptime, resource statistics, and real-time logs.
- Press `q` or `Ctrl+C` to exit.

## 5. Control Lifecycle

```bash
# Restart application
zpm restart api

# Temporarily stop application
zpm stop api

# Restart stopped application
zpm start api

# Completely purge application and its logs
zpm purge api
```
