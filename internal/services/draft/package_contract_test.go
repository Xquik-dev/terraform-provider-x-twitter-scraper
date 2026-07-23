// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package draft

import "testing"

func TestConstructorsReturnDraftTypes(t *testing.T) {
	t.Parallel()

	if _, ok := NewResource().(*DraftResource); !ok {
		t.Fatal("NewResource did not return *DraftResource")
	}
	if _, ok := NewDraftDataSource().(*DraftDataSource); !ok {
		t.Fatal("NewDraftDataSource did not return *DraftDataSource")
	}
}
