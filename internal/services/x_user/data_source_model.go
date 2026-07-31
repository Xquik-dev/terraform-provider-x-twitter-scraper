// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XUserDataSourceModel struct {
	ID                              types.String                                                             `tfsdk:"id" path:"id,required"`
	AutomatedBy                     types.String                                                             `tfsdk:"automated_by" json:"automatedBy,computed"`
	BusinessAccountAffiliatesCount  types.Int64                                                              `tfsdk:"business_account_affiliates_count" json:"businessAccountAffiliatesCount,computed"`
	CommunityRole                   types.String                                                             `tfsdk:"community_role" json:"communityRole,computed"`
	CoverPicture                    types.String                                                             `tfsdk:"cover_picture" json:"coverPicture,computed"`
	CreatedAt                       types.String                                                             `tfsdk:"created_at" json:"createdAt,computed"`
	CreatorSubscriptionsCount       types.Int64                                                              `tfsdk:"creator_subscriptions_count" json:"creatorSubscriptionsCount,computed"`
	Description                     types.String                                                             `tfsdk:"description" json:"description,computed"`
	FavouritesCount                 types.Int64                                                              `tfsdk:"favourites_count" json:"favouritesCount,computed"`
	Followers                       types.Int64                                                              `tfsdk:"followers" json:"followers,computed"`
	Following                       types.Int64                                                              `tfsdk:"following" json:"following,computed"`
	HasCustomTimelines              types.Bool                                                               `tfsdk:"has_custom_timelines" json:"hasCustomTimelines,computed"`
	HasGraduatedAccess              types.Bool                                                               `tfsdk:"has_graduated_access" json:"hasGraduatedAccess,computed"`
	HasHiddenSubscriptionsOnProfile types.Bool                                                               `tfsdk:"has_hidden_subscriptions_on_profile" json:"hasHiddenSubscriptionsOnProfile,computed"`
	IsAutomated                     types.Bool                                                               `tfsdk:"is_automated" json:"isAutomated,computed"`
	IsBlueVerified                  types.Bool                                                               `tfsdk:"is_blue_verified" json:"isBlueVerified,computed"`
	IsProfileTranslatable           types.Bool                                                               `tfsdk:"is_profile_translatable" json:"isProfileTranslatable,computed"`
	IsTranslator                    types.Bool                                                               `tfsdk:"is_translator" json:"isTranslator,computed"`
	IsVerified                      types.Bool                                                               `tfsdk:"is_verified" json:"isVerified,computed"`
	Location                        types.String                                                             `tfsdk:"location" json:"location,computed"`
	MediaCount                      types.Int64                                                              `tfsdk:"media_count" json:"mediaCount,computed"`
	Name                            types.String                                                             `tfsdk:"name" json:"name,computed"`
	ParodyCommentaryFanLabel        types.String                                                             `tfsdk:"parody_commentary_fan_label" json:"parodyCommentaryFanLabel,computed"`
	PossiblySensitive               types.Bool                                                               `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBannerURL                types.String                                                             `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfileDescriptionLanguage      types.String                                                             `tfsdk:"profile_description_language" json:"profileDescriptionLanguage,computed"`
	ProfileImageShape               types.String                                                             `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfileInterstitialType         types.String                                                             `tfsdk:"profile_interstitial_type" json:"profileInterstitialType,computed"`
	ProfilePicture                  types.String                                                             `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfileSortEnabled              types.Bool                                                               `tfsdk:"profile_sort_enabled" json:"profileSortEnabled,computed"`
	ProfileTranslatorType           types.String                                                             `tfsdk:"profile_translator_type" json:"profileTranslatorType,computed"`
	Protected                       types.Bool                                                               `tfsdk:"protected" json:"protected,computed"`
	StatusesCount                   types.Int64                                                              `tfsdk:"statuses_count" json:"statusesCount,computed"`
	SuperFollowEligible             types.Bool                                                               `tfsdk:"super_follow_eligible" json:"superFollowEligible,computed"`
	Unavailable                     types.Bool                                                               `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason               types.String                                                             `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                             types.String                                                             `tfsdk:"url" json:"url,computed"`
	Username                        types.String                                                             `tfsdk:"username" json:"username,computed"`
	Verified                        types.Bool                                                               `tfsdk:"verified" json:"verified,computed"`
	VerifiedType                    types.String                                                             `tfsdk:"verified_type" json:"verifiedType,computed"`
	PinnedTweetIDs                  customfield.List[types.String]                                           `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	ProfileBio                      customfield.Map[jsontypes.Normalized]                                    `tfsdk:"profile_bio" json:"profile_bio,computed"`
	WithheldInCountries             customfield.List[types.String]                                           `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
	AffiliatesHighlightedLabel      customfield.NestedObject[XUserAffiliatesHighlightedLabelDataSourceModel] `tfsdk:"affiliates_highlighted_label" json:"affiliatesHighlightedLabel,computed"`
	HighlightsInfo                  customfield.NestedObject[XUserHighlightsInfoDataSourceModel]             `tfsdk:"highlights_info" json:"highlightsInfo,computed"`
	IdentityVerification            customfield.NestedObject[XUserIdentityVerificationDataSourceModel]       `tfsdk:"identity_verification" json:"identityVerification,computed"`
}

type XUserAffiliatesHighlightedLabelDataSourceModel struct {
	BadgeURL             types.String `tfsdk:"badge_url" json:"badgeUrl,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	URL                  types.String `tfsdk:"url" json:"url,computed"`
	URLType              types.String `tfsdk:"url_type" json:"urlType,computed"`
	UserLabelDisplayType types.String `tfsdk:"user_label_display_type" json:"userLabelDisplayType,computed"`
	UserLabelType        types.String `tfsdk:"user_label_type" json:"userLabelType,computed"`
}

type XUserHighlightsInfoDataSourceModel struct {
	CanHighlightTweets types.Bool   `tfsdk:"can_highlight_tweets" json:"canHighlightTweets,computed"`
	HighlightedTweets  types.String `tfsdk:"highlighted_tweets" json:"highlightedTweets,computed"`
}

type XUserIdentityVerificationDataSourceModel struct {
	Description        types.String `tfsdk:"description" json:"description,computed"`
	IsIdentityVerified types.Bool   `tfsdk:"is_identity_verified" json:"isIdentityVerified,computed"`
	VerifiedSinceMsec  types.String `tfsdk:"verified_since_msec" json:"verifiedSinceMsec,computed"`
}
