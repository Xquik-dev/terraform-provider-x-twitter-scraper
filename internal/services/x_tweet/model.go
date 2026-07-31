// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XTweetModel struct {
	ID                    types.String                                    `tfsdk:"id" json:"id,computed"`
	Account               types.String                                    `tfsdk:"account" json:"account,required,no_refresh"`
	CommunityID           types.String                                    `tfsdk:"community_id" json:"community_id,optional,no_refresh"`
	IsNoteTweet           types.Bool                                      `tfsdk:"is_note_tweet" json:"is_note_tweet,optional,no_refresh"`
	ReplyToTweetID        types.String                                    `tfsdk:"reply_to_tweet_id" json:"reply_to_tweet_id,optional,no_refresh"`
	Text                  types.String                                    `tfsdk:"text" json:"text,optional,no_refresh"`
	Media                 *[]types.String                                 `tfsdk:"media" json:"media,optional,no_refresh"`
	Action                types.String                                    `tfsdk:"action" json:"action,computed,no_refresh"`
	Charged               types.Bool                                      `tfsdk:"charged" json:"charged,computed,no_refresh"`
	ChargedCredits        types.String                                    `tfsdk:"charged_credits" json:"chargedCredits,computed,no_refresh"`
	CommunityName         types.String                                    `tfsdk:"community_name" json:"communityName,computed,no_refresh"`
	CompletedAt           timetypes.RFC3339                               `tfsdk:"completed_at" json:"completedAt,computed,no_refresh" format:"date-time"`
	ConfirmationAttempts  types.Int64                                     `tfsdk:"confirmation_attempts" json:"confirmationAttempts,computed,no_refresh"`
	ConfirmationCheckedAt timetypes.RFC3339                               `tfsdk:"confirmation_checked_at" json:"confirmationCheckedAt,computed,no_refresh" format:"date-time"`
	ConfirmedAt           timetypes.RFC3339                               `tfsdk:"confirmed_at" json:"confirmedAt,computed,no_refresh" format:"date-time"`
	CreatedAt             timetypes.RFC3339                               `tfsdk:"created_at" json:"createdAt,computed,no_refresh" format:"date-time"`
	Error                 types.String                                    `tfsdk:"error" json:"error,computed,no_refresh"`
	ExpiresAt             timetypes.RFC3339                               `tfsdk:"expires_at" json:"expiresAt,computed,no_refresh" format:"date-time"`
	Idempotent            types.Bool                                      `tfsdk:"idempotent" json:"idempotent,computed,no_refresh"`
	MediaID               types.String                                    `tfsdk:"media_id" json:"mediaId,computed,no_refresh"`
	MediaURL              types.String                                    `tfsdk:"media_url" json:"mediaUrl,computed,no_refresh"`
	Message               types.String                                    `tfsdk:"message" json:"message,computed,no_refresh"`
	MessageID             types.String                                    `tfsdk:"message_id" json:"messageId,computed,no_refresh"`
	Object                types.String                                    `tfsdk:"object" json:"object,computed,no_refresh"`
	PollAfterMs           types.Int64                                     `tfsdk:"poll_after_ms" json:"pollAfterMs,computed,no_refresh"`
	RequestHash           types.String                                    `tfsdk:"request_hash" json:"requestHash,computed,no_refresh"`
	RequestID             types.String                                    `tfsdk:"request_id" json:"requestId,computed,no_refresh"`
	ResultID              types.String                                    `tfsdk:"result_id" json:"resultId,computed,no_refresh"`
	Retryable             types.Bool                                      `tfsdk:"retryable" json:"retryable,computed,no_refresh"`
	SafeToRetry           types.Bool                                      `tfsdk:"safe_to_retry" json:"safeToRetry,computed,no_refresh"`
	SendDispatched        types.Bool                                      `tfsdk:"send_dispatched" json:"sendDispatched,computed,no_refresh"`
	SendDispatchedAt      timetypes.RFC3339                               `tfsdk:"send_dispatched_at" json:"sendDispatchedAt,computed,no_refresh" format:"date-time"`
	Status                types.String                                    `tfsdk:"status" json:"status,computed,no_refresh"`
	StatusURL             types.String                                    `tfsdk:"status_url" json:"statusUrl,computed,no_refresh"`
	Success               types.Bool                                      `tfsdk:"success" json:"success,computed,no_refresh"`
	TargetID              types.String                                    `tfsdk:"target_id" json:"targetId,computed,no_refresh"`
	Terminal              types.Bool                                      `tfsdk:"terminal" json:"terminal,computed,no_refresh"`
	TweetID               types.String                                    `tfsdk:"tweet_id" json:"tweetId,computed,no_refresh"`
	UpdatedAt             timetypes.RFC3339                               `tfsdk:"updated_at" json:"updatedAt,computed,no_refresh" format:"date-time"`
	WriteActionID         types.String                                    `tfsdk:"write_action_id" json:"writeActionId,computed,no_refresh"`
	Details               customfield.Map[jsontypes.Normalized]           `tfsdk:"details" json:"details,computed,no_refresh"`
	Author                customfield.NestedObject[XTweetAuthorModel]     `tfsdk:"author" json:"author,computed"`
	Billing               customfield.NestedObject[XTweetBillingModel]    `tfsdk:"billing" json:"billing,computed,no_refresh"`
	NextAction            customfield.NestedObject[XTweetNextActionModel] `tfsdk:"next_action" json:"nextAction,computed,no_refresh"`
	Request               customfield.NestedObject[XTweetRequestModel]    `tfsdk:"request" json:"request,computed,no_refresh"`
	Result                customfield.NestedObject[XTweetResultModel]     `tfsdk:"result" json:"result,computed,no_refresh"`
	Target                customfield.NestedObject[XTweetTargetModel]     `tfsdk:"target" json:"target,computed,no_refresh"`
	Tweet                 customfield.NestedObject[XTweetTweetModel]      `tfsdk:"tweet" json:"tweet,computed"`
}

