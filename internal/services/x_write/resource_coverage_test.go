// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package xwrite

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestWriteResourceLifecycle(t *testing.T) {
	t.Parallel()

	var readCount atomic.Int64
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		statusURL := serverURL(request) + "/x/write-actions/write-action-123"
		if request.Method == http.MethodPost {
			writeCanonicalAction(writer, operations[0].action, statusURL, true)
			return
		}
		writeCanonicalAction(writer, operations[0].action, statusURL, readCount.Add(1) > 1)
	}))
	t.Cleanup(server.Close)

	client := xtwitterscraper.NewClient(
		option.WithAPIKey("test-api-key"),
		option.WithBaseURL(server.URL+"/"),
		option.WithMaxRetries(0),
	)
	resourceUnderTest := writeResourceFor(t, operations[0].action)
	ctx := context.Background()

	var metadataResponse resource.MetadataResponse
	resourceUnderTest.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "x-twitter-scraper"}, &metadataResponse)
	assertEqual(t, "resource type", metadataResponse.TypeName, "x-twitter-scraper_x_tweet")

	var schemaResponse resource.SchemaResponse
	resourceUnderTest.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
	if len(schemaResponse.Schema.Attributes) == 0 {
		t.Fatal("resource schema has no attributes")
	}

	var nilConfigureResponse resource.ConfigureResponse
	resourceUnderTest.Configure(ctx, resource.ConfigureRequest{}, &nilConfigureResponse)
	if nilConfigureResponse.Diagnostics.HasError() {
		t.Fatalf("nil configuration: %v", nilConfigureResponse.Diagnostics)
	}
	var invalidConfigureResponse resource.ConfigureResponse
	resourceUnderTest.Configure(ctx, resource.ConfigureRequest{ProviderData: "invalid"}, &invalidConfigureResponse)
	if !invalidConfigureResponse.Diagnostics.HasError() {
		t.Fatal("resource accepted invalid provider data")
	}
	var configureResponse resource.ConfigureResponse
	resourceUnderTest.Configure(ctx, resource.ConfigureRequest{ProviderData: &client}, &configureResponse)
	if configureResponse.Diagnostics.HasError() {
		t.Fatalf("configure resource: %v", configureResponse.Diagnostics)
	}

	raw := rawWriteObject(schemaResponse.Schema.Type().TerraformType(ctx), map[string]string{
		"account":         "@request-account",
		"idempotency_key": "test-key-123",
		"payload_json":    `{"text":"hello"}`,
	})
	plan := tfsdk.Plan{Raw: raw, Schema: schemaResponse.Schema}
	var createResponse resource.CreateResponse
	createResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Create(ctx, resource.CreateRequest{Plan: plan}, &createResponse)
	if createResponse.Diagnostics.HasError() {
		t.Fatalf("create write: %v", createResponse.Diagnostics)
	}

	var created resourceModel
	if diagnostics := createResponse.State.Get(ctx, &created); diagnostics.HasError() {
		t.Fatalf("read created state: %v", diagnostics)
	}
	assertEqual(t, "write action ID", created.WriteActionID.ValueString(), "write-action-123")
	assertEqual(t, "create status", created.Status.ValueString(), "success")

	var readResponse resource.ReadResponse
	readResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Read(ctx, resource.ReadRequest{State: createResponse.State}, &readResponse)
	if readResponse.Diagnostics.HasError() {
		t.Fatalf("refresh write: %v", readResponse.Diagnostics)
	}
	assertEqual(t, "refresh requests", readCount.Load(), int64(2))

	var refreshed resourceModel
	if diagnostics := readResponse.State.Get(ctx, &refreshed); diagnostics.HasError() {
		t.Fatalf("read refreshed state: %v", diagnostics)
	}
	assertEqual(t, "refresh status", refreshed.Status.ValueString(), "success")

	var updateResponse resource.UpdateResponse
	resourceUnderTest.Update(ctx, resource.UpdateRequest{}, &updateResponse)
	if !updateResponse.Diagnostics.HasError() {
		t.Fatal("write resource accepted an update")
	}
	var deleteResponse resource.DeleteResponse
	resourceUnderTest.Delete(ctx, resource.DeleteRequest{}, &deleteResponse)
	if deleteResponse.Diagnostics.HasError() {
		t.Fatalf("delete state: %v", deleteResponse.Diagnostics)
	}
}

func TestWriteResourceRejectsInvalidResponsesAndState(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		if request.Method == http.MethodPost {
			_, _ = writer.Write([]byte(`{}`))
			return
		}
		http.Error(writer, `{"error":"unavailable"}`, http.StatusServiceUnavailable)
	}))
	t.Cleanup(server.Close)

	client := xtwitterscraper.NewClient(
		option.WithBaseURL(server.URL+"/"),
		option.WithMaxRetries(0),
	)
	resourceUnderTest := writeResourceFor(t, operations[0].action)
	resourceUnderTest.client = &client
	resourceSchema := resourceSchema(context.Background(), operations[0])

	plan := tfsdk.Plan{
		Raw: rawWriteObject(resourceSchema.Type().TerraformType(context.Background()), map[string]string{
			"account":         "@request-account",
			"idempotency_key": "test-key-123",
			"payload_json":    `{"text":"hello"}`,
		}),
		Schema: resourceSchema,
	}
	var createResponse resource.CreateResponse
	createResponse.State.Schema = resourceSchema
	resourceUnderTest.Create(context.Background(), resource.CreateRequest{Plan: plan}, &createResponse)
	if !createResponse.Diagnostics.HasError() {
		t.Fatal("resource accepted an invalid create response")
	}

	state := tfsdk.State{
		Raw: rawWriteObject(resourceSchema.Type().TerraformType(context.Background()), map[string]string{
			"status_url":      server.URL + "/x/write-actions/write-action-123",
			"write_action_id": "write-action-123",
		}),
		Schema: resourceSchema,
	}
	var readResponse resource.ReadResponse
	readResponse.State.Schema = resourceSchema
	resourceUnderTest.Read(context.Background(), resource.ReadRequest{State: state}, &readResponse)
	if !readResponse.Diagnostics.HasError() {
		t.Fatal("resource accepted a failed refresh")
	}

	invalidState := tfsdk.State{
		Raw: rawWriteObject(resourceSchema.Type().TerraformType(context.Background()), map[string]string{
			"status_url":      "https://example.invalid/not-canonical",
			"write_action_id": "write-action-123",
		}),
		Schema: resourceSchema,
	}
	var invalidStateResponse resource.ReadResponse
	invalidStateResponse.State.Schema = resourceSchema
	resourceUnderTest.Read(context.Background(), resource.ReadRequest{State: invalidState}, &invalidStateResponse)
	if !invalidStateResponse.Diagnostics.HasError() {
		t.Fatal("resource accepted an invalid polling URL")
	}
}

