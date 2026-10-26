// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/coveragebranch"
)

func main() {
	if len(os.Args) != 4 {
		exitf("Run go run ./cmd/branch-coverage REPORT PROFILE MINIMUM.")
	}
	minimum, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		exitf("The minimum is not a number. %v", err)
	}

	report, err := os.ReadFile(os.Args[1])
	if err != nil {
		exitf("The report cannot be read. %v", err)
	}
	profile, err := os.Open(os.Args[2])
	if err != nil {
		exitf("The profile cannot be opened. %v", err)
	}
	defer profile.Close()

	covered, total, err := coveragebranch.MeasureExcluding(
		report,
		profile,
		"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/coveragebranch",
		"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers",
	)
	if err != nil {
		exitf("Branch coverage cannot be measured. %v", err)
	}
	if total == 0 {
		exitf("Branch report has no branches. Check the test build.")
	}

	percent := 100 * float64(covered) / float64(total)
	fmt.Printf(
		"Branch coverage is %d/%d (%.2f%%). %.2f%% is required.\n",
		covered,
		total,
		percent,
		minimum,
	)
	if percent+0.0000001 < minimum {
		os.Exit(1)
	}
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
