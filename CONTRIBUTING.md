# Contribute

Read the organization [contribution policy](https://github.com/Xquik-dev/.github/blob/main/CONTRIBUTING.md).

Report vulnerabilities through [SECURITY.md](SECURITY.md).

## Set up the repository

Install the Go version declared in `go.mod`, GoReleaser, the CodeQL CLI, jq,
ShellCheck and uv. On macOS, `./scripts/bootstrap` offers to install them from
the `Brewfile`.

Then run these commands.

```sh
./scripts/bootstrap
./scripts/build
```

The bootstrap script downloads pinned modules and verifies checksums.

## Make changes

Keep each pull request focused.

Preserve public contracts defined by the provider schemas and API SDK.

Generated files say so in their opening comment.

Change the generator when a generated contract needs correction.

Hand-maintained provider logic lives under `internal/services/x_write`.

Regenerate Terraform Registry documentation after changing a schema.

```sh
./scripts/generate-docs
```

Never include API keys, tokens, Terraform state, or private account data.

## Verify changes

Run every required check.

```sh
bin/check-local
```

It runs these checks in order.

1. `./scripts/bootstrap` to download and verify modules
2. `reuse lint` 6.2.0 through `uvx` for license metadata
3. `./scripts/lint` for formatting, module metadata, `go vet` and builds
4. ShellCheck on every bash script
5. `./scripts/test-scripts`, the tests of the scripts in `bin/` and `scripts/`
6. `./scripts/test` for race detection, 90% statement and 80% branch
   coverage, and Windows compilation
7. `./scripts/audit`, the dependency audit
8. `goreleaser check` and the Terraform Registry manifest, which must name
   protocol 6.0
9. 2 GoReleaser snapshot builds whose checksums must match, and a DWARF debug
   information check
10. Go fuzzing of the multipart encoder for 15 seconds, or for `FUZZ_TIME`,
    such as `FUZZ_TIME=2m`
11. CodeQL with the Go `security-extended` queries on a copy of the tree,
    which fails on any result
12. A check that no step changed the working tree

It stops at the first failure and says what to fix.

The audit runs govulncheck over every package and test and counts every
finding at every severity. It blocks on any finding in a module that ships in
the provider binary, the standard library included, and on any other finding
with a fixed version. An unfixable development-only finding passes only while
`audit-exceptions.json` lists it with its advisory, module, reason, and a
recheck date at most 30 days out. The list is empty today.

On a clean tree whose commit is on GitHub, `bin/check-local` posts the
`xquik-local-checks` commit status with your `gh` login. It posts success only
after every check passes, and the text names only the checks that ran.
Maintainers merge a pull request after that status passes on its exact head.
Pass `--no-status` to skip it.

Run the branch gate on its own.

```sh
./scripts/branch-coverage
```

Add regression tests for every corrected defect.

Acceptance tests may create remote resources and incur charges. Run them only
with an isolated account.

```sh
TF_ACC=1 ./scripts/test
```

## Release

A release is the annotated tag `v<version>` plus a GitHub release that holds
what the Terraform Registry reads. That is 1 zip per platform, the registry
manifest, the SHA256SUMS file and its binary GPG signature.

Maintainers release from their own machine with `bin/release-local`. It runs
from a clean `main` that matches `origin/main`, and it takes these steps.

1. Reads the version from `.release-please-manifest.json`
2. Stops when tag `v<version>` is on GitHub, the GitHub release is published,
   or the Terraform Registry already lists the version
3. Runs `bin/check-local` and the dependency audit
4. Builds the zips and SHA256SUMS with GoReleaser and checks them against
   what the registry reads
5. Signs SHA256SUMS with the provider's GPG key
6. Uploads the assets to a draft release and checks them there
7. Pushes the tag and publishes the release with notes from `CHANGELOG.md`

The GPG passphrase lives in the macOS Keychain. Set
`RELEASE_SIGNING_KEYCHAIN_SERVICE` to the name of that Keychain item. The
script stops at once when the variable is unset. The passphrase reaches gpg
through a pipe, never through a command line argument.

Run `bin/release-local --dry-run` to build every asset and print the version,
the assets with sizes and checksums, the registry manifest and the release
notes. A dry run needs no Keychain variable, posts no commit status and stops
before signing, the tag and the release.

The Terraform Registry reads the signed release from GitHub.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
