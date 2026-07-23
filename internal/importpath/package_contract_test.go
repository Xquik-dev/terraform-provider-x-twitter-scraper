// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package importpath

import "testing"

func TestParseImportIDDecodesPathSegments(t *testing.T) {
	t.Parallel()

	var value string
	diagnostics := ParseImportID("tweet%20search", "query", &value)
	if diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diagnostics)
	}
	if value != "tweet search" {
		t.Fatalf("unexpected decoded value: %q", value)
	}
}
