// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package guestwallet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GuestWalletModel struct {
	AmountMinor             types.Int64                                             `tfsdk:"amount_minor" json:"amount_minor" api:"required"`
	Currency                types.String                                            `tfsdk:"currency" json:"currency" api:"required"`
	IdempotencyKey          types.String                                            `tfsdk:"idempotency_key" json:"-" api:"required"`
	AccountRequired         types.Bool                                              `tfsdk:"account_required" json:"account_required" api:"computed"`
	APIKey                  types.String                                            `tfsdk:"api_key" json:"api_key" api:"computed"`
	CheckoutURL             types.String                                            `tfsdk:"checkout_url" json:"checkout_url" api:"computed"`
	CredentialNotice        types.String                                            `tfsdk:"credential_notice" json:"credential_notice" api:"computed"`
	Credits                 types.String                                            `tfsdk:"credits" json:"credits" api:"computed"`
	ExpiresAt               timetypes.RFC3339                                       `tfsdk:"expires_at" json:"expires_at" api:"computed" format:"date-time"`
	Instructions            types.String                                            `tfsdk:"instructions" json:"instructions" api:"computed"`
	PollAfterSeconds        types.Int64                                             `tfsdk:"poll_after_seconds" json:"poll_after_seconds" api:"computed"`
	PurchaseID              types.String                                            `tfsdk:"purchase_id" json:"purchase_id" api:"computed"`
	RequiresUserInteraction types.Bool                                              `tfsdk:"requires_user_interaction" json:"requires_user_interaction" api:"computed"`
	Status                  types.String                                            `tfsdk:"status" json:"status" api:"computed"`
	StatusURL               types.String                                            `tfsdk:"status_url" json:"status_url" api:"computed"`
	WalletID                types.String                                            `tfsdk:"wallet_id" json:"wallet_id" api:"computed"`
	Amount                  customfield.NestedObject[GuestWalletAmountModel]        `tfsdk:"amount" json:"amount" api:"computed"`
	Authorization           customfield.NestedObject[GuestWalletAuthorizationModel] `tfsdk:"authorization" json:"authorization" api:"computed"`
}

func (m GuestWalletModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m GuestWalletModel) MarshalJSONForUpdate(state GuestWalletModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type GuestWalletAmountModel struct {
	AmountMinor types.Int64  `tfsdk:"amount_minor" json:"amount_minor" api:"computed"`
	Currency    types.String `tfsdk:"currency" json:"currency" api:"computed"`
}

type GuestWalletAuthorizationModel struct {
	Header types.String `tfsdk:"header" json:"header" api:"computed"`
	Scheme types.String `tfsdk:"scheme" json:"scheme" api:"computed"`
}
