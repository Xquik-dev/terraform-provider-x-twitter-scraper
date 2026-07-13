// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet_like

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XTweetLikeModel struct {
	ID      types.String `tfsdk:"id" path:"id,required"`
	Account types.String `tfsdk:"account" json:"account,required"`
	Success types.Bool   `tfsdk:"success" json:"success,computed"`
}

func (m XTweetLikeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XTweetLikeModel) MarshalJSONForUpdate(state XTweetLikeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
