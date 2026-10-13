#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

# Runs bin/release-local in a throwaway repository whose origin is a local
# bare repository. Stubs stand in for bin/check-local and the audit. A fake
# GoReleaser builds 2 small zips and a SHA256SUMS file, fake gpg and security
# commands stand in for the GPG keyring and the Keychain, a fake curl answers
# for the Terraform Registry, and a fake GitHub CLI keeps the draft release in
# the logs folder. No test reaches GitHub, the registry, the keyring or the
# Keychain.

set -euo pipefail

# shellcheck source=scripts/tests/support.sh
source "$(dirname -- "${BASH_SOURCE[0]}")/support.sh"

readonly repository="Xquik-dev/terraform-provider-x-twitter-scraper"
readonly project="terraform-provider-x-twitter-scraper"
readonly test_service="test-signing-item"
readonly test_passphrase="test-passphrase-for-release-local"
readonly test_fingerprint="0000111122223333444455556666777788889999"

# Builds a releasable repository at the given version, 1.2.3 by default.
prepare_release() {
  local version="${1:-1.2.3}"
  copy_file bin/release-local
  copy_file scripts/release-listing
  stub bin/check-local <<'EOF'
#!/usr/bin/env bash
# Stands in for bin/check-local. It records its arguments.
echo "check-local${*:+ $*}" >>"$FAKE_LOGS/order"
EOF
  stub scripts/audit <<'EOF'
#!/usr/bin/env bash
# Stands in for scripts/audit.
echo "audit" >>"$FAKE_LOGS/order"
EOF
  stub bin/check-release-debug-info <<'EOF'
#!/usr/bin/env bash
# Stands in for bin/check-release-debug-info.
echo "check-release-debug-info" >>"$FAKE_LOGS/order"
EOF
  write_file .release-please-manifest.json <<EOF
{".": "$version"}
EOF
  write_file internal/version.go <<EOF
package internal

const PackageVersion = "$version" // x-release-please-version
EOF
  write_file terraform-registry-manifest.json <<'EOF'
{
  "version": 1,
  "metadata": {
    "protocol_versions": ["6.0"]
  }
}
EOF
  write_file CHANGELOG.md <<EOF
# Changelog

## $version (2026-10-13)

### Bug Fixes

* sign the checksums from a local runner

## 0.0.1 (2026-01-01)

* first release
EOF
  write_file .gitignore <<'EOF'
dist/
EOF
  fake go <<'EOF'
#!/bin/sh
# Stands in for go. bin/release-local only checks that it exists.
exit 0
EOF
  fake goreleaser <<'EOF'
#!/bin/sh
# Stands in for GoReleaser. Each run builds a Linux and a Windows zip and a
# SHA256SUMS file that lists them and the registry manifest, as the real
# checksum section does. Like Go, it stamps each binary with the tag on HEAD,
# or with a pseudo-version when HEAD has no tag.
# FAKE_SUMS_WITHOUT_MANIFEST leaves the manifest out,
# FAKE_WRONG_BINARY names the Linux binary wrongly and FAKE_DIRTY_BUILD
# changes a tracked file.
{
  printf 'goreleaser %s' "$*"
  [ -n "${GORELEASER_CURRENT_TAG:-}" ] && printf ' with GORELEASER_CURRENT_TAG=%s' "$GORELEASER_CURRENT_TAG"
  printf '\n'
} >>"$FAKE_LOGS/order"
version="$(sed -n 's/^const PackageVersion = "\(.*\)".*$/\1/p' internal/version.go)"
prefix="terraform-provider-x-twitter-scraper_${version}"
binary="terraform-provider-x-twitter-scraper_v${version}"
rm -rf dist
mkdir -p dist/linux dist/windows
linux_binary="$binary"
[ -z "${FAKE_WRONG_BINARY:-}" ] || linux_binary="terraform-provider-x-twitter-scraper"
stamp="$(git describe --tags --exact-match HEAD 2>/dev/null || echo "pseudo-version of $(git rev-parse --short HEAD)")"
printf 'linux %s built as %s\n' "$version" "$stamp" >"dist/linux/$linux_binary"
printf 'windows %s built as %s\n' "$version" "$stamp" >"dist/windows/$binary.exe"
touch -t 202601010000 "dist/linux/$linux_binary" "dist/windows/$binary.exe"
(cd dist/linux && zip -qX "../${prefix}_linux_amd64.zip" "$linux_binary")
(cd dist/windows && zip -qX "../${prefix}_windows_amd64.zip" "$binary.exe")
cp terraform-registry-manifest.json "dist/manifest-copy.json"
{
  (cd dist && shasum -a 256 "${prefix}_linux_amd64.zip")
  if [ -z "${FAKE_SUMS_WITHOUT_MANIFEST:-}" ]; then
    printf '%s  %s\n' "$(shasum -a 256 dist/manifest-copy.json | cut -d ' ' -f 1)" "${prefix}_manifest.json"
  fi
  (cd dist && shasum -a 256 "${prefix}_windows_amd64.zip")
} >"dist/${prefix}_SHA256SUMS"
rm dist/manifest-copy.json
cat >dist/artifacts.json <<JSON
[
  {"name": "$binary", "path": "dist/linux/$linux_binary", "type": "Binary", "goos": "linux", "goarch": "amd64"},
  {"name": "${prefix}_linux_amd64.zip", "path": "dist/${prefix}_linux_amd64.zip", "type": "Archive", "goos": "linux", "goarch": "amd64"},
  {"name": "${prefix}_windows_amd64.zip", "path": "dist/${prefix}_windows_amd64.zip", "type": "Archive", "goos": "windows", "goarch": "amd64"},
  {"name": "${prefix}_SHA256SUMS", "path": "dist/${prefix}_SHA256SUMS", "type": "Checksum"}
]
JSON
[ -z "${FAKE_DIRTY_BUILD:-}" ] || echo "changed" >>CHANGELOG.md
EOF
  fake curl <<'EOF'
#!/bin/sh
# Stands in for curl against the Terraform Registry. FAKE_REGISTRY_VERSIONS
# sets the versions it serves, and FAKE_REGISTRY_FAILS makes it fail.
echo "curl $*" >>"$FAKE_LOGS/order"
[ -z "${FAKE_REGISTRY_FAILS:-}" ] || exit 22
versions="${FAKE_REGISTRY_VERSIONS:-}"
[ -n "$versions" ] || versions='{"version": "0.0.1"}'
echo "{\"versions\": [$versions]}"
EOF
  fake security <<'EOF'
#!/bin/sh
# Stands in for the macOS security command. It records its arguments. With
# -w it prints the test passphrase, as the real command prints a password.
# FAKE_NO_KEYCHAIN_ITEM makes the item missing.
echo "security $*" >>"$FAKE_LOGS/order"
echo "$*" >>"$FAKE_LOGS/security-arguments"
[ -z "${FAKE_NO_KEYCHAIN_ITEM:-}" ] || exit 44
case " $* " in
  *" -w "*) echo "test-passphrase-for-release-local" ;;
  *) echo "keychain: \"test.keychain-db\"" ;;
