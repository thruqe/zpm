![zpm](https://raw.githubusercontent.com/thruqe/zpm/main/docs/public/zpm.jpg)

zpm is a small, sharp process manager for apps you want to keep alive, providing production-grade daemon supervision, health checking, automatic restart, log rotation, and cluster scaling with zero configuration hassle.

## Usage

If you want to install on your system, see the [Installation Guide](./scripts/README.md) to quickly install.

### Quick Install

**Linux / macOS (Bash):**
```bash
curl -sL https://raw.githubusercontent.com/thruqe/zpm/main/scripts/install.sh | bash
```

**Windows (PowerShell as Administrator):**
```powershell
irm https://raw.githubusercontent.com/thruqe/zpm/main/scripts/install-windows.ps1 | iex
```

**Homebrew (macOS & Linux):**
```bash
brew tap thruqe/tap
brew install zpm
```

**WinGet (Windows):**
```powershell
winget install thruqe.zpm
```

**Chocolatey (Windows):**
```powershell
choco install zpm
```

### Essential Commands

```bash
zpm start "bun index.ts" --name api --follow   # Start managing an application
zpm list                                      # List all managed processes and statuses
zpm tui                                       # Launch the live animated terminal dashboard
zpm stop api                                  # Stop a running process or cluster
zpm restart api                               # Restart a process or cluster
zpm purge api                                 # Stop and remove process entry and logs
zpm daemon start                              # Start background supervisor daemon
zpm startup install                           # Configure auto-start on system boot/login
```

### Ecosystem Configuration

Process configurations can be declared in `zpm.config.json` for reproducible deployments:

```json
{
  "apps": [
    {
      "name": "api",
      "command": "bun index",
      "cwd": ".",
      "instances": 2,
      "auto_restart": true,
      "restart_delay": 1000,
      "max_restarts": -1,
      "health_command": "curl -fsS http://127.0.0.1:3000/health",
      "log_max_bytes": 10485760,
      "log_backups": 5,
      "env": {
        "PORT": "3000"
      },
      "env_production": {
        "NODE_ENV": "production"
      }
    }
  ]
}
```

Run declarative environments with a single command:
```bash
zpm ecosystem start zpm.config.json --env production
```

## Features

### Native Cross-Platform Process Supervision
- **Self-Contained Native Daemons**: Independent background supervisor (`zpmd`) continuously managing process lifecycles without bulky runtime runtimes.
- **Full OS Platform Parity**: First-class support across Linux (amd64, arm64), macOS (Intel, Apple Silicon), and Windows (x64, ARM64).
- **Native Boot Integration**: Seamless service integration utilizing `systemd` user services on Linux, `LaunchAgent` on macOS, and Windows Task Scheduler.

### Autonomous Healing & Crash Recovery
- **Automatic Process Restarts**: Instant automatic respawn upon abnormal termination or crash with configurable restart backoff delays (`--restart-delay`) and thresholds (`--max-restarts`).
- **Periodic Health Verification**: Configurable active health probes (`--health`) executed at regular intervals with automated recovery if probes fail.
- **Graceful Lifecycle Signaling**: Orderly process shutdowns with configurable timeout escalations before forceful termination.

### Dynamic Cluster Scaling & Multi-Instance Execution
- **Zero-Config Cluster Spawning**: Spin up multiple concurrent workers with `--instances <n>` (e.g. `api-0`, `api-1`).
- **Unified Cluster Targeting**: Control and manage entire clusters collectively or individually (`zpm stop api`, `zpm restart api`).
- **Execution Directory Preservation**: Processes start accurately within their target working directories (`--cwd`), maintaining relative assets and module paths.

### Live Terminal Dashboard (TUI)
- **Interactive Console UI**: Real-time curses dashboard (`zpm tui`) showcasing CPU, RSS memory footprint, restart counters, and live uptime.
- **Responsive Layout Adaptation**: Automatically adapts between compact mobile-width terminals and wide desktop console viewports.
- **Keyboard Navigation**: Smooth interactive keyboard controls (`j`/`k`, arrow keys) for inspecting individual process logs and metrics.

### High-Performance Log Streaming & Rotation
- **Zero-Drop Standard I/O Capture**: Dedicated non-blocking stdout and stderr aggregation directly into organized process log vaults.
- **Configurable Log Rotation**: Automatic size-based rotation (`--log-max-size`) with configurable archive retention limits (`--log-backups`).
- **Live Output Following**: Real-time console log streaming directly during startup using `--follow` or via process log inspections.

### Zero-Overhead Self-Contained Footprint
- **Static Native Binaries**: Pure Go implementation compiled without CGO dependencies; starts instantly with microscopic memory overhead (<10 MB RAM).
- **Clean Isolated IPC**: Fast, reliable local socket/TCP communication between CLI tools and the supervision daemon.
- **Standardized State Storage**: Predictable, clean filesystem layout storing logs, state registries, and PID locks inside native platform user configuration directories.

## Contributions

If you want to help make this project better, please take the time to read this contribution [doc](./docs/CONTRIBUTING.md) and [fork](https://github.com/thruqe/zpm/fork) this repository. Then open a pull request with your changes.

If you are using any AI agent for assistance, please refer to the [AGENTS documentation](./docs/AGENTS.md) Guide.

## Acknowledgements

zpm wouldn't have been possible without these open source, community and passion-driven projects: [PM2](https://pm2.keymetrics.io/), [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Lip Gloss](https://github.com/charmbracelet/lipgloss), and the broader Go ecosystem.

## Licensing

This project is open source, see the [LICENSE](./LICENSE) file for full details.

## Support this project

If you found this project to be useful in anyway, please consider sending a tip.

<a href="https://etherscan.io/address/0x76a5307e573401f2b7189fb721a33dc483e985a8"><img src="https://cdn-icons-png.flaticon.com/128/15301/15301597.png" width="24" height="24" alt="Donate ETH" /></a> `0x76a5307e573401f2b7189fb721a33dc483e985a8`
