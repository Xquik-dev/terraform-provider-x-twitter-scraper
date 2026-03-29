// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscribe_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/subscribe"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestSubscribeModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*subscribe.SubscribeModel)(nil)
	schema := subscribe.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
