// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package draw

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DrawDataSourceModel struct {
	ID      types.String                                             `tfsdk:"id" path:"id,required"`
	Draw    customfield.NestedObject[DrawDrawDataSourceModel]        `tfsdk:"draw" json:"draw" api:"computed"`
	Winners customfield.NestedObjectList[DrawWinnersDataSourceModel] `tfsdk:"winners" json:"winners" api:"computed"`
}

type DrawDrawDataSourceModel struct {
	ID                  types.String      `tfsdk:"id" json:"id" api:"computed"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	Status              types.String      `tfsdk:"status" json:"status" api:"computed"`
	TotalEntries        types.Int64       `tfsdk:"total_entries" json:"totalEntries" api:"computed"`
	TweetAuthorUsername types.String      `tfsdk:"tweet_author_username" json:"tweetAuthorUsername" api:"computed"`
	TweetID             types.String      `tfsdk:"tweet_id" json:"tweetId" api:"computed"`
	TweetLikeCount      types.Int64       `tfsdk:"tweet_like_count" json:"tweetLikeCount" api:"computed"`
	TweetQuoteCount     types.Int64       `tfsdk:"tweet_quote_count" json:"tweetQuoteCount" api:"computed"`
	TweetReplyCount     types.Int64       `tfsdk:"tweet_reply_count" json:"tweetReplyCount" api:"computed"`
	TweetRetweetCount   types.Int64       `tfsdk:"tweet_retweet_count" json:"tweetRetweetCount" api:"computed"`
	TweetText           types.String      `tfsdk:"tweet_text" json:"tweetText" api:"computed"`
	TweetURL            types.String      `tfsdk:"tweet_url" json:"tweetUrl" api:"computed"`
	ValidEntries        types.Int64       `tfsdk:"valid_entries" json:"validEntries" api:"computed"`
	DrawnAt             timetypes.RFC3339 `tfsdk:"drawn_at" json:"drawnAt" api:"computed" format:"date-time"`
}

type DrawWinnersDataSourceModel struct {
	AuthorUsername types.String `tfsdk:"author_username" json:"authorUsername" api:"computed"`
	IsBackup       types.Bool   `tfsdk:"is_backup" json:"isBackup" api:"computed"`
	Position       types.Int64  `tfsdk:"position" json:"position" api:"computed"`
	TweetID        types.String `tfsdk:"tweet_id" json:"tweetId" api:"computed"`
}
