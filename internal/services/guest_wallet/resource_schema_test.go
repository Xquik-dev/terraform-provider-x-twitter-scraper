// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package guestwallet_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/guest_wallet"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestGuestWalletModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*guestwallet.GuestWalletModel)(nil)
	schema := guestwallet.ResourceSchema(context.TODO())
	errs := testhelpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
