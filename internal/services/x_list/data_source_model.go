// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xlist

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XListDataSourceModel struct {
	ID   types.String                                       `tfsdk:"id" path:"id,required"`
	List customfield.NestedObject[XListListDataSourceModel] `tfsdk:"list" json:"list" api:"computed"`
}

type XListListDataSourceModel struct {
	ID              types.String                                                    `tfsdk:"id" json:"id" api:"computed"`
	Name            types.String                                                    `tfsdk:"name" json:"name" api:"computed"`
	BannerURL       types.String                                                    `tfsdk:"banner_url" json:"bannerUrl" api:"computed"`
	CreatedAt       timetypes.RFC3339                                               `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	CustomBanner    customfield.NestedObject[XListListCustomBannerDataSourceModel]  `tfsdk:"custom_banner" json:"customBanner" api:"computed"`
	DefaultBanner   customfield.NestedObject[XListListDefaultBannerDataSourceModel] `tfsdk:"default_banner" json:"defaultBanner" api:"computed"`
	Description     types.String                                                    `tfsdk:"description" json:"description" api:"computed"`
	FacepileURLs    customfield.List[types.String]                                  `tfsdk:"facepile_urls" json:"facepileUrls" api:"computed"`
	MemberCount     types.Int64                                                     `tfsdk:"member_count" json:"memberCount" api:"computed"`
	MembersContext  types.String                                                    `tfsdk:"members_context" json:"membersContext" api:"computed"`
	Owner           customfield.NestedObject[XListListOwnerDataSourceModel]         `tfsdk:"owner" json:"owner" api:"computed"`
	SubscriberCount types.Int64                                                     `tfsdk:"subscriber_count" json:"subscriberCount" api:"computed"`
	URL             types.String                                                    `tfsdk:"url" json:"url" api:"computed"`
}

type XListListCustomBannerDataSourceModel struct {
	URL         types.String                                                              `tfsdk:"url" json:"url" api:"computed"`
	Height      types.Int64                                                               `tfsdk:"height" json:"height" api:"computed"`
	MediaID     types.String                                                              `tfsdk:"media_id" json:"mediaId" api:"computed"`
	MediaKey    types.String                                                              `tfsdk:"media_key" json:"mediaKey" api:"computed"`
	Palette     customfield.NestedObjectList[XListListCustomBannerPaletteDataSourceModel] `tfsdk:"palette" json:"palette" api:"computed"`
	SalientRect customfield.NestedObject[XListListCustomBannerSalientRectDataSourceModel] `tfsdk:"salient_rect" json:"salientRect" api:"computed"`
	Width       types.Int64                                                               `tfsdk:"width" json:"width" api:"computed"`
}

type XListListCustomBannerPaletteDataSourceModel struct {
	Percentage types.Float64                                                            `tfsdk:"percentage" json:"percentage" api:"computed"`
	Rgb        customfield.NestedObject[XListListCustomBannerPaletteRgbDataSourceModel] `tfsdk:"rgb" json:"rgb" api:"computed"`
}

type XListListCustomBannerPaletteRgbDataSourceModel struct {
	Blue  types.Int64 `tfsdk:"blue" json:"blue" api:"computed"`
	Green types.Int64 `tfsdk:"green" json:"green" api:"computed"`
	Red   types.Int64 `tfsdk:"red" json:"red" api:"computed"`
}

type XListListCustomBannerSalientRectDataSourceModel struct {
	Height types.Int64 `tfsdk:"height" json:"height" api:"computed"`
	Left   types.Int64 `tfsdk:"left" json:"left" api:"computed"`
	Top    types.Int64 `tfsdk:"top" json:"top" api:"computed"`
	Width  types.Int64 `tfsdk:"width" json:"width" api:"computed"`
}

type XListListDefaultBannerDataSourceModel struct {
	URL         types.String                                                               `tfsdk:"url" json:"url" api:"computed"`
	Height      types.Int64                                                                `tfsdk:"height" json:"height" api:"computed"`
	MediaID     types.String                                                               `tfsdk:"media_id" json:"mediaId" api:"computed"`
	MediaKey    types.String                                                               `tfsdk:"media_key" json:"mediaKey" api:"computed"`
	Palette     customfield.NestedObjectList[XListListDefaultBannerPaletteDataSourceModel] `tfsdk:"palette" json:"palette" api:"computed"`
	SalientRect customfield.NestedObject[XListListDefaultBannerSalientRectDataSourceModel] `tfsdk:"salient_rect" json:"salientRect" api:"computed"`
	Width       types.Int64                                                                `tfsdk:"width" json:"width" api:"computed"`
}

type XListListDefaultBannerPaletteDataSourceModel struct {
	Percentage types.Float64                                                             `tfsdk:"percentage" json:"percentage" api:"computed"`
	Rgb        customfield.NestedObject[XListListDefaultBannerPaletteRgbDataSourceModel] `tfsdk:"rgb" json:"rgb" api:"computed"`
}

type XListListDefaultBannerPaletteRgbDataSourceModel struct {
	Blue  types.Int64 `tfsdk:"blue" json:"blue" api:"computed"`
	Green types.Int64 `tfsdk:"green" json:"green" api:"computed"`
	Red   types.Int64 `tfsdk:"red" json:"red" api:"computed"`
}

type XListListDefaultBannerSalientRectDataSourceModel struct {
	Height types.Int64 `tfsdk:"height" json:"height" api:"computed"`
	Left   types.Int64 `tfsdk:"left" json:"left" api:"computed"`
	Top    types.Int64 `tfsdk:"top" json:"top" api:"computed"`
	Width  types.Int64 `tfsdk:"width" json:"width" api:"computed"`
}

type XListListOwnerDataSourceModel struct {
	ID                              types.String                                                                      `tfsdk:"id" json:"id" api:"computed"`
	Name                            types.String                                                                      `tfsdk:"name" json:"name" api:"computed"`
	Username                        types.String                                                                      `tfsdk:"username" json:"username" api:"computed"`
	AccountBasedIn                  customfield.NestedObject[XListListOwnerAccountBasedInDataSourceModel]             `tfsdk:"account_based_in" json:"accountBasedIn" api:"computed"`
	AccountBasedInAccurate          types.Bool                                                                        `tfsdk:"account_based_in_accurate" json:"accountBasedInAccurate" api:"computed"`
	AccountBasedInUnavailable       types.Bool                                                                        `tfsdk:"account_based_in_unavailable" json:"accountBasedInUnavailable" api:"computed"`
	AdvertiserAccountServiceLevels  customfield.List[types.String]                                                    `tfsdk:"advertiser_account_service_levels" json:"advertiserAccountServiceLevels" api:"computed"`
	AdvertiserAccountType           types.String                                                                      `tfsdk:"advertiser_account_type" json:"advertiserAccountType" api:"computed"`
	AffiliateUsername               types.String                                                                      `tfsdk:"affiliate_username" json:"affiliateUsername" api:"computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XListListOwnerAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel" api:"computed"`
	AutomatedBy                     types.String                                                                      `tfsdk:"automated_by" json:"automatedBy" api:"computed"`
	Birthdate                       customfield.NestedObject[XListListOwnerBirthdateDataSourceModel]                  `tfsdk:"birthdate" json:"birthdate" api:"computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                                       `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount" api:"computed"`
	BusinessProfileState            types.String                                                                      `tfsdk:"business_profile_state" json:"businessProfileState" api:"computed"`
	CommunityRole                   types.String                                                                      `tfsdk:"community_role" json:"communityRole" api:"computed"`
	ConnectedVia                    types.String                                                                      `tfsdk:"connected_via" json:"connectedVia" api:"computed"`
	ConnectedViaAccurate            types.Bool                                                                        `tfsdk:"connected_via_accurate" json:"connectedViaAccurate" api:"computed"`
	ContactInfo                     customfield.NestedObject[XListListOwnerContactInfoDataSourceModel]                `tfsdk:"contact_info" json:"contactInfo" api:"computed"`
	CoverPicture                    types.String                                                                      `tfsdk:"cover_picture" json:"coverPicture" api:"computed"`
	CreatedAt                       types.String                                                                      `tfsdk:"created_at" json:"createdAt" api:"computed"`
	CreatorSubscriptionsCount       types.Int64                                                                       `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount" api:"computed"`
	DefaultProfile                  types.Bool                                                                        `tfsdk:"default_profile" json:"defaultProfile" api:"computed"`
	DefaultProfileImage             types.Bool                                                                        `tfsdk:"default_profile_image" json:"defaultProfileImage" api:"computed"`
	Description                     types.String                                                                      `tfsdk:"description" json:"description" api:"computed"`
	Email                           types.String                                                                      `tfsdk:"email" json:"email" api:"computed"`
	FastFollowersCount              types.Int64                                                                       `tfsdk:"fast_followers_count" json:"fastFollowersCount" api:"computed"`
	FavouritesCount                 types.Int64                                                                       `tfsdk:"favourites_count" json:"favouritesCount" api:"computed"`
	Followers                       types.Int64                                                                       `tfsdk:"followers" json:"followers" api:"computed"`
	Following                       types.Int64                                                                       `tfsdk:"following" json:"following" api:"computed"`
	GeoEnabled                      types.Bool                                                                        `tfsdk:"geo_enabled" json:"geoEnabled" api:"computed"`
	GrokTranslatedBio               customfield.Map[jsontypes.Normalized]                                             `tfsdk:"grok_translated_bio" json:"grokTranslatedBio" api:"computed"`
	HasCustomTimelines              types.Bool                                                                        `tfsdk:"has_custom_timelines" json:"hasCustomTimelines" api:"computed"`
	HasExtendedProfile              types.Bool                                                                        `tfsdk:"has_extended_profile" json:"hasExtendedProfile" api:"computed"`
	HasGraduatedAccess              types.Bool                                                                        `tfsdk:"has_graduated_access" json:"hasGraduatedAccess" api:"computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                                        `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile" api:"computed"`
	HighlightsInfo                  customfield.NestedObject[XListListOwnerHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo" api:"computed"`
	IdentityVerification            customfield.NestedObject[XListListOwnerIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification" api:"computed"`
	IsAutomated                     types.Bool                                                                        `tfsdk:"is_automated" json:"isAutomated" api:"computed"`
	IsBlueVerified                  types.Bool                                                                        `tfsdk:"is_blue_verified" json:"isBlueVerified" api:"computed"`
	IsProfileTranslatable           types.Bool                                                                        `tfsdk:"is_profile_translatable" json:"isProfileTranslatable" api:"computed"`
	IsTranslationEnabled            types.Bool                                                                        `tfsdk:"is_translation_enabled" json:"isTranslationEnabled" api:"computed"`
	IsTranslator                    types.Bool                                                                        `tfsdk:"is_translator" json:"isTranslator" api:"computed"`
	IsVerified                      types.Bool                                                                        `tfsdk:"is_verified" json:"isVerified" api:"computed"`
	ListedCount                     types.Int64                                                                       `tfsdk:"listed_count" json:"listedCount" api:"computed"`
	Location                        types.String                                                                      `tfsdk:"location" json:"location" api:"computed"`
	MediaCount                      types.Int64                                                                       `tfsdk:"media_count" json:"mediaCount" api:"computed"`
	NormalFollowersCount            types.Int64                                                                       `tfsdk:"normal_followers_count" json:"normalFollowersCount" api:"computed"`
	ParodyCommentaryFanLabel        types.String                                                                      `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel" api:"computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                                    `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds" api:"computed"`
	PossiblySensitive               types.Bool                                                                        `tfsdk:"possibly_sensitive" json:"possiblySensitive" api:"computed"`
	Professional                    customfield.Map[jsontypes.Normalized]                                             `tfsdk:"professional" json:"professional" api:"computed"`
	ProfileBannerURL                types.String                                                                      `tfsdk:"profile_banner_url" json:"profileBannerUrl" api:"computed"`
	ProfileDescriptionLanguage      types.String                                                                      `tfsdk:"profile_description_language" json:"profileDescriptionLanguage" api:"computed"`
	ProfileImageShape               types.String                                                                      `tfsdk:"profile_image_shape" json:"profileImageShape" api:"computed"`
	ProfileInterstitialType         types.String                                                                      `tfsdk:"profile_interstitial_type" json:"profileInterstitialType" api:"computed"`
	ProfilePicture                  types.String                                                                      `tfsdk:"profile_picture" json:"profilePicture" api:"computed"`
	ProfilePictureOriginal          types.String                                                                      `tfsdk:"profile_picture_original" json:"profilePictureOriginal" api:"computed"`
	ProfileSortEnabled              types.Bool                                                                        `tfsdk:"profile_sort_enabled" json:"profileSortEnabled" api:"computed"`
	ProfileTranslatorType           types.String                                                                      `tfsdk:"profile_translator_type" json:"profileTranslatorType" api:"computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                             `tfsdk:"profile_bio" json:"profile_bio" api:"computed"`
	Protected                       types.Bool                                                                        `tfsdk:"protected" json:"protected" api:"computed"`
	RedirectsToUsername             types.String                                                                      `tfsdk:"redirects_to_username" json:"redirectsToUsername" api:"computed"`
	RetweetedAt                     timetypes.RFC3339                                                                 `tfsdk:"retweeted_at" json:"retweetedAt" api:"computed" format:"date-time"`
	StatusesCount                   types.Int64                                                                       `tfsdk:"statuses_count" json:"statusesCount" api:"computed"`
	SubscribersCount                types.Int64                                                                       `tfsdk:"subscribers_count" json:"subscribersCount" api:"computed"`
	SuperFollowEligible             types.Bool                                                                        `tfsdk:"super_follow_eligible" json:"superFollowEligible" api:"computed"`
	SuperFollowsUserProfileActive   types.Bool                                                                        `tfsdk:"super_follows_user_profile_active" json:"superFollowsUserProfileActive" api:"computed"`
	TipJar                          customfield.NestedObject[XListListOwnerTipJarDataSourceModel]                     `tfsdk:"tip_jar" json:"tipJar" api:"computed"`
	Unavailable                     types.Bool                                                                        `tfsdk:"unavailable" json:"unavailable" api:"computed"`
	UnavailableReason               types.String                                                                      `tfsdk:"unavailable_reason" json:"unavailableReason" api:"computed"`
	URL                             types.String                                                                      `tfsdk:"url" json:"url" api:"computed"`
	UsernameChanges                 customfield.NestedObject[XListListOwnerUsernameChangesDataSourceModel]            `tfsdk:"username_changes" json:"usernameChanges" api:"computed"`
	VerificationReason              customfield.NestedObject[XListListOwnerVerificationReasonDataSourceModel]         `tfsdk:"verification_reason" json:"verificationReason" api:"computed"`
	Verified                        types.Bool                                                                        `tfsdk:"verified" json:"verified" api:"computed"`
	VerifiedType                    types.String                                                                      `tfsdk:"verified_type" json:"verifiedType" api:"computed"`
	WithheldInCountries             customfield.List[types.String]                                                    `tfsdk:"withheld_in_countries" json:"withheldInCountries" api:"computed"`
	WithheldScope                   types.String                                                                      `tfsdk:"withheld_scope" json:"withheldScope" api:"computed"`
}

type XListListOwnerAccountBasedInDataSourceModel struct {
	Level      types.String      `tfsdk:"level" json:"level" api:"computed"`
	ObservedAt timetypes.RFC3339 `tfsdk:"observed_at" json:"observedAt" api:"computed" format:"date-time"`
	Value      types.String      `tfsdk:"value" json:"value" api:"computed"`
}

type XListListOwnerAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String                          `tfsdk:"badge_url" json:"badgeUrl" api:"computed"`
	Description          types.String                          `tfsdk:"description" json:"description" api:"computed"`
	LongDescription      customfield.Map[jsontypes.Normalized] `tfsdk:"long_description" json:"longDescription" api:"computed"`
	URL                  types.String                          `tfsdk:"url" json:"url" api:"computed"`
	URLType              types.String                          `tfsdk:"url_type" json:"urlType" api:"computed"`
	UserLabelDisplayType types.String                          `tfsdk:"user_label_display_type" json:"userLabelDisplayType" api:"computed"`
	UserLabelType        types.String                          `tfsdk:"user_label_type" json:"userLabelType" api:"computed"`
}

type XListListOwnerBirthdateDataSourceModel struct {
	Day   types.Int64 `tfsdk:"day" json:"day" api:"computed"`
	Month types.Int64 `tfsdk:"month" json:"month" api:"computed"`
	Year  types.Int64 `tfsdk:"year" json:"year" api:"computed"`
}

type XListListOwnerContactInfoDataSourceModel struct {
	Emails customfield.NestedObjectList[XListListOwnerContactInfoEmailsDataSourceModel] `tfsdk:"emails" json:"emails" api:"computed"`
	Links  customfield.NestedObjectList[XListListOwnerContactInfoLinksDataSourceModel]  `tfsdk:"links" json:"links" api:"computed"`
	Phones customfield.NestedObjectList[XListListOwnerContactInfoPhonesDataSourceModel] `tfsdk:"phones" json:"phones" api:"computed"`
}

type XListListOwnerContactInfoEmailsDataSourceModel struct {
	Domain    types.String `tfsdk:"domain" json:"domain" api:"computed"`
	Source    types.String `tfsdk:"source" json:"source" api:"computed"`
	Value     types.String `tfsdk:"value" json:"value" api:"computed"`
	WrittenAs types.String `tfsdk:"written_as" json:"writtenAs" api:"computed"`
}

type XListListOwnerContactInfoLinksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network" api:"computed"`
	Source  types.String `tfsdk:"source" json:"source" api:"computed"`
	URL     types.String `tfsdk:"url" json:"url" api:"computed"`
}

type XListListOwnerContactInfoPhonesDataSourceModel struct {
	Source    types.String `tfsdk:"source" json:"source" api:"computed"`
	Value     types.String `tfsdk:"value" json:"value" api:"computed"`
	WrittenAs types.String `tfsdk:"written_as" json:"writtenAs" api:"computed"`
}

type XListListOwnerHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets" api:"computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets" api:"computed"`
}

type XListListOwnerIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description" api:"computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified" api:"computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec" api:"computed"`
}

