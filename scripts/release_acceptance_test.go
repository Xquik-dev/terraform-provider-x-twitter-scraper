//go:build linux

// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"
)

type releaseFixture struct {
	directory   string
	services    string
	artifacts   string
	names       []string
	commit      string
	fingerprint string
	keyring     string
}

func releaseWrite(t *testing.T, name string, value []byte) {
	t.Helper()
	if err := os.WriteFile(name, value, 0o700); err != nil {
		t.Fatal(err)
	}
}

func releaseJSON(t *testing.T, name string, value any) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	releaseWrite(t, name, data)
}

func newReleaseFixture(t *testing.T, nativeAppID int) releaseFixture {
	t.Helper()
	directory := t.TempDir()
	if err := os.Mkdir(filepath.Join(directory, "scripts"), 0o700); err != nil {
		t.Fatal(err)
	}
	publisher, err := os.ReadFile("release")
	if err != nil {
		t.Fatal(err)
	}
	releaseWrite(t, filepath.Join(directory, "scripts", "release"), publisher)
	releaseWrite(t, filepath.Join(directory, ".release-please-manifest.json"), []byte(`{".":"0.7.5"}`))
	var policyApp any = nativeAppID
	if nativeAppID == 0 {
		policyApp = nil
		nativeAppID = 345678
	}
	releaseJSON(t, filepath.Join(directory, "release-policy.json"), map[string]any{"native_check_app_id": policyApp, "goreleaser_version": "2.18.2", "required_checks": []string{"Lint", "Test", "Analyze Go", "Dependency Security", "Fuzz Multipart Encoding", "Reproducible Release Build"}, "platforms": []string{"freebsd_386", "freebsd_amd64", "freebsd_arm", "freebsd_arm64", "linux_386", "linux_amd64", "linux_arm", "linux_arm64", "darwin_amd64", "darwin_arm64", "windows_386", "windows_amd64", "windows_arm64"}})
	releaseWrite(t, filepath.Join(directory, "CHANGELOG.md"), []byte("0.7.5\n"))
	releaseGit(t, directory, "init", "-q", "--initial-branch=main")
	releaseGit(t, directory, "config", "user.name", "Release test")
	releaseGit(t, directory, "config", "user.email", "release@example.test")
	releaseGit(t, directory, "config", "commit.gpgsign", "false")
	releaseGit(t, directory, "add", ".")
	releaseGit(t, directory, "commit", "-qm", "release fixture")
	releaseGit(t, directory, "tag", "v0.7.5")
	remote := filepath.Join(directory, ".git", "origin.git")
	releaseGit(t, directory, "init", "-q", "--bare", remote)
	releaseGit(t, directory, "remote", "add", "origin", remote)
	releaseGit(t, directory, "push", "-q", "origin", "HEAD:refs/heads/main", "refs/tags/v0.7.5")
	readGit := func(ref string) string {
		command := exec.Command("git", "rev-parse", ref)
		command.Dir = directory
		output, err := command.Output()
		if err != nil {
			t.Fatal(err)
		}
		return strings.TrimSpace(string(output))
	}
	fixture := releaseFixture{directory: directory, services: filepath.Join(directory, ".git", "services"), artifacts: filepath.Join(directory, ".git", "artifacts"), commit: readGit("HEAD")}
	for _, name := range []string{fixture.services, fixture.artifacts} {
		if err := os.Mkdir(name, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	checks, runs, qualified := []any{}, []any{}, []any{}
	for index, name := range []string{"Lint", "Test", "Analyze Go", "Dependency Security", "Fuzz Multipart Encoding", "Reproducible Release Build"} {
		checks = append(checks, map[string]any{"context": name, "app_id": nativeAppID})
		runs = append(runs, map[string]any{"name": name, "id": index + 1, "app": map[string]any{"id": nativeAppID, "slug": "native-release-checks"}, "head_sha": fixture.commit, "status": "completed", "conclusion": "success"})
		qualified = append(qualified, map[string]any{"name": name, "id": index + 1, "app_id": nativeAppID})
	}
	releaseJSON(t, filepath.Join(fixture.services, "tag.json"), map[string]any{"ref": "refs/tags/v0.7.5", "object": map[string]any{"type": "commit", "sha": fixture.commit}})
	releaseJSON(t, filepath.Join(fixture.services, "main.json"), map[string]any{"object": map[string]any{"sha": fixture.commit}})
	releaseJSON(t, filepath.Join(fixture.services, "comparison.json"), map[string]any{"status": "identical", "base_commit": map[string]any{"sha": fixture.commit}, "merge_base_commit": map[string]any{"sha": fixture.commit}, "commits": []any{}})
	releaseJSON(t, filepath.Join(fixture.services, "requirements.json"), map[string]any{"strict": true, "checks": checks})
	releaseJSON(t, filepath.Join(fixture.services, "runs.json"), []any{map[string]any{"check_runs": runs}})
	releaseWrite(t, filepath.Join(fixture.services, "gh"), []byte(releaseGitHubFake))
	releaseWrite(t, filepath.Join(fixture.services, "curl"), []byte(releaseRegistryFake))
	fixture.keyring = filepath.Join(directory, ".git", "gnupg")
	if err := os.Mkdir(fixture.keyring, 0o700); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		command := releaseCommand(ctx, directory, "gpgconf", "--homedir", fixture.keyring, "--kill", "gpg-agent")
		if output, err := command.CombinedOutput(); err != nil {
			t.Errorf("GPG agent cleanup: %v\n%s", err, output)
		}
	})
	fixture.gpg(t, "--passphrase", "", "--quick-generate-key", "release@example.test", "ed25519", "sign", "0")
	for _, line := range strings.Split(string(fixture.gpg(t, "--with-colons", "--list-keys")), "\n") {
		fields := strings.Split(line, ":")
		if fields[0] == "fpr" {
			fixture.fingerprint = fields[9]
			break
		}
	}
	if fixture.fingerprint == "" {
		t.Fatal("test signing key has no fingerprint")
	}
	armor := string(fixture.gpg(t, "--armor", "--export", fixture.fingerprint))
	releaseJSON(t, filepath.Join(fixture.services, "registry.json"), map[string]any{"signing_keys": map[string]any{"gpg_public_keys": []any{map[string]any{"ascii_armor": armor}}}})
	releaseJSON(t, filepath.Join(fixture.artifacts, "metadata.json"), map[string]any{"commit": fixture.commit, "version": "0.7.5", "tag": "v0.7.5"})
	prefix := "terraform-provider-x-twitter-scraper_0.7.5_"
	for _, platform := range []string{"freebsd_386", "freebsd_amd64", "freebsd_arm", "freebsd_arm64", "linux_386", "linux_amd64", "linux_arm", "linux_arm64", "darwin_amd64", "darwin_arm64", "windows_386", "windows_amd64", "windows_arm64"} {
		name := prefix + platform + ".zip"
		fixture.names = append(fixture.names, name)
		var archive bytes.Buffer
		writer := zip.NewWriter(&archive)
		binary := "terraform-provider-x-twitter-scraper_v0.7.5"
		if strings.HasPrefix(platform, "windows_") {
			binary += ".exe"
		}
		entry, err := writer.Create(binary)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := entry.Write([]byte("binary fixture for " + platform)); err != nil {
			t.Fatal(err)
		}
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
		releaseWrite(t, filepath.Join(fixture.artifacts, name), archive.Bytes())
	}
	fixture.names = append(fixture.names, prefix+"manifest.json", prefix+"source.json")
	releaseWrite(t, filepath.Join(fixture.artifacts, prefix+"manifest.json"), []byte(`{"version":1,"metadata":{"protocol_versions":["6.0"]}}`))
	releaseJSON(t, filepath.Join(fixture.artifacts, prefix+"source.json"), map[string]any{"commit": fixture.commit, "version": "0.7.5", "tree": readGit("HEAD^{tree}"), "checks": qualified})
	fixture.checksums(t)
	fixture.gpg(t, "--local-user", fixture.fingerprint, "--output", filepath.Join(fixture.artifacts, prefix+"SHA256SUMS.sig"), "--detach-sign", filepath.Join(fixture.artifacts, prefix+"SHA256SUMS"))
	fixture.names = append(fixture.names, prefix+"SHA256SUMS", prefix+"SHA256SUMS.sig")
	fixture.release(t, true, fixture.names)
	return fixture
}

