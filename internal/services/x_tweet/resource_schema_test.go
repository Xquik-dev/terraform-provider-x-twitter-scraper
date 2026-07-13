// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_tweet"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXTweetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_tweet.XTweetModel)(nil)
	schema := x_tweet.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
