// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package webhook

import "testing"

func TestNewResourceReturnsWebhookResource(t *testing.T) {
	t.Parallel()

	if _, ok := NewResource().(*WebhookResource); !ok {
		t.Fatal("NewResource did not return *WebhookResource")
	}
}
