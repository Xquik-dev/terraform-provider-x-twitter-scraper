// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package support_ticket

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	Attachments customfield.NestedObjectList[SupportTicketMessagesAttachmentsDataSourceModel] `tfsdk:"attachments" json:"attachments,computed"`
	Body        types.String                                                                  `tfsdk:"body" json:"body,computed"`
	CreatedAt   timetypes.RFC3339                                                             `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Sender      types.String                                                                  `tfsdk:"sender" json:"sender,computed"`
}

type SupportTicketMessagesAttachmentsDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType,computed"`
	Filename    types.String `tfsdk:"filename" json:"filename,computed"`
	Kind        types.String `tfsdk:"kind" json:"kind,computed"`
	PublicID    types.String `tfsdk:"public_id" json:"publicId,computed"`
	SizeBytes   types.Int64  `tfsdk:"size_bytes" json:"sizeBytes,computed"`
	Status      types.String `tfsdk:"status" json:"status,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
}
