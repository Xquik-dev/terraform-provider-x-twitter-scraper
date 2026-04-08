// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_user

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XUserDataSourceModel struct {
	ID             types.String `tfsdk:"id" path:"id,required"`
	CreatedAt      types.String `tfsdk:"created_at" json:"createdAt,computed"`
	Description    types.String `tfsdk:"description" json:"description,computed"`
	Followers      types.Int64  `tfsdk:"followers" json:"followers,computed"`
	Following      types.Int64  `tfsdk:"following" json:"following,computed"`
	Location       types.String `tfsdk:"location" json:"location,computed"`
	Name           types.String `tfsdk:"name" json:"name,computed"`
	ProfilePicture types.String `tfsdk:"profile_picture" json:"profilePicture,computed"`
	StatusesCount  types.Int64  `tfsdk:"statuses_count" json:"statusesCount,computed"`
	Username       types.String `tfsdk:"username" json:"username,computed"`
	Verified       types.Bool   `tfsdk:"verified" json:"verified,computed"`
}
