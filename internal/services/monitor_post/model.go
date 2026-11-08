// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitorpost

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MonitorPostModel struct {
	ID             types.String                   `tfsdk:"id" json:"id" api:"computed"`
	Post           types.String                   `tfsdk:"post" json:"post" api:"required,no_refresh"`
	Name           types.String                   `tfsdk:"name" json:"name" api:"optional"`
	EventTypes     customfield.List[types.String] `tfsdk:"event_types" json:"eventTypes" api:"computed_optional"`
	IsActive       types.Bool                     `tfsdk:"is_active" json:"isActive" api:"computed_optional"`
	PostID         types.String                   `tfsdk:"post_id" json:"postId" api:"computed"`
	PostURL        types.String                   `tfsdk:"post_url" json:"postUrl" api:"computed"`
	AuthorID       types.String                   `tfsdk:"author_id" json:"authorId" api:"computed"`
	AuthorUsername types.String                   `tfsdk:"author_username" json:"authorUsername" api:"computed"`
	AuthorName     types.String                   `tfsdk:"author_name" json:"authorName" api:"computed"`
	Text           types.String                   `tfsdk:"text" json:"text" api:"computed"`
	PostedAt       timetypes.RFC3339              `tfsdk:"posted_at" json:"postedAt" api:"computed" format:"date-time"`
	CreatedAt      timetypes.RFC3339              `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	NextBillingAt  timetypes.RFC3339              `tfsdk:"next_billing_at" json:"nextBillingAt" api:"computed" format:"date-time"`
	PausedReason   types.String                   `tfsdk:"paused_reason" json:"pausedReason" api:"computed"`
	PausedAt       timetypes.RFC3339              `tfsdk:"paused_at" json:"pausedAt" api:"computed" format:"date-time"`
}

func (m MonitorPostModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MonitorPostModel) MarshalJSONForUpdate(state MonitorPostModel) (data []byte, err error) {
	m.Post = state.Post
	return apijson.MarshalForPatch(m, state)
}
