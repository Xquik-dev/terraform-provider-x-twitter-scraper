// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package support_ticket

import (
	"context"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*SupportTicketResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Support ticket management",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"body": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"subject": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Description: `Available values: "open", "resolved", "closed".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"open",
						"resolved",
						"closed",
					),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"public_id": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"attachments": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[SupportTicketAttachmentsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"public_id": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Description: `Available values: "pending", "ready", "failed".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"pending",
									"ready",
									"failed",
								),
							},
						},
					},
				},
			},
			"messages": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[SupportTicketMessagesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attachments": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[SupportTicketMessagesAttachmentsModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"content_type": schema.StringAttribute{
										Description: "Validated media type.\nAvailable values: \"image/jpeg\", \"image/png\", \"image/gif\", \"image/webp\", \"video/mp4\", \"video/quicktime\", \"video/webm\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"image/jpeg",
												"image/png",
												"image/gif",
												"image/webp",
												"video/mp4",
												"video/quicktime",
												"video/webm",
											),
										},
									},
									"filename": schema.StringAttribute{
										Computed: true,
									},
									"kind": schema.StringAttribute{
										Description: "Attachment media class.\nAvailable values: \"image\", \"video\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("image", "video"),
										},
									},
									"public_id": schema.StringAttribute{
										Computed: true,
									},
									"size_bytes": schema.Int64Attribute{
										Computed: true,
									},
									"status": schema.StringAttribute{
										Description: "Storage processing state.\nAvailable values: \"pending\", \"ready\", \"failed\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"pending",
												"ready",
												"failed",
											),
										},
									},
									"url": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"body": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"sender": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *SupportTicketResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SupportTicketResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
