// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*XTweetDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"tweet_id": schema.StringAttribute{
				Required: true,
			},
			"author": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[XTweetAuthorDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"followers": schema.Int64Attribute{
						Computed: true,
					},
					"username": schema.StringAttribute{
						Computed: true,
					},
					"verified": schema.BoolAttribute{
						Computed: true,
					},
					"profile_picture": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"tweet": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[XTweetTweetDataSourceModel](ctx),
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
					"created_at": schema.StringAttribute{
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
