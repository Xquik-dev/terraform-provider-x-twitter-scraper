// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package x_list

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XListDataSourceModel struct {
	ID   types.String                                       `tfsdk:"id" path:"id,required"`
	List customfield.NestedObject[XListListDataSourceModel] `tfsdk:"list" json:"list,computed"`
}

type XListListDataSourceModel struct {
	ID              types.String                                                    `tfsdk:"id" json:"id,computed"`
	Name            types.String                                                    `tfsdk:"name" json:"name,computed"`
	BannerURL       types.String                                                    `tfsdk:"banner_url" json:"bannerUrl,computed"`
	CreatedAt       timetypes.RFC3339                                               `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	CustomBanner    customfield.NestedObject[XListListCustomBannerDataSourceModel]  `tfsdk:"custom_banner" json:"customBanner,computed"`
	DefaultBanner   customfield.NestedObject[XListListDefaultBannerDataSourceModel] `tfsdk:"default_banner" json:"defaultBanner,computed"`
	Description     types.String                                                    `tfsdk:"description" json:"description,computed"`
	FacepileURLs    customfield.List[types.String]                                  `tfsdk:"facepile_urls" json:"facepileUrls,computed"`
	MemberCount     types.Int64                                                     `tfsdk:"member_count" json:"memberCount,computed"`
	MembersContext  types.String                                                    `tfsdk:"members_context" json:"membersContext,computed"`
	Owner           customfield.NestedObject[XListListOwnerDataSourceModel]         `tfsdk:"owner" json:"owner,computed"`
	SubscriberCount types.Int64                                                     `tfsdk:"subscriber_count" json:"subscriberCount,computed"`
	URL             types.String                                                    `tfsdk:"url" json:"url,computed"`
}

type XListListCustomBannerDataSourceModel struct {
	URL         types.String                                                              `tfsdk:"url" json:"url,computed"`
	Height      types.Int64                                                               `tfsdk:"height" json:"height,computed"`
	MediaID     types.String                                                              `tfsdk:"media_id" json:"mediaId,computed"`
	MediaKey    types.String                                                              `tfsdk:"media_key" json:"mediaKey,computed"`
	Palette     customfield.NestedObjectList[XListListCustomBannerPaletteDataSourceModel] `tfsdk:"palette" json:"palette,computed"`
	SalientRect customfield.NestedObject[XListListCustomBannerSalientRectDataSourceModel] `tfsdk:"salient_rect" json:"salientRect,computed"`
	Width       types.Int64                                                               `tfsdk:"width" json:"width,computed"`
}

type XListListCustomBannerPaletteDataSourceModel struct {
	Percentage types.Float64                                                            `tfsdk:"percentage" json:"percentage,computed"`
	Rgb        customfield.NestedObject[XListListCustomBannerPaletteRgbDataSourceModel] `tfsdk:"rgb" json:"rgb,computed"`
}

type XListListCustomBannerPaletteRgbDataSourceModel struct {
	Blue  types.Int64 `tfsdk:"blue" json:"blue,computed"`
	Green types.Int64 `tfsdk:"green" json:"green,computed"`
	Red   types.Int64 `tfsdk:"red" json:"red,computed"`
}

type XListListCustomBannerSalientRectDataSourceModel struct {
	Height types.Int64 `tfsdk:"height" json:"height,computed"`
	Left   types.Int64 `tfsdk:"left" json:"left,computed"`
	Top    types.Int64 `tfsdk:"top" json:"top,computed"`
	Width  types.Int64 `tfsdk:"width" json:"width,computed"`
}

type XListListDefaultBannerDataSourceModel struct {
	URL         types.String                                                               `tfsdk:"url" json:"url,computed"`
	Height      types.Int64                                                                `tfsdk:"height" json:"height,computed"`
	MediaID     types.String                                                               `tfsdk:"media_id" json:"mediaId,computed"`
	MediaKey    types.String                                                               `tfsdk:"media_key" json:"mediaKey,computed"`
	Palette     customfield.NestedObjectList[XListListDefaultBannerPaletteDataSourceModel] `tfsdk:"palette" json:"palette,computed"`
	SalientRect customfield.NestedObject[XListListDefaultBannerSalientRectDataSourceModel] `tfsdk:"salient_rect" json:"salientRect,computed"`
	Width       types.Int64                                                                `tfsdk:"width" json:"width,computed"`
}

type XListListDefaultBannerPaletteDataSourceModel struct {
	Percentage types.Float64                                                             `tfsdk:"percentage" json:"percentage,computed"`
	Rgb        customfield.NestedObject[XListListDefaultBannerPaletteRgbDataSourceModel] `tfsdk:"rgb" json:"rgb,computed"`
}

type XListListDefaultBannerPaletteRgbDataSourceModel struct {
	Blue  types.Int64 `tfsdk:"blue" json:"blue,computed"`
	Green types.Int64 `tfsdk:"green" json:"green,computed"`
	Red   types.Int64 `tfsdk:"red" json:"red,computed"`
}

type XListListDefaultBannerSalientRectDataSourceModel struct {
	Height types.Int64 `tfsdk:"height" json:"height,computed"`
	Left   types.Int64 `tfsdk:"left" json:"left,computed"`
	Top    types.Int64 `tfsdk:"top" json:"top,computed"`
	Width  types.Int64 `tfsdk:"width" json:"width,computed"`
}

type XListListOwnerDataSourceModel struct {
	ID                              types.String                                                                      `tfsdk:"id" json:"id,computed"`
	Name                            types.String                                                                      `tfsdk:"name" json:"name,computed"`
	Username                        types.String                                                                      `tfsdk:"username" json:"username,computed"`
	AccountBasedIn                  customfield.NestedObject[XListListOwnerAccountBasedInDataSourceModel]             `tfsdk:"account_based_in" json:"accountBasedIn,computed"`
	AccountBasedInAccurate          types.Bool                                                                        `tfsdk:"account_based_in_accurate" json:"accountBasedInAccurate,computed"`
	AccountBasedInUnavailable       types.Bool                                                                        `tfsdk:"account_based_in_unavailable" json:"accountBasedInUnavailable,computed"`
	AdvertiserAccountServiceLevels  customfield.List[types.String]                                                    `tfsdk:"advertiser_account_service_levels" json:"advertiserAccountServiceLevels,computed"`
	AdvertiserAccountType           types.String                                                                      `tfsdk:"advertiser_account_type" json:"advertiserAccountType,computed"`
	AffiliateUsername               types.String                                                                      `tfsdk:"affiliate_username" json:"affiliateUsername,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XListListOwnerAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	AutomatedBy                     types.String                                                                      `tfsdk:"automated_by" json:"automatedBy,computed"`
	Birthdate                       customfield.NestedObject[XListListOwnerBirthdateDataSourceModel]                  `tfsdk:"birthdate" json:"birthdate,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                       `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	BusinessProfileState            types.String                                                                      `tfsdk:"business_profile_state" json:"businessProfileState,computed"`
	CommunityRole                   types.String                                                                      `tfsdk:"community_role" json:"communityRole,computed"`
	ConnectedVia                    types.String                                                                      `tfsdk:"connected_via" json:"connectedVia,computed"`
	ConnectedViaAccurate            types.Bool                                                                        `tfsdk:"connected_via_accurate" json:"connectedViaAccurate,computed"`
	ContactInfo                     customfield.NestedObject[XListListOwnerContactInfoDataSourceModel]                `tfsdk:"contact_info" json:"contactInfo,computed"`
	CoverPicture                    types.String                                                                      `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                                      `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                                       `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	DefaultProfile                  types.Bool                                                                        `tfsdk:"default_profile" json:"defaultProfile,computed"`
	DefaultProfileImage             types.Bool                                                                        `tfsdk:"default_profile_image" json:"defaultProfileImage,computed"`
	Description                     types.String                                                                      `tfsdk:"description" json:"description,computed"`
	Email                           types.String                                                                      `tfsdk:"email" json:"email,computed"`
	FastFollowersCount              types.Int64                                                                       `tfsdk:"fast_followers_count" json:"fastFollowersCount,computed"`
	FavouritesCount                 types.Int64                                                                       `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                                       `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                                       `tfsdk:"following" json:"following,computed"`
	GeoEnabled                      types.Bool                                                                        `tfsdk:"geo_enabled" json:"geoEnabled,computed"`
	GrokTranslatedBio               customfield.Map[jsontypes.Normalized]                                             `tfsdk:"grok_translated_bio" json:"grokTranslatedBio,computed"`
	HasCustomTimelines              types.Bool                                                                        `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasExtendedProfile              types.Bool                                                                        `tfsdk:"has_extended_profile" json:"hasExtendedProfile,computed"`
	HasGraduatedAccess              types.Bool                                                                        `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                        `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	HighlightsInfo                  customfield.NestedObject[XListListOwnerHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XListListOwnerIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
	IsAutomated                     types.Bool                                                                        `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                                        `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                                        `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslationEnabled            types.Bool                                                                        `tfsdk:"is_translation_enabled" json:"isTranslationEnabled,computed"`
	IsTranslator                    types.Bool                                                                        `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                                        `tfsdk:"is_verified" json:"isVerified,computed"`
	ListedCount                     types.Int64                                                                       `tfsdk:"listed_count" json:"listedCount,computed"`
	Location                        types.String                                                                      `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                                       `tfsdk:"media_count" json:"mediaCount,computed"`
	NormalFollowersCount            types.Int64                                                                       `tfsdk:"normal_followers_count" json:"normalFollowersCount,computed"`
	ParodyCommentaryFanLabel        types.String                                                                      `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                    `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	PossiblySensitive               types.Bool                                                                        `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	Professional                    customfield.Map[jsontypes.Normalized]                                             `tfsdk:"professional" json:"professional,computed"`
	ProfileBannerURL                types.String                                                                      `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                                      `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                                      `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                                      `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                                      `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfilePictureOriginal          types.String                                                                      `tfsdk:"profile_picture_original" json:"profilePictureOriginal,computed"`
	ProfileSortEnabled              types.Bool                                                                        `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                                      `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                             `tfsdk:"profile_bio" json:"profile_bio,computed"`
	Protected                       types.Bool                                                                        `tfsdk:"protected" json:"protected,computed"`
	RedirectsToUsername             types.String                                                                      `tfsdk:"redirects_to_username" json:"redirectsToUsername,computed"`
	RetweetedAt                     timetypes.RFC3339                                                                 `tfsdk:"retweeted_at" json:"retweetedAt,computed" format:"date-time"`
	StatusesCount                   types.Int64                                                                       `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SubscribersCount                types.Int64                                                                       `tfsdk:"subscribers_count" json:"subscribersCount,computed"`
	SuperFollowEligible             types.Bool                                                                        `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	SuperFollowsUserProfileActive   types.Bool                                                                        `tfsdk:"super_follows_user_profile_active" json:"superFollowsUserProfileActive,computed"`
	TipJar                          customfield.NestedObject[XListListOwnerTipJarDataSourceModel]                     `tfsdk:"tip_jar" json:"tipJar,computed"`
	Unavailable                     types.Bool                                                                        `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                                      `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                                      `tfsdk:"url" json:"url,computed"`
	UsernameChanges                 customfield.NestedObject[XListListOwnerUsernameChangesDataSourceModel]            `tfsdk:"username_changes" json:"usernameChanges,computed"`
	VerificationReason              customfield.NestedObject[XListListOwnerVerificationReasonDataSourceModel]         `tfsdk:"verification_reason" json:"verificationReason,computed"`
	Verified                        types.Bool                                                                        `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                                      `tfsdk:"verified_type" json:"verifiedType,computed"`
	WithheldInCountries             customfield.List[types.String]                                                    `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
	WithheldScope                   types.String                                                                      `tfsdk:"withheld_scope" json:"withheldScope,computed"`
}

type XListListOwnerAccountBasedInDataSourceModel struct {
	Level      types.String      `tfsdk:"level" json:"level,computed"`
	ObservedAt timetypes.RFC3339 `tfsdk:"observed_at" json:"observedAt,computed" format:"date-time"`
	Value      types.String      `tfsdk:"value" json:"value,computed"`
}

type XListListOwnerAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String                          `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String                          `tfsdk:"description" json:"description,computed"`
	LongDescription      customfield.Map[jsontypes.Normalized] `tfsdk:"long_description" json:"longDescription,computed"`
	URL                  types.String                          `tfsdk:"url" json:"url,computed"`
	URLType              types.String                          `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String                          `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String                          `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XListListOwnerBirthdateDataSourceModel struct {
	Day   types.Int64 `tfsdk:"day" json:"day,computed"`
	Month types.Int64 `tfsdk:"month" json:"month,computed"`
	Year  types.Int64 `tfsdk:"year" json:"year,computed"`
}

type XListListOwnerContactInfoDataSourceModel struct {
	Emails customfield.NestedObjectList[XListListOwnerContactInfoEmailsDataSourceModel] `tfsdk:"emails" json:"emails,computed"`
	Links  customfield.NestedObjectList[XListListOwnerContactInfoLinksDataSourceModel]  `tfsdk:"links" json:"links,computed"`
	Phones customfield.NestedObjectList[XListListOwnerContactInfoPhonesDataSourceModel] `tfsdk:"phones" json:"phones,computed"`
}

type XListListOwnerContactInfoEmailsDataSourceModel struct {
	Domain    types.String `tfsdk:"domain" json:"domain,computed"`
	Source    types.String `tfsdk:"source" json:"source,computed"`
	Value     types.String `tfsdk:"value" json:"value,computed"`
	WrittenAs types.String `tfsdk:"written_as" json:"writtenAs,computed"`
}

type XListListOwnerContactInfoLinksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
	Source  types.String `tfsdk:"source" json:"source,computed"`
	URL     types.String `tfsdk:"url" json:"url,computed"`
}

type XListListOwnerContactInfoPhonesDataSourceModel struct {
	Source    types.String `tfsdk:"source" json:"source,computed"`
	Value     types.String `tfsdk:"value" json:"value,computed"`
	WrittenAs types.String `tfsdk:"written_as" json:"writtenAs,computed"`
}

type XListListOwnerHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XListListOwnerIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}

type XListListOwnerTipJarDataSourceModel struct {
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

type XListListOwnerUsernameChangesDataSourceModel struct {
	Count         types.Int64       `tfsdk:"count" json:"count,computed"`
	LastChangedAt timetypes.RFC3339 `tfsdk:"last_changed_at" json:"lastChangedAt,computed" format:"date-time"`
}

type XListListOwnerVerificationReasonDataSourceModel struct {
	Entities customfield.NestedObjectList[XListListOwnerVerificationReasonEntitiesDataSourceModel] `tfsdk:"entities" json:"entities,computed"`
	Text     types.String                                                                          `tfsdk:"text" json:"text,computed"`
}

type XListListOwnerVerificationReasonEntitiesDataSourceModel struct {
	FromIndex types.Int64  `tfsdk:"from_index" json:"fromIndex,computed"`
	ToIndex   types.Int64  `tfsdk:"to_index" json:"toIndex,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	URLType   types.String `tfsdk:"url_type" json:"urlType,computed"`
}
