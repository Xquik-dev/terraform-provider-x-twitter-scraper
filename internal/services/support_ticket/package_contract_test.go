// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package supportticket

import "testing"

func TestConstructorsReturnSupportTicketTypes(t *testing.T) {
	t.Parallel()

	if _, ok := NewResource().(*SupportTicketResource); !ok {
		t.Fatal("NewResource did not return *SupportTicketResource")
	}
	if _, ok := NewSupportTicketDataSource().(*SupportTicketDataSource); !ok {
		t.Fatal("NewSupportTicketDataSource did not return *SupportTicketDataSource")
	}
}
