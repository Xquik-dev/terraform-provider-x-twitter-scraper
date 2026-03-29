// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_platform_link

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
)

type BotPlatformLinkModel struct {
	ID             types.String      `tfsdk:"id" json:"id,computed"`
	Platform       types.String      `tfsdk:"platform" json:"platform,required"`
	PlatformUserID types.String      `tfsdk:"platform_user_id" json:"platformUserId,required"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
}

func (m BotPlatformLinkModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m BotPlatformLinkModel) MarshalJSONForUpdate(state BotPlatformLinkModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
