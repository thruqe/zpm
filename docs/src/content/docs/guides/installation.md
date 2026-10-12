---
title: Installation
description: How to install ZPM on Linux, macOS, and Windows.
---

import { Tabs, TabItem } from '@astrojs/starlight/components';

ZPM consists of two binaries:
- **`zpm`**: The client command-line interface and terminal dashboard.
- **`zpmd`**: The background supervision daemon that keeps processes alive.

## One-Liner Installers

<Tabs>
  <TabItem label="Linux / macOS (Bash)">
```bash
curl -sL https://raw.githubusercontent.com/thruqe/zpm/main/scripts/install.sh | bash
```
  </TabItem>
  <TabItem label="Windows (PowerShell as Admin)">
```powershell
irm https://raw.githubusercontent.com/thruqe/zpm/main/scripts/install-windows.ps1 | iex
```
  </TabItem>
</Tabs>

---

## Package Managers

### Homebrew (macOS & Linux)

Install via the official Thruqe Homebrew tap:

```bash
brew tap thruqe/tap
brew install zpm
```

Or install directly in a single command:

```bash
brew install thruqe/tap/zpm
```

### WinGet (Windows)

Install using the Windows Package Manager:

```powershell
winget install thruqe.zpm
```

### Chocolatey (Windows)

Install via Chocolatey:

```powershell
choco install zpm
```

---

## Build from Source

Requirements: Go 1.22 or newer.

```bash
git clone https://github.com/thruqe/zpm.git
cd zpm

# Build daemon and CLI
go build -o daemon/zpmd ./daemon
go build -o src/zpm ./src

# Link or copy to your PATH
ln -sf "$PWD/src/zpm" ~/.local/bin/zpm
ln -sf "$PWD/src/zpm" ~/.local/bin/zpp
```

Keep `zpmd` beside `zpm`, or set the `ZPMD_PATH` environment variable:

```bash
export ZPMD_PATH=/path/to/daemon/zpmd
```

---

## Platform Auto-Start Configuration

ZPM integrates natively with OS init systems:

| Platform | Init Manager | Setup Command |
| :------- | :----------- | :------------ |
| **Linux** | systemd user service | `zpm startup install` |
| **macOS** | LaunchAgent | `zpm startup install` |
| **Windows** | Task Scheduler | `zpm startup install` (Run as Admin) |

To remove startup entries:
```bash
zpm startup uninstall
```