func (fixture releaseFixture) gpg(t *testing.T, arguments ...string) []byte {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	command := releaseCommand(ctx, fixture.directory, "gpg", append([]string{"--homedir", fixture.keyring, "--batch", "--pinentry-mode", "loopback"}, arguments...)...)
	var stderr strings.Builder
	command.Stderr = &stderr
	output, err := command.Output()
	if err != nil {
		t.Fatalf("gpg: %v\n%s", err, stderr.String())
	}
	return output
}

func (fixture releaseFixture) checksums(t *testing.T) {
	t.Helper()
	var data strings.Builder
	for _, name := range fixture.names {
		if strings.HasSuffix(name, "SHA256SUMS") || strings.HasSuffix(name, "SHA256SUMS.sig") {
			continue
		}
		content, err := os.ReadFile(filepath.Join(fixture.artifacts, name))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintf(&data, "%x  %s\n", sha256.Sum256(content), name)
	}
	releaseWrite(t, filepath.Join(fixture.artifacts, "terraform-provider-x-twitter-scraper_0.7.5_SHA256SUMS"), []byte(data.String()))
}

func (fixture releaseFixture) release(t *testing.T, draft bool, names []string) {
	t.Helper()
	assets := []any{}
	for _, name := range names {
		content, err := os.ReadFile(filepath.Join(fixture.artifacts, name))
		if err != nil {
			t.Fatal(err)
		}
		assets = append(assets, map[string]any{"name": name, "digest": fmt.Sprintf("sha256:%x", sha256.Sum256(content))})
	}
	releaseJSON(t, filepath.Join(fixture.services, "release.json"), map[string]any{"isDraft": draft, "tagName": "v0.7.5", "targetCommitish": fixture.commit, "assets": assets})
}

