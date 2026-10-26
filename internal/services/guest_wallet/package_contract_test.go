// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package guestwallet

import "testing"

func TestNewResourceReturnsGuestWalletResource(t *testing.T) {
	t.Parallel()

	if _, ok := NewResource().(*GuestWalletResource); !ok {
		t.Fatal("NewResource did not return *GuestWalletResource")
	}
}
