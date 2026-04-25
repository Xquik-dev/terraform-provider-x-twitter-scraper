// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_account

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type XAccountModel struct {
	ID                types.String      `tfsdk:"id" json:"id,computed"`
	Email             types.String      `tfsdk:"email" json:"email,required,no_refresh"`
	Password          types.String      `tfsdk:"password" json:"password,required,no_refresh"`
	Username          types.String      `tfsdk:"username" json:"username,required,no_refresh"`
	ProxyCountry      types.String      `tfsdk:"proxy_country" json:"proxy_country,optional"`
	TotpSecret        types.String      `tfsdk:"totp_secret" json:"totp_secret,optional,no_refresh"`
	CookiesObtainedAt timetypes.RFC3339 `tfsdk:"cookies_obtained_at" json:"cookiesObtainedAt,computed" format:"date-time"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Health            types.String      `tfsdk:"health" json:"health,computed"`
	LoginCountry      types.String      `tfsdk:"login_country" json:"loginCountry,computed,no_refresh"`
	Status            types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
	XUserID           types.String      `tfsdk:"x_user_id" json:"xUserId,computed"`
	XUsername         types.String      `tfsdk:"x_username" json:"xUsername,computed"`
}

func (m XAccountModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XAccountModel) MarshalJSONForUpdate(state XAccountModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
