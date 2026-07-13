// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_write_action

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*XWriteActionDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "X write actions (tweets, likes, follows, DMs)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"action": schema.StringAttribute{
				Computed: true,
			},
			"charged": schema.BoolAttribute{
				Computed: true,
			},
			"charged_credits": schema.StringAttribute{
				Computed: true,
			},
			"confirmation_attempts": schema.Int64Attribute{
				Computed: true,
			},
			"confirmation_checked_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"confirmation_source": schema.StringAttribute{
				Computed: true,
			},
			"confirmed_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"message": schema.StringAttribute{
				Computed: true,
			},
			"message_id": schema.StringAttribute{
				Computed: true,
			},
			"retryable": schema.BoolAttribute{
				Computed: true,
			},
			"send_dispatched": schema.BoolAttribute{
				Computed: true,
			},
			"send_dispatched_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: `Available values: "success", "failed", "pending_confirmation".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"success",
						"failed",
						"pending_confirmation",
					),
				},
			},
			"target_id": schema.StringAttribute{
				Computed: true,
			},
			"tweet_id": schema.StringAttribute{
				Computed: true,
			},
			"write_action_id": schema.StringAttribute{
				Computed: true,
			},
			"media": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[XWriteActionMediaDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"x_write_action_count": schema.Int64Attribute{
						Computed: true,
					},
					"credits": schema.StringAttribute{
						Computed: true,
					},
					"kind": schema.StringAttribute{
						Description: `Available values: "none", "image", "video".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"none",
								"image",
								"video",
							),
						},
					},
					"total_bytes": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *XWriteActionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *XWriteActionDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
