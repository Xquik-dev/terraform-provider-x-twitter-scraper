// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package style_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/style"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestStyleModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*style.StyleModel)(nil)
	schema := style.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
