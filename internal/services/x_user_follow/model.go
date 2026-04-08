// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user_follow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type XUserFollowModel struct {
	ID      types.String `tfsdk:"id" path:"id,required"`
	Account types.String `tfsdk:"account" json:"account,required"`
	Success types.Bool   `tfsdk:"success" json:"success,computed"`
}

func (m XUserFollowModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XUserFollowModel) MarshalJSONForUpdate(state XUserFollowModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
