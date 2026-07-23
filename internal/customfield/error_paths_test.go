// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package customfield

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type unsupportedCoverageObject struct {
	Channel chan string `tfsdk:"channel"`
}

type privateCoverageObject struct {
	hidden string `tfsdk:"hidden"`
	Name   string `tfsdk:"name"`
}

type nestedUnsupportedCoverageObject struct {
	Nested unsupportedCoverageObject `tfsdk:"nested"`
}

type listUnsupportedCoverageObject struct {
	Items []unsupportedCoverageObject `tfsdk:"items"`
}

type mapUnsupportedCoverageObject struct {
	Items map[string]unsupportedCoverageObject `tfsdk:"items"`
}

func TestNestedObjectRejectsInvalidSchemasAndValues(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	assertPanics(t, "invalid object type", func() {
		NewNestedObjectType[unsupportedCoverageObject](ctx)
	})
	assertPanics(t, "invalid null object", func() {
		NullObject[unsupportedCoverageObject](ctx)
	})
	assertPanics(t, "invalid unknown object", func() {
		UnknownObject[unsupportedCoverageObject](ctx)
	})
	assertPanics(t, "invalid object value", func() {
		NewObject(ctx, &unsupportedCoverageObject{})
	})
	assertPanics(t, "invalid object conversion", func() {
		NestedObjectType[unsupportedCoverageObject]{}.ValueFromObject(
			ctx,
			basetypes.NewObjectValueMust(
				map[string]attr.Type{},
				map[string]attr.Value{},
			),
		)
	})

	objectType := NewNestedObjectType[coverageObject](ctx)
	wrongObject := wrongCoverageObject()
	_, diagnostics := objectType.ValueFromObject(ctx, wrongObject)
	requireDiagnostics(t, diagnostics)

	if _, err := objectType.ValueFromTerraform(
		ctx,
		tftypes.NewValue(tftypes.String, "wrong"),
	); err == nil {
		t.Fatal("object type accepted a string Terraform value")
	}

	value := NestedObject[coverageObject]{ObjectValue: wrongObject}
	if _, diagnostics := value.Value(ctx); !diagnostics.HasError() {
		t.Fatal("object value accepted mismatched attributes")
	}

	attributes, diagnostics := StructToAttributes[privateCoverageObject](ctx)
	requireNoDiagnostics(t, diagnostics)
	if len(attributes) != 1 {
		t.Fatalf("private field was included: %v", attributes)
	}
	for name, model := range map[string]any{
		"nested": nestedUnsupportedCoverageObject{},
		"list":   listUnsupportedCoverageObject{},
		"map":    mapUnsupportedCoverageObject{},
	} {
		t.Run(name, func(t *testing.T) {
			_, nestedDiagnostics := StructFromAttributesGeneric(
				ctx,
				reflect.ValueOf(model),
			)
			requireDiagnostics(t, nestedDiagnostics)
		})
	}
}

