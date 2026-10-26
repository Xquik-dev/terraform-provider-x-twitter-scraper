// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xwriteaction

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XWriteActionDataSourceModel struct {
	ID                    types.String                                                    `tfsdk:"id" path:"id,required"`
	Action                types.String                                                    `tfsdk:"action" json:"action" api:"computed"`
	Charged               types.Bool                                                      `tfsdk:"charged" json:"charged" api:"computed"`
	ChargedCredits        types.String                                                    `tfsdk:"charged_credits" json:"chargedCredits" api:"computed"`
	CommunityID           types.String                                                    `tfsdk:"community_id" json:"communityId" api:"computed"`
	CommunityName         types.String                                                    `tfsdk:"community_name" json:"communityName" api:"computed"`
	CompletedAt           timetypes.RFC3339                                               `tfsdk:"completed_at" json:"completedAt" api:"computed" format:"date-time"`
	ConfirmationAttempts  types.Int64                                                     `tfsdk:"confirmation_attempts" json:"confirmationAttempts" api:"computed"`
	ConfirmationCheckedAt timetypes.RFC3339                                               `tfsdk:"confirmation_checked_at" json:"confirmationCheckedAt" api:"computed" format:"date-time"`
	ConfirmedAt           timetypes.RFC3339                                               `tfsdk:"confirmed_at" json:"confirmedAt" api:"computed" format:"date-time"`
	CreatedAt             timetypes.RFC3339                                               `tfsdk:"created_at" json:"createdAt" api:"computed" format:"date-time"`
	Error                 types.String                                                    `tfsdk:"error" json:"error" api:"computed"`
	ExpiresAt             timetypes.RFC3339                                               `tfsdk:"expires_at" json:"expiresAt" api:"computed" format:"date-time"`
	Idempotent            types.Bool                                                      `tfsdk:"idempotent" json:"idempotent" api:"computed"`
	MediaID               types.String                                                    `tfsdk:"media_id" json:"mediaId" api:"computed"`
	MediaURL              types.String                                                    `tfsdk:"media_url" json:"mediaUrl" api:"computed"`
	Message               types.String                                                    `tfsdk:"message" json:"message" api:"computed"`
	MessageID             types.String                                                    `tfsdk:"message_id" json:"messageId" api:"computed"`
	Object                types.String                                                    `tfsdk:"object" json:"object" api:"computed"`
	PollAfterMs           types.Int64                                                     `tfsdk:"poll_after_ms" json:"pollAfterMs" api:"computed"`
	RequestHash           types.String                                                    `tfsdk:"request_hash" json:"requestHash" api:"computed"`
	RequestID             types.String                                                    `tfsdk:"request_id" json:"requestId" api:"computed"`
	ResultID              types.String                                                    `tfsdk:"result_id" json:"resultId" api:"computed"`
	Retryable             types.Bool                                                      `tfsdk:"retryable" json:"retryable" api:"computed"`
	SafeToRetry           types.Bool                                                      `tfsdk:"safe_to_retry" json:"safeToRetry" api:"computed"`
	SendDispatched        types.Bool                                                      `tfsdk:"send_dispatched" json:"sendDispatched" api:"computed"`
	SendDispatchedAt      timetypes.RFC3339                                               `tfsdk:"send_dispatched_at" json:"sendDispatchedAt" api:"computed" format:"date-time"`
	Status                types.String                                                    `tfsdk:"status" json:"status" api:"computed"`
	StatusURL             types.String                                                    `tfsdk:"status_url" json:"statusUrl" api:"computed"`
	Success               types.Bool                                                      `tfsdk:"success" json:"success" api:"computed"`
	TargetID              types.String                                                    `tfsdk:"target_id" json:"targetId" api:"computed"`
	Terminal              types.Bool                                                      `tfsdk:"terminal" json:"terminal" api:"computed"`
	TweetID               types.String                                                    `tfsdk:"tweet_id" json:"tweetId" api:"computed"`
	UpdatedAt             timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updatedAt" api:"computed" format:"date-time"`
	WriteActionID         types.String                                                    `tfsdk:"write_action_id" json:"writeActionId" api:"computed"`
	Details               customfield.Map[jsontypes.Normalized]                           `tfsdk:"details" json:"details" api:"computed"`
	Media                 customfield.Map[jsontypes.Normalized]                           `tfsdk:"media" json:"media" api:"computed"`
	Account               customfield.NestedObject[XWriteActionAccountDataSourceModel]    `tfsdk:"account" json:"account" api:"computed"`
	Billing               customfield.NestedObject[XWriteActionBillingDataSourceModel]    `tfsdk:"billing" json:"billing" api:"computed"`
	NextAction            customfield.NestedObject[XWriteActionNextActionDataSourceModel] `tfsdk:"next_action" json:"nextAction" api:"computed"`
	Request               customfield.NestedObject[XWriteActionRequestDataSourceModel]    `tfsdk:"request" json:"request" api:"computed"`
	Result                customfield.NestedObject[XWriteActionResultDataSourceModel]     `tfsdk:"result" json:"result" api:"computed"`
	Target                customfield.NestedObject[XWriteActionTargetDataSourceModel]     `tfsdk:"target" json:"target" api:"computed"`
}

type XWriteActionAccountDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id" api:"computed"`
	Username types.String `tfsdk:"username" json:"username" api:"computed"`
}

type XWriteActionBillingDataSourceModel struct {
	Charged        types.Bool   `tfsdk:"charged" json:"charged" api:"computed"`
	ChargedCredits types.String `tfsdk:"charged_credits" json:"chargedCredits" api:"computed"`
	PlannedCredits types.String `tfsdk:"planned_credits" json:"plannedCredits" api:"computed"`
	Status         types.String `tfsdk:"status" json:"status" api:"computed"`
}

type XWriteActionNextActionDataSourceModel struct {
	Type                      types.String `tfsdk:"type" json:"type" api:"computed"`
	AfterMs                   types.Int64  `tfsdk:"after_ms" json:"afterMs" api:"computed"`
	RequiresNewIdempotencyKey types.Bool   `tfsdk:"requires_new_idempotency_key" json:"requiresNewIdempotencyKey" api:"computed"`
	URL                       types.String `tfsdk:"url" json:"url" api:"computed"`
}

type XWriteActionRequestDataSourceModel struct {
	Hash    types.String                          `tfsdk:"hash" json:"hash" api:"computed"`
	Payload customfield.Map[jsontypes.Normalized] `tfsdk:"payload" json:"payload" api:"computed"`
}

type XWriteActionResultDataSourceModel struct {
	ID    types.String `tfsdk:"id" json:"id" api:"computed"`
	State types.String `tfsdk:"state" json:"state" api:"computed"`
	Type  types.String `tfsdk:"type" json:"type" api:"computed"`
}

type XWriteActionTargetDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id" api:"computed"`
	Type types.String `tfsdk:"type" json:"type" api:"computed"`
}
