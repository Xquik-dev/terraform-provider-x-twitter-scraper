// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"context"

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
				Description: "Array of event types to subscribe to.",
				Required:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"tweet.new",
							"tweet.reply",
							"tweet.retweet",
							"tweet.quote",
							"tweet.media",
							"tweet.link",
							"tweet.poll",
							"tweet.mention",
							"tweet.hashtag",
							"tweet.longform",
							"profile.avatar.changed",
							"profile.banner.changed",
							"profile.name.changed",
							"profile.username.changed",
							"profile.bio.changed",
							"profile.location.changed",
							"profile.url.changed",
							"profile.verified.changed",
							"profile.protected.changed",
							"profile.pinned_tweet.changed",
							"profile.unavailable.changed",
						),
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
				Description: "Endpoint delivery state. needs_attention means delivery stopped after repeated failures.\nAvailable values: \"active\", \"paused\", \"needs_attention\".",
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
