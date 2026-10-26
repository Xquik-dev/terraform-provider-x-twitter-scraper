// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package draft

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DraftModel struct {
	ID        types.String      `tfsdk:"id" json:"id" api:"computed"`
	Text      types.String      `tfsdk:"text" json:"text" api:"required"`
	Goal      types.String      `tfsdk:"goal" json:"goal" api:"optional"`
	Topic     types.String      `tfsdk:"topic" json:"topic" api:"optional"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updatedAt" api:"computed" format:"date-time"`
}

func (m DraftModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DraftModel) MarshalJSONForUpdate(state DraftModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