func (fixture releaseFixture) run(t *testing.T, wantExit int, wantError string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	command := releaseCommand(ctx, fixture.directory, "bash", "scripts/release", "publish", fixture.artifacts)
	command.Env = append(os.Environ(), "PATH="+fixture.services+":"+os.Getenv("PATH"), "RELEASE_FIXTURES="+fixture.services, "GPG_FINGERPRINT="+fixture.fingerprint, "GNUPGHOME="+fixture.keyring, "RELEASE_COMMIT="+fixture.commit)
	var stdout, stderr strings.Builder
	command.Stdout, command.Stderr = &stdout, &stderr
	err := command.Run()
	exit := 0
	if err != nil {
		failure, ok := err.(*exec.ExitError)
		if !ok {
			t.Fatal(err)
		}
		exit = failure.ExitCode()
	}
	var result struct {
		Status string `json:"status"`
	}
	wantStatus := "completed"
	if wantExit != 0 {
		wantStatus = "failed"
	}
	if exit != wantExit || !strings.Contains(stderr.String(), wantError) || json.Unmarshal([]byte(stdout.String()), &result) != nil || result.Status != wantStatus {
		t.Fatalf("publisher exit %d; stdout %q; stderr %s; want exit %d and %q", exit, stdout.String(), stderr.String(), wantExit, wantError)
	}
}

func (fixture releaseFixture) mutations(t *testing.T, want []string) {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(fixture.services, "mutations"))
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	var got []string
	if len(content) != 0 {
		got = strings.Split(strings.TrimSpace(string(content)), "\n")
	}
	if !slices.Equal(got, want) {
		t.Fatalf("GitHub mutations: %v; want %v", got, want)
	}
}

