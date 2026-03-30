// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/event"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestEventDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*event.EventDataSourceModel)(nil)
	schema := event.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
