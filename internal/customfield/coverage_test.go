// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package customfield

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type coverageObject struct {
	Name  types.String `tfsdk:"name"`
	Count types.Int64  `tfsdk:"count"`
}

func TestPrimitiveCollectionContracts(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	stringValue := types.StringValue("value")

	t.Run("list", func(t *testing.T) {
		listType := NewListType[basetypes.StringValue](ctx)
		assertTrue(t, "list type equality", listType.Equal(NewListType[basetypes.StringValue](ctx)))
		assertFalse(t, "list type mismatch", listType.Equal(types.StringType))
		assertNotEmpty(t, "list type string", listType.String())
		assertNotNil(t, "list value type", listType.ValueType(ctx))
		_, diags := listType.NullValue(ctx)
		requireNoDiagnostics(t, diags)

		nullValue, diags := listType.ValueFromList(ctx, basetypes.NewListNull(types.StringType))
		requireNoDiagnostics(t, diags)
		assertTrue(t, "null list", nullValue.IsNull())
		unknownValue, diags := listType.ValueFromList(ctx, basetypes.NewListUnknown(types.StringType))
		requireNoDiagnostics(t, diags)
		assertTrue(t, "unknown list", unknownValue.IsUnknown())

		known := NewListMust[basetypes.StringValue](ctx, []attr.Value{stringValue})
		converted, diags := listType.ValueFromList(ctx, known.ListValue)
		requireNoDiagnostics(t, diags)
		assertFalse(t, "known list null", converted.IsNull())
		terraformValue, err := known.ToTerraformValue(ctx)
		requireNoError(t, err)
		_, err = listType.ValueFromTerraform(ctx, terraformValue)
		requireNoError(t, err)

		var zero List[basetypes.StringValue]
		_, err = zero.ToTerraformValue(ctx)
		requireNoError(t, err)
		assertTrue(t, "list null helper", known.NullValue(ctx).IsNull())
		assertTrue(t, "list unknown helper", known.UnknownValue(ctx).IsUnknown())
		assertFalse(t, "list known helper", known.KnownValue(ctx, []basetypes.StringValue{stringValue}).IsNull())
		assertFalse(t, "known list state", known.IsNullOrUnknown())
		assertTrue(t, "unknown list state", UnknownList[basetypes.StringValue](ctx).IsNullOrUnknown())
		assertTrue(t, "list equality", known.Equal(known))
		assertFalse(t, "list value mismatch", known.Equal(types.StringValue("value")))
		assertNotNil(t, "list value type", known.Type(ctx))
		values, diags := known.Value(ctx)
		requireNoDiagnostics(t, diags)
		assertEqual(t, "list length", len(values), 1)
		_, diags = known.ValueAttr(ctx)
		requireNoDiagnostics(t, diags)

		_, diags = NewList[basetypes.StringValue](ctx, "invalid")
		requireDiagnostics(t, diags)
		assertPanics(t, "invalid list must", func() {
			NewListMust[basetypes.StringValue](ctx, []attr.Value{types.Int64Value(1)})
		})
	})

	t.Run("set", func(t *testing.T) {
		setType := NewSetType[basetypes.StringValue](ctx)
		assertTrue(t, "set type equality", setType.Equal(NewSetType[basetypes.StringValue](ctx)))
		assertFalse(t, "set type mismatch", setType.Equal(types.StringType))
		assertNotEmpty(t, "set type string", setType.String())
		assertNotNil(t, "set value type", setType.ValueType(ctx))
		_, diags := setType.NullValue(ctx)
		requireNoDiagnostics(t, diags)

		nullValue, diags := setType.ValueFromSet(ctx, basetypes.NewSetNull(types.StringType))
		requireNoDiagnostics(t, diags)
		assertTrue(t, "null set", nullValue.IsNull())
		unknownValue, diags := setType.ValueFromSet(ctx, basetypes.NewSetUnknown(types.StringType))
		requireNoDiagnostics(t, diags)
		assertTrue(t, "unknown set", unknownValue.IsUnknown())

		known := NewSetMust[basetypes.StringValue](ctx, []attr.Value{stringValue})
		converted, diags := setType.ValueFromSet(ctx, known.SetValue)
		requireNoDiagnostics(t, diags)
		assertFalse(t, "known set null", converted.IsNull())
		terraformValue, err := known.ToTerraformValue(ctx)
		requireNoError(t, err)
		_, err = setType.ValueFromTerraform(ctx, terraformValue)
		requireNoError(t, err)

		var zero Set[basetypes.StringValue]
		_, err = zero.ToTerraformValue(ctx)
		requireNoError(t, err)
		assertTrue(t, "set null helper", known.NullValue(ctx).IsNull())
		assertTrue(t, "set unknown helper", known.UnknownValue(ctx).IsUnknown())
		assertFalse(t, "set known helper", known.KnownValue(ctx, []basetypes.StringValue{stringValue}).IsNull())
		assertFalse(t, "known set state", known.IsNullOrUnknown())
		assertTrue(t, "unknown set state", UnknownSet[basetypes.StringValue](ctx).IsNullOrUnknown())
		assertTrue(t, "set equality", known.Equal(known))
		assertFalse(t, "set value mismatch", known.Equal(types.StringValue("value")))
		assertNotNil(t, "set value type", known.Type(ctx))
		values, diags := known.Value(ctx)
		requireNoDiagnostics(t, diags)
		assertEqual(t, "set length", len(values), 1)
		_, diags = known.ValueAttr(ctx)
		requireNoDiagnostics(t, diags)

		_, diags = NewSet[basetypes.StringValue](ctx, "invalid")
		requireDiagnostics(t, diags)
		assertPanics(t, "invalid set must", func() {
			NewSetMust[basetypes.StringValue](ctx, []attr.Value{types.Int64Value(1)})
		})
	})

	t.Run("map", func(t *testing.T) {
		mapType := NewMapType[basetypes.StringValue](ctx)
		assertTrue(t, "map type equality", mapType.Equal(NewMapType[basetypes.StringValue](ctx)))
		assertFalse(t, "map type mismatch", mapType.Equal(types.StringType))
		assertNotEmpty(t, "map type string", mapType.String())
		assertNotNil(t, "map value type", mapType.ValueType(ctx))
		_, diags := mapType.NullValue(ctx)
		requireNoDiagnostics(t, diags)

		nullValue, diags := mapType.ValueFromMap(ctx, basetypes.NewMapNull(types.StringType))
		requireNoDiagnostics(t, diags)
		assertTrue(t, "null map", nullValue.IsNull())
		unknownValue, diags := mapType.ValueFromMap(ctx, basetypes.NewMapUnknown(types.StringType))
		requireNoDiagnostics(t, diags)
		assertTrue(t, "unknown map", unknownValue.IsUnknown())

		known := NewMapMust(ctx, map[string]basetypes.StringValue{"key": stringValue})
		converted, diags := mapType.ValueFromMap(ctx, known.MapValue)
		requireNoDiagnostics(t, diags)
		assertFalse(t, "known map null", converted.IsNull())
		terraformValue, err := known.ToTerraformValue(ctx)
		requireNoError(t, err)
		_, err = mapType.ValueFromTerraform(ctx, terraformValue)
		requireNoError(t, err)

		var zero Map[basetypes.StringValue]
		_, err = zero.ToTerraformValue(ctx)
		requireNoError(t, err)
		assertTrue(t, "map null helper", known.NullValue(ctx).IsNull())
		assertTrue(t, "map unknown helper", known.UnknownValue(ctx).IsUnknown())
		assertFalse(t, "map known helper", known.KnownValue(ctx, map[string]basetypes.StringValue{"key": stringValue}).IsNull())
		assertTrue(t, "map equality", known.Equal(known))
		assertFalse(t, "map value mismatch", known.Equal(types.StringValue("value")))
		assertNotNil(t, "map value type", known.Type(ctx))
		values, diags := known.Value(ctx)
		requireNoDiagnostics(t, diags)
		assertEqual(t, "map length", len(values), 1)
		_, diags = known.ValueAttr(ctx)
		requireNoDiagnostics(t, diags)

		_, diags = NewMap[basetypes.StringValue](ctx, "invalid")
		requireDiagnostics(t, diags)
		_, diags = NewMap[basetypes.StringValue](ctx, map[string]attr.Value{"key": types.Int64Value(1)})
		requireDiagnostics(t, diags)
	})
}

