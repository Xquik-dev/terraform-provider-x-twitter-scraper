// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monitor_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestMonitorModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*monitor.MonitorModel)(nil)
	schema := monitor.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
