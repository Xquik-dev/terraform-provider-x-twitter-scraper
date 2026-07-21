// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XTweetDataSourceModel struct {
	ID     types.String                                         `tfsdk:"id" path:"id,required"`
	Tweet  customfield.NestedObject[XTweetTweetDataSourceModel] `tfsdk:"tweet" json:"tweet,computed"`
	Author jsontypes.Normalized                                 `tfsdk:"author" json:"author,computed"`
}

type XTweetTweetDataSourceModel struct {
	ID                types.String                                                          `tfsdk:"id" json:"id,computed"`
	BookmarkCount     types.Int64                                                           `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount         types.Int64                                                           `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount        types.Int64                                                           `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount        types.Int64                                                           `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount      types.Int64                                                           `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text              types.String                                                          `tfsdk:"text" json:"text,computed"`
	ViewCount         types.Int64                                                           `tfsdk:"view_count" json:"viewCount,computed"`
	Author            jsontypes.Normalized                                                  `tfsdk:"author" json:"author,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                          `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                          `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                         `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                 `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                          `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                          `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                          `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                            `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                            `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                            `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                            `tfsdk:"is_reply" json:"isReply,computed"`
	Lang              types.String                                                          `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetMediaDataSourceModel]         `tfsdk:"media" json:"media,computed"`
	QuotedTweet       customfield.NestedObject[XTweetTweetQuotedTweetDataSourceModel]       `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	RetweetedTweet    customfield.NestedObject[XTweetTweetRetweetedTweetDataSourceModel]    `tfsdk:"retweeted_tweet" json:"retweeted_tweet,computed"`
	Source            types.String                                                          `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                          `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetContentDisclosureAIGeneratedDataSourceModel struct {
	CanEdit             types.Bool   `tfsdk:"can_edit" json:"canEdit,computed"`
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetMediaDataSourceModel struct {
	MediaURL      types.String                                                               `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type          types.String                                                               `tfsdk:"type" json:"type,computed"`
	URL           types.String                                                               `tfsdk:"url" json:"url,computed"`
	VideoVariants customfield.NestedObjectList[XTweetTweetMediaVideoVariantsDataSourceModel] `tfsdk:"video_variants" json:"videoVariants,computed"`
}

type XTweetTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}

type XTweetTweetQuotedTweetDataSourceModel struct {
	ID                types.String                                                                     `tfsdk:"id" json:"id,computed"`
	BookmarkCount     types.Int64                                                                      `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount         types.Int64                                                                      `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount        types.Int64                                                                      `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount        types.Int64                                                                      `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount      types.Int64                                                                      `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text              types.String                                                                     `tfsdk:"text" json:"text,computed"`
	ViewCount         types.Int64                                                                      `tfsdk:"view_count" json:"viewCount,computed"`
	Author            customfield.NestedObject[XTweetTweetQuotedTweetAuthorDataSourceModel]            `tfsdk:"author" json:"author,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                                     `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                                     `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                                    `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                            `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                                     `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                                     `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                                     `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                                       `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                                       `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                                       `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                                       `tfsdk:"is_reply" json:"isReply,computed"`
	Lang              types.String                                                                     `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetQuotedTweetMediaDataSourceModel]         `tfsdk:"media" json:"media,computed"`
	Source            types.String                                                                     `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                                     `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                                     `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetQuotedTweetAuthorDataSourceModel struct {
	ID                  types.String                          `tfsdk:"id" json:"id,computed"`
	Name                types.String                          `tfsdk:"name" json:"name,computed"`
	Username            types.String                          `tfsdk:"username" json:"username,computed"`
	AutomatedBy         types.String                          `tfsdk:"automated_by" json:"automatedBy,computed"`
	CanDm               types.Bool                            `tfsdk:"can_dm" json:"canDm,computed"`
	CommunityRole       types.String                          `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture        types.String                          `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt           types.String                          `tfsdk:"created_at" json:"createdAt,computed"`
	Description         types.String                          `tfsdk:"description" json:"description,computed"`
	FavouritesCount     types.Int64                           `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers           types.Int64                           `tfsdk:"followers" json:"followers,computed"`
	Following           types.Int64                           `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines  types.Bool                            `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	IsAutomated         types.Bool                            `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified      types.Bool                            `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsTranslator        types.Bool                            `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified          types.Bool                            `tfsdk:"is_verified" json:"isVerified,computed"`
	Location            types.String                          `tfsdk:"location" json:"location,computed"`
	MediaCount          types.Int64                           `tfsdk:"media_count" json:"mediaCount,computed"`
	PinnedTweetIDs      customfield.List[types.String]        `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive   types.Bool                            `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio          customfield.Map[jsontypes.Normalized] `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL    types.String                          `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfilePicture      types.String                          `tfsdk:"profile_picture" json:"profilePicture,computed"`
	Protected           types.Bool                            `tfsdk:"protected" json:"protected,computed"`
	StatusesCount       types.Int64                           `tfsdk:"statuses_count" json:"statusesCount,computed"`
	Unavailable         types.Bool                            `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason   types.String                          `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                 types.String                          `tfsdk:"url" json:"url,computed"`
	Verified            types.Bool                            `tfsdk:"verified" json:"verified,computed"`
	VerifiedType        types.String                          `tfsdk:"verified_type" json:"verifiedType,computed"`
	ViewerFollowedBy    types.Bool                            `tfsdk:"viewer_followed_by" json:"viewerFollowedBy,computed"`
	ViewerFollowing     types.Bool                            `tfsdk:"viewer_following" json:"viewerFollowing,computed"`
	WithheldInCountries customfield.List[types.String]        `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAIGeneratedDataSourceModel struct {
	CanEdit             types.Bool   `tfsdk:"can_edit" json:"canEdit,computed"`
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetQuotedTweetMediaDataSourceModel struct {
	MediaURL      types.String                                                                          `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type          types.String                                                                          `tfsdk:"type" json:"type,computed"`
	URL           types.String                                                                          `tfsdk:"url" json:"url,computed"`
	VideoVariants customfield.NestedObjectList[XTweetTweetQuotedTweetMediaVideoVariantsDataSourceModel] `tfsdk:"video_variants" json:"videoVariants,computed"`
}

type XTweetTweetQuotedTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}

type XTweetTweetRetweetedTweetDataSourceModel struct {
	ID                types.String                                                                        `tfsdk:"id" json:"id,computed"`
	BookmarkCount     types.Int64                                                                         `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount         types.Int64                                                                         `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount        types.Int64                                                                         `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount        types.Int64                                                                         `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount      types.Int64                                                                         `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text              types.String                                                                        `tfsdk:"text" json:"text,computed"`
	ViewCount         types.Int64                                                                         `tfsdk:"view_count" json:"viewCount,computed"`
	Author            customfield.NestedObject[XTweetTweetRetweetedTweetAuthorDataSourceModel]            `tfsdk:"author" json:"author,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                                        `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                                        `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                                       `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                               `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                                        `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                                        `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                                        `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                                          `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                                          `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                                          `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                                          `tfsdk:"is_reply" json:"isReply,computed"`
	Lang              types.String                                                                        `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaDataSourceModel]         `tfsdk:"media" json:"media,computed"`
	Source            types.String                                                                        `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                                        `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                                        `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetRetweetedTweetAuthorDataSourceModel struct {
	ID                  types.String                          `tfsdk:"id" json:"id,computed"`
	Name                types.String                          `tfsdk:"name" json:"name,computed"`
	Username            types.String                          `tfsdk:"username" json:"username,computed"`
	AutomatedBy         types.String                          `tfsdk:"automated_by" json:"automatedBy,computed"`
	CanDm               types.Bool                            `tfsdk:"can_dm" json:"canDm,computed"`
	CommunityRole       types.String                          `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture        types.String                          `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt           types.String                          `tfsdk:"created_at" json:"createdAt,computed"`
	Description         types.String                          `tfsdk:"description" json:"description,computed"`
	FavouritesCount     types.Int64                           `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers           types.Int64                           `tfsdk:"followers" json:"followers,computed"`
	Following           types.Int64                           `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines  types.Bool                            `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	IsAutomated         types.Bool                            `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified      types.Bool                            `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsTranslator        types.Bool                            `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified          types.Bool                            `tfsdk:"is_verified" json:"isVerified,computed"`
	Location            types.String                          `tfsdk:"location" json:"location,computed"`
	MediaCount          types.Int64                           `tfsdk:"media_count" json:"mediaCount,computed"`
	PinnedTweetIDs      customfield.List[types.String]        `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive   types.Bool                            `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio          customfield.Map[jsontypes.Normalized] `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL    types.String                          `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfilePicture      types.String                          `tfsdk:"profile_picture" json:"profilePicture,computed"`
	Protected           types.Bool                            `tfsdk:"protected" json:"protected,computed"`
	StatusesCount       types.Int64                           `tfsdk:"statuses_count" json:"statusesCount,computed"`
	Unavailable         types.Bool                            `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason   types.String                          `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                 types.String                          `tfsdk:"url" json:"url,computed"`
	Verified            types.Bool                            `tfsdk:"verified" json:"verified,computed"`
	VerifiedType        types.String                          `tfsdk:"verified_type" json:"verifiedType,computed"`
	ViewerFollowedBy    types.Bool                            `tfsdk:"viewer_followed_by" json:"viewerFollowedBy,computed"`
	ViewerFollowing     types.Bool                            `tfsdk:"viewer_following" json:"viewerFollowing,computed"`
	WithheldInCountries customfield.List[types.String]        `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAIGeneratedDataSourceModel struct {
	CanEdit             types.Bool   `tfsdk:"can_edit" json:"canEdit,computed"`
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetRetweetedTweetMediaDataSourceModel struct {
	MediaURL      types.String                                                                             `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type          types.String                                                                             `tfsdk:"type" json:"type,computed"`
	URL           types.String                                                                             `tfsdk:"url" json:"url,computed"`
	VideoVariants customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaVideoVariantsDataSourceModel] `tfsdk:"video_variants" json:"videoVariants,computed"`
}

type XTweetTweetRetweetedTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}
