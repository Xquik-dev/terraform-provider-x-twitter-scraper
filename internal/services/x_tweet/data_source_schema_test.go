// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_tweet"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXTweetDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_tweet.XTweetDataSourceModel)(nil)
	schema := x_tweet.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
