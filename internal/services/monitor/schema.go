// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apienum"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MonitorResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "X account monitoring with 1-second checks",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"username": schema.StringAttribute{
				Description:   "X username (without @)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"event_types": schema.ListAttribute{
				Description: "Array of event types to subscribe to.\n" + apienum.Describe(apienum.EventType),
				Required:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(apienum.EventType...),
					),
				},
				ElementType: types.StringType,
			},
			"is_active": schema.BoolAttribute{
				Optional: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
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
			"x_user_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *MonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MonitorResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
