// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_bookmark_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_bookmark"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestXBookmarksDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_bookmark.XBookmarksDataSourceModel)(nil)
	schema := x_bookmark.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
