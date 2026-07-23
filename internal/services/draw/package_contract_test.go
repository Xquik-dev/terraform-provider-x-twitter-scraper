// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package draw

import "testing"

func TestNewDrawDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewDrawDataSource().(*DrawDataSource); !ok {
		t.Fatal("NewDrawDataSource did not return *DrawDataSource")
	}
}