func TestNestedObjectCollectionContracts(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	item := coverageObject{Name: types.StringValue("name"), Count: types.Int64Value(1)}
	object := NewObjectMust(ctx, &item)

	objectType := NewNestedObjectType[coverageObject](ctx)
	assertTrue(t, "object type equality", objectType.Equal(NewNestedObjectType[coverageObject](ctx)))
	assertFalse(t, "object type mismatch", objectType.Equal(types.StringType))
	assertNotEmpty(t, "object type string", objectType.String())
	assertNotNil(t, "object value type", objectType.ValueType(ctx))
	_, diags := objectType.NullValue(ctx)
	requireNoDiagnostics(t, diags)
	_, diags = objectType.ValueFromObject(ctx, basetypes.NewObjectNull(objectType.AttrTypes))
	requireNoDiagnostics(t, diags)
	_, diags = objectType.ValueFromObject(ctx, basetypes.NewObjectUnknown(objectType.AttrTypes))
	requireNoDiagnostics(t, diags)
	_, diags = objectType.ValueFromObject(ctx, object.ObjectValue)
	requireNoDiagnostics(t, diags)
	terraformValue, err := object.ToTerraformValue(ctx)
	requireNoError(t, err)
	_, err = objectType.ValueFromTerraform(ctx, terraformValue)
	requireNoError(t, err)

	var zeroObject NestedObject[coverageObject]
	_, diags = zeroObject.ToObjectValue(ctx)
	requireNoDiagnostics(t, diags)
	_, err = zeroObject.ToTerraformValue(ctx)
	requireNoError(t, err)
	assertTrue(t, "object null helper", object.NullValue(ctx).IsNull())
	assertTrue(t, "object unknown helper", object.UnknownValue(ctx).IsUnknown())
	assertFalse(t, "object known helper", object.KnownValue(ctx, &item).IsNull())
	assertTrue(t, "object equality", object.Equal(object))
	assertFalse(t, "object value mismatch", object.Equal(types.StringValue("value")))
	assertNotNil(t, "object value type", object.Type(ctx))
	value, diags := object.Value(ctx)
	requireNoDiagnostics(t, diags)
	assertEqual(t, "object name", value.Name.ValueString(), "name")
	_, diags = object.ValueAny(ctx)
	requireNoDiagnostics(t, diags)
	nullValue, diags := NullObject[coverageObject](ctx).Value(ctx)
	requireNoDiagnostics(t, diags)
	if nullValue != nil {
		t.Fatal("null object returned a value")
	}

	t.Run("list", func(t *testing.T) {
		collectionType := NewNestedObjectListType[coverageObject](ctx)
		exerciseNestedCollectionType(t, collectionType.String(), collectionType.Equal(NewNestedObjectListType[coverageObject](ctx)), collectionType.Equal(types.StringType))
		assertNotNil(t, "object list value type", collectionType.ValueType(ctx))
		_, nestedDiags := collectionType.NullValue(ctx)
		requireNoDiagnostics(t, nestedDiags)
		_, nestedDiags = collectionType.ValueFromList(ctx, basetypes.NewListNull(object.Type(ctx)))
		requireNoDiagnostics(t, nestedDiags)
		_, nestedDiags = collectionType.ValueFromList(ctx, basetypes.NewListUnknown(object.Type(ctx)))
		requireNoDiagnostics(t, nestedDiags)

		known := NewObjectListMust(ctx, []coverageObject{item})
		_, nestedDiags = collectionType.ValueFromList(ctx, known.ListValue)
		requireNoDiagnostics(t, nestedDiags)
		tfValue, nestedErr := known.ToTerraformValue(ctx)
		requireNoError(t, nestedErr)
		_, nestedErr = collectionType.ValueFromTerraform(ctx, tfValue)
		requireNoError(t, nestedErr)

		var zero NestedObjectList[coverageObject]
		_, nestedErr = zero.ToTerraformValue(ctx)
		requireNoError(t, nestedErr)
		assertTrue(t, "object list null helper", known.NullValue(ctx).IsNull())
		assertTrue(t, "object list unknown helper", known.UnknownValue(ctx).IsUnknown())
		assertFalse(t, "object list known helper", known.KnownValue(ctx, []coverageObject{item}).IsNull())
		assertTrue(t, "object list equality", known.Equal(known))
		assertFalse(t, "object list mismatch", known.Equal(types.StringValue("value")))
		assertNotNil(t, "object list type", known.Type(ctx))
		_, nestedDiags = known.AsStructSlice(ctx)
		requireNoDiagnostics(t, nestedDiags)
		values, nestedDiags := known.AsStructSliceT(ctx)
		requireNoDiagnostics(t, nestedDiags)
		assertEqual(t, "object list length", len(values), 1)
		_, nestedDiags = known.Value(ctx)
		requireNoDiagnostics(t, nestedDiags)

		_, nestedDiags = NewObjectListFromAttributes[coverageObject](ctx, known.Elements())
		requireNoDiagnostics(t, nestedDiags)
		reflectValue := reflect.ValueOf([]coverageObject{item})
		_, nestedDiags = NewObjectListFromValue[coverageObject](ctx, reflectValue)
		requireNoDiagnostics(t, nestedDiags)
		_ = NewObjectListFromValueMust[coverageObject](ctx, reflectValue)
	})

	t.Run("map", func(t *testing.T) {
		collectionType := NewNestedObjectMapType[coverageObject](ctx)
		exerciseNestedCollectionType(t, collectionType.String(), collectionType.Equal(NewNestedObjectMapType[coverageObject](ctx)), collectionType.Equal(types.StringType))
		assertNotNil(t, "object map value type", collectionType.ValueType(ctx))
		_, nestedDiags := collectionType.NullValue(ctx)
		requireNoDiagnostics(t, nestedDiags)
		_, nestedDiags = collectionType.ValueFromMap(ctx, basetypes.NewMapNull(object.Type(ctx)))
		requireNoDiagnostics(t, nestedDiags)
		_, nestedDiags = collectionType.ValueFromMap(ctx, basetypes.NewMapUnknown(object.Type(ctx)))
		requireNoDiagnostics(t, nestedDiags)

		known := NewObjectMapMust(ctx, map[string]coverageObject{"key": item})
		_, nestedDiags = collectionType.ValueFromMap(ctx, known.MapValue)
		requireNoDiagnostics(t, nestedDiags)
		tfValue, nestedErr := known.ToTerraformValue(ctx)
		requireNoError(t, nestedErr)
		_, nestedErr = collectionType.ValueFromTerraform(ctx, tfValue)
		requireNoError(t, nestedErr)

		var zero NestedObjectMap[coverageObject]
		_, nestedErr = zero.ToTerraformValue(ctx)
		requireNoError(t, nestedErr)
		assertTrue(t, "object map null helper", known.NullValue(ctx).IsNull())
		assertTrue(t, "object map unknown helper", known.UnknownValue(ctx).IsUnknown())
		assertFalse(t, "object map known helper", known.KnownValue(ctx, map[string]coverageObject{"key": item}).IsNull())
		assertTrue(t, "object map equality", known.Equal(known))
		assertFalse(t, "object map mismatch", known.Equal(types.StringValue("value")))
		assertNotNil(t, "object map type", known.Type(ctx))
		_, nestedDiags = known.AsStructMap(ctx)
		requireNoDiagnostics(t, nestedDiags)
		values, nestedDiags := known.AsStructMapT(ctx)
		requireNoDiagnostics(t, nestedDiags)
		assertEqual(t, "object map length", len(values), 1)
		_, nestedDiags = known.Value(ctx)
		requireNoDiagnostics(t, nestedDiags)
		_, nestedDiags = NewObjectMapFromAttributes[coverageObject](ctx, known.Elements())
		requireNoDiagnostics(t, nestedDiags)
	})

	t.Run("set", func(t *testing.T) {
		collectionType := NewNestedObjectSetType[coverageObject](ctx)
		exerciseNestedCollectionType(t, collectionType.String(), collectionType.Equal(NewNestedObjectSetType[coverageObject](ctx)), collectionType.Equal(types.StringType))
		assertNotNil(t, "object set value type", collectionType.ValueType(ctx))
		_, nestedDiags := collectionType.NullValue(ctx)
		requireNoDiagnostics(t, nestedDiags)
		_, nestedDiags = collectionType.ValueFromSet(ctx, basetypes.NewSetNull(object.Type(ctx)))
		requireNoDiagnostics(t, nestedDiags)
		_, nestedDiags = collectionType.ValueFromSet(ctx, basetypes.NewSetUnknown(object.Type(ctx)))
		requireNoDiagnostics(t, nestedDiags)

		known := NewObjectSetMust(ctx, []coverageObject{item})
		_, nestedDiags = collectionType.ValueFromSet(ctx, known.SetValue)
		requireNoDiagnostics(t, nestedDiags)
		tfValue, nestedErr := known.ToTerraformValue(ctx)
		requireNoError(t, nestedErr)
		_, nestedErr = collectionType.ValueFromTerraform(ctx, tfValue)
		requireNoError(t, nestedErr)

		var zero NestedObjectSet[coverageObject]
		_, nestedErr = zero.ToTerraformValue(ctx)
		requireNoError(t, nestedErr)
		assertTrue(t, "object set null helper", known.NullValue(ctx).IsNull())
		assertTrue(t, "object set unknown helper", known.UnknownValue(ctx).IsUnknown())
		assertFalse(t, "object set known helper", known.KnownValue(ctx, []coverageObject{item}).IsNull())
		assertFalse(t, "known object set state", known.IsNullOrUnknown())
		assertTrue(t, "null object set state", NullObjectSet[coverageObject](ctx).IsNullOrUnknown())
		assertTrue(t, "object set equality", known.Equal(known))
		assertFalse(t, "object set mismatch", known.Equal(types.StringValue("value")))
		assertNotNil(t, "object set type", known.Type(ctx))
		_, nestedDiags = known.AsStructSlice(ctx)
		requireNoDiagnostics(t, nestedDiags)
		values, nestedDiags := known.AsStructSliceT(ctx)
		requireNoDiagnostics(t, nestedDiags)
		assertEqual(t, "object set length", len(values), 1)
		_, nestedDiags = known.Value(ctx)
		requireNoDiagnostics(t, nestedDiags)
		_, nestedDiags = NewObjectSetFromAttributes[coverageObject](ctx, known.Elements())
		requireNoDiagnostics(t, nestedDiags)
	})
}

