// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_profile_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_profile"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXProfileModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_profile.XProfileModel)(nil)
	schema := x_profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
