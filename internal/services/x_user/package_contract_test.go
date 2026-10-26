// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package xuser

import "testing"

func TestNewXUserDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewXUserDataSource().(*XUserDataSource); !ok {
		t.Fatal("NewXUserDataSource did not return *XUserDataSource")
	}
}
