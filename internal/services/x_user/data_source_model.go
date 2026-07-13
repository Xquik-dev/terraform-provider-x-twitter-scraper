// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XUserDataSourceModel struct {
	ID                  types.String                          `tfsdk:"id" path:"id,required"`
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
	Name                types.String                          `tfsdk:"name" json:"name,computed"`
	PossiblySensitive   types.Bool                            `tfsdk:"possibly_sensitive" json:"possiblySensitive,computed"`
	ProfileBannerURL    types.String                          `tfsdk:"profile_banner_url" json:"profileBannerUrl,computed"`
	ProfilePicture      types.String                          `tfsdk:"profile_picture" json:"profilePicture,computed"`
	Protected           types.Bool                            `tfsdk:"protected" json:"protected,computed"`
	StatusesCount       types.Int64                           `tfsdk:"statuses_count" json:"statusesCount,computed"`
	Unavailable         types.Bool                            `tfsdk:"unavailable" json:"unavailable,computed"`
	UnavailableReason   types.String                          `tfsdk:"unavailable_reason" json:"unavailableReason,computed"`
	URL                 types.String                          `tfsdk:"url" json:"url,computed"`
	Username            types.String                          `tfsdk:"username" json:"username,computed"`
	Verified            types.Bool                            `tfsdk:"verified" json:"verified,computed"`
	VerifiedType        types.String                          `tfsdk:"verified_type" json:"verifiedType,computed"`
	ViewerFollowedBy    types.Bool                            `tfsdk:"viewer_followed_by" json:"viewerFollowedBy,computed"`
	ViewerFollowing     types.Bool                            `tfsdk:"viewer_following" json:"viewerFollowing,computed"`
	PinnedTweetIDs      customfield.List[types.String]        `tfsdk:"pinned_tweet_ids" json:"pinnedTweetIds,computed"`
	ProfileBio          customfield.Map[jsontypes.Normalized] `tfsdk:"profile_bio" json:"profile_bio,computed"`
	WithheldInCountries customfield.List[types.String]        `tfsdk:"withheld_in_countries" json:"withheldInCountries,computed"`
}
