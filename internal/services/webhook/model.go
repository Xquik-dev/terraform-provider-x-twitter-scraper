// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebhookModel struct {
	ID                  types.String      `tfsdk:"id" json:"id" api:"computed"`
	URL                 types.String      `tfsdk:"url" json:"url" api:"required"`
	EventTypes          *[]types.String   `tfsdk:"event_types" json:"eventTypes" api:"required"`
	IsActive            types.Bool        `tfsdk:"is_active" json:"isActive" api:"optional"`
	ConsecutiveFailures types.Int64       `tfsdk:"consecutive_failures" json:"consecutiveFailures" api:"computed"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	DeliveryStatus      types.String      `tfsdk:"delivery_status" json:"deliveryStatus" api:"computed"`
	FailureHardCap      types.Int64       `tfsdk:"failure_hard_cap" json:"failureHardCap" api:"computed"`
	Secret              types.String      `tfsdk:"secret" json:"secret" api:"computed"`
}

func (m WebhookModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WebhookModel) MarshalJSONForUpdate(state WebhookModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
