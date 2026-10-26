// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package xwrite

import (
	"context"
	"fmt"
	"io"
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
		if _, err := io.WriteString(writer, writeActionJSON(operations[0].action, serverURL(request)+"/x/write-actions/write-action-123", failedWrite)); err != nil {
			t.Errorf("write response: %v", err)
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

type writeActionState struct {
	status         string
	terminal       bool
	charged        bool
	chargedCredits string
	billingStatus  string
	result         string
	nextAction     string
	sendDispatched bool
	success        bool
}

var (
	succeededWrite = writeActionState{
		status: "success", terminal: true, charged: true, chargedCredits: "10", billingStatus: "charged",
		result: `{"id":"result-123","state":"confirmed","type":"state_change"}`, nextAction: "verify_result",
		sendDispatched: true, success: true,
	}
	acceptedWrite = writeActionState{
		status: "accepted", chargedCredits: "10", billingStatus: "pending",
		result: `{}`, nextAction: "poll", sendDispatched: true,
	}
	failedWrite = writeActionState{
		status: "failed", terminal: true, billingStatus: "not_charged",
		result: `{"id":"result-123","state":"confirmed","type":"state_change"}`, nextAction: "verify_result",
	}
)

func writeActionJSON(action string, statusURL string, state writeActionState) string {
	return fmt.Sprintf(`{"object":"x_write_action","id":"write-action-123","writeActionId":"write-action-123",`+
		`"action":%q,"status":%q,"terminal":%t,"retryable":false,"safeToRetry":false,"statusUrl":%q,"pollAfterMs":0,`+
		`"charged":%t,"chargedCredits":%q,"targetId":"target-123",`+
		`"billing":{"status":%q,"charged":%t,"chargedCredits":%q,"plannedCredits":"10"},`+
		`"request":{"hash":"request-hash","payload":{"account":"sanitized"}},`+
		`"account":{"id":"account-123","username":"response-account"},"result":%s,`+
		`"target":{"id":"target-123","type":"tweet"},`+
		`"nextAction":{"type":%q,"afterMs":0,"requiresNewIdempotencyKey":false,"url":%q},`+
		`"requestId":"request-123","sendDispatched":%t,"success":%t}`,
		action, state.status, state.terminal, statusURL,
		state.charged, state.chargedCredits,
		state.billingStatus, state.charged, state.chargedCredits,
		state.result, state.nextAction, statusURL, state.sendDispatched, state.success)
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
