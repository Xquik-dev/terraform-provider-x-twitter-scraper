// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package x_job

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*XJobDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Public X job listings, search and details",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"job": schema.SingleNestedAttribute{
				Description: "Public X job listing. A search returns the fields down to the salary, plus company & account. A lookup adds work type, seniority, function, team & descriptions. A field X leaves empty is left out.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XJobJobDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Job ID.",
						Computed:    true,
					},
					"url": schema.StringAttribute{
						Description: "The job page on X.",
						Computed:    true,
					},
					"account": schema.SingleNestedAttribute{
						Description: "The X account that posted a job. A job has none when X imported it from a job board or no longer shows the account.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XJobJobAccountDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "X user ID of the account.",
								Computed:    true,
							},
							"verified": schema.BoolAttribute{
								Description: "Whether X shows a verification badge on the account.",
								Computed:    true,
							},
							"description": schema.StringAttribute{
								Description: "Bio of the account. A job lookup returns it.",
								Computed:    true,
							},
							"name": schema.StringAttribute{
								Description: "Display name of the account that posted the job.",
								Computed:    true,
							},
							"profile_image_shape": schema.StringAttribute{
								Description: "Avatar shape, such as Circle or Square.",
								Computed:    true,
							},
							"profile_picture": schema.StringAttribute{
								Description: "Avatar URL of the account that posted the job.",
								Computed:    true,
							},
							"profile_picture_original": schema.StringAttribute{
								Description: "The job account's avatar at full size.",
								Computed:    true,
							},
							"username": schema.StringAttribute{
								Description: "Handle without @.",
								Computed:    true,
							},
							"verified_type": schema.StringAttribute{
								Description: "Badge type, such as Business or Government.",
								Computed:    true,
							},
						},
					},
					"apply_url": schema.StringAttribute{
						Description: "Where applicants apply, off X.",
						Computed:    true,
					},
					"company": schema.SingleNestedAttribute{
						Description: "The company profile of a job X imported from a job board.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[XJobJobCompanyDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "X company profile ID.",
								Computed:    true,
							},
							"logo_url": schema.StringAttribute{
								Description: "Company logo URL.",
								Computed:    true,
							},
							"name": schema.StringAttribute{
								Description: "Company name.",
								Computed:    true,
							},
						},
					},
					"description": schema.StringAttribute{
						Description: "Full description as plain text, 1 line per paragraph.",
						Computed:    true,
					},
					"description_markdown": schema.StringAttribute{
						Description: "Full description as Markdown, with headings, lists, links, bold & italics.",
						Computed:    true,
					},
					"employment_type": schema.StringAttribute{
						Description: "Contract type (full_time, full_time_contract, part_time or contract_to_hire).",
						Computed:    true,
					},
					"featured": schema.BoolAttribute{
						Description: "Whether X features the job.",
						Computed:    true,
					},
					"formatted_salary": schema.StringAttribute{
						Description: "Salary as X shows it.",
						Computed:    true,
					},
					"job_function": schema.StringAttribute{
						Description: "Job function, such as software_engineering.",
						Computed:    true,
					},
					"location": schema.StringAttribute{
						Description: "Location as the employer wrote it.",
						Computed:    true,
					},
					"location_type": schema.StringAttribute{
						Description: "Where the work happens (onsite, remote or hybrid).",
						Computed:    true,
					},
					"posted_at": schema.StringAttribute{
						Description: "When the employer posted the job.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"salary_currency_code": schema.StringAttribute{
						Description: "ISO 4217 currency code of the salary.",
						Computed:    true,
					},
					"salary_interval": schema.StringAttribute{
						Description: "Pay period of the salary (annually or hourly).",
						Computed:    true,
					},
					"salary_max": schema.Float64Attribute{
						Description: "Highest salary, in salaryCurrencyCode per salaryInterval.",
						Computed:    true,
					},
					"salary_min": schema.Float64Attribute{
						Description: "Lowest salary, in salaryCurrencyCode per salaryInterval.",
						Computed:    true,
					},
					"seniority_level": schema.StringAttribute{
						Description: "Seniority (intern, entry_level, junior, mid_level, senior, lead, manager or executive).",
						Computed:    true,
					},
					"short_description": schema.StringAttribute{
						Description: "Short summary of the job, as plain text.",
						Computed:    true,
					},
					"team": schema.StringAttribute{
						Description: "Team inside the company.",
						Computed:    true,
					},
					"title": schema.StringAttribute{
						Description: "Job title.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *XJobDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *XJobDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
