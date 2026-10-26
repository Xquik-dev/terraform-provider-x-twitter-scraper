// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtweet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XTweetDataSourceModel struct {
	ID     types.String                                          `tfsdk:"id" path:"id,required"`
	Author customfield.NestedObject[XTweetAuthorDataSourceModel] `tfsdk:"author" json:"author" api:"computed"`
	Tweet  customfield.NestedObject[XTweetTweetDataSourceModel]  `tfsdk:"tweet" json:"tweet" api:"computed"`
}

type XTweetAuthorDataSourceModel struct {
	ID                              types.String                                                                    `tfsdk:"id" json:"id" api:"computed"`
	Name                            types.String                                                                    `tfsdk:"name" json:"name" api:"computed"`
	Username                        types.String                                                                    `tfsdk:"username" json:"username" api:"computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetAuthorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel" api:"computed"`
	AutomatedBy                     types.String                                                                    `tfsdk:"automated_by" json:"automatedBy" api:"computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                     `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount" api:"computed"`
	CommunityRole                   types.String                                                                    `tfsdk:"community_role" json:"communityRole" api:"computed"`
	CoverPicture                    types.String                                                                    `tfsdk:"cover_picture" json:"coverPicture" api:"computed"`
	CreatedAt                       types.String                                                                    `tfsdk:"created_at" json:"createdAt" api:"computed"`
	CreatorSubscriptionsCount       types.Int64                                                                     `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount" api:"computed"`
	Description                     types.String                                                                    `tfsdk:"description" json:"description" api:"computed"`
	FavouritesCount                 types.Int64                                                                     `tfsdk:"favourites_count" json:"favouritesCount" api:"computed"`
	Followers                       types.Int64                                                                     `tfsdk:"followers" json:"followers" api:"computed"`
	Following                       types.Int64                                                                     `tfsdk:"following" json:"following" api:"computed"`
	HasCustomTimelines              types.Bool                                                                      `tfsdk:"has_custom_timelines" json:"hasCustomTimelines" api:"computed"`
	HasGraduatedAccess              types.Bool                                                                      `tfsdk:"has_graduated_access" json:"hasGraduatedAccess" api:"computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                      `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile" api:"computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetAuthorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo" api:"computed"`
	IdentityVerification            customfield.NestedObject[XTweetAuthorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification" api:"computed"`
	IsAutomated                     types.Bool                                                                      `tfsdk:"is_automated" json:"isAutomated" api:"computed"`
	IsBlueVerified                  types.Bool                                                                      `tfsdk:"is_blue_verified" json:"isBlueVerified" api:"computed"`
	IsProfileTranslatable           types.Bool                                                                      `tfsdk:"is_profile_translatable" json:"isProfileTranslatable" api:"computed"`
	IsTranslator                    types.Bool                                                                      `tfsdk:"is_translator" json:"isTranslator" api:"computed"`
	IsVerified                      types.Bool                                                                      `tfsdk:"is_verified" json:"isVerified" api:"computed"`
	Location                        types.String                                                                    `tfsdk:"location" json:"location" api:"computed"`
	MediaCount                      types.Int64                                                                     `tfsdk:"media_count" json:"mediaCount" api:"computed"`
	ParodyCommentaryFanLabel        types.String                                                                    `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel" api:"computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                  `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds" api:"computed"`
	PossiblySensitive               types.Bool                                                                      `tfsdk:"possibly_sensitive" json:"possiblySensitive" api:"computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                           `tfsdk:"profile_bio" json:"profile_bio" api:"computed"`
	ProfileBannerURL                types.String                                                                    `tfsdk:"profile_banner_url" json:"profileBannerUrl" api:"computed"`
	ProfileDescriptionLanguage      types.String                                                                    `tfsdk:"profile_description_language" json:"profileDescriptionLanguage" api:"computed"`
	ProfileImageShape               types.String                                                                    `tfsdk:"profile_image_shape" json:"profileImageShape" api:"computed"`
	ProfileInterstitialType         types.String                                                                    `tfsdk:"profile_interstitial_type" json:"profileInterstitialType" api:"computed"`
	ProfilePicture                  types.String                                                                    `tfsdk:"profile_picture" json:"profilePicture" api:"computed"`
	ProfileSortEnabled              types.Bool                                                                      `tfsdk:"profile_sort_enabled" json:"profileSortEnabled" api:"computed"`
	ProfileTranslatorType           types.String                                                                    `tfsdk:"profile_translator_type" json:"profileTranslatorType" api:"computed"`
	Protected                       types.Bool                                                                      `tfsdk:"protected" json:"protected" api:"computed"`
	StatusesCount                   types.Int64                                                                     `tfsdk:"statuses_count" json:"statusesCount" api:"computed"`
	SuperFollowEligible             types.Bool                                                                      `tfsdk:"super_follow_eligible" json:"superFollowEligible" api:"computed"`
	Unavailable                     types.Bool                                                                      `tfsdk:"unavailable" json:"unavailable" api:"computed"`
	UnavailableReason               types.String                                                                    `tfsdk:"unavailable_reason" json:"unavailableReason" api:"computed"`
	URL                             types.String                                                                    `tfsdk:"url" json:"url" api:"computed"`
	Verified                        types.Bool                                                                      `tfsdk:"verified" json:"verified" api:"computed"`
	VerifiedType                    types.String                                                                    `tfsdk:"verified_type" json:"verifiedType" api:"computed"`
	WithheldInCountries             customfield.List[types.String]                                                  `tfsdk:"withheld_in_countries" json:"withheldInCountries" api:"computed"`
}

type XTweetAuthorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl" api:"computed"`
	Description          types.String `tfsdk:"description" json:"description" api:"computed"`
	URL                  types.String `tfsdk:"url" json:"url" api:"computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType" api:"computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType" api:"computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType" api:"computed"`
}

type XTweetAuthorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets" api:"computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets" api:"computed"`
}

type XTweetAuthorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description" api:"computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified" api:"computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec" api:"computed"`
}

