// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package x_broadcast

import "testing"

func TestNewXBroadcastDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewXBroadcastDataSource().(*XBroadcastDataSource); !ok {
		t.Fatal("NewXBroadcastDataSource did not return *XBroadcastDataSource")
	}
}
