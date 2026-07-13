// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user_follow

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
