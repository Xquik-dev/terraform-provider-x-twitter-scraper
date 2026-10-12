# Contribute

Read the organization [contribution policy](https://github.com/Xquik-dev/.github/blob/main/CONTRIBUTING.md).

Report vulnerabilities through [SECURITY.md](SECURITY.md).

## Set up the repository

Install the Go version declared in `go.mod`, GoReleaser, and `jq`.
On macOS, `./scripts/bootstrap` offers to install Go and GoReleaser from the
`Brewfile`.

Run:

```sh
./scripts/bootstrap
./scripts/build
```

The bootstrap script downloads pinned modules and verifies checksums.

## Make changes

Keep each pull request focused.

Preserve public contracts defined by the provider schemas and API SDK.

Generated files identify their generator in the opening comment.

Change the generator when a generated contract needs correction.

Hand-maintained provider logic lives under `internal/services/x_write`.

Regenerate Terraform Registry documentation after changing a schema:

```sh
./scripts/generate-docs
```

Never include API keys, tokens, Terraform state, or private account data.

## Verify changes

Run every required check:

```sh
bin/check-local
```

It runs, in order:

1. `./scripts/bootstrap` to download and verify modules
2. `./scripts/lint` for formatting, module metadata, `go vet`, and builds
3. `reuse lint` for license metadata, when `reuse` or `uv` is installed
4. `./scripts/test` for race detection, 90% statement and 80% branch
   coverage, and Windows compilation
5. A GoReleaser snapshot build and a DWARF debug information check
6. `./scripts/vulncheck` for reachable known vulnerabilities
7. A second snapshot build that must match the first checksums
8. `goreleaser check` and the Terraform Registry manifest

It stops at the first failure and says what to fix. Set `FUZZ_TIME=2m` to
also fuzz the multipart encoder.

On a clean tree whose commit is on GitHub, `bin/check-local` posts the
`xquik-local-checks` commit status. Maintainers merge a pull request after
that status passes. Pass `--no-status` to skip it.

Run the branch gate independently:

```sh
./scripts/branch-coverage
```

Add regression tests for every corrected defect.

Acceptance tests may create remote resources and incur charges. Run them only
with an isolated account:

```sh
TF_ACC=1 ./scripts/test
```

## Release

Maintainers release from their own machine with `bin/release-local`. It runs
from a clean `main` that matches `origin/main`, and it:

1. Reads the version from `.release-please-manifest.json`
2. Stops when tag `v<version>` already exists on GitHub
3. Runs `bin/check-local` and `./scripts/vulncheck`
4. Builds the release with GoReleaser
5. Stops when the Terraform Registry already lists the version
6. Signs the checksum file with the provider's GPG key
7. Pushes the annotated tag and publishes the GitHub release with notes from
   `CHANGELOG.md`

The Terraform Registry reads the signed release from GitHub.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
