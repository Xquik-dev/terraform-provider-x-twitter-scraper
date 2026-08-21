# OpenSSF Best Practices evidence

Official assessment: [bestpractices.dev project 13729](https://www.bestpractices.dev/projects/13729).

Assessment date: 2026-08-20.

## Eligibility

This public, released Terraform provider meets the OpenSSF eligibility rules.

## Verified technical controls

| Area | Evidence |
| --- | --- |
| License | Apache-2.0, REUSE 3.3 metadata, and bundled BSD-3-Clause terms |
| Contribution process | `CONTRIBUTING.md`, DCO sign-off, and independent review rules |
| Security reporting | Private GitHub reporting, response targets, and safe harbor in `SECURITY.md` |
| Build | Pinned Go modules, checksum verification, and `./scripts/build` |
| Tests | `./scripts/test` runs race-enabled unit and regression tests |
| Statement coverage | `./scripts/coverage` enforces 90%; latest full result is 93.6% |
| Branch coverage | `./scripts/branch-coverage` enforces 80%; current result is 794/963 (82.45%) |
| Static analysis | `go vet` and CodeQL security-extended queries |
| Dynamic analysis | Scheduled Go fuzzing and race detection |
| Dependency review | Dependabot, `go mod verify`, and `govulncheck` |
| CI | Pull requests from this repository and forks run every required job |
| Releases | Signed checksums, exact tags, and GitHub artifact attestations |
| Reproducibility | CI builds each GoReleaser snapshot twice and compares checksums |
| Documentation | Generated Registry pages match registered schemas |
| Portability | CI compiles every package for Windows AMD64 |
| Two-factor authentication | The Xquik-dev organization requires 2FA |

Verify locally:

```sh
./scripts/bootstrap
./scripts/lint
./scripts/test
uvx --from reuse==5.1.1 reuse lint
go run golang.org/x/vuln/cmd/govulncheck@v1.6.0 ./...
```

## Branch coverage evidence

The branch gate pins gocove, rebuilds every package, and instruments control
flow. Native Go coverage confirms executed ranges. The local verifier corrects
nested-block associations and excludes only unshipped test support and itself.

The current result is 794 of 963 branches, or 82.45%.

## Outstanding Gold blockers

These criteria require human or organizational evidence:

| Gold Criterion | Current Evidence | Required Action |
| --- | --- | --- |
| Access continuity | No evidence proves two people hold required release access | Grant and verify access for another maintainer |
| Bus factor | Git history shows one significant contributor | Add another significant contributor |
| Unassociated contributors | Fewer than two contributors are organization-independent | Accept qualifying external contributions |
| Independent modification review | Historical changes do not prove the required 50% threshold | Require another person to review qualifying pull requests |
| Human security review | No completed independent review exists within five years | Commission and publish a scoped human review |

The repository must not claim Gold while any mandatory criterion remains unmet.

## Maintenance

CI checks licensing, tests, coverage, security, and reproducibility. Scheduled
CodeQL, fuzzing, Scorecard, and dependency updates detect drift. Reassess this
file before each major release. Update bestpractices.dev only with public evidence.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
