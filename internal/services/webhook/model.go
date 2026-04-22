// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebhookModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	URL        types.String      `tfsdk:"url" json:"url,required"`
	EventTypes *[]types.String   `tfsdk:"event_types" json:"eventTypes,required"`
	IsActive   types.Bool        `tfsdk:"is_active" json:"isActive,optional"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Secret     types.String      `tfsdk:"secret" json:"secret,computed"`
}

func (m WebhookModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WebhookModel) MarshalJSONForUpdate(state WebhookModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