type XListListOwnerTipJarDataSourceModel struct {
	BandcampHandle types.String `tfsdk:"bandcamp_handle" json:"bandcampHandle" api:"computed"`
	BitcoinHandle  types.String `tfsdk:"bitcoin_handle" json:"bitcoinHandle" api:"computed"`
	CashAppHandle  types.String `tfsdk:"cash_app_handle" json:"cashAppHandle" api:"computed"`
	EthereumHandle types.String `tfsdk:"ethereum_handle" json:"ethereumHandle" api:"computed"`
	GofundmeHandle types.String `tfsdk:"gofundme_handle" json:"gofundmeHandle" api:"computed"`
	IsEnabled      types.Bool   `tfsdk:"is_enabled" json:"isEnabled" api:"computed"`
	PatreonHandle  types.String `tfsdk:"patreon_handle" json:"patreonHandle" api:"computed"`
	PayPalHandle   types.String `tfsdk:"pay_pal_handle" json:"payPalHandle" api:"computed"`
	VenmoHandle    types.String `tfsdk:"venmo_handle" json:"venmoHandle" api:"computed"`
}

type XListListOwnerUsernameChangesDataSourceModel struct {
	Count         types.Int64       `tfsdk:"count" json:"count" api:"computed"`
	LastChangedAt timetypes.RFC3339 `tfsdk:"last_changed_at" json:"lastChangedAt" api:"computed" format:"date-time"`
}

type XListListOwnerVerificationReasonDataSourceModel struct {
	Entities customfield.NestedObjectList[XListListOwnerVerificationReasonEntitiesDataSourceModel] `tfsdk:"entities" json:"entities" api:"computed"`
	Text     types.String                                                                          `tfsdk:"text" json:"text" api:"computed"`
}

type XListListOwnerVerificationReasonEntitiesDataSourceModel struct {
	FromIndex types.Int64  `tfsdk:"from_index" json:"fromIndex" api:"computed"`
	ToIndex   types.Int64  `tfsdk:"to_index" json:"toIndex" api:"computed"`
	URL       types.String `tfsdk:"url" json:"url" api:"computed"`
	URLType   types.String `tfsdk:"url_type" json:"urlType" api:"computed"`
}
