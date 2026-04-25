// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package support_ticket_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/support_ticket"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/test_helpers"
)

func TestSupportTicketDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*support_ticket.SupportTicketDataSourceModel)(nil)
	schema := support_ticket.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
