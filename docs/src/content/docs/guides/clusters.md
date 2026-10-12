---
title: Clusters
description: Scale processes across multiple worker instances.
---

ZPM supports multi-instance process clustering to utilize multiple CPU cores and achieve high availability.

## Spawning Instances

Use the `--instances` flag to spin up multiple workers:

```bash
zpm start "node server.js" --name web --instances 4
```

ZPM automatically names individual instances with numeric suffixes:
- `web-0`
- `web-1`
- `web-2`
- `web-3`

## Managing Cluster Groups

When issuing lifecycle commands, using the base cluster name affects all worker instances:

```bash
# Restart all 4 workers
zpm restart web

# Stop all 4 workers
zpm stop web

# Purge the entire cluster
zpm purge web
```

You can also target individual instances if necessary:

```bash
zpm restart web-0
zpm stop web-2
```

## Zero-Downtime Rolling Restarts

To reboot a cluster with zero downtime, restart instances sequentially:

```bash
for i in {0..3}; do
  zpm restart web-$i
  sleep 2
done
```
