// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package event

import "testing"

func TestNewEventDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewEventDataSource().(*EventDataSource); !ok {
		t.Fatal("NewEventDataSource did not return *EventDataSource")
	}
}
