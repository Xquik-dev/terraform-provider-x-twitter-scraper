// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XTweetDataSourceModel struct {
	ID     types.String                                          `tfsdk:"id" path:"id,required"`
	Author customfield.NestedObject[XTweetAuthorDataSourceModel] `tfsdk:"author" json:"author,computed"`
	Tweet  customfield.NestedObject[XTweetTweetDataSourceModel]  `tfsdk:"tweet" json:"tweet,computed"`
}

type XTweetAuthorDataSourceModel struct {
	ID                              types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                    `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                    `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetAuthorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                    `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                     `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                                    `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                                    `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                    `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                     `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                                    `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                                     `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                     `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                     `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                                      `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                                      `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                      `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetAuthorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XTweetAuthorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                      `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                      `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                      `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                                      `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                      `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                                    `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                     `tfsdk:"media_count" json:"mediaCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                    `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                  `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                      `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                           `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL                types.String                                                                    `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                    `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                    `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                    `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                    `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                                      `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                    `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                                      `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                                     `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                                      `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                                      `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                    `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                    `tfsdk:"url" json:"url,computed"`
	Verified                        types.Bool                                                                      `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                    `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                  `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}

type XTweetAuthorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XTweetAuthorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XTweetAuthorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
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
	Article           customfield.NestedObject[XTweetTweetArticleDataSourceModel]           `tfsdk:"article" json:"article,computed"`
	Author            customfield.NestedObject[XTweetTweetAuthorDataSourceModel]            `tfsdk:"author" json:"author,computed"`
	Card              customfield.NestedObject[XTweetTweetCardDataSourceModel]              `tfsdk:"card" json:"card,computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetCommunityNoteDataSourceModel]     `tfsdk:"community_note" json:"communityNote,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                          `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                          `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                         `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Edit              customfield.NestedObject[XTweetTweetEditDataSourceModel]              `tfsdk:"edit" json:"edit,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                 `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                          `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                          `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                          `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                            `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                            `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                            `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                            `tfsdk:"is_reply" json:"isReply,computed"`
	IsTranslatable    types.Bool                                                            `tfsdk:"is_translatable" json:"isTranslatable,computed"`
	Lang              types.String                                                          `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetMediaDataSourceModel]         `tfsdk:"media" json:"media,computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetNoteTweetDataSourceModel]         `tfsdk:"note_tweet" json:"noteTweet,computed"`
	Place             customfield.NestedObject[XTweetTweetPlaceDataSourceModel]             `tfsdk:"place" json:"place,computed"`
	PossiblySensitive types.Bool                                                            `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetPreviousCountsDataSourceModel]    `tfsdk:"previous_counts" json:"previousCounts,computed"`
	QuotedTweet       customfield.NestedObject[XTweetTweetQuotedTweetDataSourceModel]       `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	RetweetedTweet    customfield.NestedObject[XTweetTweetRetweetedTweetDataSourceModel]    `tfsdk:"retweeted_tweet" json:"retweeted_tweet,computed"`
	Source            types.String                                                          `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                          `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                          `tfsdk:"url" json:"url,computed"`
	ViewState         types.String                                                          `tfsdk:"view_state" json:"viewState,computed"`
}

type XTweetTweetArticleDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl,computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText,computed"`
	Title         types.String `tfsdk:"title" json:"title,computed"`
}

type XTweetTweetAuthorDataSourceModel struct {
	ID                              types.String                                                                         `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                         `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                         `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetAuthorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                         `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                          `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                                         `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                                         `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                         `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                          `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                                         `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                                          `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                          `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                          `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                                           `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                                           `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                           `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetAuthorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetAuthorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                           `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                           `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                           `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                                           `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                           `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                                         `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                          `tfsdk:"media_count" json:"mediaCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                         `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                       `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                           `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL                types.String                                                                         `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                         `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                         `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                         `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                         `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                                           `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                         `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                                           `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                                          `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                                           `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                                           `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                         `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                         `tfsdk:"url" json:"url,computed"`
	Verified                        types.Bool                                                                           `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                         `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                       `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}

type XTweetTweetAuthorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XTweetTweetAuthorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XTweetTweetAuthorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XTweetTweetCardDataSourceModel struct {
	ID            types.String                          `tfsdk:"id" json:"id,computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues,computed"`
	Name          types.String                          `tfsdk:"name" json:"name,computed"`
	URL           types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetCommunityNoteDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl,computed"`
	Footer         types.String `tfsdk:"footer" json:"footer,computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle,computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle,computed"`
	Title          types.String `tfsdk:"title" json:"title,computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle,computed"`
}

type XTweetTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetContentDisclosureAIGeneratedDataSourceModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetEditDataSourceModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs,computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds,computed"`
}

type XTweetTweetMediaDataSourceModel struct {
	MediaURL           types.String                                                                            `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type               types.String                                                                            `tfsdk:"type" json:"type,computed"`
	URL                types.String                                                                            `tfsdk:"url" json:"url,computed"`
	ID                 types.String                                                                            `tfsdk:"id" json:"id,computed"`
	AllowDownload      types.Bool                                                                              `tfsdk:"allow_download" json:"allowDownload,computed"`
	AltText            types.String                                                                            `tfsdk:"alt_text" json:"altText,computed"`
	AspectRatio        customfield.List[types.Int64]                                                           `tfsdk:"aspect_ratio" json:"aspectRatio,computed"`
	AvailabilityStatus types.String                                                                            `tfsdk:"availability_status" json:"availabilityStatus,computed"`
	DisplayURL         types.String                                                                            `tfsdk:"display_url" json:"displayUrl,computed"`
	DurationMillis     types.Int64                                                                             `tfsdk:"duration_millis" json:"durationMillis,computed"`
	ExpandedURL        types.String                                                                            `tfsdk:"expanded_url" json:"expandedUrl,computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetMediaFaceRectsDataSourceModel]] `tfsdk:"face_rects" json:"faceRects,computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetMediaFocusRectsDataSourceModel]                 `tfsdk:"focus_rects" json:"focusRects,computed"`
	Height             types.Int64                                                                             `tfsdk:"height" json:"height,computed"`
	Indices            customfield.List[types.Int64]                                                           `tfsdk:"indices" json:"indices,computed"`
	MediaKey           types.String                                                                            `tfsdk:"media_key" json:"mediaKey,computed"`
	Monetizable        types.Bool                                                                              `tfsdk:"monetizable" json:"monetizable,computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetMediaSizesDataSourceModel]                       `tfsdk:"sizes" json:"sizes,computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetMediaVideoVariantsDataSourceModel]              `tfsdk:"video_variants" json:"videoVariants,computed"`
	Width              types.Int64                                                                             `tfsdk:"width" json:"width,computed"`
}

type XTweetTweetMediaFaceRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h,required"`
	W types.Int64 `tfsdk:"w" json:"w,required"`
	X types.Int64 `tfsdk:"x" json:"x,required"`
	Y types.Int64 `tfsdk:"y" json:"y,required"`
}

type XTweetTweetMediaFocusRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h,computed"`
	W types.Int64 `tfsdk:"w" json:"w,computed"`
	X types.Int64 `tfsdk:"x" json:"x,computed"`
	Y types.Int64 `tfsdk:"y" json:"y,computed"`
}

type XTweetTweetMediaSizesDataSourceModel struct {
	H      types.Int64  `tfsdk:"h" json:"h,computed"`
	Resize types.String `tfsdk:"resize" json:"resize,computed"`
	W      types.Int64  `tfsdk:"w" json:"w,computed"`
}

type XTweetTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}

type XTweetTweetNoteTweetDataSourceModel struct {
	Text         types.String                                                                  `tfsdk:"text" json:"text,computed"`
	ID           types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Entities     customfield.Map[jsontypes.Normalized]                                         `tfsdk:"entities" json:"entities,computed"`
	IsExpandable types.Bool                                                                    `tfsdk:"is_expandable" json:"isExpandable,computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetNoteTweetRichtextTagsDataSourceModel] `tfsdk:"richtext_tags" json:"richtextTags,computed"`
}

type XTweetTweetNoteTweetRichtextTagsDataSourceModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex,computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types,computed"`
}

type XTweetTweetPlaceDataSourceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id,computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox,computed"`
	Country     types.String                          `tfsdk:"country" json:"country,computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode,computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName,computed"`
	Name        types.String                          `tfsdk:"name" json:"name,computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType,computed"`
	URL         types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetPreviousCountsDataSourceModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount,computed"`
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
	Article           customfield.NestedObject[XTweetTweetQuotedTweetArticleDataSourceModel]           `tfsdk:"article" json:"article,computed"`
	Author            customfield.NestedObject[XTweetTweetQuotedTweetAuthorDataSourceModel]            `tfsdk:"author" json:"author,computed"`
	Card              customfield.NestedObject[XTweetTweetQuotedTweetCardDataSourceModel]              `tfsdk:"card" json:"card,computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetQuotedTweetCommunityNoteDataSourceModel]     `tfsdk:"community_note" json:"communityNote,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                                     `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                                     `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                                    `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Edit              customfield.NestedObject[XTweetTweetQuotedTweetEditDataSourceModel]              `tfsdk:"edit" json:"edit,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                            `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                                     `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                                     `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                                     `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                                       `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                                       `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                                       `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                                       `tfsdk:"is_reply" json:"isReply,computed"`
	IsTranslatable    types.Bool                                                                       `tfsdk:"is_translatable" json:"isTranslatable,computed"`
	Lang              types.String                                                                     `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetQuotedTweetMediaDataSourceModel]         `tfsdk:"media" json:"media,computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetQuotedTweetNoteTweetDataSourceModel]         `tfsdk:"note_tweet" json:"noteTweet,computed"`
	Place             customfield.NestedObject[XTweetTweetQuotedTweetPlaceDataSourceModel]             `tfsdk:"place" json:"place,computed"`
	PossiblySensitive types.Bool                                                                       `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetQuotedTweetPreviousCountsDataSourceModel]    `tfsdk:"previous_counts" json:"previousCounts,computed"`
	QuotedTweet       customfield.NormalizedDynamicValue                                               `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	RetweetedTweet    customfield.NormalizedDynamicValue                                               `tfsdk:"retweeted_tweet" json:"retweeted_tweet,computed"`
	Source            types.String                                                                     `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                                     `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                                     `tfsdk:"url" json:"url,computed"`
	ViewState         types.String                                                                     `tfsdk:"view_state" json:"viewState,computed"`
}

type XTweetTweetQuotedTweetArticleDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl,computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText,computed"`
	Title         types.String `tfsdk:"title" json:"title,computed"`
}

type XTweetTweetQuotedTweetAuthorDataSourceModel struct {
	ID                              types.String                                                                                    `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                                    `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                                    `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetQuotedTweetAuthorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                                    `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                                     `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                                                    `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                                                    `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                                    `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                                     `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                                                    `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                                                     `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                                     `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                                     `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                                                      `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                                                      `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                                      `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetQuotedTweetAuthorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetQuotedTweetAuthorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                                      `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                                      `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                                      `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                                                      `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                                      `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                                                    `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                                     `tfsdk:"media_count" json:"mediaCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                                    `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                                  `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                                      `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                           `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL                types.String                                                                                    `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                                    `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                                    `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                                    `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                                    `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                                                      `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                                    `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                                                      `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                                                     `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                                                      `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                                                      `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                                    `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                                    `tfsdk:"url" json:"url,computed"`
	Verified                        types.Bool                                                                                      `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                                    `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                                  `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}

type XTweetTweetQuotedTweetAuthorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XTweetTweetQuotedTweetAuthorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XTweetTweetQuotedTweetAuthorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XTweetTweetQuotedTweetCardDataSourceModel struct {
	ID            types.String                          `tfsdk:"id" json:"id,computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues,computed"`
	Name          types.String                          `tfsdk:"name" json:"name,computed"`
	URL           types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetQuotedTweetCommunityNoteDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl,computed"`
	Footer         types.String `tfsdk:"footer" json:"footer,computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle,computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle,computed"`
	Title          types.String `tfsdk:"title" json:"title,computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAIGeneratedDataSourceModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetQuotedTweetEditDataSourceModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs,computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds,computed"`
}

