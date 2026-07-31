// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*XTweetDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"author": schema.SingleNestedAttribute{
				Description: "Tweet author profile. The lookup route always includes follower count and verification state. Other profile fields appear when available.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XTweetAuthorDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetAuthorAffiliatesHighlightedLabelDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetAuthorHighlightsInfoDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetAuthorIdentityVerificationDataSourceModel](ctx),
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
			"tweet": schema.SingleNestedAttribute{
				Description: "Full tweet with text, engagement metrics, media, and metadata. A zero metric can mean X did not report the count.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XTweetTweetDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetArticleDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetAuthorDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetAuthorAffiliatesHighlightedLabelDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetAuthorHighlightsInfoDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetAuthorIdentityVerificationDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetCardDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetCommunityNoteDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetContentDisclosureDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"advertising": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[XTweetTweetContentDisclosureAdvertisingDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"is_paid_promotion": schema.BoolAttribute{
										Description: "True when X labels the tweet as paid promotion content.",
										Computed:    true,
									},
								},
							},
							"ai_generated": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[XTweetTweetContentDisclosureAIGeneratedDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetEditDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[XTweetTweetMediaDataSourceModel](ctx),
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
									CustomType:  customfield.NewMapType[customfield.NestedObjectList[XTweetTweetMediaFaceRectsDataSourceModel]](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[XTweetTweetMediaFocusRectsDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectMapType[XTweetTweetMediaSizesDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[XTweetTweetMediaVideoVariantsDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetNoteTweetDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[XTweetTweetNoteTweetRichtextTagsDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetPlaceDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetPreviousCountsDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetArticleDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetAuthorDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetAuthorAffiliatesHighlightedLabelDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetAuthorHighlightsInfoDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetAuthorIdentityVerificationDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetCardDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetCommunityNoteDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetContentDisclosureDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"advertising": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XTweetTweetQuotedTweetContentDisclosureAdvertisingDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"is_paid_promotion": schema.BoolAttribute{
												Description: "True when X labels the tweet as paid promotion content.",
												Computed:    true,
											},
										},
									},
									"ai_generated": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XTweetTweetQuotedTweetContentDisclosureAIGeneratedDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetEditDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[XTweetTweetQuotedTweetMediaDataSourceModel](ctx),
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
											CustomType:  customfield.NewMapType[customfield.NestedObjectList[XTweetTweetQuotedTweetMediaFaceRectsDataSourceModel]](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[XTweetTweetQuotedTweetMediaFocusRectsDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectMapType[XTweetTweetQuotedTweetMediaSizesDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[XTweetTweetQuotedTweetMediaVideoVariantsDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetNoteTweetDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectListType[XTweetTweetQuotedTweetNoteTweetRichtextTagsDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetPlaceDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetQuotedTweetPreviousCountsDataSourceModel](ctx),
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
								Description: "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
								Computed:    true,
								CustomType:  customfield.NormalizedDynamicType{},
							},
							"retweeted_tweet": schema.DynamicAttribute{
								Description: "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
								Computed:    true,
								CustomType:  customfield.NormalizedDynamicType{},
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
						CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetArticleDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetAuthorDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetAuthorAffiliatesHighlightedLabelDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetAuthorHighlightsInfoDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetAuthorIdentityVerificationDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetCardDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetCommunityNoteDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetContentDisclosureDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"advertising": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XTweetTweetRetweetedTweetContentDisclosureAdvertisingDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"is_paid_promotion": schema.BoolAttribute{
												Description: "True when X labels the tweet as paid promotion content.",
												Computed:    true,
											},
										},
									},
									"ai_generated": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XTweetTweetRetweetedTweetContentDisclosureAIGeneratedDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetEditDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[XTweetTweetRetweetedTweetMediaDataSourceModel](ctx),
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
											CustomType:  customfield.NewMapType[customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaFaceRectsDataSourceModel]](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[XTweetTweetRetweetedTweetMediaFocusRectsDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectMapType[XTweetTweetRetweetedTweetMediaSizesDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[XTweetTweetRetweetedTweetMediaVideoVariantsDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetNoteTweetDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectListType[XTweetTweetRetweetedTweetNoteTweetRichtextTagsDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetPlaceDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectType[XTweetTweetRetweetedTweetPreviousCountsDataSourceModel](ctx),
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
								Description: "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
								Computed:    true,
								CustomType:  customfield.NormalizedDynamicType{},
							},
							"retweeted_tweet": schema.DynamicAttribute{
								Description: "Quoted or retweeted tweet context. Every object includes id, text, and engagement metrics. A zero metric can mean X did not report the count. Author, media, and conversation fields appear when available.",
								Computed:    true,
								CustomType:  customfield.NormalizedDynamicType{},
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

func (d *XTweetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *XTweetDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
