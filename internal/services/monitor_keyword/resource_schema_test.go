// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitorkeyword_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor_keyword"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestMonitorKeywordModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*monitorkeyword.MonitorKeywordModel)(nil)
	schema := monitorkeyword.ResourceSchema(context.TODO())
	errs := testhelpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