type XTweetTweetQuotedTweetMediaDataSourceModel struct {
	MediaURL           types.String                                                                                       `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type               types.String                                                                                       `tfsdk:"type" json:"type,computed"`
	URL                types.String                                                                                       `tfsdk:"url" json:"url,computed"`
	ID                 types.String                                                                                       `tfsdk:"id" json:"id,computed"`
	AllowDownload      types.Bool                                                                                         `tfsdk:"allow_download" json:"allowDownload,computed"`
	AltText            types.String                                                                                       `tfsdk:"alt_text" json:"altText,computed"`
	AspectRatio        customfield.List[types.Int64]                                                                      `tfsdk:"aspect_ratio" json:"aspectRatio,computed"`
	AvailabilityStatus types.String                                                                                       `tfsdk:"availability_status" json:"availabilityStatus,computed"`
	DisplayURL         types.String                                                                                       `tfsdk:"display_url" json:"displayUrl,computed"`
	DurationMillis     types.Int64                                                                                        `tfsdk:"duration_millis" json:"durationMillis,computed"`
	ExpandedURL        types.String                                                                                       `tfsdk:"expanded_url" json:"expandedUrl,computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetQuotedTweetMediaFaceRectsDataSourceModel]] `tfsdk:"face_rects" json:"faceRects,computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetQuotedTweetMediaFocusRectsDataSourceModel]                 `tfsdk:"focus_rects" json:"focusRects,computed"`
	Height             types.Int64                                                                                        `tfsdk:"height" json:"height,computed"`
	Indices            customfield.List[types.Int64]                                                                      `tfsdk:"indices" json:"indices,computed"`
	MediaKey           types.String                                                                                       `tfsdk:"media_key" json:"mediaKey,computed"`
	Monetizable        types.Bool                                                                                         `tfsdk:"monetizable" json:"monetizable,computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetQuotedTweetMediaSizesDataSourceModel]                       `tfsdk:"sizes" json:"sizes,computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetQuotedTweetMediaVideoVariantsDataSourceModel]              `tfsdk:"video_variants" json:"videoVariants,computed"`
	Width              types.Int64                                                                                        `tfsdk:"width" json:"width,computed"`
}

type XTweetTweetQuotedTweetMediaFaceRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h,required"`
	W types.Int64 `tfsdk:"w" json:"w,required"`
	X types.Int64 `tfsdk:"x" json:"x,required"`
	Y types.Int64 `tfsdk:"y" json:"y,required"`
}

type XTweetTweetQuotedTweetMediaFocusRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h,computed"`
	W types.Int64 `tfsdk:"w" json:"w,computed"`
	X types.Int64 `tfsdk:"x" json:"x,computed"`
	Y types.Int64 `tfsdk:"y" json:"y,computed"`
}

type XTweetTweetQuotedTweetMediaSizesDataSourceModel struct {
	H      types.Int64  `tfsdk:"h" json:"h,computed"`
	Resize types.String `tfsdk:"resize" json:"resize,computed"`
	W      types.Int64  `tfsdk:"w" json:"w,computed"`
}

type XTweetTweetQuotedTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}

type XTweetTweetQuotedTweetNoteTweetDataSourceModel struct {
	Text         types.String                                                                             `tfsdk:"text" json:"text,computed"`
	ID           types.String                                                                             `tfsdk:"id" json:"id,computed"`
	Entities     customfield.Map[jsontypes.Normalized]                                                    `tfsdk:"entities" json:"entities,computed"`
	IsExpandable types.Bool                                                                               `tfsdk:"is_expandable" json:"isExpandable,computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetQuotedTweetNoteTweetRichtextTagsDataSourceModel] `tfsdk:"richtext_tags" json:"richtextTags,computed"`
}

type XTweetTweetQuotedTweetNoteTweetRichtextTagsDataSourceModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex,computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types,computed"`
}

