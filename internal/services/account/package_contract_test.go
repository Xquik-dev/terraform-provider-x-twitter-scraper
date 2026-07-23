// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package account

import "testing"

func TestNewAccountDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewAccountDataSource().(*AccountDataSource); !ok {
		t.Fatal("NewAccountDataSource did not return *AccountDataSource")
	}
}