func (fixture releaseFixture) verifyAssets(t *testing.T) {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(fixture.services, "release.json"))
	if err != nil {
		t.Fatal(err)
	}
	var release struct {
		Assets []struct {
			Name   string `json:"name"`
			Digest string `json:"digest"`
		} `json:"assets"`
	}
	if err := json.Unmarshal(content, &release); err != nil {
		t.Fatal(err)
	}
	want := map[string]string{}
	for _, name := range fixture.names {
		content, err := os.ReadFile(filepath.Join(fixture.artifacts, name))
		if err != nil {
			t.Fatal(err)
		}
		want[name] = fmt.Sprintf("sha256:%x", sha256.Sum256(content))
	}
	for _, asset := range release.Assets {
		if want[asset.Name] == "" || want[asset.Name] != asset.Digest {
			t.Fatalf("unexpected published asset: %+v", asset)
		}
		delete(want, asset.Name)
	}
	if len(want) != 0 {
		t.Fatalf("missing published assets: %v", want)
	}
}

func releaseEdit(t *testing.T, name, expression string) {
	t.Helper()
	output, err := exec.Command("jq", expression, name).Output()
	if err != nil {
		t.Fatal(err)
	}
	releaseWrite(t, name, output)
}

func TestReleaseRequiresCanonicalSource(t *testing.T) {
	for _, scenario := range []struct{ name, file, change, want string }{
		{"wrong tag", "tag.json", `.object.sha = "1111111111111111111111111111111111111111"`, "Publish the matching immutable release tag first."},
		{"different tag ref", "tag.json", `.ref = "refs/tags/v0.7.4"`, "Publish the matching immutable release tag first."},
		{"canonical main excludes candidate", "comparison.json", `.status = "diverged" | .merge_base_commit.sha = "1111111111111111111111111111111111111111"`, "Merge the release into the canonical repository first."},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			fixture := newReleaseFixture(t, 345678)
			releaseEdit(t, filepath.Join(fixture.services, scenario.file), scenario.change)
			fixture.run(t, 1, scenario.want)
			fixture.mutations(t, []string{})
		})
	}
}

func TestReleaseAcceptsCanonicalAncestorWithoutCommitList(t *testing.T) {
	fixture := newReleaseFixture(t, 345678)
	releaseEdit(t, filepath.Join(fixture.services, "main.json"), `.object.sha = "1111111111111111111111111111111111111111"`)
	releaseEdit(t, filepath.Join(fixture.services, "comparison.json"), `.status = "ahead" | .merge_base_commit = .base_commit | .commits = []`)
	fixture.run(t, 0, "")
}

func TestReleaseRequiresCurrentNativeChecks(t *testing.T) {
	for _, scenario := range []struct{ name, file, change, want string }{
		{"null policy", "requirements.json", `.`, "Confirm the native check publisher binding before releasing."},
		{"wrong App", "requirements.json", `.checks[0].app_id = 345679`, "Bind every release check to its required GitHub App."},
		{"Actions App", "requirements.json", `.checks[].app_id = 15368`, "Bind every release check to its required GitHub App."},
		{"stale head", "runs.json", `.[0].check_runs[0].head_sha = "1111111111111111111111111111111111111111"`, "Complete every required check on the exact release commit."},
		{"missing context", "runs.json", `.[0].check_runs |= .[1:]`, "Complete every required check on the exact release commit."},
		{"newer failure", "runs.json", `.[0].check_runs += [(.[0].check_runs[0] | .id = 100 | .conclusion = "failure")]`, "Complete every required check on the exact release commit."},
		{"reduced protection", "requirements.json", `.checks |= .[0:1]`, "Restore the complete required release checks."},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			appID := 345678
			if scenario.name == "null policy" {
				appID = 0
			}
			if scenario.name == "Actions App" {
				appID = 15368
			}
			fixture := newReleaseFixture(t, appID)
			releaseEdit(t, filepath.Join(fixture.services, scenario.file), scenario.change)
			fixture.run(t, 1, scenario.want)
			fixture.mutations(t, []string{})
		})
	}
}

func TestReleaseRequiresRegisteredSigner(t *testing.T) {
	for _, scenario := range []string{"empty registry", "different registered key"} {
		t.Run(scenario, func(t *testing.T) {
			fixture := newReleaseFixture(t, 345678)
			keys := []any{}
			if scenario == "different registered key" {
				fixture.gpg(t, "--passphrase", "", "--quick-generate-key", "other@example.test", "ed25519", "sign", "0")
				armor := string(fixture.gpg(t, "--armor", "--export", "other@example.test"))
				keys = append(keys, map[string]any{"ascii_armor": armor})
			}
			releaseJSON(t, filepath.Join(fixture.services, "registry.json"), map[string]any{"signing_keys": map[string]any{"gpg_public_keys": keys}})
			fixture.run(t, 1, "Use the signing key registered with Terraform.")
			fixture.mutations(t, []string{})
		})
	}
}

