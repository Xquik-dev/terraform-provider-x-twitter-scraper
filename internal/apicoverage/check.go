// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package apicoverage compares the operations in the public API spec with the
// requests that the provider sends.
package apicoverage

import (
	"cmp"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	defaultSpecSource = "https://xquik.com/openapi.json"
	SpecVariable      = "XQUIK_OPENAPI_SPEC"
	SDKModule         = "github.com/Xquik-dev/x-twitter-scraper-go"
	exclusionFile     = "excluded-operations.json"
	servicesRoot      = "internal/services"
	missingNextStep   = "Add a resource or data source for each operation above, or list it in excluded-operations.json with a reason."
)

func check(root, source string, enums map[string][]string, client *http.Client) ([]string, []string, error) {
	excluded, err := readExclusions(os.DirFS(root))
	if err != nil {
		return nil, nil, err
	}
	document, err := readSpec(source, client)
	if err != nil {
		return nil, nil, err
	}
	api, err := specOperations(document, source)
	if err != nil {
		return nil, nil, err
	}
	directory, err := sdkDir(root)
	if err != nil {
		return nil, nil, err
	}
	index, err := readSDK(os.DirFS(directory))
	if err != nil {
		return nil, nil, err
	}
	requests, problems, err := providerRequests(root, index)
	if err != nil || len(problems) > 0 {
		return problems, nil, err
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
	enumFailures, enumSummaries := compareEnums(document, enums)
	if failures := slices.Concat(missing, stale, enumFailures); len(failures) > 0 {
		return failures, nil, nil
	}
	covered := len(api) - len(excluded)
	summary := fmt.Sprintf("The provider covers %d of %d API operations, and %s lists the other %d.", covered, len(api), exclusionFile, len(excluded))
	return nil, append([]string{summary}, enumSummaries...), nil
}

func Sentence(err error) string {
	text := err.Error()
	first, size := utf8.DecodeRuneInString(text)
	return string(unicode.ToUpper(first)) + text[size:] + "."
}

func Run(source, root string, enums map[string][]string, client *http.Client, out, failures io.Writer) int {
	lines, summaries, err := check(root, cmp.Or(source, defaultSpecSource), enums, client)
	if err != nil {
		lines = append(lines, Sentence(err))
	}
	for _, line := range lines {
		fmt.Fprintln(failures, line)
	}
	if len(lines) > 0 {
		return 1
	}
	for _, line := range summaries {
		fmt.Fprintln(out, line)
	}
	return 0
}
