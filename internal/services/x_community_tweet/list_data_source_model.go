// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_community_tweet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-go"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type XCommunityTweetsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[XCommunityTweetsItemsDataSourceModel] `json:"data,computed"`
}

type XCommunityTweetsDataSourceModel struct {
	Q         types.String                                                       `tfsdk:"q" query:"q,required"`
	QueryType types.String                                                       `tfsdk:"query_type" query:"queryType,optional"`
	MaxItems  types.Int64                                                        `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[XCommunityTweetsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *XCommunityTweetsDataSourceModel) toListParams(_ context.Context) (params xtwitterscraper.XCommunityTweetListParams, diags diag.Diagnostics) {
	params = xtwitterscraper.XCommunityTweetListParams{
		Q: m.Q.ValueString(),
	}

	if !m.QueryType.IsNull() {
		params.QueryType = param.NewOpt(m.QueryType.ValueString())
	}

	return
}

type XCommunityTweetsItemsDataSourceModel struct {
	HasNextPage types.Bool                                                          `tfsdk:"has_next_page" json:"has_next_page,computed"`
	NextCursor  types.String                                                        `tfsdk:"next_cursor" json:"next_cursor,computed"`
	Tweets      customfield.NestedObjectList[XCommunityTweetsTweetsDataSourceModel] `tfsdk:"tweets" json:"tweets,computed"`
}

type XCommunityTweetsTweetsDataSourceModel struct {
	ID            types.String                                                          `tfsdk:"id" json:"id,computed"`
	Text          types.String                                                          `tfsdk:"text" json:"text,computed"`
	Author        customfield.NestedObject[XCommunityTweetsTweetsAuthorDataSourceModel] `tfsdk:"author" json:"author,computed"`
	BookmarkCount types.Int64                                                           `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	CreatedAt     types.String                                                          `tfsdk:"created_at" json:"createdAt,computed"`
	IsNoteTweet   types.Bool                                                            `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	LikeCount     types.Int64                                                           `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64                                                           `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64                                                           `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64                                                           `tfsdk:"retweet_count" json:"retweetCount,computed"`
	ViewCount     types.Int64                                                           `tfsdk:"view_count" json:"viewCount,computed"`
}

type XCommunityTweetsTweetsAuthorDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Username types.String `tfsdk:"username" json:"username,computed"`
	Verified types.Bool   `tfsdk:"verified" json:"verified,computed"`
}
