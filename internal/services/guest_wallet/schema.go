// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guest_wallet

import (
	"context"
	"regexp"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*GuestWalletResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Accountless prepaid access for paid read endpoints",
		Attributes: map[string]schema.Attribute{
			"amount_minor": schema.Int64Attribute{
				Description: "USD cents accepted for this checkout.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(1000, 25000),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"currency": schema.StringAttribute{
				Description: `Available values: "usd".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("usd"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"idempotency_key": schema.StringAttribute{
				Description: "Cryptographically random UUID v4. Reuse it only to retry the same wallet and amount request.",
				Required:    true,
				Sensitive:   true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}$`),
						"must be a UUID v4",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_required": schema.BoolAttribute{
				Computed: true,
			},
			"api_key": schema.StringAttribute{
				Description: "Paid-read bearer credential returned only by initial creation. Store it as a secret. Never place it in a URL or log.",
				Computed:    true,
				Sensitive:   true,
			},
			"checkout_url": schema.StringAttribute{
				Description: "Raw Stripe-hosted checkout URL for user interaction.",
				Computed:    true,
			},
			"credential_notice": schema.StringAttribute{
				Description: `Available values: "Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available.".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."),
				},
			},
			"credits": schema.StringAttribute{
				Description: "Credits granted after verified payment.",
				Computed:    true,
			},
			"expires_at": schema.StringAttribute{
				Description: "Time when the pending checkout expires.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"instructions": schema.StringAttribute{
				Description: `Available values: "Give checkout_url to the user. They must complete payment on Stripe. Never submit payment for them. After payment, poll status_url every poll_after_seconds until latest_purchase.status is no longer pending.".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Give checkout_url to the user. They must complete payment on Stripe. Never submit payment for them. After payment, poll status_url every poll_after_seconds until latest_purchase.status is no longer pending."),
				},
			},
			"poll_after_seconds": schema.Int64Attribute{
				Description: "Wait at least this long before polling status_url.\nAvailable values: 2.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.OneOf(2),
				},
			},
			"purchase_id": schema.StringAttribute{
				Computed: true,
			},
			"requires_user_interaction": schema.BoolAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "creating", "pending", "paid", "expired", "failed", "refunded", "disputed".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"creating",
						"pending",
						"paid",
						"expired",
						"failed",
						"refunded",
						"disputed",
					),
				},
			},
			"status_url": schema.StringAttribute{
				Description: `Available values: "https://xquik.com/api/v1/guest-wallets/status".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("https://xquik.com/api/v1/guest-wallets/status"),
				},
			},
			"wallet_id": schema.StringAttribute{
				Computed: true,
			},
			"amount": schema.SingleNestedAttribute{
				Description: "Confirmed USD amount for a guest wallet purchase.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[GuestWalletAmountModel](ctx),
				Attributes: map[string]schema.Attribute{
					"amount_minor": schema.Int64Attribute{
						Description: "USD amount in cents. Accepted range is $10-$250.",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(1000, 25000),
						},
					},
					"currency": schema.StringAttribute{
						Description: `Available values: "usd".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("usd"),
						},
					},
				},
			},
			"authorization": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[GuestWalletAuthorizationModel](ctx),
				Attributes: map[string]schema.Attribute{
					"header": schema.StringAttribute{
						Description: `Available values: "Authorization".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("Authorization"),
						},
					},
					"scheme": schema.StringAttribute{
						Description: `Available values: "Bearer".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("Bearer"),
						},
					},
				},
			},
		},
	}
}

func (r *GuestWalletResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *GuestWalletResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
