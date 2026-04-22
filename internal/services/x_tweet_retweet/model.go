// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet_retweet

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type XTweetRetweetModel struct {
	ID      types.String `tfsdk:"id" path:"id,required"`
	Account types.String `tfsdk:"account" json:"account,required"`
	Success types.Bool   `tfsdk:"success" json:"success,computed"`
}

func (m XTweetRetweetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XTweetRetweetModel) MarshalJSONForUpdate(state XTweetRetweetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
