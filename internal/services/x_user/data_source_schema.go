// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*XUserDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
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
			"name": schema.StringAttribute{
				Computed: true,
			},
			"parody_commentary_fan_label": schema.StringAttribute{
				Computed: true,
			},
			"possibly_sensitive": schema.BoolAttribute{
				Computed: true,
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
			"username": schema.StringAttribute{
				Computed: true,
			},
			"verified": schema.BoolAttribute{
				Computed: true,
			},
			"verified_type": schema.StringAttribute{
				Computed: true,
			},
			"pinned_tweet_ids": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"profile_bio": schema.MapAttribute{
				Description: "Structured profile bio with entity annotations",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"withheld_in_countries": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
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
				Description: "Identity verification metadata displayed by X.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XUserIdentityVerificationDataSourceModel](ctx),
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
		},
	}
}

func (d *XUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *XUserDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