type XTweetTweetDataSourceModel struct {
	ID                types.String                                                          `tfsdk:"id" json:"id" api:"computed"`
	BookmarkCount     types.Int64                                                           `tfsdk:"bookmark_count" json:"bookmarkCount" api:"computed"`
	LikeCount         types.Int64                                                           `tfsdk:"like_count" json:"likeCount" api:"computed"`
	QuoteCount        types.Int64                                                           `tfsdk:"quote_count" json:"quoteCount" api:"computed"`
	ReplyCount        types.Int64                                                           `tfsdk:"reply_count" json:"replyCount" api:"computed"`
	RetweetCount      types.Int64                                                           `tfsdk:"retweet_count" json:"retweetCount" api:"computed"`
	Text              types.String                                                          `tfsdk:"text" json:"text" api:"computed"`
	ViewCount         types.Int64                                                           `tfsdk:"view_count" json:"viewCount" api:"computed"`
	Article           customfield.NestedObject[XTweetTweetArticleDataSourceModel]           `tfsdk:"article" json:"article" api:"computed"`
	Author            customfield.NestedObject[XTweetTweetAuthorDataSourceModel]            `tfsdk:"author" json:"author" api:"computed"`
	Card              customfield.NestedObject[XTweetTweetCardDataSourceModel]              `tfsdk:"card" json:"card" api:"computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetCommunityNoteDataSourceModel]     `tfsdk:"community_note" json:"communityNote" api:"computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure" api:"computed"`
	ConversationID    types.String                                                          `tfsdk:"conversation_id" json:"conversationId" api:"computed"`
	CreatedAt         types.String                                                          `tfsdk:"created_at" json:"createdAt" api:"computed"`
	DisplayTextRange  customfield.List[types.Int64]                                         `tfsdk:"display_text_range" json:"displayTextRange" api:"computed"`
	Edit              customfield.NestedObject[XTweetTweetEditDataSourceModel]              `tfsdk:"edit" json:"edit" api:"computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                 `tfsdk:"entities" json:"entities" api:"computed"`
	InReplyToID       types.String                                                          `tfsdk:"in_reply_to_id" json:"inReplyToId" api:"computed"`
	InReplyToUserID   types.String                                                          `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId" api:"computed"`
	InReplyToUsername types.String                                                          `tfsdk:"in_reply_to_username" json:"inReplyToUsername" api:"computed"`
	IsLimitedReply    types.Bool                                                            `tfsdk:"is_limited_reply" json:"isLimitedReply" api:"computed"`
	IsNoteTweet       types.Bool                                                            `tfsdk:"is_note_tweet" json:"isNoteTweet" api:"computed"`
	IsQuoteStatus     types.Bool                                                            `tfsdk:"is_quote_status" json:"isQuoteStatus" api:"computed"`
	IsReply           types.Bool                                                            `tfsdk:"is_reply" json:"isReply" api:"computed"`
	IsTranslatable    types.Bool                                                            `tfsdk:"is_translatable" json:"isTranslatable" api:"computed"`
	Lang              types.String                                                          `tfsdk:"lang" json:"lang" api:"computed"`
	Media             customfield.NestedObjectList[XTweetTweetMediaDataSourceModel]         `tfsdk:"media" json:"media" api:"computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetNoteTweetDataSourceModel]         `tfsdk:"note_tweet" json:"noteTweet" api:"computed"`
	Place             customfield.NestedObject[XTweetTweetPlaceDataSourceModel]             `tfsdk:"place" json:"place" api:"computed"`
	PossiblySensitive types.Bool                                                            `tfsdk:"possibly_sensitive" json:"possiblySensitive" api:"computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetPreviousCountsDataSourceModel]    `tfsdk:"previous_counts" json:"previousCounts" api:"computed"`
	QuotedTweet       customfield.NestedObject[XTweetTweetQuotedTweetDataSourceModel]       `tfsdk:"quoted_tweet" json:"quoted_tweet" api:"computed"`
	RetweetedTweet    customfield.NestedObject[XTweetTweetRetweetedTweetDataSourceModel]    `tfsdk:"retweeted_tweet" json:"retweeted_tweet" api:"computed"`
	Source            types.String                                                          `tfsdk:"source" json:"source" api:"computed"`
	Type              types.String                                                          `tfsdk:"type" json:"type" api:"computed"`
	URL               types.String                                                          `tfsdk:"url" json:"url" api:"computed"`
	ViewState         types.String                                                          `tfsdk:"view_state" json:"viewState" api:"computed"`
}

type XTweetTweetArticleDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id" api:"computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl" api:"computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText" api:"computed"`
	Title         types.String `tfsdk:"title" json:"title" api:"computed"`
}

type XTweetTweetAuthorDataSourceModel struct {
	ID                              types.String                                                                         `tfsdk:"id" json:"id" api:"computed"`
	Name                            types.String                                                                         `tfsdk:"name" json:"name" api:"computed"`
	Username                        types.String                                                                         `tfsdk:"username" json:"username" api:"computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetAuthorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel" api:"computed"`
	AutomatedBy                     types.String                                                                         `tfsdk:"automated_by" json:"automatedBy" api:"computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                          `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount" api:"computed"`
	CommunityRole                   types.String                                                                         `tfsdk:"community_role" json:"communityRole" api:"computed"`
	CoverPicture                    types.String                                                                         `tfsdk:"cover_picture" json:"coverPicture" api:"computed"`
	CreatedAt                       types.String                                                                         `tfsdk:"created_at" json:"createdAt" api:"computed"`
	CreatorSubscriptionsCount       types.Int64                                                                          `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount" api:"computed"`
	Description                     types.String                                                                         `tfsdk:"description" json:"description" api:"computed"`
	FavouritesCount                 types.Int64                                                                          `tfsdk:"favourites_count" json:"favouritesCount" api:"computed"`
	Followers                       types.Int64                                                                          `tfsdk:"followers" json:"followers" api:"computed"`
	Following                       types.Int64                                                                          `tfsdk:"following" json:"following" api:"computed"`
	HasCustomTimelines              types.Bool                                                                           `tfsdk:"has_custom_timelines" json:"hasCustomTimelines" api:"computed"`
	HasGraduatedAccess              types.Bool                                                                           `tfsdk:"has_graduated_access" json:"hasGraduatedAccess" api:"computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                           `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile" api:"computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetAuthorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo" api:"computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetAuthorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification" api:"computed"`
	IsAutomated                     types.Bool                                                                           `tfsdk:"is_automated" json:"isAutomated" api:"computed"`
	IsBlueVerified                  types.Bool                                                                           `tfsdk:"is_blue_verified" json:"isBlueVerified" api:"computed"`
	IsProfileTranslatable           types.Bool                                                                           `tfsdk:"is_profile_translatable" json:"isProfileTranslatable" api:"computed"`
	IsTranslator                    types.Bool                                                                           `tfsdk:"is_translator" json:"isTranslator" api:"computed"`
	IsVerified                      types.Bool                                                                           `tfsdk:"is_verified" json:"isVerified" api:"computed"`
	Location                        types.String                                                                         `tfsdk:"location" json:"location" api:"computed"`
	MediaCount                      types.Int64                                                                          `tfsdk:"media_count" json:"mediaCount" api:"computed"`
	ParodyCommentaryFanLabel        types.String                                                                         `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel" api:"computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                       `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds" api:"computed"`
	PossiblySensitive               types.Bool                                                                           `tfsdk:"possibly_sensitive" json:"possiblySensitive" api:"computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                `tfsdk:"profile_bio" json:"profile_bio" api:"computed"`
	ProfileBannerURL                types.String                                                                         `tfsdk:"profile_banner_url" json:"profileBannerUrl" api:"computed"`
	ProfileDescriptionLanguage      types.String                                                                         `tfsdk:"profile_description_language" json:"profileDescriptionLanguage" api:"computed"`
	ProfileImageShape               types.String                                                                         `tfsdk:"profile_image_shape" json:"profileImageShape" api:"computed"`
	ProfileInterstitialType         types.String                                                                         `tfsdk:"profile_interstitial_type" json:"profileInterstitialType" api:"computed"`
	ProfilePicture                  types.String                                                                         `tfsdk:"profile_picture" json:"profilePicture" api:"computed"`
	ProfileSortEnabled              types.Bool                                                                           `tfsdk:"profile_sort_enabled" json:"profileSortEnabled" api:"computed"`
	ProfileTranslatorType           types.String                                                                         `tfsdk:"profile_translator_type" json:"profileTranslatorType" api:"computed"`
	Protected                       types.Bool                                                                           `tfsdk:"protected" json:"protected" api:"computed"`
	StatusesCount                   types.Int64                                                                          `tfsdk:"statuses_count" json:"statusesCount" api:"computed"`
	SuperFollowEligible             types.Bool                                                                           `tfsdk:"super_follow_eligible" json:"superFollowEligible" api:"computed"`
	Unavailable                     types.Bool                                                                           `tfsdk:"unavailable" json:"unavailable" api:"computed"`
	UnavailableReason               types.String                                                                         `tfsdk:"unavailable_reason" json:"unavailableReason" api:"computed"`
	URL                             types.String                                                                         `tfsdk:"url" json:"url" api:"computed"`
	Verified                        types.Bool                                                                           `tfsdk:"verified" json:"verified" api:"computed"`
	VerifiedType                    types.String                                                                         `tfsdk:"verified_type" json:"verifiedType" api:"computed"`
	WithheldInCountries             customfield.List[types.String]                                                       `tfsdk:"withheld_in_countries" json:"withheldInCountries" api:"computed"`
}

type XTweetTweetAuthorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl" api:"computed"`
	Description          types.String `tfsdk:"description" json:"description" api:"computed"`
	URL                  types.String `tfsdk:"url" json:"url" api:"computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType" api:"computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType" api:"computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType" api:"computed"`
}

type XTweetTweetAuthorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets" api:"computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets" api:"computed"`
}

type XTweetTweetAuthorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description" api:"computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified" api:"computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec" api:"computed"`
}

type XTweetTweetCardDataSourceModel struct {
	ID            types.String                          `tfsdk:"id" json:"id" api:"computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues" api:"computed"`
	Name          types.String                          `tfsdk:"name" json:"name" api:"computed"`
	URL           types.String                          `tfsdk:"url" json:"url" api:"computed"`
}

type XTweetTweetCommunityNoteDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id" api:"computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl" api:"computed"`
	Footer         types.String `tfsdk:"footer" json:"footer" api:"computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle" api:"computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle" api:"computed"`
	Title          types.String `tfsdk:"title" json:"title" api:"computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle" api:"computed"`
}

type XTweetTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising" api:"computed"`
	AIGenerated customfield.NestedObject[XTweetTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated" api:"computed"`
}

type XTweetTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion" api:"computed"`
}

type XTweetTweetContentDisclosureAIGeneratedDataSourceModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource" api:"computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia" api:"computed"`
}

type XTweetTweetEditDataSourceModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs" api:"computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds" api:"computed"`
}

type XTweetTweetMediaDataSourceModel struct {
	MediaURL           types.String                                                                            `tfsdk:"media_url" json:"mediaUrl" api:"computed"`
	Type               types.String                                                                            `tfsdk:"type" json:"type" api:"computed"`
	URL                types.String                                                                            `tfsdk:"url" json:"url" api:"computed"`
	ID                 types.String                                                                            `tfsdk:"id" json:"id" api:"computed"`
	AllowDownload      types.Bool                                                                              `tfsdk:"allow_download" json:"allowDownload" api:"computed"`
	AltText            types.String                                                                            `tfsdk:"alt_text" json:"altText" api:"computed"`
	AspectRatio        customfield.List[types.Int64]                                                           `tfsdk:"aspect_ratio" json:"aspectRatio" api:"computed"`
	AvailabilityStatus types.String                                                                            `tfsdk:"availability_status" json:"availabilityStatus" api:"computed"`
	DisplayURL         types.String                                                                            `tfsdk:"display_url" json:"displayUrl" api:"computed"`
	DurationMillis     types.Int64                                                                             `tfsdk:"duration_millis" json:"durationMillis" api:"computed"`
	ExpandedURL        types.String                                                                            `tfsdk:"expanded_url" json:"expandedUrl" api:"computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetMediaFaceRectsDataSourceModel]] `tfsdk:"face_rects" json:"faceRects" api:"computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetMediaFocusRectsDataSourceModel]                 `tfsdk:"focus_rects" json:"focusRects" api:"computed"`
	Height             types.Int64                                                                             `tfsdk:"height" json:"height" api:"computed"`
	Indices            customfield.List[types.Int64]                                                           `tfsdk:"indices" json:"indices" api:"computed"`
	MediaKey           types.String                                                                            `tfsdk:"media_key" json:"mediaKey" api:"computed"`
	Monetizable        types.Bool                                                                              `tfsdk:"monetizable" json:"monetizable" api:"computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetMediaSizesDataSourceModel]                       `tfsdk:"sizes" json:"sizes" api:"computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetMediaVideoVariantsDataSourceModel]              `tfsdk:"video_variants" json:"videoVariants" api:"computed"`
	Width              types.Int64                                                                             `tfsdk:"width" json:"width" api:"computed"`
}

type XTweetTweetMediaFaceRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h" api:"required"`
	W types.Int64 `tfsdk:"w" json:"w" api:"required"`
	X types.Int64 `tfsdk:"x" json:"x" api:"required"`
	Y types.Int64 `tfsdk:"y" json:"y" api:"required"`
}

type XTweetTweetMediaFocusRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h" api:"computed"`
	W types.Int64 `tfsdk:"w" json:"w" api:"computed"`
	X types.Int64 `tfsdk:"x" json:"x" api:"computed"`
	Y types.Int64 `tfsdk:"y" json:"y" api:"computed"`
}

type XTweetTweetMediaSizesDataSourceModel struct {
	H      types.Int64  `tfsdk:"h" json:"h" api:"computed"`
	Resize types.String `tfsdk:"resize" json:"resize" api:"computed"`
	W      types.Int64  `tfsdk:"w" json:"w" api:"computed"`
}

type XTweetTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType" api:"computed"`
	URL         types.String `tfsdk:"url" json:"url" api:"computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate" api:"computed"`
}

