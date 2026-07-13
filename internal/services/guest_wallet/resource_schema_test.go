// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guest_wallet_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/guest_wallet"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestGuestWalletModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*guest_wallet.GuestWalletModel)(nil)
	schema := guest_wallet.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
