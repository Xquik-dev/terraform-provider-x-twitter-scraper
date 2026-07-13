// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package draft_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/draft"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestDraftModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*draft.DraftModel)(nil)
	schema := draft.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