type XTweetTweetQuotedTweetPlaceDataSourceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id,computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox,computed"`
	Country     types.String                          `tfsdk:"country" json:"country,computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode,computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName,computed"`
	Name        types.String                          `tfsdk:"name" json:"name,computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType,computed"`
	URL         types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetQuotedTweetPreviousCountsDataSourceModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount,computed"`
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
	Article           customfield.NestedObject[XTweetTweetRetweetedTweetArticleDataSourceModel]           `tfsdk:"article" json:"article,computed"`
	Author            customfield.NestedObject[XTweetTweetRetweetedTweetAuthorDataSourceModel]            `tfsdk:"author" json:"author,computed"`
	Card              customfield.NestedObject[XTweetTweetRetweetedTweetCardDataSourceModel]              `tfsdk:"card" json:"card,computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetRetweetedTweetCommunityNoteDataSourceModel]     `tfsdk:"community_note" json:"communityNote,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureDataSourceModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                                        `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                                        `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                                       `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Edit              customfield.NestedObject[XTweetTweetRetweetedTweetEditDataSourceModel]              `tfsdk:"edit" json:"edit,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                               `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                                        `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                                        `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                                        `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                                          `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                                          `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                                          `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                                          `tfsdk:"is_reply" json:"isReply,computed"`
	IsTranslatable    types.Bool                                                                          `tfsdk:"is_translatable" json:"isTranslatable,computed"`
	Lang              types.String                                                                        `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaDataSourceModel]         `tfsdk:"media" json:"media,computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetRetweetedTweetNoteTweetDataSourceModel]         `tfsdk:"note_tweet" json:"noteTweet,computed"`
	Place             customfield.NestedObject[XTweetTweetRetweetedTweetPlaceDataSourceModel]             `tfsdk:"place" json:"place,computed"`
	PossiblySensitive types.Bool                                                                          `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetRetweetedTweetPreviousCountsDataSourceModel]    `tfsdk:"previous_counts" json:"previousCounts,computed"`
	QuotedTweet       customfield.NormalizedDynamicValue                                                  `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	RetweetedTweet    customfield.NormalizedDynamicValue                                                  `tfsdk:"retweeted_tweet" json:"retweeted_tweet,computed"`
	Source            types.String                                                                        `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                                        `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                                        `tfsdk:"url" json:"url,computed"`
	ViewState         types.String                                                                        `tfsdk:"view_state" json:"viewState,computed"`
}

type XTweetTweetRetweetedTweetArticleDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl,computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText,computed"`
	Title         types.String `tfsdk:"title" json:"title,computed"`
}

type XTweetTweetRetweetedTweetAuthorDataSourceModel struct {
	ID                              types.String                                                                                       `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                                       `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                                       `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetRetweetedTweetAuthorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                                       `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                                        `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                                                       `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                                                       `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                                       `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                                        `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                                                       `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                                                        `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                                        `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                                        `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                                                         `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                                                         `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                                         `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetRetweetedTweetAuthorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetRetweetedTweetAuthorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                                         `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                                         `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                                         `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                                                         `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                                         `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                                                       `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                                        `tfsdk:"media_count" json:"mediaCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                                       `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                                     `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                                         `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                              `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL                types.String                                                                                       `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                                       `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                                       `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                                       `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                                       `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                                                         `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                                       `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                                                         `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                                                        `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                                                         `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                                                         `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                                       `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                                       `tfsdk:"url" json:"url,computed"`
	Verified                        types.Bool                                                                                         `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                                       `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                                     `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}

type XTweetTweetRetweetedTweetAuthorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XTweetTweetRetweetedTweetAuthorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XTweetTweetRetweetedTweetAuthorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XTweetTweetRetweetedTweetCardDataSourceModel struct {
	ID            types.String                          `tfsdk:"id" json:"id,computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues,computed"`
	Name          types.String                          `tfsdk:"name" json:"name,computed"`
	URL           types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetRetweetedTweetCommunityNoteDataSourceModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl,computed"`
	Footer         types.String `tfsdk:"footer" json:"footer,computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle,computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle,computed"`
	Title          types.String `tfsdk:"title" json:"title,computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureDataSourceModel struct {
	Advertising customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAdvertisingDataSourceModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAIGeneratedDataSourceModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAdvertisingDataSourceModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAIGeneratedDataSourceModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetRetweetedTweetEditDataSourceModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs,computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds,computed"`
}

