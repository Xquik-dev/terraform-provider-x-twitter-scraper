// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package activity

import "testing"

func TestNewActivityDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewActivityDataSource().(*ActivityDataSource); !ok {
		t.Fatal("NewActivityDataSource did not return *ActivityDataSource")
	}
}