func TestStructTypeValidation(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	type nested struct {
		Name string `tfsdk:"name"`
	}
	type valid struct {
		Nested *nested            `tfsdk:"nested"`
		List   []*nested          `tfsdk:"list"`
		Map    map[string]*nested `tfsdk:"map"`
		Text   string             `tfsdk:"text"`
		Count  uint               `tfsdk:"count"`
		Ratio  float64            `tfsdk:"ratio"`
		Flag   bool               `tfsdk:"flag"`
		Skip   string             `tfsdk:"-"`
	}
	attributes, diags := StructToAttributes[valid](ctx)
	requireNoDiagnostics(t, diags)
	assertEqual(t, "valid attribute count", len(attributes), 7)

	_, diags = StructToAttributes[int](ctx)
	requireDiagnostics(t, diags)
	_, diags = StructToAttributes[struct{ Missing string }](ctx)
	requireDiagnostics(t, diags)
	_, diags = StructToAttributes[struct {
		Invalid map[int]string `tfsdk:"invalid"`
	}](ctx)
	assertEqual(t, "map-key diagnostic count", len(diags), 1)
	_, diags = StructToAttributes[struct {
		Invalid chan string `tfsdk:"invalid"`
	}](ctx)
	assertEqual(t, "unsupported-type diagnostic count", len(diags), 1)
}

