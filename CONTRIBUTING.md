# Contributing

Thanks for improving the Xquik Terraform provider.

Read the organization [contribution policy](https://github.com/Xquik-dev/.github/blob/main/CONTRIBUTING.md).

Report vulnerabilities through [SECURITY.md](SECURITY.md).

## Set Up

Install the Go version declared in `go.mod`.

Then run:

```sh
./scripts/bootstrap
./scripts/build
```

The bootstrap script downloads pinned modules and verifies their checksums.

## Make Changes

Keep each pull request focused.

Preserve public contracts defined by the provider schemas and API SDK.

Most generated files identify Stainless in their opening comment.

Prefer generator changes when a generated contract needs correction.

Hand-maintained provider logic lives under `internal/services/x_write`.

Regenerate Registry documentation after changing schemas:

```sh
./scripts/generate-docs
```

Never include API keys, tokens, Terraform state, or private account data.

## Verify Changes

Run every required check:

```sh
./scripts/bootstrap
./scripts/lint
./scripts/test
```

The test command runs race detection.

It enforces 90% statement coverage and 80% branch coverage.

It also verifies Windows compilation.

Run the branch gate independently:

```sh
./scripts/branch-coverage
```

Add regression tests for every corrected defect.

Acceptance tests can create remote resources and incur charges.

Run them only with an isolated account:

```sh
TF_ACC=1 ./scripts/test
```

## Submit Changes

Sign each non-trivial commit under the DCO:

```sh
git commit -s
```

Open a pull request and resolve every blocking comment.

A person other than the author must approve non-trivial changes.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
