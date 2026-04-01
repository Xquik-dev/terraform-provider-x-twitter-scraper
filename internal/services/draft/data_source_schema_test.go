// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package draft_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/draft"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestDraftDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*draft.DraftDataSourceModel)(nil)
	schema := draft.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
