// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_dm_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_dm"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestXDmModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_dm.XDmModel)(nil)
	schema := x_dm.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
