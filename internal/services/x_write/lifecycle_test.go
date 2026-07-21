package x_write

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestAllCanonicalWritesDispatchOnceAndPollUntilTerminal(t *testing.T) {
	for _, op := range operations {
		op := op
		t.Run(op.action, func(t *testing.T) {
			t.Parallel()

			const (
				requestAccount = "@request-account"
				responseID     = "account-123"
				responseUser   = "response-account"
				writeActionID  = "write-action-123"
				idempotencyKey = "test-key-123"
				targetID       = "target/123"
			)

			var dispatchCount atomic.Int64
			var pollCount atomic.Int64
			var capturedMethod string
			var capturedPath string
			var capturedKey string
			var capturedBody map[string]any

			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.Header().Set("Content-Type", "application/json")
				if strings.HasPrefix(request.URL.Path, "/x/write-actions/") {
					pollCount.Add(1)
					writeCanonicalAction(writer, op.action, serverURL(request)+"/x/write-actions/"+writeActionID, true)
					return
				}

				dispatchCount.Add(1)
				capturedMethod = request.Method
				capturedPath = request.URL.EscapedPath()
				capturedKey = request.Header.Get("Idempotency-Key")
				body, err := io.ReadAll(request.Body)
				if err != nil {
					t.Errorf("read request body: %v", err)
				}
				if err := json.Unmarshal(body, &capturedBody); err != nil {
					t.Errorf("decode request body: %v", err)
				}
				writer.WriteHeader(http.StatusAccepted)
				writeCanonicalAction(writer, op.action, serverURL(request)+"/x/write-actions/"+writeActionID, false)
			}))
			defer server.Close()

			client := xtwitterscraper.NewClient(
				option.WithBaseURL(server.URL+"/"),
				option.WithAPIKey("test-api-key"),
			)
			payload, err := json.Marshal(op.testPayload)
			if err != nil {
				t.Fatalf("marshal payload: %v", err)
			}
			action, err := executeWrite(context.Background(), &client, op, writeRequest{
				account:        requestAccount,
				idempotencyKey: idempotencyKey,
				targetID:       targetID,
				payloadJSON:    string(payload),
			})
			if err != nil {
				t.Fatalf("execute write: %v", err)
			}

			assertEqual(t, "one dispatch", dispatchCount.Load(), int64(1))
			assertEqual(t, "one poll", pollCount.Load(), int64(1))
			assertEqual(t, "method", capturedMethod, op.method)
			assertEqual(t, "path", capturedPath, "/"+op.path(targetID))
			assertEqual(t, "idempotency key", capturedKey, idempotencyKey)
			assertEqual(t, "request account", capturedBody["account"], requestAccount)
			assertEqual(t, "write ID", action.ID, writeActionID)
			assertEqual(t, "writeActionId", action.WriteActionID, writeActionID)
			assertEqual(t, "action", action.Action, op.action)
			assertEqual(t, "terminal", action.Terminal, true)
			assertEqual(t, "status", action.Status, "success")
			assertEqual(t, "success", action.Success, true)
			assertEqual(t, "safe retry", action.SafeToRetry, false)
			assertEqual(t, "billing", action.Billing.Status, "charged")
			assertEqual(t, "billing charged", action.Billing.Charged, true)
			assertEqual(t, "billing charged credits", action.Billing.ChargedCredits, "10")
			assertEqual(t, "billing planned credits", action.Billing.PlannedCredits, "10")
			assertEqual(t, "request hash", action.Request.Hash, "request-hash")
			assertEqual(t, "request payload", string(action.Request.Payload), `{"account":"sanitized"}`)
			assertEqual(t, "result ID", action.Result.ID, "result-123")
			assertEqual(t, "target ID", action.Target.ID, "target-123")
			assertEqual(t, "target type", action.Target.Type, "tweet")
			assertEqual(t, "next action type", action.NextAction.Type, "verify_result")
			assertEqual(t, "response account ID", action.Account.ID, responseID)
			assertEqual(t, "response account username", action.Account.Username, responseUser)
			assertEqual(t, "request account remains separate", requestAccount, "@request-account")
			for field, expected := range op.testPayload {
				assertEqual(t, "payload "+field, capturedBody[field], expected)
			}
		})
	}
}

func TestWriteDispatchDoesNotBlindRetryUnknownState(t *testing.T) {
	var dispatchCount atomic.Int64
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		dispatchCount.Add(1)
		http.Error(writer, `{"error":"unavailable"}`, http.StatusServiceUnavailable)
	}))
	defer server.Close()

	client := xtwitterscraper.NewClient(option.WithBaseURL(server.URL + "/"))
	_, err := executeWrite(context.Background(), &client, operations[0], writeRequest{
		account:        "@request-account",
		idempotencyKey: "test-key-123",
		payloadJSON:    `{"text":"hello"}`,
	})
	if err == nil {
		t.Fatal("expected dispatch error")
	}
	if !strings.Contains(err.Error(), "same Idempotency-Key") {
		t.Fatalf("error does not preserve exact-replay guidance: %v", err)
	}
	assertEqual(t, "dispatch attempts", dispatchCount.Load(), int64(1))
}

func TestOperationValidationRejectsAccountCollisionAndInvalidFields(t *testing.T) {
	_, err := operations[0].requestBody("@account", `{"account":"other","text":"hello"}`)
	if err == nil || !strings.Contains(err.Error(), "must not contain account") {
		t.Fatalf("expected account collision error, got %v", err)
	}

	_, err = operations[0].requestBody("@account", `{"typo":"hello"}`)
	if err == nil || !strings.Contains(err.Error(), "not valid") {
		t.Fatalf("expected invalid field error, got %v", err)
	}

	_, err = operations[0].requestBody("@account", `{}`)
	if err == nil || !strings.Contains(err.Error(), "at least one") {
		t.Fatalf("expected missing content error, got %v", err)
	}
}

