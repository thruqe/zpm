---
title: Configuration Reference
description: Schema specification for zpm.config.json ecosystem files.
---

Schema specification for the declarative ecosystem configuration file (`zpm.config.json`).

## Top-Level Schema

The configuration file requires a root `apps` array:

```json
{
  "apps": [
    {
      "name": "api",
      "command": "bun run index.ts"
    }
  ]
}
```

## Application Object Properties

| Property | Type | Default | Description |
| :------- | :--- | :------ | :---------- |
| `name` | string | *required* | Unique identifier for the process. |
| `command` | string | *required* | The executable command to run. |
| `cwd` | string | `"."` | Directory to execute the command from. |
| `instances` | integer | `1` | Number of clustered worker instances to spawn. |
| `auto_restart` | boolean | `true` | Whether to automatically revive crashed processes. |
| `restart_delay` | integer | `1000` | Delay in milliseconds before attempting restart. |
| `max_restarts` | integer | `-1` | Maximum crash restart attempts (-1 for infinite). |
| `health_command` | string | `""` | Shell command probe executed to test application health. |
| `log_max_bytes` | integer | `10485760` | Size threshold in bytes before rotating log files. |
| `log_backups` | integer | `5` | Number of historical rotated log files to retain. |
| `env` | object | `{}` | Key-value mapping of default environment variables. |
| `env_<name>` | object | `{}` | Environment profile variables loaded when `--env <name>` is passed. |
