// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package event

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EventDataSourceModel struct {
	ID         types.String                          `tfsdk:"id" path:"id,required"`
	MonitorID  types.String                          `tfsdk:"monitor_id" json:"monitorId" api:"computed"`
	OccurredAt timetypes.RFC3339                     `tfsdk:"occurred_at" json:"occurredAt" api:"computed" format:"date-time"`
	Type       types.String                          `tfsdk:"type" json:"type" api:"computed"`
	Username   types.String                          `tfsdk:"username" json:"username" api:"computed"`
	XEventID   types.String                          `tfsdk:"x_event_id" json:"xEventId" api:"computed"`
	Data       customfield.Map[jsontypes.Normalized] `tfsdk:"data" json:"data" api:"computed"`
}
