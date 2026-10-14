// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTerraformConsumerLoadsProviderSchema(t *testing.T) {
	workspace := t.TempDir()
	mirror := filepath.Join(workspace, "mirror")
	platform := runtime.GOOS + "_" + runtime.GOARCH
	packageDirectory := filepath.Join(mirror, "registry.terraform.io", "xquik-dev", "x-twitter-scraper", "0.0.1", platform)
	if err := os.MkdirAll(packageDirectory, 0o700); err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(packageDirectory, "terraform-provider-x-twitter-scraper_v0.0.1")
	build := exec.Command("go", "build", "-cover", "-coverpkg=./...", "-ldflags=-X main.version=0.0.1", "-o", binary, ".")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build provider: %v\n%s", err, output)
	}
	configuration := `terraform {
  required_providers {
    x-twitter-scraper = {
      source = "Xquik-dev/x-twitter-scraper"
      version = "0.0.1"
    }
  }
}`
	cliConfiguration := `provider_installation {
  filesystem_mirror {
    path = ` + strconv.Quote(mirror) + `
    include = ["registry.terraform.io/xquik-dev/x-twitter-scraper"]
  }
}`
	cliPath := filepath.Join(workspace, "terraform.rc")
	counters := os.Getenv("XQUIK_GO_COVERAGE_DIR")
	if counters == "" {
		counters = t.TempDir()
	}
	for path, contents := range map[string]string{
		filepath.Join(workspace, "main.tf"): configuration,
		cliPath:                             cliConfiguration,
	} {
		if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	run := func(args ...string) []byte {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		command := exec.CommandContext(ctx, "terraform", args...)
		command.Dir = workspace
		for _, variable := range os.Environ() {
			if !strings.HasPrefix(variable, "TF_") && !strings.HasPrefix(variable, "X_TWITTER_SCRAPER_") && !strings.HasPrefix(variable, "GOCOVERDIR=") {
				command.Env = append(command.Env, variable)
			}
		}
		command.Env = append(command.Env, "TF_CLI_CONFIG_FILE="+cliPath, "CHECKPOINT_DISABLE=1", "GOCOVERDIR="+counters)
		var stdout, stderr bytes.Buffer
		command.Stdout, command.Stderr = &stdout, &stderr
		if err := command.Run(); err != nil {
			t.Fatalf("terraform %v: %v\nstdout: %s\nstderr: %s", args, err, stdout.Bytes(), stderr.Bytes())
		}
		return stdout.Bytes()
	}
	run("init", "-backend=false", "-input=false", "-no-color")
	var report struct {
		Schemas map[string]struct {
			Provider struct {
				Block struct {
					Attributes map[string]struct {
						Sensitive bool `json:"sensitive"`
					} `json:"attributes"`
				} `json:"block"`
			} `json:"provider"`
			Resources   map[string]json.RawMessage `json:"resource_schemas"`
			DataSources map[string]json.RawMessage `json:"data_source_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal(run("providers", "schema", "-json"), &report); err != nil {
		t.Fatal(err)
	}
	schema, exists := report.Schemas["registry.terraform.io/xquik-dev/x-twitter-scraper"]
	if !exists || !schema.Provider.Block.Attributes["api_key"].Sensitive || !schema.Provider.Block.Attributes["bearer_token"].Sensitive {
		t.Fatal("Terraform did not receive the provider's sensitive credential schema")
	}
	if _, exists := schema.Resources["x-twitter-scraper_x_tweet"]; !exists {
		t.Fatal("Terraform did not receive the tweet resource schema")
	}
	if _, exists := schema.DataSources["x-twitter-scraper_x_user"]; !exists {
		t.Fatal("Terraform did not receive the user data source schema")
	}
}
