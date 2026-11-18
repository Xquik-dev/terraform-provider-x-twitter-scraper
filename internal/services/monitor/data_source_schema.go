// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apienum"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MonitorDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "X account monitoring with 1-second checks",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"is_active": schema.BoolAttribute{
				Computed: true,
			},
			"next_billing_at": schema.StringAttribute{
				Description: "Next hourly credit charge time for this account monitor.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"paused_at": schema.StringAttribute{
				Description: "When Xquik automatically paused this monitor.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"paused_reason": schema.StringAttribute{
				Description: "Why Xquik automatically paused this monitor.\nThe only available value is \"x_user_not_found\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("x_user_not_found"),
				},
			},
			"username": schema.StringAttribute{
				Computed: true,
			},
			"x_user_id": schema.StringAttribute{
				Computed: true,
			},
			"event_types": schema.ListAttribute{
				Description: "Array of event types to subscribe to.\n" + apienum.Describe(apienum.EventType),
				Computed:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(apienum.EventType...),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *MonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MonitorDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
