// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package importpath

import "testing"

func TestParseIDDecodesTheID(t *testing.T) {
	t.Parallel()

	id, diags := ParseImportID("tweet%20search")
	if diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if id != "tweet search" {
		t.Fatalf("unexpected ID: %q", id)
	}
}

func TestParseIDRejectsInvalidIDs(t *testing.T) {
	t.Parallel()

	if _, diagnostics := ParseImportID("%"); !diagnostics.HasError() {
		t.Fatal("invalid URL encoding produced no diagnostic")
	}
	if _, diagnostics := ParseImportID("one/two"); !diagnostics.HasError() {
		t.Fatal("an ID with 2 segments produced no diagnostic")
	}
}
