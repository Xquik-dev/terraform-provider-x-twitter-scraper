// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.ResourceWithConfigValidators = (*XProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "X write actions (tweets, likes, follows, DMs)",
		Attributes: map[string]schema.Attribute{
			"account": schema.StringAttribute{
				Description: "X account (@username or ID) to update profile",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Bio description",
				Optional:    true,
			},
			"location": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Description: "Display name",
				Optional:    true,
			},
			"url": schema.StringAttribute{
				Description: "Website URL",
				Optional:    true,
			},
			"success": schema.BoolAttribute{
				Computed: true,
			},
		},
	}
}

func (r *XProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *XProfileResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
