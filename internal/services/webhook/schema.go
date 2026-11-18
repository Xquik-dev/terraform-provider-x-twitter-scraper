// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"context"
	"slices"

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

var _ resource.ResourceWithConfigValidators = (*WebhookResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	eventTypes := slices.Concat(apienum.EventType, apienum.PostEventType)
	return schema.Schema{
		MarkdownDescription: "Webhook endpoint management and delivery",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"url": schema.StringAttribute{
				Description: "HTTPS URL",
				Required:    true,
			},
			"event_types": schema.ListAttribute{
				Description: "Array of event types to subscribe to.\n" + apienum.Describe(eventTypes),
				Required:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(eventTypes...),
					),
				},
				ElementType: types.StringType,
			},
			"is_active": schema.BoolAttribute{
				Optional: true,
			},
			"consecutive_failures": schema.Int64Attribute{
				Description: "Consecutive failed delivery attempts since the last success.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"delivery_status": schema.StringAttribute{
				Description: "Endpoint delivery state. needs_attention means delivery stopped after repeated failures.\nAvailable values are \"active\", \"paused\" and \"needs_attention\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"paused",
						"needs_attention",
					),
				},
			},
			"failure_hard_cap": schema.Int64Attribute{
				Description: "Consecutive delivery failures that pause the endpoint.",
				Computed:    true,
			},
			"secret": schema.StringAttribute{
				Description: "Plaintext HMAC signing secret returned only at creation.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func (r *WebhookResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WebhookResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
