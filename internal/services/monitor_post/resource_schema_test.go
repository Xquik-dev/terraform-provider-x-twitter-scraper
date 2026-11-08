// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitorpost_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor_post"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestMonitorPostModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*monitorpost.MonitorPostModel)(nil)
	schema := monitorpost.ResourceSchema(context.TODO())
	errs := testhelpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
