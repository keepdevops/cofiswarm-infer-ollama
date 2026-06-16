# cofiswarm-infer-ollama

Cofiswarm component: `infer-ollama`.

- Layout: [REPO-STANDARD-LAYOUT](https://github.com/keepdevops/cofiswarmdev/blob/main/docs/REPO-STANDARD-LAYOUT.md)
- Migration: [MIGRATION-SPRINTS](https://github.com/keepdevops/cofiswarmdev/blob/main/docs/MIGRATION-SPRINTS.md)

## FHS paths

| Path | Purpose |
|------|---------|
| `/etc/cofiswarm/infer-ollama/` | config |
| `/var/lib/cofiswarm/infer-ollama/` | state |
| `/var/log/cofiswarm/infer-ollama/` | logs |

## Test

```bash
./test/scripts/assert-layout.sh infer-ollama
```
