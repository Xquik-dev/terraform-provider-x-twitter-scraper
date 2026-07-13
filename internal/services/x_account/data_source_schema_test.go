// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_account_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_account"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXAccountDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_account.XAccountDataSourceModel)(nil)
	schema := x_account.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
