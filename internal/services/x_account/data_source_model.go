// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xaccount

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XAccountDataSourceModel struct {
	ID        types.String      `tfsdk:"id" path:"id,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	Health    types.String      `tfsdk:"health" json:"health" api:"computed"`
	Status    types.String      `tfsdk:"status" json:"status" api:"computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updatedAt" api:"computed" format:"date-time"`
	XUserID   types.String      `tfsdk:"x_user_id" json:"xUserId" api:"computed"`
	XUsername types.String      `tfsdk:"x_username" json:"xUsername" api:"computed"`
}
