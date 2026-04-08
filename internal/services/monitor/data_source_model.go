// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monitor

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type MonitorDataSourceModel struct {
	ID         types.String                   `tfsdk:"id" path:"id,required"`
	CreatedAt  timetypes.RFC3339              `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	IsActive   types.Bool                     `tfsdk:"is_active" json:"isActive,computed"`
	Username   types.String                   `tfsdk:"username" json:"username,computed"`
	XUserID    types.String                   `tfsdk:"x_user_id" json:"xUserId,computed"`
	EventTypes customfield.List[types.String] `tfsdk:"event_types" json:"eventTypes,computed"`
}
