// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Account info and settings",
		Attributes: map[string]schema.Attribute{
			"monitors_allowed": schema.Int64Attribute{
				Computed: true,
			},
			"monitors_used": schema.Int64Attribute{
				Computed: true,
			},
			"plan": schema.StringAttribute{
				Description: `Available values: "active", "inactive".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "inactive"),
				},
			},
			"credit_info": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AccountCreditInfoDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"auto_topup_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"balance": schema.Int64Attribute{
						Computed: true,
					},
					"lifetime_purchased": schema.Int64Attribute{
						Computed: true,
					},
					"lifetime_used": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *AccountDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccountDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