func TestWriteResourceRejectsInvalidPlan(t *testing.T) {
	t.Parallel()

	resourceUnderTest := writeResourceFor(t, operations[0].action)
	resourceSchema := resourceSchema(context.Background(), operations[0])
	plan := tfsdk.Plan{
		Raw:    tftypes.NewValue(tftypes.String, "not-an-object"),
		Schema: resourceSchema,
	}
	var response resource.CreateResponse
	response.State.Schema = resourceSchema
	resourceUnderTest.Create(context.Background(), resource.CreateRequest{Plan: plan}, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("resource accepted an invalid plan")
	}
}

func TestWriteResourceReportsTerminalFailure(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		response := canonicalAction(operations[0].action, serverURL(request)+"/x/write-actions/write-action-123")
		response.Status = "failed"
		response.Terminal = true
		response.Success = false
		response.SendDispatched = false
		response.Billing.Status = "not_charged"
		response.Billing.Charged = false
		response.Billing.ChargedCredits = ""
		response.Billing.PlannedCredits = "10"
		response.Charged = false
		response.ChargedCredits = ""
		if err := json.NewEncoder(writer).Encode(response); err != nil {
			t.Errorf("encode response: %v", err)
		}
	}))
	t.Cleanup(server.Close)

	client := xtwitterscraper.NewClient(
		option.WithBaseURL(server.URL+"/"),
		option.WithMaxRetries(0),
	)
	resourceUnderTest := writeResourceFor(t, operations[0].action)
	resourceUnderTest.client = &client
	resourceSchema := resourceSchema(context.Background(), operations[0])
	plan := tfsdk.Plan{
		Raw: rawWriteObject(resourceSchema.Type().TerraformType(context.Background()), map[string]string{
			"account":         "@request-account",
			"idempotency_key": "test-key-123",
			"payload_json":    `{"text":"hello"}`,
		}),
		Schema: resourceSchema,
	}
	var response resource.CreateResponse
	response.State.Schema = resourceSchema
	resourceUnderTest.Create(context.Background(), resource.CreateRequest{Plan: plan}, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("failed write produced no diagnostic")
	}
	if !strings.Contains(response.Diagnostics.Errors()[0].Summary(), "write failed") {
		t.Fatalf("failed write diagnostic = %v", response.Diagnostics)
	}
}

func TestUnknownWriteResourcePanics(t *testing.T) {
	t.Parallel()

	defer func() {
		if recover() == nil {
			t.Fatal("unknown write action did not panic")
		}
	}()
	_ = resourceFor("unknown")
}

func canonicalAction(action string, statusURL string) writeAction {
	response := writeAction{
		Object: "x_write_action", ID: "write-action-123", WriteActionID: "write-action-123",
		Action: action, Status: "success", Terminal: true, StatusURL: statusURL,
		Charged: true, ChargedCredits: "10", TargetID: "target-123",
		RequestID: "request-123", SendDispatched: true, Success: true,
	}
	response.Billing.Status = "charged"
	response.Billing.Charged = true
	response.Billing.ChargedCredits = "10"
	response.Billing.PlannedCredits = "10"
	response.Request.Hash = "request-hash"
	response.Request.Payload = json.RawMessage(`{"account":"sanitized"}`)
	response.Account.ID = "account-123"
	response.Account.Username = "response-account"
	response.Result.ID, response.Result.State, response.Result.Type = "result-123", "confirmed", "state_change"
	response.Target.ID, response.Target.Type = "target-123", "tweet"
	response.NextAction.Type = "verify_result"
	response.NextAction.URL = statusURL
	return response
}

func writeResourceFor(t *testing.T, action string) *writeResource {
	t.Helper()
	resourceUnderTest, ok := resourceFor(action).(*writeResource)
	if !ok {
		t.Fatalf("resource for %s is not a write resource", action)
	}
	return resourceUnderTest
}

func rawWriteObject(terraformType tftypes.Type, stringsByName map[string]string) tftypes.Value {
	object, ok := terraformType.(tftypes.Object)
	if !ok {
		panic(fmt.Sprintf("schema type %T is not an object", terraformType))
	}
	attributes := make(map[string]tftypes.Value, len(object.AttributeTypes))
	for name, attributeType := range object.AttributeTypes {
		if stringValue, exists := stringsByName[name]; exists && attributeType.Is(tftypes.String) {
			attributes[name] = tftypes.NewValue(attributeType, stringValue)
			continue
		}
		attributes[name] = tftypes.NewValue(attributeType, nil)
	}
	return tftypes.NewValue(object, attributes)
}
