// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestProviderContract(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	providerFactory := NewProvider("test")
	providerUnderTest := providerFactory()

	var metadataResponse provider.MetadataResponse
	providerUnderTest.Metadata(ctx, provider.MetadataRequest{}, &metadataResponse)
	assertEqual(t, "provider type", metadataResponse.TypeName, "x-twitter-scraper")
	assertEqual(t, "provider version", metadataResponse.Version, "test")

	var schemaResponse provider.SchemaResponse
	providerUnderTest.Schema(ctx, provider.SchemaRequest{}, &schemaResponse)
	if len(schemaResponse.Schema.Attributes) != 3 {
		t.Fatalf("provider attribute count = %d, want 3", len(schemaResponse.Schema.Attributes))
	}

	config := tfsdk.Config{
		Raw: rawObject(schemaResponse.Schema.Type(), map[string]string{
			"api_key":      "test-api-key",
			"base_url":     "https://example.invalid/",
			"bearer_token": "test-bearer-token",
		}),
		Schema: schemaResponse.Schema,
	}
	var configureResponse provider.ConfigureResponse
	providerUnderTest.Configure(ctx, provider.ConfigureRequest{Config: config}, &configureResponse)
	if configureResponse.Diagnostics.HasError() {
		t.Fatalf("configure provider: %v", configureResponse.Diagnostics)
	}
	if configureResponse.ResourceData == nil || configureResponse.DataSourceData == nil {
		t.Fatal("provider did not configure resource and data source clients")
	}

	validatorProvider, ok := providerUnderTest.(provider.ProviderWithConfigValidators)
	if !ok {
		t.Fatal("provider does not expose configuration validators")
	}
	if validators := validatorProvider.ConfigValidators(ctx); len(validators) != 0 {
		t.Fatalf("provider validator count = %d, want 0", len(validators))
	}
}

func TestProviderUsesEnvironmentFallbacks(t *testing.T) {
	t.Setenv("X_TWITTER_SCRAPER_BASE_URL", "https://example.invalid/")
	t.Setenv("X_TWITTER_SCRAPER_API_KEY", "test-api-key")
	t.Setenv("X_TWITTER_SCRAPER_BEARER_TOKEN", "test-bearer-token")

	ctx := context.Background()
	providerUnderTest := NewProvider("test")()
	var schemaResponse provider.SchemaResponse
	providerUnderTest.Schema(ctx, provider.SchemaRequest{}, &schemaResponse)
	config := tfsdk.Config{
		Raw:    rawObject(schemaResponse.Schema.Type(), nil),
		Schema: schemaResponse.Schema,
	}
	var response provider.ConfigureResponse
	providerUnderTest.Configure(ctx, provider.ConfigureRequest{Config: config}, &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("configure provider from environment: %v", response.Diagnostics)
	}
	if response.ResourceData == nil || response.DataSourceData == nil {
		t.Fatal("environment configuration did not create a client")
	}
}

