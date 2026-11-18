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
	source, err := apicoverage.GenerateEnums(os.Getenv(apicoverage.SpecVariable), http.DefaultClient, os.Args[1:])
	if err == nil {
		err = os.WriteFile("enums.go", source, 0o644)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, apicoverage.Sentence(err))
		os.Exit(1)
	}
}
