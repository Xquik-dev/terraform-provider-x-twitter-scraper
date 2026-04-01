// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XTweetDataSourceModel struct {
	TweetID types.String                                          `tfsdk:"tweet_id" path:"tweetId,required"`
	Author  customfield.NestedObject[XTweetAuthorDataSourceModel] `tfsdk:"author" json:"author,computed"`
	Tweet   customfield.NestedObject[XTweetTweetDataSourceModel]  `tfsdk:"tweet" json:"tweet,computed"`
}

type XTweetAuthorDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	Followers      types.Int64  `tfsdk:"followers" json:"followers,computed"`
	Username       types.String `tfsdk:"username" json:"username,computed"`
	Verified       types.Bool   `tfsdk:"verified" json:"verified,computed"`
	ProfilePicture types.String `tfsdk:"profile_picture" json:"profilePicture,computed"`
}

type XTweetTweetDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	BookmarkCount types.Int64  `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount     types.Int64  `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64  `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64  `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64  `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text          types.String `tfsdk:"text" json:"text,computed"`
	ViewCount     types.Int64  `tfsdk:"view_count" json:"viewCount,computed"`
	CreatedAt     types.String `tfsdk:"created_at" json:"createdAt,computed"`
}
