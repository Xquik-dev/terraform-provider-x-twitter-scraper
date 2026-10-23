// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor_post

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

var _ resource.ResourceWithConfigValidators = (*MonitorPostResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Watches 1 post's replies, quotes, counts, edits & deletion",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"post": schema.StringAttribute{
				Description:   "Post link or ID. x.com, twitter.com, mobile & /i/web/status/ links work with any query string.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Your own label for the post monitor. Its events carry it as monitorName.",
				Optional:    true,
			},
			"event_types": schema.ListAttribute{
				Description: "Post event types to send. Without it, a monitor sends new replies, new quotes, like & retweet changes, edits & deletion.",
				Optional:    true,
				Computed:    true,
				Validators: []validator.List{
					listvalidator.SizeBetween(1, 10),
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"post.reply.new",
							"post.quote.new",
							"post.likes.changed",
							"post.retweets.changed",
							"post.replies.changed",
							"post.quotes.changed",
							"post.bookmarks.changed",
							"post.views.changed",
							"post.edited",
							"post.deleted",
						),
					),
				},
				ElementType: types.StringType,
			},
			"is_active": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"post_id": schema.StringAttribute{
				Description: "ID of the post you gave.",
				Computed:    true,
			},
			"post_url": schema.StringAttribute{
				Computed: true,
			},
			"author_id": schema.StringAttribute{
				Description: "Permanent X user ID of the post's author.",
				Computed:    true,
			},
			"author_username": schema.StringAttribute{
				Computed: true,
			},
			"author_name": schema.StringAttribute{
				Computed: true,
			},
			"text": schema.StringAttribute{
				Description: "Text of the newest version Xquik read.",
				Computed:    true,
			},
			"posted_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"next_billing_at": schema.StringAttribute{
				Description: "Next hourly credit charge time for this post monitor.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"paused_reason": schema.StringAttribute{
				Description: "Why Xquik paused this post monitor, such as insufficient_credits or x_post_deleted.",
				Computed:    true,
			},
			"paused_at": schema.StringAttribute{
				Description: "When Xquik paused this post monitor.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *MonitorPostResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MonitorPostResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
