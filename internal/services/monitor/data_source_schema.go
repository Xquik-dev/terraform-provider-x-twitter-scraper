// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monitor

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MonitorDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "X account monitoring with 1-second checks",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"is_active": schema.BoolAttribute{
				Computed: true,
			},
			"next_billing_at": schema.StringAttribute{
				Description: "Next hourly credit charge time for this account monitor.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"username": schema.StringAttribute{
				Computed: true,
			},
			"x_user_id": schema.StringAttribute{
				Computed: true,
			},
			"event_types": schema.ListAttribute{
				Description: "Array of event types to subscribe to.",
				Computed:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"tweet.new",
							"tweet.reply",
							"tweet.retweet",
							"tweet.quote",
							"tweet.media",
							"tweet.link",
							"tweet.poll",
							"tweet.mention",
							"tweet.hashtag",
							"tweet.longform",
							"profile.avatar.changed",
							"profile.banner.changed",
							"profile.name.changed",
							"profile.username.changed",
							"profile.bio.changed",
							"profile.location.changed",
							"profile.url.changed",
							"profile.verified.changed",
							"profile.protected.changed",
							"profile.pinned_tweet.changed",
							"profile.unavailable.changed",
						),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *MonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MonitorDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
