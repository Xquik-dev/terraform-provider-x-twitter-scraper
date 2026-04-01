// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package extraction

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ExtractionDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Bulk data extraction (20 tool types)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"after": schema.StringAttribute{
				Description: "Cursor for pagination",
				Optional:    true,
			},
			"limit": schema.Int64Attribute{
				Computed: true,
				Optional: true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1000),
				},
			},
			"has_more": schema.BoolAttribute{
				Computed: true,
			},
			"next_cursor": schema.StringAttribute{
				Computed: true,
			},
			"job": schema.MapAttribute{
				Description: "Extraction job metadata — shape varies by tool type (JSON)",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"results": schema.ListAttribute{
				Computed:   true,
				CustomType: customfield.NewListType[customfield.Map[jsontypes.Normalized]](ctx),
				ElementType: types.MapType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
		},
	}
}

func (d *ExtractionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ExtractionDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
