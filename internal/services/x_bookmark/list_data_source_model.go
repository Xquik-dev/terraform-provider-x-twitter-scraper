// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_bookmark

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-go"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type XBookmarksDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[XBookmarksItemsDataSourceModel] `json:"data,computed"`
}

type XBookmarksDataSourceModel struct {
	FolderID types.String                                                 `tfsdk:"folder_id" query:"folderId,optional"`
	MaxItems types.Int64                                                  `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[XBookmarksItemsDataSourceModel] `tfsdk:"items"`
}

func (m *XBookmarksDataSourceModel) toListParams(_ context.Context) (params xtwitterscraper.XBookmarkListParams, diags diag.Diagnostics) {
	params = xtwitterscraper.XBookmarkListParams{}

	if !m.FolderID.IsNull() {
		params.FolderID = param.NewOpt(m.FolderID.ValueString())
	}

	return
}

type XBookmarksItemsDataSourceModel struct {
	HasNextPage types.Bool                                                    `tfsdk:"has_next_page" json:"has_next_page,computed"`
	NextCursor  types.String                                                  `tfsdk:"next_cursor" json:"next_cursor,computed"`
	Tweets      customfield.NestedObjectList[XBookmarksTweetsDataSourceModel] `tfsdk:"tweets" json:"tweets,computed"`
}

type XBookmarksTweetsDataSourceModel struct {
	ID            types.String                                                    `tfsdk:"id" json:"id,computed"`
	Text          types.String                                                    `tfsdk:"text" json:"text,computed"`
	Author        customfield.NestedObject[XBookmarksTweetsAuthorDataSourceModel] `tfsdk:"author" json:"author,computed"`
	BookmarkCount types.Int64                                                     `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	CreatedAt     types.String                                                    `tfsdk:"created_at" json:"createdAt,computed"`
	IsNoteTweet   types.Bool                                                      `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	LikeCount     types.Int64                                                     `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64                                                     `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64                                                     `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64                                                     `tfsdk:"retweet_count" json:"retweetCount,computed"`
	ViewCount     types.Int64                                                     `tfsdk:"view_count" json:"viewCount,computed"`
}

type XBookmarksTweetsAuthorDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Username types.String `tfsdk:"username" json:"username,computed"`
	Verified types.Bool   `tfsdk:"verified" json:"verified,computed"`
}
