// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*XTweetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account": schema.StringAttribute{
				Description:   "X account (@username or account ID)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"community_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"is_note_tweet": schema.BoolAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"reply_to_tweet_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"text": schema.StringAttribute{
				Description:   "Tweet text (optional when media is provided)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"media": schema.ListAttribute{
				Description:   "Array of public media URLs to attach. Supports up to 4 images or exactly 1 MP4 video up to 100 MB. Each URL must be publicly reachable. Attached media adds 2 credits per started MB across all files.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
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
			"author": schema.SingleNestedAttribute{
				Description: "Tweet author profile. The lookup route always includes follower count and verification state. Other profile fields appear when available.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XTweetAuthorModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"username": schema.StringAttribute{
						Computed: true,
					},
					"affiliates_highlighted_label": schema.SingleNestedAttribute{
						Description: "Organization affiliation label shown on an X profile.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetAuthorAffiliatesHighlightedLabelModel](ctx),
						Attributes: map[string]schema.Attribute{
							"badge_url": schema.StringAttribute{
								Computed: true,
							},
							"description": schema.StringAttribute{
								Computed: true,
							},
							"url": schema.StringAttribute{
								Computed: true,
							},
							"url_type": schema.StringAttribute{
								Computed: true,
							},
							"user_label_display_type": schema.StringAttribute{
								Computed: true,
							},
							"user_label_type": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"automated_by": schema.StringAttribute{
						Computed: true,
					},
					"business_account_affiliates_count": schema.Int64Attribute{
						Computed: true,
					},
					"community_role": schema.StringAttribute{
						Description: "Community role when returned by community member reads",
						Computed:    true,
					},
					"cover_picture": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"creator_subscriptions_count": schema.Int64Attribute{
						Computed: true,
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"favourites_count": schema.Int64Attribute{
						Computed: true,
					},
					"followers": schema.Int64Attribute{
						Computed: true,
					},
					"following": schema.Int64Attribute{
						Computed: true,
					},
					"has_custom_timelines": schema.BoolAttribute{
						Computed: true,
					},
					"has_graduated_access": schema.BoolAttribute{
						Computed: true,
					},
					"has_hidden_subscriptions_on_profile": schema.BoolAttribute{
						Computed: true,
					},
					"highlights_info": schema.SingleNestedAttribute{
						Description: "Profile highlight availability and count metadata.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetAuthorHighlightsInfoModel](ctx),
						Attributes: map[string]schema.Attribute{
							"can_highlight_tweets": schema.BoolAttribute{
								Computed: true,
							},
							"highlighted_tweets": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"identity_verification": schema.SingleNestedAttribute{
						Description: "Identity verification metadata displayed by X.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetAuthorIdentityVerificationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								Computed: true,
							},
							"is_identity_verified": schema.BoolAttribute{
								Computed: true,
							},
							"verified_since_msec": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"is_automated": schema.BoolAttribute{
						Computed: true,
					},
					"is_blue_verified": schema.BoolAttribute{
						Description: "Whether X shows a blue verification badge",
						Computed:    true,
					},
					"is_profile_translatable": schema.BoolAttribute{
						Computed: true,
					},
					"is_translator": schema.BoolAttribute{
						Computed: true,
					},
					"is_verified": schema.BoolAttribute{
						Description: "Whether X marks the profile as verified",
						Computed:    true,
					},
					"location": schema.StringAttribute{
						Computed: true,
					},
					"media_count": schema.Int64Attribute{
						Computed: true,
					},
					"parody_commentary_fan_label": schema.StringAttribute{
						Computed: true,
					},
					"pinned_tweet_ids": schema.ListAttribute{
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"possibly_sensitive": schema.BoolAttribute{
						Computed: true,
					},
					"profile_bio": schema.MapAttribute{
						Description: "Structured profile bio with entity annotations",
						Computed:    true,
						CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
						ElementType: jsontypes.NormalizedType{},
					},
					"profile_banner_url": schema.StringAttribute{
						Description: "Original X profile banner field when available",
						Computed:    true,
					},
					"profile_description_language": schema.StringAttribute{
						Computed: true,
					},
					"profile_image_shape": schema.StringAttribute{
						Computed: true,
					},
					"profile_interstitial_type": schema.StringAttribute{
						Computed: true,
					},
					"profile_picture": schema.StringAttribute{
						Computed: true,
					},
					"profile_sort_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"profile_translator_type": schema.StringAttribute{
						Computed: true,
					},
					"protected": schema.BoolAttribute{
						Description: "Whether the profile protects its posts",
						Computed:    true,
					},
					"statuses_count": schema.Int64Attribute{
						Computed: true,
					},
					"super_follow_eligible": schema.BoolAttribute{
						Computed: true,
					},
					"unavailable": schema.BoolAttribute{
						Computed: true,
					},
					"unavailable_reason": schema.StringAttribute{
						Computed: true,
					},
					"url": schema.StringAttribute{
						Computed: true,
					},
					"verified": schema.BoolAttribute{
						Computed: true,
					},
					"verified_type": schema.StringAttribute{
						Computed: true,
					},
					"withheld_in_countries": schema.ListAttribute{
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"followers": schema.Int64Attribute{
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"verified": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
			"billing": schema.SingleNestedAttribute{
				Description: "plannedCredits is the approved maximum. chargedCredits comes from the settled credit ledger. Pending or failed writes are not charged.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XTweetBillingModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[XTweetNextActionModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[XTweetRequestModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[XTweetResultModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[XTweetTargetModel](ctx),
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
			"tweet": schema.SingleNestedAttribute{
				Description: "Full tweet with text, engagement metrics, media, and metadata. A zero metric can mean X did not report the count.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XTweetTweetModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"bookmark_count": schema.Int64Attribute{
						Computed: true,
					},
					"like_count": schema.Int64Attribute{
						Computed: true,
					},
					"quote_count": schema.Int64Attribute{
						Computed: true,
					},
					"reply_count": schema.Int64Attribute{
						Computed: true,
					},
					"retweet_count": schema.Int64Attribute{
						Computed: true,
					},
					"text": schema.StringAttribute{
						Computed: true,
					},
					"view_count": schema.Int64Attribute{
						Computed: true,
					},
					"article": schema.SingleNestedAttribute{
						Description: "Article metadata attached to a tweet.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetArticleModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"cover_media_url": schema.StringAttribute{
								Computed: true,
							},
							"preview_text": schema.StringAttribute{
								Computed: true,
							},
							"title": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"author": schema.SingleNestedAttribute{
						Description: "Tweet author profile. The lookup route always includes follower count and verification state. Other profile fields appear when available.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetAuthorModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"username": schema.StringAttribute{
								Computed: true,
							},
							"affiliates_highlighted_label": schema.SingleNestedAttribute{
								Description: "Organization affiliation label shown on an X profile.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetAuthorAffiliatesHighlightedLabelModel](ctx),
								Attributes: map[string]schema.Attribute{
									"badge_url": schema.StringAttribute{
										Computed: true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"url": schema.StringAttribute{
										Computed: true,
									},
									"url_type": schema.StringAttribute{
										Computed: true,
									},
									"user_label_display_type": schema.StringAttribute{
										Computed: true,
									},
									"user_label_type": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"automated_by": schema.StringAttribute{
								Computed: true,
							},
							"business_account_affiliates_count": schema.Int64Attribute{
								Computed: true,
							},
							"community_role": schema.StringAttribute{
								Description: "Community role when returned by community member reads",
								Computed:    true,
							},
							"cover_picture": schema.StringAttribute{
								Computed: true,
							},
							"created_at": schema.StringAttribute{
								Computed: true,
							},
							"creator_subscriptions_count": schema.Int64Attribute{
								Computed: true,
							},
							"description": schema.StringAttribute{
								Computed: true,
							},
							"favourites_count": schema.Int64Attribute{
								Computed: true,
							},
							"followers": schema.Int64Attribute{
								Computed: true,
							},
							"following": schema.Int64Attribute{
								Computed: true,
							},
							"has_custom_timelines": schema.BoolAttribute{
								Computed: true,
							},
							"has_graduated_access": schema.BoolAttribute{
								Computed: true,
							},
							"has_hidden_subscriptions_on_profile": schema.BoolAttribute{
								Computed: true,
							},
							"highlights_info": schema.SingleNestedAttribute{
								Description: "Profile highlight availability and count metadata.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetAuthorHighlightsInfoModel](ctx),
								Attributes: map[string]schema.Attribute{
									"can_highlight_tweets": schema.BoolAttribute{
										Computed: true,
									},
									"highlighted_tweets": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"identity_verification": schema.SingleNestedAttribute{
								Description: "Identity verification metadata displayed by X.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetAuthorIdentityVerificationModel](ctx),
								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										Computed: true,
									},
									"is_identity_verified": schema.BoolAttribute{
										Computed: true,
									},
									"verified_since_msec": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"is_automated": schema.BoolAttribute{
								Computed: true,
							},
							"is_blue_verified": schema.BoolAttribute{
								Description: "Whether X shows a blue verification badge",
								Computed:    true,
							},
							"is_profile_translatable": schema.BoolAttribute{
								Computed: true,
							},
							"is_translator": schema.BoolAttribute{
								Computed: true,
							},
							"is_verified": schema.BoolAttribute{
								Description: "Whether X marks the profile as verified",
								Computed:    true,
							},
							"location": schema.StringAttribute{
								Computed: true,
							},
							"media_count": schema.Int64Attribute{
								Computed: true,
							},
							"parody_commentary_fan_label": schema.StringAttribute{
								Computed: true,
							},
							"pinned_tweet_ids": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"possibly_sensitive": schema.BoolAttribute{
								Computed: true,
							},
							"profile_bio": schema.MapAttribute{
								Description: "Structured profile bio with entity annotations",
								Computed:    true,
								CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
								ElementType: jsontypes.NormalizedType{},
							},
							"profile_banner_url": schema.StringAttribute{
								Description: "Original X profile banner field when available",
								Computed:    true,
							},
							"profile_description_language": schema.StringAttribute{
								Computed: true,
							},
							"profile_image_shape": schema.StringAttribute{
								Computed: true,
							},
							"profile_interstitial_type": schema.StringAttribute{
								Computed: true,
							},
							"profile_picture": schema.StringAttribute{
								Computed: true,
							},
							"profile_sort_enabled": schema.BoolAttribute{
								Computed: true,
							},
							"profile_translator_type": schema.StringAttribute{
								Computed: true,
							},
							"protected": schema.BoolAttribute{
								Description: "Whether the profile protects its posts",
								Computed:    true,
							},
							"statuses_count": schema.Int64Attribute{
								Computed: true,
							},
							"super_follow_eligible": schema.BoolAttribute{
								Computed: true,
							},
							"unavailable": schema.BoolAttribute{
								Computed: true,
							},
							"unavailable_reason": schema.StringAttribute{
								Computed: true,
							},
							"url": schema.StringAttribute{
								Computed: true,
							},
							"verified": schema.BoolAttribute{
								Computed: true,
							},
							"verified_type": schema.StringAttribute{
								Computed: true,
							},
							"withheld_in_countries": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"followers": schema.Int64Attribute{
								Computed: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},
							"verified": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"card": schema.SingleNestedAttribute{
						Description: "Public card metadata attached to a tweet.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetCardModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"binding_values": schema.MapAttribute{
								Computed:    true,
								CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
								ElementType: jsontypes.NormalizedType{},
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"url": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"community_note": schema.SingleNestedAttribute{
						Description: "Community Note presentation metadata returned by X.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetCommunityNoteModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"destination_url": schema.StringAttribute{
								Computed: true,
							},
							"footer": schema.StringAttribute{
								Computed: true,
							},
							"short_title": schema.StringAttribute{
								Computed: true,
							},
							"subtitle": schema.StringAttribute{
								Computed: true,
							},
							"title": schema.StringAttribute{
								Computed: true,
							},
							"visual_style": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"content_disclosure": schema.SingleNestedAttribute{
						Description: "Content disclosure metadata shown by X when a tweet is labeled as paid partnership content or AI-generated media.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetContentDisclosureModel](ctx),
						Attributes: map[string]schema.Attribute{
							"advertising": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[XTweetTweetContentDisclosureAdvertisingModel](ctx),
								Attributes: map[string]schema.Attribute{
									"is_paid_promotion": schema.BoolAttribute{
										Description: "True when X labels the tweet as paid promotion content.",
										Computed:    true,
									},
								},
							},
							"ai_generated": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[XTweetTweetContentDisclosureAIGeneratedModel](ctx),
								Attributes: map[string]schema.Attribute{
									"detection_source": schema.StringAttribute{
										Description: "Source of the AI-generated media disclosure.",
										Computed:    true,
									},
									"has_ai_generated_media": schema.BoolAttribute{
										Description: "True when X labels the tweet as containing AI-generated media.",
										Computed:    true,
									},
								},
							},
						},
					},
					"conversation_id": schema.StringAttribute{
						Description: "ID of the root tweet in the conversation thread",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"display_text_range": schema.ListAttribute{
						Description: "Start and end offsets for rendered tweet text",
						Computed:    true,
						CustomType:  customfield.NewListType[types.Int64](ctx),
						ElementType: types.Int64Type,
					},
					"edit": schema.SingleNestedAttribute{
						Description: "Edit history metadata returned by X.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetEditModel](ctx),
						Attributes: map[string]schema.Attribute{
							"editable_until_msecs": schema.StringAttribute{
								Computed: true,
							},
							"edit_tweet_ids": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"entities": schema.MapAttribute{
						Description: "Parsed entities from the tweet text (URLs, mentions, hashtags, media)",
						Computed:    true,
						CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
						ElementType: jsontypes.NormalizedType{},
					},
					"in_reply_to_id": schema.StringAttribute{
						Description: "Tweet ID being replied to",
						Computed:    true,
					},
					"in_reply_to_user_id": schema.StringAttribute{
						Description: "User ID being replied to",
						Computed:    true,
					},
					"in_reply_to_username": schema.StringAttribute{
						Description: "Username being replied to",
						Computed:    true,
					},
					"is_limited_reply": schema.BoolAttribute{
						Description: "Whether replies are limited for this tweet",
						Computed:    true,
					},
					"is_note_tweet": schema.BoolAttribute{
						Description: "Whether this is a Note Tweet (long-form post, up to 25,000 characters)",
						Computed:    true,
					},
					"is_quote_status": schema.BoolAttribute{
						Description: "Whether this tweet quotes another tweet",
						Computed:    true,
					},
					"is_reply": schema.BoolAttribute{
						Description: "Whether this tweet is a reply to another tweet",
						Computed:    true,
					},
					"is_translatable": schema.BoolAttribute{
						Computed: true,
					},
					"lang": schema.StringAttribute{
						Description: "Tweet language code",
						Computed:    true,
					},
					"media": schema.ListNestedAttribute{
						Description: "Attached media items, omitted when the tweet has no media",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[XTweetTweetMediaModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"media_url": schema.StringAttribute{
									Description: "Media preview URL",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: `Available values: "photo", "video", "animated_gif".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"photo",
											"video",
											"animated_gif",
										),
									},
								},
								"url": schema.StringAttribute{
									Description: "X media link from the tweet",
									Computed:    true,
								},
								"id": schema.StringAttribute{
									Description: "X media entity ID.",
									Computed:    true,
								},
								"allow_download": schema.BoolAttribute{
									Description: "Whether X permits direct media download.",
									Computed:    true,
								},
								"alt_text": schema.StringAttribute{
									Description: "Accessibility text supplied for the media.",
									Computed:    true,
								},
								"aspect_ratio": schema.ListAttribute{
									Description: "Video aspect ratio as width and height.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.Int64](ctx),
									ElementType: types.Int64Type,
								},
								"availability_status": schema.StringAttribute{
									Description: "Media availability state reported by X.",
									Computed:    true,
								},
								"display_url": schema.StringAttribute{
									Description: "Display-friendly media URL reported by X.",
									Computed:    true,
								},
								"duration_millis": schema.Int64Attribute{
									Description: "Video duration in milliseconds.",
									Computed:    true,
								},
								"expanded_url": schema.StringAttribute{
									Description: "Expanded X media URL.",
									Computed:    true,
								},
								"face_rects": schema.MapAttribute{
									Description: "Face-aware crop rectangles grouped by media size.",
									Computed:    true,
									CustomType:  customfield.NewMapType[customfield.NestedObjectList[XTweetTweetMediaFaceRectsModel]](ctx),
									ElementType: types.ListType{
										ElemType: types.ObjectType{
											AttrTypes: map[string]attr.Type{"h": schema.Int64Attribute{
												Required: true,
											}.GetType(), "w": schema.Int64Attribute{
												Required: true,
											}.GetType(), "x": schema.Int64Attribute{
												Required: true,
											}.GetType(), "y": schema.Int64Attribute{
												Required: true,
											}.GetType()},
										},
									},
								},
								"focus_rects": schema.ListNestedAttribute{
									Description: "Suggested image crops reported by X.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[XTweetTweetMediaFocusRectsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"h": schema.Int64Attribute{
												Computed: true,
											},
											"w": schema.Int64Attribute{
												Computed: true,
											},
											"x": schema.Int64Attribute{
												Computed: true,
											},
											"y": schema.Int64Attribute{
												Computed: true,
											},
										},
									},
								},
								"height": schema.Int64Attribute{
									Description: "Original media height.",
									Computed:    true,
								},
								"indices": schema.ListAttribute{
									Description: "Media entity offsets in the tweet text.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.Int64](ctx),
									ElementType: types.Int64Type,
								},
								"media_key": schema.StringAttribute{
									Description: "Stable X media key.",
									Computed:    true,
								},
								"monetizable": schema.BoolAttribute{
									Description: "Whether X reports the media as monetizable.",
									Computed:    true,
								},
								"sizes": schema.MapNestedAttribute{
									Description: "Named media renditions and resize modes.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectMapType[XTweetTweetMediaSizesModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"h": schema.Int64Attribute{
												Computed: true,
											},
											"resize": schema.StringAttribute{
												Computed: true,
											},
											"w": schema.Int64Attribute{
												Computed: true,
											},
										},
									},
								},
								"video_variants": schema.ListNestedAttribute{
									Description: "Available video encodings, ordered as returned",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[XTweetTweetMediaVideoVariantsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"content_type": schema.StringAttribute{
												Computed: true,
											},
											"url": schema.StringAttribute{
												Computed: true,
											},
											"bitrate": schema.Int64Attribute{
												Computed: true,
											},
										},
									},
								},
								"width": schema.Int64Attribute{
									Description: "Original media width.",
									Computed:    true,
								},
							},
						},
					},
					"note_tweet": schema.SingleNestedAttribute{
						Description: "Complete Note Tweet content and rich-text metadata.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetNoteTweetModel](ctx),
						Attributes: map[string]schema.Attribute{
							"text": schema.StringAttribute{
								Computed: true,
							},
							"id": schema.StringAttribute{
								Computed: true,
							},
							"entities": schema.MapAttribute{
								Computed:    true,
								CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
								ElementType: jsontypes.NormalizedType{},
							},
							"is_expandable": schema.BoolAttribute{
								Computed: true,
							},
							"richtext_tags": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[XTweetTweetNoteTweetRichtextTagsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"from_index": schema.Int64Attribute{
											Computed: true,
										},
										"to_index": schema.Int64Attribute{
											Computed: true,
										},
										"types": schema.ListAttribute{
											Computed:    true,
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
									},
								},
							},
						},
					},
					"place": schema.SingleNestedAttribute{
						Description: "Public place metadata attached to a tweet.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetPlaceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"bounding_box": schema.MapAttribute{
								Computed:    true,
								CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
								ElementType: jsontypes.NormalizedType{},
							},
							"country": schema.StringAttribute{
								Computed: true,
							},
							"country_code": schema.StringAttribute{
								Computed: true,
							},
							"full_name": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"place_type": schema.StringAttribute{
								Computed: true,
							},
							"url": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"possibly_sensitive": schema.BoolAttribute{
						Computed: true,
					},
					"previous_counts": schema.SingleNestedAttribute{
						Description: "Engagement counts retained from a prior tweet edit.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetPreviousCountsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"bookmark_count": schema.Int64Attribute{
								Computed: true,
							},
							"like_count": schema.Int64Attribute{
								Computed: true,
							},
							"quote_count": schema.Int64Attribute{
								Computed: true,
							},
							"reply_count": schema.Int64Attribute{
								Computed: true,
							},
							"retweet_count": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
					"quoted_tweet": schema.SingleNestedAttribute{
						Description: "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"bookmark_count": schema.Int64Attribute{
								Computed: true,
							},
							"like_count": schema.Int64Attribute{
								Computed: true,
							},
							"quote_count": schema.Int64Attribute{
								Computed: true,
							},
							"reply_count": schema.Int64Attribute{
								Computed: true,
							},
							"retweet_count": schema.Int64Attribute{
								Computed: true,
							},
							"text": schema.StringAttribute{
								Computed: true,
							},
							"view_count": schema.Int64Attribute{
								Computed: true,
							},
							"article": schema.SingleNestedAttribute{
								Description: "Article metadata attached to a tweet.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetArticleModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"cover_media_url": schema.StringAttribute{
										Computed: true,
									},
									"preview_text": schema.StringAttribute{
										Computed: true,
									},
									"title": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"author": schema.SingleNestedAttribute{
								Description: "X user profile with bio, follower counts, and verification status.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetAuthorModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"username": schema.StringAttribute{
										Computed: true,
									},
									"affiliates_highlighted_label": schema.SingleNestedAttribute{
										Description: "Organization affiliation label shown on an X profile.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetAuthorAffiliatesHighlightedLabelModel](ctx),
										Attributes: map[string]schema.Attribute{
											"badge_url": schema.StringAttribute{
												Computed: true,
											},
											"description": schema.StringAttribute{
												Computed: true,
											},
											"url": schema.StringAttribute{
												Computed: true,
											},
											"url_type": schema.StringAttribute{
												Computed: true,
											},
											"user_label_display_type": schema.StringAttribute{
												Computed: true,
											},
											"user_label_type": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"automated_by": schema.StringAttribute{
										Computed: true,
									},
									"business_account_affiliates_count": schema.Int64Attribute{
										Computed: true,
									},
									"community_role": schema.StringAttribute{
										Description: "Community role when returned by community member reads",
										Computed:    true,
									},
									"cover_picture": schema.StringAttribute{
										Computed: true,
									},
									"created_at": schema.StringAttribute{
										Computed: true,
									},
									"creator_subscriptions_count": schema.Int64Attribute{
										Computed: true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"favourites_count": schema.Int64Attribute{
										Computed: true,
									},
									"followers": schema.Int64Attribute{
										Computed: true,
									},
									"following": schema.Int64Attribute{
										Computed: true,
									},
									"has_custom_timelines": schema.BoolAttribute{
										Computed: true,
									},
									"has_graduated_access": schema.BoolAttribute{
										Computed: true,
									},
									"has_hidden_subscriptions_on_profile": schema.BoolAttribute{
										Computed: true,
									},
									"highlights_info": schema.SingleNestedAttribute{
										Description: "Profile highlight availability and count metadata.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetAuthorHighlightsInfoModel](ctx),
										Attributes: map[string]schema.Attribute{
											"can_highlight_tweets": schema.BoolAttribute{
												Computed: true,
											},
											"highlighted_tweets": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"identity_verification": schema.SingleNestedAttribute{
										Description: "Identity verification metadata displayed by X.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetAuthorIdentityVerificationModel](ctx),
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"is_identity_verified": schema.BoolAttribute{
												Computed: true,
											},
											"verified_since_msec": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"is_automated": schema.BoolAttribute{
										Computed: true,
									},
									"is_blue_verified": schema.BoolAttribute{
										Description: "Whether X shows a blue verification badge",
										Computed:    true,
									},
									"is_profile_translatable": schema.BoolAttribute{
										Computed: true,
									},
									"is_translator": schema.BoolAttribute{
										Computed: true,
									},
									"is_verified": schema.BoolAttribute{
										Description: "Whether X marks the profile as verified",
										Computed:    true,
									},
									"location": schema.StringAttribute{
										Computed: true,
									},
									"media_count": schema.Int64Attribute{
										Computed: true,
									},
									"parody_commentary_fan_label": schema.StringAttribute{
										Computed: true,
									},
									"pinned_tweet_ids": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"possibly_sensitive": schema.BoolAttribute{
										Computed: true,
									},
									"profile_bio": schema.MapAttribute{
										Description: "Structured profile bio with entity annotations",
										Computed:    true,
										CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"profile_banner_url": schema.StringAttribute{
										Description: "Original X profile banner field when available",
										Computed:    true,
									},
									"profile_description_language": schema.StringAttribute{
										Computed: true,
									},
									"profile_image_shape": schema.StringAttribute{
										Computed: true,
									},
									"profile_interstitial_type": schema.StringAttribute{
										Computed: true,
									},
									"profile_picture": schema.StringAttribute{
										Computed: true,
									},
									"profile_sort_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"profile_translator_type": schema.StringAttribute{
										Computed: true,
									},
									"protected": schema.BoolAttribute{
										Description: "Whether the profile protects its posts",
										Computed:    true,
									},
									"statuses_count": schema.Int64Attribute{
										Computed: true,
									},
									"super_follow_eligible": schema.BoolAttribute{
										Computed: true,
									},
									"unavailable": schema.BoolAttribute{
										Computed: true,
									},
									"unavailable_reason": schema.StringAttribute{
										Computed: true,
									},
									"url": schema.StringAttribute{
										Computed: true,
									},
									"verified": schema.BoolAttribute{
										Computed: true,
									},
									"verified_type": schema.StringAttribute{
										Computed: true,
									},
									"withheld_in_countries": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
								},
							},
							"card": schema.SingleNestedAttribute{
								Description: "Public card metadata attached to a tweet.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetCardModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"binding_values": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"url": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"community_note": schema.SingleNestedAttribute{
								Description: "Community Note presentation metadata returned by X.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetCommunityNoteModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"destination_url": schema.StringAttribute{
										Computed: true,
									},
									"footer": schema.StringAttribute{
										Computed: true,
									},
									"short_title": schema.StringAttribute{
										Computed: true,
									},
									"subtitle": schema.StringAttribute{
										Computed: true,
									},
									"title": schema.StringAttribute{
										Computed: true,
									},
									"visual_style": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"content_disclosure": schema.SingleNestedAttribute{
								Description: "Content disclosure metadata shown by X when a tweet is labeled as paid partnership content or AI-generated media.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetContentDisclosureModel](ctx),
								Attributes: map[string]schema.Attribute{
									"advertising": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XTweetTweetQuotedTweetContentDisclosureAdvertisingModel](ctx),
										Attributes: map[string]schema.Attribute{
											"is_paid_promotion": schema.BoolAttribute{
												Description: "True when X labels the tweet as paid promotion content.",
												Computed:    true,
											},
										},
									},
									"ai_generated": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XTweetTweetQuotedTweetContentDisclosureAIGeneratedModel](ctx),
										Attributes: map[string]schema.Attribute{
											"detection_source": schema.StringAttribute{
												Description: "Source of the AI-generated media disclosure.",
												Computed:    true,
											},
											"has_ai_generated_media": schema.BoolAttribute{
												Description: "True when X labels the tweet as containing AI-generated media.",
												Computed:    true,
											},
										},
									},
								},
							},
							"conversation_id": schema.StringAttribute{
								Computed: true,
							},
							"created_at": schema.StringAttribute{
								Computed: true,
							},
							"display_text_range": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.Int64](ctx),
								ElementType: types.Int64Type,
							},
							"edit": schema.SingleNestedAttribute{
								Description: "Edit history metadata returned by X.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetEditModel](ctx),
								Attributes: map[string]schema.Attribute{
									"editable_until_msecs": schema.StringAttribute{
										Computed: true,
									},
									"edit_tweet_ids": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
								},
							},
							"entities": schema.MapAttribute{
								Computed:    true,
								CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
								ElementType: jsontypes.NormalizedType{},
							},
							"in_reply_to_id": schema.StringAttribute{
								Computed: true,
							},
							"in_reply_to_user_id": schema.StringAttribute{
								Computed: true,
							},
							"in_reply_to_username": schema.StringAttribute{
								Computed: true,
							},
							"is_limited_reply": schema.BoolAttribute{
								Computed: true,
							},
							"is_note_tweet": schema.BoolAttribute{
								Computed: true,
							},
							"is_quote_status": schema.BoolAttribute{
								Computed: true,
							},
							"is_reply": schema.BoolAttribute{
								Computed: true,
							},
							"is_translatable": schema.BoolAttribute{
								Computed: true,
							},
							"lang": schema.StringAttribute{
								Computed: true,
							},
							"media": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[XTweetTweetQuotedTweetMediaModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"media_url": schema.StringAttribute{
											Description: "Media preview URL",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: `Available values: "photo", "video", "animated_gif".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"photo",
													"video",
													"animated_gif",
												),
											},
										},
										"url": schema.StringAttribute{
											Description: "X media link from the tweet",
											Computed:    true,
										},
										"id": schema.StringAttribute{
											Description: "X media entity ID.",
											Computed:    true,
										},
										"allow_download": schema.BoolAttribute{
											Description: "Whether X permits direct media download.",
											Computed:    true,
										},
										"alt_text": schema.StringAttribute{
											Description: "Accessibility text supplied for the media.",
											Computed:    true,
										},
										"aspect_ratio": schema.ListAttribute{
											Description: "Video aspect ratio as width and height.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.Int64](ctx),
											ElementType: types.Int64Type,
										},
										"availability_status": schema.StringAttribute{
											Description: "Media availability state reported by X.",
											Computed:    true,
										},
										"display_url": schema.StringAttribute{
											Description: "Display-friendly media URL reported by X.",
											Computed:    true,
										},
										"duration_millis": schema.Int64Attribute{
											Description: "Video duration in milliseconds.",
											Computed:    true,
										},
										"expanded_url": schema.StringAttribute{
											Description: "Expanded X media URL.",
											Computed:    true,
										},
										"face_rects": schema.MapAttribute{
											Description: "Face-aware crop rectangles grouped by media size.",
											Computed:    true,
											CustomType:  customfield.NewMapType[customfield.NestedObjectList[XTweetTweetQuotedTweetMediaFaceRectsModel]](ctx),
											ElementType: types.ListType{
												ElemType: types.ObjectType{
													AttrTypes: map[string]attr.Type{"h": schema.Int64Attribute{
														Required: true,
													}.GetType(), "w": schema.Int64Attribute{
														Required: true,
													}.GetType(), "x": schema.Int64Attribute{
														Required: true,
													}.GetType(), "y": schema.Int64Attribute{
														Required: true,
													}.GetType()},
												},
											},
										},
										"focus_rects": schema.ListNestedAttribute{
											Description: "Suggested image crops reported by X.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[XTweetTweetQuotedTweetMediaFocusRectsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"h": schema.Int64Attribute{
														Computed: true,
													},
													"w": schema.Int64Attribute{
														Computed: true,
													},
													"x": schema.Int64Attribute{
														Computed: true,
													},
													"y": schema.Int64Attribute{
														Computed: true,
													},
												},
											},
										},
										"height": schema.Int64Attribute{
											Description: "Original media height.",
											Computed:    true,
										},
										"indices": schema.ListAttribute{
											Description: "Media entity offsets in the tweet text.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.Int64](ctx),
											ElementType: types.Int64Type,
										},
										"media_key": schema.StringAttribute{
											Description: "Stable X media key.",
											Computed:    true,
										},
										"monetizable": schema.BoolAttribute{
											Description: "Whether X reports the media as monetizable.",
											Computed:    true,
										},
										"sizes": schema.MapNestedAttribute{
											Description: "Named media renditions and resize modes.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectMapType[XTweetTweetQuotedTweetMediaSizesModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"h": schema.Int64Attribute{
														Computed: true,
													},
													"resize": schema.StringAttribute{
														Computed: true,
													},
													"w": schema.Int64Attribute{
														Computed: true,
													},
												},
											},
										},
										"video_variants": schema.ListNestedAttribute{
											Description: "Available video encodings, ordered as returned",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[XTweetTweetQuotedTweetMediaVideoVariantsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"content_type": schema.StringAttribute{
														Computed: true,
													},
													"url": schema.StringAttribute{
														Computed: true,
													},
													"bitrate": schema.Int64Attribute{
														Computed: true,
													},
												},
											},
										},
										"width": schema.Int64Attribute{
											Description: "Original media width.",
											Computed:    true,
										},
									},
								},
							},
							"note_tweet": schema.SingleNestedAttribute{
								Description: "Complete Note Tweet content and rich-text metadata.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetNoteTweetModel](ctx),
								Attributes: map[string]schema.Attribute{
									"text": schema.StringAttribute{
										Computed: true,
									},
									"id": schema.StringAttribute{
										Computed: true,
									},
									"entities": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"is_expandable": schema.BoolAttribute{
										Computed: true,
									},
									"richtext_tags": schema.ListNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectListType[XTweetTweetQuotedTweetNoteTweetRichtextTagsModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"from_index": schema.Int64Attribute{
													Computed: true,
												},
												"to_index": schema.Int64Attribute{
													Computed: true,
												},
												"types": schema.ListAttribute{
													Computed:    true,
													CustomType:  customfield.NewListType[types.String](ctx),
													ElementType: types.StringType,
												},
											},
										},
									},
								},
							},
							"place": schema.SingleNestedAttribute{
								Description: "Public place metadata attached to a tweet.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetPlaceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"bounding_box": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"country": schema.StringAttribute{
										Computed: true,
									},
									"country_code": schema.StringAttribute{
										Computed: true,
									},
									"full_name": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"place_type": schema.StringAttribute{
										Computed: true,
									},
									"url": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"possibly_sensitive": schema.BoolAttribute{
								Computed: true,
							},
							"previous_counts": schema.SingleNestedAttribute{
								Description: "Engagement counts retained from a prior tweet edit.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetPreviousCountsModel](ctx),
								Attributes: map[string]schema.Attribute{
									"bookmark_count": schema.Int64Attribute{
										Computed: true,
									},
									"like_count": schema.Int64Attribute{
										Computed: true,
									},
									"quote_count": schema.Int64Attribute{
										Computed: true,
									},
									"reply_count": schema.Int64Attribute{
										Computed: true,
									},
									"retweet_count": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
							"quoted_tweet": schema.DynamicAttribute{
								Description:   "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
								Computed:      true,
								CustomType:    customfield.NormalizedDynamicType{},
								PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
							},
							"retweeted_tweet": schema.DynamicAttribute{
								Description:   "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
								Computed:      true,
								CustomType:    customfield.NormalizedDynamicType{},
								PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
							},
							"source": schema.StringAttribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
							"url": schema.StringAttribute{
								Computed: true,
							},
							"view_state": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"retweeted_tweet": schema.SingleNestedAttribute{
						Description: "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"bookmark_count": schema.Int64Attribute{
								Computed: true,
							},
							"like_count": schema.Int64Attribute{
								Computed: true,
							},
							"quote_count": schema.Int64Attribute{
								Computed: true,
							},
							"reply_count": schema.Int64Attribute{
								Computed: true,
							},
							"retweet_count": schema.Int64Attribute{
								Computed: true,
							},
							"text": schema.StringAttribute{
								Computed: true,
							},
							"view_count": schema.Int64Attribute{
								Computed: true,
							},
							"article": schema.SingleNestedAttribute{
								Description: "Article metadata attached to a tweet.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetArticleModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"cover_media_url": schema.StringAttribute{
										Computed: true,
									},
									"preview_text": schema.StringAttribute{
										Computed: true,
									},
									"title": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"author": schema.SingleNestedAttribute{
								Description: "X user profile with bio, follower counts, and verification status.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetAuthorModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"username": schema.StringAttribute{
										Computed: true,
									},
									"affiliates_highlighted_label": schema.SingleNestedAttribute{
										Description: "Organization affiliation label shown on an X profile.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetAuthorAffiliatesHighlightedLabelModel](ctx),
										Attributes: map[string]schema.Attribute{
											"badge_url": schema.StringAttribute{
												Computed: true,
											},
											"description": schema.StringAttribute{
												Computed: true,
											},
											"url": schema.StringAttribute{
												Computed: true,
											},
											"url_type": schema.StringAttribute{
												Computed: true,
											},
											"user_label_display_type": schema.StringAttribute{
												Computed: true,
											},
											"user_label_type": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"automated_by": schema.StringAttribute{
										Computed: true,
									},
									"business_account_affiliates_count": schema.Int64Attribute{
										Computed: true,
									},
									"community_role": schema.StringAttribute{
										Description: "Community role when returned by community member reads",
										Computed:    true,
									},
									"cover_picture": schema.StringAttribute{
										Computed: true,
									},
									"created_at": schema.StringAttribute{
										Computed: true,
									},
									"creator_subscriptions_count": schema.Int64Attribute{
										Computed: true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"favourites_count": schema.Int64Attribute{
										Computed: true,
									},
									"followers": schema.Int64Attribute{
										Computed: true,
									},
									"following": schema.Int64Attribute{
										Computed: true,
									},
									"has_custom_timelines": schema.BoolAttribute{
										Computed: true,
									},
									"has_graduated_access": schema.BoolAttribute{
										Computed: true,
									},
									"has_hidden_subscriptions_on_profile": schema.BoolAttribute{
										Computed: true,
									},
									"highlights_info": schema.SingleNestedAttribute{
										Description: "Profile highlight availability and count metadata.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetAuthorHighlightsInfoModel](ctx),
										Attributes: map[string]schema.Attribute{
											"can_highlight_tweets": schema.BoolAttribute{
												Computed: true,
											},
											"highlighted_tweets": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"identity_verification": schema.SingleNestedAttribute{
										Description: "Identity verification metadata displayed by X.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetAuthorIdentityVerificationModel](ctx),
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"is_identity_verified": schema.BoolAttribute{
												Computed: true,
											},
											"verified_since_msec": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"is_automated": schema.BoolAttribute{
										Computed: true,
									},
									"is_blue_verified": schema.BoolAttribute{
										Description: "Whether X shows a blue verification badge",
										Computed:    true,
									},
									"is_profile_translatable": schema.BoolAttribute{
										Computed: true,
									},
									"is_translator": schema.BoolAttribute{
										Computed: true,
									},
									"is_verified": schema.BoolAttribute{
										Description: "Whether X marks the profile as verified",
										Computed:    true,
									},
									"location": schema.StringAttribute{
										Computed: true,
									},
									"media_count": schema.Int64Attribute{
										Computed: true,
									},
									"parody_commentary_fan_label": schema.StringAttribute{
										Computed: true,
									},
									"pinned_tweet_ids": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"possibly_sensitive": schema.BoolAttribute{
										Computed: true,
									},
									"profile_bio": schema.MapAttribute{
										Description: "Structured profile bio with entity annotations",
										Computed:    true,
										CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"profile_banner_url": schema.StringAttribute{
										Description: "Original X profile banner field when available",
										Computed:    true,
									},
									"profile_description_language": schema.StringAttribute{
										Computed: true,
									},
									"profile_image_shape": schema.StringAttribute{
										Computed: true,
									},
									"profile_interstitial_type": schema.StringAttribute{
										Computed: true,
									},
									"profile_picture": schema.StringAttribute{
										Computed: true,
									},
									"profile_sort_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"profile_translator_type": schema.StringAttribute{
										Computed: true,
									},
									"protected": schema.BoolAttribute{
										Description: "Whether the profile protects its posts",
										Computed:    true,
									},
									"statuses_count": schema.Int64Attribute{
										Computed: true,
									},
									"super_follow_eligible": schema.BoolAttribute{
										Computed: true,
									},
									"unavailable": schema.BoolAttribute{
										Computed: true,
									},
									"unavailable_reason": schema.StringAttribute{
										Computed: true,
									},
									"url": schema.StringAttribute{
										Computed: true,
									},
									"verified": schema.BoolAttribute{
										Computed: true,
									},
									"verified_type": schema.StringAttribute{
										Computed: true,
									},
									"withheld_in_countries": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
								},
							},
							"card": schema.SingleNestedAttribute{
								Description: "Public card metadata attached to a tweet.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetCardModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"binding_values": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"url": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"community_note": schema.SingleNestedAttribute{
								Description: "Community Note presentation metadata returned by X.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetCommunityNoteModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"destination_url": schema.StringAttribute{
										Computed: true,
									},
									"footer": schema.StringAttribute{
										Computed: true,
									},
									"short_title": schema.StringAttribute{
										Computed: true,
									},
									"subtitle": schema.StringAttribute{
										Computed: true,
									},
									"title": schema.StringAttribute{
										Computed: true,
									},
									"visual_style": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"content_disclosure": schema.SingleNestedAttribute{
								Description: "Content disclosure metadata shown by X when a tweet is labeled as paid partnership content or AI-generated media.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetContentDisclosureModel](ctx),
								Attributes: map[string]schema.Attribute{
									"advertising": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XTweetTweetRetweetedTweetContentDisclosureAdvertisingModel](ctx),
										Attributes: map[string]schema.Attribute{
											"is_paid_promotion": schema.BoolAttribute{
												Description: "True when X labels the tweet as paid promotion content.",
												Computed:    true,
											},
										},
									},
									"ai_generated": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XTweetTweetRetweetedTweetContentDisclosureAIGeneratedModel](ctx),
										Attributes: map[string]schema.Attribute{
											"detection_source": schema.StringAttribute{
												Description: "Source of the AI-generated media disclosure.",
												Computed:    true,
											},
											"has_ai_generated_media": schema.BoolAttribute{
												Description: "True when X labels the tweet as containing AI-generated media.",
												Computed:    true,
											},
										},
									},
								},
							},
							"conversation_id": schema.StringAttribute{
								Computed: true,
							},
							"created_at": schema.StringAttribute{
								Computed: true,
							},
							"display_text_range": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.Int64](ctx),
								ElementType: types.Int64Type,
							},
							"edit": schema.SingleNestedAttribute{
								Description: "Edit history metadata returned by X.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetEditModel](ctx),
								Attributes: map[string]schema.Attribute{
									"editable_until_msecs": schema.StringAttribute{
										Computed: true,
									},
									"edit_tweet_ids": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
								},
							},
							"entities": schema.MapAttribute{
								Computed:    true,
								CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
								ElementType: jsontypes.NormalizedType{},
							},
							"in_reply_to_id": schema.StringAttribute{
								Computed: true,
							},
							"in_reply_to_user_id": schema.StringAttribute{
								Computed: true,
							},
							"in_reply_to_username": schema.StringAttribute{
								Computed: true,
							},
							"is_limited_reply": schema.BoolAttribute{
								Computed: true,
							},
							"is_note_tweet": schema.BoolAttribute{
								Computed: true,
							},
							"is_quote_status": schema.BoolAttribute{
								Computed: true,
							},
							"is_reply": schema.BoolAttribute{
								Computed: true,
							},
							"is_translatable": schema.BoolAttribute{
								Computed: true,
							},
							"lang": schema.StringAttribute{
								Computed: true,
							},
							"media": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[XTweetTweetRetweetedTweetMediaModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"media_url": schema.StringAttribute{
											Description: "Media preview URL",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: `Available values: "photo", "video", "animated_gif".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"photo",
													"video",
													"animated_gif",
												),
											},
										},
										"url": schema.StringAttribute{
											Description: "X media link from the tweet",
											Computed:    true,
										},
										"id": schema.StringAttribute{
											Description: "X media entity ID.",
											Computed:    true,
										},
										"allow_download": schema.BoolAttribute{
											Description: "Whether X permits direct media download.",
											Computed:    true,
										},
										"alt_text": schema.StringAttribute{
											Description: "Accessibility text supplied for the media.",
											Computed:    true,
										},
										"aspect_ratio": schema.ListAttribute{
											Description: "Video aspect ratio as width and height.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.Int64](ctx),
											ElementType: types.Int64Type,
										},
										"availability_status": schema.StringAttribute{
											Description: "Media availability state reported by X.",
											Computed:    true,
										},
										"display_url": schema.StringAttribute{
											Description: "Display-friendly media URL reported by X.",
											Computed:    true,
										},
										"duration_millis": schema.Int64Attribute{
											Description: "Video duration in milliseconds.",
											Computed:    true,
										},
										"expanded_url": schema.StringAttribute{
											Description: "Expanded X media URL.",
											Computed:    true,
										},
										"face_rects": schema.MapAttribute{
											Description: "Face-aware crop rectangles grouped by media size.",
											Computed:    true,
											CustomType:  customfield.NewMapType[customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaFaceRectsModel]](ctx),
											ElementType: types.ListType{
												ElemType: types.ObjectType{
													AttrTypes: map[string]attr.Type{"h": schema.Int64Attribute{
														Required: true,
													}.GetType(), "w": schema.Int64Attribute{
														Required: true,
													}.GetType(), "x": schema.Int64Attribute{
														Required: true,
													}.GetType(), "y": schema.Int64Attribute{
														Required: true,
													}.GetType()},
												},
											},
										},
										"focus_rects": schema.ListNestedAttribute{
											Description: "Suggested image crops reported by X.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[XTweetTweetRetweetedTweetMediaFocusRectsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"h": schema.Int64Attribute{
														Computed: true,
													},
													"w": schema.Int64Attribute{
														Computed: true,
													},
													"x": schema.Int64Attribute{
														Computed: true,
													},
													"y": schema.Int64Attribute{
														Computed: true,
													},
												},
											},
										},
										"height": schema.Int64Attribute{
											Description: "Original media height.",
											Computed:    true,
										},
										"indices": schema.ListAttribute{
											Description: "Media entity offsets in the tweet text.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.Int64](ctx),
											ElementType: types.Int64Type,
										},
										"media_key": schema.StringAttribute{
											Description: "Stable X media key.",
											Computed:    true,
										},
										"monetizable": schema.BoolAttribute{
											Description: "Whether X reports the media as monetizable.",
											Computed:    true,
										},
										"sizes": schema.MapNestedAttribute{
											Description: "Named media renditions and resize modes.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectMapType[XTweetTweetRetweetedTweetMediaSizesModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"h": schema.Int64Attribute{
														Computed: true,
													},
													"resize": schema.StringAttribute{
														Computed: true,
													},
													"w": schema.Int64Attribute{
														Computed: true,
													},
												},
											},
										},
										"video_variants": schema.ListNestedAttribute{
											Description: "Available video encodings, ordered as returned",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[XTweetTweetRetweetedTweetMediaVideoVariantsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"content_type": schema.StringAttribute{
														Computed: true,
													},
													"url": schema.StringAttribute{
														Computed: true,
													},
													"bitrate": schema.Int64Attribute{
														Computed: true,
													},
												},
											},
										},
										"width": schema.Int64Attribute{
											Description: "Original media width.",
											Computed:    true,
										},
									},
								},
							},
							"note_tweet": schema.SingleNestedAttribute{
								Description: "Complete Note Tweet content and rich-text metadata.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetNoteTweetModel](ctx),
								Attributes: map[string]schema.Attribute{
									"text": schema.StringAttribute{
										Computed: true,
									},
									"id": schema.StringAttribute{
										Computed: true,
									},
									"entities": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"is_expandable": schema.BoolAttribute{
										Computed: true,
									},
									"richtext_tags": schema.ListNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectListType[XTweetTweetRetweetedTweetNoteTweetRichtextTagsModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"from_index": schema.Int64Attribute{
													Computed: true,
												},
												"to_index": schema.Int64Attribute{
													Computed: true,
												},
												"types": schema.ListAttribute{
													Computed:    true,
													CustomType:  customfield.NewListType[types.String](ctx),
													ElementType: types.StringType,
												},
											},
										},
									},
								},
							},
							"place": schema.SingleNestedAttribute{
								Description: "Public place metadata attached to a tweet.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetPlaceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"bounding_box": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"country": schema.StringAttribute{
										Computed: true,
									},
									"country_code": schema.StringAttribute{
										Computed: true,
									},
									"full_name": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"place_type": schema.StringAttribute{
										Computed: true,
									},
									"url": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"possibly_sensitive": schema.BoolAttribute{
								Computed: true,
							},
							"previous_counts": schema.SingleNestedAttribute{
								Description: "Engagement counts retained from a prior tweet edit.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetPreviousCountsModel](ctx),
								Attributes: map[string]schema.Attribute{
									"bookmark_count": schema.Int64Attribute{
										Computed: true,
									},
									"like_count": schema.Int64Attribute{
										Computed: true,
									},
									"quote_count": schema.Int64Attribute{
										Computed: true,
									},
									"reply_count": schema.Int64Attribute{
										Computed: true,
									},
									"retweet_count": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
							"quoted_tweet": schema.DynamicAttribute{
								Description:   "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
								Computed:      true,
								CustomType:    customfield.NormalizedDynamicType{},
								PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
							},
							"retweeted_tweet": schema.DynamicAttribute{
								Description:   "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
								Computed:      true,
								CustomType:    customfield.NormalizedDynamicType{},
								PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
							},
							"source": schema.StringAttribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
							"url": schema.StringAttribute{
								Computed: true,
							},
							"view_state": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"source": schema.StringAttribute{
						Description: "Client application used to post this tweet",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "Tweet result type",
						Computed:    true,
					},
					"url": schema.StringAttribute{
						Description: "Tweet permalink URL",
						Computed:    true,
					},
					"view_state": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *XTweetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *XTweetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
