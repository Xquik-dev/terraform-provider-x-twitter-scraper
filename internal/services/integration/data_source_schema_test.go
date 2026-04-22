// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/integration"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestIntegrationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*integration.IntegrationDataSourceModel)(nil)
	schema := integration.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
