// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
	"time"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/test_helpers"
	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

	_, err = operations[0].requestBody("@account", `{`)
	if err == nil || !strings.Contains(err.Error(), "JSON object") {
		t.Fatalf("expected malformed JSON error, got %v", err)
	}
	_, err = operations[0].requestBody("@account", `null`)
	if err == nil || !strings.Contains(err.Error(), "at least one") {
		t.Fatalf("expected null payload error, got %v", err)
	}
	if present([]any{}) {
		t.Fatal("empty list is present")
	}
	if !present(true) {
		t.Fatal("boolean value is not present")
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

func TestWritePayloadAttributesAreSensitive(t *testing.T) {
	for _, op := range operations {
		op := op
		t.Run(op.action, func(t *testing.T) {
			t.Parallel()
			writeSchema := resourceSchema(context.Background(), op)
			for _, name := range []string{"payload_json", "request_payload_json"} {
				attribute, ok := writeSchema.Attributes[name].(schema.StringAttribute)
				if !ok {
					t.Fatalf("%s is not a string attribute", name)
				}
				if !attribute.Sensitive {
					t.Errorf("%s must be sensitive", name)
				}
			}
		})
	}
}

func TestCanonicalPollPathPreservesSDKBasePath(t *testing.T) {
	const writeActionID = "write-action-123"
	var requestedPath string
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestedPath = request.URL.Path
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	path, err := canonicalPollPath(server.URL+"/api/v1/x/write-actions/"+writeActionID, writeActionID)
	if err != nil {
		t.Fatalf("canonicalPollPath: %v", err)
	}
	client := xtwitterscraper.NewClient(option.WithBaseURL(server.URL + "/api/v1/"))
	var response map[string]any
	if err := client.Get(context.Background(), path, nil, &response, option.WithMaxRetries(0)); err != nil {
		t.Fatalf("Client.Get: %v", err)
	}
	assertEqual(t, "relative poll path", path, "x/write-actions/"+writeActionID)
	assertEqual(t, "SDK request path", requestedPath, "/api/v1/x/write-actions/"+writeActionID)
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

func TestCanonicalResponseValidationRejectsContractViolations(t *testing.T) {
	t.Parallel()

	valid := validCanonicalAction()
	cases := []struct {
		name      string
		mutate    func(*writeAction)
		wantError string
	}{
		{name: "object", mutate: func(action *writeAction) { action.Object = "other" }, wantError: "unexpected object"},
		{name: "missing ID", mutate: func(action *writeAction) { action.ID = "" }, wantError: "omitted id"},
		{name: "different IDs", mutate: func(action *writeAction) { action.ID = "other" }, wantError: "differ"},
		{name: "action", mutate: func(action *writeAction) { action.Action = "other" }, wantError: "does not match"},
		{name: "status", mutate: func(action *writeAction) { action.Status = "other" }, wantError: "unknown status"},
		{name: "terminal", mutate: func(action *writeAction) { action.Terminal = false }, wantError: "terminal=false disagree"},
		{name: "billing status", mutate: func(action *writeAction) { action.Billing.Status = "other" }, wantError: "unknown billing status"},
		{name: "billing fields", mutate: func(action *writeAction) { action.Charged = false }, wantError: "billing compatibility"},
		{name: "target fields", mutate: func(action *writeAction) { action.TargetID = "other" }, wantError: "target compatibility"},
		{name: "retry guidance", mutate: func(action *writeAction) { action.SafeToRetry = true }, wantError: "dispatched write safe to retry"},
		{name: "new key guidance", mutate: func(action *writeAction) {
			action.NextAction.RequiresNewIdempotencyKey = true
		}, wantError: "new Idempotency-Key"},
		{name: "success", mutate: func(action *writeAction) { action.Success = false }, wantError: "success=false disagree"},
		{name: "status URL", mutate: func(action *writeAction) { action.StatusURL = "" }, wantError: "omitted statusUrl"},
	}

	for _, testCase := range cases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			action := valid
			testCase.mutate(&action)
			err := validateAction(operations[0], action)
			if err == nil || !strings.Contains(err.Error(), testCase.wantError) {
				t.Fatalf("validateAction() error = %v, want text %q", err, testCase.wantError)
			}
		})
	}
}

