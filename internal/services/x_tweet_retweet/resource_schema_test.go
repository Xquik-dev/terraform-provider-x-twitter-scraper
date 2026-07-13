// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet_retweet_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_tweet_retweet"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXTweetRetweetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_tweet_retweet.XTweetRetweetModel)(nil)
	schema := x_tweet_retweet.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
