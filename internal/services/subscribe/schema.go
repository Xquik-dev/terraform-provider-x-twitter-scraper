// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscribe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*SubscribeResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Subscription & billing",
		Attributes: map[string]schema.Attribute{
			"message": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "checkout_created", "already_subscribed", "payment_issue".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"checkout_created",
						"already_subscribed",
						"payment_issue",
					),
				},
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SubscribeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SubscribeResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
