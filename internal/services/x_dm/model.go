// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_dm

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type XDmModel struct {
	UserID           types.String    `tfsdk:"user_id" path:"userId,required"`
	Account          types.String    `tfsdk:"account" json:"account,required"`
	Text             types.String    `tfsdk:"text" json:"text,required"`
	ReplyToMessageID types.String    `tfsdk:"reply_to_message_id" json:"reply_to_message_id,optional"`
	MediaIDs         *[]types.String `tfsdk:"media_ids" json:"media_ids,optional"`
	MessageID        types.String    `tfsdk:"message_id" json:"messageId,computed"`
	Success          types.Bool      `tfsdk:"success" json:"success,computed"`
}

func (m XDmModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XDmModel) MarshalJSONForUpdate(state XDmModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
