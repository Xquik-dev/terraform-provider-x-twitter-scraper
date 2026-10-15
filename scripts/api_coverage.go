//go:build ignore

// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apicoverage"
)

func main() {
	var module struct{ Dir string }
	output, err := exec.Command("go", "mod", "download", "-json", apicoverage.SDKModule).Output()
	if err != nil || json.Unmarshal(output, &module) != nil || module.Dir == "" {
		fmt.Fprintf(os.Stderr, "The coverage check could not find the source of %s. Run go mod download, then rerun.\n", apicoverage.SDKModule)
		os.Exit(1)
	}
	os.Exit(apicoverage.Run(os.Getenv(apicoverage.SpecVariable), os.DirFS("."), os.DirFS(module.Dir), http.DefaultClient, os.Stdout, os.Stderr))
}
