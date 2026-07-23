// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package schemata

import (
	"strings"
	"testing"
)

func TestDescriptionFormatsScopesAndMarkdown(t *testing.T) {
	t.Parallel()

	description := Description{
		Scopes:              []string{"tweet`search"},
		MarkdownDescription: "Search X posts.",
	}
	text := description.String()
	for _, expected := range []string{
		"Accepted Permissions",
		"`tweet\\`search`",
		"Search X posts.",
	} {
		if !strings.Contains(text, expected) {
			t.Fatalf("description does not contain %q: %q", expected, text)
		}
	}
}
