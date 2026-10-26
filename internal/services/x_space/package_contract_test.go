// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package xspace

import "testing"

func TestNewXSpaceDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewXSpaceDataSource().(*XSpaceDataSource); !ok {
		t.Fatal("NewXSpaceDataSource did not return *XSpaceDataSource")
	}
}
