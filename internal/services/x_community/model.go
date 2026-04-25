// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_community

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type XCommunityModel struct {
	ID            types.String `tfsdk:"id" path:"id,optional"`
	Account       types.String `tfsdk:"account" json:"account,required"`
	Name          types.String `tfsdk:"name" json:"name,required"`
	Description   types.String `tfsdk:"description" json:"description,optional"`
	CommunityID   types.String `tfsdk:"community_id" json:"communityId,computed"`
	CommunityName types.String `tfsdk:"community_name" json:"communityName,computed"`
	Success       types.Bool   `tfsdk:"success" json:"success,computed"`
}

func (m XCommunityModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XCommunityModel) MarshalJSONForUpdate(state XCommunityModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
