// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package x_user

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*XUserDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "X profile data lookup",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"username": schema.StringAttribute{
				Computed: true,
			},
			"account_based_in": schema.SingleNestedAttribute{
				Description: "Null when X shows no label.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserAccountBasedInDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"level": schema.StringAttribute{
						Description: `Available values are "country" and "region".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"country",
								"region",
							),
						},
					},
					"observed_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"value": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"account_based_in_accurate": schema.BoolAttribute{
				Description: "False when X notes accountBasedIn may be inaccurate.",
				Computed:    true,
			},
			"account_based_in_unavailable": schema.BoolAttribute{
				Description: "True when X withheld accountBasedIn from this read. Retry later.",
				Computed:    true,
			},
			"advertiser_account_service_levels": schema.ListAttribute{
				Description: "X Ads services the account uses, such as `analytics`, `media_studio` or `smb`. Empty for an account without ads.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"advertiser_account_type": schema.StringAttribute{
				Description: "The account's X Ads type. `none` for an account without ads, or `promotable_user` for one that may run them.",
				Computed:    true,
			},
			"affiliate_username": schema.StringAttribute{
				Description: "Handle of the organization X shows the account is affiliated with, from its About page.",
				Computed:    true,
			},
			"affiliates_highlighted_label": schema.SingleNestedAttribute{
				Description: "Organization affiliation label shown on an X profile.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserAffiliatesHighlightedLabelDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"badge_url": schema.StringAttribute{
						Computed: true,
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"long_description": schema.MapAttribute{
						Description: "Public text, ranges, references, and mention data.",
						Computed:    true,
						CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
						ElementType: jsontypes.NormalizedType{},
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
			"birthdate": schema.SingleNestedAttribute{
				Description: "The parts of the birthday the account made public.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserBirthdateDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"day": schema.Int64Attribute{
						Computed: true,
					},
					"month": schema.Int64Attribute{
						Computed: true,
					},
					"year": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"business_account_affiliates_count": schema.Int64Attribute{
				Computed: true,
			},
			"business_profile_state": schema.StringAttribute{
				Description: "Whether X treats the profile as a business profile. `none` or `enabled`.",
				Computed:    true,
			},
			"community_role": schema.StringAttribute{
				Description: "Community role when returned by community member reads",
				Computed:    true,
			},
			"connected_via": schema.StringAttribute{
				Description: "How the account first connected to X.",
				Computed:    true,
			},
			"connected_via_accurate": schema.BoolAttribute{
				Description: "False when X notes that country may be inaccurate.",
				Computed:    true,
			},
			"contact_info": schema.SingleNestedAttribute{
				Description: "Emails, phone numbers & links the account publishes on its own profile. Omitted when it publishes none.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserContactInfoDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"emails": schema.ListNestedAttribute{
						Description: "Emails in the bio, name or location, lowercased.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[XUserContactInfoEmailsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The part after the @ sign.",
									Computed:    true,
								},
								"source": schema.StringAttribute{
									Description: "Profile field that holds the email.\nAvailable values are \"bio\", \"name\" and \"location\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"bio",
											"name",
											"location",
										),
									},
								},
								"value": schema.StringAttribute{
									Description: "The email address.",
									Computed:    true,
								},
								"written_as": schema.StringAttribute{
									Description: "The account's own spelling, sent when it spelled the email out.",
									Computed:    true,
								},
							},
						},
					},
					"links": schema.ListNestedAttribute{
						Description: "The profile's website & the links in its bio.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[XUserContactInfoLinksDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"network": schema.StringAttribute{
									Description: "Site the link belongs to, such as `linkedin`, `github` or `telegram`. Any other site is `website`.",
									Computed:    true,
								},
								"source": schema.StringAttribute{
									Description: "`url` for the website field, `bio` for a bio link.\nAvailable values are \"url\" and \"bio\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"url",
											"bio",
										),
									},
								},
								"url": schema.StringAttribute{
									Description: "The link, as X expands it.",
									Computed:    true,
								},
							},
						},
					},
					"phones": schema.ListNestedAttribute{
						Description: "Phone numbers on the profile. Only a spelling that can be nothing else counts.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[XUserContactInfoPhonesDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"source": schema.StringAttribute{
									Description: "Profile field that holds the number.\nAvailable values are \"bio\", \"name\", \"location\" and \"url\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"bio",
											"name",
											"location",
											"url",
										),
									},
								},
								"value": schema.StringAttribute{
									Description: "The digits, with a leading + when the account wrote a country code.",
									Computed:    true,
								},
								"written_as": schema.StringAttribute{
									Description: "The account's own spelling of the number.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
			"cover_picture": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Description: "Account creation time.",
				Computed:    true,
			},
			"creator_subscriptions_count": schema.Int64Attribute{
				Computed: true,
			},
			"default_profile": schema.BoolAttribute{
				Description: "Whether the profile keeps X's default theme.",
				Computed:    true,
			},
			"default_profile_image": schema.BoolAttribute{
				Description: "Whether the profile keeps X's default avatar.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"email": schema.StringAttribute{
				Description: "First email the account publishes on its profile, lowercased. Omitted when it publishes none.",
				Computed:    true,
			},
			"fast_followers_count": schema.Int64Attribute{
				Description: "Followers X counts as fast followers.",
				Computed:    true,
			},
			"favourites_count": schema.Int64Attribute{
				Computed: true,
			},
			"followers": schema.Int64Attribute{
				Description: "Omitted when X sends no count, as for the account that first posted reused media.",
				Computed:    true,
			},
			"following": schema.Int64Attribute{
				Computed: true,
			},
			"geo_enabled": schema.BoolAttribute{
				Description: "Whether the profile may add a location to posts.",
				Computed:    true,
			},
			"grok_translated_bio": schema.MapAttribute{
				Description: "Public profile bio translation returned by X",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"has_custom_timelines": schema.BoolAttribute{
				Computed: true,
			},
			"has_extended_profile": schema.BoolAttribute{
				Description: "Whether the profile filled X's extended fields, such as a birth date. True or false.",
				Computed:    true,
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
				CustomType:  customfield.NewNestedObjectType[XUserHighlightsInfoDataSourceModel](ctx),
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
				Description: "Identity verification X shows on the profile & its About page.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserIdentityVerificationDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Computed: true,
					},
					"is_identity_verified": schema.BoolAttribute{
						Description: "Whether the owner verified their ID with X.",
						Computed:    true,
					},
					"verified_since_msec": schema.StringAttribute{
						Description: "When X verified the account, in milliseconds since 1970, as text.",
						Computed:    true,
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
			"is_translation_enabled": schema.BoolAttribute{
				Description: "Whether the account turned on X's translation of its posts. True or false.",
				Computed:    true,
			},
			"is_translator": schema.BoolAttribute{
				Computed: true,
			},
			"is_verified": schema.BoolAttribute{
				Description: "Whether X marks the profile as verified",
				Computed:    true,
			},
			"listed_count": schema.Int64Attribute{
				Description: "Public lists that include the profile.",
				Computed:    true,
			},
			"location": schema.StringAttribute{
				Description: "Account owner's public profile location text",
				Computed:    true,
			},
			"media_count": schema.Int64Attribute{
				Computed: true,
			},
			"normal_followers_count": schema.Int64Attribute{
				Description: "Followers minus fast followers.",
				Computed:    true,
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
			"professional": schema.MapAttribute{
				Description: "Professional metadata with category display settings",
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
				Description: "Profile picture at 48 px, as X sends it.",
				Computed:    true,
			},
			"profile_picture_original": schema.StringAttribute{
				Description: "The same profile picture at full size, as the account uploaded it. Download this one for the image.",
				Computed:    true,
			},
			"profile_sort_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"profile_translator_type": schema.StringAttribute{
				Computed: true,
			},
			"profile_bio": schema.MapAttribute{
				Description: "Structured profile bio with entity annotations",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"protected": schema.BoolAttribute{
				Description: "Whether the profile protects its posts",
				Computed:    true,
			},
			"redirects_to_username": schema.StringAttribute{
				Description: "Handle x.com sends this profile's visitors to, such as the new handle of a renamed account whose old handle now holds a placeholder. Missing when x.com shows the profile itself.",
				Computed:    true,
			},
			"retweeted_at": schema.StringAttribute{
				Description: "UTC repost time with includeRetweetTimestamp. Null if the newest profile page has no match or lookup failed. Otherwise omitted.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"statuses_count": schema.Int64Attribute{
				Computed: true,
			},
			"subscribers_count": schema.Int64Attribute{
				Description: "Accounts subscribed to the profile's paid Subscriptions. 0 when it has none.",
				Computed:    true,
			},
			"super_follow_eligible": schema.BoolAttribute{
				Computed: true,
			},
			"super_follows_user_profile_active": schema.BoolAttribute{
				Description: "Whether X marks the subscription profile as active.",
				Computed:    true,
			},
			"tip_jar": schema.SingleNestedAttribute{
				Description: "Public payment and creator-support handles shown on X.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserTipJarDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"bandcamp_handle": schema.StringAttribute{
						Computed: true,
					},
					"bitcoin_handle": schema.StringAttribute{
						Computed: true,
					},
					"cash_app_handle": schema.StringAttribute{
						Computed: true,
					},
					"ethereum_handle": schema.StringAttribute{
						Computed: true,
					},
					"gofundme_handle": schema.StringAttribute{
						Computed: true,
					},
					"is_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"patreon_handle": schema.StringAttribute{
						Computed: true,
					},
					"pay_pal_handle": schema.StringAttribute{
						Computed: true,
					},
					"venmo_handle": schema.StringAttribute{
						Computed: true,
					},
				},
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
			"username_changes": schema.SingleNestedAttribute{
				Description: "How often the account changed its username.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserUsernameChangesDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"count": schema.Int64Attribute{
						Computed: true,
					},
					"last_changed_at": schema.StringAttribute{
						Description: "Time of the latest change, omitted when there is none.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
				},
			},
			"verification_reason": schema.SingleNestedAttribute{
				Description: "X's text on why the profile is verified, with its links.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserVerificationReasonDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"entities": schema.ListNestedAttribute{
						Description: "Links inside the text, by character range.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[XUserVerificationReasonEntitiesDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from_index": schema.Int64Attribute{
									Computed: true,
								},
								"to_index": schema.Int64Attribute{
									Computed: true,
								},
								"url": schema.StringAttribute{
									Computed: true,
								},
								"url_type": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"text": schema.StringAttribute{
						Computed: true,
					},
				},
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
			"withheld_scope": schema.StringAttribute{
				Description: "Whether X withholds a post or user",
				Computed:    true,
			},
		},
	}
}

func (d *XUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *XUserDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
