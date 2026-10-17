// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package x_space

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XSpaceDataSourceModel struct {
	ID    types.String                                         `tfsdk:"id" path:"id,required"`
	Space customfield.NestedObject[XSpaceSpaceDataSourceModel] `tfsdk:"space" json:"space,computed"`
}

type XSpaceSpaceDataSourceModel struct {
	ID                     types.String                                                      `tfsdk:"id" json:"id,computed"`
	Admins                 customfield.NestedObjectList[XSpaceSpaceAdminsDataSourceModel]    `tfsdk:"admins" json:"admins,computed"`
	Listeners              customfield.NestedObjectList[XSpaceSpaceListenersDataSourceModel] `tfsdk:"listeners" json:"listeners,computed"`
	Speakers               customfield.NestedObjectList[XSpaceSpaceSpeakersDataSourceModel]  `tfsdk:"speakers" json:"speakers,computed"`
	ContentType            types.String                                                      `tfsdk:"content_type" json:"contentType,computed"`
	ConversationControls   types.Int64                                                       `tfsdk:"conversation_controls" json:"conversationControls,computed"`
	CreatedAt              timetypes.RFC3339                                                 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Creator                customfield.NestedObject[XSpaceSpaceCreatorDataSourceModel]       `tfsdk:"creator" json:"creator,computed"`
	DisallowJoin           types.Bool                                                        `tfsdk:"disallow_join" json:"disallowJoin,computed"`
	EndedAt                timetypes.RFC3339                                                 `tfsdk:"ended_at" json:"endedAt,computed" format:"date-time"`
	IsAvailableForClipping types.Bool                                                        `tfsdk:"is_available_for_clipping" json:"isAvailableForClipping,computed"`
	IsAvailableForReplay   types.Bool                                                        `tfsdk:"is_available_for_replay" json:"isAvailableForReplay,computed"`
	IsEmployeeOnly         types.Bool                                                        `tfsdk:"is_employee_only" json:"isEmployeeOnly,computed"`
	IsLocked               types.Bool                                                        `tfsdk:"is_locked" json:"isLocked,computed"`
	MaxAdminCapacity       types.Int64                                                       `tfsdk:"max_admin_capacity" json:"maxAdminCapacity,computed"`
	MediaKey               types.String                                                      `tfsdk:"media_key" json:"mediaKey,computed"`
	MentionedUserIDs       customfield.List[types.String]                                    `tfsdk:"mentioned_user_ids" json:"mentionedUserIds,computed"`
	NoIncognito            types.Bool                                                        `tfsdk:"no_incognito" json:"noIncognito,computed"`
	ParticipantCount       types.Int64                                                       `tfsdk:"participant_count" json:"participantCount,computed"`
	ScheduledStart         timetypes.RFC3339                                                 `tfsdk:"scheduled_start" json:"scheduledStart,computed" format:"date-time"`
	Sharings               customfield.NestedObjectList[XSpaceSpaceSharingsDataSourceModel]  `tfsdk:"sharings" json:"sharings,computed"`
	StartedAt              timetypes.RFC3339                                                 `tfsdk:"started_at" json:"startedAt,computed" format:"date-time"`
	State                  types.String                                                      `tfsdk:"state" json:"state,computed"`
	Title                  types.String                                                      `tfsdk:"title" json:"title,computed"`
	TotalLiveListeners     types.Int64                                                       `tfsdk:"total_live_listeners" json:"totalLiveListeners,computed"`
	TotalReplayWatched     types.Int64                                                       `tfsdk:"total_replay_watched" json:"totalReplayWatched,computed"`
	Tweet                  customfield.NormalizedDynamicValue                                `tfsdk:"tweet" json:"tweet,computed"`
	TweetID                types.String                                                      `tfsdk:"tweet_id" json:"tweetId,computed"`
	UpdatedAt              timetypes.RFC3339                                                 `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
}

type XSpaceSpaceAdminsDataSourceModel struct {
	ID                         types.String                                                                         `tfsdk:"id" json:"id,computed"`
	Name                       types.String                                                                         `tfsdk:"name" json:"name,computed"`
	Username                   types.String                                                                         `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel customfield.NestedObject[XSpaceSpaceAdminsAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	IsBlueVerified             types.Bool                                                                           `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsMutedByAdmin             types.Bool                                                                           `tfsdk:"is_muted_by_admin" json:"isMutedByAdmin,computed"`
	IsMutedByGuest             types.Bool                                                                           `tfsdk:"is_muted_by_guest" json:"isMutedByGuest,computed"`
	IsVerified                 types.Bool                                                                           `tfsdk:"is_verified" json:"isVerified,computed"`
	JoinedAt                   timetypes.RFC3339                                                                    `tfsdk:"joined_at" json:"joinedAt,computed" format:"date-time"`
	PeriscopeUserID            types.String                                                                         `tfsdk:"periscope_user_id" json:"periscopeUserId,computed"`
	ProfilePicture             types.String                                                                         `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfilePictureOriginal     types.String                                                                         `tfsdk:"profile_picture_original" json:"profilePictureOriginal,computed"`
	VerifiedType               types.String                                                                         `tfsdk:"verified_type" json:"verifiedType,computed"`
}

type XSpaceSpaceAdminsAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String                          `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String                          `tfsdk:"description" json:"description,computed"`
	LongDescription      customfield.Map[jsontypes.Normalized] `tfsdk:"long_description" json:"longDescription,computed"`
	URL                  types.String                          `tfsdk:"url" json:"url,computed"`
	URLType              types.String                          `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String                          `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String                          `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XSpaceSpaceListenersDataSourceModel struct {
	ID                         types.String                                                                            `tfsdk:"id" json:"id,computed"`
	Name                       types.String                                                                            `tfsdk:"name" json:"name,computed"`
	Username                   types.String                                                                            `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel customfield.NestedObject[XSpaceSpaceListenersAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	IsBlueVerified             types.Bool                                                                              `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsMutedByAdmin             types.Bool                                                                              `tfsdk:"is_muted_by_admin" json:"isMutedByAdmin,computed"`
	IsMutedByGuest             types.Bool                                                                              `tfsdk:"is_muted_by_guest" json:"isMutedByGuest,computed"`
	IsVerified                 types.Bool                                                                              `tfsdk:"is_verified" json:"isVerified,computed"`
	JoinedAt                   timetypes.RFC3339                                                                       `tfsdk:"joined_at" json:"joinedAt,computed" format:"date-time"`
	PeriscopeUserID            types.String                                                                            `tfsdk:"periscope_user_id" json:"periscopeUserId,computed"`
	ProfilePicture             types.String                                                                            `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfilePictureOriginal     types.String                                                                            `tfsdk:"profile_picture_original" json:"profilePictureOriginal,computed"`
	VerifiedType               types.String                                                                            `tfsdk:"verified_type" json:"verifiedType,computed"`
}

type XSpaceSpaceListenersAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String                          `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String                          `tfsdk:"description" json:"description,computed"`
	LongDescription      customfield.Map[jsontypes.Normalized] `tfsdk:"long_description" json:"longDescription,computed"`
	URL                  types.String                          `tfsdk:"url" json:"url,computed"`
	URLType              types.String                          `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String                          `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String                          `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XSpaceSpaceSpeakersDataSourceModel struct {
	ID                         types.String                                                                           `tfsdk:"id" json:"id,computed"`
	Name                       types.String                                                                           `tfsdk:"name" json:"name,computed"`
	Username                   types.String                                                                           `tfsdk:"username" json:"username,computed"`
	AffiliatesHighlightedLabel customfield.NestedObject[XSpaceSpaceSpeakersAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	IsBlueVerified             types.Bool                                                                             `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsMutedByAdmin             types.Bool                                                                             `tfsdk:"is_muted_by_admin" json:"isMutedByAdmin,computed"`
	IsMutedByGuest             types.Bool                                                                             `tfsdk:"is_muted_by_guest" json:"isMutedByGuest,computed"`
	IsVerified                 types.Bool                                                                             `tfsdk:"is_verified" json:"isVerified,computed"`
	JoinedAt                   timetypes.RFC3339                                                                      `tfsdk:"joined_at" json:"joinedAt,computed" format:"date-time"`
	PeriscopeUserID            types.String                                                                           `tfsdk:"periscope_user_id" json:"periscopeUserId,computed"`
	ProfilePicture             types.String                                                                           `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfilePictureOriginal     types.String                                                                           `tfsdk:"profile_picture_original" json:"profilePictureOriginal,computed"`
	VerifiedType               types.String                                                                           `tfsdk:"verified_type" json:"verifiedType,computed"`
}

type XSpaceSpaceSpeakersAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String                          `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String                          `tfsdk:"description" json:"description,computed"`
	LongDescription      customfield.Map[jsontypes.Normalized] `tfsdk:"long_description" json:"longDescription,computed"`
	URL                  types.String                          `tfsdk:"url" json:"url,computed"`
	URLType              types.String                          `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String                          `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String                          `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XSpaceSpaceCreatorDataSourceModel struct {
	ID                              types.String                                                                          `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                          `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                          `tfsdk:"username" json:"username,computed"`
	AccountBasedIn                  customfield.NestedObject[XSpaceSpaceCreatorAccountBasedInDataSourceModel]             `tfsdk:"account_based_in" json:"accountBasedIn,computed"`
	AccountBasedInAccurate          types.Bool                                                                            `tfsdk:"account_based_in_accurate" json:"accountBasedInAccurate,computed"`
	AccountBasedInUnavailable       types.Bool                                                                            `tfsdk:"account_based_in_unavailable" json:"accountBasedInUnavailable,computed"`
	AdvertiserAccountServiceLevels  customfield.List[types.String]                                                        `tfsdk:"advertiser_account_service_levels" json:"advertiserAccountServiceLevels,computed"`
	AdvertiserAccountType           types.String                                                                          `tfsdk:"advertiser_account_type" json:"advertiserAccountType,computed"`
	AffiliateUsername               types.String                                                                          `tfsdk:"affiliate_username" json:"affiliateUsername,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XSpaceSpaceCreatorAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                          `tfsdk:"automated_by" json:"automatedBy,computed"`
	Birthdate                       customfield.NestedObject[XSpaceSpaceCreatorBirthdateDataSourceModel]                  `tfsdk:"birthdate" json:"birthdate,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                           `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	BusinessProfileState            types.String                                                                          `tfsdk:"business_profile_state" json:"businessProfileState,computed"`
	CommunityRole                   types.String                                                                          `tfsdk:"community_role" json:"communityRole,computed"`
	ConnectedVia                    types.String                                                                          `tfsdk:"connected_via" json:"connectedVia,computed"`
	ConnectedViaAccurate            types.Bool                                                                            `tfsdk:"connected_via_accurate" json:"connectedViaAccurate,computed"`
	ContactInfo                     customfield.NestedObject[XSpaceSpaceCreatorContactInfoDataSourceModel]                `tfsdk:"contact_info" json:"contactInfo,computed"`
	CoverPicture                    types.String                                                                          `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                          `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                           `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	DefaultProfile                  types.Bool                                                                            `tfsdk:"default_profile" json:"defaultProfile,computed"`
	DefaultProfileImage             types.Bool                                                                            `tfsdk:"default_profile_image" json:"defaultProfileImage,computed"`
	Description                     types.String                                                                          `tfsdk:"description" json:"description,computed"`
	Email                           types.String                                                                          `tfsdk:"email" json:"email,computed"`
	FastFollowersCount              types.Int64                                                                           `tfsdk:"fast_followers_count" json:"fastFollowersCount,computed"`
	FavouritesCount                 types.Int64                                                                           `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                           `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                           `tfsdk:"following" json:"following,computed"`
	GeoEnabled                      types.Bool                                                                            `tfsdk:"geo_enabled" json:"geoEnabled,computed"`
	GrokTranslatedBio               customfield.Map[jsontypes.Normalized]                                                 `tfsdk:"grok_translated_bio" json:"grokTranslatedBio,computed"`
	HasCustomTimelines              types.Bool                                                                            `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasExtendedProfile              types.Bool                                                                            `tfsdk:"has_extended_profile" json:"hasExtendedProfile,computed"`
	HasGraduatedAccess              types.Bool                                                                            `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                            `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XSpaceSpaceCreatorHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XSpaceSpaceCreatorIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                            `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                            `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                            `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslationEnabled            types.Bool                                                                            `tfsdk:"is_translation_enabled" json:"isTranslationEnabled,computed"`
	IsTranslator                    types.Bool                                                                            `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                            `tfsdk:"is_verified" json:"isVerified,computed"`
	ListedCount                     types.Int64                                                                           `tfsdk:"listed_count" json:"listedCount,computed"`
	Location                        types.String                                                                          `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                           `tfsdk:"media_count" json:"mediaCount,computed"`
	NormalFollowersCount            types.Int64                                                                           `tfsdk:"normal_followers_count" json:"normalFollowersCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                          `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                        `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                            `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	Professional                    customfield.Map[jsontypes.Normalized]                                                 `tfsdk:"professional" json:"professional,computed"`
	ProfileBannerURL                types.String                                                                          `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                          `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                          `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                          `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                          `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfilePictureOriginal          types.String                                                                          `tfsdk:"profile_picture_original" json:"profilePictureOriginal,computed"`
	ProfileSortEnabled              types.Bool                                                                            `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                          `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                 `tfsdk:"profile_bio" json:"profile_bio,computed"`
	Protected                       types.Bool                                                                            `tfsdk:"protected" json:"protected,computed"`
	RedirectsToUsername             types.String                                                                          `tfsdk:"redirects_to_username" json:"redirectsToUsername,computed"`
	RetweetedAt                     timetypes.RFC3339                                                                     `tfsdk:"retweeted_at" json:"retweetedAt,computed" format:"date-time"`
	StatusesCount                   types.Int64                                                                           `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SubscribersCount                types.Int64                                                                           `tfsdk:"subscribers_count" json:"subscribersCount,computed"`
	SuperFollowEligible             types.Bool                                                                            `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	SuperFollowsUserProfileActive   types.Bool                                                                            `tfsdk:"super_follows_user_profile_active" json:"superFollowsUserProfileActive,computed"`
	TipJar                          customfield.NestedObject[XSpaceSpaceCreatorTipJarDataSourceModel]                     `tfsdk:"tip_jar" json:"tipJar,computed"`
	Unavailable                     types.Bool                                                                            `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                          `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                          `tfsdk:"url" json:"url,computed"`
	UsernameChanges                 customfield.NestedObject[XSpaceSpaceCreatorUsernameChangesDataSourceModel]            `tfsdk:"username_changes" json:"usernameChanges,computed"`
	VerificationReason              customfield.NestedObject[XSpaceSpaceCreatorVerificationReasonDataSourceModel]         `tfsdk:"verification_reason" json:"verificationReason,computed"`
	Verified                        types.Bool                                                                            `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                          `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                        `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
	WithheldScope                   types.String                                                                          `tfsdk:"withheld_scope" json:"withheldScope,computed"`
}

type XSpaceSpaceCreatorAccountBasedInDataSourceModel struct {
	Level      types.String      `tfsdk:"level" json:"level,computed"`
	ObservedAt timetypes.RFC3339 `tfsdk:"observed_at" json:"observedAt,computed" format:"date-time"`
	Value      types.String      `tfsdk:"value" json:"value,computed"`
}

type XSpaceSpaceCreatorAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String                          `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String                          `tfsdk:"description" json:"description,computed"`
	LongDescription      customfield.Map[jsontypes.Normalized] `tfsdk:"long_description" json:"longDescription,computed"`
	URL                  types.String                          `tfsdk:"url" json:"url,computed"`
	URLType              types.String                          `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String                          `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String                          `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XSpaceSpaceCreatorBirthdateDataSourceModel struct {
	Day   types.Int64 `tfsdk:"day" json:"day,computed"`
	Month types.Int64 `tfsdk:"month" json:"month,computed"`
	Year  types.Int64 `tfsdk:"year" json:"year,computed"`
}

type XSpaceSpaceCreatorContactInfoDataSourceModel struct {
	Emails customfield.NestedObjectList[XSpaceSpaceCreatorContactInfoEmailsDataSourceModel] `tfsdk:"emails" json:"emails,computed"`
	Links  customfield.NestedObjectList[XSpaceSpaceCreatorContactInfoLinksDataSourceModel]  `tfsdk:"links" json:"links,computed"`
	Phones customfield.NestedObjectList[XSpaceSpaceCreatorContactInfoPhonesDataSourceModel] `tfsdk:"phones" json:"phones,computed"`
}

type XSpaceSpaceCreatorContactInfoEmailsDataSourceModel struct {
	Domain    types.String `tfsdk:"domain" json:"domain,computed"`
	Source    types.String `tfsdk:"source" json:"source,computed"`
	Value     types.String `tfsdk:"value" json:"value,computed"`
	WrittenAs types.String `tfsdk:"written_as" json:"writtenAs,computed"`
}

type XSpaceSpaceCreatorContactInfoLinksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
	Source  types.String `tfsdk:"source" json:"source,computed"`
	URL     types.String `tfsdk:"url" json:"url,computed"`
}