func (m XTweetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XTweetModel) MarshalJSONForUpdate(state XTweetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type XTweetAuthorModel struct {
	ID                              types.String                                                          `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                          `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                          `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetAuthorAffiliatesHighlightedLabelModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                          `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                           `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                          `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                          `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                          `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                           `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                          `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                           `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                           `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                           `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                            `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                            `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                            `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetAuthorHighlightsInfoModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XTweetAuthorIdentityVerificationModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                            `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                            `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                            `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                            `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                            `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                          `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                           `tfsdk:"media_count" json:"mediaCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                          `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                        `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                            `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                 `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL                types.String                                                          `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                          `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                          `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                          `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                          `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                            `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                          `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                            `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                           `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                            `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                            `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                          `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                          `tfsdk:"url" json:"url,computed"`
	Verified                        types.Bool                                                            `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                          `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                        `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
	Followers                       types.Int64                                                           `tfsdk:"followers" json:"followers,computed"`
	Verified                        types.Bool                                                            `tfsdk:"verified" json:"verified,computed"`
}

type XTweetAuthorAffiliatesHighlightedLabelModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XTweetAuthorHighlightsInfoModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XTweetAuthorIdentityVerificationModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XTweetBillingModel struct {
	Charged        types.Bool   `tfsdk:"charged" json:"charged,computed"`
	ChargedCredits types.String `tfsdk:"charged_credits" json:"chargedCredits,computed"`
	PlannedCredits types.String `tfsdk:"planned_credits" json:"plannedCredits,computed"`
	Status         types.String `tfsdk:"status" json:"status,computed"`
}

type XTweetNextActionModel struct {
	Type                      types.String `tfsdk:"type" json:"type,computed"`
	AfterMs                   types.Int64  `tfsdk:"after_ms" json:"afterMs,computed"`
	RequiresNewIdempotencyKey types.Bool   `tfsdk:"requires_new_idempotency_key" json:"requiresNewIdempotencyKey,computed"`
	URL                       types.String `tfsdk:"url" json:"url,computed"`
}

type XTweetRequestModel struct {
	Hash    types.String                          `tfsdk:"hash" json:"hash,computed"`
	Payload customfield.Map[jsontypes.Normalized] `tfsdk:"payload" json:"payload,computed"`
}

type XTweetResultModel struct {
	ID    types.String `tfsdk:"id" json:"id,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
}

type XTweetTargetModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type XTweetTweetModel struct {
	ID                types.String                                                `tfsdk:"id" json:"id,computed"`
	BookmarkCount     types.Int64                                                 `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount         types.Int64                                                 `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount        types.Int64                                                 `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount        types.Int64                                                 `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount      types.Int64                                                 `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text              types.String                                                `tfsdk:"text" json:"text,computed"`
	ViewCount         types.Int64                                                 `tfsdk:"view_count" json:"viewCount,computed"`
	Article           customfield.NestedObject[XTweetTweetArticleModel]           `tfsdk:"article" json:"article,computed"`
	Author            customfield.NestedObject[XTweetTweetAuthorModel]            `tfsdk:"author" json:"author,computed"`
	Card              customfield.NestedObject[XTweetTweetCardModel]              `tfsdk:"card" json:"card,computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetCommunityNoteModel]     `tfsdk:"community_note" json:"communityNote,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetContentDisclosureModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                               `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Edit              customfield.NestedObject[XTweetTweetEditModel]              `tfsdk:"edit" json:"edit,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                       `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                  `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                  `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                  `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                  `tfsdk:"is_reply" json:"isReply,computed"`
	IsTranslatable    types.Bool                                                  `tfsdk:"is_translatable" json:"isTranslatable,computed"`
	Lang              types.String                                                `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetMediaModel]         `tfsdk:"media" json:"media,computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetNoteTweetModel]         `tfsdk:"note_tweet" json:"noteTweet,computed"`
	Place             customfield.NestedObject[XTweetTweetPlaceModel]             `tfsdk:"place" json:"place,computed"`
	PossiblySensitive types.Bool                                                  `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetPreviousCountsModel]    `tfsdk:"previous_counts" json:"previousCounts,computed"`
	QuotedTweet       customfield.NestedObject[XTweetTweetQuotedTweetModel]       `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	RetweetedTweet    customfield.NestedObject[XTweetTweetRetweetedTweetModel]    `tfsdk:"retweeted_tweet" json:"retweeted_tweet,computed"`
	Source            types.String                                                `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                `tfsdk:"url" json:"url,computed"`
	ViewState         types.String                                                `tfsdk:"view_state" json:"viewState,computed"`
}

type XTweetTweetArticleModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl,computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText,computed"`
	Title         types.String `tfsdk:"title" json:"title,computed"`
}

type XTweetTweetAuthorModel struct {
	ID                              types.String                                                               `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                               `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                               `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetAuthorAffiliatesHighlightedLabelModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                               `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                               `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                               `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                               `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                               `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                                `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                                 `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                                 `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                 `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetAuthorHighlightsInfoModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetAuthorIdentityVerificationModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                 `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                 `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                 `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                                 `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                 `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                               `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                `tfsdk:"media_count" json:"mediaCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                               `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                             `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                 `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                      `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL                types.String                                                               `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                               `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                               `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                               `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                               `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                                 `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                               `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                                 `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                                `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                                 `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                                 `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                               `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                               `tfsdk:"url" json:"url,computed"`
	Verified                        types.Bool                                                                 `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                               `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                             `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
	Followers                       types.Int64                                                                `tfsdk:"followers" json:"followers,computed"`
	Verified                        types.Bool                                                                 `tfsdk:"verified" json:"verified,computed"`
}

type XTweetTweetAuthorAffiliatesHighlightedLabelModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XTweetTweetAuthorHighlightsInfoModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XTweetTweetAuthorIdentityVerificationModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XTweetTweetCardModel struct {
	ID            types.String                          `tfsdk:"id" json:"id,computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues,computed"`
	Name          types.String                          `tfsdk:"name" json:"name,computed"`
	URL           types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetCommunityNoteModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl,computed"`
	Footer         types.String `tfsdk:"footer" json:"footer,computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle,computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle,computed"`
	Title          types.String `tfsdk:"title" json:"title,computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle,computed"`
}

type XTweetTweetContentDisclosureModel struct {
	Advertising customfield.NestedObject[XTweetTweetContentDisclosureAdvertisingModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetContentDisclosureAIGeneratedModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetContentDisclosureAdvertisingModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetContentDisclosureAIGeneratedModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetEditModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs,computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds,computed"`
}

type XTweetTweetMediaModel struct {
	MediaURL           types.String                                                                  `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type               types.String                                                                  `tfsdk:"type" json:"type,computed"`
	URL                types.String                                                                  `tfsdk:"url" json:"url,computed"`
	ID                 types.String                                                                  `tfsdk:"id" json:"id,computed"`
	AllowDownload      types.Bool                                                                    `tfsdk:"allow_download" json:"allowDownload,computed"`
	AltText            types.String                                                                  `tfsdk:"alt_text" json:"altText,computed"`
	AspectRatio        customfield.List[types.Int64]                                                 `tfsdk:"aspect_ratio" json:"aspectRatio,computed"`
	AvailabilityStatus types.String                                                                  `tfsdk:"availability_status" json:"availabilityStatus,computed"`
	DisplayURL         types.String                                                                  `tfsdk:"display_url" json:"displayUrl,computed"`
	DurationMillis     types.Int64                                                                   `tfsdk:"duration_millis" json:"durationMillis,computed"`
	ExpandedURL        types.String                                                                  `tfsdk:"expanded_url" json:"expandedUrl,computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetMediaFaceRectsModel]] `tfsdk:"face_rects" json:"faceRects,computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetMediaFocusRectsModel]                 `tfsdk:"focus_rects" json:"focusRects,computed"`
	Height             types.Int64                                                                   `tfsdk:"height" json:"height,computed"`
	Indices            customfield.List[types.Int64]                                                 `tfsdk:"indices" json:"indices,computed"`
	MediaKey           types.String                                                                  `tfsdk:"media_key" json:"mediaKey,computed"`
	Monetizable        types.Bool                                                                    `tfsdk:"monetizable" json:"monetizable,computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetMediaSizesModel]                       `tfsdk:"sizes" json:"sizes,computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetMediaVideoVariantsModel]              `tfsdk:"video_variants" json:"videoVariants,computed"`
	Width              types.Int64                                                                   `tfsdk:"width" json:"width,computed"`
}

type XTweetTweetMediaFaceRectsModel struct {
	H types.Int64 `tfsdk:"h" json:"h,required"`
	W types.Int64 `tfsdk:"w" json:"w,required"`
	X types.Int64 `tfsdk:"x" json:"x,required"`
	Y types.Int64 `tfsdk:"y" json:"y,required"`
}

type XTweetTweetMediaFocusRectsModel struct {
	H types.Int64 `tfsdk:"h" json:"h,computed"`
	W types.Int64 `tfsdk:"w" json:"w,computed"`
	X types.Int64 `tfsdk:"x" json:"x,computed"`
	Y types.Int64 `tfsdk:"y" json:"y,computed"`
}

type XTweetTweetMediaSizesModel struct {
	H      types.Int64  `tfsdk:"h" json:"h,computed"`
	Resize types.String `tfsdk:"resize" json:"resize,computed"`
	W      types.Int64  `tfsdk:"w" json:"w,computed"`
}

type XTweetTweetMediaVideoVariantsModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}