type XTweetTweetRetweetedTweetMediaDataSourceModel struct {
	MediaURL           types.String                                                                                          `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type               types.String                                                                                          `tfsdk:"type" json:"type,computed"`
	URL                types.String                                                                                          `tfsdk:"url" json:"url,computed"`
	ID                 types.String                                                                                          `tfsdk:"id" json:"id,computed"`
	AllowDownload      types.Bool                                                                                            `tfsdk:"allow_download" json:"allowDownload,computed"`
	AltText            types.String                                                                                          `tfsdk:"alt_text" json:"altText,computed"`
	AspectRatio        customfield.List[types.Int64]                                                                         `tfsdk:"aspect_ratio" json:"aspectRatio,computed"`
	AvailabilityStatus types.String                                                                                          `tfsdk:"availability_status" json:"availabilityStatus,computed"`
	DisplayURL         types.String                                                                                          `tfsdk:"display_url" json:"displayUrl,computed"`
	DurationMillis     types.Int64                                                                                           `tfsdk:"duration_millis" json:"durationMillis,computed"`
	ExpandedURL        types.String                                                                                          `tfsdk:"expanded_url" json:"expandedUrl,computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaFaceRectsDataSourceModel]] `tfsdk:"face_rects" json:"faceRects,computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaFocusRectsDataSourceModel]                 `tfsdk:"focus_rects" json:"focusRects,computed"`
	Height             types.Int64                                                                                           `tfsdk:"height" json:"height,computed"`
	Indices            customfield.List[types.Int64]                                                                         `tfsdk:"indices" json:"indices,computed"`
	MediaKey           types.String                                                                                          `tfsdk:"media_key" json:"mediaKey,computed"`
	Monetizable        types.Bool                                                                                            `tfsdk:"monetizable" json:"monetizable,computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetRetweetedTweetMediaSizesDataSourceModel]                       `tfsdk:"sizes" json:"sizes,computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaVideoVariantsDataSourceModel]              `tfsdk:"video_variants" json:"videoVariants,computed"`
	Width              types.Int64                                                                                           `tfsdk:"width" json:"width,computed"`
}

type XTweetTweetRetweetedTweetMediaFaceRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h,required"`
	W types.Int64 `tfsdk:"w" json:"w,required"`
	X types.Int64 `tfsdk:"x" json:"x,required"`
	Y types.Int64 `tfsdk:"y" json:"y,required"`
}

type XTweetTweetRetweetedTweetMediaFocusRectsDataSourceModel struct {
	H types.Int64 `tfsdk:"h" json:"h,computed"`
	W types.Int64 `tfsdk:"w" json:"w,computed"`
	X types.Int64 `tfsdk:"x" json:"x,computed"`
	Y types.Int64 `tfsdk:"y" json:"y,computed"`
}

type XTweetTweetRetweetedTweetMediaSizesDataSourceModel struct {
	H      types.Int64  `tfsdk:"h" json:"h,computed"`
	Resize types.String `tfsdk:"resize" json:"resize,computed"`
	W      types.Int64  `tfsdk:"w" json:"w,computed"`
}

type XTweetTweetRetweetedTweetMediaVideoVariantsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}

type XTweetTweetRetweetedTweetNoteTweetDataSourceModel struct {
	Text         types.String                                                                                `tfsdk:"text" json:"text,computed"`
	ID           types.String                                                                                `tfsdk:"id" json:"id,computed"`
	Entities     customfield.Map[jsontypes.Normalized]                                                       `tfsdk:"entities" json:"entities,computed"`
	IsExpandable types.Bool                                                                                  `tfsdk:"is_expandable" json:"isExpandable,computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetRetweetedTweetNoteTweetRichtextTagsDataSourceModel] `tfsdk:"richtext_tags" json:"richtextTags,computed"`
}

type XTweetTweetRetweetedTweetNoteTweetRichtextTagsDataSourceModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex,computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types,computed"`
}

type XTweetTweetRetweetedTweetPlaceDataSourceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id,computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox,computed"`
	Country     types.String                          `tfsdk:"country" json:"country,computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode,computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName,computed"`
	Name        types.String                          `tfsdk:"name" json:"name,computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType,computed"`
	URL         types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetRetweetedTweetPreviousCountsDataSourceModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount,computed"`
}
