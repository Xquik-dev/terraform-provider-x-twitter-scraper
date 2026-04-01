// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package support_ticket

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SupportTicketModel struct {
	ID        types.String                                             `tfsdk:"id" path:"id,optional"`
	Body      types.String                                             `tfsdk:"body" json:"body,required,no_refresh"`
	Subject   types.String                                             `tfsdk:"subject" json:"subject,required"`
	Status    types.String                                             `tfsdk:"status" json:"status,optional"`
	CreatedAt timetypes.RFC3339                                        `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	PublicID  types.String                                             `tfsdk:"public_id" json:"publicId,computed"`
	UpdatedAt timetypes.RFC3339                                        `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
	Messages  customfield.NestedObjectList[SupportTicketMessagesModel] `tfsdk:"messages" json:"messages,computed"`
}

func (m SupportTicketModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SupportTicketModel) MarshalJSONForUpdate(state SupportTicketModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type SupportTicketMessagesModel struct {
	Body      types.String      `tfsdk:"body" json:"body,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Sender    types.String      `tfsdk:"sender" json:"sender,computed"`
}
