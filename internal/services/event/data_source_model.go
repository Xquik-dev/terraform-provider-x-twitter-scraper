// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EventDataSourceModel struct {
	ID         types.String                          `tfsdk:"id" path:"id,required"`
	MonitorID  types.String                          `tfsdk:"monitor_id" json:"monitorId,computed"`
	OccurredAt timetypes.RFC3339                     `tfsdk:"occurred_at" json:"occurredAt,computed" format:"date-time"`
	Type       types.String                          `tfsdk:"type" json:"type,computed"`
	Username   types.String                          `tfsdk:"username" json:"username,computed"`
	XEventID   types.String                          `tfsdk:"x_event_id" json:"xEventId,computed"`
	Data       customfield.Map[jsontypes.Normalized] `tfsdk:"data" json:"data,computed"`
}
