// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package importpath_test

import (
	"reflect"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/importpath"
)

func TestParseID(t *testing.T) {
	a, b, c, d := false, int64(0), float64(0), ""
	diags := importpath.ParseImportID("true/1/1.1/hi", "///", &a, &b, &c, &d)
	results := []any{a, b, c, d}

	if diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if !reflect.DeepEqual(results, []any{true, int64(1), 1.1, "hi"}) {
		t.Fatalf("unexpected value: %v", results)
	}
}

func TestParseIDRejectsInvalidSegments(t *testing.T) {
	t.Parallel()

	var boolean bool
	if diagnostics := importpath.ParseImportID("no", "boolean", &boolean); !diagnostics.HasError() {
		t.Fatal("invalid bool produced no diagnostic")
	}
	var integer int64
	if diagnostics := importpath.ParseImportID("not-an-int", "integer", &integer); !diagnostics.HasError() {
		t.Fatal("invalid int produced no diagnostic")
	}
	var floatingPoint float64
	if diagnostics := importpath.ParseImportID("not-a-float", "float", &floatingPoint); !diagnostics.HasError() {
		t.Fatal("invalid float produced no diagnostic")
	}
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
	assertPanics(t, func() {
		var unsupported uint64
		importpath.ParseImportID("1", "unsupported", &unsupported)
	})
}

func assertPanics(t *testing.T, action func()) {
	t.Helper()
	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("operation did not panic")
		}
	}()
	action()
}
