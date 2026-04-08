// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_community_tweet_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_community_tweet"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestXCommunityTweetsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_community_tweet.XCommunityTweetsDataSourceModel)(nil)
	schema := x_community_tweet.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
