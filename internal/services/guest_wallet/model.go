// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guest_wallet

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GuestWalletModel struct {
	AmountMinor             types.Int64                                             `tfsdk:"amount_minor" json:"amount_minor,required"`
	Currency                types.String                                            `tfsdk:"currency" json:"currency,required"`
	AccountRequired         types.Bool                                              `tfsdk:"account_required" json:"account_required,computed"`
	APIKey                  types.String                                            `tfsdk:"api_key" json:"api_key,computed"`
	CheckoutURL             types.String                                            `tfsdk:"checkout_url" json:"checkout_url,computed"`
	CredentialNotice        types.String                                            `tfsdk:"credential_notice" json:"credential_notice,computed"`
	Credits                 types.String                                            `tfsdk:"credits" json:"credits,computed"`
	ExpiresAt               timetypes.RFC3339                                       `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Instructions            types.String                                            `tfsdk:"instructions" json:"instructions,computed"`
	PollAfterSeconds        types.Int64                                             `tfsdk:"poll_after_seconds" json:"poll_after_seconds,computed"`
	PurchaseID              types.String                                            `tfsdk:"purchase_id" json:"purchase_id,computed"`
	RequiresUserInteraction types.Bool                                              `tfsdk:"requires_user_interaction" json:"requires_user_interaction,computed"`
	Status                  types.String                                            `tfsdk:"status" json:"status,computed"`
	StatusURL               types.String                                            `tfsdk:"status_url" json:"status_url,computed"`
	WalletID                types.String                                            `tfsdk:"wallet_id" json:"wallet_id,computed"`
	Amount                  customfield.NestedObject[GuestWalletAmountModel]        `tfsdk:"amount" json:"amount,computed"`
	Authorization           customfield.NestedObject[GuestWalletAuthorizationModel] `tfsdk:"authorization" json:"authorization,computed"`
}

func (m GuestWalletModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m GuestWalletModel) MarshalJSONForUpdate(state GuestWalletModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type GuestWalletAmountModel struct {
	AmountMinor types.Int64  `tfsdk:"amount_minor" json:"amount_minor,computed"`
	Currency    types.String `tfsdk:"currency" json:"currency,computed"`
}

type GuestWalletAuthorizationModel struct {
	Header types.String `tfsdk:"header" json:"header,computed"`
	Scheme types.String `tfsdk:"scheme" json:"scheme,computed"`
}
