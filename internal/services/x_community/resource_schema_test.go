// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_community_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_community"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXCommunityModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_community.XCommunityModel)(nil)
	schema := x_community.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
