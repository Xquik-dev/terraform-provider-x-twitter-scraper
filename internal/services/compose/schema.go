// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compose

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ComposeResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Tweet composition, drafts, writing styles & radar",
		Attributes: map[string]schema.Attribute{
			"step": schema.StringAttribute{
				Description: "Workflow step\nAvailable values: \"compose\", \"refine\", \"score\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"compose",
						"refine",
						"score",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"additional_context": schema.StringAttribute{
				Description:   "Extra context or URLs (refine)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"call_to_action": schema.StringAttribute{
				Description:   "Desired call to action (refine)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"draft": schema.StringAttribute{
				Description:   "Tweet draft text to evaluate (score)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"goal": schema.StringAttribute{
				Description: "Optimization goal\nAvailable values: \"engagement\", \"followers\", \"authority\", \"conversation\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"engagement",
						"followers",
						"authority",
						"conversation",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"has_link": schema.BoolAttribute{
				Description:   "Whether a link is attached (score)",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"has_media": schema.BoolAttribute{
				Description:   "Whether media is attached (score)",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"media_type": schema.StringAttribute{
				Description: "Media type (refine)\nAvailable values: \"photo\", \"video\", \"none\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"photo",
						"video",
						"none",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"style_username": schema.StringAttribute{
				Description:   "Cached style username for voice matching (compose)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"tone": schema.StringAttribute{
				Description:   "Desired tone (refine)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"topic": schema.StringAttribute{
				Description:   "Tweet topic (compose, refine)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"feedback": schema.StringAttribute{
				Description: "AI feedback on the draft",
				Computed:    true,
			},
			"score": schema.Float64Attribute{
				Description: "Engagement score (0-100)",
				Computed:    true,
			},
			"text": schema.StringAttribute{
				Description: "Generated or refined tweet text",
				Computed:    true,
			},
			"suggestions": schema.ListAttribute{
				Description: "Improvement suggestions",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (r *ComposeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ComposeResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
