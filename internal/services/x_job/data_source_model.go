// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package x_job

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XJobDataSourceModel struct {
	ID  types.String                                     `tfsdk:"id" path:"id,required"`
	Job customfield.NestedObject[XJobJobDataSourceModel] `tfsdk:"job" json:"job,computed"`
}

type XJobJobDataSourceModel struct {
	ID                  types.String                                            `tfsdk:"id" json:"id,computed"`
	URL                 types.String                                            `tfsdk:"url" json:"url,computed"`
	Account             customfield.NestedObject[XJobJobAccountDataSourceModel] `tfsdk:"account" json:"account,computed"`
	ApplyURL            types.String                                            `tfsdk:"apply_url" json:"applyUrl,computed"`
	Company             customfield.NestedObject[XJobJobCompanyDataSourceModel] `tfsdk:"company" json:"company,computed"`
	Description         types.String                                            `tfsdk:"description" json:"description,computed"`
	DescriptionMarkdown types.String                                            `tfsdk:"description_markdown" json:"descriptionMarkdown,computed"`
	EmploymentType      types.String                                            `tfsdk:"employment_type" json:"employmentType,computed"`
	Featured            types.Bool                                              `tfsdk:"featured" json:"featured,computed"`
	FormattedSalary     types.String                                            `tfsdk:"formatted_salary" json:"formattedSalary,computed"`
	JobFunction         types.String                                            `tfsdk:"job_function" json:"jobFunction,computed"`
	Location            types.String                                            `tfsdk:"location" json:"location,computed"`
	LocationType        types.String                                            `tfsdk:"location_type" json:"locationType,computed"`
	PostedAt            timetypes.RFC3339                                       `tfsdk:"posted_at" json:"postedAt,computed" format:"date-time"`
	SalaryCurrencyCode  types.String                                            `tfsdk:"salary_currency_code" json:"salaryCurrencyCode,computed"`
	SalaryInterval      types.String                                            `tfsdk:"salary_interval" json:"salaryInterval,computed"`
	SalaryMax           types.Float64                                           `tfsdk:"salary_max" json:"salaryMax,computed"`
	SalaryMin           types.Float64                                           `tfsdk:"salary_min" json:"salaryMin,computed"`
	SeniorityLevel      types.String                                            `tfsdk:"seniority_level" json:"seniorityLevel,computed"`
	ShortDescription    types.String                                            `tfsdk:"short_description" json:"shortDescription,computed"`
	Team                types.String                                            `tfsdk:"team" json:"team,computed"`
	Title               types.String                                            `tfsdk:"title" json:"title,computed"`
}

type XJobJobAccountDataSourceModel struct {
	ID                     types.String `tfsdk:"id" json:"id,computed"`
	Verified               types.Bool   `tfsdk:"verified" json:"verified,computed"`
	Description            types.String `tfsdk:"description" json:"description,computed"`
	Name                   types.String `tfsdk:"name" json:"name,computed"`
	ProfileImageShape      types.String `tfsdk:"profile_image_shape" json:"profileImageShape,computed"`
	ProfilePicture         types.String `tfsdk:"profile_picture" json:"profilePicture,computed"`
	ProfilePictureOriginal types.String `tfsdk:"profile_picture_original" json:"profilePictureOriginal,computed"`
	Username               types.String `tfsdk:"username" json:"username,computed"`
	VerifiedType           types.String `tfsdk:"verified_type" json:"verifiedType,computed"`
}

type XJobJobCompanyDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	LogoURL types.String `tfsdk:"logo_url" json:"logoUrl,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
}
