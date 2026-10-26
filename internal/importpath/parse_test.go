// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package importpath_test

import (
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/importpath"
)

func TestParseID(t *testing.T) {
	a, b, c, d := false, int64(0), float64(0), ""
	diags := importpath.ParseImportID("true/1/1.1/hi", "///", importpath.Bool(&a), importpath.Int64(&b), importpath.Float64(&c), importpath.String(&d))
	if diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if !a || b != 1 || c != 1.1 || d != "hi" {
		t.Fatalf("unexpected values: %t %d %g %q", a, b, c, d)
	}
}

func TestParseIDRejectsInvalidSegments(t *testing.T) {
	t.Parallel()

	var boolean bool
	if diagnostics := importpath.ParseImportID("no", "boolean", importpath.Bool(&boolean)); !diagnostics.HasError() {
		t.Fatal("invalid bool produced no diagnostic")
	}
	var integer int64
	if diagnostics := importpath.ParseImportID("not-an-int", "integer", importpath.Int64(&integer)); !diagnostics.HasError() {
		t.Fatal("invalid int produced no diagnostic")
	}
	var floatingPoint float64
	if diagnostics := importpath.ParseImportID("not-a-float", "float", importpath.Float64(&floatingPoint)); !diagnostics.HasError() {
		t.Fatal("invalid float produced no diagnostic")
	}
	var text string
	if diagnostics := importpath.ParseImportID("%", "text", importpath.String(&text)); !diagnostics.HasError() {
		t.Fatal("invalid URL encoding produced no diagnostic")
	}
	if diagnostics := importpath.ParseImportID("one", "one/two", importpath.String(&text), importpath.String(&text)); !diagnostics.HasError() {
		t.Fatal("wrong segment count produced no diagnostic")
	}
}

func TestParseIDRejectsInvalidCallers(t *testing.T) {
	t.Parallel()

	assertPanics(t, func() {
		var text string
		importpath.ParseImportID("one", "one/two", importpath.String(&text))
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