type XTweetTweetNoteTweetDataSourceModel struct {
	Text         types.String                                                                  `tfsdk:"text" json:"text" api:"computed"`
	ID           types.String                                                                  `tfsdk:"id" json:"id" api:"computed"`
	Entities     customfield.Map[jsontypes.Normalized]                                         `tfsdk:"entities" json:"entities" api:"computed"`
	IsExpandable types.Bool                                                                    `tfsdk:"is_expandable" json:"isExpandable" api:"computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetNoteTweetRichtextTagsDataSourceModel] `tfsdk:"richtext_tags" json:"richtextTags" api:"computed"`
}

type XTweetTweetNoteTweetRichtextTagsDataSourceModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex" api:"computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex" api:"computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types" api:"computed"`
}

type XTweetTweetPlaceDataSourceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id" api:"computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox" api:"computed"`
	Country     types.String                          `tfsdk:"country" json:"country" api:"computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode" api:"computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName" api:"computed"`
	Name        types.String                          `tfsdk:"name" json:"name" api:"computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType" api:"computed"`
	URL         types.String                          `tfsdk:"url" json:"url" api:"computed"`
}

type XTweetTweetPreviousCountsDataSourceModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount" api:"computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount" api:"computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount" api:"computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount" api:"computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount" api:"computed"`
}

type XTweetTweetQuotedTweetDataSourceModel struct {
	ID                types.String                                                                     `tfsdk:"id" json:"id" api:"computed"`
	BookmarkCount     types.Int64                                                                      `tfsdk:"bookmark_count" json:"bookmarkCount" api:"computed"`
	LikeCount         types.Int64                                                                      `tfsdk:"like_count" json:"likeCount" api:"computed"`
	QuoteCount        types.Int64                                                                      `tfsdk:"quote_count" json:"quoteCount" api:"computed"`
	ReplyCount        types.Int64                                                                      `tfsdk:"reply_count" json:"replyCount" api:"computed"`
	RetweetCount      types.Int64                                                                      `tfsdk:"retweet_count" json:"retweetCount" api:"computed"`
	Text              types.String                                                                     `tfsdk:"text" json:"text" api:"computed"`
	ViewCount         types.Int64                                                                      `tfsdk:"view_count" json:"viewCount" api:"computed"`
	Article           customfield.NestedObject[XTweetTweetQuotedTweetArticleDataSourceModel]           `tfsdk:"article" json:"article" api:"computed"`
	Author            customfield.NestedObject[XTweetTweetQuotedTweetAuthorDataSourceModel]            `tfsdk:"author" json:"author" api:"computed"`
	Card              customfield.NestedObject[XTweetTweetQuotedTweetCardDataSourceModel]              `tfsdk:"card" json:"card" api:"computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetQuotedTweetCommunityNoteDataSourceModel]     `tfsdk:"community_note" json:"communityNote" api:"computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure" api:"computed"`
	ConversationID    types.String                                                                     `tfsdk:"conversation_id" json:"conversationId" api:"computed"`
	CreatedAt         types.String                                                                     `tfsdk:"created_at" json:"createdAt" api:"computed"`
	DisplayTextRange  customfield.List[types.Int64]                                                    `tfsdk:"display_text_range" json:"displayTextRange" api:"computed"`
	Edit              customfield.NestedObject[XTweetTweetQuotedTweetEditDataSourceModel]              `tfsdk:"edit" json:"edit" api:"computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                            `tfsdk:"entities" json:"entities" api:"computed"`
	InReplyToID       types.String                                                                     `tfsdk:"in_reply_to_id" json:"inReplyToId" api:"computed"`
	InReplyToUserID   types.String                                                                     `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId" api:"computed"`
	InReplyToUsername types.String                                                                     `tfsdk:"in_reply_to_username" json:"inReplyToUsername" api:"computed"`
	IsLimitedReply    types.Bool                                                                       `tfsdk:"is_limited_reply" json:"isLimitedReply" api:"computed"`
	IsNoteTweet       types.Bool                                                                       `tfsdk:"is_note_tweet" json:"isNoteTweet" api:"computed"`
	IsQuoteStatus     types.Bool                                                                       `tfsdk:"is_quote_status" json:"isQuoteStatus" api:"computed"`
	IsReply           types.Bool                                                                       `tfsdk:"is_reply" json:"isReply" api:"computed"`
	IsTranslatable    types.Bool                                                                       `tfsdk:"is_translatable" json:"isTranslatable" api:"computed"`
	Lang              types.String                                                                     `tfsdk:"lang" json:"lang" api:"computed"`
	Media             customfield.NestedObjectList[XTweetTweetQuotedTweetMediaDataSourceModel]         `tfsdk:"media" json:"media" api:"computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetQuotedTweetNoteTweetDataSourceModel]         `tfsdk:"note_tweet" json:"noteTweet" api:"computed"`
	Place             customfield.NestedObject[XTweetTweetQuotedTweetPlaceDataSourceModel]             `tfsdk:"place" json:"place" api:"computed"`
	PossiblySensitive types.Bool                                                                       `tfsdk:"possibly_sensitive" json:"possiblySensitive" api:"computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetQuotedTweetPreviousCountsDataSourceModel]    `tfsdk:"previous_counts" json:"previousCounts" api:"computed"`
	QuotedTweet       customfield.NormalizedDynamicValue                                               `tfsdk:"quoted_tweet" json:"quoted_tweet" api:"computed"`
	RetweetedTweet    customfield.NormalizedDynamicValue                                               `tfsdk:"retweeted_tweet" json:"retweeted_tweet" api:"computed"`
	Source            types.String                                                                     `tfsdk:"source" json:"source" api:"computed"`
	Type              types.String                                                                     `tfsdk:"type" json:"type" api:"computed"`
	URL               types.String                                                                     `tfsdk:"url" json:"url" api:"computed"`
	ViewState         types.String                                                                     `tfsdk:"view_state" json:"viewState" api:"computed"`
}

type XTweetTweetQuotedTweetArticleDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id" api:"computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl" api:"computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText" api:"computed"`
	Title         types.String `tfsdk:"title" json:"title" api:"computed"`
}

