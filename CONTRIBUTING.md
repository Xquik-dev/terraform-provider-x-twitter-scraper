# Contribute

Read the organization [contribution policy](https://github.com/Xquik-dev/.github/blob/main/CONTRIBUTING.md).

Report vulnerabilities through [SECURITY.md](SECURITY.md).

## Set up the repository

Install the Go version declared in `go.mod`.

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
./scripts/bootstrap
./scripts/lint
./scripts/test
```

The test command runs race detection, enforces 90% statement and 80% branch
coverage, and verifies Windows compilation.

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

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
