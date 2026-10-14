//go:build linux

// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
)

func releaseCommand(ctx context.Context, directory, program string, arguments ...string) *exec.Cmd {
	command := exec.CommandContext(ctx, program, arguments...)
	command.Dir = directory
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	command.Cancel = func() error { return syscall.Kill(-command.Process.Pid, syscall.SIGKILL) }
	command.WaitDelay = 5 * time.Second
	return command
}

func releaseGit(t *testing.T, directory string, arguments ...string) {
	t.Helper()
	command := exec.Command("git", arguments...)
	command.Dir = directory
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %v\n%s", arguments, err, output)
	}
}

func TestReleaseRefusesInvalidSourceBeforePreparingArtifacts(t *testing.T) {
	for _, scenario := range []struct {
		name string
		want string
	}{
		{"missing tag", "Create the matching immutable release tag first."},
		{"tag on another commit", "Create the matching immutable release tag first."},
		{"unmerged commit", "Merge the release into main first."},
		{"dirty source", "Commit verified changes before release preparation."},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			directory := t.TempDir()
			if err := os.Mkdir(filepath.Join(directory, "scripts"), 0o700); err != nil {
				t.Fatal(err)
			}
			publisher, err := os.ReadFile("release")
			if err != nil && !os.IsNotExist(err) {
				t.Fatal(err)
			}
			if err == nil {
				if err := os.WriteFile(filepath.Join(directory, "scripts", "release"), publisher, 0o700); err != nil {
					t.Fatal(err)
				}
			}
			if err := os.WriteFile(filepath.Join(directory, ".release-please-manifest.json"), []byte("{\".\":\"0.7.5\"}\n"), 0o600); err != nil {
				t.Fatal(err)
			}
			releaseGit(t, directory, "init", "-q", "--initial-branch=main")
			releaseGit(t, directory, "config", "user.name", "Release test")
			releaseGit(t, directory, "config", "user.email", "release@example.test")
			releaseGit(t, directory, "config", "commit.gpgsign", "false")
			releaseGit(t, directory, "add", ".")
			releaseGit(t, directory, "commit", "-qm", "release fixture")
			remote := filepath.Join(directory, ".git", "origin.git")
			releaseGit(t, directory, "init", "-q", "--bare", remote)
			releaseGit(t, directory, "remote", "add", "origin", remote)
			releaseGit(t, directory, "push", "-q", "origin", "HEAD:refs/heads/main")
			if scenario.name == "tag on another commit" {
				releaseGit(t, directory, "tag", "v0.7.5")
				releaseGit(t, directory, "commit", "--allow-empty", "-qm", "candidate")
				releaseGit(t, directory, "push", "-q", "origin", "HEAD:refs/heads/main")
			}
			if scenario.name == "unmerged commit" {
				releaseGit(t, directory, "commit", "--allow-empty", "-qm", "unmerged candidate")
			}
			if scenario.name != "missing tag" && scenario.name != "tag on another commit" {
				releaseGit(t, directory, "tag", "v0.7.5")
			}
			if scenario.name != "missing tag" {
				releaseGit(t, directory, "push", "-q", "origin", "refs/tags/v0.7.5")
			}
			if scenario.name == "dirty source" {
				if err := os.WriteFile(filepath.Join(directory, "uncommitted"), []byte("dirty\n"), 0o600); err != nil {
					t.Fatal(err)
				}
			}
			workspace := t.TempDir()
			services := filepath.Join(directory, ".git", "external-services")
			if err := os.Mkdir(services, 0o700); err != nil {
				t.Fatal(err)
			}
			releaseWrite(t, filepath.Join(services, "gh"), []byte("#!/usr/bin/env bash\necho 'Unexpected external request.' >&2\nexit 1\n"))
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			command := releaseCommand(ctx, directory, "bash", "scripts/release", "publish")
			command.Env = append(os.Environ(), "TMPDIR="+workspace, "PATH="+services+":"+os.Getenv("PATH"))
			var stdout, stderr strings.Builder
			command.Stdout, command.Stderr = &stdout, &stderr
			err = command.Run()
			if failure, ok := err.(*exec.ExitError); !ok || failure.ExitCode() != 1 {
				t.Fatalf("publisher exit: %v\n%s\n%s", err, stdout.String(), stderr.String())
			}
			var result struct {
				Status string `json:"status"`
			}
			if err := json.Unmarshal([]byte(stdout.String()), &result); err != nil || result.Status != "failed" || !strings.Contains(stderr.String(), scenario.want) {
				t.Fatalf("publisher result: %q\n%s; want %q", stdout.String(), stderr.String(), scenario.want)
			}
			if _, err := os.Stat(filepath.Join(directory, "dist")); !os.IsNotExist(err) {
				t.Fatalf("publisher prepared artifacts before validating source: %v", err)
			}
			if entries, err := os.ReadDir(workspace); err != nil || len(entries) != 0 {
				t.Fatalf("publisher prepared temporary artifacts before validating source: %v, %v", entries, err)
			}
		})
	}
}

func TestReleaseConfigurationDoesNotRequireActions(t *testing.T) {
	directory := t.TempDir()
	if err := os.Mkdir(filepath.Join(directory, "bin"), 0o700); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"bin/check-release-environment", "bin/check-release-debug-info", "bin/run-goreleaser", ".goreleaser.yml", "release-please-config.json", "terraform-registry-manifest.json", "release-policy.json"} {
		data, err := os.ReadFile(filepath.Join("..", name))
		if err != nil {
			t.Fatal(err)
		}
		releaseWrite(t, filepath.Join(directory, name), data)
	}
	releaseGit(t, directory, "init", "--quiet", "--initial-branch=main")
	releaseGit(t, directory, "add", ".")
	releaseGit(t, directory, "-c", "user.name=Release test", "-c", "user.email=release@example.com", "-c", "commit.gpgsign=false", "commit", "--quiet", "-m", "Release fixture")
	releaseGit(t, directory, "remote", "add", "origin", "https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper.git")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	command := releaseCommand(ctx, directory, "bash", "bin/check-release-environment")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("release configuration: %v\n%s", err, output)
	}
}
