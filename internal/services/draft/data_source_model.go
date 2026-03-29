// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package draft

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DraftDataSourceModel struct {
	ID        types.String      `tfsdk:"id" path:"id,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Goal      types.String      `tfsdk:"goal" json:"goal,computed"`
	Text      types.String      `tfsdk:"text" json:"text,computed"`
	Topic     types.String      `tfsdk:"topic" json:"topic,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
}
