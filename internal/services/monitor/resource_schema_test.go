// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
	errs := testhelpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
