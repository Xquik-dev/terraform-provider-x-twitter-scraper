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
	Billing               customfield.NestedObject[XTweetBillingModel]    `tfsdk:"billing" json:"billing,computed,no_refresh"`
	NextAction            customfield.NestedObject[XTweetNextActionModel] `tfsdk:"next_action" json:"nextAction,computed,no_refresh"`
	Request               customfield.NestedObject[XTweetRequestModel]    `tfsdk:"request" json:"request,computed,no_refresh"`
	Result                customfield.NestedObject[XTweetResultModel]     `tfsdk:"result" json:"result,computed,no_refresh"`
	Target                customfield.NestedObject[XTweetTargetModel]     `tfsdk:"target" json:"target,computed,no_refresh"`
	Tweet                 customfield.NestedObject[XTweetTweetModel]      `tfsdk:"tweet" json:"tweet,computed"`
	Author                jsontypes.Normalized                            `tfsdk:"author" json:"author,computed"`
}

func (m XTweetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XTweetModel) MarshalJSONForUpdate(state XTweetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
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
	Author            jsontypes.Normalized                                        `tfsdk:"author" json:"author,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetContentDisclosureModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                               `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                       `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                  `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                  `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                  `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                  `tfsdk:"is_reply" json:"isReply,computed"`
	Lang              types.String                                                `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetMediaModel]         `tfsdk:"media" json:"media,computed"`
	QuotedTweet       customfield.NestedObject[XTweetTweetQuotedTweetModel]       `tfsdk:"quoted_tweet" json:"quoted_tweet,computed"`
	RetweetedTweet    customfield.NestedObject[XTweetTweetRetweetedTweetModel]    `tfsdk:"retweeted_tweet" json:"retweeted_tweet,computed"`
	Source            types.String                                                `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetContentDisclosureModel struct {
	Advertising customfield.NestedObject[XTweetTweetContentDisclosureAdvertisingModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetContentDisclosureAIGeneratedModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetContentDisclosureAdvertisingModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetContentDisclosureAIGeneratedModel struct {
	CanEdit             types.Bool   `tfsdk:"can_edit" json:"canEdit,computed"`
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetMediaModel struct {
	MediaURL      types.String                                                     `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type          types.String                                                     `tfsdk:"type" json:"type,computed"`
	URL           types.String                                                     `tfsdk:"url" json:"url,computed"`
	VideoVariants customfield.NestedObjectList[XTweetTweetMediaVideoVariantsModel] `tfsdk:"video_variants" json:"videoVariants,computed"`
}

type XTweetTweetMediaVideoVariantsModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
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
	Author            customfield.NestedObject[XTweetTweetQuotedTweetAuthorModel]            `tfsdk:"author" json:"author,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                           `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                           `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                          `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                  `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                           `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                           `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                           `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                             `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                             `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                             `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                             `tfsdk:"is_reply" json:"isReply,computed"`
	Lang              types.String                                                           `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetQuotedTweetMediaModel]         `tfsdk:"media" json:"media,computed"`
	Source            types.String                                                           `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                           `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                           `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetQuotedTweetAuthorModel struct {
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

type XTweetTweetQuotedTweetContentDisclosureModel struct {
	Advertising customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAdvertisingModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetQuotedTweetContentDisclosureAIGeneratedModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAdvertisingModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetQuotedTweetContentDisclosureAIGeneratedModel struct {
	CanEdit             types.Bool   `tfsdk:"can_edit" json:"canEdit,computed"`
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetQuotedTweetMediaModel struct {
	MediaURL      types.String                                                                `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type          types.String                                                                `tfsdk:"type" json:"type,computed"`
	URL           types.String                                                                `tfsdk:"url" json:"url,computed"`
	VideoVariants customfield.NestedObjectList[XTweetTweetQuotedTweetMediaVideoVariantsModel] `tfsdk:"video_variants" json:"videoVariants,computed"`
}

type XTweetTweetQuotedTweetMediaVideoVariantsModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
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
	Author            customfield.NestedObject[XTweetTweetRetweetedTweetAuthorModel]            `tfsdk:"author" json:"author,computed"`
	ContentDisclosure customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureModel] `tfsdk:"content_disclosure" json:"contentDisclosure,computed"`
	ConversationID    types.String                                                              `tfsdk:"conversation_id" json:"conversationId,computed"`
	CreatedAt         types.String                                                              `tfsdk:"created_at" json:"createdAt,computed"`
	DisplayTextRange  customfield.List[types.Int64]                                             `tfsdk:"display_text_range" json:"displayTextRange,computed"`
	Entities          customfield.Map[jsontypes.Normalized]                                     `tfsdk:"entities" json:"entities,computed"`
	InReplyToID       types.String                                                              `tfsdk:"in_reply_to_id" json:"inReplyToId,computed"`
	InReplyToUserID   types.String                                                              `tfsdk:"in_reply_to_user_id" json:"inReplyToUserId,computed"`
	InReplyToUsername types.String                                                              `tfsdk:"in_reply_to_username" json:"inReplyToUsername,computed"`
	IsLimitedReply    types.Bool                                                                `tfsdk:"is_limited_reply" json:"isLimitedReply,computed"`
	IsNoteTweet       types.Bool                                                                `tfsdk:"is_note_tweet" json:"isNoteTweet,computed"`
	IsQuoteStatus     types.Bool                                                                `tfsdk:"is_quote_status" json:"isQuoteStatus,computed"`
	IsReply           types.Bool                                                                `tfsdk:"is_reply" json:"isReply,computed"`
	Lang              types.String                                                              `tfsdk:"lang" json:"lang,computed"`
	Media             customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaModel]         `tfsdk:"media" json:"media,computed"`
	Source            types.String                                                              `tfsdk:"source" json:"source,computed"`
	Type              types.String                                                              `tfsdk:"type" json:"type,computed"`
	URL               types.String                                                              `tfsdk:"url" json:"url,computed"`
}

type XTweetTweetRetweetedTweetAuthorModel struct {
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

type XTweetTweetRetweetedTweetContentDisclosureModel struct {
	Advertising customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAdvertisingModel] `tfsdk:"advertising" json:"advertising,computed"`
	AIGenerated customfield.NestedObject[XTweetTweetRetweetedTweetContentDisclosureAIGeneratedModel] `tfsdk:"ai_generated" json:"aiGenerated,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAdvertisingModel struct {
	IsPaidPromotion types.Bool `tfsdk:"is_paid_promotion" json:"isPaidPromotion,computed"`
}

type XTweetTweetRetweetedTweetContentDisclosureAIGeneratedModel struct {
	CanEdit             types.Bool   `tfsdk:"can_edit" json:"canEdit,computed"`
	DetectionSource     types.String `tfsdk:"detection_source" json:"detectionSource,computed"`
	HasAIGeneratedMedia types.Bool   `tfsdk:"has_ai_generated_media" json:"hasAiGeneratedMedia,computed"`
}

type XTweetTweetRetweetedTweetMediaModel struct {
	MediaURL      types.String                                                                   `tfsdk:"media_url" json:"mediaUrl,computed"`
	Type          types.String                                                                   `tfsdk:"type" json:"type,computed"`
	URL           types.String                                                                   `tfsdk:"url" json:"url,computed"`
	VideoVariants customfield.NestedObjectList[XTweetTweetRetweetedTweetMediaVideoVariantsModel] `tfsdk:"video_variants" json:"videoVariants,computed"`
}

type XTweetTweetRetweetedTweetMediaVideoVariantsModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
	Bitrate     types.Int64  `tfsdk:"bitrate" json:"bitrate,computed"`
}