func TestNestedObjectCollectionsRejectInvalidSchemasAndValues(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	wrongObject := wrongCoverageObject()

	t.Run("list", func(t *testing.T) {
		assertPanics(t, "invalid list type", func() {
			NewNestedObjectListType[unsupportedCoverageObject](ctx)
		})
		assertPanics(t, "invalid null list", func() {
			NullObjectList[unsupportedCoverageObject](ctx)
		})
		assertPanics(t, "invalid unknown list", func() {
			UnknownObjectList[unsupportedCoverageObject](ctx)
		})
		assertPanics(t, "invalid list values", func() {
			NewObjectList(ctx, []unsupportedCoverageObject{{}})
		})
		assertPanics(t, "invalid list conversion", func() {
			NestedObjectListType[unsupportedCoverageObject]{}.ValueFromList(
				ctx,
				basetypes.NewListValueMust(
					types.StringType,
					[]attr.Value{types.StringValue("wrong")},
				),
			)
		})

		listType := NewNestedObjectListType[coverageObject](ctx)
		wrongList := basetypes.NewListValueMust(
			wrongObject.Type(ctx),
			[]attr.Value{wrongObject},
		)
		_, diagnostics := listType.ValueFromList(ctx, wrongList)
		requireDiagnostics(t, diagnostics)
		_, diagnostics = NewObjectListFromAttributes[coverageObject](
			ctx,
			[]attr.Value{wrongObject},
		)
		requireDiagnostics(t, diagnostics)
		if _, diagnostics := (NestedObjectList[coverageObject]{ListValue: wrongList}).AsStructSliceT(ctx); !diagnostics.HasError() {
			t.Fatal("object list accepted mismatched attributes")
		}
		if _, err := listType.ValueFromTerraform(
			ctx,
			tftypes.NewValue(tftypes.String, "wrong"),
		); err == nil {
			t.Fatal("object list accepted a string Terraform value")
		}
		assertPanics(t, "invalid reflected list", func() {
			NewObjectListFromValueMust[coverageObject](
				ctx,
				reflect.ValueOf([]unsupportedCoverageObject{{}}),
			)
		})
	})

	t.Run("map", func(t *testing.T) {
		assertPanics(t, "invalid null map", func() {
			NullObjectMap[unsupportedCoverageObject](ctx)
		})
		assertPanics(t, "invalid unknown map", func() {
			UnknownObjectMap[unsupportedCoverageObject](ctx)
		})
		assertPanics(t, "invalid map values", func() {
			NewObjectMap(ctx, map[string]unsupportedCoverageObject{"key": {}})
		})
		assertPanics(t, "invalid map conversion", func() {
			NestedObjectMapType[unsupportedCoverageObject]{}.ValueFromMap(
				ctx,
				basetypes.NewMapValueMust(
					types.StringType,
					map[string]attr.Value{
						"key": types.StringValue("wrong"),
					},
				),
			)
		})

		mapType := NewNestedObjectMapType[coverageObject](ctx)
		wrongMap := basetypes.NewMapValueMust(
			wrongObject.Type(ctx),
			map[string]attr.Value{"key": wrongObject},
		)
		_, diagnostics := mapType.ValueFromMap(ctx, wrongMap)
		requireDiagnostics(t, diagnostics)
		_, diagnostics = NewObjectMapFromAttributes[coverageObject](
			ctx,
			map[string]attr.Value{"key": wrongObject},
		)
		requireDiagnostics(t, diagnostics)
		if _, diagnostics := (NestedObjectMap[coverageObject]{MapValue: wrongMap}).AsStructMapT(ctx); !diagnostics.HasError() {
			t.Fatal("object map accepted mismatched attributes")
		}
		if _, err := mapType.ValueFromTerraform(
			ctx,
			tftypes.NewValue(tftypes.String, "wrong"),
		); err == nil {
			t.Fatal("object map accepted a string Terraform value")
		}
	})

	t.Run("set", func(t *testing.T) {
		assertPanics(t, "invalid null set", func() {
			NullObjectSet[unsupportedCoverageObject](ctx)
		})
		assertPanics(t, "invalid unknown set", func() {
			UnknownObjectSet[unsupportedCoverageObject](ctx)
		})
		assertPanics(t, "invalid set values", func() {
			NewObjectSet(ctx, []unsupportedCoverageObject{{}})
		})
		assertPanics(t, "invalid set conversion", func() {
			NestedObjectSetType[unsupportedCoverageObject]{}.ValueFromSet(
				ctx,
				basetypes.NewSetValueMust(
					types.StringType,
					[]attr.Value{types.StringValue("wrong")},
				),
			)
		})

		setType := NewNestedObjectSetType[coverageObject](ctx)
		wrongSet := basetypes.NewSetValueMust(
			wrongObject.Type(ctx),
			[]attr.Value{wrongObject},
		)
		_, diagnostics := setType.ValueFromSet(ctx, wrongSet)
		requireDiagnostics(t, diagnostics)
		_, diagnostics = NewObjectSetFromAttributes[coverageObject](
			ctx,
			[]attr.Value{wrongObject},
		)
		requireDiagnostics(t, diagnostics)
		if _, diagnostics := (NestedObjectSet[coverageObject]{SetValue: wrongSet}).AsStructSliceT(ctx); !diagnostics.HasError() {
			t.Fatal("object set accepted mismatched attributes")
		}
		if _, err := setType.ValueFromTerraform(
			ctx,
			tftypes.NewValue(tftypes.String, "wrong"),
		); err == nil {
			t.Fatal("object set accepted a string Terraform value")
		}
	})
}

