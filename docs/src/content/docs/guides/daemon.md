---
title: Daemon & OS Boot
description: Daemon supervision architecture and system startup integration.
---

ZPM utilizes a decoupled architecture where client commands communicate with `zpmd`, an independent background supervisor running as a lightweight daemon.

## Managing the Daemon

```bash
# Start the supervisor daemon
zpm daemon start

# Stop the supervisor daemon
zpm daemon stop

# Restart/Reload the supervisor daemon
zpm daemon reload
```

## Running on System Boot

Configure ZPM to initialize automatically on user logon or reboot:

```bash
zpm startup install
```

### Platform Implementations
- **Linux**: Creates and enables a systemd user unit at `~/.config/systemd/user/zpmd.service`.
- **macOS**: Registers a LaunchAgent property list at `~/Library/LaunchAgents/com.zpm.daemon.plist`.
- **Windows**: Registers an ONLOGON task in the Windows Task Scheduler under `ZPM Daemon`.

To disable and clean up startup services:

```bash
zpm startup uninstall
```
