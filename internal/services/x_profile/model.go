// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_profile

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XProfileModel struct {
	ID                    types.String                                      `tfsdk:"id" json:"id,computed"`
	Account               types.String                                      `tfsdk:"account" json:"account,required"`
	Description           types.String                                      `tfsdk:"description" json:"description,optional"`
	Location              types.String                                      `tfsdk:"location" json:"location,optional"`
	Name                  types.String                                      `tfsdk:"name" json:"name,optional"`
	URL                   types.String                                      `tfsdk:"url" json:"url,optional"`
	Action                types.String                                      `tfsdk:"action" json:"action,computed"`
	Charged               types.Bool                                        `tfsdk:"charged" json:"charged,computed"`
	ChargedCredits        types.String                                      `tfsdk:"charged_credits" json:"chargedCredits,computed"`
	CommunityID           types.String                                      `tfsdk:"community_id" json:"communityId,computed"`
	CommunityName         types.String                                      `tfsdk:"community_name" json:"communityName,computed"`
	CompletedAt           timetypes.RFC3339                                 `tfsdk:"completed_at" json:"completedAt,computed" format:"date-time"`
	ConfirmationAttempts  types.Int64                                       `tfsdk:"confirmation_attempts" json:"confirmationAttempts,computed"`
	ConfirmationCheckedAt timetypes.RFC3339                                 `tfsdk:"confirmation_checked_at" json:"confirmationCheckedAt,computed" format:"date-time"`
	ConfirmedAt           timetypes.RFC3339                                 `tfsdk:"confirmed_at" json:"confirmedAt,computed" format:"date-time"`
	CreatedAt             timetypes.RFC3339                                 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Error                 types.String                                      `tfsdk:"error" json:"error,computed"`
	ExpiresAt             timetypes.RFC3339                                 `tfsdk:"expires_at" json:"expiresAt,computed" format:"date-time"`
	Idempotent            types.Bool                                        `tfsdk:"idempotent" json:"idempotent,computed"`
	MediaID               types.String                                      `tfsdk:"media_id" json:"mediaId,computed"`
	MediaURL              types.String                                      `tfsdk:"media_url" json:"mediaUrl,computed"`
	Message               types.String                                      `tfsdk:"message" json:"message,computed"`
	MessageID             types.String                                      `tfsdk:"message_id" json:"messageId,computed"`
	Object                types.String                                      `tfsdk:"object" json:"object,computed"`
	PollAfterMs           types.Int64                                       `tfsdk:"poll_after_ms" json:"pollAfterMs,computed"`
	RequestHash           types.String                                      `tfsdk:"request_hash" json:"requestHash,computed"`
	RequestID             types.String                                      `tfsdk:"request_id" json:"requestId,computed"`
	ResultID              types.String                                      `tfsdk:"result_id" json:"resultId,computed"`
	Retryable             types.Bool                                        `tfsdk:"retryable" json:"retryable,computed"`
	SafeToRetry           types.Bool                                        `tfsdk:"safe_to_retry" json:"safeToRetry,computed"`
	SendDispatched        types.Bool                                        `tfsdk:"send_dispatched" json:"sendDispatched,computed"`
	SendDispatchedAt      timetypes.RFC3339                                 `tfsdk:"send_dispatched_at" json:"sendDispatchedAt,computed" format:"date-time"`
	Status                types.String                                      `tfsdk:"status" json:"status,computed"`
	StatusURL             types.String                                      `tfsdk:"status_url" json:"statusUrl,computed"`
	Success               types.Bool                                        `tfsdk:"success" json:"success,computed"`
	TargetID              types.String                                      `tfsdk:"target_id" json:"targetId,computed"`
	Terminal              types.Bool                                        `tfsdk:"terminal" json:"terminal,computed"`
	TweetID               types.String                                      `tfsdk:"tweet_id" json:"tweetId,computed"`
	UpdatedAt             timetypes.RFC3339                                 `tfsdk:"updated_at" json:"updatedAt,computed" format:"date-time"`
	WriteActionID         types.String                                      `tfsdk:"write_action_id" json:"writeActionId,computed"`
	Details               customfield.Map[jsontypes.Normalized]             `tfsdk:"details" json:"details,computed"`
	Media                 customfield.Map[jsontypes.Normalized]             `tfsdk:"media" json:"media,computed"`
	Billing               customfield.NestedObject[XProfileBillingModel]    `tfsdk:"billing" json:"billing,computed"`
	NextAction            customfield.NestedObject[XProfileNextActionModel] `tfsdk:"next_action" json:"nextAction,computed"`
	Request               customfield.NestedObject[XProfileRequestModel]    `tfsdk:"request" json:"request,computed"`
	Result                customfield.NestedObject[XProfileResultModel]     `tfsdk:"result" json:"result,computed"`
	Target                customfield.NestedObject[XProfileTargetModel]     `tfsdk:"target" json:"target,computed"`
}

func (m XProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m XProfileModel) MarshalJSONForUpdate(state XProfileModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type XProfileBillingModel struct {
	Charged        types.Bool   `tfsdk:"charged" json:"charged,computed"`
	ChargedCredits types.String `tfsdk:"charged_credits" json:"chargedCredits,computed"`
	PlannedCredits types.String `tfsdk:"planned_credits" json:"plannedCredits,computed"`
	Status         types.String `tfsdk:"status" json:"status,computed"`
}

type XProfileNextActionModel struct {
	Type                      types.String `tfsdk:"type" json:"type,computed"`
	AfterMs                   types.Int64  `tfsdk:"after_ms" json:"afterMs,computed"`
	RequiresNewIdempotencyKey types.Bool   `tfsdk:"requires_new_idempotency_key" json:"requiresNewIdempotencyKey,computed"`
	URL                       types.String `tfsdk:"url" json:"url,computed"`
}

type XProfileRequestModel struct {
	Hash    types.String                          `tfsdk:"hash" json:"hash,computed"`
	Payload customfield.Map[jsontypes.Normalized] `tfsdk:"payload" json:"payload,computed"`
}

type XProfileResultModel struct {
	ID    types.String `tfsdk:"id" json:"id,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
}

type XProfileTargetModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}
