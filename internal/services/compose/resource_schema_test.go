// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compose_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/compose"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestComposeModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compose.ComposeModel)(nil)
	schema := compose.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
