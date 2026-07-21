// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account": schema.StringAttribute{
				Description:   "X account (@username or account ID)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"attachment_url": schema.StringAttribute{
				Optional:      true,
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
			"charged": schema.BoolAttribute{
				Computed: true,
			},
			"charged_credits": schema.StringAttribute{
				Description: "Credits charged for this tweet. Text-only tweets and replies cost 30 credits; attached media adds 2 credits per started MB.",
				Computed:    true,
			},
			"success": schema.BoolAttribute{
				Computed: true,
			},
			"tweet_id": schema.StringAttribute{
				Computed: true,
			},
			"write_action_id": schema.StringAttribute{
				Computed: true,
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
					"author": schema.StringAttribute{
						Description: "Tweet author profile. The lookup route always includes follower count and verification state. Other profile fields appear when available.",
						Computed:    true,
						CustomType:  jsontypes.NormalizedType{},
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
									"can_edit": schema.BoolAttribute{
										Description: "Whether the disclosure can be edited on X.",
										Computed:    true,
									},
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
									"automated_by": schema.StringAttribute{
										Computed: true,
									},
									"can_dm": schema.BoolAttribute{
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
									"is_automated": schema.BoolAttribute{
										Computed: true,
									},
									"is_blue_verified": schema.BoolAttribute{
										Description: "Whether X shows a blue verification badge",
										Computed:    true,
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
									"profile_picture": schema.StringAttribute{
										Computed: true,
									},
									"protected": schema.BoolAttribute{
										Description: "Whether the profile protects its posts",
										Computed:    true,
									},
									"statuses_count": schema.Int64Attribute{
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
									"viewer_followed_by": schema.BoolAttribute{
										Description: "Whether this profile follows the authenticated viewer",
										Computed:    true,
									},
									"viewer_following": schema.BoolAttribute{
										Description: "Whether the authenticated viewer follows this profile",
										Computed:    true,
									},
									"withheld_in_countries": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
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
											"can_edit": schema.BoolAttribute{
												Description: "Whether the disclosure can be edited on X.",
												Computed:    true,
											},
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
									},
								},
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
									"automated_by": schema.StringAttribute{
										Computed: true,
									},
									"can_dm": schema.BoolAttribute{
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
									"is_automated": schema.BoolAttribute{
										Computed: true,
									},
									"is_blue_verified": schema.BoolAttribute{
										Description: "Whether X shows a blue verification badge",
										Computed:    true,
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
									"profile_picture": schema.StringAttribute{
										Computed: true,
									},
									"protected": schema.BoolAttribute{
										Description: "Whether the profile protects its posts",
										Computed:    true,
									},
									"statuses_count": schema.Int64Attribute{
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
									"viewer_followed_by": schema.BoolAttribute{
										Description: "Whether this profile follows the authenticated viewer",
										Computed:    true,
									},
									"viewer_following": schema.BoolAttribute{
										Description: "Whether the authenticated viewer follows this profile",
										Computed:    true,
									},
									"withheld_in_countries": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
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
											"can_edit": schema.BoolAttribute{
												Description: "Whether the disclosure can be edited on X.",
												Computed:    true,
											},
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
									},
								},
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
				},
			},
			"author": schema.StringAttribute{
				Description: "Tweet author profile. The lookup route always includes follower count and verification state. Other profile fields appear when available.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
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
