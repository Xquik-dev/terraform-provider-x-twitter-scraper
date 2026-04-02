// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XTweetModel struct {
	Account        types.String    `tfsdk:"account" json:"account,required"`
	Text           types.String    `tfsdk:"text" json:"text,required"`
	AttachmentURL  types.String    `tfsdk:"attachment_url" json:"attachment_url,optional"`
	CommunityID    types.String    `tfsdk:"community_id" json:"community_id,optional"`
	IsNoteTweet    types.Bool      `tfsdk:"is_note_tweet" json:"is_note_tweet,optional"`
	ReplyToTweetID types.String    `tfsdk:"reply_to_tweet_id" json:"reply_to_tweet_id,optional"`
	MediaIDs       *[]types.String `tfsdk:"media_ids" json:"media_ids,optional"`
	Success        types.Bool      `tfsdk:"success" json:"success,computed"`
	TweetID        types.String    `tfsdk:"tweet_id" json:"tweetId,computed"`
}

func (m XTweetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XTweetModel) MarshalJSONForUpdate(state XTweetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
