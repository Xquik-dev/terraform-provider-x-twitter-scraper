// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_tweet"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestXTweetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_tweet.XTweetModel)(nil)
	schema := x_tweet.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