type XTweetTweetQuotedTweetAuthorDataSourceModel struct {
	ID                              types.String                                                                                    `tfsdk:"id" json:"id" api:"computed"`
	Name                            types.String                                                                                    `tfsdk:"name" json:"name" api:"computed"`
	Username                        types.String                                                                                    `tfsdk:"username" json:"username" api:"computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetQuotedTweetAuthorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel" api:"computed"`
	AutomatedBy                     types.String                                                                                    `tfsdk:"automated_by" json:"automatedBy" api:"computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                                     `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount" api:"computed"`
	CommunityRole                   types.String                                                                                    `tfsdk:"community_role" json:"communityRole" api:"computed"`
	CoverPicture                    types.String                                                                                    `tfsdk:"cover_picture" json:"coverPicture" api:"computed"`
	CreatedAt                       types.String                                                                                    `tfsdk:"created_at" json:"createdAt" api:"computed"`
	CreatorSubscriptionsCount       types.Int64                                                                                     `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount" api:"computed"`
	Description                     types.String                                                                                    `tfsdk:"description" json:"description" api:"computed"`
	FavouritesCount                 types.Int64                                                                                     `tfsdk:"favourites_count" json:"favouritesCount" api:"computed"`
	Followers                       types.Int64                                                                                     `tfsdk:"followers" json:"followers" api:"computed"`
	Following                       types.Int64                                                                                     `tfsdk:"following" json:"following" api:"computed"`
	HasCustomTimelines              types.Bool                                                                                      `tfsdk:"has_custom_timelines" json:"hasCustomTimelines" api:"computed"`
	HasGraduatedAccess              types.Bool                                                                                      `tfsdk:"has_graduated_access" json:"hasGraduatedAccess" api:"computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                                      `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile" api:"computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetQuotedTweetAuthorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo" api:"computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetQuotedTweetAuthorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification" api:"computed"`
	IsAutomated                     types.Bool                                                                                      `tfsdk:"is_automated" json:"isAutomated" api:"computed"`
	IsBlueVerified                  types.Bool                                                                                      `tfsdk:"is_blue_verified" json:"isBlueVerified" api:"computed"`
	IsProfileTranslatable           types.Bool                                                                                      `tfsdk:"is_profile_translatable" json:"isProfileTranslatable" api:"computed"`
	IsTranslator                    types.Bool                                                                                      `tfsdk:"is_translator" json:"isTranslator" api:"computed"`
	IsVerified                      types.Bool                                                                                      `tfsdk:"is_verified" json:"isVerified" api:"computed"`
	Location                        types.String                                                                                    `tfsdk:"location" json:"location" api:"computed"`
	MediaCount                      types.Int64                                                                                     `tfsdk:"media_count" json:"mediaCount" api:"computed"`
	ParodyCommentaryFanLabel        types.String                                                                                    `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel" api:"computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                                  `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds" api:"computed"`
	PossiblySensitive               types.Bool                                                                                      `tfsdk:"possibly_sensitive" json:"possiblySensitive" api:"computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                           `tfsdk:"profile_bio" json:"profile_bio" api:"computed"`
	ProfileBannerURL                types.String                                                                                    `tfsdk:"profile_banner_url" json:"profileBannerUrl" api:"computed"`
	ProfileDescriptionLanguage      types.String                                                                                    `tfsdk:"profile_description_language" json:"profileDescriptionLanguage" api:"computed"`
	ProfileImageShape               types.String                                                                                    `tfsdk:"profile_image_shape" json:"profileImageShape" api:"computed"`
	ProfileInterstitialType         types.String                                                                                    `tfsdk:"profile_interstitial_type" json:"profileInterstitialType" api:"computed"`
	ProfilePicture                  types.String                                                                                    `tfsdk:"profile_picture" json:"profilePicture" api:"computed"`
	ProfileSortEnabled              types.Bool                                                                                      `tfsdk:"profile_sort_enabled" json:"profileSortEnabled" api:"computed"`
	ProfileTranslatorType           types.String                                                                                    `tfsdk:"profile_translator_type" json:"profileTranslatorType" api:"computed"`
	Protected                       types.Bool                                                                                      `tfsdk:"protected" json:"protected" api:"computed"`
	StatusesCount                   types.Int64                                                                                     `tfsdk:"statuses_count" json:"statusesCount" api:"computed"`
	SuperFollowEligible             types.Bool                                                                                      `tfsdk:"super_follow_eligible" json:"superFollowEligible" api:"computed"`
	Unavailable                     types.Bool                                                                                      `tfsdk:"unavailable" json:"unavailable" api:"computed"`
	UnavailableReason               types.String                                                                                    `tfsdk:"unavailable_reason" json:"unavailableReason" api:"computed"`
	URL                             types.String                                                                                    `tfsdk:"url" json:"url" api:"computed"`
	Verified                        types.Bool                                                                                      `tfsdk:"verified" json:"verified" api:"computed"`
	VerifiedType                    types.String                                                                                    `tfsdk:"verified_type" json:"verifiedType" api:"computed"`
	WithheldInCountries             customfield.List[types.String]                                                                  `tfsdk:"withheld_in_countries" json:"withheldInCountries" api:"computed"`
}

type XTweetTweetQuotedTweetAuthorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl" api:"computed"`
	Description          types.String `tfsdk:"description" json:"description" api:"computed"`
	URL                  types.String `tfsdk:"url" json:"url" api:"computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType" api:"computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType" api:"computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType" api:"computed"`
}

type XTweetTweetQuotedTweetAuthorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets" api:"computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets" api:"computed"`
}

type XTweetTweetQuotedTweetAuthorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description" api:"computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified" api:"computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec" api:"computed"`
}

type XTweetTweetQuotedTweetCardDataSourceModel struct {
	ID            types.String                          `tfsdk:"id" json:"id" api:"computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues" api:"computed"`
	Name          types.String                          `tfsdk:"name" json:"name" api:"computed"`
	URL           types.String                          `tfsdk:"url" json:"url" api:"computed"`
}

type XTweetTweetQuotedTweetCommunityNoteDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id" api:"computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl" api:"computed"`
	Footer         types.String `tfsdk:"footer" json:"footer" api:"computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle" api:"computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle" api:"computed"`
	Title          types.String `tfsdk:"title" json:"title" api:"computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle" api:"computed"`
}

type XTweetTweetQuotedTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising" api:"computed"`
	AIGenerated customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated" api:"computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion" api:"computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAIGeneratedDataSourceModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource" api:"computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia" api:"computed"`
}

type XTweetTweetQuotedTweetEditDataSourceModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs" api:"computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds" api:"computed"`
}

