//go:build ignore

// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apicoverage"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "The coverage check could not read the current folder. Run it from the repository root.")
		os.Exit(1)
	}
	os.Exit(apicoverage.Run(os.Getenv(apicoverage.SpecVariable), root, http.DefaultClient, os.Stdout, os.Stderr))
}
