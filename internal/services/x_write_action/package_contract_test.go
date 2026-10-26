// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package xwriteaction

import "testing"

func TestNewXWriteActionDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewXWriteActionDataSource().(*XWriteActionDataSource); !ok {
		t.Fatal("NewXWriteActionDataSource did not return *XWriteActionDataSource")
	}
}
