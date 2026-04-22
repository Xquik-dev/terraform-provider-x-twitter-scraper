// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type AccountDataSourceModel struct {
	MonitorsAllowed types.Int64                                                `tfsdk:"monitors_allowed" json:"monitorsAllowed,computed"`
	MonitorsUsed    types.Int64                                                `tfsdk:"monitors_used" json:"monitorsUsed,computed"`
	Plan            types.String                                               `tfsdk:"plan" json:"plan,computed"`
	CreditInfo      customfield.NestedObject[AccountCreditInfoDataSourceModel] `tfsdk:"credit_info" json:"creditInfo,computed"`
}

type AccountCreditInfoDataSourceModel struct {
	AutoTopupEnabled  types.Bool  `tfsdk:"auto_topup_enabled" json:"autoTopupEnabled,computed"`
	Balance           types.Int64 `tfsdk:"balance" json:"balance,computed"`
	LifetimePurchased types.Int64 `tfsdk:"lifetime_purchased" json:"lifetimePurchased,computed"`
	LifetimeUsed      types.Int64 `tfsdk:"lifetime_used" json:"lifetimeUsed,computed"`
}