type XSpaceSpaceCreatorContactInfoPhonesDataSourceModel struct {
	Source    types.String `tfsdk:"source" json:"source,computed"`
	Value     types.String `tfsdk:"value" json:"value,computed"`
	WrittenAs types.String `tfsdk:"written_as" json:"writtenAs,computed"`
}

type XSpaceSpaceCreatorHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XSpaceSpaceCreatorIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XSpaceSpaceCreatorTipJarDataSourceModel struct {
	BandcampHandle types.String `tfsdk:"bandcamp_handle" json:"bandcampHandle,computed"`
	BitcoinHandle  types.String `tfsdk:"bitcoin_handle" json:"bitcoinHandle,computed"`
	CashAppHandle  types.String `tfsdk:"cash_app_handle" json:"cashAppHandle,computed"`
	EthereumHandle types.String `tfsdk:"ethereum_handle" json:"ethereumHandle,computed"`
	GofundmeHandle types.String `tfsdk:"gofundme_handle" json:"gofundmeHandle,computed"`
	IsEnabled      types.Bool   `tfsdk:"is_enabled" json:"isEnabled,computed"`
	PatreonHandle  types.String `tfsdk:"patreon_handle" json:"patreonHandle,computed"`
	PayPalHandle   types.String `tfsdk:"pay_pal_handle" json:"payPalHandle,computed"`
	VenmoHandle    types.String `tfsdk:"venmo_handle" json:"venmoHandle,computed"`
}

