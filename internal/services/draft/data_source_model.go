// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package draft

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DraftDataSourceModel struct {
	ID        types.String      `tfsdk:"id" path:"id,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	Goal      types.String      `tfsdk:"goal" json:"goal" api:"computed"`
	Text      types.String      `tfsdk:"text" json:"text" api:"computed"`
	Topic     types.String      `tfsdk:"topic" json:"topic" api:"computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updatedAt" api:"computed" format:"date-time"`
}
