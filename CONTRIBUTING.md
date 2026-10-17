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

1. `go run ./scripts/api_coverage.go`, the [API coverage](#api-coverage)
   check
2. `./scripts/bootstrap` to download and verify modules
3. `reuse lint` 6.2.0 through `uvx` for license metadata
4. `./scripts/lint` for formatting, module metadata, `go vet` and builds
5. ShellCheck on every bash script
6. `./scripts/test-scripts`, the tests of the scripts in `bin/` and `scripts/`
7. `./scripts/test` for race detection, 90% statement and 80% branch
   coverage, and Windows compilation
8. `./scripts/audit`, the dependency audit
9. `goreleaser check` and the Terraform Registry manifest, which must name
   protocol 6.0
10. 2 GoReleaser snapshot builds whose checksums must match, and a DWARF debug
    information check
11. Go fuzzing of the multipart encoder for 15 seconds, or for `FUZZ_TIME`,
    such as `FUZZ_TIME=2m`
12. CodeQL with the Go `security-extended` queries on a copy of the tree,
    which fails on any result and on any file CodeQL could not read
13. A check that no step changed the working tree

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

## API coverage

`bin/check-local` starts with `go run ./scripts/api_coverage.go`. It compares
the operations in the public API spec with the requests that resources and data
sources send. It fails when the API has an operation that the provider neither
sends nor lists in `excluded-operations.json`. An operation that the provider
sends and the spec lacks is fine.

The check reads the spec from https://xquik.com/openapi.json. Set
`XQUIK_OPENAPI_SPEC` to a local file path or another URL to read another copy.
The spec must be JSON.

It loads the Go packages under `internal/services` with their types, as the
compiler sees them. Tests and `testdata` folders stay out.

- Every use of a Go SDK method on `Client` or on a `Service` type must call it.
  A call through a variable, an embedded client or a function literal counts.
- A service method, such as `r.client.X.Tweets.Get`, counts as the request it
  sends. The check reads that method in the Go SDK source of the version in
  `go.mod`, by the rules of the Go SDK's own check in `cmd/api-coverage`.
- A raw call, such as `r.client.Get` or `client.Execute`, needs a readable
  path. That is a constant, a `fmt.Sprintf` call with a constant format and a
  value for each verb, a variable assigned once to 1 of those, or a package
  function whose body returns 1 of those in 1 line. `Execute` also needs a
  constant HTTP method. The check drops a query string after `?`.
- `client.Execute(ctx, op.method, op.path(id), ...)` reads an operation table.
  The method comes from a field of `op`. The path comes from a field of `op`,
  or from a method of `op` whose every return gives 1 field, alone or as the
  `fmt.Sprintf` format. Every literal of that type in the package needs a
  constant HTTP method and a constant path format in those fields. The
  `x_write` operations work this way.

The check puts the path of the Go SDK's default base URL, from
`option.WithEnvironmentProduction`, before each path. It fails on a call it
cannot read and names its file and line. Path parameter names do not matter,
so `{id}` matches `{tweetId}`.

Fix a failure by adding a resource or data source that sends the missing
operation. An operation that the provider should not send goes in
`excluded-operations.json` with a reason. Remove an entry when the API drops
its operation or the provider sends it. The check fails on a stale entry.

`excluded-operations.json` is a list of groups. Each group has a non-empty
`reason` and at least 1 operation. Write an operation as an upper case method,
a space and the absolute path from the spec. List each operation once.

```json
[
  {
    "reason": "Searches the Xquik docs at https://xquik.com/ask, outside the SDK base URL https://xquik.com/api/v1.",
    "operations": ["GET /ask", "POST /ask"]
  }
]
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