type XTweetTweetNoteTweetModel struct {
	Text         types.String                                                        `tfsdk:"text" json:"text,computed"`
	ID           types.String                                                        `tfsdk:"id" json:"id,computed"`
	Entities     customfield.Map[jsontypes.Normalized]                               `tfsdk:"entities" json:"entities,computed"`
	IsExpandable types.Bool                                                          `tfsdk:"is_expandable" json:"isExpandable,computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetNoteTweetRichtextTagsModel] `tfsdk:"richtext_tags" json:"richtextTags,computed"`
}

type XTweetTweetNoteTweetRichtextTagsModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex,computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types,computed"`
}

type XTweetTweetPlaceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id,computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox,computed"`
	Country     types.String                          `tfsdk:"country" json:"country,computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode,computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName,computed"`
	Name        types.String                          `tfsdk:"name" json:"name,computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType,computed"`
	URL         types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetPreviousCountsModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount,computed"`
}

type XTweetTweetQuotedTweetModel struct {
	ID                types.String                                                           `tfsdk:"id" json:"id,computed"`
	BookmarkCount     types.Int64                                                            `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount         types.Int64                                                            `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount        types.Int64                                                            `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount        types.Int64                                                            `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount      types.Int64                                                            `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text              types.String                                                           `tfsdk:"text" json:"text,computed"`
	ViewCount         types.Int64                                                            `tfsdk:"view_count" json:"viewCount,computed"`
	Article           customfield.NestedObject[XTweetTweetQuotedTweetArticleModel]           `tfsdk:"article" json:"article,computed"`
	Author            customfield.NestedObject[XTweetTweetQuotedTweetAuthorModel]            `tfsdk:"author" json:"author,computed"`
	Card              customfield.NestedObject[XTweetTweetQuotedTweetCardModel]              `tfsdk:"card" json:"card,computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetQuotedTweetCommunityNoteModel]     `tfsdk:"community_note" json:"communityNote,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                           `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                           `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                          `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Edit              customfield.NestedObject[XTweetTweetQuotedTweetEditModel]              `tfsdk:"edit" json:"edit,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                  `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                           `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                           `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                           `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                             `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                             `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                             `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                             `tfsdk:"is_reply" json:"isReply,computed"`
	IsTranslatable    types.Bool                                                             `tfsdk:"is_translatable" json:"isTranslatable,computed"`
	Lang              types.String                                                           `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetQuotedTweetMediaModel]         `tfsdk:"media" json:"media,computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetQuotedTweetNoteTweetModel]         `tfsdk:"note_tweet" json:"noteTweet,computed"`
	Place             customfield.NestedObject[XTweetTweetQuotedTweetPlaceModel]             `tfsdk:"place" json:"place,computed"`
	PossiblySensitive types.Bool                                                             `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetQuotedTweetPreviousCountsModel]    `tfsdk:"previous_counts" json:"previousCounts,computed"`
	QuotedTweet       customfield.NormalizedDynamicValue                                     `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	RetweetedTweet    customfield.NormalizedDynamicValue                                     `tfsdk:"retweeted_tweet" json:"retweeted_tweet,computed"`
	Source            types.String                                                           `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                           `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                           `tfsdk:"url" json:"url,computed"`
	ViewState         types.String                                                           `tfsdk:"view_state" json:"viewState,computed"`
}

type XTweetTweetQuotedTweetArticleModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl,computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText,computed"`
	Title         types.String `tfsdk:"title" json:"title,computed"`
}

type XTweetTweetQuotedTweetAuthorModel struct {
	ID                              types.String                                                                          `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                          `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                          `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetQuotedTweetAuthorAffiliatesHighlightedLabelModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                          `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                           `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                                          `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                                          `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                          `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                           `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                                          `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                                           `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                           `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                           `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                                            `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                                            `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                            `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetQuotedTweetAuthorHighlightsInfoModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetQuotedTweetAuthorIdentityVerificationModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                            `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                            `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                            `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                                            `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                            `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                                          `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                           `tfsdk:"media_count" json:"mediaCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                          `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                        `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                            `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                 `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL                types.String                                                                          `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                          `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                          `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                          `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                          `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                                            `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                          `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                                            `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                                           `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                                            `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                                            `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                          `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                          `tfsdk:"url" json:"url,computed"`
	Verified                        types.Bool                                                                            `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                          `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                        `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}

type XTweetTweetQuotedTweetAuthorAffiliatesHighlightedLabelModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XTweetTweetQuotedTweetAuthorHighlightsInfoModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XTweetTweetQuotedTweetAuthorIdentityVerificationModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XTweetTweetQuotedTweetCardModel struct {
	ID            types.String                          `tfsdk:"id" json:"id,computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues,computed"`
	Name          types.String                          `tfsdk:"name" json:"name,computed"`
	URL           types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetQuotedTweetCommunityNoteModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl,computed"`
	Footer         types.String `tfsdk:"footer" json:"footer,computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle,computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle,computed"`
	Title          types.String `tfsdk:"title" json:"title,computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureModel struct {
	Advertising customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAdvertisingModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAIGeneratedModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAdvertisingModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAIGeneratedModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetQuotedTweetEditModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs,computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds,computed"`
}

