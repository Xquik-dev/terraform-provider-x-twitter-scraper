//go:build linux

// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type goreleaserFixture struct {
	root, tools, caller string
	policy              map[string]any
}

func newGoReleaserFixture(t *testing.T, replacement map[string][]byte) goreleaserFixture {
	t.Helper()
	installed, err := exec.LookPath("goreleaser")
	if err != nil {
		t.Fatal("put the verified GoReleaser distribution on PATH:", err)
	}
	fixture := goreleaserFixture{root: t.TempDir(), tools: t.TempDir(), caller: t.TempDir()}
	for _, name := range []string{"goreleaser", "cosign", "goreleaser_Linux_x86_64.tar.gz", "checksums.txt", "checksums.txt.sigstore.json"} {
		path := filepath.Join(fixture.tools, name)
		if contents, replaced := replacement[name]; replaced {
			if contents != nil {
				releaseWrite(t, path, contents)
			}
		} else if err := os.Symlink(filepath.Join(filepath.Dir(installed), name), path); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Mkdir(filepath.Join(fixture.root, "bin"), 0o700); err != nil {
		t.Fatal(err)
	}
	launcher, err := os.ReadFile(filepath.Join("..", "bin", "run-goreleaser"))
	if err == nil {
		releaseWrite(t, filepath.Join(fixture.root, "bin", "run-goreleaser"), launcher)
	} else if !os.IsNotExist(err) {
		t.Fatal(err)
	}
	fixture.policy = map[string]any{
		"goreleaser_version": "2.18.2",
		"goreleaser_distribution": map[string]any{
			"platform":                "Linux_x86_64",
			"cosign_version":          "3.1.3",
			"certificate_identity":    "https://github.com/goreleaser/goreleaser/.github/workflows/release.yml@refs/tags/v2.18.2",
			"certificate_oidc_issuer": "https://token.actions.githubusercontent.com",
			"sha256": map[string]any{
				"goreleaser_Linux_x86_64.tar.gz": "0a96edc9d9bc594e4a41cc4d59467c182062910ab24d9d1f6dd7b667d32606d3",
				"checksums.txt":                  "0818c962a5a5e49abb7f721f4fef851d92aca7253821c1eb5496338e812c0a47",
				"checksums.txt.sigstore.json":    "47d83affc90012d6de53268fde64d5c5defe32e6155a30c71dc22eed37222d0d",
				"goreleaser":                     "d9fc6cdfaa7e8d44b072ade96e0d63b74857848d8890e4a52f8d69bcf5e49b7b",
				"cosign":                         "4629c757b7618056f8ddd7e2625ae9fdd94c0372a65049520bc7d9df9efc7f71",
			},
		},
	}
	return fixture
}

func (fixture goreleaserFixture) run(t *testing.T, wantExit int, arguments ...string) (string, string) {
	t.Helper()
	releaseJSON(t, filepath.Join(fixture.root, "release-policy.json"), fixture.policy)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	command := releaseCommand(ctx, fixture.caller, "bash", append([]string{filepath.Join(fixture.root, "bin", "run-goreleaser")}, arguments...)...)
	command.Env = append(os.Environ(), "PATH="+fixture.tools+":"+os.Getenv("PATH"), "XDG_CACHE_HOME="+filepath.Join(fixture.root, "sigstore-cache"))
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
	if exit != wantExit {
		t.Fatalf("launcher exit %d; stdout %s; stderr %s; want %d", exit, stdout.String(), stderr.String(), wantExit)
	}
	return stdout.String(), stderr.String()
}

func TestGoReleaserLauncherPreservesCallerConfiguration(t *testing.T) {
	fixture := newGoReleaserFixture(t, nil)
	releaseWrite(t, filepath.Join(fixture.caller, "caller.yml"), []byte("version: 2\nproject_name: caller-fixture\n"))
	releaseWrite(t, filepath.Join(fixture.caller, ".goreleaser.yml"), []byte("unknown_configuration_field: true\n"))
	releaseGit(t, fixture.caller, "init", "--quiet", "--initial-branch=main")
	releaseGit(t, fixture.caller, "add", ".")
	releaseGit(t, fixture.caller, "-c", "user.name=Release test", "-c", "user.email=release@example.com", "-c", "commit.gpgsign=false", "commit", "--quiet", "-m", "Caller configuration")
	releaseGit(t, fixture.caller, "remote", "add", "origin", "https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper.git")
	fixture.run(t, 0, "check", "--config", "caller.yml")
}

func TestGoReleaserLauncherPreservesCLIError(t *testing.T) {
	fixture := newGoReleaserFixture(t, nil)
	_, stderr := fixture.run(t, 1, "not-a-goreleaser-command")
	if !strings.Contains(stderr, "unknown command") {
		t.Fatalf("missing GoReleaser CLI error: %s", stderr)
	}
}

func TestGoReleaserLauncherRequiresQualifiedDistribution(t *testing.T) {
	for _, scenario := range []struct {
		name, file, field, value string
	}{
		{name: "missing archive", file: "goreleaser_Linux_x86_64.tar.gz"},
		{name: "corrupt archive", file: "goreleaser_Linux_x86_64.tar.gz", value: "corrupt"},
		{name: "corrupt checksums", file: "checksums.txt", value: "corrupt"},
		{name: "corrupt bundle", file: "checksums.txt.sigstore.json", value: "corrupt"},
		{name: "corrupt executable", file: "goreleaser", value: "corrupt"},
		{name: "corrupt verifier", file: "cosign", value: "corrupt"},
		{name: "wrong signer", field: "certificate_identity", value: "https://github.com/other/project/.github/workflows/release.yml@refs/tags/v2.18.2"},
		{name: "wrong issuer", field: "certificate_oidc_issuer", value: "https://accounts.google.com"},
		{name: "wrong runtime", field: "goreleaser_version", value: "2.18.1"},
		{name: "wrong verifier runtime", field: "cosign_version", value: "3.1.2"},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			replacements := map[string][]byte{}
			if scenario.file != "" {
				if scenario.value == "" {
					replacements[scenario.file] = nil
				} else {
					replacements[scenario.file] = []byte(scenario.value)
				}
			}
			fixture := newGoReleaserFixture(t, replacements)
			if scenario.field == "goreleaser_version" {
				fixture.policy[scenario.field] = scenario.value
			} else if scenario.field != "" {
				fixture.policy["goreleaser_distribution"].(map[string]any)[scenario.field] = scenario.value
			}
			stdout, _ := fixture.run(t, 1, "--version")
			if strings.Contains(stdout, "GitVersion:") {
				t.Fatalf("unqualified runtime reached the caller: %s", stdout)
			}
		})
	}
}
