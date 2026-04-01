// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/integration"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestIntegrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*integration.IntegrationModel)(nil)
	schema := integration.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