func TestReleaseRequiresRecordedNativeQualification(t *testing.T) {
	for _, scenario := range []string{"unknown recorded result", "wrong recorded App", "failed recorded result"} {
		t.Run(scenario, func(t *testing.T) {
			fixture := newReleaseFixture(t, 345678)
			source := filepath.Join(fixture.artifacts, "terraform-provider-x-twitter-scraper_0.7.5_source.json")
			switch scenario {
			case "unknown recorded result":
				releaseEdit(t, source, `.checks[0].id = 999`)
			case "wrong recorded App":
				releaseEdit(t, source, `.checks[0].app_id = 345679`)
			case "failed recorded result":
				releaseEdit(t, filepath.Join(fixture.services, "runs.json"), `.[0].check_runs += [(.[0].check_runs[0] | .id = 100)] | .[0].check_runs[0].conclusion = "failure"`)
			}
			fixture.checksums(t)
			checksums := filepath.Join(fixture.artifacts, "terraform-provider-x-twitter-scraper_0.7.5_SHA256SUMS")
			fixture.gpg(t, "--yes", "--local-user", fixture.fingerprint, "--output", checksums+".sig", "--detach-sign", checksums)
			fixture.release(t, true, fixture.names)
			fixture.run(t, 1, "Verify the recorded native check results.")
			fixture.mutations(t, []string{})
		})
	}
}

func TestReleaseRequiresCompleteArtifactSet(t *testing.T) {
	for _, scenario := range []string{"missing archive", "extra archive", "omitted checksum", "corrupt archive", "duplicate checksum"} {
		t.Run(scenario, func(t *testing.T) {
			fixture := newReleaseFixture(t, 345678)
			checksums := filepath.Join(fixture.artifacts, "terraform-provider-x-twitter-scraper_0.7.5_SHA256SUMS")
			switch scenario {
			case "missing archive":
				if err := os.Remove(filepath.Join(fixture.artifacts, fixture.names[0])); err != nil {
					t.Fatal(err)
				}
			case "extra archive":
				releaseWrite(t, filepath.Join(fixture.artifacts, "terraform-provider-x-twitter-scraper_0.7.5_windows_arm.zip"), []byte("unsupported archive"))
			case "omitted checksum", "duplicate checksum":
				content, err := os.ReadFile(checksums)
				if err != nil {
					t.Fatal(err)
				}
				lines := strings.Split(string(content), "\n")
				if scenario == "omitted checksum" {
					lines = lines[1:]
				} else {
					lines = append(lines[:len(lines)-1], lines[0], "")
				}
				releaseWrite(t, checksums, []byte(strings.Join(lines, "\n")))
				fixture.gpg(t, "--yes", "--local-user", fixture.fingerprint, "--output", checksums+".sig", "--detach-sign", checksums)
				fixture.release(t, true, fixture.names)
			case "corrupt archive":
				releaseWrite(t, filepath.Join(fixture.artifacts, fixture.names[0]), []byte("corrupted"))
			}
			fixture.run(t, 1, "Verify the exact release artifact set.")
			fixture.mutations(t, []string{})
		})
	}
}

func TestReleaseVerifiesUploadedDigestsBeforePublication(t *testing.T) {
	fixture := newReleaseFixture(t, 345678)
	fixture.release(t, true, fixture.names[1:])
	releaseWrite(t, filepath.Join(fixture.services, "corrupt-upload"), []byte("true"))
	fixture.run(t, 1, "Verify every uploaded artifact before publishing.")
	content, err := os.ReadFile(filepath.Join(fixture.services, "release.json"))
	if err != nil {
		t.Fatal(err)
	}
	var release struct {
		IsDraft bool `json:"isDraft"`
	}
	if json.Unmarshal(content, &release) != nil || !release.IsDraft {
		t.Fatal("corrupt upload was published")
	}
}

