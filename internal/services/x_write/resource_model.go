package x_write

import "github.com/hashicorp/terraform-plugin-framework/types"

type resourceModel struct {
	ID                      types.String `tfsdk:"id"`
	Action                  types.String `tfsdk:"action"`
	Account                 types.String `tfsdk:"account"`
	IdempotencyKey          types.String `tfsdk:"idempotency_key"`
	TargetID                types.String `tfsdk:"target_id"`
	PayloadJSON             types.String `tfsdk:"payload_json"`
	Object                  types.String `tfsdk:"object"`
	WriteActionID           types.String `tfsdk:"write_action_id"`
	Status                  types.String `tfsdk:"status"`
	Terminal                types.Bool   `tfsdk:"terminal"`
	Retryable               types.Bool   `tfsdk:"retryable"`
	SafeToRetry             types.Bool   `tfsdk:"safe_to_retry"`
	StatusURL               types.String `tfsdk:"status_url"`
	PollAfterMs             types.Int64  `tfsdk:"poll_after_ms"`
	Charged                 types.Bool   `tfsdk:"charged"`
	ChargedCredits          types.String `tfsdk:"charged_credits"`
	TargetIDTopLevel        types.String `tfsdk:"target_id_response"`
	BillingStatus           types.String `tfsdk:"billing_status"`
	BillingCharged          types.Bool   `tfsdk:"billing_charged"`
	BillingChargedCredits   types.String `tfsdk:"billing_charged_credits"`
	BillingPlannedCredits   types.String `tfsdk:"billing_planned_credits"`
	RequestHash             types.String `tfsdk:"request_hash"`
	RequestPayloadJSON      types.String `tfsdk:"request_payload_json"`
	RequestID               types.String `tfsdk:"request_id"`
	ResponseAccountID       types.String `tfsdk:"response_account_id"`
	ResponseAccountUsername types.String `tfsdk:"response_account_username"`
	ResultID                types.String `tfsdk:"result_id"`
	ResultState             types.String `tfsdk:"result_state"`
	ResultType              types.String `tfsdk:"result_type"`
	TargetIDResponse        types.String `tfsdk:"response_target_id"`
	TargetType              types.String `tfsdk:"response_target_type"`
	NextActionType          types.String `tfsdk:"next_action_type"`
	NextActionAfterMs       types.Int64  `tfsdk:"next_action_after_ms"`
	NextActionRequiresKey   types.Bool   `tfsdk:"next_action_requires_new_idempotency_key"`
	NextActionURL           types.String `tfsdk:"next_action_url"`
	SendDispatched          types.Bool   `tfsdk:"send_dispatched"`
	Success                 types.Bool   `tfsdk:"success"`
	Error                   types.String `tfsdk:"error"`
	Message                 types.String `tfsdk:"message"`
}

func (m *resourceModel) apply(action writeAction) {
	m.ID = types.StringValue(action.WriteActionID)
	m.Action = types.StringValue(action.Action)
	m.Object = types.StringValue(action.Object)
	m.WriteActionID = types.StringValue(action.WriteActionID)
	m.Status = types.StringValue(action.Status)
	m.Terminal = types.BoolValue(action.Terminal)
	m.Retryable = types.BoolValue(action.Retryable)
	m.SafeToRetry = types.BoolValue(action.SafeToRetry)
	m.StatusURL = types.StringValue(action.StatusURL)
	m.PollAfterMs = types.Int64Value(action.PollAfterMs)
	m.Charged = types.BoolValue(action.Charged)
	m.ChargedCredits = types.StringValue(action.ChargedCredits)
	m.TargetIDTopLevel = types.StringValue(action.TargetID)
	m.BillingStatus = types.StringValue(action.Billing.Status)
	m.BillingCharged = types.BoolValue(action.Billing.Charged)
	m.BillingChargedCredits = types.StringValue(action.Billing.ChargedCredits)
	m.BillingPlannedCredits = types.StringValue(action.Billing.PlannedCredits)
	m.RequestHash = types.StringValue(action.Request.Hash)
	m.RequestPayloadJSON = types.StringValue(string(action.Request.Payload))
	m.RequestID = types.StringValue(action.RequestID)
	m.ResponseAccountID = types.StringValue(action.Account.ID)
	m.ResponseAccountUsername = types.StringValue(action.Account.Username)
	m.ResultID = types.StringValue(action.Result.ID)
	m.ResultState = types.StringValue(action.Result.State)
	m.ResultType = types.StringValue(action.Result.Type)
	m.TargetIDResponse = types.StringValue(action.Target.ID)
	m.TargetType = types.StringValue(action.Target.Type)
	m.NextActionType = types.StringValue(action.NextAction.Type)
	m.NextActionAfterMs = types.Int64Value(action.NextAction.AfterMs)
	m.NextActionRequiresKey = types.BoolValue(action.NextAction.RequiresNewIdempotencyKey)
	m.NextActionURL = types.StringValue(action.NextAction.URL)
	m.SendDispatched = types.BoolValue(action.SendDispatched)
	m.Success = types.BoolValue(action.Success)
	m.Error = types.StringValue(action.Error)
	m.Message = types.StringValue(action.Message)
}
