// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"encoding/json"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/tidwall/gjson"
)

type coverageUnmarshaler string

func (value *coverageUnmarshaler) UnmarshalJSON(data []byte) error {
	*value = coverageUnmarshaler(strings.Trim(string(data), `"`))
	return nil
}

type unknownString string

func (unknownString) IsKnown() bool {
	return false
}

func TestUnmarshalInputValidationAndRootBehavior(t *testing.T) {
	t.Parallel()

	var target struct {
		Value string `json:"value"`
	}
	for name, unmarshal := range map[string]func([]byte, any) error{
		"default":  Unmarshal,
		"computed": UnmarshalComputed,
		"root":     UnmarshalRoot,
	} {
		t.Run(name, func(t *testing.T) {
			if err := unmarshal([]byte(`{`), &target); err == nil {
				t.Fatal("malformed JSON was accepted")
			}
			if err := unmarshal([]byte(`{}`), nil); err == nil {
				t.Fatal("nil destination was accepted")
			}
			if err := unmarshal([]byte(`{}`), target); err == nil {
				t.Fatal("non-pointer destination was accepted")
			}
		})
	}

	var custom coverageUnmarshaler
	if err := Unmarshal([]byte(`"decoded"`), &custom); err != nil {
		t.Fatalf("custom unmarshaler: %v", err)
	}
	if custom != "decoded" {
		t.Fatalf("custom unmarshaler value = %q, want decoded", custom)
	}

	var root coverageUnmarshaler
	if err := UnmarshalRoot([]byte(`"root"`), &root); err != nil {
		t.Fatalf("root unmarshaler: %v", err)
	}
	if root != "root" {
		t.Fatalf("root value = %q, want root", root)
	}
}

func TestDecoderPrimitiveFailuresAndSpecialTypes(t *testing.T) {
	t.Parallel()

	builder := &decoderBuilder{dateFormat: "2006-01-02"}
	tests := []struct {
		name  string
		value any
		json  string
	}{
		{name: "string", value: new(string), json: `{}`},
		{name: "bool", value: new(bool), json: `"invalid"`},
		{name: "int", value: new(int), json: `"invalid"`},
		{name: "uint", value: new(uint), json: `-1`},
		{name: "float", value: new(float64), json: `"invalid"`},
	}
	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			value := reflect.ValueOf(testCase.value).Elem()
			decoder := builder.newPrimitiveTypeDecoder(value.Type())
			if err := decoder(gjson.Parse(testCase.json), value, &decoderState{strict: true, exactness: exact}); err == nil {
				t.Fatal("strict decoder accepted an invalid value")
			}
		})
	}

	var unknown unknownString
	unknownValue := reflect.ValueOf(&unknown).Elem()
	if err := builder.newPrimitiveTypeDecoder(unknownValue.Type())(
		gjson.Parse(`"value"`),
		unknownValue,
		&decoderState{strict: true, exactness: exact},
	); err == nil {
		t.Fatal("unknown enum value was accepted")
	}

	var unsupported complex64
	if err := builder.newPrimitiveTypeDecoder(reflect.TypeOf(unsupported))(
		gjson.Parse(`1`),
		reflect.ValueOf(&unsupported).Elem(),
		&decoderState{},
	); err == nil {
		t.Fatal("unsupported primitive type was accepted")
	}

	var number big.Float
	bigFloatDecoder := builder.newBigFloatDecoder(reflect.TypeOf(number))
	if err := bigFloatDecoder(gjson.Parse(`1.25`), reflect.ValueOf(&number).Elem(), &decoderState{}); err != nil {
		t.Fatalf("decode big float: %v", err)
	}
	if number.String() != "1.25" {
		t.Fatalf("big float = %s, want 1.25", number.String())
	}
	if err := bigFloatDecoder(gjson.Parse(`"invalid"`), reflect.ValueOf(&number).Elem(), &decoderState{}); err == nil {
		t.Fatal("invalid big float was accepted")
	}

	var keyed map[int]string
	keyedValue := reflect.ValueOf(&keyed).Elem()
	if err := builder.newMapDecoder(keyedValue.Type())(
		gjson.Parse(`{"one":"value"}`),
		keyedValue,
		&decoderState{},
	); err == nil {
		t.Fatal("non-string JSON map key was accepted")
	}

	var values []string
	valuesValue := reflect.ValueOf(&values).Elem()
	if err := builder.newArrayTypeDecoder(valuesValue.Type())(
		gjson.Parse(`{}`),
		valuesValue,
		&decoderState{},
	); err == nil {
		t.Fatal("object was accepted as an array")
	}
}

func TestDecoderTerraformInferenceAndInternalGuards(t *testing.T) {
	t.Parallel()

	builder := &decoderBuilder{}
	for name, raw := range map[string]string{
		"null":          `null`,
		"true":          `true`,
		"false":         `false`,
		"integer":       `1`,
		"float":         `1.25`,
		"string":        `"value"`,
		"homogeneous":   `[1,2]`,
		"heterogeneous": `[1,"two"]`,
		"object":        `{"key":"value"}`,
	} {
		t.Run(name, func(t *testing.T) {
			value, err := builder.inferTerraformAttrFromValue(gjson.Parse(raw))
			if err != nil {
				t.Fatalf("infer Terraform value: %v", err)
			}
			if value == nil {
				t.Fatal("inferred Terraform value is nil")
			}
		})
	}

	invalid := gjson.Result{Type: gjson.JSON, Raw: `invalid`}
	if _, err := builder.inferTerraformAttrFromValue(invalid); err == nil {
		t.Fatal("invalid JSON result was inferred")
	}
	elementType, values, err := builder.parseArrayOfValues(gjson.Parse(`[1,2]`))
	if err != nil {
		t.Fatalf("parse array values: %v", err)
	}
	if elementType == nil || len(values) != 2 {
		t.Fatalf("array inference returned type %v and %d values", elementType, len(values))
	}

	state := &decoderState{exactness: exact}
	if guardStrict(state, false) {
		t.Fatal("false strict guard failed")
	}
	if guardStrict(state, true) {
		t.Fatal("loose strict guard rejected a value")
	}
	if state.exactness != loose {
		t.Fatalf("exactness = %v, want loose", state.exactness)
	}
	state.strict = true
	if !guardStrict(state, true) {
		t.Fatal("strict guard accepted a value")
	}
	if canParseAsNumber("not-a-number") {
		t.Fatal("invalid number was accepted")
	}

	type private struct {
		value string
	}
	holder := private{}
	setUnexportedField(reflect.ValueOf(&holder).Elem().FieldByName("value"), "set")
	if holder.value != "set" {
		t.Fatalf("private value = %q, want set", holder.value)
	}
}

func TestUnionRegistrationAndDiagnosticFormatting(t *testing.T) {
	t.Parallel()

	unionType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	RegisterUnion(
		unionType,
		"type",
		UnionVariant{TypeFilter: gjson.String, DiscriminatorValue: "value", Type: reflect.TypeOf("")},
	)
	entry, ok := unionRegistry[unionType]
	if !ok || entry.discriminatorKey != "type" || len(entry.variants) != 1 {
		t.Fatal("union registration was not preserved")
	}

	if err := errorFromDiagnostics(nil); err != nil {
		t.Fatalf("nil diagnostics returned an error: %v", err)
	}
	diagnostics := diag.Diagnostics{
		diag.NewErrorDiagnostic("summary", "detail"),
	}
	err := errorFromDiagnostics(diagnostics)
	if err == nil || !strings.Contains(err.Error(), "summary detail") {
		t.Fatalf("diagnostic error = %v", err)
	}
}
