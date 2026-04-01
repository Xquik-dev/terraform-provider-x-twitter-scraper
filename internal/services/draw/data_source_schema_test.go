// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package draw_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/draw"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestDrawDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*draw.DrawDataSourceModel)(nil)
	schema := draw.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
