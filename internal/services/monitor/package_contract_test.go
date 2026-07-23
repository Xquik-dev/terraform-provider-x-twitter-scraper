// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package monitor

import "testing"

func TestConstructorsReturnMonitorTypes(t *testing.T) {
	t.Parallel()

	if _, ok := NewResource().(*MonitorResource); !ok {
		t.Fatal("NewResource did not return *MonitorResource")
	}
	if _, ok := NewMonitorDataSource().(*MonitorDataSource); !ok {
		t.Fatal("NewMonitorDataSource did not return *MonitorDataSource")
	}
}
