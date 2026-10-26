// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package supportticket

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SupportTicketDataSourceModel struct {
	ID        types.String                                                       `tfsdk:"id" path:"id,required"`
	CreatedAt timetypes.RFC3339                                                  `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	PublicID  types.String                                                       `tfsdk:"public_id" json:"publicId" api:"computed"`
	Status    types.String                                                       `tfsdk:"status" json:"status" api:"computed"`
	Subject   types.String                                                       `tfsdk:"subject" json:"subject" api:"computed"`
	UpdatedAt timetypes.RFC3339                                                  `tfsdk:"updated_at" json:"updatedAt" api:"computed" format:"date-time"`
	Messages  customfield.NestedObjectList[SupportTicketMessagesDataSourceModel] `tfsdk:"messages" json:"messages" api:"computed"`
}

type SupportTicketMessagesDataSourceModel struct {
	Attachments customfield.NestedObjectList[SupportTicketMessagesAttachmentsDataSourceModel] `tfsdk:"attachments" json:"attachments" api:"computed"`
	Body        types.String                                                                  `tfsdk:"body" json:"body" api:"computed"`
	CreatedAt   timetypes.RFC3339                                                             `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	Sender      types.String                                                                  `tfsdk:"sender" json:"sender" api:"computed"`
}

type SupportTicketMessagesAttachmentsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType" api:"computed"`
	Filename    types.String `tfsdk:"filename" json:"filename" api:"computed"`
	Kind        types.String `tfsdk:"kind" json:"kind" api:"computed"`
	PublicID    types.String `tfsdk:"public_id" json:"publicId" api:"computed"`
	SizeBytes   types.Int64  `tfsdk:"size_bytes" json:"sizeBytes" api:"computed"`
	Status      types.String `tfsdk:"status" json:"status" api:"computed"`
	URL         types.String `tfsdk:"url" json:"url" api:"computed"`
}
