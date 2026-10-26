// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xjob

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XJobDataSourceModel struct {
	ID  types.String                                     `tfsdk:"id" path:"id,required"`
	Job customfield.NestedObject[XJobJobDataSourceModel] `tfsdk:"job" json:"job" api:"computed"`
}

type XJobJobDataSourceModel struct {
	ID                  types.String                                            `tfsdk:"id" json:"id" api:"computed"`
	URL                 types.String                                            `tfsdk:"url" json:"url" api:"computed"`
	Account             customfield.NestedObject[XJobJobAccountDataSourceModel] `tfsdk:"account" json:"account" api:"computed"`
	ApplyURL            types.String                                            `tfsdk:"apply_url" json:"applyUrl" api:"computed"`
	Company             customfield.NestedObject[XJobJobCompanyDataSourceModel] `tfsdk:"company" json:"company" api:"computed"`
	Description         types.String                                            `tfsdk:"description" json:"description" api:"computed"`
	DescriptionMarkdown types.String                                            `tfsdk:"description_markdown" json:"descriptionMarkdown" api:"computed"`
	EmploymentType      types.String                                            `tfsdk:"employment_type" json:"employmentType" api:"computed"`
	Featured            types.Bool                                              `tfsdk:"featured" json:"featured" api:"computed"`
	FormattedSalary     types.String                                            `tfsdk:"formatted_salary" json:"formattedSalary" api:"computed"`
	JobFunction         types.String                                            `tfsdk:"job_function" json:"jobFunction" api:"computed"`
	Location            types.String                                            `tfsdk:"location" json:"location" api:"computed"`
	LocationType        types.String                                            `tfsdk:"location_type" json:"locationType" api:"computed"`
	PostedAt            timetypes.RFC3339                                       `tfsdk:"posted_at" json:"postedAt" api:"computed" format:"date-time"`
	SalaryCurrencyCode  types.String                                            `tfsdk:"salary_currency_code" json:"salaryCurrencyCode" api:"computed"`
	SalaryInterval      types.String                                            `tfsdk:"salary_interval" json:"salaryInterval" api:"computed"`
	SalaryMax           types.Float64                                           `tfsdk:"salary_max" json:"salaryMax" api:"computed"`
	SalaryMin           types.Float64                                           `tfsdk:"salary_min" json:"salaryMin" api:"computed"`
	SeniorityLevel      types.String                                            `tfsdk:"seniority_level" json:"seniorityLevel" api:"computed"`
	ShortDescription    types.String                                            `tfsdk:"short_description" json:"shortDescription" api:"computed"`
	Team                types.String                                            `tfsdk:"team" json:"team" api:"computed"`
	Title               types.String                                            `tfsdk:"title" json:"title" api:"computed"`
}

type XJobJobAccountDataSourceModel struct {
	ID                     types.String `tfsdk:"id" json:"id" api:"computed"`
	Verified               types.Bool   `tfsdk:"verified" json:"verified" api:"computed"`
	Description            types.String `tfsdk:"description" json:"description" api:"computed"`
	Name                   types.String `tfsdk:"name" json:"name" api:"computed"`
	ProfileImageShape      types.String `tfsdk:"profile_image_shape" json:"profileImageShape" api:"computed"`
	ProfilePicture         types.String `tfsdk:"profile_picture" json:"profilePicture" api:"computed"`
	ProfilePictureOriginal types.String `tfsdk:"profile_picture_original" json:"profilePictureOriginal" api:"computed"`
	Username               types.String `tfsdk:"username" json:"username" api:"computed"`
	VerifiedType           types.String `tfsdk:"verified_type" json:"verifiedType" api:"computed"`
}

type XJobJobCompanyDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id" api:"computed"`
	LogoURL types.String `tfsdk:"logo_url" json:"logoUrl" api:"computed"`
	Name    types.String `tfsdk:"name" json:"name" api:"computed"`
}
