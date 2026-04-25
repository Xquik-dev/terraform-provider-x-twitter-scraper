// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_user"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestXUserDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_user.XUserDataSourceModel)(nil)
	schema := x_user.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
