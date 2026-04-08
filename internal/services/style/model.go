// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package style

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type StyleModel struct {
	ID           types.String         `tfsdk:"id" path:"id,required"`
	Label        types.String         `tfsdk:"label" json:"label,required,no_refresh"`
	Tweets       *[]*StyleTweetsModel `tfsdk:"tweets" json:"tweets,required"`
	FetchedAt    timetypes.RFC3339    `tfsdk:"fetched_at" json:"fetchedAt,computed" format:"date-time"`
	IsOwnAccount types.Bool           `tfsdk:"is_own_account" json:"isOwnAccount,computed"`
	TweetCount   types.Int64          `tfsdk:"tweet_count" json:"tweetCount,computed"`
	XUsername    types.String         `tfsdk:"x_username" json:"xUsername,computed"`
}

func (m StyleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StyleModel) MarshalJSONForUpdate(state StyleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type StyleTweetsModel struct {
	Text types.String `tfsdk:"text" json:"text,required"`
}
