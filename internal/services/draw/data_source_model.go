// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package draw

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DrawDataSourceModel struct {
	ID      types.String                                             `tfsdk:"id" path:"id,required"`
	Draw    customfield.NestedObject[DrawDrawDataSourceModel]        `tfsdk:"draw" json:"draw,computed"`
	Winners customfield.NestedObjectList[DrawWinnersDataSourceModel] `tfsdk:"winners" json:"winners,computed"`
}

type DrawDrawDataSourceModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Status              types.String      `tfsdk:"status" json:"status,computed"`
	TotalEntries        types.Int64       `tfsdk:"total_entries" json:"totalEntries,computed"`
	TweetAuthorUsername types.String      `tfsdk:"tweet_author_username" json:"tweetAuthorUsername,computed"`
	TweetID             types.String      `tfsdk:"tweet_id" json:"tweetId,computed"`
	TweetLikeCount      types.Int64       `tfsdk:"tweet_like_count" json:"tweetLikeCount,computed"`
	TweetQuoteCount     types.Int64       `tfsdk:"tweet_quote_count" json:"tweetQuoteCount,computed"`
	TweetReplyCount     types.Int64       `tfsdk:"tweet_reply_count" json:"tweetReplyCount,computed"`
	TweetRetweetCount   types.Int64       `tfsdk:"tweet_retweet_count" json:"tweetRetweetCount,computed"`
	TweetText           types.String      `tfsdk:"tweet_text" json:"tweetText,computed"`
	TweetURL            types.String      `tfsdk:"tweet_url" json:"tweetUrl,computed"`
	ValidEntries        types.Int64       `tfsdk:"valid_entries" json:"validEntries,computed"`
	DrawnAt             timetypes.RFC3339 `tfsdk:"drawn_at" json:"drawnAt,computed" format:"date-time"`
}

type DrawWinnersDataSourceModel struct {
	AuthorUsername types.String `tfsdk:"author_username" json:"authorUsername,computed"`
	IsBackup       types.Bool   `tfsdk:"is_backup" json:"isBackup,computed"`
	Position       types.Int64  `tfsdk:"position" json:"position,computed"`
	TweetID        types.String `tfsdk:"tweet_id" json:"tweetId,computed"`
}
