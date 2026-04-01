// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountDataSourceModel struct {
	MonitorsAllowed types.Int64                                                   `tfsdk:"monitors_allowed" json:"monitorsAllowed,computed"`
	MonitorsUsed    types.Int64                                                   `tfsdk:"monitors_used" json:"monitorsUsed,computed"`
	Plan            types.String                                                  `tfsdk:"plan" json:"plan,computed"`
	CurrentPeriod   customfield.NestedObject[AccountCurrentPeriodDataSourceModel] `tfsdk:"current_period" json:"currentPeriod,computed"`
}

type AccountCurrentPeriodDataSourceModel struct {
	End          timetypes.RFC3339 `tfsdk:"end" json:"end,computed" format:"date-time"`
	Start        timetypes.RFC3339 `tfsdk:"start" json:"start,computed" format:"date-time"`
	UsagePercent types.Float64     `tfsdk:"usage_percent" json:"usagePercent,computed"`
}
