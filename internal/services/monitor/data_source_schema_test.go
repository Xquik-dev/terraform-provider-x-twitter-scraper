// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monitor_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestMonitorDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*monitor.MonitorDataSourceModel)(nil)
	schema := monitor.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
