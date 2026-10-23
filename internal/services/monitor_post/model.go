// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor_post

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MonitorPostModel struct {
	ID             types.String      `tfsdk:"id" json:"id,computed"`
	Post           types.String      `tfsdk:"post" json:"post,required,no_refresh"`
	Name           types.String      `tfsdk:"name" json:"name,optional"`
	EventTypes     *[]types.String   `tfsdk:"event_types" json:"eventTypes,computed_optional"`
	IsActive       types.Bool        `tfsdk:"is_active" json:"isActive,computed_optional"`
	PostID         types.String      `tfsdk:"post_id" json:"postId,computed"`
	PostURL        types.String      `tfsdk:"post_url" json:"postUrl,computed"`
	AuthorID       types.String      `tfsdk:"author_id" json:"authorId,computed"`
	AuthorUsername types.String      `tfsdk:"author_username" json:"authorUsername,computed"`
	AuthorName     types.String      `tfsdk:"author_name" json:"authorName,computed"`
	Text           types.String      `tfsdk:"text" json:"text,computed"`
	PostedAt       timetypes.RFC3339 `tfsdk:"posted_at" json:"postedAt,computed" format:"date-time"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	NextBillingAt  timetypes.RFC3339 `tfsdk:"next_billing_at" json:"nextBillingAt,computed" format:"date-time"`
	PausedReason   types.String      `tfsdk:"paused_reason" json:"pausedReason,computed"`
	PausedAt       timetypes.RFC3339 `tfsdk:"paused_at" json:"pausedAt,computed" format:"date-time"`
}

func (m MonitorPostModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MonitorPostModel) MarshalJSONForUpdate(state MonitorPostModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
