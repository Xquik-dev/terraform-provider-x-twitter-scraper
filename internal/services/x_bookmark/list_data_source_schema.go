// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_bookmark

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*XBookmarksDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "X data lookups (subscription required)",
		Attributes: map[string]schema.Attribute{
			"folder_id": schema.StringAttribute{
				Description: "Optional bookmark folder ID",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[XBookmarksItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"has_next_page": schema.BoolAttribute{
							Computed: true,
						},
						"next_cursor": schema.StringAttribute{
							Computed: true,
						},
						"tweets": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[XBookmarksTweetsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"text": schema.StringAttribute{
										Computed: true,
									},
									"author": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[XBookmarksTweetsAuthorDataSourceModel](ctx),
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
											"verified": schema.BoolAttribute{
												Computed: true,
											},
										},
									},
									"bookmark_count": schema.Int64Attribute{
										Computed: true,
									},
									"created_at": schema.StringAttribute{
										Computed: true,
									},
									"is_note_tweet": schema.BoolAttribute{
										Description: "True for Note Tweets (long-form content, up to 25,000 characters)",
										Computed:    true,
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
									"view_count": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *XBookmarksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *XBookmarksDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
