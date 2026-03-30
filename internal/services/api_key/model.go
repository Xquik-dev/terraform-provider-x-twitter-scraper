// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIKeyModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Name      types.String      `tfsdk:"name" json:"name,optional"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	FullKey   types.String      `tfsdk:"full_key" json:"fullKey,computed"`
	Prefix    types.String      `tfsdk:"prefix" json:"prefix,computed"`
}

func (m APIKeyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m APIKeyModel) MarshalJSONForUpdate(state APIKeyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