func TestRegisteredResourcesAndDataSources(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{}`))
	}))
	t.Cleanup(server.Close)

	client := xtwitterscraper.NewClient(
		option.WithAPIKey("test-api-key"),
		option.WithBaseURL(server.URL+"/"),
		option.WithMaxRetries(0),
	)
	providerUnderTest := NewProvider("test")().(*XTwitterScraperProvider)

	for _, factory := range providerUnderTest.Resources(context.Background()) {
		resourceUnderTest := factory()
		t.Run(fmt.Sprintf("resource/%T", resourceUnderTest), func(t *testing.T) {
			exerciseResource(t, resourceUnderTest, &client)
		})
	}

	for _, factory := range providerUnderTest.DataSources(context.Background()) {
		dataSourceUnderTest := factory()
		t.Run(fmt.Sprintf("data_source/%T", dataSourceUnderTest), func(t *testing.T) {
			exerciseDataSource(t, dataSourceUnderTest, &client)
		})
	}
}

func TestRegisteredSurfaceRejectsInvalidAPIResponses(t *testing.T) {
	t.Parallel()

	providerUnderTest := NewProvider("test")().(*XTwitterScraperProvider)
	scenarios := []struct {
		name       string
		statusCode int
		body       string
	}{
		{name: "server error", statusCode: http.StatusInternalServerError, body: `{"error":"test"}`},
		{name: "malformed response", statusCode: http.StatusOK, body: `{`},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.name, func(t *testing.T) {
			client, closeServer := responseClient(scenario.statusCode, scenario.body)
			t.Cleanup(closeServer)

			for _, factory := range providerUnderTest.Resources(context.Background()) {
				resourceUnderTest := factory()
				if strings.Contains(fmt.Sprintf("%T", resourceUnderTest), "x_write.") {
					continue
				}
				t.Run(fmt.Sprintf("resource/%T", resourceUnderTest), func(t *testing.T) {
					exerciseRejectedResourceResponse(t, resourceUnderTest, &client)
				})
			}
			for _, factory := range providerUnderTest.DataSources(context.Background()) {
				dataSourceUnderTest := factory()
				t.Run(fmt.Sprintf("data_source/%T", dataSourceUnderTest), func(t *testing.T) {
					exerciseRejectedDataSourceResponse(t, dataSourceUnderTest, &client)
				})
			}
		})
	}
}

func TestRegisteredResourcesHandleNotFoundResponses(t *testing.T) {
	t.Parallel()

	client, closeServer := responseClient(http.StatusNotFound, `{"error":"not found"}`)
	t.Cleanup(closeServer)
	providerUnderTest := NewProvider("test")().(*XTwitterScraperProvider)

	for _, factory := range providerUnderTest.Resources(context.Background()) {
		resourceUnderTest := factory()
		if strings.Contains(fmt.Sprintf("%T", resourceUnderTest), "x_write.") {
			continue
		}
		t.Run(fmt.Sprintf("%T", resourceUnderTest), func(t *testing.T) {
			ctx := context.Background()
			schemaResponse := resourceSchema(t, resourceUnderTest)
			configureResource(t, resourceUnderTest, &client)
			raw := rawObject(schemaResponse.Schema.Type(), lifecycleStrings())
			state := tfsdk.State{Raw: raw, Schema: schemaResponse.Schema}

			if resourceSupportsRead(resourceUnderTest) {
				var readResponse resource.ReadResponse
				readResponse.State.Schema = schemaResponse.Schema
				resourceUnderTest.Read(ctx, resource.ReadRequest{State: state}, &readResponse)
				if readResponse.Diagnostics.HasError() {
					t.Fatalf("not-found read returned an error: %v", readResponse.Diagnostics)
				}
				if len(readResponse.Diagnostics) == 0 {
					t.Fatal("not-found read returned no warning")
				}
			}

			var deleteResponse resource.DeleteResponse
			deleteResponse.State.Schema = schemaResponse.Schema
			resourceUnderTest.Delete(ctx, resource.DeleteRequest{State: state}, &deleteResponse)
		})
	}
}

func exerciseResource(t *testing.T, resourceUnderTest resource.Resource, client *xtwitterscraper.Client) {
	t.Helper()

	ctx := context.Background()
	var metadataResponse resource.MetadataResponse
	resourceUnderTest.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "x-twitter-scraper"}, &metadataResponse)
	if metadataResponse.TypeName == "" {
		t.Fatal("resource metadata type is empty")
	}

	var schemaResponse resource.SchemaResponse
	resourceUnderTest.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
	if len(schemaResponse.Schema.Attributes) == 0 {
		t.Fatal("resource schema has no attributes")
	}

	configurable, ok := resourceUnderTest.(resource.ResourceWithConfigure)
	if !ok {
		t.Fatal("resource does not implement configuration")
	}
	var nilConfigureResponse resource.ConfigureResponse
	configurable.Configure(ctx, resource.ConfigureRequest{}, &nilConfigureResponse)
	if nilConfigureResponse.Diagnostics.HasError() {
		t.Fatalf("nil resource configuration: %v", nilConfigureResponse.Diagnostics)
	}
	var invalidConfigureResponse resource.ConfigureResponse
	configurable.Configure(ctx, resource.ConfigureRequest{ProviderData: "invalid"}, &invalidConfigureResponse)
	if !invalidConfigureResponse.Diagnostics.HasError() {
		t.Fatal("resource accepted an invalid provider client")
	}
	var configureResponse resource.ConfigureResponse
	configurable.Configure(ctx, resource.ConfigureRequest{ProviderData: client}, &configureResponse)
	if configureResponse.Diagnostics.HasError() {
		t.Fatalf("configure resource: %v", configureResponse.Diagnostics)
	}
	if validatorResource, ok := resourceUnderTest.(resource.ResourceWithConfigValidators); ok {
		if validators := validatorResource.ConfigValidators(ctx); len(validators) != 0 {
			t.Fatalf("resource validator count = %d, want 0", len(validators))
		}
	}

	raw := rawObject(schemaResponse.Schema.Type(), lifecycleStrings())
	plan := tfsdk.Plan{Raw: raw, Schema: schemaResponse.Schema}
	state := tfsdk.State{Raw: raw, Schema: schemaResponse.Schema}

	var createResponse resource.CreateResponse
	createResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Create(ctx, resource.CreateRequest{Plan: plan}, &createResponse)

	var readResponse resource.ReadResponse
	readResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Read(ctx, resource.ReadRequest{State: state}, &readResponse)

	var updateResponse resource.UpdateResponse
	updateResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Update(ctx, resource.UpdateRequest{Plan: plan, State: state}, &updateResponse)

	var deleteResponse resource.DeleteResponse
	deleteResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Delete(ctx, resource.DeleteRequest{State: state}, &deleteResponse)

	if importer, ok := resourceUnderTest.(resource.ResourceWithImportState); ok {
		var invalidResponse resource.ImportStateResponse
		invalidResponse.State.Schema = schemaResponse.Schema
		importer.ImportState(ctx, resource.ImportStateRequest{ID: "too/many"}, &invalidResponse)
		if !invalidResponse.Diagnostics.HasError() {
			t.Fatal("resource accepted a malformed import identifier")
		}

		var importResponse resource.ImportStateResponse
		importResponse.State.Schema = schemaResponse.Schema
		importer.ImportState(ctx, resource.ImportStateRequest{ID: "test"}, &importResponse)
	}

	if modifier, ok := resourceUnderTest.(resource.ResourceWithModifyPlan); ok {
		modifier.ModifyPlan(ctx, resource.ModifyPlanRequest{}, &resource.ModifyPlanResponse{})
	}
	if upgrader, ok := resourceUnderTest.(resource.ResourceWithUpgradeState); ok {
		if upgrades := upgrader.UpgradeState(ctx); len(upgrades) != 0 {
			t.Fatalf("resource state upgrade count = %d, want 0", len(upgrades))
		}
	}
}

func exerciseRejectedResourceResponse(t *testing.T, resourceUnderTest resource.Resource, client *xtwitterscraper.Client) {
	t.Helper()

	ctx := context.Background()
	schemaResponse := resourceSchema(t, resourceUnderTest)
	configureResource(t, resourceUnderTest, client)
	raw := rawObject(schemaResponse.Schema.Type(), lifecycleStrings())
	plan := tfsdk.Plan{Raw: raw, Schema: schemaResponse.Schema}
	state := tfsdk.State{Raw: raw, Schema: schemaResponse.Schema}

	var createResponse resource.CreateResponse
	createResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Create(ctx, resource.CreateRequest{Plan: plan}, &createResponse)
	if !createResponse.Diagnostics.HasError() {
		t.Fatal("resource accepted an invalid create response")
	}

	var readResponse resource.ReadResponse
	readResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Read(ctx, resource.ReadRequest{State: state}, &readResponse)
	if resourceSupportsRead(resourceUnderTest) && !readResponse.Diagnostics.HasError() {
		t.Fatal("resource accepted an invalid read response")
	}

	var updateResponse resource.UpdateResponse
	updateResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Update(ctx, resource.UpdateRequest{Plan: plan, State: state}, &updateResponse)

	var deleteResponse resource.DeleteResponse
	deleteResponse.State.Schema = schemaResponse.Schema
	resourceUnderTest.Delete(ctx, resource.DeleteRequest{State: state}, &deleteResponse)

	if importer, ok := resourceUnderTest.(resource.ResourceWithImportState); ok {
		var importResponse resource.ImportStateResponse
		importResponse.State.Schema = schemaResponse.Schema
		importer.ImportState(ctx, resource.ImportStateRequest{ID: "test"}, &importResponse)
		if !importResponse.Diagnostics.HasError() {
			t.Fatal("resource accepted an invalid import response")
		}
	}
}

func exerciseRejectedDataSourceResponse(t *testing.T, dataSourceUnderTest datasource.DataSource, client *xtwitterscraper.Client) {
	t.Helper()

	ctx := context.Background()
	var schemaResponse datasource.SchemaResponse
	dataSourceUnderTest.Schema(ctx, datasource.SchemaRequest{}, &schemaResponse)
	configurable := dataSourceUnderTest.(datasource.DataSourceWithConfigure)
	var configureResponse datasource.ConfigureResponse
	configurable.Configure(ctx, datasource.ConfigureRequest{ProviderData: client}, &configureResponse)
	if configureResponse.Diagnostics.HasError() {
		t.Fatalf("configure data source: %v", configureResponse.Diagnostics)
	}
	if validatorDataSource, ok := dataSourceUnderTest.(datasource.DataSourceWithConfigValidators); ok {
		if validators := validatorDataSource.ConfigValidators(ctx); len(validators) != 0 {
			t.Fatalf("data source validator count = %d, want 0", len(validators))
		}
	}

	raw := rawObject(schemaResponse.Schema.Type(), lifecycleStrings())
	var readResponse datasource.ReadResponse
	readResponse.State.Schema = schemaResponse.Schema
	dataSourceUnderTest.Read(ctx, datasource.ReadRequest{
		Config: tfsdk.Config{Raw: raw, Schema: schemaResponse.Schema},
	}, &readResponse)
	if !readResponse.Diagnostics.HasError() {
		t.Fatal("data source accepted an invalid API response")
	}
}

func resourceSchema(t *testing.T, resourceUnderTest resource.Resource) resource.SchemaResponse {
	t.Helper()
	var schemaResponse resource.SchemaResponse
	resourceUnderTest.Schema(context.Background(), resource.SchemaRequest{}, &schemaResponse)
	if len(schemaResponse.Schema.Attributes) == 0 {
		t.Fatal("resource schema has no attributes")
	}
	return schemaResponse
}

func configureResource(t *testing.T, resourceUnderTest resource.Resource, client *xtwitterscraper.Client) {
	t.Helper()
	configurable := resourceUnderTest.(resource.ResourceWithConfigure)
	var configureResponse resource.ConfigureResponse
	configurable.Configure(
		context.Background(),
		resource.ConfigureRequest{ProviderData: client},
		&configureResponse,
	)
	if configureResponse.Diagnostics.HasError() {
		t.Fatalf("configure resource: %v", configureResponse.Diagnostics)
	}
}

func responseClient(statusCode int, body string) (xtwitterscraper.Client, func()) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(statusCode)
		_, _ = writer.Write([]byte(body))
	}))
	client := xtwitterscraper.NewClient(
		option.WithAPIKey("test-api-key"),
		option.WithBaseURL(server.URL+"/"),
		option.WithMaxRetries(0),
	)
	return client, server.Close
}

func resourceSupportsRead(resourceUnderTest resource.Resource) bool {
	typeName := fmt.Sprintf("%T", resourceUnderTest)
	return !strings.Contains(typeName, "compose.") &&
		!strings.Contains(typeName, "guest_wallet.") &&
		!strings.Contains(typeName, "webhook.")
}

func exerciseDataSource(t *testing.T, dataSourceUnderTest datasource.DataSource, client *xtwitterscraper.Client) {
	t.Helper()

	ctx := context.Background()
	var metadataResponse datasource.MetadataResponse
	dataSourceUnderTest.Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: "x-twitter-scraper"}, &metadataResponse)
	if metadataResponse.TypeName == "" {
		t.Fatal("data source metadata type is empty")
	}

	var schemaResponse datasource.SchemaResponse
	dataSourceUnderTest.Schema(ctx, datasource.SchemaRequest{}, &schemaResponse)
	if len(schemaResponse.Schema.Attributes) == 0 {
		t.Fatal("data source schema has no attributes")
	}

	configurable, ok := dataSourceUnderTest.(datasource.DataSourceWithConfigure)
	if !ok {
		t.Fatal("data source does not implement configuration")
	}
	var nilConfigureResponse datasource.ConfigureResponse
	configurable.Configure(ctx, datasource.ConfigureRequest{}, &nilConfigureResponse)
	if nilConfigureResponse.Diagnostics.HasError() {
		t.Fatalf("nil data source configuration: %v", nilConfigureResponse.Diagnostics)
	}
	var invalidConfigureResponse datasource.ConfigureResponse
	configurable.Configure(ctx, datasource.ConfigureRequest{ProviderData: "invalid"}, &invalidConfigureResponse)
	if !invalidConfigureResponse.Diagnostics.HasError() {
		t.Fatal("data source accepted an invalid provider client")
	}
	var configureResponse datasource.ConfigureResponse
	configurable.Configure(ctx, datasource.ConfigureRequest{ProviderData: client}, &configureResponse)
	if configureResponse.Diagnostics.HasError() {
		t.Fatalf("configure data source: %v", configureResponse.Diagnostics)
	}

	raw := rawObject(schemaResponse.Schema.Type(), lifecycleStrings())
	var readResponse datasource.ReadResponse
	readResponse.State.Schema = schemaResponse.Schema
	dataSourceUnderTest.Read(ctx, datasource.ReadRequest{
		Config: tfsdk.Config{Raw: raw, Schema: schemaResponse.Schema},
	}, &readResponse)
}

func lifecycleStrings() map[string]string {
	return map[string]string{
		"account":         "@test",
		"id":              "test",
		"idempotency_key": "test-key",
		"keyword":         "test",
		"monitor_id":      "test",
		"payload_json":    `{"text":"hello"}`,
		"query":           "test",
		"status_url":      "x/write-actions/test",
		"target_id":       "test",
		"write_action_id": "test",
	}
}

func rawObject(objectType attr.Type, stringsByName map[string]string) tftypes.Value {
	terraformType := objectType.TerraformType(context.Background())
	object, ok := terraformType.(tftypes.Object)
	if !ok {
		panic(fmt.Sprintf("schema type %T is not an object", terraformType))
	}

	attributes := make(map[string]tftypes.Value, len(object.AttributeTypes))
	for name, attributeType := range object.AttributeTypes {
		value := any(nil)
		if stringValue, ok := stringsByName[name]; ok && attributeType.Is(tftypes.String) {
			value = stringValue
		}
		attributes[name] = tftypes.NewValue(attributeType, value)
	}
	return tftypes.NewValue(object, attributes)
}

func assertEqual[T comparable](t *testing.T, label string, actual T, expected T) {
	t.Helper()
	if actual != expected {
		t.Fatalf("%s = %v, want %v", label, actual, expected)
	}
}
