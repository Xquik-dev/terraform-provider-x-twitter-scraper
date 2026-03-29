// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_media_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_media"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestXMediaModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_media.XMediaModel)(nil)
	schema := x_media.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
