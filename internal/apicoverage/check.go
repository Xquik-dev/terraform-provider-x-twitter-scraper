// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package apicoverage compares the operations in the public API spec with the
// requests that the provider sends.
package apicoverage

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)

const (
	defaultSpecSource = "https://xquik.com/openapi.json"
	SpecVariable      = "XQUIK_OPENAPI_SPEC"
	SDKModule         = "github.com/Xquik-dev/x-twitter-scraper-go"
	exclusionFile     = "excluded-operations.json"
	servicesRoot      = "internal/services"
	missingNextStep   = "Add a resource or data source for each operation above, or list it in excluded-operations.json with a reason."
)

func check(root, source string, client *http.Client) ([]string, string, error) {
	excluded, err := readExclusions(os.DirFS(root))
	if err != nil {
		return nil, "", err
	}
	api, err := specOperations(source, client)
	if err != nil {
		return nil, "", err
	}
	directory, err := sdkDir(root)
	if err != nil {
		return nil, "", err
	}
	index, err := readSDK(os.DirFS(directory))
	if err != nil {
		return nil, "", err
	}
	requests, problems, err := providerRequests(root, index)
	if err != nil || len(problems) > 0 {
		return problems, "", err
	}
	implemented := operations{}
	for _, found := range requests {
		target, _, _ := strings.Cut(found.path, "?")
		full := joinPath(index.basePath, target)
		implemented[operationKey(found.method, full)] = found.method + " " + full
	}
	var missing, stale []string
	for key, label := range api {
		if implemented[key] == "" && excluded[key] == "" {
			missing = append(missing, label+" is in the API but not in the provider.")
		}
	}
	for key, text := range excluded {
		switch {
		case api[key] == "":
			stale = append(stale, fmt.Sprintf("%s lists %s, which the API no longer has. Remove it.", exclusionFile, text))
		case implemented[key] != "":
			stale = append(stale, fmt.Sprintf("%s lists %s, which the provider already implements. Remove it.", exclusionFile, text))
		}
	}
	slices.Sort(missing)
	slices.Sort(stale)
	if len(missing) > 0 {
		missing = append(missing, missingNextStep)
	}
	if failures := append(missing, stale...); len(failures) > 0 {
		return failures, "", nil
	}
	covered := len(api) - len(excluded)
	return nil, fmt.Sprintf("The provider covers %d of %d API operations, and %s lists the other %d.", covered, len(api), exclusionFile, len(excluded)), nil
}

func Run(source, root string, client *http.Client, out, failures io.Writer) int {
	if source == "" {
		source = defaultSpecSource
	}
	lines, summary, err := check(root, source, client)
	if err != nil {
		lines = append(lines, err.Error())
	}
	for _, line := range lines {
		fmt.Fprintln(failures, line)
	}
	if len(lines) > 0 {
		return 1
	}
	fmt.Fprintln(out, summary)
	return 0
}
