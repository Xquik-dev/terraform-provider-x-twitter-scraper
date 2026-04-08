// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IntegrationModel struct {
	ID               types.String                     `tfsdk:"id" json:"id,computed"`
	Type             types.String                     `tfsdk:"type" json:"type,required"`
	Config           *IntegrationConfigModel          `tfsdk:"config" json:"config,required,no_refresh"`
	Name             types.String                     `tfsdk:"name" json:"name,required"`
	EventTypes       *[]types.String                  `tfsdk:"event_types" json:"eventTypes,required"`
	IsActive         types.Bool                       `tfsdk:"is_active" json:"isActive,optional"`
	ScopeAllMonitors types.Bool                       `tfsdk:"scope_all_monitors" json:"scopeAllMonitors,optional"`
	SilentPush       types.Bool                       `tfsdk:"silent_push" json:"silentPush,optional"`
	Filters          *map[string]jsontypes.Normalized `tfsdk:"filters" json:"filters,optional"`
	MessageTemplate  *map[string]jsontypes.Normalized `tfsdk:"message_template" json:"messageTemplate,optional,no_refresh"`
	CreatedAt        timetypes.RFC3339                `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
}

func (m IntegrationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m IntegrationModel) MarshalJSONForUpdate(state IntegrationModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type IntegrationConfigModel struct {
	ChatID types.String `tfsdk:"chat_id" json:"chatId,required"`
}
