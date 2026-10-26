// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package supportticket_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/support_ticket"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestSupportTicketModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*supportticket.SupportTicketModel)(nil)
	schema := supportticket.ResourceSchema(context.TODO())
	errs := testhelpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
