// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compose

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComposeModel struct {
	Step              types.String `tfsdk:"step" json:"step,required"`
	AdditionalContext types.String `tfsdk:"additional_context" json:"additionalContext,optional"`
	CallToAction      types.String `tfsdk:"call_to_action" json:"callToAction,optional"`
	Draft             types.String `tfsdk:"draft" json:"draft,optional"`
	Goal              types.String `tfsdk:"goal" json:"goal,optional"`
	HasLink           types.Bool   `tfsdk:"has_link" json:"hasLink,optional"`
	HasMedia          types.Bool   `tfsdk:"has_media" json:"hasMedia,optional"`
	MediaType         types.String `tfsdk:"media_type" json:"mediaType,optional"`
	StyleUsername     types.String `tfsdk:"style_username" json:"styleUsername,optional"`
	Tone              types.String `tfsdk:"tone" json:"tone,optional"`
	Topic             types.String `tfsdk:"topic" json:"topic,optional"`
}

func (m ComposeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComposeModel) MarshalJSONForUpdate(state ComposeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
