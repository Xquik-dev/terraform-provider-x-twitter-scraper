// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package compose

import "testing"

func TestNewResourceReturnsComposeResource(t *testing.T) {
	t.Parallel()

	if _, ok := NewResource().(*ComposeResource); !ok {
		t.Fatal("NewResource did not return *ComposeResource")
	}
}
