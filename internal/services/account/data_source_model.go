// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package account

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountDataSourceModel struct {
	MonitorsAllowed types.Int64                                                    `tfsdk:"monitors_allowed" json:"monitorsAllowed" api:"computed"`
	MonitorsUsed    types.Int64                                                    `tfsdk:"monitors_used" json:"monitorsUsed" api:"computed"`
	Plan            types.String                                                   `tfsdk:"plan" json:"plan" api:"computed"`
	XUsername       types.String                                                   `tfsdk:"x_username" json:"xUsername" api:"computed"`
	CreditInfo      customfield.NestedObject[AccountCreditInfoDataSourceModel]     `tfsdk:"credit_info" json:"creditInfo" api:"computed"`
	MonitorBilling  customfield.NestedObject[AccountMonitorBillingDataSourceModel] `tfsdk:"monitor_billing" json:"monitorBilling" api:"computed"`
}

type AccountCreditInfoDataSourceModel struct {
	AutoTopupAmountDollars types.Float64 `tfsdk:"auto_topup_amount_dollars" json:"autoTopupAmountDollars" api:"computed"`
	AutoTopupEnabled       types.Bool    `tfsdk:"auto_topup_enabled" json:"autoTopupEnabled" api:"computed"`
	AutoTopupThreshold     types.String  `tfsdk:"auto_topup_threshold" json:"autoTopupThreshold" api:"computed"`
	Balance                types.String  `tfsdk:"balance" json:"balance" api:"computed"`
	LifetimePurchased      types.String  `tfsdk:"lifetime_purchased" json:"lifetimePurchased" api:"computed"`
	LifetimeUsed           types.String  `tfsdk:"lifetime_used" json:"lifetimeUsed" api:"computed"`
}

type AccountMonitorBillingDataSourceModel struct {
	ActiveDailyEstimate         types.String `tfsdk:"active_daily_estimate" json:"activeDailyEstimate" api:"computed"`
	ActiveHourlyBurn            types.String `tfsdk:"active_hourly_burn" json:"activeHourlyBurn" api:"computed"`
	CreditsPerActiveMonitorDay  types.String `tfsdk:"credits_per_active_monitor_day" json:"creditsPerActiveMonitorDay" api:"computed"`
	CreditsPerActiveMonitorHour types.String `tfsdk:"credits_per_active_monitor_hour" json:"creditsPerActiveMonitorHour" api:"computed"`
	EventsIncluded              types.Bool   `tfsdk:"events_included" json:"eventsIncluded" api:"computed"`
	InstantCheckIntervalSeconds types.Int64  `tfsdk:"instant_check_interval_seconds" json:"instantCheckIntervalSeconds" api:"computed"`
	UnlimitedSlots              types.Bool   `tfsdk:"unlimited_slots" json:"unlimitedSlots" api:"computed"`
}
