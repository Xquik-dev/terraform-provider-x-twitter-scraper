// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_community_join_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_community_join"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXCommunityJoinModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_community_join.XCommunityJoinModel)(nil)
	schema := x_community_join.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
