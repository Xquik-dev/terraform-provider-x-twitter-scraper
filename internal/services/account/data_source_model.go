// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountDataSourceModel struct {
	MonitorsAllowed types.Int64                                                    `tfsdk:"monitors_allowed" json:"monitorsAllowed,computed"`
	MonitorsUsed    types.Int64                                                    `tfsdk:"monitors_used" json:"monitorsUsed,computed"`
	Plan            types.String                                                   `tfsdk:"plan" json:"plan,computed"`
	XUsername       types.String                                                   `tfsdk:"x_username" json:"xUsername,computed"`
	CreditInfo      customfield.NestedObject[AccountCreditInfoDataSourceModel]     `tfsdk:"credit_info" json:"creditInfo,computed"`
	MonitorBilling  customfield.NestedObject[AccountMonitorBillingDataSourceModel] `tfsdk:"monitor_billing" json:"monitorBilling,computed"`
}

type AccountCreditInfoDataSourceModel struct {
	AutoTopupAmountDollars types.Float64 `tfsdk:"auto_topup_amount_dollars" json:"autoTopupAmountDollars,computed"`
	AutoTopupEnabled       types.Bool    `tfsdk:"auto_topup_enabled" json:"autoTopupEnabled,computed"`
	AutoTopupThreshold     types.String  `tfsdk:"auto_topup_threshold" json:"autoTopupThreshold,computed"`
	Balance                types.String  `tfsdk:"balance" json:"balance,computed"`
	LifetimePurchased      types.String  `tfsdk:"lifetime_purchased" json:"lifetimePurchased,computed"`
	LifetimeUsed           types.String  `tfsdk:"lifetime_used" json:"lifetimeUsed,computed"`
}

type AccountMonitorBillingDataSourceModel struct {
	ActiveDailyEstimate         types.String `tfsdk:"active_daily_estimate" json:"activeDailyEstimate,computed"`
	ActiveHourlyBurn            types.String `tfsdk:"active_hourly_burn" json:"activeHourlyBurn,computed"`
	CreditsPerActiveMonitorDay  types.String `tfsdk:"credits_per_active_monitor_day" json:"creditsPerActiveMonitorDay,computed"`
	CreditsPerActiveMonitorHour types.String `tfsdk:"credits_per_active_monitor_hour" json:"creditsPerActiveMonitorHour,computed"`
	EventsIncluded              types.Bool   `tfsdk:"events_included" json:"eventsIncluded,computed"`
	InstantCheckIntervalSeconds types.Int64  `tfsdk:"instant_check_interval_seconds" json:"instantCheckIntervalSeconds,computed"`
	UnlimitedSlots              types.Bool   `tfsdk:"unlimited_slots" json:"unlimitedSlots,computed"`
}