func TestOperationRegistryIsCompleteAndUnique(t *testing.T) {
	if len(operations) != 18 {
		t.Fatalf("operation count = %d, want 18", len(operations))
	}
	actions := map[string]bool{}
	resources := map[string]bool{}
	for _, op := range operations {
		if actions[op.action] {
			t.Errorf("duplicate action %s", op.action)
		}
		if resources[op.resourceSuffix] {
			t.Errorf("duplicate resource suffix %s", op.resourceSuffix)
		}
		actions[op.action] = true
		resources[op.resourceSuffix] = true
	}
}

func TestAllWriteResourceSchemasMatchStableStateModel(t *testing.T) {
	for _, op := range operations {
		op := op
		t.Run(op.action, func(t *testing.T) {
			t.Parallel()
			errs := test_helpers.ValidateResourceModelSchemaIntegrity((*resourceModel)(nil), resourceSchema(context.Background(), op))
			errs.Report(t)
		})
	}
}

func TestIdempotencyKeyValidationMatchesCanonicalContract(t *testing.T) {
	valid := []string{"!", "test-key-123", strings.Repeat("x", 255)}
	for _, key := range valid {
		if !validIdempotencyKey(key) {
			t.Errorf("expected key of length %d to be valid", len(key))
		}
	}
	invalid := []string{"", "contains space", "contains\nnewline", strings.Repeat("x", 256), "non-ascii-é"}
	for _, key := range invalid {
		if validIdempotencyKey(key) {
			t.Errorf("expected %q to be invalid", key)
		}
	}
}

func TestStableStatePreservesRequestAndResponseAccountsSeparately(t *testing.T) {
	model := resourceModel{
		Account:        types.StringValue("@request-account"),
		IdempotencyKey: types.StringValue("test-key-123"),
		PayloadJSON:    types.StringValue(`{"text":"hello"}`),
	}
	var action writeAction
	action.WriteActionID = "write-action-123"
	action.Account.ID = "response-account-id"
	action.Account.Username = "response-account"

	model.apply(action)
	assertEqual(t, "request account", model.Account.ValueString(), "@request-account")
	assertEqual(t, "response account ID", model.ResponseAccountID.ValueString(), "response-account-id")
	assertEqual(t, "response account username", model.ResponseAccountUsername.ValueString(), "response-account")
	assertEqual(t, "stable state ID", model.ID.ValueString(), "write-action-123")
	assertEqual(t, "write action ID", model.WriteActionID.ValueString(), "write-action-123")
	assertEqual(t, "idempotency key", model.IdempotencyKey.ValueString(), "test-key-123")
	assertEqual(t, "request payload", model.PayloadJSON.ValueString(), `{"text":"hello"}`)
}

func TestCanonicalPollPathRejectsDifferentWrite(t *testing.T) {
	_, err := canonicalPollPath("https://xquik.com/api/v1/x/write-actions/other", "expected")
	if err == nil {
		t.Fatal("expected mismatched polling URL error")
	}
}

func TestCanonicalResponseRejectsUnsafeRetryGuidance(t *testing.T) {
	var action writeAction
	action.Object = "x_write_action"
	action.ID = "write-action-123"
	action.WriteActionID = action.ID
	action.Action = operations[0].action
	action.Status = "failed"
	action.Terminal = true
	action.StatusURL = "https://xquik.com/api/v1/x/write-actions/write-action-123"
	action.SafeToRetry = true
	action.SendDispatched = true
	action.Success = false
	action.Billing.Status = "not_charged"

	err := validateAction(operations[0], action)
	if err == nil || !strings.Contains(err.Error(), "dispatched write safe to retry") {
		t.Fatalf("expected unsafe retry guidance error, got %v", err)
	}
}

func writeCanonicalAction(writer http.ResponseWriter, action string, statusURL string, terminal bool) {
	status := "accepted"
	billingStatus := "pending"
	result := map[string]any{}
	nextActionType := "poll"
	if terminal {
		status = "success"
		billingStatus = "charged"
		result = map[string]any{"id": "result-123", "state": "confirmed", "type": "state_change"}
		nextActionType = "verify_result"
	}
	response := map[string]any{
		"object": "x_write_action", "id": "write-action-123", "writeActionId": "write-action-123",
		"action": action, "status": status, "terminal": terminal, "retryable": false,
		"safeToRetry": false, "statusUrl": statusURL, "pollAfterMs": 0,
		"charged": terminal, "chargedCredits": "10", "targetId": "target-123",
		"billing": map[string]any{"status": billingStatus, "charged": terminal, "chargedCredits": "10", "plannedCredits": "10"},
		"request": map[string]any{"hash": "request-hash", "payload": map[string]any{"account": "sanitized"}},
		"account": map[string]any{"id": "account-123", "username": "response-account"}, "result": result,
		"target":     map[string]any{"id": "target-123", "type": "tweet"},
		"nextAction": map[string]any{"type": nextActionType, "afterMs": 0, "requiresNewIdempotencyKey": false, "url": statusURL},
		"requestId":  "request-123", "sendDispatched": true, "success": terminal,
	}
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		panic(fmt.Sprintf("encode response: %v", err))
	}
}

func serverURL(request *http.Request) string {
	return "http://" + request.Host
}

func assertEqual[T comparable](t *testing.T, name string, got T, want T) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %v, want %v", name, got, want)
	}
}