esac
EOF
  fake gpg <<'EOF'
#!/bin/sh
# Stands in for gpg. It records its arguments and what it reads on standard
# input. It lists 1 secret key unless FAKE_NO_KEY is set, writes a binary
# signature unless FAKE_ARMORED is set, and verifies it unless
# FAKE_BAD_SIGNATURE is set. The binary signature starts with the header
# bytes of a real one, which hold a null byte.
echo "gpg $*" >>"$FAKE_LOGS/order"
echo "$*" >>"$FAKE_LOGS/gpg-arguments"
fingerprint="0000111122223333444455556666777788889999"
case " $* " in
  *" --list-secret-keys "*)
    [ -z "${FAKE_NO_KEY:-}" ] || exit 2
    echo "sec:u:4096:1:4900C25E85324F28:1700000000:::u:::scESC:::+:::23::0:"
    echo "fpr:::::::::$fingerprint:"
    ;;
  *" --detach-sign "*)
    cat >>"$FAKE_LOGS/gpg-input"
    output=""
    previous=""
    for argument in "$@"; do
      [ "$previous" = "--output" ] && output="$argument"
      previous="$argument"
    done
    if [ -n "${FAKE_ARMORED:-}" ]; then
      echo "-----BEGIN PGP SIGNATURE-----" >"$output"
    else
      printf '\211\002\117\004\000binary signature\n' >"$output"
    fi
    ;;
  *" --verify "*)
    [ -n "${FAKE_BAD_SIGNATURE:-}" ] && exit 1
    echo "[GNUPG:] VALIDSIG $fingerprint 2026-10-13 1760000000 0 4 0 1 10 00 $fingerprint"
    ;;
  *)
    echo "unexpected gpg $*" >&2
    exit 97
    ;;
