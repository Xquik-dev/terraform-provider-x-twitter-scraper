// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtweet_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_tweet"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
)

func TestXTweetDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*xtweet.XTweetDataSourceModel)(nil)
	schema := xtweet.DataSourceSchema(context.TODO())
	errs := testhelpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
