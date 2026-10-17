// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package x_list

import "testing"

func TestNewXListDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewXListDataSource().(*XListDataSource); !ok {
		t.Fatal("NewXListDataSource did not return *XListDataSource")
	}
}
