// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package compose

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComposeModel struct {
	Step              types.String `tfsdk:"step" json:"step" api:"required"`
	AdditionalContext types.String `tfsdk:"additional_context" json:"additionalContext" api:"optional"`
	CallToAction      types.String `tfsdk:"call_to_action" json:"callToAction" api:"optional"`
	Draft             types.String `tfsdk:"draft" json:"draft" api:"optional"`
	Goal              types.String `tfsdk:"goal" json:"goal" api:"optional"`
	HasLink           types.Bool   `tfsdk:"has_link" json:"hasLink" api:"optional"`
	HasMedia          types.Bool   `tfsdk:"has_media" json:"hasMedia" api:"optional"`
	MediaType         types.String `tfsdk:"media_type" json:"mediaType" api:"optional"`
	StyleUsername     types.String `tfsdk:"style_username" json:"styleUsername" api:"optional"`
	Tone              types.String `tfsdk:"tone" json:"tone" api:"optional"`
	Topic             types.String `tfsdk:"topic" json:"topic" api:"optional"`
}

func (m ComposeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComposeModel) MarshalJSONForUpdate(state ComposeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