func TestExecuteWriteRejectsInvalidRequests(t *testing.T) {
	t.Parallel()

	client := xtwitterscraper.NewClient(option.WithBaseURL("https://example.invalid/"))
	cases := []struct {
		name      string
		operation operation
		request   writeRequest
		wantError string
	}{
		{
			name:      "empty account",
			operation: operations[0],
			request:   writeRequest{idempotencyKey: "test-key", payloadJSON: `{"text":"hello"}`},
			wantError: "account must not be empty",
		},
		{
			name:      "invalid idempotency key",
			operation: operations[0],
			request:   writeRequest{account: "@account", idempotencyKey: "not valid", payloadJSON: `{"text":"hello"}`},
			wantError: "visible ASCII",
		},
		{
			name:      "missing target",
			operation: operations[1],
			request:   writeRequest{account: "@account", idempotencyKey: "test-key"},
			wantError: "target_id is required",
		},
	}
	for _, testCase := range cases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			_, err := executeWrite(context.Background(), &client, testCase.operation, testCase.request)
			if err == nil || !strings.Contains(err.Error(), testCase.wantError) {
				t.Fatalf("executeWrite() error = %v, want text %q", err, testCase.wantError)
			}
		})
	}
}

func TestPollingFailuresPreserveSafeRecoveryGuidance(t *testing.T) {
	t.Parallel()

	valid := validCanonicalAction()
	valid.Status = "accepted"
	valid.Terminal = false
	valid.Success = false
	valid.Billing.Status = "pending"
	valid.PollAfterMs = 1_000

	cancelledContext, cancel := context.WithCancel(context.Background())
	cancel()
	client := xtwitterscraper.NewClient(option.WithBaseURL("https://example.invalid/"))
	_, err := pollUntilTerminal(cancelledContext, &client, operations[0], valid)
	if err == nil || !strings.Contains(err.Error(), "still non-terminal") {
		t.Fatalf("cancelled poll error = %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		http.Error(writer, `{"error":"unavailable"}`, http.StatusServiceUnavailable)
	}))
	t.Cleanup(server.Close)
	valid.StatusURL = server.URL + "/x/write-actions/" + valid.WriteActionID
	valid.PollAfterMs = 0
	client = xtwitterscraper.NewClient(
		option.WithBaseURL(server.URL+"/"),
		option.WithMaxRetries(0),
	)
	_, err = pollUntilTerminal(context.Background(), &client, operations[0], valid)
	if err == nil || !strings.Contains(err.Error(), "do not redispatch") {
		t.Fatalf("failed poll error = %v", err)
	}
}

func TestPollRejectsInvalidCanonicalResponse(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{}`))
	}))
	t.Cleanup(server.Close)
	action := validCanonicalAction()
	action.Status = "accepted"
	action.Terminal = false
	action.Success = false
	action.Billing.Status = "pending"
	action.StatusURL = server.URL + "/x/write-actions/" + action.WriteActionID
	client := xtwitterscraper.NewClient(
		option.WithBaseURL(server.URL+"/"),
		option.WithMaxRetries(0),
	)
	_, err := pollUntilTerminal(context.Background(), &client, operations[0], action)
	if err == nil || !strings.Contains(err.Error(), "unexpected object") {
		t.Fatalf("invalid canonical poll error = %v", err)
	}
}

func TestCanonicalPollPathRejectsMalformedURL(t *testing.T) {
	t.Parallel()

	_, err := canonicalPollPath("https://example.com/%", "write-action-123")
	if err == nil || !strings.Contains(err.Error(), "invalid statusUrl") {
		t.Fatalf("canonicalPollPath() error = %v", err)
	}
}

func TestPollingRejectsMalformedStatusURL(t *testing.T) {
	t.Parallel()

	action := validCanonicalAction()
	action.Status = "accepted"
	action.Terminal = false
	action.Success = false
	action.Billing.Status = "pending"
	action.StatusURL = "https://example.com/%"
	client := xtwitterscraper.NewClient(option.WithBaseURL("https://example.invalid/"))
	_, err := pollUntilTerminal(context.Background(), &client, operations[0], action)
	if err == nil || !strings.Contains(err.Error(), "invalid statusUrl") {
		t.Fatalf("pollUntilTerminal() error = %v", err)
	}
}

func TestWaitForPollNormalizesDelayAndHonorsCancellation(t *testing.T) {
	t.Parallel()

	if err := waitForPoll(context.Background(), -1); err != nil {
		t.Fatalf("negative polling delay: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := waitForPoll(ctx, int64(maxPollDuration/time.Millisecond)+1); err == nil {
		t.Fatal("cancelled maximum polling delay returned no error")
	}
}

func validCanonicalAction() writeAction {
	action := writeAction{
		Object:         "x_write_action",
		ID:             "write-action-123",
		WriteActionID:  "write-action-123",
		Action:         operations[0].action,
		Status:         "success",
		Terminal:       true,
		StatusURL:      "https://example.com/x/write-actions/write-action-123",
		Charged:        true,
		ChargedCredits: "10",
		TargetID:       "target-123",
		SendDispatched: true,
		Success:        true,
	}
	action.Billing.Status = "charged"
	action.Billing.Charged = true
	action.Billing.ChargedCredits = "10"
	action.Target.ID = "target-123"
	return action
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
