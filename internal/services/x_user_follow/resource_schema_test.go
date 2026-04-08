// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user_follow_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_user_follow"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestXUserFollowModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*x_user_follow.XUserFollowModel)(nil)
	schema := x_user_follow.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