esac
EOF
  fake gh <<'EOF'
#!/bin/sh
# Stands in for the GitHub CLI. FAKE_RELEASES holds the releases GitHub
# lists, and FAKE_PUBLISHED makes the tag endpoint return a published
# release. GitHub's list can leave out a published release, so the tests
# keep it out of FAKE_RELEASES. FAKE_TAG_LOOKUP_FAILS makes the tag endpoint
# fail. The draft's assets live in the logs folder, unless FAKE_VIEW_ASSETS
# or FAKE_DOWNLOADED_SUMS override what GitHub returns.
echo "gh $*" >>"$FAKE_LOGS/order"
case "$1 $2" in
  "api repos/Xquik-dev/terraform-provider-x-twitter-scraper/releases/tags/"*)
    if [ -n "${FAKE_TAG_LOOKUP_FAILS:-}" ]; then
      echo "gh: Bad Gateway (HTTP 502)" >&2
      exit 1
    elif [ -n "${FAKE_PUBLISHED:-}" ]; then
      echo '{"tag_name": "v1.2.3", "draft": false}'
    else
      echo "gh: Not Found (HTTP 404)" >&2
      exit 1
    fi
    ;;
  "auth status") exit 0 ;;
  "api --paginate") echo "${FAKE_RELEASES:-[[]]}" ;;
  "release create") : >"$FAKE_LOGS/draft" ;;
  "release upload")
    mkdir -p "$FAKE_LOGS/draft-assets"
    shift 3
    for argument in "$@"; do
      case "$argument" in
        dist/*) cp "$argument" "$FAKE_LOGS/draft-assets/" ;;
      esac
    done
    ;;
  "release view")
    if [ -n "${FAKE_VIEW_ASSETS:-}" ]; then
      echo "$FAKE_VIEW_ASSETS"
    else
      ls "$FAKE_LOGS/draft-assets" | jq -R . | jq -cs 'sort'
    fi
    ;;
  "release download")
    directory=""
    previous=""
    for argument in "$@"; do
      [ "$previous" = "--dir" ] && directory="$argument"
      previous="$argument"
    done
    cp "$FAKE_LOGS"/draft-assets/*_SHA256SUMS "$FAKE_LOGS"/draft-assets/*_SHA256SUMS.sig "$directory/"
    if [ -n "${FAKE_DOWNLOADED_SUMS:-}" ]; then
      echo "$FAKE_DOWNLOADED_SUMS" >"$directory/terraform-provider-x-twitter-scraper_1.2.3_SHA256SUMS"
    fi
    ;;
  "release edit") ;;
  *)
    echo "unexpected gh $*" >&2
    exit 97
    ;;
esac
EOF
  commit_repo
  add_origin
}

remote_tags() {
  sandbox_git ls-remote --tags origin
}

order_without_paths() {
  log_of order | sed "s|$sandbox|SANDBOX|g; s|--notes-file [^ ]*|--notes-file NOTES|g; s|--dir [^ ]*|--dir DRAFT|g"
}

# The steps before bin/check-local that every run takes.
readonly lookups="gh auth status --hostname github.com
gh api repos/$repository/releases/tags/v1.2.3
gh api --paginate --slurp repos/$repository/releases?per_page=100
curl -fsS --max-time 30 --proto =https https://registry.terraform.io/v1/providers/Xquik-dev/x-twitter-scraper/versions
gpg --batch --with-colons --list-secret-keys 4900C25E85324F28"

test_a_dry_run_builds_every_asset_prints_the_listing_and_releases_nothing() {
  local digest
  prepare_release

  run_in_repo bin/release-local --dry-run

  assert_equal "" "$run_stderr" "stderr"
  assert_equal 0 "$run_status" "the exit code"
  digest="$(cd "$sandbox_repo/dist" && shasum -a 256 "${project}_1.2.3_linux_amd64.zip" | cut -d ' ' -f 1)"
  assert_contains "$run_stdout" "GPG key 4900C25E85324F28 is in the keyring. A dry run never signs." "stdout"
  assert_contains "$run_stdout" "Version  1.2.3
Tag  v1.2.3
Registry address  registry.terraform.io/Xquik-dev/x-twitter-scraper
Protocol versions  6.0
Platforms
  linux_amd64  ${project}_v1.2.3
  windows_amd64  ${project}_v1.2.3.exe
Release assets
  ${project}_1.2.3_linux_amd64.zip  " "stdout"
  assert_contains "$run_stdout" " bytes  sha256 $digest" "stdout"
  assert_contains "$run_stdout" "  ${project}_1.2.3_manifest.json  73 bytes  sha256 " "stdout"
  assert_contains "$run_stdout" "  ${project}_1.2.3_SHA256SUMS.sig  made by bin/release-local after this listing
Asset count  5 with the signature, 2 platforms
GitHub release title  v1.2.3
GitHub release notes from CHANGELOG.md

  ### Bug Fixes

  * sign the checksums from a local runner" "stdout"
  assert_ends_with "$run_stdout" "Dry run finished. Nothing was published, tagged or released." "stdout"
  assert_equal "$lookups
check-local --no-status
audit
goreleaser release --clean --skip=publish,validate,sign with GORELEASER_CURRENT_TAG=v1.2.3
check-release-debug-info" "$(order_without_paths)" "the steps"
  assert_equal "" "$(log_of security-arguments)" "the Keychain calls"
  assert_equal "" "$(remote_tags)" "the tags on origin"
  assert_equal "" "$(sandbox_git tag)" "the local tags"
}

test_a_dry_run_lists_the_same_checksums_as_a_release_of_the_same_commit() {
  local sums="${project}_1.2.3_SHA256SUMS" dry_run_checksums
  prepare_release

  run_in_repo bin/release-local --dry-run
  dry_run_checksums="$(cat "$sandbox_repo/dist/$sums")"
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service")
  run_in_repo bin/release-local

  assert_equal 0 "$run_status" "the exit code of the release"
  assert_contains "$(cat "$sandbox_repo/dist/linux/${project}_v1.2.3")" "built as v1.2.3" "the released Linux binary"
  assert_equal "$dry_run_checksums" "$(cat "$sandbox_logs/draft-assets/$sums")" "the uploaded checksums"
}

test_a_dry_run_keeps_a_local_tag_that_was_there_before() {
  prepare_release
  sandbox_git tag --annotate v1.2.3 --message "Release v1.2.3"

  run_in_repo bin/release-local --dry-run

  assert_equal 0 "$run_status" "the exit code"
  assert_contains "$run_stdout" "==> Reusing the local annotated tag v1.2.3" "stdout"
  assert_equal "tag" "$(sandbox_git cat-file -t v1.2.3)" "the local tag type"
  assert_equal "" "$(remote_tags)" "the tags on origin"
}

test_a_failed_dry_run_removes_the_tag_it_created() {
  prepare_release
  run_env=("FAKE_WRONG_BINARY=1")

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "" "$(sandbox_git tag)" "the local tags"
}

test_a_dry_run_needs_no_keychain_variable_and_no_signing_key() {
  prepare_release
  run_env=("FAKE_NO_KEY=1")

  run_in_repo bin/release-local --dry-run

  assert_equal 0 "$run_status" "the exit code"
  assert_contains "$run_stdout" "GPG key 4900C25E85324F28 is not in the keyring. A real release would stop here." "stdout"
  assert_ends_with "$run_stdout" "Dry run finished. Nothing was published, tagged or released." "stdout"
}

test_a_dry_run_reports_what_github_and_the_registry_already_have_and_goes_on() {
  prepare_release
  run_env=("FAKE_PUBLISHED=1" 'FAKE_REGISTRY_VERSIONS={"version": "1.2.3"}')

  run_in_repo bin/release-local --dry-run

  assert_equal 0 "$run_status" "the exit code"
  assert_contains "$run_stdout" "GitHub release v1.2.3 is already published. A real release would stop here." "stdout"
  assert_contains "$run_stdout" "The Terraform Registry already has 1.2.3. A real release would stop here." "stdout"
  assert_ends_with "$run_stdout" "Dry run finished. Nothing was published, tagged or released." "stdout"
}

test_a_real_run_without_the_keychain_variable_stops_at_once() {
  prepare_release

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "RELEASE_SIGNING_KEYCHAIN_SERVICE is not set. Set it to the Keychain item that holds the GPG passphrase, then rerun." "$run_stderr" "stderr"
  assert_equal "" "$run_stdout" "stdout"
  assert_equal "" "$(log_of order)" "the steps"
}

test_a_real_run_with_an_empty_keychain_variable_stops_at_once() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=")

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "RELEASE_SIGNING_KEYCHAIN_SERVICE is not set. Set it to the Keychain item that holds the GPG passphrase, then rerun." "$run_stderr" "stderr"
  assert_equal "" "$(log_of order)" "the steps"
}

# Puts the 2 commands of 1 pipeline in pipeline order. The shell starts both
# at once, so either may record its call first.
pipeline_in_order() {
  awk -v first="$1" -v second="$2" '
    held != "" && index($0, first) == 1 { print; print held; held = ""; next }
    held != "" { print held; held = "" }
    index($0, second) == 1 { held = $0; next }
    { print }
    END { if (held != "") print held }
  '
}

test_a_release_signs_through_a_pipe_uploads_then_tags_and_publishes() {
  local sums="${project}_1.2.3_SHA256SUMS"
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service")

  run_in_repo bin/release-local

  assert_equal "" "$run_stderr" "stderr"
  assert_equal 0 "$run_status" "the exit code"
  assert_ends_with "$run_stdout" "Released v1.2.3 at https://github.com/$repository/releases/tag/v1.2.3. The Terraform Registry reads it from GitHub." "stdout"
  assert_equal "$lookups
security find-generic-password -s $test_service
check-local
audit
goreleaser release --clean --skip=publish,validate,sign with GORELEASER_CURRENT_TAG=v1.2.3
check-release-debug-info
security find-generic-password -s $test_service -w
gpg --batch --yes --pinentry-mode loopback --passphrase-fd 0 --local-user $test_fingerprint --output dist/$sums.sig --detach-sign dist/$sums
gpg --batch --status-fd 1 --verify dist/$sums.sig dist/$sums
gh release create v1.2.3 --repo $repository --draft --title v1.2.3 --target $(sandbox_git rev-parse HEAD) --notes-file NOTES
gh release upload v1.2.3 dist/$sums dist/$sums.sig dist/${project}_1.2.3_linux_amd64.zip dist/${project}_1.2.3_manifest.json dist/${project}_1.2.3_windows_amd64.zip --repo $repository --clobber
gh release view v1.2.3 --repo $repository --json assets --jq [.assets[].name] | sort
gh release download v1.2.3 --repo $repository --dir DRAFT --pattern $sums --pattern $sums.sig
gh release edit v1.2.3 --repo $repository --draft=false --verify-tag --title v1.2.3 --notes-file NOTES --latest" \
    "$(order_without_paths | pipeline_in_order "security find-generic-password -s $test_service -w" "gpg --batch --yes --pinentry-mode loopback")" "the steps"
  assert_equal "$test_passphrase" "$(log_of gpg-input)" "what gpg read on standard input"
  assert_not_contains "$(log_of gpg-arguments)" "$test_passphrase" "the gpg arguments"
  assert_not_contains "$(log_of security-arguments)" "$test_passphrase" "the security arguments"
  assert_not_contains "$(log_of order)" "$test_passphrase" "the commands that ran"
  assert_not_contains "$run_stdout$run_stderr" "$test_passphrase" "the output"
  cmp -s "$sandbox_logs/draft-assets/$sums.sig" <(printf '\211\002\117\004\000binary signature\n') ||
    fail_test "The uploaded signature differs from the one gpg wrote."
  assert_contains "$(remote_tags)" "refs/tags/v1.2.3" "the tags on origin"
  assert_equal "tag" "$(sandbox_git cat-file -t v1.2.3)" "the tag type"
}

test_a_release_reuses_an_existing_draft() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" 'FAKE_RELEASES=[[{"tag_name": "v1.2.3", "draft": true}]]')

  run_in_repo bin/release-local

  assert_equal 0 "$run_status" "the exit code"
  assert_not_contains "$(log_of order)" "gh release create" "the steps"
  assert_contains "$(log_of order)" "gh release upload v1.2.3" "the steps"
}

test_a_missing_signing_key_stops_a_real_run_before_the_checks() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" "FAKE_NO_KEY=1")

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "GPG secret key 4900C25E85324F28 is not in the keyring. Import it, then rerun." "$run_stderr" "stderr"
  assert_not_contains "$(log_of order)" "check-local" "the steps"
}

test_a_missing_keychain_item_stops_a_real_run_before_the_checks() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" "FAKE_NO_KEYCHAIN_ITEM=1")

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "The Keychain has no item named by RELEASE_SIGNING_KEYCHAIN_SERVICE. Add the GPG passphrase there, then rerun." "$run_stderr" "stderr"
  assert_not_contains "$(log_of order)" "check-local" "the steps"
}

test_a_signature_that_does_not_verify_stops_before_the_tag_and_the_upload() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" "FAKE_BAD_SIGNATURE=1")

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "The signature does not verify with GPG key 4900C25E85324F28. Check the key, then rerun." "$run_stderr" "stderr"
  assert_not_contains "$(log_of order)" "gh release" "the steps"
  assert_equal "" "$(sandbox_git tag)" "the local tags"
}

test_an_armored_signature_stops_before_the_tag_and_the_upload() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" "FAKE_ARMORED=1")

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "gpg wrote an ASCII armored signature. Remove armor from the gpg configuration, then rerun." "$run_stderr" "stderr"
  assert_not_contains "$(log_of order)" "gh release" "the steps"
}

test_a_published_release_stops_a_real_run_before_the_checks() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" "FAKE_PUBLISHED=1")

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "GitHub release v1.2.3 is already published. Bump the version, then rerun." "$run_stderr" "stderr"
  assert_not_contains "$(log_of order)" "check-local" "the steps"
}

test_a_version_the_registry_has_stops_a_real_run_before_the_checks() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" 'FAKE_REGISTRY_VERSIONS={"version": "1.2.3"}')

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "The Terraform Registry already has 1.2.3. Bump the version, then rerun." "$run_stderr" "stderr"
  assert_not_contains "$(log_of order)" "check-local" "the steps"
}

test_a_registry_that_does_not_answer_stops_the_run() {
  prepare_release
  run_env=("FAKE_REGISTRY_FAILS=1")

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "The Terraform Registry did not answer. Check the network, then rerun." "$run_stderr" "stderr"
}

test_a_failed_release_lookup_stops_the_run_before_the_checks() {
  prepare_release
  run_env=("FAKE_TAG_LOOKUP_FAILS=1")

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "GitHub did not answer for release v1.2.3. Check the network and gh auth status, then rerun." "$run_stderr" "stderr"
  assert_not_contains "$(log_of order)" "check-local" "the steps"
}

test_two_drafts_for_the_tag_stop_the_run() {
  prepare_release
  run_env=('FAKE_RELEASES=[[{"tag_name": "v1.2.3", "draft": true}, {"tag_name": "v1.2.3", "draft": true}]]')

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "Several draft releases use v1.2.3. Delete the extra drafts on GitHub, then rerun." "$run_stderr" "stderr"
}

test_a_draft_missing_an_asset_after_the_upload_stops_before_the_tag() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" "FAKE_VIEW_ASSETS=[\"${project}_1.2.3_SHA256SUMS\"]")

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "Draft release v1.2.3 does not hold exactly the built assets. Fix the draft on GitHub, then rerun." "$run_stderr" "stderr"
  assert_equal "" "$(remote_tags)" "the tags on origin"
  assert_not_contains "$(log_of order)" "gh release edit" "the steps"
}

test_checksums_that_differ_on_the_draft_stop_before_the_tag() {
  prepare_release
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service" "FAKE_DOWNLOADED_SUMS=0000  ${project}_1.2.3_linux_amd64.zip")

  run_in_repo bin/release-local

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "The SHA256SUMS files on draft v1.2.3 differ from this build. Rerun to upload them again." "$run_stderr" "stderr"
  assert_equal "" "$(remote_tags)" "the tags on origin"
}

test_a_prerelease_is_published_without_the_latest_label() {
  prepare_release 1.3.0-beta.1
  run_env=("RELEASE_SIGNING_KEYCHAIN_SERVICE=$test_service")

  run_in_repo bin/release-local

  assert_equal 0 "$run_status" "the exit code"
  assert_contains "$(log_of order)" "--draft=false --verify-tag --title v1.3.0-beta.1" "the steps"
  assert_ends_with "$(log_of order)" "--prerelease --latest=false" "the steps"
}

test_checksums_without_the_registry_manifest_stop_the_run() {
  prepare_release
  run_env=("FAKE_SUMS_WITHOUT_MANIFEST=1")

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "dist/${project}_1.2.3_SHA256SUMS must list every zip and the manifest, and nothing else. Check the checksum section of .goreleaser.yml." "$run_stderr" "stderr"
  assert_not_contains "$run_stdout" "Dry run finished" "stdout"
}

test_a_zip_without_the_registry_binary_name_stops_the_run() {
  prepare_release
  run_env=("FAKE_WRONG_BINARY=1")

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "${project}_1.2.3_linux_amd64.zip must hold the binary ${project}_v1.2.3. Check the binary name in .goreleaser.yml." "$run_stderr" "stderr"
}

test_a_build_that_changes_tracked_files_stops_the_run() {
  prepare_release
  run_env=("FAKE_DIRTY_BUILD=1")

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "The build changed tracked files. Fix the build so it leaves the tree alone, then rerun." "$run_stderr" "stderr"
}

test_a_tag_that_origin_already_has_stops_the_run() {
  prepare_release
  sandbox_git tag v1.2.3
  sandbox_git push --quiet origin v1.2.3
  sandbox_git tag -d v1.2.3 >/dev/null

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "Tag v1.2.3 already exists on origin. Bump the version, then rerun." "$run_stderr" "stderr"
}

test_a_working_tree_with_changes_stops_the_run() {
  prepare_release
  write_file notes.txt <<'EOF'
Uncommitted work.
EOF

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "The working tree has changes. Commit or remove them, then rerun." "$run_stderr" "stderr"
}

test_a_branch_other_than_main_stops_the_run() {
  prepare_release
  sandbox_git switch --quiet -c feature

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "The current branch is not main. Switch to main, then rerun." "$run_stderr" "stderr"
}

test_a_main_that_origin_lacks_stops_the_run() {
  prepare_release
  write_file extra.txt <<'EOF'
Not pushed.
EOF
  sandbox_git add extra.txt
  sandbox_git commit --quiet --message "Local only"

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "main differs from origin/main. Pull or push main, then rerun." "$run_stderr" "stderr"
}

test_a_version_without_a_changelog_section_stops_the_run() {
  prepare_release
  write_file CHANGELOG.md <<'EOF'
# Changelog

## 0.0.1 (2026-01-01)

* first release
EOF
  sandbox_git commit --quiet --all --message "Drop the section"
  sandbox_git push --quiet origin main

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "CHANGELOG.md has no section for 1.2.3. Merge the release pull request first." "$run_stderr" "stderr"
}

test_a_version_file_that_disagrees_stops_the_run() {
  prepare_release
  write_file internal/version.go <<'EOF'
package internal

const PackageVersion = "1.2.2" // x-release-please-version
EOF
  sandbox_git commit --quiet --all --message "Old version"
  sandbox_git push --quiet origin main

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_equal "internal/version.go does not say 1.2.3. Merge the release pull request first." "$run_stderr" "stderr"
}

test_an_unknown_option_stops_with_the_usage_line() {
  prepare_release

  run_in_repo bin/release-local --publish

  assert_equal 2 "$run_status" "the exit code"
  assert_equal "Unknown option --publish. Run bin/release-local or bin/release-local --dry-run." "$run_stderr" "stderr"
  assert_equal "" "$(log_of order)" "the steps"
}

run_tests
