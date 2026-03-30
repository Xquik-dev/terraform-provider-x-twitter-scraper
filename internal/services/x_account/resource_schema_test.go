// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_account_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_account"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXAccountModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_account.XAccountModel)(nil)
	schema := x_account.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
