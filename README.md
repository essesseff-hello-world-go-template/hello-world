# hello-world Application (Go)

Golang (Go) implementation with trunk-based development and event-driven deployments via essesseff platform.

## Architecture

- **Branch Strategy**: Single `main` branch (trunk-based)
- **Auto-Deploy**: DEV only
- **Manual Deploy**: QA, STAGING, PROD (via essesseff)

## Development Workflow

```bash
# 1. Create feature branch
git checkout -b feature/my-feature

# 2. Make changes and commit
git commit -am "Add feature"

# 3. Push and create PR
git push origin feature/my-feature

# 4. After review, merge to main
# This triggers automatic build and deploy to DEV

# 5. Use essesseff UI for promotions:
#    - Developer declares Release Candidate (RC)
#    - QA accepts RC → deploys to QA (or alternatively rejects the promotion of the RC to QA)
#    - QA marks as Stable (or alternatively rejects the promotion to Stable)
#    - Release Engineer deploys from Stable Release to STAGING/PROD
```

## Local Development

```bash
# Run locally
go run main.go

# Build container
docker build -t hello-world:local .
docker run -p 8080:8080 hello-world:local
```

## Endpoints

- `/` - Main page
- `/health` - Health check
- `/ready` - Readiness check

## Related Repositories

- Source: hello-world (this repo)
- Config DEV: hello-world-config-dev
- Config QA: hello-world-config-qa
- Config STAGING: hello-world-config-staging
- Config PROD: hello-world-config-prod
- Argo CD Apps: hello-world-app-of-apps
