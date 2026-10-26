// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package supportticket

import (
	"bytes"
	"errors"
	"mime/multipart"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apiform"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SupportTicketModel struct {
	ID          types.String                                                `tfsdk:"id" path:"id,optional"`
	Body        types.String                                                `tfsdk:"body" json:"body" api:"required,no_refresh"`
	Subject     types.String                                                `tfsdk:"subject" json:"subject" api:"required"`
	Status      types.String                                                `tfsdk:"status" json:"status" api:"optional"`
	CreatedAt   timetypes.RFC3339                                           `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	PublicID    types.String                                                `tfsdk:"public_id" json:"publicId" api:"computed"`
	UpdatedAt   timetypes.RFC3339                                           `tfsdk:"updated_at" json:"updatedAt" api:"computed" format:"date-time"`
	Attachments customfield.NestedObjectList[SupportTicketAttachmentsModel] `tfsdk:"attachments" json:"attachments" api:"computed,no_refresh"`
	Messages    customfield.NestedObjectList[SupportTicketMessagesModel]    `tfsdk:"messages" json:"messages" api:"computed"`
}

func (r SupportTicketModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		if e := writer.Close(); e != nil {
			err = errors.Join(err, e)
		}
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type SupportTicketAttachmentsModel struct {
	PublicID types.String `tfsdk:"public_id" json:"publicId" api:"computed"`
	Status   types.String `tfsdk:"status" json:"status" api:"computed"`
}

type SupportTicketMessagesModel struct {
	Attachments customfield.NestedObjectList[SupportTicketMessagesAttachmentsModel] `tfsdk:"attachments" json:"attachments" api:"computed"`
	Body        types.String                                                        `tfsdk:"body" json:"body" api:"computed"`
	CreatedAt   timetypes.RFC3339                                                   `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	Sender      types.String                                                        `tfsdk:"sender" json:"sender" api:"computed"`
}

type SupportTicketMessagesAttachmentsModel struct {
	ContentType types.String `tfsdk:"content_type" json:"contentType" api:"computed"`
	Filename    types.String `tfsdk:"filename" json:"filename" api:"computed"`
	Kind        types.String `tfsdk:"kind" json:"kind" api:"computed"`
	PublicID    types.String `tfsdk:"public_id" json:"publicId" api:"computed"`
	SizeBytes   types.Int64  `tfsdk:"size_bytes" json:"sizeBytes" api:"computed"`
	Status      types.String `tfsdk:"status" json:"status" api:"computed"`
	URL         types.String `tfsdk:"url" json:"url" api:"computed"`
}
