// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package style

import "testing"

func TestConstructorsReturnStyleTypes(t *testing.T) {
	t.Parallel()

	if _, ok := NewResource().(*StyleResource); !ok {
		t.Fatal("NewResource did not return *StyleResource")
	}
	if _, ok := NewStyleDataSource().(*StyleDataSource); !ok {
		t.Fatal("NewStyleDataSource did not return *StyleDataSource")
	}
}
