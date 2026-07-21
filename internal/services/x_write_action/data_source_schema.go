// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_write_action

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*XWriteActionDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "X write actions (tweets, likes, follows, DMs)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"action": schema.StringAttribute{
				Description: `Available values: "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet", "follow", "unfollow", "remove_follower", "send_dm", "upload_media", "update_profile", "update_avatar", "update_banner", "create_community", "delete_community", "join_community", "leave_community".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"create_tweet",
						"delete_tweet",
						"like",
						"unlike",
						"retweet",
						"unretweet",
						"follow",
						"unfollow",
						"remove_follower",
						"send_dm",
						"upload_media",
						"update_profile",
						"update_avatar",
						"update_banner",
						"create_community",
						"delete_community",
						"join_community",
						"leave_community",
					),
				},
			},
			"charged": schema.BoolAttribute{
				Computed: true,
			},
			"charged_credits": schema.StringAttribute{
				Computed: true,
			},
			"community_id": schema.StringAttribute{
				Description: "Compatibility field for a confirmed community ID.",
				Computed:    true,
			},
			"community_name": schema.StringAttribute{
				Description: "Confirmed community name when available.",
				Computed:    true,
			},
			"completed_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"confirmation_attempts": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"confirmation_checked_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"confirmed_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"error": schema.StringAttribute{
				Computed: true,
			},
			"expires_at": schema.StringAttribute{
				Description: "Deadline for resolving a non-terminal write. This is not the Idempotency-Key retention deadline.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"idempotent": schema.BoolAttribute{
				Computed: true,
			},
			"media_id": schema.StringAttribute{
				Description: "Compatibility field for a confirmed media upload ID.",
				Computed:    true,
			},
			"media_url": schema.StringAttribute{
				Description: "Public media URL when the upload creates one.",
				Computed:    true,
			},
			"message": schema.StringAttribute{
				Computed: true,
			},
			"message_id": schema.StringAttribute{
				Description: "Compatibility field for a confirmed direct message ID.",
				Computed:    true,
			},
			"object": schema.StringAttribute{
				Description: `Available values: "x_write_action".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("x_write_action"),
				},
			},
			"poll_after_ms": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"request_hash": schema.StringAttribute{
				Computed: true,
			},
			"request_id": schema.StringAttribute{
				Computed: true,
			},
			"result_id": schema.StringAttribute{
				Description: "Compatibility result ID for other write actions.",
				Computed:    true,
			},
			"retryable": schema.BoolAttribute{
				Description: "True only when a new attempt can reasonably succeed.",
				Computed:    true,
			},
			"safe_to_retry": schema.BoolAttribute{
				Description: "True only when no write was dispatched and a new idempotency key may be used.",
				Computed:    true,
			},
			"send_dispatched": schema.BoolAttribute{
				Computed: true,
			},
			"send_dispatched_at": schema.StringAttribute{
				Description: "Dispatch timestamp when the write reached execution.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: `Available values: "accepted", "dispatching", "pending_confirmation", "success", "failed", "expired".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"accepted",
						"dispatching",
						"pending_confirmation",
						"success",
						"failed",
						"expired",
					),
				},
			},
			"status_url": schema.StringAttribute{
				Computed: true,
			},
			"success": schema.BoolAttribute{
				Computed: true,
			},
			"target_id": schema.StringAttribute{
				Computed: true,
			},
			"terminal": schema.BoolAttribute{
				Computed: true,
			},
			"tweet_id": schema.StringAttribute{
				Description: "Compatibility field for a confirmed tweet result ID.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"write_action_id": schema.StringAttribute{
				Computed: true,
			},
			"details": schema.MapAttribute{
				Description: "Structured recovery context for a failed write.",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"media": schema.MapAttribute{
				Description: "Media count, kind, size, and billing details when used.",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"account": schema.SingleNestedAttribute{
				Description: "Connected account selected for the write.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XWriteActionAccountDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"username": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"billing": schema.SingleNestedAttribute{
				Description: "plannedCredits is the approved maximum. chargedCredits comes from the settled credit ledger. Pending or failed writes are not charged.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XWriteActionBillingDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"charged": schema.BoolAttribute{
						Computed: true,
					},
					"charged_credits": schema.StringAttribute{
						Computed: true,
					},
					"planned_credits": schema.StringAttribute{
						Computed: true,
					},
					"status": schema.StringAttribute{
						Description: `Available values: "not_charged", "pending", "charged", "charge_failed", "refunded".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"not_charged",
								"pending",
								"charged",
								"charge_failed",
								"refunded",
							),
						},
					},
				},
			},
			"next_action": schema.SingleNestedAttribute{
				Description: "Exact follow-up an API client or agent should perform.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XWriteActionNextActionDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: `Available values: "poll", "retry", "verify_result", "fix_request".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"poll",
								"retry",
								"verify_result",
								"fix_request",
							),
						},
					},
					"after_ms": schema.Int64Attribute{
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"requires_new_idempotency_key": schema.BoolAttribute{
						Computed: true,
					},
					"url": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"request": schema.SingleNestedAttribute{
				Description: "Stable fingerprint and sanitized payload for replay checks.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XWriteActionRequestDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"hash": schema.StringAttribute{
						Description: "Stable hash of account, action, target, and payload.",
						Computed:    true,
					},
					"payload": schema.MapAttribute{
						Description: "Exact sanitized payload dispatched for this action.",
						Computed:    true,
						CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
						ElementType: jsontypes.NormalizedType{},
					},
				},
			},
			"result": schema.SingleNestedAttribute{
				Description: "Confirmed result produced by the write, when available.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XWriteActionResultDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"state": schema.StringAttribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Description: `Available values: "tweet", "direct_message", "media", "community", "state_change".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"tweet",
								"direct_message",
								"media",
								"community",
								"state_change",
							),
						},
					},
				},
			},
			"target": schema.SingleNestedAttribute{
				Description: "Existing X resource targeted by the write, when applicable.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XWriteActionTargetDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Description: `Available values: "tweet", "user", "community".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"tweet",
								"user",
								"community",
							),
						},
					},
				},
			},
		},
	}
}

func (d *XWriteActionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *XWriteActionDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
