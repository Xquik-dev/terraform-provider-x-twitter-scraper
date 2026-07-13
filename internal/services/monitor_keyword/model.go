// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monitor_keyword

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MonitorKeywordModel struct {
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	Query         types.String      `tfsdk:"query" json:"query,required"`
	EventTypes    *[]types.String   `tfsdk:"event_types" json:"eventTypes,required"`
	IsActive      types.Bool        `tfsdk:"is_active" json:"isActive,optional"`
	CreatedAt     timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	NextBillingAt timetypes.RFC3339 `tfsdk:"next_billing_at" json:"nextBillingAt,computed" format:"date-time"`
}

func (m MonitorKeywordModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MonitorKeywordModel) MarshalJSONForUpdate(state MonitorKeywordModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
