// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package draw_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/draw"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestDrawDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*draw.DrawDataSourceModel)(nil)
	schema := draw.DataSourceSchema(context.TODO())
	errs := testhelpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
