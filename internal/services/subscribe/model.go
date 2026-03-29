// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscribe

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type SubscribeModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
	Status  types.String `tfsdk:"status" json:"status,computed"`
	URL     types.String `tfsdk:"url" json:"url,computed"`
}

func (m SubscribeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SubscribeModel) MarshalJSONForUpdate(state SubscribeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
