// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package activity

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ActivityDataSourceModel struct {
	ID         types.String                                                 `tfsdk:"id" path:"id,required"`
	Kind       types.String                                                 `tfsdk:"kind" json:"kind" api:"computed"`
	Object     types.String                                                 `tfsdk:"object" json:"object" api:"computed"`
	RequestID  types.String                                                 `tfsdk:"request_id" json:"requestId" api:"computed"`
	Status     types.String                                                 `tfsdk:"status" json:"status" api:"computed"`
	Type       types.String                                                 `tfsdk:"type" json:"type" api:"computed"`
	URL        types.String                                                 `tfsdk:"url" json:"url" api:"computed"`
	Details    customfield.Map[jsontypes.Normalized]                        `tfsdk:"details" json:"details" api:"computed"`
	Caller     customfield.NestedObject[ActivityCallerDataSourceModel]      `tfsdk:"caller" json:"caller" api:"computed"`
	Charge     customfield.NestedObject[ActivityChargeDataSourceModel]      `tfsdk:"charge" json:"charge" api:"computed"`
	Request    customfield.NestedObject[ActivityRequestDataSourceModel]     `tfsdk:"request" json:"request" api:"computed"`
	Timestamps customfield.NestedObject[ActivityTimestampsDataSourceModel]  `tfsdk:"timestamps" json:"timestamps" api:"computed"`
	Related    customfield.NestedObjectList[ActivityRelatedDataSourceModel] `tfsdk:"related" json:"related" api:"computed"`
}

type ActivityCallerDataSourceModel struct {
	AuthMethod  types.String                                                       `tfsdk:"auth_method" json:"authMethod" api:"computed"`
	APIKey      customfield.NestedObject[ActivityCallerAPIKeyDataSourceModel]      `tfsdk:"api_key" json:"apiKey" api:"computed"`
	OAuthClient customfield.NestedObject[ActivityCallerOAuthClientDataSourceModel] `tfsdk:"oauth_client" json:"oauthClient" api:"computed"`
}

type ActivityCallerAPIKeyDataSourceModel struct {
	Name   types.String `tfsdk:"name" json:"name" api:"computed"`
	Prefix types.String `tfsdk:"prefix" json:"prefix" api:"computed"`
}

type ActivityCallerOAuthClientDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id" api:"computed"`
	Name types.String `tfsdk:"name" json:"name" api:"computed"`
}

type ActivityChargeDataSourceModel struct {
	Credits      types.Int64       `tfsdk:"credits" json:"credits" api:"computed"`
	BalanceAfter types.Int64       `tfsdk:"balance_after" json:"balanceAfter" api:"computed"`
	At           timetypes.RFC3339 `tfsdk:"at" json:"at" api:"computed" format:"date-time"`
}

type ActivityRequestDataSourceModel struct {
	Method     types.String      `tfsdk:"method" json:"method" api:"computed"`
	Path       types.String      `tfsdk:"path" json:"path" api:"computed"`
	HTTPStatus types.Int64       `tfsdk:"http_status" json:"httpStatus" api:"computed"`
	DurationMs types.Int64       `tfsdk:"duration_ms" json:"durationMs" api:"computed"`
	IPAddress  types.String      `tfsdk:"ip_address" json:"ipAddress" api:"computed"`
	UserAgent  types.String      `tfsdk:"user_agent" json:"userAgent" api:"computed"`
	At         timetypes.RFC3339 `tfsdk:"at" json:"at" api:"computed" format:"date-time"`
}

type ActivityTimestampsDataSourceModel struct {
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	UpdatedAt        timetypes.RFC3339 `tfsdk:"updated_at" json:"updatedAt" api:"computed" format:"date-time"`
	CompletedAt      timetypes.RFC3339 `tfsdk:"completed_at" json:"completedAt" api:"computed" format:"date-time"`
	ConfirmedAt      timetypes.RFC3339 `tfsdk:"confirmed_at" json:"confirmedAt" api:"computed" format:"date-time"`
	ExpiresAt        timetypes.RFC3339 `tfsdk:"expires_at" json:"expiresAt" api:"computed" format:"date-time"`
	OccurredAt       timetypes.RFC3339 `tfsdk:"occurred_at" json:"occurredAt" api:"computed" format:"date-time"`
	DiscoveredAt     timetypes.RFC3339 `tfsdk:"discovered_at" json:"discoveredAt" api:"computed" format:"date-time"`
	PostedAt         timetypes.RFC3339 `tfsdk:"posted_at" json:"postedAt" api:"computed" format:"date-time"`
	FirstAttemptedAt timetypes.RFC3339 `tfsdk:"first_attempted_at" json:"firstAttemptedAt" api:"computed" format:"date-time"`
	DeliveredAt      timetypes.RFC3339 `tfsdk:"delivered_at" json:"deliveredAt" api:"computed" format:"date-time"`
	NextRetryAt      timetypes.RFC3339 `tfsdk:"next_retry_at" json:"nextRetryAt" api:"computed" format:"date-time"`
	ExhaustedAt      timetypes.RFC3339 `tfsdk:"exhausted_at" json:"exhaustedAt" api:"computed" format:"date-time"`
	DrawnAt          timetypes.RFC3339 `tfsdk:"drawn_at" json:"drawnAt" api:"computed" format:"date-time"`
	LastUsedAt       timetypes.RFC3339 `tfsdk:"last_used_at" json:"lastUsedAt" api:"computed" format:"date-time"`
	PausedAt         timetypes.RFC3339 `tfsdk:"paused_at" json:"pausedAt" api:"computed" format:"date-time"`
}

type ActivityRelatedDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id" api:"computed"`
	Type      types.String      `tfsdk:"type" json:"type" api:"computed"`
	Status    types.String      `tfsdk:"status" json:"status" api:"computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	URL       types.String      `tfsdk:"url" json:"url" api:"computed"`
}