func TestNormalizedDynamicContracts(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	normalizedType := NormalizedDynamicType{}
	assertTrue(t, "dynamic type equality", normalizedType.Equal(NormalizedDynamicType{}))
	assertFalse(t, "dynamic type mismatch", normalizedType.Equal(types.StringType))
	assertNotEmpty(t, "dynamic type string", normalizedType.String())
	assertNotNil(t, "dynamic value type", normalizedType.ValueType(ctx))

	dynamicString := types.DynamicValue(types.StringValue("value"))
	normalizedValue, diags := normalizedType.ValueFromDynamic(ctx, dynamicString)
	requireNoDiagnostics(t, diags)
	converted := normalizedValue.(NormalizedDynamicValue)
	_, diags = converted.ToDynamicValue(ctx)
	requireNoDiagnostics(t, diags)
	assertNotNil(t, "normalized value type", converted.Type(ctx))

	terraformValue, err := converted.ToTerraformValue(ctx)
	requireNoError(t, err)
	_, err = normalizedType.ValueFromTerraform(ctx, terraformValue)
	requireNoError(t, err)
	_, err = RawNormalizedDynamicValue(types.DynamicNull()).ToTerraformValue(ctx)
	requireNoError(t, err)
	_, err = RawNormalizedDynamicValue(types.DynamicUnknown()).ToTerraformValue(ctx)
	requireNoError(t, err)

	for name, value := range map[string]attr.Value{
		"list": basetypes.ListValue{},
		"map":  basetypes.MapValue{},
		"set":  basetypes.SetValue{},
	} {
		t.Run(name, func(t *testing.T) {
			_, nestedErr := RawNormalizedDynamicValueFrom(value).ToTerraformValue(ctx)
			requireNoError(t, nestedErr)
		})
	}

	modifier := NormalizeDynamicPlanModifier()
	assertEqual(t, "modifier description", modifier.Description(ctx), "")
	assertEqual(t, "modifier markdown", modifier.MarkdownDescription(ctx), "")
	var equalResponse planmodifier.DynamicResponse
	modifier.PlanModifyDynamic(ctx, planmodifier.DynamicRequest{
		PlanValue:  dynamicString,
		StateValue: dynamicString,
	}, &equalResponse)
	requireNoDiagnostics(t, equalResponse.Diagnostics)
	assertTrue(t, "equal plan preserved", equalResponse.PlanValue.Equal(dynamicString))

	mapValue := types.MapValueMust(types.StringType, map[string]attr.Value{"key": types.StringValue("value")})
	var invalidResponse planmodifier.DynamicResponse
	modifier.PlanModifyDynamic(ctx, planmodifier.DynamicRequest{
		PlanValue: types.DynamicValue(mapValue),
	}, &invalidResponse)
	requireDiagnostics(t, invalidResponse.Diagnostics)
}

