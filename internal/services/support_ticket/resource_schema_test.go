// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package support_ticket_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/support_ticket"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestSupportTicketModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*support_ticket.SupportTicketModel)(nil)
	schema := support_ticket.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
