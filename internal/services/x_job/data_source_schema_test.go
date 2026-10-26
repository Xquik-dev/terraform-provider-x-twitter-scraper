// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xjob_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_job"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXJobDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*xjob.XJobDataSourceModel)(nil)
	schema := xjob.DataSourceSchema(context.TODO())
	errs := testhelpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
