// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type IntegrationDataSourceModel struct {
	ID               types.String                          `tfsdk:"id" path:"id,required"`
	CreatedAt        timetypes.RFC3339                     `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	IsActive         types.Bool                            `tfsdk:"is_active" json:"isActive,computed"`
	MessageTemplate  types.String                          `tfsdk:"message_template" json:"messageTemplate,computed"`
	Name             types.String                          `tfsdk:"name" json:"name,computed"`
	ScopeAllMonitors types.Bool                            `tfsdk:"scope_all_monitors" json:"scopeAllMonitors,computed"`
	SilentPush       types.Bool                            `tfsdk:"silent_push" json:"silentPush,computed"`
	Type             types.String                          `tfsdk:"type" json:"type,computed"`
	Config           customfield.Map[jsontypes.Normalized] `tfsdk:"config" json:"config,computed"`
	EventTypes       customfield.List[types.String]        `tfsdk:"event_types" json:"eventTypes,computed"`
	Filters          customfield.Map[jsontypes.Normalized] `tfsdk:"filters" json:"filters,computed"`
}
