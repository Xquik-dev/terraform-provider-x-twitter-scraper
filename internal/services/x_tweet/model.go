// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XTweetModel struct {
	ID             types.String                                `tfsdk:"id" path:"id,optional"`
	Account        types.String                                `tfsdk:"account" json:"account,required,no_refresh"`
	AttachmentURL  types.String                                `tfsdk:"attachment_url" json:"attachment_url,optional,no_refresh"`
	CommunityID    types.String                                `tfsdk:"community_id" json:"community_id,optional,no_refresh"`
	IsNoteTweet    types.Bool                                  `tfsdk:"is_note_tweet" json:"is_note_tweet,optional,no_refresh"`
	ReplyToTweetID types.String                                `tfsdk:"reply_to_tweet_id" json:"reply_to_tweet_id,optional,no_refresh"`
	Text           types.String                                `tfsdk:"text" json:"text,optional,no_refresh"`
	Media          *[]types.String                             `tfsdk:"media" json:"media,optional,no_refresh"`
	MediaIDs       *[]types.String                             `tfsdk:"media_ids" json:"media_ids,optional,no_refresh"`
	Success        types.Bool                                  `tfsdk:"success" json:"success,computed,no_refresh"`
	TweetID        types.String                                `tfsdk:"tweet_id" json:"tweetId,computed,no_refresh"`
	Author         customfield.NestedObject[XTweetAuthorModel] `tfsdk:"author" json:"author,computed"`
	Tweet          customfield.NestedObject[XTweetTweetModel]  `tfsdk:"tweet" json:"tweet,computed"`
}

func (m XTweetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XTweetModel) MarshalJSONForUpdate(state XTweetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type XTweetAuthorModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	Followers      types.Int64  `tfsdk:"followers" json:"followers,computed"`
	Username       types.String `tfsdk:"username" json:"username,computed"`
	Verified       types.Bool   `tfsdk:"verified" json:"verified,computed"`
	ProfilePicture types.String `tfsdk:"profile_picture" json:"profilePicture,computed"`
}

type XTweetTweetModel struct {
	ID             types.String                                        `tfsdk:"id" json:"id,computed"`
	BookmarkCount  types.Int64                                         `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount      types.Int64                                         `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount     types.Int64                                         `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount     types.Int64                                         `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount   types.Int64                                         `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text           types.String                                        `tfsdk:"text" json:"text,computed"`
	ViewCount      types.Int64                                         `tfsdk:"view_count" json:"viewCount,computed"`
	ConversationID types.String                                        `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt      types.String                                        `tfsdk:"created_at" json:"createdAt,computed"`
	Entities       customfield.Map[jsontypes.Normalized]               `tfsdk:"entities" json:"entities,computed"`
	IsNoteTweet    types.Bool                                          `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus  types.Bool                                          `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply        types.Bool                                          `tfsdk:"is_reply" json:"isReply,computed"`
	Media          customfield.NestedObjectList[XTweetTweetMediaModel] `tfsdk:"media" json:"media,computed"`
	QuotedTweet    customfield.Map[jsontypes.Normalized]               `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	Source         types.String                                        `tfsdk:"source" json:"source,computed"`
}

type XTweetTweetMediaModel struct {
	MediaURL types.String `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type     types.String `tfsdk:"type" json:"type,computed"`
	URL      types.String `tfsdk:"url" json:"url,computed"`
}
