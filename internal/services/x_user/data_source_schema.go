// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
		Description: "X profile data lookup",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
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
			"name": schema.StringAttribute{
				Computed: true,
			},
			"possibly_sensitive": schema.BoolAttribute{
				Computed: true,
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
			"username": schema.StringAttribute{
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
		},
	}
}

func (d *XUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *XUserDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
