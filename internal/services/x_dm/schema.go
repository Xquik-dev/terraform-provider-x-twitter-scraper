// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_dm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*XDmResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account": schema.StringAttribute{
				Description: "X account (@username or account ID)",
				Required:    true,
			},
			"text": schema.StringAttribute{
				Required: true,
			},
			"reply_to_message_id": schema.StringAttribute{
				Optional: true,
			},
			"media_ids": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"message_id": schema.StringAttribute{
				Computed: true,
			},
			"success": schema.BoolAttribute{
				Computed: true,
			},
		},
	}
}

func (r *XDmResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *XDmResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
