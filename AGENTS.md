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

See `RELEASE.md`. Summary: trigger `prepare-release.yaml` (`workflow_dispatch` against
`develop`), which computes the next version from conventional commits via git-cliff and opens
a `release/<version>` PR into `master`. Merging that PR tags the release
(`tag-release.yaml`), which triggers `release.yaml` - re-runs tests, creates a GitHub
Release, and merges `master` back into `develop`.

## Platform Conventions and Decisions

This repo is a submodule of `sweetrpg/platform`. Platform-wide conventions and Architecture
Decision Records live there:

- checked out inside the platform tree: `../../docs/README.md` (convention index) and
  `../../docs/adr/README.md` (platform ADRs, `PADR-*`)
- standalone or module-only checkout:
  <https://github.com/sweetrpg/platform/tree/master/docs> and
  <https://github.com/sweetrpg/platform/tree/master/docs/adr>

Accepted `PADR-*` records are binding constraints. Read the ADR index before proposing a
structural change; if a task needs to contradict an accepted record, stop and say so -
propose a superseding ADR rather than working around it. This repo's own service-local
decisions are `ADR-*` and belong in `docs/adr/` here.

Before implementing any feature, complete the
[feature work checkpoint](https://github.com/sweetrpg/platform/blob/master/docs/openspec.md#feature-work-checkpoint)
and record the answers in the issue or PR.
