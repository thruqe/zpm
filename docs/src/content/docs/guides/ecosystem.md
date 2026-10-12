---
title: Ecosystem Config
description: Declarative multi-app configurations with environment profiles.
---

Instead of running individual CLI commands, you can declare all your services in a single `zpm.config.json` configuration file.

## Example Configuration

Create `zpm.config.json` in the root of your project:

```json
{
  "apps": [
    {
      "name": "api",
      "command": "bun run server.ts",
      "cwd": "./backend",
      "instances": 2,
      "auto_restart": true,
      "restart_delay": 1000,
      "max_restarts": -1,
      "health_command": "curl -fsS http://127.0.0.1:4000/health",
      "log_max_bytes": 10485760,
      "log_backups": 5,
      "env": {
        "PORT": "4000"
      },
      "env_production": {
        "NODE_ENV": "production",
        "PORT": "8080"
      }
    },
    {
      "name": "worker",
      "command": "python queue_worker.py",
      "cwd": "./workers",
      "instances": 1,
      "auto_restart": true
    }
  ]
}
```

## Running Ecosystem Files

Start all defined applications:

```bash
zpm ecosystem start zpm.config.json
```

Or activate a specific environment profile:

```bash
zpm ecosystem start zpm.config.json --env production
```

This injects variables from `env_production` while overriding defaults in `env`.
