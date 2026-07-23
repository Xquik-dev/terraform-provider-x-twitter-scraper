// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package customvalidator

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TestAllowedSubtypesPreservesAllowedTypes(t *testing.T) {
	t.Parallel()

	validator, ok := AllowedSubtypes(basetypes.StringType{}).(subtypesValidator)
	if !ok {
		t.Fatal("AllowedSubtypes did not return subtypesValidator")
	}
	if got := validator.Description(context.Background()); got != "The following types are allowed: basetypes.StringType." {
		t.Fatalf("unexpected description: %q", got)
	}
}

func TestValidateDynamicAcceptsUnresolvedValues(t *testing.T) {
	t.Parallel()

	subtypes := AllowedSubtypes(basetypes.StringType{})
	for name, value := range map[string]basetypes.DynamicValue{
		"null":               basetypes.NewDynamicNull(),
		"unknown":            basetypes.NewDynamicUnknown(),
		"underlying null":    basetypes.NewDynamicValue(basetypes.NewStringNull()),
		"underlying unknown": basetypes.NewDynamicValue(basetypes.NewStringUnknown()),
	} {
		t.Run(name, func(t *testing.T) {
			response := validator.DynamicResponse{}
			subtypes.ValidateDynamic(
				context.Background(),
				validator.DynamicRequest{ConfigValue: value},
				&response,
			)
			if response.Diagnostics.HasError() {
				t.Fatalf("unresolved value produced diagnostics: %v", response.Diagnostics)
			}
		})
	}
}

func TestCompatibleAcceptsNarrowNumericTypes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	if !compatible(ctx, basetypes.Int32Type{}, basetypes.NewInt64Value(1)) {
		t.Fatal("Int32 type rejected an integral value")
	}
	if !compatible(ctx, basetypes.Float32Type{}, basetypes.NewFloat64Value(1.5)) {
		t.Fatal("Float32 type rejected a floating-point value")
	}
}
