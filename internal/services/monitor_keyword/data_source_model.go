// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitorkeyword

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MonitorKeywordDataSourceModel struct {
	ID            types.String                   `tfsdk:"id" path:"id,required"`
	CreatedAt     timetypes.RFC3339              `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	IsActive      types.Bool                     `tfsdk:"is_active" json:"isActive" api:"computed"`
	NextBillingAt timetypes.RFC3339              `tfsdk:"next_billing_at" json:"nextBillingAt" api:"computed" format:"date-time"`
	Query         types.String                   `tfsdk:"query" json:"query" api:"computed"`
	EventTypes    customfield.List[types.String] `tfsdk:"event_types" json:"eventTypes" api:"computed"`
}
