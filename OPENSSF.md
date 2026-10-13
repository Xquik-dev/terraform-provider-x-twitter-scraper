# OpenSSF Best Practices evidence

The official assessment is [bestpractices.dev project 13729](https://www.bestpractices.dev/projects/13729).

The assessment date is 2026-08-20.

## Eligibility

This public, released Terraform provider meets the OpenSSF eligibility rules.

## Verified technical controls

| Area | Evidence |
| --- | --- |
| License | Apache-2.0, REUSE 3.3 metadata, and bundled BSD-3-Clause terms |
| Contribution process | `CONTRIBUTING.md`, DCO sign-off, and `bin/check-local` before every merge |
| Security reporting | Private GitHub reporting, response targets, and safe harbor in `SECURITY.md` |
| Build | Pinned Go modules, checksum verification, and `./scripts/build` |
| Tests | `./scripts/test` runs race-enabled unit and regression tests |
| Statement coverage | `./scripts/coverage` enforces 90%. The latest full result is 93.6% |
| Branch coverage | `./scripts/branch-coverage` enforces 80%. The current result is 794/963 (82.45%) |
| Static analysis | `go vet`, ShellCheck and CodeQL security-extended queries through `bin/check-local` |
| Dynamic analysis | Race detection, plus 15 seconds of Go fuzzing in every `bin/check-local` run |
| Dependency review | `go mod verify`, and `govulncheck` at every severity with reviewed exceptions in `audit-exceptions.json` |
| Required checks | `bin/check-local` runs every check and posts the `xquik-local-checks` commit status |
| Releases | `bin/release-local` publishes GPG-signed checksums and exact annotated tags, with a `--dry-run` that stops before signing |
| Reproducibility | `bin/check-local` builds each GoReleaser snapshot twice and compares checksums. A dry run builds with the release tag, so its checksums equal the release's |
| Documentation | Generated Registry pages match registered schemas |
| Portability | `./scripts/test` compiles every package for Windows AMD64 |
| Two-factor authentication | The Xquik-dev organization requires 2FA |

Verify locally.

```sh
bin/check-local
```

## Branch coverage evidence

The branch gate pins gocove, rebuilds every package, and instruments control
flow. Native Go coverage confirms executed ranges. The local verifier corrects
nested-block associations and excludes only unshipped test support and itself.

The current result is 794 of 963 branches, or 82.45%.

## Outstanding Gold blockers

These criteria require human or organizational evidence.

| Gold Criterion | Current Evidence | Required Action |
| --- | --- | --- |
| Access continuity | No evidence proves two people hold required release access | Grant and verify access for another maintainer |
| Bus factor | Git history shows one significant contributor | Add another significant contributor |
| Unassociated contributors | Fewer than two contributors are organization-independent | Accept qualifying external contributions |
| Independent modification review | Historical changes do not prove the required 50% threshold | Require another person to review qualifying pull requests |
| Human security review | No completed independent review exists within five years | Commission and publish a scoped human review |

The repository must not claim Gold while any mandatory criterion remains unmet.

## Maintenance

`bin/check-local` checks licensing, tests, coverage, security, and
reproducibility before each merge and release. Reassess this file before each
major release. Update bestpractices.dev only with public evidence.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