type XTweetTweetQuotedTweetMediaDataSourceModel struct {
	MediaURL           types.String                                                                                       `tfsdk:"media_url" json:"mediaUrl" api:"computed"`
	Type               types.String                                                                                       `tfsdk:"type" json:"type" api:"computed"`
	URL                types.String                                                                                       `tfsdk:"url" json:"url" api:"computed"`
	ID                 types.String                                                                                       `tfsdk:"id" json:"id" api:"computed"`
	AllowDownload      types.Bool                                                                                         `tfsdk:"allow_download" json:"allowDownload" api:"computed"`
	AltText            types.String                                                                                       `tfsdk:"alt_text" json:"altText" api:"computed"`
	AspectRatio        customfield.List[types.Int64]                                                                      `tfsdk:"aspect_ratio" json:"aspectRatio" api:"computed"`
	AvailabilityStatus types.String                                                                                       `tfsdk:"availability_status" json:"availabilityStatus" api:"computed"`
	DisplayURL         types.String                                                                                       `tfsdk:"display_url" json:"displayUrl" api:"computed"`
	DurationMillis     types.Int64                                                                                        `tfsdk:"duration_millis" json:"durationMillis" api:"computed"`
	ExpandedURL        types.String                                                                                       `tfsdk:"expanded_url" json:"expandedUrl" api:"computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetQuotedTweetMediaFaceRectsDataSourceModel]] `tfsdk:"face_rects" json:"faceRects" api:"computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetQuotedTweetMediaFocusRectsDataSourceModel]                 `tfsdk:"focus_rects" json:"focusRects" api:"computed"`
	Height             types.Int64                                                                                        `tfsdk:"height" json:"height" api:"computed"`
	Indices            customfield.List[types.Int64]                                                                      `tfsdk:"indices" json:"indices" api:"computed"`
	MediaKey           types.String                                                                                       `tfsdk:"media_key" json:"mediaKey" api:"computed"`
	Monetizable        types.Bool                                                                                         `tfsdk:"monetizable" json:"monetizable" api:"computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetQuotedTweetMediaSizesDataSourceModel]                       `tfsdk:"sizes" json:"sizes" api:"computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetQuotedTweetMediaVideoVariantsDataSourceModel]              `tfsdk:"video_variants" json:"videoVariants" api:"computed"`
	Width              types.Int64                                                                                        `tfsdk:"width" json:"width" api:"computed"`
}

type XTweetTweetQuotedTweetMediaFaceRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h" api:"required"`
	W types.Int64 `tfsdk:"w" json:"w" api:"required"`
	X types.Int64 `tfsdk:"x" json:"x" api:"required"`
	Y types.Int64 `tfsdk:"y" json:"y" api:"required"`
}

type XTweetTweetQuotedTweetMediaFocusRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h" api:"computed"`
	W types.Int64 `tfsdk:"w" json:"w" api:"computed"`
	X types.Int64 `tfsdk:"x" json:"x" api:"computed"`
	Y types.Int64 `tfsdk:"y" json:"y" api:"computed"`
}

type XTweetTweetQuotedTweetMediaSizesDataSourceModel struct {
	H      types.Int64  `tfsdk:"h" json:"h" api:"computed"`
	Resize types.String `tfsdk:"resize" json:"resize" api:"computed"`
	W      types.Int64  `tfsdk:"w" json:"w" api:"computed"`
}

type XTweetTweetQuotedTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType" api:"computed"`
	URL         types.String `tfsdk:"url" json:"url" api:"computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate" api:"computed"`
}

type XTweetTweetQuotedTweetNoteTweetDataSourceModel struct {
	Text         types.String                                                                             `tfsdk:"text" json:"text" api:"computed"`
	ID           types.String                                                                             `tfsdk:"id" json:"id" api:"computed"`
	Entities     customfield.Map[jsontypes.Normalized]                                                    `tfsdk:"entities" json:"entities" api:"computed"`
	IsExpandable types.Bool                                                                               `tfsdk:"is_expandable" json:"isExpandable" api:"computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetQuotedTweetNoteTweetRichtextTagsDataSourceModel] `tfsdk:"richtext_tags" json:"richtextTags" api:"computed"`
}

type XTweetTweetQuotedTweetNoteTweetRichtextTagsDataSourceModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex" api:"computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex" api:"computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types" api:"computed"`
}

type XTweetTweetQuotedTweetPlaceDataSourceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id" api:"computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox" api:"computed"`
	Country     types.String                          `tfsdk:"country" json:"country" api:"computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode" api:"computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName" api:"computed"`
	Name        types.String                          `tfsdk:"name" json:"name" api:"computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType" api:"computed"`
	URL         types.String                          `tfsdk:"url" json:"url" api:"computed"`
}

type XTweetTweetQuotedTweetPreviousCountsDataSourceModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount" api:"computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount" api:"computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount" api:"computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount" api:"computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount" api:"computed"`
}

type XTweetTweetRetweetedTweetDataSourceModel struct {
	ID                types.String                                                                        `tfsdk:"id" json:"id" api:"computed"`
	BookmarkCount     types.Int64                                                                         `tfsdk:"bookmark_count" json:"bookmarkCount" api:"computed"`
	LikeCount         types.Int64                                                                         `tfsdk:"like_count" json:"likeCount" api:"computed"`
	QuoteCount        types.Int64                                                                         `tfsdk:"quote_count" json:"quoteCount" api:"computed"`
	ReplyCount        types.Int64                                                                         `tfsdk:"reply_count" json:"replyCount" api:"computed"`
	RetweetCount      types.Int64                                                                         `tfsdk:"retweet_count" json:"retweetCount" api:"computed"`
	Text              types.String                                                                        `tfsdk:"text" json:"text" api:"computed"`
	ViewCount         types.Int64                                                                         `tfsdk:"view_count" json:"viewCount" api:"computed"`
	Article           customfield.NestedObject[XTweetTweetRetweetedTweetArticleDataSourceModel]           `tfsdk:"article" json:"article" api:"computed"`
	Author            customfield.NestedObject[XTweetTweetRetweetedTweetAuthorDataSourceModel]            `tfsdk:"author" json:"author" api:"computed"`
	Card              customfield.NestedObject[XTweetTweetRetweetedTweetCardDataSourceModel]              `tfsdk:"card" json:"card" api:"computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetRetweetedTweetCommunityNoteDataSourceModel]     `tfsdk:"community_note" json:"communityNote" api:"computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure" api:"computed"`
	ConversationID    types.String                                                                        `tfsdk:"conversation_id" json:"conversationId" api:"computed"`
	CreatedAt         types.String                                                                        `tfsdk:"created_at" json:"createdAt" api:"computed"`
	DisplayTextRange  customfield.List[types.Int64]                                                       `tfsdk:"display_text_range" json:"displayTextRange" api:"computed"`
	Edit              customfield.NestedObject[XTweetTweetRetweetedTweetEditDataSourceModel]              `tfsdk:"edit" json:"edit" api:"computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                               `tfsdk:"entities" json:"entities" api:"computed"`
	InReplyToID       types.String                                                                        `tfsdk:"in_reply_to_id" json:"inReplyToId" api:"computed"`
	InReplyToUserID   types.String                                                                        `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId" api:"computed"`
	InReplyToUsername types.String                                                                        `tfsdk:"in_reply_to_username" json:"inReplyToUsername" api:"computed"`
	IsLimitedReply    types.Bool                                                                          `tfsdk:"is_limited_reply" json:"isLimitedReply" api:"computed"`
	IsNoteTweet       types.Bool                                                                          `tfsdk:"is_note_tweet" json:"isNoteTweet" api:"computed"`
	IsQuoteStatus     types.Bool                                                                          `tfsdk:"is_quote_status" json:"isQuoteStatus" api:"computed"`
	IsReply           types.Bool                                                                          `tfsdk:"is_reply" json:"isReply" api:"computed"`
	IsTranslatable    types.Bool                                                                          `tfsdk:"is_translatable" json:"isTranslatable" api:"computed"`
	Lang              types.String                                                                        `tfsdk:"lang" json:"lang" api:"computed"`
	Media             customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaDataSourceModel]         `tfsdk:"media" json:"media" api:"computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetRetweetedTweetNoteTweetDataSourceModel]         `tfsdk:"note_tweet" json:"noteTweet" api:"computed"`
	Place             customfield.NestedObject[XTweetTweetRetweetedTweetPlaceDataSourceModel]             `tfsdk:"place" json:"place" api:"computed"`
	PossiblySensitive types.Bool                                                                          `tfsdk:"possibly_sensitive" json:"possiblySensitive" api:"computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetRetweetedTweetPreviousCountsDataSourceModel]    `tfsdk:"previous_counts" json:"previousCounts" api:"computed"`
	QuotedTweet       customfield.NormalizedDynamicValue                                                  `tfsdk:"quoted_tweet" json:"quoted_tweet" api:"computed"`
	RetweetedTweet    customfield.NormalizedDynamicValue                                                  `tfsdk:"retweeted_tweet" json:"retweeted_tweet" api:"computed"`
	Source            types.String                                                                        `tfsdk:"source" json:"source" api:"computed"`
	Type              types.String                                                                        `tfsdk:"type" json:"type" api:"computed"`
	URL               types.String                                                                        `tfsdk:"url" json:"url" api:"computed"`
	ViewState         types.String                                                                        `tfsdk:"view_state" json:"viewState" api:"computed"`
}

type XTweetTweetRetweetedTweetArticleDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id" api:"computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl" api:"computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText" api:"computed"`
	Title         types.String `tfsdk:"title" json:"title" api:"computed"`
}

type XTweetTweetRetweetedTweetAuthorDataSourceModel struct {
	ID                              types.String                                                                                       `tfsdk:"id" json:"id" api:"computed"`
	Name                            types.String                                                                                       `tfsdk:"name" json:"name" api:"computed"`
	Username                        types.String                                                                                       `tfsdk:"username" json:"username" api:"computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetRetweetedTweetAuthorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel" api:"computed"`
	AutomatedBy                     types.String                                                                                       `tfsdk:"automated_by" json:"automatedBy" api:"computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                                        `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount" api:"computed"`
	CommunityRole                   types.String                                                                                       `tfsdk:"community_role" json:"communityRole" api:"computed"`
	CoverPicture                    types.String                                                                                       `tfsdk:"cover_picture" json:"coverPicture" api:"computed"`
	CreatedAt                       types.String                                                                                       `tfsdk:"created_at" json:"createdAt" api:"computed"`
	CreatorSubscriptionsCount       types.Int64                                                                                        `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount" api:"computed"`
	Description                     types.String                                                                                       `tfsdk:"description" json:"description" api:"computed"`
	FavouritesCount                 types.Int64                                                                                        `tfsdk:"favourites_count" json:"favouritesCount" api:"computed"`
	Followers                       types.Int64                                                                                        `tfsdk:"followers" json:"followers" api:"computed"`
	Following                       types.Int64                                                                                        `tfsdk:"following" json:"following" api:"computed"`
	HasCustomTimelines              types.Bool                                                                                         `tfsdk:"has_custom_timelines" json:"hasCustomTimelines" api:"computed"`
	HasGraduatedAccess              types.Bool                                                                                         `tfsdk:"has_graduated_access" json:"hasGraduatedAccess" api:"computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                                         `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile" api:"computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetRetweetedTweetAuthorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo" api:"computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetRetweetedTweetAuthorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification" api:"computed"`
	IsAutomated                     types.Bool                                                                                         `tfsdk:"is_automated" json:"isAutomated" api:"computed"`
	IsBlueVerified                  types.Bool                                                                                         `tfsdk:"is_blue_verified" json:"isBlueVerified" api:"computed"`
	IsProfileTranslatable           types.Bool                                                                                         `tfsdk:"is_profile_translatable" json:"isProfileTranslatable" api:"computed"`
	IsTranslator                    types.Bool                                                                                         `tfsdk:"is_translator" json:"isTranslator" api:"computed"`
	IsVerified                      types.Bool                                                                                         `tfsdk:"is_verified" json:"isVerified" api:"computed"`
	Location                        types.String                                                                                       `tfsdk:"location" json:"location" api:"computed"`
	MediaCount                      types.Int64                                                                                        `tfsdk:"media_count" json:"mediaCount" api:"computed"`
	ParodyCommentaryFanLabel        types.String                                                                                       `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel" api:"computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                                     `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds" api:"computed"`
	PossiblySensitive               types.Bool                                                                                         `tfsdk:"possibly_sensitive" json:"possiblySensitive" api:"computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                              `tfsdk:"profile_bio" json:"profile_bio" api:"computed"`
	ProfileBannerURL                types.String                                                                                       `tfsdk:"profile_banner_url" json:"profileBannerUrl" api:"computed"`
	ProfileDescriptionLanguage      types.String                                                                                       `tfsdk:"profile_description_language" json:"profileDescriptionLanguage" api:"computed"`
	ProfileImageShape               types.String                                                                                       `tfsdk:"profile_image_shape" json:"profileImageShape" api:"computed"`
	ProfileInterstitialType         types.String                                                                                       `tfsdk:"profile_interstitial_type" json:"profileInterstitialType" api:"computed"`
	ProfilePicture                  types.String                                                                                       `tfsdk:"profile_picture" json:"profilePicture" api:"computed"`
	ProfileSortEnabled              types.Bool                                                                                         `tfsdk:"profile_sort_enabled" json:"profileSortEnabled" api:"computed"`
	ProfileTranslatorType           types.String                                                                                       `tfsdk:"profile_translator_type" json:"profileTranslatorType" api:"computed"`
	Protected                       types.Bool                                                                                         `tfsdk:"protected" json:"protected" api:"computed"`
	StatusesCount                   types.Int64                                                                                        `tfsdk:"statuses_count" json:"statusesCount" api:"computed"`
	SuperFollowEligible             types.Bool                                                                                         `tfsdk:"super_follow_eligible" json:"superFollowEligible" api:"computed"`
	Unavailable                     types.Bool                                                                                         `tfsdk:"unavailable" json:"unavailable" api:"computed"`
	UnavailableReason               types.String                                                                                       `tfsdk:"unavailable_reason" json:"unavailableReason" api:"computed"`
	URL                             types.String                                                                                       `tfsdk:"url" json:"url" api:"computed"`
	Verified                        types.Bool                                                                                         `tfsdk:"verified" json:"verified" api:"computed"`
	VerifiedType                    types.String                                                                                       `tfsdk:"verified_type" json:"verifiedType" api:"computed"`
	WithheldInCountries             customfield.List[types.String]                                                                     `tfsdk:"withheld_in_countries" json:"withheldInCountries" api:"computed"`
}

type XTweetTweetRetweetedTweetAuthorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl" api:"computed"`
	Description          types.String `tfsdk:"description" json:"description" api:"computed"`
	URL                  types.String `tfsdk:"url" json:"url" api:"computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType" api:"computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType" api:"computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType" api:"computed"`
}

type XTweetTweetRetweetedTweetAuthorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets" api:"computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets" api:"computed"`
}

type XTweetTweetRetweetedTweetAuthorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description" api:"computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified" api:"computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec" api:"computed"`
}

type XTweetTweetRetweetedTweetCardDataSourceModel struct {
	ID            types.String                          `tfsdk:"id" json:"id" api:"computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues" api:"computed"`
	Name          types.String                          `tfsdk:"name" json:"name" api:"computed"`
	URL           types.String                          `tfsdk:"url" json:"url" api:"computed"`
}

type XTweetTweetRetweetedTweetCommunityNoteDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id" api:"computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl" api:"computed"`
	Footer         types.String `tfsdk:"footer" json:"footer" api:"computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle" api:"computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle" api:"computed"`
	Title          types.String `tfsdk:"title" json:"title" api:"computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle" api:"computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising" api:"computed"`
	AIGenerated customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated" api:"computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion" api:"computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAIGeneratedDataSourceModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource" api:"computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia" api:"computed"`
}

type XTweetTweetRetweetedTweetEditDataSourceModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs" api:"computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds" api:"computed"`
}

type XTweetTweetRetweetedTweetMediaDataSourceModel struct {
	MediaURL           types.String                                                                                          `tfsdk:"media_url" json:"mediaUrl" api:"computed"`
	Type               types.String                                                                                          `tfsdk:"type" json:"type" api:"computed"`
	URL                types.String                                                                                          `tfsdk:"url" json:"url" api:"computed"`
	ID                 types.String                                                                                          `tfsdk:"id" json:"id" api:"computed"`
	AllowDownload      types.Bool                                                                                            `tfsdk:"allow_download" json:"allowDownload" api:"computed"`
	AltText            types.String                                                                                          `tfsdk:"alt_text" json:"altText" api:"computed"`
	AspectRatio        customfield.List[types.Int64]                                                                         `tfsdk:"aspect_ratio" json:"aspectRatio" api:"computed"`
	AvailabilityStatus types.String                                                                                          `tfsdk:"availability_status" json:"availabilityStatus" api:"computed"`
	DisplayURL         types.String                                                                                          `tfsdk:"display_url" json:"displayUrl" api:"computed"`
	DurationMillis     types.Int64                                                                                           `tfsdk:"duration_millis" json:"durationMillis" api:"computed"`
	ExpandedURL        types.String                                                                                          `tfsdk:"expanded_url" json:"expandedUrl" api:"computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaFaceRectsDataSourceModel]] `tfsdk:"face_rects" json:"faceRects" api:"computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaFocusRectsDataSourceModel]                 `tfsdk:"focus_rects" json:"focusRects" api:"computed"`
	Height             types.Int64                                                                                           `tfsdk:"height" json:"height" api:"computed"`
	Indices            customfield.List[types.Int64]                                                                         `tfsdk:"indices" json:"indices" api:"computed"`
	MediaKey           types.String                                                                                          `tfsdk:"media_key" json:"mediaKey" api:"computed"`
	Monetizable        types.Bool                                                                                            `tfsdk:"monetizable" json:"monetizable" api:"computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetRetweetedTweetMediaSizesDataSourceModel]                       `tfsdk:"sizes" json:"sizes" api:"computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaVideoVariantsDataSourceModel]              `tfsdk:"video_variants" json:"videoVariants" api:"computed"`
	Width              types.Int64                                                                                           `tfsdk:"width" json:"width" api:"computed"`
}

type XTweetTweetRetweetedTweetMediaFaceRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h" api:"required"`
	W types.Int64 `tfsdk:"w" json:"w" api:"required"`
	X types.Int64 `tfsdk:"x" json:"x" api:"required"`
	Y types.Int64 `tfsdk:"y" json:"y" api:"required"`
}

type XTweetTweetRetweetedTweetMediaFocusRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h" api:"computed"`
	W types.Int64 `tfsdk:"w" json:"w" api:"computed"`
	X types.Int64 `tfsdk:"x" json:"x" api:"computed"`
	Y types.Int64 `tfsdk:"y" json:"y" api:"computed"`
}

type XTweetTweetRetweetedTweetMediaSizesDataSourceModel struct {
	H      types.Int64  `tfsdk:"h" json:"h" api:"computed"`
	Resize types.String `tfsdk:"resize" json:"resize" api:"computed"`
	W      types.Int64  `tfsdk:"w" json:"w" api:"computed"`
}

type XTweetTweetRetweetedTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType" api:"computed"`
	URL         types.String `tfsdk:"url" json:"url" api:"computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate" api:"computed"`
}

type XTweetTweetRetweetedTweetNoteTweetDataSourceModel struct {
	Text         types.String                                                                                `tfsdk:"text" json:"text" api:"computed"`
	ID           types.String                                                                                `tfsdk:"id" json:"id" api:"computed"`
	Entities     customfield.Map[jsontypes.Normalized]                                                       `tfsdk:"entities" json:"entities" api:"computed"`
	IsExpandable types.Bool                                                                                  `tfsdk:"is_expandable" json:"isExpandable" api:"computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetRetweetedTweetNoteTweetRichtextTagsDataSourceModel] `tfsdk:"richtext_tags" json:"richtextTags" api:"computed"`
}

type XTweetTweetRetweetedTweetNoteTweetRichtextTagsDataSourceModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex" api:"computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex" api:"computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types" api:"computed"`
}

type XTweetTweetRetweetedTweetPlaceDataSourceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id" api:"computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox" api:"computed"`
	Country     types.String                          `tfsdk:"country" json:"country" api:"computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode" api:"computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName" api:"computed"`
	Name        types.String                          `tfsdk:"name" json:"name" api:"computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType" api:"computed"`
	URL         types.String                          `tfsdk:"url" json:"url" api:"computed"`
}

type XTweetTweetRetweetedTweetPreviousCountsDataSourceModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount" api:"computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount" api:"computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount" api:"computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount" api:"computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount" api:"computed"`
}
