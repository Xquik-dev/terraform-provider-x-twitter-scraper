// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package x_write

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func resourceSchema(_ context.Context, op operation) schema.Schema {
	requiresReplace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	return schema.Schema{
		Description: "Executes one " + op.action + " write and stores only its durable canonical write-action record.",
		Attributes: map[string]schema.Attribute{
			"id":                        schema.StringAttribute{Computed: true, Description: "Stable canonical write-action ID."},
			"action":                    schema.StringAttribute{Computed: true},
			"account":                   schema.StringAttribute{Required: true, Description: "X account used for the request. This is preserved separately from response_account_*.", PlanModifiers: requiresReplace, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}},
			"idempotency_key":           schema.StringAttribute{Required: true, Sensitive: true, Description: "Visible ASCII key unique to this intended write. Reuse only for an exact replay.", PlanModifiers: requiresReplace, Validators: []validator.String{stringvalidator.LengthBetween(1, 255), stringvalidator.RegexMatches(regexp.MustCompile(`^[!-~]+$`), "must contain visible ASCII only")}},
			"target_id":                 schema.StringAttribute{Optional: true, Description: "Target Tweet, user, or community ID for targeted writes.", PlanModifiers: requiresReplace},
			"payload_json":              schema.StringAttribute{Optional: true, Sensitive: true, Description: "Operation-specific JSON object. Do not include account.", PlanModifiers: requiresReplace},
			"object":                    schema.StringAttribute{Computed: true},
			"write_action_id":           schema.StringAttribute{Computed: true},
			"status":                    schema.StringAttribute{Computed: true},
			"terminal":                  schema.BoolAttribute{Computed: true},
			"retryable":                 schema.BoolAttribute{Computed: true},
			"safe_to_retry":             schema.BoolAttribute{Computed: true},
			"status_url":                schema.StringAttribute{Computed: true},
			"poll_after_ms":             schema.Int64Attribute{Computed: true},
			"charged":                   schema.BoolAttribute{Computed: true},
			"charged_credits":           schema.StringAttribute{Computed: true},
			"target_id_response":        schema.StringAttribute{Computed: true},
			"billing_status":            schema.StringAttribute{Computed: true},
			"billing_charged":           schema.BoolAttribute{Computed: true},
			"billing_charged_credits":   schema.StringAttribute{Computed: true},
			"billing_planned_credits":   schema.StringAttribute{Computed: true},
			"request_hash":              schema.StringAttribute{Computed: true},
			"request_payload_json":      schema.StringAttribute{Computed: true, Sensitive: true, Description: "Sanitized payload recorded by the canonical write action."},
			"request_id":                schema.StringAttribute{Computed: true},
			"response_account_id":       schema.StringAttribute{Computed: true},
			"response_account_username": schema.StringAttribute{Computed: true},
			"result_id":                 schema.StringAttribute{Computed: true},
			"result_state":              schema.StringAttribute{Computed: true},
			"result_type":               schema.StringAttribute{Computed: true},
			"response_target_id":        schema.StringAttribute{Computed: true},
			"response_target_type":      schema.StringAttribute{Computed: true},
			"next_action_type":          schema.StringAttribute{Computed: true},
			"next_action_after_ms":      schema.Int64Attribute{Computed: true},
			"next_action_requires_new_idempotency_key": schema.BoolAttribute{Computed: true},
			"next_action_url":                          schema.StringAttribute{Computed: true},
			"send_dispatched":                          schema.BoolAttribute{Computed: true},
			"success":                                  schema.BoolAttribute{Computed: true},
			"error":                                    schema.StringAttribute{Computed: true},
			"message":                                  schema.StringAttribute{Computed: true},
		},
	}
}
