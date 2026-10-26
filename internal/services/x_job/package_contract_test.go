// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package xjob

import "testing"

func TestNewXJobDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewXJobDataSource().(*XJobDataSource); !ok {
		t.Fatal("NewXJobDataSource did not return *XJobDataSource")
	}
}
