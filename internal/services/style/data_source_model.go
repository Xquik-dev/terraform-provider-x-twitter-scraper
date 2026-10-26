// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package style

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StyleDataSourceModel struct {
	ID           types.String                                             `tfsdk:"id" path:"id,required"`
	FetchedAt    timetypes.RFC3339                                        `tfsdk:"fetched_at" json:"fetchedAt" api:"computed" format:"date-time"`
	IsOwnAccount types.Bool                                               `tfsdk:"is_own_account" json:"isOwnAccount" api:"computed"`
	TweetCount   types.Int64                                              `tfsdk:"tweet_count" json:"tweetCount" api:"computed"`
	XUsername    types.String                                             `tfsdk:"x_username" json:"xUsername" api:"computed"`
	Tweets       customfield.NestedObjectList[StyleTweetsDataSourceModel] `tfsdk:"tweets" json:"tweets" api:"computed"`
}

type StyleTweetsDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id" api:"computed"`
	Text           types.String `tfsdk:"text" json:"text" api:"computed"`
	AuthorUsername types.String `tfsdk:"author_username" json:"authorUsername" api:"computed"`
	CreatedAt      types.String `tfsdk:"created_at" json:"createdAt" api:"computed"`
}
