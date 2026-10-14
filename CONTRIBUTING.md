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

## Publish a release

Merge verified source & create its matching immutable tag.
Bind every required check to the approved native publisher.
`release-policy.json` refuses releases until that App is configured.
Require trusted native proof binding artifact digests to qualified source.
Confirm that integration before configuring the App.

Prepare builds on an authorized runner:

Install the Linux x86-64 distribution pinned in `release-policy.json`.
Keep GoReleaser, Cosign, the archive, checksums & signature bundle together.
Add their directory to `PATH` before running checks or preparing builds.
`bin/run-goreleaser` verifies their hashes, signer, issuer & runtime versions.

```sh
./scripts/release prepare
```

Transfer the prepared directory with its verified SHA-256 inventory.
Set `GPG_FINGERPRINT` to the registry's registered public-key fingerprint.
On macOS, set `RELEASE_SIGNING_KEYCHAIN_SERVICE` to the existing Keychain service name.

```sh
./scripts/release publish /path/to/prepared/artifacts
./scripts/release verify /path/to/prepared/artifacts
```

Matching published assets stay unchanged. Conflicting assets stop publication.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
