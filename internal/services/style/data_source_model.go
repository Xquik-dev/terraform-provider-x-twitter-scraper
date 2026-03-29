// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package style

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type StyleDataSourceModel struct {
	Username     types.String                                             `tfsdk:"username" path:"username,required"`
	FetchedAt    timetypes.RFC3339                                        `tfsdk:"fetched_at" json:"fetchedAt,computed" format:"date-time"`
	IsOwnAccount types.Bool                                               `tfsdk:"is_own_account" json:"isOwnAccount,computed"`
	TweetCount   types.Int64                                              `tfsdk:"tweet_count" json:"tweetCount,computed"`
	XUsername    types.String                                             `tfsdk:"x_username" json:"xUsername,computed"`
	Tweets       customfield.NestedObjectList[StyleTweetsDataSourceModel] `tfsdk:"tweets" json:"tweets,computed"`
}

type StyleTweetsDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	Text           types.String `tfsdk:"text" json:"text,computed"`
	AuthorUsername types.String `tfsdk:"author_username" json:"authorUsername,computed"`
	CreatedAt      types.String `tfsdk:"created_at" json:"createdAt,computed"`
}
