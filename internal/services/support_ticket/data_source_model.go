// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package support_ticket

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type SupportTicketDataSourceModel struct {
	ID        types.String                                                       `tfsdk:"id" path:"id,required"`
	CreatedAt timetypes.RFC3339                                                  `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	PublicID  types.String                                                       `tfsdk:"public_id" json:"publicId,computed"`
	Status    types.String                                                       `tfsdk:"status" json:"status,computed"`
	Subject   types.String                                                       `tfsdk:"subject" json:"subject,computed"`
	UpdatedAt timetypes.RFC3339                                                  `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
	Messages  customfield.NestedObjectList[SupportTicketMessagesDataSourceModel] `tfsdk:"messages" json:"messages,computed"`
}

type SupportTicketMessagesDataSourceModel struct {
	Body      types.String      `tfsdk:"body" json:"body,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Sender    types.String      `tfsdk:"sender" json:"sender,computed"`
}
