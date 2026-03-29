// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package style_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/style"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestStyleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*style.StyleDataSourceModel)(nil)
	schema := style.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