type XTweetTweetQuotedTweetMediaModel struct {
	MediaURL           types.String                                                                             `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type               types.String                                                                             `tfsdk:"type" json:"type,computed"`
	URL                types.String                                                                             `tfsdk:"url" json:"url,computed"`
	ID                 types.String                                                                             `tfsdk:"id" json:"id,computed"`
	AllowDownload      types.Bool                                                                               `tfsdk:"allow_download" json:"allowDownload,computed"`
	AltText            types.String                                                                             `tfsdk:"alt_text" json:"altText,computed"`
	AspectRatio        customfield.List[types.Int64]                                                            `tfsdk:"aspect_ratio" json:"aspectRatio,computed"`
	AvailabilityStatus types.String                                                                             `tfsdk:"availability_status" json:"availabilityStatus,computed"`
	DisplayURL         types.String                                                                             `tfsdk:"display_url" json:"displayUrl,computed"`
	DurationMillis     types.Int64                                                                              `tfsdk:"duration_millis" json:"durationMillis,computed"`
	ExpandedURL        types.String                                                                             `tfsdk:"expanded_url" json:"expandedUrl,computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetQuotedTweetMediaFaceRectsModel]] `tfsdk:"face_rects" json:"faceRects,computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetQuotedTweetMediaFocusRectsModel]                 `tfsdk:"focus_rects" json:"focusRects,computed"`
	Height             types.Int64                                                                              `tfsdk:"height" json:"height,computed"`
	Indices            customfield.List[types.Int64]                                                            `tfsdk:"indices" json:"indices,computed"`
	MediaKey           types.String                                                                             `tfsdk:"media_key" json:"mediaKey,computed"`
	Monetizable        types.Bool                                                                               `tfsdk:"monetizable" json:"monetizable,computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetQuotedTweetMediaSizesModel]                       `tfsdk:"sizes" json:"sizes,computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetQuotedTweetMediaVideoVariantsModel]              `tfsdk:"video_variants" json:"videoVariants,computed"`
	Width              types.Int64                                                                              `tfsdk:"width" json:"width,computed"`
}

