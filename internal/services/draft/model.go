// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package draft

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type DraftModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Text      types.String      `tfsdk:"text" json:"text,required"`
	Goal      types.String      `tfsdk:"goal" json:"goal,optional"`
	Topic     types.String      `tfsdk:"topic" json:"topic,optional"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
}

func (m DraftModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DraftModel) MarshalJSONForUpdate(state DraftModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
