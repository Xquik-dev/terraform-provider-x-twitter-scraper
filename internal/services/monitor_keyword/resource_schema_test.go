// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monitor_keyword_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor_keyword"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestMonitorKeywordModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*monitor_keyword.MonitorKeywordModel)(nil)
	schema := monitor_keyword.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
