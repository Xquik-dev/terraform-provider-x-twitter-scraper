// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*XUserDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "X data lookups (subscription required)",
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"followers": schema.Int64Attribute{
				Computed: true,
			},
			"following": schema.Int64Attribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"location": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"profile_picture": schema.StringAttribute{
				Computed: true,
			},
			"statuses_count": schema.Int64Attribute{
				Computed: true,
			},
			"verified": schema.BoolAttribute{
				Computed: true,
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
