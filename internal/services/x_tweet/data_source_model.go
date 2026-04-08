// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type XTweetDataSourceModel struct {
	ID     types.String                                          `tfsdk:"id" path:"id,required"`
	Author customfield.NestedObject[XTweetAuthorDataSourceModel] `tfsdk:"author" json:"author,computed"`
	Tweet  customfield.NestedObject[XTweetTweetDataSourceModel]  `tfsdk:"tweet" json:"tweet,computed"`
}

type XTweetAuthorDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	Followers      types.Int64  `tfsdk:"followers" json:"followers,computed"`
	Username       types.String `tfsdk:"username" json:"username,computed"`
	Verified       types.Bool   `tfsdk:"verified" json:"verified,computed"`
	ProfilePicture types.String `tfsdk:"profile_picture" json:"profilePicture,computed"`
}

type XTweetTweetDataSourceModel struct {
	ID             types.String                                                  `tfsdk:"id" json:"id,computed"`
	BookmarkCount  types.Int64                                                   `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount      types.Int64                                                   `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount     types.Int64                                                   `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount     types.Int64                                                   `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount   types.Int64                                                   `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text           types.String                                                  `tfsdk:"text" json:"text,computed"`
	ViewCount      types.Int64                                                   `tfsdk:"view_count" json:"viewCount,computed"`
	ConversationID types.String                                                  `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt      types.String                                                  `tfsdk:"created_at" json:"createdAt,computed"`
	Entities       customfield.Map[jsontypes.Normalized]                         `tfsdk:"entities" json:"entities,computed"`
	IsNoteTweet    types.Bool                                                    `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus  types.Bool                                                    `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply        types.Bool                                                    `tfsdk:"is_reply" json:"isReply,computed"`
	Media          customfield.NestedObjectList[XTweetTweetMediaDataSourceModel] `tfsdk:"media" json:"media,computed"`
	QuotedTweet    customfield.Map[jsontypes.Normalized]                         `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	Source         types.String                                                  `tfsdk:"source" json:"source,computed"`
}

type XTweetTweetMediaDataSourceModel struct {
	MediaURL types.String `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type     types.String `tfsdk:"type" json:"type,computed"`
	URL      types.String `tfsdk:"url" json:"url,computed"`
}