type XTweetTweetQuotedTweetMediaFaceRectsModel struct {
	H types.Int64 `tfsdk:"h" json:"h,required"`
	W types.Int64 `tfsdk:"w" json:"w,required"`
	X types.Int64 `tfsdk:"x" json:"x,required"`
	Y types.Int64 `tfsdk:"y" json:"y,required"`
}

type XTweetTweetQuotedTweetMediaFocusRectsModel struct {
	H types.Int64 `tfsdk:"h" json:"h,computed"`
	W types.Int64 `tfsdk:"w" json:"w,computed"`
	X types.Int64 `tfsdk:"x" json:"x,computed"`
	Y types.Int64 `tfsdk:"y" json:"y,computed"`
}

type XTweetTweetQuotedTweetMediaSizesModel struct {
	H      types.Int64  `tfsdk:"h" json:"h,computed"`
	Resize types.String `tfsdk:"resize" json:"resize,computed"`
	W      types.Int64  `tfsdk:"w" json:"w,computed"`
}

type XTweetTweetQuotedTweetMediaVideoVariantsModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}

type XTweetTweetQuotedTweetNoteTweetModel struct {
	Text         types.String                                                                   `tfsdk:"text" json:"text,computed"`
	ID           types.String                                                                   `tfsdk:"id" json:"id,computed"`
	Entities     customfield.Map[jsontypes.Normalized]                                          `tfsdk:"entities" json:"entities,computed"`
	IsExpandable types.Bool                                                                     `tfsdk:"is_expandable" json:"isExpandable,computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetQuotedTweetNoteTweetRichtextTagsModel] `tfsdk:"richtext_tags" json:"richtextTags,computed"`
}

type XTweetTweetQuotedTweetNoteTweetRichtextTagsModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex,computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types,computed"`
}

type XTweetTweetQuotedTweetPlaceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id,computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox,computed"`
	Country     types.String                          `tfsdk:"country" json:"country,computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode,computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName,computed"`
	Name        types.String                          `tfsdk:"name" json:"name,computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType,computed"`
	URL         types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetQuotedTweetPreviousCountsModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount,computed"`
}

