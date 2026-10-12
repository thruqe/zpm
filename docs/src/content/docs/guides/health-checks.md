---
title: Health Checks & Auto-Healing
description: Configure active probes and automated crash recovery.
---

ZPM actively monitors your applications and ensures continuous uptime through automated crash handling and custom health verification commands.

## Crash Detection & Restart Backoff

By default, any process that terminates unexpectedly is immediately revived by the supervisor.

You can configure restart behavior per process:

```bash
zpm start "python app.py" --name worker \
  --restart-delay 1500 \
  --max-restarts 10
```

- `--restart-delay <ms>`: Delay in milliseconds before attempting to restart after exit (defaults to 1000ms).
- `--max-restarts <n>`: Maximum consecutive restart attempts before marking as `crashed` (-1 for infinite).

## Custom Health Probes

Provide an executable command or HTTP check with `--health`:

```bash
zpm start "bun index.ts" --name api \
  --health "curl -fsS http://127.0.0.1:3000/health"
```

### How Probes Work
1. The supervisor daemon runs the `--health` probe every 10 seconds.
2. If the probe exits with code `0`, the process is marked as healthy.
3. If the probe fails (non-zero exit code) and auto-restart is enabled, ZPM automatically restarts the process to recover from deadlock or unresponsive states.
