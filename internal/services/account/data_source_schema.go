// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Account info and settings",
		Attributes: map[string]schema.Attribute{
			"monitors_allowed": schema.Int64Attribute{
				Description:        "Deprecated. Monitor slots are unlimited, so this is always Number.MAX_SAFE_INTEGER.",
				Computed:           true,
				DeprecationMessage: "Monitor slots are unlimited. Use monitorBilling.unlimitedSlots instead.",
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
			"x_username": schema.StringAttribute{
				Description: "Linked X username, omitted when no X account is connected.",
				Computed:    true,
			},
			"credit_info": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AccountCreditInfoDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"auto_topup_amount_dollars": schema.Float64Attribute{
						Description: "Dollar amount charged when automatic top-up runs.",
						Computed:    true,
					},
					"auto_topup_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"auto_topup_threshold": schema.StringAttribute{
						Description: "Bigint string threshold that triggers automatic top-up when enabled.",
						Computed:    true,
					},
					"balance": schema.StringAttribute{
						Description: "Bigint string to preserve precision above Number.MAX_SAFE_INTEGER.",
						Computed:    true,
					},
					"lifetime_purchased": schema.StringAttribute{
						Description: "Total purchased credits as a bigint string.",
						Computed:    true,
					},
					"lifetime_used": schema.StringAttribute{
						Description: "Total consumed credits as a bigint string.",
						Computed:    true,
					},
				},
			},
			"monitor_billing": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AccountMonitorBillingDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"active_daily_estimate": schema.StringAttribute{
						Description: "Estimated daily credits for currently active monitors.",
						Computed:    true,
					},
					"active_hourly_burn": schema.StringAttribute{
						Description: "Credits charged each hour for currently active monitors.",
						Computed:    true,
					},
					"credits_per_active_monitor_day": schema.StringAttribute{
						Description: "Estimated daily credits for 1 active instant monitor.",
						Computed:    true,
					},
					"credits_per_active_monitor_hour": schema.StringAttribute{
						Description: "Hourly credits charged for 1 active instant monitor.",
						Computed:    true,
					},
					"events_included": schema.BoolAttribute{
						Description: "Webhook and event deliveries are included in monitor billing.",
						Computed:    true,
					},
					"instant_check_interval_seconds": schema.Int64Attribute{
						Description: "Active monitors check every 1 second.",
						Computed:    true,
					},
					"unlimited_slots": schema.BoolAttribute{
						Description: "Monitor slot count is unlimited.",
						Computed:    true,
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
