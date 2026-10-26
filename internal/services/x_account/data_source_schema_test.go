// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xaccount_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_account"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXAccountDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*xaccount.XAccountDataSourceModel)(nil)
	schema := xaccount.DataSourceSchema(context.TODO())
	errs := testhelpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
