// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*IntegrationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Push notification integrations (Telegram)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				Description: `Available values: "telegram".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("telegram"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"config": schema.SingleNestedAttribute{
				Description: "Integration config (e.g. Telegram chatId)",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"chat_id": schema.StringAttribute{
						Required: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"event_types": schema.ListAttribute{
				Required: true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"tweet.new",
							"tweet.reply",
							"tweet.retweet",
							"tweet.quote",
							"follower.gained",
							"follower.lost",
						),
					),
				},
				ElementType: types.StringType,
			},
			"is_active": schema.BoolAttribute{
				Optional: true,
			},
			"scope_all_monitors": schema.BoolAttribute{
				Optional: true,
			},
			"silent_push": schema.BoolAttribute{
				Optional: true,
			},
			"filters": schema.MapAttribute{
				Description: "Event filter rules (JSON)",
				Optional:    true,
				ElementType: jsontypes.NormalizedType{},
			},
			"message_template": schema.MapAttribute{
				Description: "Custom message template (JSON)",
				Optional:    true,
				ElementType: jsontypes.NormalizedType{},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *IntegrationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *IntegrationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
