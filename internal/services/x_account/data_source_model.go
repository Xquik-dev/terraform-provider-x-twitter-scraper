// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_account

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XAccountDataSourceModel struct {
	ID                types.String      `tfsdk:"id" path:"id,required"`
	CookiesObtainedAt timetypes.RFC3339 `tfsdk:"cookies_obtained_at" json:"cookiesObtainedAt,computed" format:"date-time"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Health            types.String      `tfsdk:"health" json:"health,computed"`
	ProxyCountry      types.String      `tfsdk:"proxy_country" json:"proxyCountry,computed"`
	Status            types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
	XUserID           types.String      `tfsdk:"x_user_id" json:"xUserId,computed"`
	XUsername         types.String      `tfsdk:"x_username" json:"xUsername,computed"`
}
