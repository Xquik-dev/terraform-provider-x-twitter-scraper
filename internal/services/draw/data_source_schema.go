// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package draw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*DrawDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Giveaway draws from tweet replies",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"draw": schema.SingleNestedAttribute{
				Description: "Full giveaway draw with tweet metrics, entries, and timing.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[DrawDrawDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"status": schema.StringAttribute{
						Computed: true,
					},
					"total_entries": schema.Int64Attribute{
						Computed: true,
					},
					"tweet_author_username": schema.StringAttribute{
						Computed: true,
					},
					"tweet_id": schema.StringAttribute{
						Computed: true,
					},
					"tweet_like_count": schema.Int64Attribute{
						Computed: true,
					},
					"tweet_quote_count": schema.Int64Attribute{
						Computed: true,
					},
					"tweet_reply_count": schema.Int64Attribute{
						Computed: true,
					},
					"tweet_retweet_count": schema.Int64Attribute{
						Computed: true,
					},
					"tweet_text": schema.StringAttribute{
						Computed: true,
					},
					"tweet_url": schema.StringAttribute{
						Computed: true,
					},
					"valid_entries": schema.Int64Attribute{
						Computed: true,
					},
					"drawn_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
				},
			},
			"winners": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[DrawWinnersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"author_username": schema.StringAttribute{
							Computed: true,
						},
						"is_backup": schema.BoolAttribute{
							Computed: true,
						},
						"position": schema.Int64Attribute{
							Computed: true,
						},
						"tweet_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *DrawDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DrawDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