func TestReleaseRetryPreservesImmutableAssets(t *testing.T) {
	for _, scenario := range []string{"matching draft", "matching published release", "newer successful check", "conflicting asset", "missing published asset"} {
		t.Run(scenario, func(t *testing.T) {
			fixture := newReleaseFixture(t, 345678)
			switch scenario {
			case "matching draft":
				fixture.release(t, true, fixture.names[1:])
			case "matching published release":
				fixture.release(t, false, fixture.names)
			case "newer successful check":
				fixture.release(t, false, fixture.names)
				releaseEdit(t, filepath.Join(fixture.services, "runs.json"), `.[0].check_runs += [(.[0].check_runs[0] | .id = 100)]`)
			case "conflicting asset":
				releaseJSON(t, filepath.Join(fixture.services, "release.json"), map[string]any{"isDraft": true, "tagName": "v0.7.5", "targetCommitish": fixture.commit, "assets": []any{map[string]any{"name": fixture.names[0], "digest": "sha256:conflicting"}}})
			case "missing published asset":
				fixture.release(t, false, fixture.names[1:])
			}
			if scenario == "conflicting asset" {
				fixture.run(t, 1, "Conflicting release artifact:")
				fixture.mutations(t, []string{})
				return
			}
			if scenario == "missing published asset" {
				fixture.run(t, 1, "Keep published release artifacts immutable.")
				fixture.mutations(t, []string{})
				return
			}
			fixture.run(t, 0, "")
			content, err := os.ReadFile(filepath.Join(fixture.services, "release.json"))
			if err != nil {
				t.Fatal(err)
			}
			var release struct {
				IsDraft bool `json:"isDraft"`
				Assets  []struct {
					Name string `json:"name"`
				} `json:"assets"`
			}
			if json.Unmarshal(content, &release) != nil || release.IsDraft || len(release.Assets) != 17 {
				t.Fatalf("incomplete published release: %s", content)
			}
			fixture.verifyAssets(t)
			if scenario == "matching draft" {
				fixture.mutations(t, []string{"upload:terraform-provider-x-twitter-scraper_0.7.5_freebsd_386.zip", "edit:"})
			} else {
				fixture.mutations(t, []string{})
			}
		})
	}
}

func TestReleasePublishesCompleteDraft(t *testing.T) {
	fixture := newReleaseFixture(t, 345678)
	releaseWrite(t, filepath.Join(fixture.services, "release.json"), []byte("null"))
	fixture.run(t, 0, "")
	fixture.verifyAssets(t)
	content, err := os.ReadFile(filepath.Join(fixture.services, "release.json"))
	if err != nil {
		t.Fatal(err)
	}
	var release struct {
		IsDraft bool `json:"isDraft"`
	}
	if json.Unmarshal(content, &release) != nil || release.IsDraft {
		t.Fatalf("draft was not published: %s", content)
	}
	content, err = os.ReadFile(filepath.Join(fixture.services, "mutations"))
	if err != nil {
		t.Fatal(err)
	}
	mutations := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(mutations) != 19 || mutations[0] != "create:" || mutations[len(mutations)-1] != "edit:" {
		t.Fatalf("GitHub publication phases: %v", mutations)
	}
	uploads := mutations[1 : len(mutations)-1]
	want := []string{}
	for _, name := range fixture.names {
		want = append(want, "upload:"+name)
	}
	slices.Sort(uploads)
	slices.Sort(want)
	if !slices.Equal(uploads, want) {
		t.Fatalf("GitHub uploads: %v; want %v", uploads, want)
	}
}

