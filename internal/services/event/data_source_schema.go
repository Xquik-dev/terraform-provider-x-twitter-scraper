// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EventDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Activity events from monitored accounts",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"monitor_id": schema.StringAttribute{
				Computed: true,
			},
			"occurred_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"type": schema.StringAttribute{
				Description: `Available values: "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote", "follower.gained", "follower.lost".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"tweet.new",
						"tweet.reply",
						"tweet.retweet",
						"tweet.quote",
						"follower.gained",
						"follower.lost",
					),
				},
			},
			"username": schema.StringAttribute{
				Computed: true,
			},
			"x_event_id": schema.StringAttribute{
				Computed: true,
			},
			"data": schema.MapAttribute{
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
		},
	}
}

func (d *EventDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EventDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
