// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package activity

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ActivityDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Every activity keeps a full record you can fetch & export.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Activity ID. It is the id of an Activity log row and the last part of the dashboard's activity page URL.",
				Required:    true,
			},
			"kind": schema.StringAttribute{
				Description: "Kind of the activity. Only a signed-in dashboard session sees an `api_key` record.",
				Computed:    true,
			},
			"object": schema.StringAttribute{
				Computed: true,
			},
			"request_id": schema.StringAttribute{
				Description: "ID of the API request behind this activity.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Description: "Type of the activity within its kind, the same as the type of its Activity log row. Examples are a write action type, an extraction tool type, a monitor change, a monitor event type or a ledger entry type.",
				Computed:    true,
			},
			"url": schema.StringAttribute{
				Description: "Dashboard page of this activity.",
				Computed:    true,
			},
			"details": schema.MapAttribute{
				Description: "Fields of this kind of activity, such as its target. A field can hold JSON, such as the data of a monitor event or the payload of a webhook delivery.",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"caller": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ActivityCallerDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"auth_method": schema.StringAttribute{
						Computed: true,
					},
					"api_key": schema.SingleNestedAttribute{
						Description: "Name and prefix of the API key used. Never its value.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ActivityCallerAPIKeyDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Computed: true,
							},
							"prefix": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"oauth_client": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ActivityCallerOAuthClientDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
			"charge": schema.SingleNestedAttribute{
				Description: "Credits charged for this activity.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ActivityChargeDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"credits": schema.Int64Attribute{
						Computed: true,
					},
					"balance_after": schema.Int64Attribute{
						Computed: true,
					},
					"at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
				},
			},
			"request": schema.SingleNestedAttribute{
				Description: "The API request behind this activity. Null for activity from before Oct 30, 2026, and for activity no API request starts, such as a monitor event, a webhook delivery or a credit grant.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ActivityRequestDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"method": schema.StringAttribute{
						Computed: true,
					},
					"path": schema.StringAttribute{
						Computed: true,
					},
					"http_status": schema.Int64Attribute{
						Computed: true,
					},
					"duration_ms": schema.Int64Attribute{
						Computed: true,
					},
					"ip_address": schema.StringAttribute{
						Computed: true,
					},
					"user_agent": schema.StringAttribute{
						Computed: true,
					},
					"at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
				},
			},
			"timestamps": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ActivityTimestampsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"updated_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"completed_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"confirmed_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"expires_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"occurred_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"discovered_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"posted_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"first_attempted_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"delivered_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"next_retry_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"exhausted_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"drawn_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"last_used_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"paused_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
				},
			},
			"related": schema.ListNestedAttribute{
				Description: "Up to 5 other activities on the same subject within 7 days, newest first.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ActivityRelatedDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"url": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *ActivityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ActivityDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