const releaseGitHubFake = `#!/usr/bin/env bash
set -euo pipefail
if [[ "$1" == api ]]; then
  endpoint=""
  for argument in "$@"; do [[ "$argument" != repos/* ]] || endpoint="$argument"; done
  repo='repos/Xquik-dev/terraform-provider-x-twitter-scraper'
  case "$endpoint" in
    "$repo/git/ref/tags/v0.7.5") file=tag.json ;;
    "$repo/git/ref/heads/main") file=main.json ;;
    "$repo/compare/$RELEASE_COMMIT...$(jq -r .object.sha "$RELEASE_FIXTURES/main.json")") file=comparison.json ;;
    "$repo/branches/main/protection/required_status_checks") file=requirements.json ;;
    "$repo/commits/$RELEASE_COMMIT/check-runs?filter=all&per_page=100") file=runs.json ;;
    "$repo/releases?per_page=100") jq -c 'if . == null then [[]] else [[{tag_name:.tagName,target_commitish:.targetCommitish,draft:.isDraft}]] end' "$RELEASE_FIXTURES/release.json"; exit ;;
    *) echo "Unexpected GitHub endpoint: $endpoint" >&2; exit 1 ;;
  esac
  if [[ "$*" == *'--jq .object.sha'* ]]; then jq -r '.object.sha' "$RELEASE_FIXTURES/$file"; else cat "$RELEASE_FIXTURES/$file"; fi
elif [[ "$1 $2" == 'release create' ]]; then
  draft=false
  for argument in "$@"; do
    case "$argument" in --draft) draft=true ;; --draft=false) exit 1 ;; esac
  done
  [[ "$3" == v0.7.5 && "$*" == *'--repo Xquik-dev/terraform-provider-x-twitter-scraper'* && "$*" == *"--target $RELEASE_COMMIT"* && "$draft" == true && "$*" == *'--verify-tag'* ]]
  printf 'create:\n' >> "$RELEASE_FIXTURES/mutations"
  jq -n --arg commit "$RELEASE_COMMIT" '{isDraft:true,tagName:"v0.7.5",targetCommitish:$commit,assets:[]}' > "$RELEASE_FIXTURES/release.json"
elif [[ "$1 $2" == 'release view' ]]; then
  [[ "$3" == v0.7.5 && "$*" == *'--repo Xquik-dev/terraform-provider-x-twitter-scraper'* ]]
  cat "$RELEASE_FIXTURES/release.json"
elif [[ "$1 $2" == 'release upload' ]]; then
  [[ "$3" == v0.7.5 && "$*" == *'--repo Xquik-dev/terraform-provider-x-twitter-scraper'* ]]
  name="$(basename "$4")"
  printf 'upload:%s\n' "$name" >> "$RELEASE_FIXTURES/mutations"
  digest="sha256:$(shasum -a 256 "$4" | cut -d ' ' -f 1)"
  [[ ! -e "$RELEASE_FIXTURES/corrupt-upload" ]] || digest=sha256:corrupted
  jq --arg name "$name" --arg digest "$digest" '.assets += [{name:$name,digest:$digest}]' "$RELEASE_FIXTURES/release.json" > "$RELEASE_FIXTURES/next.json"
  mv "$RELEASE_FIXTURES/next.json" "$RELEASE_FIXTURES/release.json"
elif [[ "$1 $2" == 'release edit' ]]; then
  [[ "$3" == v0.7.5 && "$*" == *'--repo Xquik-dev/terraform-provider-x-twitter-scraper'* && "$*" == *'--draft=false'* && "$*" == *'--verify-tag'* ]]
  printf 'edit:\n' >> "$RELEASE_FIXTURES/mutations"
  jq '.isDraft = false' "$RELEASE_FIXTURES/release.json" > "$RELEASE_FIXTURES/next.json"
  mv "$RELEASE_FIXTURES/next.json" "$RELEASE_FIXTURES/release.json"
else
  echo 'Unexpected GitHub mutation.' >&2; exit 1
fi
`

const releaseRegistryFake = `#!/usr/bin/env bash
set -euo pipefail
case "${@: -1}" in
  https://registry.terraform.io/v1/providers/xquik-dev/x-twitter-scraper/versions) printf '%s\n' '{"versions":[{"version":"0.7.4"}]}' ;;
  https://registry.terraform.io/v1/providers/xquik-dev/x-twitter-scraper/0.7.4/download/linux/amd64) cat "$RELEASE_FIXTURES/registry.json" ;;
  *) echo 'Unexpected Terraform registry request.' >&2; exit 1 ;;
esac
`
