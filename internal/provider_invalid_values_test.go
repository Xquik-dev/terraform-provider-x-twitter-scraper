// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestRegisteredSurfaceRejectsInvalidTerraformValues(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	providerUnderTest := NewProvider("test")().(*XTwitterScraperProvider)
	invalidRaw := tftypes.NewValue(tftypes.String, "invalid")

	for _, factory := range providerUnderTest.Resources(ctx) {
		resourceUnderTest := factory()
		t.Run(fmt.Sprintf("resource/%T", resourceUnderTest), func(t *testing.T) {
			schemaResponse := resourceSchema(t, resourceUnderTest)
			invalidPlan := tfsdk.Plan{Raw: invalidRaw, Schema: schemaResponse.Schema}
			invalidState := tfsdk.State{Raw: invalidRaw, Schema: schemaResponse.Schema}

			var createResponse resource.CreateResponse
			createResponse.State.Schema = schemaResponse.Schema
			resourceUnderTest.Create(
				ctx,
				resource.CreateRequest{Plan: invalidPlan},
				&createResponse,
			)
			if !createResponse.Diagnostics.HasError() {
				t.Fatal("Create accepted an invalid Terraform plan")
			}

			var readResponse resource.ReadResponse
			readResponse.State.Schema = schemaResponse.Schema
			resourceUnderTest.Read(
				ctx,
				resource.ReadRequest{State: invalidState},
				&readResponse,
			)
			if resourceSupportsRead(resourceUnderTest) &&
				!readResponse.Diagnostics.HasError() {
				t.Fatal("Read accepted invalid Terraform state")
			}

			var updateResponse resource.UpdateResponse
			updateResponse.State.Schema = schemaResponse.Schema
			resourceUnderTest.Update(
				ctx,
				resource.UpdateRequest{
					Plan:  invalidPlan,
					State: invalidState,
				},
				&updateResponse,
			)
			if resourceSupportsUpdate(resourceUnderTest) &&
				!updateResponse.Diagnostics.HasError() {
				t.Fatal("Update accepted invalid Terraform values")
			}
			if resourceSupportsUpdate(resourceUnderTest) {
				validPlan := tfsdk.Plan{
					Raw: rawObject(
						schemaResponse.Schema.Type(),
						lifecycleStrings(),
					),
					Schema: schemaResponse.Schema,
				}
				var invalidStateResponse resource.UpdateResponse
				invalidStateResponse.State.Schema = schemaResponse.Schema
				resourceUnderTest.Update(
					ctx,
					resource.UpdateRequest{
						Plan:  validPlan,
						State: invalidState,
					},
					&invalidStateResponse,
				)
				if !invalidStateResponse.Diagnostics.HasError() {
					t.Fatal("Update accepted invalid Terraform state")
				}
			}

			var deleteResponse resource.DeleteResponse
			deleteResponse.State.Schema = schemaResponse.Schema
			resourceUnderTest.Delete(
				ctx,
				resource.DeleteRequest{State: invalidState},
				&deleteResponse,
			)
			if resourceSupportsDelete(resourceUnderTest) &&
				!deleteResponse.Diagnostics.HasError() {
				t.Fatal("Delete accepted invalid Terraform state")
			}
		})
	}

	for _, factory := range providerUnderTest.DataSources(ctx) {
		dataSourceUnderTest := factory()
		t.Run(fmt.Sprintf("data_source/%T", dataSourceUnderTest), func(t *testing.T) {
			var schemaResponse datasource.SchemaResponse
			dataSourceUnderTest.Schema(
				ctx,
				datasource.SchemaRequest{},
				&schemaResponse,
			)
			var readResponse datasource.ReadResponse
			readResponse.State.Schema = schemaResponse.Schema
			dataSourceUnderTest.Read(
				ctx,
				datasource.ReadRequest{
					Config: tfsdk.Config{
						Raw:    invalidRaw,
						Schema: schemaResponse.Schema,
					},
				},
				&readResponse,
			)
			if !readResponse.Diagnostics.HasError() {
				t.Fatal("Read accepted invalid Terraform configuration")
			}
		})
	}
}

func resourceSupportsUpdate(resourceUnderTest resource.Resource) bool {
	typeName := fmt.Sprintf("%T", resourceUnderTest)
	return strings.Contains(typeName, "monitor.") ||
		strings.Contains(typeName, "monitor_keyword.") ||
		strings.Contains(typeName, "style.") ||
		strings.Contains(typeName, "support_ticket.") ||
		strings.Contains(typeName, "webhook.")
}

func resourceSupportsDelete(resourceUnderTest resource.Resource) bool {
	typeName := fmt.Sprintf("%T", resourceUnderTest)
	return strings.Contains(typeName, "draft.") ||
		strings.Contains(typeName, "monitor_keyword.") ||
		strings.Contains(typeName, "style.")
}
