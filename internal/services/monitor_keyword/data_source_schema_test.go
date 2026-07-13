// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monitor_keyword_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor_keyword"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestMonitorKeywordDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*monitor_keyword.MonitorKeywordDataSourceModel)(nil)
	schema := monitor_keyword.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
