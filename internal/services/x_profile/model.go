// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_profile

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XProfileModel struct {
	Account     types.String `tfsdk:"account" json:"account,required"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Location    types.String `tfsdk:"location" json:"location,optional"`
	Name        types.String `tfsdk:"name" json:"name,optional"`
	URL         types.String `tfsdk:"url" json:"url,optional"`
	Success     types.Bool   `tfsdk:"success" json:"success,computed"`
}

func (m XProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XProfileModel) MarshalJSONForUpdate(state XProfileModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
