// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_platform_link_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/bot_platform_link"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestBotPlatformLinkModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*bot_platform_link.BotPlatformLinkModel)(nil)
	schema := bot_platform_link.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