func exerciseNestedCollectionType(t *testing.T, typeName string, equal bool, mismatched bool) {
	t.Helper()
	assertNotEmpty(t, "nested collection type string", typeName)
	assertTrue(t, "nested collection type equality", equal)
	assertFalse(t, "nested collection type mismatch", mismatched)
}

func requireNoDiagnostics(t *testing.T, diags diag.Diagnostics) {
	t.Helper()
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
}

func requireDiagnostics(t *testing.T, diags diag.Diagnostics) {
	t.Helper()
	if !diags.HasError() {
		t.Fatal("expected error diagnostics")
	}
}

func requireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertPanics(t *testing.T, label string, function func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatalf("%s did not panic", label)
		}
	}()
	function()
}

func assertTrue(t *testing.T, label string, value bool) {
	t.Helper()
	if !value {
		t.Fatalf("%s = false, want true", label)
	}
}

func assertFalse(t *testing.T, label string, value bool) {
	t.Helper()
	if value {
		t.Fatalf("%s = true, want false", label)
	}
}

func assertNotEmpty(t *testing.T, label string, value string) {
	t.Helper()
	if value == "" {
		t.Fatalf("%s is empty", label)
	}
}

func assertNotNil(t *testing.T, label string, value any) {
	t.Helper()
	if value == nil {
		t.Fatalf("%s is nil", label)
	}
}

func assertEqual[T comparable](t *testing.T, label string, actual T, expected T) {
	t.Helper()
	if actual != expected {
		t.Fatalf("%s = %v, want %v", label, actual, expected)
	}
}
