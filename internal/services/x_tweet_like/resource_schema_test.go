// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet_like_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_tweet_like"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestXTweetLikeModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_tweet_like.XTweetLikeModel)(nil)
	schema := x_tweet_like.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