type XTweetTweetRetweetedTweetModel struct {
	ID                types.String                                                              `tfsdk:"id" json:"id,computed"`
	BookmarkCount     types.Int64                                                               `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount         types.Int64                                                               `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount        types.Int64                                                               `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount        types.Int64                                                               `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount      types.Int64                                                               `tfsdk:"retweet_count" json:"retweetCount,computed"`
	Text              types.String                                                              `tfsdk:"text" json:"text,computed"`
	ViewCount         types.Int64                                                               `tfsdk:"view_count" json:"viewCount,computed"`
	Article           customfield.NestedObject[XTweetTweetRetweetedTweetArticleModel]           `tfsdk:"article" json:"article,computed"`
	Author            customfield.NestedObject[XTweetTweetRetweetedTweetAuthorModel]            `tfsdk:"author" json:"author,computed"`
	Card              customfield.NestedObject[XTweetTweetRetweetedTweetCardModel]              `tfsdk:"card" json:"card,computed"`
	CommunityNote     customfield.NestedObject[XTweetTweetRetweetedTweetCommunityNoteModel]     `tfsdk:"community_note" json:"communityNote,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                              `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                              `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                             `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Edit              customfield.NestedObject[XTweetTweetRetweetedTweetEditModel]              `tfsdk:"edit" json:"edit,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                     `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                              `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                              `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                              `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                                `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                                `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                                `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                                `tfsdk:"is_reply" json:"isReply,computed"`
	IsTranslatable    types.Bool                                                                `tfsdk:"is_translatable" json:"isTranslatable,computed"`
	Lang              types.String                                                              `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaModel]         `tfsdk:"media" json:"media,computed"`
	NoteTweet         customfield.NestedObject[XTweetTweetRetweetedTweetNoteTweetModel]         `tfsdk:"note_tweet" json:"noteTweet,computed"`
	Place             customfield.NestedObject[XTweetTweetRetweetedTweetPlaceModel]             `tfsdk:"place" json:"place,computed"`
	PossiblySensitive types.Bool                                                                `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	PreviousCounts    customfield.NestedObject[XTweetTweetRetweetedTweetPreviousCountsModel]    `tfsdk:"previous_counts" json:"previousCounts,computed"`
	QuotedTweet       customfield.NormalizedDynamicValue                                        `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	RetweetedTweet    customfield.NormalizedDynamicValue                                        `tfsdk:"retweeted_tweet" json:"retweeted_tweet,computed"`
	Source            types.String                                                              `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                              `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                              `tfsdk:"url" json:"url,computed"`
	ViewState         types.String                                                              `tfsdk:"view_state" json:"viewState,computed"`
}

type XTweetTweetRetweetedTweetArticleModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CoverMediaURL types.String `tfsdk:"cover_media_url" json:"coverMediaUrl,computed"`
	PreviewText   types.String `tfsdk:"preview_text" json:"previewText,computed"`
	Title         types.String `tfsdk:"title" json:"title,computed"`
}

type XTweetTweetRetweetedTweetAuthorModel struct {
	ID                              types.String                                                                             `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                             `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                             `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XTweetTweetRetweetedTweetAuthorAffiliatesHighlightedLabelModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                             `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                              `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                                             `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                                             `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                             `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                              `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                                             `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                                              `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                              `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                              `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                                               `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                                               `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                               `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XTweetTweetRetweetedTweetAuthorHighlightsInfoModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XTweetTweetRetweetedTweetAuthorIdentityVerificationModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                               `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                               `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                               `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                                               `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                               `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                                             `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                              `tfsdk:"media_count" json:"mediaCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                             `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                           `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                               `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                    `tfsdk:"profile_bio" json:"profile_bio,computed"`
	ProfileBannerURL                types.String                                                                             `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                             `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                             `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                             `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                             `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                                               `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                             `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                                               `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                                              `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                                               `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                                               `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                             `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                             `tfsdk:"url" json:"url,computed"`
	Verified                        types.Bool                                                                               `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                             `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                           `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}

type XTweetTweetRetweetedTweetAuthorAffiliatesHighlightedLabelModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XTweetTweetRetweetedTweetAuthorHighlightsInfoModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XTweetTweetRetweetedTweetAuthorIdentityVerificationModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XTweetTweetRetweetedTweetCardModel struct {
	ID            types.String                          `tfsdk:"id" json:"id,computed"`
	BindingValues customfield.Map[jsontypes.Normalized] `tfsdk:"binding_values" json:"bindingValues,computed"`
	Name          types.String                          `tfsdk:"name" json:"name,computed"`
	URL           types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetRetweetedTweetCommunityNoteModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	DestinationURL types.String `tfsdk:"destination_url" json:"destinationUrl,computed"`
	Footer         types.String `tfsdk:"footer" json:"footer,computed"`
	ShortTitle     types.String `tfsdk:"short_title" json:"shortTitle,computed"`
	Subtitle       types.String `tfsdk:"subtitle" json:"subtitle,computed"`
	Title          types.String `tfsdk:"title" json:"title,computed"`
	VisualStyle    types.String `tfsdk:"visual_style" json:"visualStyle,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureModel struct {
	Advertising customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAdvertisingModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAIGeneratedModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAdvertisingModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAIGeneratedModel struct {
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetRetweetedTweetEditModel struct {
	EditableUntilMsecs types.String                   `tfsdk:"editable_until_msecs" json:"editableUntilMsecs,computed"`
	EditTweetIDs       customfield.List[types.String] `tfsdk:"edit_tweet_ids" json:"editTweetIds,computed"`
}

type XTweetTweetRetweetedTweetMediaModel struct {
	MediaURL           types.String                                                                                `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type               types.String                                                                                `tfsdk:"type" json:"type,computed"`
	URL                types.String                                                                                `tfsdk:"url" json:"url,computed"`
	ID                 types.String                                                                                `tfsdk:"id" json:"id,computed"`
	AllowDownload      types.Bool                                                                                  `tfsdk:"allow_download" json:"allowDownload,computed"`
	AltText            types.String                                                                                `tfsdk:"alt_text" json:"altText,computed"`
	AspectRatio        customfield.List[types.Int64]                                                               `tfsdk:"aspect_ratio" json:"aspectRatio,computed"`
	AvailabilityStatus types.String                                                                                `tfsdk:"availability_status" json:"availabilityStatus,computed"`
	DisplayURL         types.String                                                                                `tfsdk:"display_url" json:"displayUrl,computed"`
	DurationMillis     types.Int64                                                                                 `tfsdk:"duration_millis" json:"durationMillis,computed"`
	ExpandedURL        types.String                                                                                `tfsdk:"expanded_url" json:"expandedUrl,computed"`
	FaceRects          customfield.Map[customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaFaceRectsModel]] `tfsdk:"face_rects" json:"faceRects,computed"`
	FocusRects         customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaFocusRectsModel]                 `tfsdk:"focus_rects" json:"focusRects,computed"`
	Height             types.Int64                                                                                 `tfsdk:"height" json:"height,computed"`
	Indices            customfield.List[types.Int64]                                                               `tfsdk:"indices" json:"indices,computed"`
	MediaKey           types.String                                                                                `tfsdk:"media_key" json:"mediaKey,computed"`
	Monetizable        types.Bool                                                                                  `tfsdk:"monetizable" json:"monetizable,computed"`
	Sizes              customfield.NestedObjectMap[XTweetTweetRetweetedTweetMediaSizesModel]                       `tfsdk:"sizes" json:"sizes,computed"`
	VideoVariants      customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaVideoVariantsModel]              `tfsdk:"video_variants" json:"videoVariants,computed"`
	Width              types.Int64                                                                                 `tfsdk:"width" json:"width,computed"`
}

type XTweetTweetRetweetedTweetMediaFaceRectsModel struct {
	H types.Int64 `tfsdk:"h" json:"h,required"`
	W types.Int64 `tfsdk:"w" json:"w,required"`
	X types.Int64 `tfsdk:"x" json:"x,required"`
	Y types.Int64 `tfsdk:"y" json:"y,required"`
}

type XTweetTweetRetweetedTweetMediaFocusRectsModel struct {
	H types.Int64 `tfsdk:"h" json:"h,computed"`
	W types.Int64 `tfsdk:"w" json:"w,computed"`
	X types.Int64 `tfsdk:"x" json:"x,computed"`
	Y types.Int64 `tfsdk:"y" json:"y,computed"`
}

type XTweetTweetRetweetedTweetMediaSizesModel struct {
	H      types.Int64  `tfsdk:"h" json:"h,computed"`
	Resize types.String `tfsdk:"resize" json:"resize,computed"`
	W      types.Int64  `tfsdk:"w" json:"w,computed"`
}

type XTweetTweetRetweetedTweetMediaVideoVariantsModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}

type XTweetTweetRetweetedTweetNoteTweetModel struct {
	Text         types.String                                                                      `tfsdk:"text" json:"text,computed"`
	ID           types.String                                                                      `tfsdk:"id" json:"id,computed"`
	Entities     customfield.Map[jsontypes.Normalized]                                             `tfsdk:"entities" json:"entities,computed"`
	IsExpandable types.Bool                                                                        `tfsdk:"is_expandable" json:"isExpandable,computed"`
	RichtextTags customfield.NestedObjectList[XTweetTweetRetweetedTweetNoteTweetRichtextTagsModel] `tfsdk:"richtext_tags" json:"richtextTags,computed"`
}

type XTweetTweetRetweetedTweetNoteTweetRichtextTagsModel struct {
	FromIndex types.Int64                    `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64                    `tfsdk:"to_index" json:"toIndex,computed"`
	Types     customfield.List[types.String] `tfsdk:"types" json:"types,computed"`
}

type XTweetTweetRetweetedTweetPlaceModel struct {
	ID          types.String                          `tfsdk:"id" json:"id,computed"`
	BoundingBox customfield.Map[jsontypes.Normalized] `tfsdk:"bounding_box" json:"boundingBox,computed"`
	Country     types.String                          `tfsdk:"country" json:"country,computed"`
	CountryCode types.String                          `tfsdk:"country_code" json:"countryCode,computed"`
	FullName    types.String                          `tfsdk:"full_name" json:"fullName,computed"`
	Name        types.String                          `tfsdk:"name" json:"name,computed"`
	PlaceType   types.String                          `tfsdk:"place_type" json:"placeType,computed"`
	URL         types.String                          `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetRetweetedTweetPreviousCountsModel struct {
	BookmarkCount types.Int64 `tfsdk:"bookmark_count" json:"bookmarkCount,computed"`
	LikeCount     types.Int64 `tfsdk:"like_count" json:"likeCount,computed"`
	QuoteCount    types.Int64 `tfsdk:"quote_count" json:"quoteCount,computed"`
	ReplyCount    types.Int64 `tfsdk:"reply_count" json:"replyCount,computed"`
	RetweetCount  types.Int64 `tfsdk:"retweet_count" json:"retweetCount,computed"`
}