func TestPrimitiveCollectionsRejectMismatchedValues(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	terraformString := tftypes.NewValue(tftypes.String, "wrong")

	listType := NewListType[basetypes.StringValue](ctx)
	_, diagnostics := listType.ValueFromList(
		ctx,
		basetypes.NewListValueMust(
			types.Int64Type,
			[]attr.Value{types.Int64Value(1)},
		),
	)
	requireDiagnostics(t, diagnostics)
	if _, err := listType.ValueFromTerraform(ctx, terraformString); err == nil {
		t.Fatal("list accepted a string Terraform value")
	}

	setType := NewSetType[basetypes.StringValue](ctx)
	_, diagnostics = setType.ValueFromSet(
		ctx,
		basetypes.NewSetValueMust(
			types.Int64Type,
			[]attr.Value{types.Int64Value(1)},
		),
	)
	requireDiagnostics(t, diagnostics)
	if _, err := setType.ValueFromTerraform(ctx, terraformString); err == nil {
		t.Fatal("set accepted a string Terraform value")
	}

	mapType := NewMapType[basetypes.StringValue](ctx)
	_, diagnostics = mapType.ValueFromMap(
		ctx,
		basetypes.NewMapValueMust(
			types.Int64Type,
			map[string]attr.Value{"key": types.Int64Value(1)},
		),
	)
	requireDiagnostics(t, diagnostics)
	if _, err := mapType.ValueFromTerraform(ctx, terraformString); err == nil {
		t.Fatal("map accepted a string Terraform value")
	}
}

func TestDynamicValidationAndSemanticEqualityBranches(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	equal, diagnostics := semanticEquals(ctx, nil, nil)
	requireNoDiagnostics(t, diagnostics)
	if !equal {
		t.Fatal("nil values were not semantically equal")
	}

	equal, diagnostics = semanticEquals(
		ctx,
		types.Float32Value(1.5),
		types.Float64Value(1.5),
	)
	requireNoDiagnostics(t, diagnostics)
	if !equal {
		t.Fatal("equal floating-point values were not semantically equal")
	}

	for name, value := range map[string]attr.Value{
		"set": basetypes.NewSetValueMust(
			types.StringType,
			[]attr.Value{types.StringValue("value")},
		),
		"list": basetypes.NewListValueMust(
			types.StringType,
			[]attr.Value{types.StringValue("value")},
		),
		"object": basetypes.NewObjectValueMust(
			map[string]attr.Type{"key": types.StringType},
			map[string]attr.Value{"key": types.StringValue("value")},
		),
	} {
		t.Run(name, func(t *testing.T) {
			nestedDiagnostics := validate(ctx, value)
			if name == "set" {
				requireDiagnostics(t, nestedDiagnostics)
				return
			}
			requireNoDiagnostics(t, nestedDiagnostics)
		})
	}
}

func wrongCoverageObject() basetypes.ObjectValue {
	return basetypes.NewObjectValueMust(
		map[string]attr.Type{
			"name":  types.StringType,
			"count": types.StringType,
		},
		map[string]attr.Value{
			"name":  types.StringValue("name"),
			"count": types.StringValue("wrong"),
		},
	)
}