type XSpaceSpaceCreatorUsernameChangesDataSourceModel struct {
	Count         types.Int64       `tfsdk:"count" json:"count,computed"`
	LastChangedAt timetypes.RFC3339 `tfsdk:"last_changed_at" json:"lastChangedAt,computed" format:"date-time"`
}

type XSpaceSpaceCreatorVerificationReasonDataSourceModel struct {
	Entities customfield.NestedObjectList[XSpaceSpaceCreatorVerificationReasonEntitiesDataSourceModel] `tfsdk:"entities" json:"entities,computed"`
	Text     types.String                                                                              `tfsdk:"text" json:"text,computed"`
}

type XSpaceSpaceCreatorVerificationReasonEntitiesDataSourceModel struct {
	FromIndex types.Int64  `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64  `tfsdk:"to_index" json:"toIndex,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	URLType   types.String `tfsdk:"url_type" json:"urlType,computed"`
}

type XSpaceSpaceSharingsDataSourceModel struct {
	ID        types.String                                                         `tfsdk:"id" json:"id,computed"`
	SharedAt  timetypes.RFC3339                                                    `tfsdk:"shared_at" json:"sharedAt,computed" format:"date-time"`
	SharedBy  customfield.NestedObject[XSpaceSpaceSharingsSharedByDataSourceModel] `tfsdk:"shared_by" json:"sharedBy,computed"`
	Tweet     jsontypes.Normalized                                                 `tfsdk:"tweet" json:"tweet,computed"`
	UpdatedAt timetypes.RFC3339                                                    `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
}

type XSpaceSpaceSharingsSharedByDataSourceModel struct {
	ID                              types.String                                                                                   `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                                   `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                                   `tfsdk:"username" json:"username,computed"`
	AccountBasedIn                  customfield.NestedObject[XSpaceSpaceSharingsSharedByAccountBasedInDataSourceModel]             `tfsdk:"account_based_in" json:"accountBasedIn,computed"`
	AccountBasedInAccurate          types.Bool                                                                                     `tfsdk:"account_based_in_accurate" json:"accountBasedInAccurate,computed"`
	AccountBasedInUnavailable       types.Bool                                                                                     `tfsdk:"account_based_in_unavailable" json:"accountBasedInUnavailable,computed"`
	AdvertiserAccountServiceLevels  customfield.List[types.String]                                                                 `tfsdk:"advertiser_account_service_levels" json:"advertiserAccountServiceLevels,computed"`
	AdvertiserAccountType           types.String                                                                                   `tfsdk:"advertiser_account_type" json:"advertiserAccountType,computed"`
	AffiliateUsername               types.String                                                                                   `tfsdk:"affiliate_username" json:"affiliateUsername,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XSpaceSpaceSharingsSharedByAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                                   `tfsdk:"automated_by" json:"automatedBy,computed"`
	Birthdate                       customfield.NestedObject[XSpaceSpaceSharingsSharedByBirthdateDataSourceModel]                  `tfsdk:"birthdate" json:"birthdate,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                                    `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	BusinessProfileState            types.String                                                                                   `tfsdk:"business_profile_state" json:"businessProfileState,computed"`
	CommunityRole                   types.String                                                                                   `tfsdk:"community_role" json:"communityRole,computed"`
	ConnectedVia                    types.String                                                                                   `tfsdk:"connected_via" json:"connectedVia,computed"`
	ConnectedViaAccurate            types.Bool                                                                                     `tfsdk:"connected_via_accurate" json:"connectedViaAccurate,computed"`
	ContactInfo                     customfield.NestedObject[XSpaceSpaceSharingsSharedByContactInfoDataSourceModel]                `tfsdk:"contact_info" json:"contactInfo,computed"`
	CoverPicture                    types.String                                                                                   `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                                   `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                                    `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	DefaultProfile                  types.Bool                                                                                     `tfsdk:"default_profile" json:"defaultProfile,computed"`
	DefaultProfileImage             types.Bool                                                                                     `tfsdk:"default_profile_image" json:"defaultProfileImage,computed"`
	Description                     types.String                                                                                   `tfsdk:"description" json:"description,computed"`
	Email                           types.String                                                                                   `tfsdk:"email" json:"email,computed"`
	FastFollowersCount              types.Int64                                                                                    `tfsdk:"fast_followers_count" json:"fastFollowersCount,computed"`
	FavouritesCount                 types.Int64                                                                                    `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                                    `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                                    `tfsdk:"following" json:"following,computed"`
	GeoEnabled                      types.Bool                                                                                     `tfsdk:"geo_enabled" json:"geoEnabled,computed"`
	GrokTranslatedBio               customfield.Map[jsontypes.Normalized]                                                          `tfsdk:"grok_translated_bio" json:"grokTranslatedBio,computed"`
	HasCustomTimelines              types.Bool                                                                                     `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasExtendedProfile              types.Bool                                                                                     `tfsdk:"has_extended_profile" json:"hasExtendedProfile,computed"`
	HasGraduatedAccess              types.Bool                                                                                     `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                                     `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XSpaceSpaceSharingsSharedByHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XSpaceSpaceSharingsSharedByIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                                     `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                                     `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                                     `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslationEnabled            types.Bool                                                                                     `tfsdk:"is_translation_enabled" json:"isTranslationEnabled,computed"`
	IsTranslator                    types.Bool                                                                                     `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                                     `tfsdk:"is_verified" json:"isVerified,computed"`
	ListedCount                     types.Int64                                                                                    `tfsdk:"listed_count" json:"listedCount,computed"`
	Location                        types.String                                                                                   `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                                    `tfsdk:"media_count" json:"mediaCount,computed"`
	NormalFollowersCount            types.Int64                                                                                    `tfsdk:"normal_followers_count" json:"normalFollowersCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                                   `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                                 `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                                     `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	Professional                    customfield.Map[jsontypes.Normalized]                                                          `tfsdk:"professional" json:"professional,computed"`
	ProfileBannerURL                types.String                                                                                   `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                                   `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                                   `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                                   `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                                   `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfilePictureOriginal          types.String                                                                                   `tfsdk:"profile_picture_original" json:"profilePictureOriginal,computed"`
	ProfileSortEnabled              types.Bool                                                                                     `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                                   `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                                          `tfsdk:"profile_bio" json:"profile_bio,computed"`
	Protected                       types.Bool                                                                                     `tfsdk:"protected" json:"protected,computed"`
	RedirectsToUsername             types.String                                                                                   `tfsdk:"redirects_to_username" json:"redirectsToUsername,computed"`
	RetweetedAt                     timetypes.RFC3339                                                                              `tfsdk:"retweeted_at" json:"retweetedAt,computed" format:"date-time"`
	StatusesCount                   types.Int64                                                                                    `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SubscribersCount                types.Int64                                                                                    `tfsdk:"subscribers_count" json:"subscribersCount,computed"`
	SuperFollowEligible             types.Bool                                                                                     `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	SuperFollowsUserProfileActive   types.Bool                                                                                     `tfsdk:"super_follows_user_profile_active" json:"superFollowsUserProfileActive,computed"`
	TipJar                          customfield.NestedObject[XSpaceSpaceSharingsSharedByTipJarDataSourceModel]                     `tfsdk:"tip_jar" json:"tipJar,computed"`
	Unavailable                     types.Bool                                                                                     `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                                   `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                                   `tfsdk:"url" json:"url,computed"`
	UsernameChanges                 customfield.NestedObject[XSpaceSpaceSharingsSharedByUsernameChangesDataSourceModel]            `tfsdk:"username_changes" json:"usernameChanges,computed"`
	VerificationReason              customfield.NestedObject[XSpaceSpaceSharingsSharedByVerificationReasonDataSourceModel]         `tfsdk:"verification_reason" json:"verificationReason,computed"`
	Verified                        types.Bool                                                                                     `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                                   `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                                 `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
	WithheldScope                   types.String                                                                                   `tfsdk:"withheld_scope" json:"withheldScope,computed"`
}

