// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_write_action

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XWriteActionDataSourceModel struct {
	ID                    types.String                                               `tfsdk:"id" path:"id,required"`
	Action                types.String                                               `tfsdk:"action" json:"action,computed"`
	Charged               types.Bool                                                 `tfsdk:"charged" json:"charged,computed"`
	ChargedCredits        types.String                                               `tfsdk:"charged_credits" json:"chargedCredits,computed"`
	ConfirmationAttempts  types.Int64                                                `tfsdk:"confirmation_attempts" json:"confirmationAttempts,computed"`
	ConfirmationCheckedAt timetypes.RFC3339                                          `tfsdk:"confirmation_checked_at" json:"confirmationCheckedAt,computed" format:"date-time"`
	ConfirmationSource    types.String                                               `tfsdk:"confirmation_source" json:"confirmationSource,computed"`
	ConfirmedAt           timetypes.RFC3339                                          `tfsdk:"confirmed_at" json:"confirmedAt,computed" format:"date-time"`
	CreatedAt             timetypes.RFC3339                                          `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Message               types.String                                               `tfsdk:"message" json:"message,computed"`
	MessageID             types.String                                               `tfsdk:"message_id" json:"messageId,computed"`
	Retryable             types.Bool                                                 `tfsdk:"retryable" json:"retryable,computed"`
	SendDispatched        types.Bool                                                 `tfsdk:"send_dispatched" json:"sendDispatched,computed"`
	SendDispatchedAt      timetypes.RFC3339                                          `tfsdk:"send_dispatched_at" json:"sendDispatchedAt,computed" format:"date-time"`
	Status                types.String                                               `tfsdk:"status" json:"status,computed"`
	TargetID              types.String                                               `tfsdk:"target_id" json:"targetId,computed"`
	TweetID               types.String                                               `tfsdk:"tweet_id" json:"tweetId,computed"`
	WriteActionID         types.String                                               `tfsdk:"write_action_id" json:"writeActionId,computed"`
	Media                 customfield.NestedObject[XWriteActionMediaDataSourceModel] `tfsdk:"media" json:"media,computed"`
}

type XWriteActionMediaDataSourceModel struct {
	XWriteActionCount types.Int64  `tfsdk:"x_write_action_count" json:"count,computed"`
	Credits           types.String `tfsdk:"credits" json:"credits,computed"`
	Kind              types.String `tfsdk:"kind" json:"kind,computed"`
	TotalBytes        types.String `tfsdk:"total_bytes" json:"totalBytes,computed"`
}
