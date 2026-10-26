// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package xaccount

import "testing"

func TestNewXAccountDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewXAccountDataSource().(*XAccountDataSource); !ok {
		t.Fatal("NewXAccountDataSource did not return *XAccountDataSource")
	}
}
