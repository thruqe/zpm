---
title: Logs & Rotation
description: Inspect real-time process output and manage automatic log rotation.
---

ZPM aggregates standard output (`stdout`) and standard error (`stderr`) for every managed process directly into dedicated log files.

## Viewing Logs

### Real-Time Streaming

Follow output in real time during startup:
```bash
zpm start "node server.js" --name api --follow
```

Or view existing logs:
```bash
zpm logs api
```

### Log Storage Location

Logs are stored in your platform's user configuration directory:
- **Linux**: `~/.config/zpm/logs/<process>.log`
- **macOS**: `~/Library/Application Support/zpm/logs/<process>.log`
- **Windows**: `%APPDATA%\zpm\logs\<process>.log`

## Automatic Log Rotation

Prevent unbounded disk growth with automatic size-based rotation:

```bash
zpm start "node server.js" --name api \
  --log-max-size 20mb \
  --log-backups 5
```

- `--log-max-size`: Maximum size before a log file is rotated (e.g. `10mb`, `500kb`).
- `--log-backups`: Number of rotated historical archive files to keep (e.g. `api.log.1`, `api.log.2`).
