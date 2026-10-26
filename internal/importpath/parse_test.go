// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package importpath_test

import (
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/importpath"
)

func TestParseID(t *testing.T) {
	first, second := "", ""
	diags := importpath.ParseImportID("tweet%20search/hi", "query/name", &first, &second)
	if diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if first != "tweet search" || second != "hi" {
		t.Fatalf("unexpected values: %q %q", first, second)
	}
}

func TestParseIDRejectsInvalidSegments(t *testing.T) {
	t.Parallel()

	var text string
	if diagnostics := importpath.ParseImportID("%", "text", &text); !diagnostics.HasError() {
		t.Fatal("invalid URL encoding produced no diagnostic")
	}
	if diagnostics := importpath.ParseImportID("one", "one/two", &text, &text); !diagnostics.HasError() {
		t.Fatal("wrong segment count produced no diagnostic")
	}
}

func TestParseIDRejectsInvalidCallers(t *testing.T) {
	t.Parallel()

	assertPanics(t, func() {
		var text string
		importpath.ParseImportID("one", "one/two", &text)
	})
}

func assertPanics(t *testing.T, action func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("operation did not panic")
		}
	}()
	action()
}
