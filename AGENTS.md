# AGENTS.md

This file provides guidance to Claude Code, Codex, GitHub Copilot, and other AI coding agents
working in this repository.

## About This Project

`mongodb.go` (formerly `db.go` - the GitHub repo was renamed; the module path is
`github.com/sweetrpg/mongodb.go`) provides generic MongoDB access helpers (Get, Query, Insert,
Update, Delete) plus connection setup/teardown for sweetrpg's Go services. It depends only on
`common.go` for logging.

## Consumers

Depended on by `api-core.go`, `catalog-data.go`, `catalog-api`, and other sweetrpg data-access
layers. If a repo's `go.mod` still requires `github.com/sweetrpg/db.go`, that's a stale
reference to this repo's old name and should be updated to `github.com/sweetrpg/mongodb.go`.

## Committing Code

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>
```

## Branches and Workflow

* `develop` - integration branch, default branch, target for all PRs.
* `master` - latest released state, nothing committed directly.
* `feature/*`, `fix/*` branched from `develop`; `hotfix/*` branched from `master`.

See `CONTRIBUTING.md` for the full workflow, including running the database-backed test suite
locally.

## Running Checks Locally

```bash
docker run --rm -d -p 27017:27017 --name mongodb-test mongo:7.0
export TEST_DB_URI="mongodb://localhost:27017/unit-tests"
export TEST_COLLECTION=unit-tests
go build -v ./...
go vet ./...
go test -v -coverprofile coverage.out ./...
docker stop mongodb-test
```

## Releases

Merges to `develop` auto-tag a patch release via CI (`.github/workflows/go-ci.yml`). Use the
"Bump version" workflow (`.github/workflows/bump-version.yml`, manually dispatched) for a minor
or major bump instead.
