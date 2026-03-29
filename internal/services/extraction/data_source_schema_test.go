// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package extraction_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/extraction"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestExtractionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*extraction.ExtractionDataSourceModel)(nil)
	schema := extraction.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
