// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monitor

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MonitorModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	Username   types.String      `tfsdk:"username" json:"username,required"`
	EventTypes *[]types.String   `tfsdk:"event_types" json:"eventTypes,required"`
	IsActive   types.Bool        `tfsdk:"is_active" json:"isActive,optional"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	XUserID    types.String      `tfsdk:"x_user_id" json:"xUserId,computed"`
}

func (m MonitorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MonitorModel) MarshalJSONForUpdate(state MonitorModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
