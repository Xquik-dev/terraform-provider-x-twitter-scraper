// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebhookModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	URL                 types.String      `tfsdk:"url" json:"url,required"`
	EventTypes          *[]types.String   `tfsdk:"event_types" json:"eventTypes,required"`
	IsActive            types.Bool        `tfsdk:"is_active" json:"isActive,optional"`
	ConsecutiveFailures types.Int64       `tfsdk:"consecutive_failures" json:"consecutiveFailures,computed"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeliveryStatus      types.String      `tfsdk:"delivery_status" json:"deliveryStatus,computed"`
	FailureHardCap      types.Int64       `tfsdk:"failure_hard_cap" json:"failureHardCap,computed"`
	Secret              types.String      `tfsdk:"secret" json:"secret,computed"`
}

func (m WebhookModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WebhookModel) MarshalJSONForUpdate(state WebhookModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
