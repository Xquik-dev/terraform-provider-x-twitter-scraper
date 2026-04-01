// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_community_join

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XCommunityJoinModel struct {
	ID            types.String `tfsdk:"id" path:"id,required"`
	Account       types.String `tfsdk:"account" json:"account,required"`
	CommunityID   types.String `tfsdk:"community_id" json:"communityId,computed"`
	CommunityName types.String `tfsdk:"community_name" json:"communityName,computed"`
	Success       types.Bool   `tfsdk:"success" json:"success,computed"`
}

func (m XCommunityJoinModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XCommunityJoinModel) MarshalJSONForUpdate(state XCommunityJoinModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