type XSpaceSpaceSharingsSharedByAccountBasedInDataSourceModel struct {
	Level      types.String      `tfsdk:"level" json:"level,computed"`
	ObservedAt timetypes.RFC3339 `tfsdk:"observed_at" json:"observedAt,computed" format:"date-time"`
	Value      types.String      `tfsdk:"value" json:"value,computed"`
}

type XSpaceSpaceSharingsSharedByAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String                          `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String                          `tfsdk:"description" json:"description,computed"`
	LongDescription      customfield.Map[jsontypes.Normalized] `tfsdk:"long_description" json:"longDescription,computed"`
	URL                  types.String                          `tfsdk:"url" json:"url,computed"`
	URLType              types.String                          `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String                          `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String                          `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XSpaceSpaceSharingsSharedByBirthdateDataSourceModel struct {
	Day   types.Int64 `tfsdk:"day" json:"day,computed"`
	Month types.Int64 `tfsdk:"month" json:"month,computed"`
	Year  types.Int64 `tfsdk:"year" json:"year,computed"`
}

type XSpaceSpaceSharingsSharedByContactInfoDataSourceModel struct {
	Emails customfield.NestedObjectList[XSpaceSpaceSharingsSharedByContactInfoEmailsDataSourceModel] `tfsdk:"emails" json:"emails,computed"`
	Links  customfield.NestedObjectList[XSpaceSpaceSharingsSharedByContactInfoLinksDataSourceModel]  `tfsdk:"links" json:"links,computed"`
	Phones customfield.NestedObjectList[XSpaceSpaceSharingsSharedByContactInfoPhonesDataSourceModel] `tfsdk:"phones" json:"phones,computed"`
}

