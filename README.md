# mongodb.go

[![CI](https://github.com/sweetrpg/mongodb.go/actions/workflows/ci.yaml/badge.svg)](https://github.com/sweetrpg/mongodb.go/actions/workflows/ci.yaml)
[![License](https://img.shields.io/github/license/sweetrpg/mongodb.go.svg)](https://img.shields.io/github/license/sweetrpg/mongodb.go.svg)
[![Issues](https://img.shields.io/github/issues/sweetrpg/mongodb.go.svg)](https://img.shields.io/github/issues/sweetrpg/mongodb.go.svg)
[![PRs](https://img.shields.io/github/issues-pr/sweetrpg/mongodb.go.svg)](https://img.shields.io/github/issues-pr/sweetrpg/mongodb.go.svg)
[![Dependabot](https://badgen.net/github/dependabot/sweetrpg/mongodb.go)](https://badgen.net/github/dependabot/sweetrpg/mongodb.go)

Generic MongoDB access layer for sweetrpg's Go services: connection setup/teardown and
`Get`/`Query`/`Insert`/`Update`/`Delete` functions parametrized over any model type.

## Install

```bash
go get github.com/sweetrpg/mongodb.go
```

## Packages

- `database` - `SetupDatabase`/`TeardownDatabase` (connection lifecycle, configured via
  `DB_URI` or the individual `DB_HOST`/`DB_USER`/`DB_PW`/etc. environment variables), and
  generic `Get[T]`/`Query[T]`/`Insert[T]`/`Update[T]`/`Delete[T]` functions
- `constants` - environment variable names and query-paging defaults

## Documentation

Package documentation: [pkg.go.dev/github.com/sweetrpg/mongodb.go](https://pkg.go.dev/github.com/sweetrpg/mongodb.go).
Test coverage reports are published to [sweetrpg.github.io/mongodb.go](https://sweetrpg.github.io/mongodb.go)
on every merge to `develop`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the development workflow (including running the
MongoDB-backed test suite locally) and [RELEASE.md](RELEASE.md) for how versions get cut.
