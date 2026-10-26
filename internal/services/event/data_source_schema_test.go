// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
	errs := testhelpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