type XSpaceSpaceSharingsSharedByContactInfoEmailsDataSourceModel struct {
	Domain    types.String `tfsdk:"domain" json:"domain,computed"`
	Source    types.String `tfsdk:"source" json:"source,computed"`
	Value     types.String `tfsdk:"value" json:"value,computed"`
	WrittenAs types.String `tfsdk:"written_as" json:"writtenAs,computed"`
}

type XSpaceSpaceSharingsSharedByContactInfoLinksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
	Source  types.String `tfsdk:"source" json:"source,computed"`
	URL     types.String `tfsdk:"url" json:"url,computed"`
}

type XSpaceSpaceSharingsSharedByContactInfoPhonesDataSourceModel struct {
	Source    types.String `tfsdk:"source" json:"source,computed"`
	Value     types.String `tfsdk:"value" json:"value,computed"`
	WrittenAs types.String `tfsdk:"written_as" json:"writtenAs,computed"`
}

type XSpaceSpaceSharingsSharedByHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XSpaceSpaceSharingsSharedByIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XSpaceSpaceSharingsSharedByTipJarDataSourceModel struct {
	BandcampHandle types.String `tfsdk:"bandcamp_handle" json:"bandcampHandle,computed"`
	BitcoinHandle  types.String `tfsdk:"bitcoin_handle" json:"bitcoinHandle,computed"`
	CashAppHandle  types.String `tfsdk:"cash_app_handle" json:"cashAppHandle,computed"`
	EthereumHandle types.String `tfsdk:"ethereum_handle" json:"ethereumHandle,computed"`
	GofundmeHandle types.String `tfsdk:"gofundme_handle" json:"gofundmeHandle,computed"`
	IsEnabled      types.Bool   `tfsdk:"is_enabled" json:"isEnabled,computed"`
	PatreonHandle  types.String `tfsdk:"patreon_handle" json:"patreonHandle,computed"`
	PayPalHandle   types.String `tfsdk:"pay_pal_handle" json:"payPalHandle,computed"`
	VenmoHandle    types.String `tfsdk:"venmo_handle" json:"venmoHandle,computed"`
}

type XSpaceSpaceSharingsSharedByUsernameChangesDataSourceModel struct {
	Count         types.Int64       `tfsdk:"count" json:"count,computed"`
	LastChangedAt timetypes.RFC3339 `tfsdk:"last_changed_at" json:"lastChangedAt,computed" format:"date-time"`
}

type XSpaceSpaceSharingsSharedByVerificationReasonDataSourceModel struct {
	Entities customfield.NestedObjectList[XSpaceSpaceSharingsSharedByVerificationReasonEntitiesDataSourceModel] `tfsdk:"entities" json:"entities,computed"`
	Text     types.String                                                                                       `tfsdk:"text" json:"text,computed"`
}

type XSpaceSpaceSharingsSharedByVerificationReasonEntitiesDataSourceModel struct {
	FromIndex types.Int64  `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64  `tfsdk:"to_index" json:"toIndex,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	URLType   types.String `tfsdk:"url_type" json:"urlType,computed"`
}
