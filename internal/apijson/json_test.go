// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
	type_helpers "github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/types"
)

func P[T any](v T) *T { return &v }

type TfsdkStructs struct {
	BoolValue        types.Bool                                        `tfsdk:"tfsdk_bool_value" json:"bool_value"`
	StringValue      types.String                                      `tfsdk:"tfsdk_string_value" json:"string_value"`
	Data             *EmbeddedTfsdkStruct                              `tfsdk:"tfsdk_data" json:"data"`
	DataObject       customfield.NestedObject[EmbeddedTfsdkStruct]     `tfsdk:"tfsdk_data_object" json:"data_object"`
	ListObject       customfield.List[types.String]                    `tfsdk:"tfsdk_list_object" json:"list_object"`
	NestedObjectList customfield.NestedObjectList[EmbeddedTfsdkStruct] `tfsdk:"tfsdk_nested_object_list" json:"nested_object_list"`
	SetObject        customfield.Set[types.String]                     `tfsdk:"tfsdk_set_object" json:"set_object"`
	NestedObjectSet  customfield.NestedObjectSet[EmbeddedTfsdkStruct]  `tfsdk:"tfsdk_nested_object_set" json:"nested_object_set"`
	MapObject        customfield.Map[types.String]                     `tfsdk:"tfsdk_map_object" json:"map_object"`
	NestedObjectMap  customfield.NestedObjectMap[EmbeddedTfsdkStruct]  `tfsdk:"tfsdk_nested_object_map" json:"nested_object_map"`
	FloatValue       types.Float64                                     `tfsdk:"tfsdk_float_value" json:"float_value"`
	OptionalArray    *[]types.String                                   `tfsdk:"tfsdk_optional_array" json:"optional_array"`
}

type EmbeddedTfsdkStruct struct {
	EmbeddedString types.String                                 `tfsdk:"tfsdk_embedded_string" json:"embedded_string" api:"required"`
	EmbeddedInt    types.Int64                                  `tfsdk:"tfsdk_embedded_int" json:"embedded_int" api:"optional"`
	DataObject     customfield.NestedObject[DoubleNestedStruct] `tfsdk:"tfsdk_data_object" json:"data_object" api:"optional"`
}

type DoubleNestedStruct struct {
	NestedInt types.Int64 `tfsdk:"tfsdk_nested_int" json:"nested_int"`
}

type Primitives struct {
	A bool    `json:"a"`
	B int     `json:"b"`
	C uint    `json:"c"`
	D float64 `json:"d"`
	E float32 `json:"e"`
	F []int   `json:"f"`
}

type PrimitivePointers struct {
	A *bool    `json:"a"`
	B *int     `json:"b"`
	C *uint    `json:"c"`
	D *float64 `json:"d"`
	E *float32 `json:"e"`
	F *[]int   `json:"f"`
}

type Slices struct {
	Slice []Primitives `json:"slices"`
}

type DateTime struct {
	Date     time.Time `json:"date" format:"date"`
	DateTime time.Time `json:"date-time" format:"date-time"`
}

type DateTimeCustom struct {
	DateCustom     timetypes.RFC3339 `json:"date" format:"date"`
	DateTimeCustom timetypes.RFC3339 `json:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool                            `json:"a"`
	Extras map[string]jsontypes.Normalized `json:"-" api:"extras"`
}

type TypedAdditionalProperties struct {
	A      bool           `json:"a"`
	Extras map[string]int `json:"-" api:"extras"`
}

type EmbeddedStructs struct {
	AdditionalProperties
	A      *int                            `json:"number2"`
	Extras map[string]jsontypes.Normalized `json:"-" api:"extras"`
}

type Recursive struct {
	Name  string     `json:"name"`
	Child *Recursive `json:"child"`
}

type Inline struct {
	InlineField Primitives `json:"-" api:"inline"`
}

type InlineArray struct {
	InlineField []string `json:"-" api:"inline"`
}

type EncodeStateForUnknownStruct struct {
	NormalField                     types.String `tfsdk:"normal_field" json:"normal_field"`
	ComputedWithForceEncode         types.String `tfsdk:"computed_force_encode" json:"computed_force_encode" api:"computed,force_encode"`
	ComputedWithStateEncode         types.String `tfsdk:"computed_state_encode" json:"computed_state_encode" api:"computed,force_encode,encode_state_for_unknown"`
	ComputedOptionalWithStateEncode types.String `tfsdk:"computed_optional_state_encode" json:"computed_optional_state_encode" api:"computed_optional,encode_state_for_unknown"`
	ComputedRegular                 types.String `tfsdk:"computed_regular" json:"computed_regular" api:"computed"`
	ComputedOptionalRegular         types.String `tfsdk:"computed_optional_regular" json:"computed_optional_regular" api:"computed_optional"`
}

type ResultEnvelope struct {
	Result RecordsModel `json:"result"`
}

type RecordsModel struct {
	A types.String `tfsdk:"tfsdk_a" json:"a"`
	B types.String `tfsdk:"tfsdk_b" json:"b"`
	C types.String `tfsdk:"tfsdk_c" json:"c" api:"computed"`
}

func DropDiagnostic[resType any](res resType, diags diag.Diagnostics) resType {
	for _, d := range diags {
		panic(fmt.Sprintf("%s: %s", d.Summary(), d.Detail()))
	}
	return res
}

type JSONModel struct {
	Arr  jsontypes.Normalized            `tfsdk:"tfsdk_arr" json:"arr"`
	Bol  jsontypes.Normalized            `tfsdk:"tfsdk_bol" json:"bol"`
	Map  jsontypes.Normalized            `tfsdk:"tfsdk_map" json:"map"`
	Nil  jsontypes.Normalized            `tfsdk:"tfsdk_nil" json:"nil"`
	Num  jsontypes.Normalized            `tfsdk:"tfsdk_num" json:"num"`
	Str  jsontypes.Normalized            `tfsdk:"tfsdk_str" json:"str"`
	Arr2 []jsontypes.Normalized          `tfsdk:"tfsdk_arr2" json:"arr2"`
	Map2 map[string]jsontypes.Normalized `tfsdk:"tfsdk_map2" json:"map2"`
}

func time2time(t time.Time) timetypes.RFC3339 {
	return timetypes.NewRFC3339TimePointerValue(&t)
}

var ctx = context.TODO()

var tests = map[string]struct {
	buf string
	val reflect.Value
}{
	"true":               {"true", reflect.ValueOf(true)},
	"false":              {"false", reflect.ValueOf(false)},
	"int":                {"1", reflect.ValueOf(1)},
	"int_bigger":         {"12324", reflect.ValueOf(12324)},
	"int_string_coerce":  {`"65"`, reflect.ValueOf(65)},
	"int_boolean_coerce": {"true", reflect.ValueOf(1)},
	"int64":              {"1", reflect.ValueOf(int64(1))},
	"int64_huge":         {"123456789123456789", reflect.ValueOf(int64(123456789123456789))},
	"uint":               {"1", reflect.ValueOf(uint(1))},
	"uint_bigger":        {"12324", reflect.ValueOf(uint(12324))},
	"uint_coerce":        {`"65"`, reflect.ValueOf(uint(65))},
	"float_1.54":         {"1.54", reflect.ValueOf(float32(1.54))},
	"float_1.89":         {"1.89", reflect.ValueOf(float64(1.89))},
	"string":             {`"str"`, reflect.ValueOf("str")},
	"string_int_coerce":  {`12`, reflect.ValueOf("12")},
	"array_string":       {`["foo","bar"]`, reflect.ValueOf([]string{"foo", "bar"})},
	"array_int":          {`[1,2]`, reflect.ValueOf([]int{1, 2})},
	"array_int_coerce":   {`["1",2]`, reflect.ValueOf([]int{1, 2})},

	"ptr_true":               {"true", reflect.ValueOf(P(true))},
	"ptr_false":              {"false", reflect.ValueOf(P(false))},
	"ptr_int":                {"1", reflect.ValueOf(P(1))},
	"ptr_int_bigger":         {"12324", reflect.ValueOf(P(12324))},
	"ptr_int_string_coerce":  {`"65"`, reflect.ValueOf(P(65))},
	"ptr_int_boolean_coerce": {"true", reflect.ValueOf(P(1))},
	"ptr_int64":              {"1", reflect.ValueOf(P(int64(1)))},
	"ptr_int64_huge":         {"123456789123456789", reflect.ValueOf(P(int64(123456789123456789)))},
	"ptr_uint":               {"1", reflect.ValueOf(P(uint(1)))},
	"ptr_uint_bigger":        {"12324", reflect.ValueOf(P(uint(12324)))},
	"ptr_uint_coerce":        {`"65"`, reflect.ValueOf(P(uint(65)))},
	"ptr_float_1.54":         {"1.54", reflect.ValueOf(P(float32(1.54)))},
	"ptr_float_1.89":         {"1.89", reflect.ValueOf(P(float64(1.89)))},

	"date_time":             {`"2007-03-01T13:00:00Z"`, reflect.ValueOf(time.Date(2007, time.March, 1, 13, 0, 0, 0, time.UTC))},
	"date_time_nano_coerce": {`"2007-03-01T13:03:05.123456789Z"`, reflect.ValueOf(time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC))},

	"date_time_missing_t_coerce":              {`"2007-03-01 13:03:05Z"`, reflect.ValueOf(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC))},
	"date_time_missing_timezone_coerce":       {`"2007-03-01T13:03:05"`, reflect.ValueOf(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC))},
	"date_time_missing_timezone_colon_coerce": {`"2007-03-01T13:03:05+0100"`, reflect.ValueOf(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.FixedZone("", 60*60)))},
	"date_time_nano_missing_t_coerce":         {`"2007-03-01 13:03:05.123456789Z"`, reflect.ValueOf(time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC))},

	"date_time_custom":             {`"2007-03-01T13:00:00Z"`, reflect.ValueOf(time2time(time.Date(2007, time.March, 1, 13, 0, 0, 0, time.UTC)))},
	"date_time_custom_nano_coerce": {`"2007-03-01T13:03:05.123456789Z"`, reflect.ValueOf(timetypes.NewRFC3339ValueMust("2007-03-01T13:03:05.123456789Z"))},

	"date_time_custom_missing_t_coerce":              {`"2007-03-01 13:03:05Z"`, reflect.ValueOf(time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC)))},
	"date_time_custom_missing_timezone_coerce":       {`"2007-03-01T13:03:05"`, reflect.ValueOf(time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC)))},
	"date_time_custom_missing_timezone_colon_coerce": {`"2007-03-01T13:03:05+0100"`, reflect.ValueOf(time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.FixedZone("", 60*60))))},
	"date_time_custom_nano_missing_t_coerce":         {`"2007-03-01 13:03:05.123456789Z"`, reflect.ValueOf(timetypes.NewRFC3339ValueMust("2007-03-01T13:03:05.123456789Z"))},

	"map_string":                       {`{"foo":"bar"}`, reflect.ValueOf(map[string]string{"foo": "bar"})},
	"map_string_with_sjson_path_chars": {`{":a.b.c*:d*-1e.f":"bar"}`, reflect.ValueOf(map[string]string{":a.b.c*:d*-1e.f": "bar"})},

	"primitive_struct": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}`,
		reflect.ValueOf(Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}),
	},

	"slices": {
		`{"slices":[{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}]}`,
		reflect.ValueOf(Slices{
			Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		}),
	},

	"primitive_pointer_struct": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4,5]}`,
		reflect.ValueOf(PrimitivePointers{
			A: P(false),
			B: P(237628372683),
			C: P(uint(654)),
			D: P(9999.43),
			E: P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		}),
	},

	"datetime_struct": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		reflect.ValueOf(DateTime{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		}),
	},

	"datetime_custom_struct": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		reflect.ValueOf(DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: time2time(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
		}),
	},

	"additional_properties": {
		`{"a":true,"bar":"value","foo":true}`,
		reflect.ValueOf(AdditionalProperties{
			A: true,
			Extras: map[string]jsontypes.Normalized{
				"bar": jsontypes.NewNormalizedValue(`"value"`),
				"foo": jsontypes.NewNormalizedValue("true"),
			},
		}),
	},

	"recursive_struct": {
		`{"child":{"name":"Alex"},"name":"Robert"}`,
		reflect.ValueOf(Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}}),
	},

	"tfsdk_null_string":  {"", reflect.ValueOf(types.StringNull())},
	"tfsdk_null_int":     {"", reflect.ValueOf(types.Int64Null())},
	"tfsdk_null_float":   {"", reflect.ValueOf(types.Float64Null())},
	"tfsdk_null_bool":    {"", reflect.ValueOf(types.BoolNull())},
	"tfsdk_null_dynamic": {"", reflect.ValueOf(types.DynamicNull())},

	"tfsdk_string":             {`"hey"`, reflect.ValueOf(types.StringValue("hey"))},
	"tfsdk_true":               {"true", reflect.ValueOf(types.BoolValue(true))},
	"tfsdk_false":              {"false", reflect.ValueOf(types.BoolValue(false))},
	"tfsdk_int":                {"1", reflect.ValueOf(types.Int64Value(1))},
	"tfsdk_int_bigger":         {"12324", reflect.ValueOf(types.Int64Value(12324))},
	"tfsdk_int_string_coerce":  {`"65"`, reflect.ValueOf(types.Int64Value(65))},
	"tfsdk_int_boolean_coerce": {"true", reflect.ValueOf(types.BoolValue(true))},
	"tfsdk_float_1.54":         {"1.54", reflect.ValueOf(type_helpers.NewFloat64ValueFromStringUnsafe("1.54"))},
	"tfsdk_float_1.89":         {"1.89", reflect.ValueOf(type_helpers.NewFloat64ValueFromStringUnsafe("1.89"))},
	"tfsdk_array_ptr":          {"[\"hi\",null]", reflect.ValueOf(&[]types.String{types.StringValue("hi"), types.StringNull()})},
	"tfsdk_dynamic_string":     {`"hey"`, reflect.ValueOf(types.DynamicValue(types.StringValue("hey")))},
	"tfsdk_dynamic_int":        {"5", reflect.ValueOf(types.DynamicValue(types.Int64Value(5)))},

	"tfsdk_list": {
		"[1,2,3]",
		reflect.ValueOf(types.ListValueMust(
			basetypes.Int64Type{},
			[]attr.Value{basetypes.NewInt64Value(1), basetypes.NewInt64Value(2), basetypes.NewInt64Value(3)},
		)),
	},

	"tfsdk_object": {
		`{"baz":4,"foo":"bar"}`,
		reflect.ValueOf(types.ObjectValueMust(
			map[string]attr.Type{"baz": basetypes.Int64Type{}, "foo": basetypes.StringType{}},
			map[string]attr.Value{"baz": basetypes.NewInt64Value(4), "foo": basetypes.NewStringValue("bar")},
		)),
	},

	"tfsdk_dynamic_object": {
		`{"baz":4,"foo":"bar"}`,
		reflect.ValueOf(types.DynamicValue(
			types.ObjectValueMust(
				map[string]attr.Type{"baz": basetypes.Int64Type{}, "foo": basetypes.StringType{}},
				map[string]attr.Value{"baz": basetypes.NewInt64Value(4), "foo": basetypes.NewStringValue("bar")},
			),
		)),
	},

	"embedded_tfsdk_struct": {
		`{"bool_value":true,` +
			`"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},` +
			`"data_object":{"data_object":{"nested_int":19},"embedded_int":18,"embedded_string":"embedded_data_string_value"},` +
			`"float_value":3.14,` +
			`"list_object":["hi_list","there_list"],` +
			`"map_object":{"hi_map":"there_map"},` +
			`"nested_object_list":[{"embedded_int":20,"embedded_string":"nested_object_string"}],` +
			`"nested_object_map":{"nested_object_map_key":{"embedded_int":21,"embedded_string":"nested_object_string_in_map"}},` +
			`"nested_object_set":[{"embedded_int":21,"embedded_string":"nested_object_string_in_set"}],` +
			`"optional_array":["hi","there"],` +
			`"set_object":["hi_set","there_set"],` +
			`"string_value":"string_value"}`,
		reflect.ValueOf(TfsdkStructs{
			BoolValue:   types.BoolValue(true),
			StringValue: types.StringValue("string_value"),
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
			DataObject: customfield.NewObjectMust(ctx, &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_data_string_value"),
				EmbeddedInt:    types.Int64Value(18),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(19),
				}),
			}),
			ListObject: customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("hi_list"), types.StringValue("there_list")}),
			NestedObjectList: customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{{
				EmbeddedString: types.StringValue("nested_object_string"),
				EmbeddedInt:    types.Int64Value(20),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			}}),
			SetObject: customfield.NewSetMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("hi_set"), types.StringValue("there_set")}),
			NestedObjectSet: customfield.NewObjectSetMust(ctx, []EmbeddedTfsdkStruct{{
				EmbeddedString: types.StringValue("nested_object_string_in_set"),
				EmbeddedInt:    types.Int64Value(21),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			}}),
			MapObject: customfield.NewMapMust[basetypes.StringValue](ctx, map[string]types.String{"hi_map": types.StringValue("there_map")}),
			NestedObjectMap: customfield.NewObjectMapMust(ctx, map[string]EmbeddedTfsdkStruct{"nested_object_map_key": {
				EmbeddedString: types.StringValue("nested_object_string_in_map"),
				EmbeddedInt:    types.Int64Value(21),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			}}),
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
		}),
	},

	"customfield_null_object": {
		"",
		reflect.ValueOf(customfield.NullObject[DoubleNestedStruct](ctx)),
	},

	"json_struct_nil1": {`{}`, reflect.ValueOf(JSONModel{})},
	"json_struct_nil2": {`{}`, reflect.ValueOf(JSONModel{})},
}

type ListWithNestedObj struct {
	A customfield.NestedObjectList[Embedded2] `tfsdk:"tfsdk_a" json:"a"`
}

type Embedded2 struct {
	B types.String      `tfsdk:"tfsdk_b" json:"b"`
	C *Inner            `tfsdk:"tfsdk_c" json:"c"`
	D *[]*Inner         `tfsdk:"tfsdk_d" json:"d"`
	E []string          `tfsdk:"tfsdk_e" json:"e"`
	F *map[string]Inner `tfsdk:"tfsdk_f" json:"f"`
}

type Inner struct {
	D types.String `tfsdk:"tfsdk_d" json:"d"`
}

var decodeOnlyTests = map[string]struct {
	buf string
	val reflect.Value
}{
	"datetime_custom_struct_fractional_seconds": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05.036Z"}`,
		reflect.ValueOf(DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: timetypes.NewRFC3339ValueMust("2006-01-02T15:04:05.036Z"),
		}),
	},

	"datetime_custom_struct_trailing_zeros": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05.500+00:00"}`,
		reflect.ValueOf(DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: timetypes.NewRFC3339ValueMust("2006-01-02T15:04:05.500+00:00"),
		}),
	},

	"datetime_custom_struct_comma_fraction": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05,5Z"}`,
		reflect.ValueOf(DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: timetypes.NewRFC3339ValueMust("2006-01-02T15:04:05.5Z"),
		}),
	},

	"datetime_custom_struct_lenient_fraction_offset": {
		`{"date":"2006-01-02","date-time":"2006-01-02 15:04:05.5+0100"}`,
		reflect.ValueOf(DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: timetypes.NewRFC3339ValueMust("2006-01-02T15:04:05.5+01:00"),
		}),
	},

	"datetime_custom_struct_offset": {
		`{"date":"2006-01-02","date-time":"2006-01-02T18:04:05.5+03:00"}`,
		reflect.ValueOf(DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: timetypes.NewRFC3339ValueMust("2006-01-02T18:04:05.5+03:00"),
		}),
	},

	"tfsdk_struct_decode": {
		`{"result":{"c":"7887590e1967befa70f48ffe9f61ce80","a":"88281d6015751d6172e7313b0c665b5e","extra":"property","another":2,"b":"http://example.com/example.html\t20"}}`,
		reflect.ValueOf(ResultEnvelope{RecordsModel{
			A: types.StringValue("88281d6015751d6172e7313b0c665b5e"),
			B: types.StringValue("http://example.com/example.html\t20"),
			C: types.StringValue("7887590e1967befa70f48ffe9f61ce80"),
		}}),
	},

	"embedded_tfsdk_struct_nil": {
		`{}`,
		reflect.ValueOf(TfsdkStructs{
			BoolValue:        types.BoolNull(),
			StringValue:      types.StringNull(),
			Data:             nil,
			DataObject:       customfield.NullObject[EmbeddedTfsdkStruct](ctx),
			ListObject:       customfield.NullList[basetypes.StringValue](ctx),
			NestedObjectList: customfield.NullObjectList[EmbeddedTfsdkStruct](ctx),
			SetObject:        customfield.NullSet[basetypes.StringValue](ctx),
			NestedObjectSet:  customfield.NullObjectSet[EmbeddedTfsdkStruct](ctx),
			MapObject:        customfield.NullMap[basetypes.StringValue](ctx),
			NestedObjectMap:  customfield.NullObjectMap[EmbeddedTfsdkStruct](ctx),
			FloatValue:       types.Float64Null(),
			OptionalArray:    nil,
		}),
	},

	"json_struct_decode": {
		`{"arr":[true,1,"one"],"arr2":[true,1,"one"],"bol":false,"map":{"nil":null,"bol":false,"str":"two"},"map2":{"bol":false,"nil":null,"str":"two"},"nil":null,"num":2,"str":"two"}`,
		reflect.ValueOf(JSONModel{
			Arr:  jsontypes.NewNormalizedValue(`[true,1,"one"]`),
			Bol:  jsontypes.NewNormalizedValue("false"),
			Map:  jsontypes.NewNormalizedValue(`{"nil":null,"bol":false,"str":"two"}`),
			Nil:  jsontypes.NewNormalizedNull(),
			Num:  jsontypes.NewNormalizedValue("2"),
			Str:  jsontypes.NewNormalizedValue(`"two"`),
			Arr2: []jsontypes.Normalized{jsontypes.NewNormalizedValue("true"), jsontypes.NewNormalizedValue("1"), jsontypes.NewNormalizedValue(`"one"`)},
			Map2: map[string]jsontypes.Normalized{"nil": jsontypes.NewNormalizedNull(), "bol": jsontypes.NewNormalizedValue("false"), "str": jsontypes.NewNormalizedValue(`"two"`)},
		}),
	},

	"json_struct_nil3": {`{"nil":null}`, reflect.ValueOf(JSONModel{})},

	"nested_object_list_missing_nested_field": {
		`{"a":[{"b":"foo"}]}`,
		reflect.ValueOf(ListWithNestedObj{
			A: customfield.NewObjectListMust(ctx, []Embedded2{
				{
					B: types.StringValue("foo"),
					C: nil,
					D: nil,
					E: nil,
					F: nil,
				},
			}),
		}),
	},
}

var encodeOnlyTests = map[string]struct {
	buf string
	val reflect.Value
}{
	"tfsdk_struct_encode": {
		`{"result":{"a":"88281d6015751d6172e7313b0c665b5e","b":"http://example.com/example.html\t20"}}`,
		reflect.ValueOf(ResultEnvelope{RecordsModel{
			A: types.StringValue("88281d6015751d6172e7313b0c665b5e"),
			B: types.StringValue("http://example.com/example.html\t20"),
			C: types.StringValue("7887590e1967befa70f48ffe9f61ce80"),
		}}),
	},

	"embedded_tfsdk_struct_nil": {
		`{}`,
		reflect.ValueOf(TfsdkStructs{
			BoolValue:   types.BoolNull(),
			StringValue: types.StringNull(),
			FloatValue:  types.Float64Null(),
		}),
	},

	"json_struct_encode": {
		`{"arr":[true,1,"one"],"arr2":[true,1,"one"],"bol":false,"map":{"nil":null,"bol":false,"str":"two"},"map2":{"bol":false,"nil":null,"str":"two"},"nil":null,"num":2,"str":"two"}`,
		reflect.ValueOf(JSONModel{
			Arr:  jsontypes.NewNormalizedValue(`[true,1,"one"]`),
			Bol:  jsontypes.NewNormalizedValue("false"),
			Map:  jsontypes.NewNormalizedValue(`{"nil":null,"bol":false,"str":"two"}`),
			Nil:  jsontypes.NewNormalizedValue("null"),
			Num:  jsontypes.NewNormalizedValue("2"),
			Str:  jsontypes.NewNormalizedValue(`"two"`),
			Arr2: []jsontypes.Normalized{jsontypes.NewNormalizedValue("true"), jsontypes.NewNormalizedValue("1"), jsontypes.NewNormalizedValue(`"one"`)},
			Map2: map[string]jsontypes.Normalized{"nil": jsontypes.NewNormalizedNull(), "bol": jsontypes.NewNormalizedValue("false"), "str": jsontypes.NewNormalizedValue(`"two"`)},
		}),
	},

	"json_struct_nil3": {`{"nil":null}`, reflect.ValueOf(JSONModel{Nil: jsontypes.NewNormalizedValue("null")})},

	"tfsdk_dynamic_number": {"5", reflect.ValueOf(types.DynamicValue(type_helpers.NewNumberValueFromStringUnsafe("5")))},

	"tfsdk_dynamic_tuple": {
		`[5,"hi"]`,
		reflect.ValueOf(types.DynamicValue(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		))),
	},

	"tfsdk_tuple": {
		`[5,"hi"]`,
		reflect.ValueOf(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		)),
	},

	"tfsdk_nested_tuple": {
		`[10,["hey","there"]]`,
		reflect.ValueOf(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.ListType{ElemType: basetypes.StringType{}}},
			[]attr.Value{basetypes.NewInt64Value(10), types.ListValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("hey"), basetypes.NewStringValue("there")})},
		)),
	},

	"complex_nested_list_object": {
		`{"a":[{"b":"foo","c":{"d":"pointer_inner"},"d":[{"d":"list_pointer_inner_1"},{"d":"list_pointer_inner_2"}],"e":["a","b"],"f":{"a_key":{"d":"a_value"}}}]}`,
		reflect.ValueOf(ListWithNestedObj{
			A: customfield.NewObjectListMust(ctx, []Embedded2{
				{
					B: types.StringValue("foo"),
					C: P(Inner{D: types.StringValue("pointer_inner")}),
					D: P([]*Inner{P(Inner{D: types.StringValue("list_pointer_inner_1")}), P(Inner{D: types.StringValue("list_pointer_inner_2")})}),
					E: []string{"a", "b"},
					F: P(map[string]Inner{
						"a_key": {D: types.StringValue("a_value")},
					}),
				},
			}),
		}),
	},

	"nested_map_pointer": {
		`{"outer":[{"a":{"a.b.*":"*"}}]}`,
		reflect.ValueOf(struct {
			Outer *[]*structWithMap `json:"outer" api:"required"`
		}{
			Outer: P([]*structWithMap{
				P(structWithMap{A: P(map[string]types.String{"a.b.*": types.StringValue("*")})}),
			}),
		}),
	},
}

type structWithMap struct {
	A *map[string]types.String `json:"a" api:"required"`
}

func TestDecode(t *testing.T) {
	spew.Config.SortKeys = true
	for name, test := range merge(tests, decodeOnlyTests) {
		t.Run(name, func(t *testing.T) {
			resultValue := reflect.New(test.val.Type())
			if err := unmarshalValue([]byte(test.buf), resultValue); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", resultValue, err)
			}
			result := resultValue.Elem()
			if !decodedEqual(result, test.val) {
				t.Fatalf("incorrect deserialization for '%s':\nexpected:\n%s\nactual:\n%s\n", test.buf, spew.Sdump(test.val.Interface()), spew.Sdump(result.Interface()))
			}
		})
	}
}

func unmarshalValue(raw []byte, to reflect.Value) error {
	return (&decoderBuilder{dateFormat: time.RFC3339}).unmarshal(raw, to)
}

func marshalValues(e encoder, plan reflect.Value, state reflect.Value) ([]byte, error) {
	e.dateFormat = time.RFC3339
	return e.marshal(plan, state)
}

func decodedEqual(result, want reflect.Value) bool {
	resultTime, isTime := reflect.TypeAssert[time.Time](result)
	wantTime, wantIsTime := reflect.TypeAssert[time.Time](want)
	if !isTime || !wantIsTime {
		return reflect.DeepEqual(result.Interface(), want.Interface())
	}
	_, resultOffset := resultTime.Zone()
	_, wantOffset := wantTime.Zone()
	return resultTime.Equal(wantTime) && resultOffset == wantOffset
}

func TestEncode(t *testing.T) {
	for name, test := range merge(tests, encodeOnlyTests) {
		if strings.HasSuffix(name, "_coerce") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			raw, err := marshalValues(encoder{}, test.val, test.val)
			if err != nil {
				t.Fatalf("serialization of %v failed with error %v", test.val, err)
			}
			if string(raw) != test.buf {
				var expected, actual string
				errExpected := formatJSON(test.buf, &expected)
				if errExpected != nil {
					panic(fmt.Sprintf("invalid expected JSON:\n%s\n%v", test.buf, errExpected))
				}
				errActual := formatJSON(string(raw), &actual)
				if errActual != nil {
					t.Fatalf("invalid actual JSON:\n%s\n%v", string(raw), errActual)
				}
				t.Fatalf("expected:\n%s\nto serialize to \n%s\n but got \n%s\n", spew.Sdump(test.val.Interface()), expected, actual)
			}
		})

	}
}

var updateTests = map[string]struct {
	state         reflect.Value
	plan          reflect.Value
	expected      string
	expectedPatch string
}{
	"true":           {reflect.ValueOf(true), reflect.ValueOf(true), "true", ""},
	"terraform_true": {reflect.ValueOf(types.BoolValue(true)), reflect.ValueOf(types.BoolValue(true)), "true", ""},

	"null to true":   {reflect.ValueOf(types.BoolNull()), reflect.ValueOf(types.BoolValue(true)), "true", "true"},
	"null to false":  {reflect.ValueOf(types.BoolNull()), reflect.ValueOf(types.BoolValue(false)), "false", "false"},
	"false to true":  {reflect.ValueOf(types.BoolValue(false)), reflect.ValueOf(types.BoolValue(true)), "true", "true"},
	"unset bool":     {reflect.ValueOf(types.BoolValue(false)), reflect.ValueOf(types.BoolNull()), "null", "null"},
	"omit null bool": {reflect.ValueOf(types.BoolNull()), reflect.ValueOf(types.BoolNull()), "", ""},

	"string set":           {reflect.ValueOf(types.StringNull()), reflect.ValueOf(types.StringValue("two")), `"two"`, `"two"`},
	"null to empty string": {reflect.ValueOf(types.StringNull()), reflect.ValueOf(types.StringValue("")), `""`, `""`},
	"string update":        {reflect.ValueOf(types.StringValue("one")), reflect.ValueOf(types.StringValue("two")), `"two"`, `"two"`},
	"unset string":         {reflect.ValueOf(types.StringValue("hey")), reflect.ValueOf(types.StringNull()), "null", "null"},
	"omit null string":     {reflect.ValueOf(types.StringNull()), reflect.ValueOf(types.StringNull()), "", ""},
	"string unchanged":     {reflect.ValueOf(types.StringValue("one")), reflect.ValueOf(types.StringValue("one")), `"one"`, ""},

	"null to zero int": {reflect.ValueOf(types.Int64Null()), reflect.ValueOf(types.Int64Value(0)), "0", "0"},
	"int set":          {reflect.ValueOf(types.Int64Null()), reflect.ValueOf(types.Int64Value(42)), "42", "42"},
	"int update":       {reflect.ValueOf(types.Int64Value(42)), reflect.ValueOf(types.Int64Value(43)), "43", "43"},
	"unset int":        {reflect.ValueOf(types.Int64Value(42)), reflect.ValueOf(types.Int64Null()), "null", "null"},
	"omit null int":    {reflect.ValueOf(types.Int64Null()), reflect.ValueOf(types.Int64Null()), "", ""},
	"int unchanged":    {reflect.ValueOf(types.Int64Value(42)), reflect.ValueOf(types.Int64Value(42)), "42", ""},

	"tuple set": {
		reflect.ValueOf(types.TupleNull([]attr.Type{types.Int64Type, types.StringType})),
		reflect.ValueOf(types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")})),
		`[1,"two"]`,
		`[1,"two"]`,
	},
	"tuple update": {
		reflect.ValueOf(types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")})),
		reflect.ValueOf(types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("three")})),
		`[1,"three"]`,
		`[1,"three"]`,
	},
	"tuple unset": {
		reflect.ValueOf(types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")})),
		reflect.ValueOf(types.TupleNull([]attr.Type{types.Int64Type, types.StringType})),
		`null`,
		`null`,
	},
	"tuple omit null": {
		reflect.ValueOf(types.TupleNull([]attr.Type{types.Int64Type, types.StringType})),
		reflect.ValueOf(types.TupleNull([]attr.Type{types.Int64Type, types.StringType})),
		``,
		``,
	},
	"tuple unchanged": {
		reflect.ValueOf(types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")})),
		reflect.ValueOf(types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")})),
		`[1,"two"]`,
		``,
	},

	"dynamic omit null":                     {reflect.ValueOf(types.DynamicNull()), reflect.ValueOf(types.DynamicNull()), "", ""},
	"dynamic omit underlying null state":    {reflect.ValueOf(types.DynamicValue(types.Int64Null())), reflect.ValueOf(types.DynamicNull()), "", ""},
	"dynamic omit underlying null plan":     {reflect.ValueOf(types.DynamicNull()), reflect.ValueOf(types.DynamicValue(types.Int64Null())), "", ""},
	"dynamic omit unknown":                  {reflect.ValueOf(types.DynamicUnknown()), reflect.ValueOf(types.DynamicUnknown()), "", ""},
	"dynamic omit underlying unknown state": {reflect.ValueOf(types.DynamicValue(types.Int64Unknown())), reflect.ValueOf(types.DynamicUnknown()), "", ""},
	"dynamic omit underlying unknown plan":  {reflect.ValueOf(types.DynamicUnknown()), reflect.ValueOf(types.DynamicValue(types.Int64Unknown())), "", ""},
	"dynamic unset null":                    {reflect.ValueOf(types.DynamicValue(types.Int64Value(4))), reflect.ValueOf(types.DynamicNull()), "null", "null"},
	"dynamic int set":                       {reflect.ValueOf(types.DynamicNull()), reflect.ValueOf(types.DynamicValue(types.Int64Value(5))), "5", "5"},
	"dynamic int update":                    {reflect.ValueOf(types.DynamicValue(types.Int64Value(4))), reflect.ValueOf(types.DynamicValue(types.Int64Value(5))), "5", "5"},
	"dynamic int unchanged":                 {reflect.ValueOf(types.DynamicValue(types.Int64Value(4))), reflect.ValueOf(types.DynamicValue(types.Int64Value(4))), "4", ""},

	"dynamic list to tuple conversion": {
		reflect.ValueOf(types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}))),
		reflect.ValueOf(types.DynamicValue(types.TupleValueMust([]attr.Type{types.StringType, types.StringType}, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}))),
		`["foo","bar"]`,
		``,
	},

	"normalized list to tuple conversion": {
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}))),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust([]attr.Type{types.StringType, types.StringType}, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}))),
		`["foo","bar"]`,
		``,
	},

	"dynamic tuple to list conversion": {
		reflect.ValueOf(types.DynamicValue(types.TupleValueMust([]attr.Type{types.StringType, types.StringType}, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}))),
		reflect.ValueOf(types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}))),
		`["foo","bar"]`,
		``,
	},

	"normalized dynamic tuple to list conversion": {
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust([]attr.Type{types.StringType, types.StringType}, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}))),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}))),
		`["foo","bar"]`,
		``,
	},

	"dynamic list to heterogeneous tuple": {
		reflect.ValueOf(types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hello"), types.StringValue("world")}))),
		reflect.ValueOf(types.DynamicValue(types.TupleValueMust([]attr.Type{types.StringType, types.Int64Type}, []attr.Value{types.StringValue("hello"), types.Int64Value(42)}))),
		`["hello",42]`,
		`["hello",42]`,
	},

	"normalized dynamic list to heterogeneous tuple": {
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hello"), types.StringValue("world")}))),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust([]attr.Type{types.StringType, types.Int64Type}, []attr.Value{types.StringValue("hello"), types.Int64Value(42)}))),
		`["hello",42]`,
		`["hello",42]`,
	},

	"set struct fields": {
		reflect.ValueOf(TfsdkStructs{}),
		reflect.ValueOf(TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		}),
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
	},

	"update some struct fields": {
		reflect.ValueOf(TfsdkStructs{
			BoolValue:   types.BoolValue(true),
			StringValue: types.StringValue("string_value"),
			FloatValue:  type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
		}),
		reflect.ValueOf(TfsdkStructs{
			BoolValue:   types.BoolValue(false),
			StringValue: types.StringValue("another_string"),
			FloatValue:  type_helpers.NewFloat64ValueFromStringUnsafe("1.14"),
		}),
		`{"bool_value":false,"float_value":1.14,"string_value":"another_string"}`,
		`{"bool_value":false,"float_value":1.14,"string_value":"another_string"}`,
	},

	"unset nested struct fields": {
		reflect.ValueOf(TfsdkStructs{
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedInt: types.Int64Value(17),
			},
		}),
		reflect.ValueOf(TfsdkStructs{
			OptionalArray: &[]types.String{types.StringValue("hi")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedInt: types.Int64Null(),
			},
		}),
		`{"data":{"embedded_int":null},"optional_array":["hi"]}`,
		`{"data":{"embedded_int":null},"optional_array":["hi"]}`,
	},

	"unset struct fields": {
		reflect.ValueOf(TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
		}),
		reflect.ValueOf(TfsdkStructs{}),
		`{"bool_value":null,"data":null,"float_value":null,"optional_array":null,"string_value":null}`,
		`{"bool_value":null,"data":null,"float_value":null,"optional_array":null,"string_value":null}`,
	},

	"nested object null to non-null": {
		reflect.ValueOf(TfsdkStructs{
			DataObject: customfield.NullObject[EmbeddedTfsdkStruct](ctx),
		}),
		reflect.ValueOf(TfsdkStructs{
			DataObject: customfield.NewObjectMust(ctx, &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("new_value"),
				EmbeddedInt:    types.Int64Value(42),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			}),
		}),
		`{"data_object":{"embedded_int":42,"embedded_string":"new_value"}}`,
		`{"data_object":{"embedded_int":42,"embedded_string":"new_value"}}`,
	},

	"nested object non-null to null": {
		reflect.ValueOf(TfsdkStructs{
			DataObject: customfield.NewObjectMust(ctx, &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("old_value"),
				EmbeddedInt:    types.Int64Value(10),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			}),
		}),
		reflect.ValueOf(TfsdkStructs{
			DataObject: customfield.NullObject[EmbeddedTfsdkStruct](ctx),
		}),
		`{"data_object":null}`,
		`{"data_object":null}`,
	},

	"nested object with null state preserves null": {
		reflect.ValueOf(customfield.NullObject[TfsdkStructs](ctx)),
		reflect.ValueOf(customfield.NullObject[TfsdkStructs](ctx)),
		``,
		``,
	},

	"nested object list null to non-null": {
		reflect.ValueOf(customfield.NullObjectList[TfsdkStructs](ctx)),
		reflect.ValueOf(customfield.NewObjectListMust(ctx, []TfsdkStructs{
			{
				BoolValue:   types.BoolValue(true),
				StringValue: types.StringValue("test"),
			},
		})),
		`[{"bool_value":true,"string_value":"test"}]`,
		`[{"bool_value":true,"string_value":"test"}]`,
	},

	"nested object list non-null to null": {
		reflect.ValueOf(customfield.NewObjectListMust(ctx, []TfsdkStructs{
			{
				BoolValue:   types.BoolValue(true),
				StringValue: types.StringValue("test"),
			},
		})),
		reflect.ValueOf(customfield.NullObjectList[TfsdkStructs](ctx)),
		`null`,
		`null`,
	},

	"nested object list with null state preserves null": {
		reflect.ValueOf(customfield.NullObjectList[TfsdkStructs](ctx)),
		reflect.ValueOf(customfield.NullObjectList[TfsdkStructs](ctx)),
		``,
		``,
	},

	"nested object map null to non-null": {
		reflect.ValueOf(customfield.NullObjectMap[TfsdkStructs](ctx)),
		reflect.ValueOf(customfield.NewObjectMapMust(ctx, map[string]TfsdkStructs{
			"key1": {
				BoolValue:   types.BoolValue(true),
				StringValue: types.StringValue("test"),
			},
		})),
		`{"key1":{"bool_value":true,"string_value":"test"}}`,
		`{"key1":{"bool_value":true,"string_value":"test"}}`,
	},

	"nested object map non-null to null": {
		reflect.ValueOf(customfield.NewObjectMapMust(ctx, map[string]TfsdkStructs{
			"key1": {
				BoolValue:   types.BoolValue(true),
				StringValue: types.StringValue("test"),
			},
		})),
		reflect.ValueOf(customfield.NullObjectMap[TfsdkStructs](ctx)),
		`null`,
		`null`,
	},

	"nested object map with null state preserves null": {
		reflect.ValueOf(customfield.NullObjectMap[TfsdkStructs](ctx)),
		reflect.ValueOf(customfield.NullObjectMap[TfsdkStructs](ctx)),
		``,
		``,
	},

	"nested object set null to non-null": {
		reflect.ValueOf(customfield.NullObjectSet[TfsdkStructs](ctx)),
		reflect.ValueOf(customfield.NewObjectSetMust(ctx, []TfsdkStructs{
			{
				BoolValue:   types.BoolValue(true),
				StringValue: types.StringValue("test"),
			},
		})),
		`[{"bool_value":true,"string_value":"test"}]`,
		`[{"bool_value":true,"string_value":"test"}]`,
	},

	"nested object set non-null to null": {
		reflect.ValueOf(customfield.NewObjectSetMust(ctx, []TfsdkStructs{
			{
				BoolValue:   types.BoolValue(true),
				StringValue: types.StringValue("test"),
			},
		})),
		reflect.ValueOf(customfield.NullObjectSet[TfsdkStructs](ctx)),
		`null`,
		`null`,
	},

	"nested object set with null state preserves null": {
		reflect.ValueOf(customfield.NullObjectSet[TfsdkStructs](ctx)),
		reflect.ValueOf(customfield.NullObjectSet[TfsdkStructs](ctx)),
		``,
		``,
	},

	"list null to non-null": {
		reflect.ValueOf(customfield.NullList[types.String](ctx)),
		reflect.ValueOf(customfield.NewListMust[types.String](ctx, []attr.Value{
			types.StringValue("test1"),
			types.StringValue("test2"),
		})),
		`["test1","test2"]`,
		`["test1","test2"]`,
	},

	"list non-null to null": {
		reflect.ValueOf(customfield.NewListMust[types.String](ctx, []attr.Value{
			types.StringValue("test1"),
			types.StringValue("test2"),
		})),
		reflect.ValueOf(customfield.NullList[types.String](ctx)),
		`null`,
		`null`,
	},

	"list with null state preserves null": {
		reflect.ValueOf(customfield.NullList[types.String](ctx)),
		reflect.ValueOf(customfield.NullList[types.String](ctx)),
		``,
		``,
	},

	"map null to non-null": {
		reflect.ValueOf(customfield.NullMap[types.String](ctx)),
		reflect.ValueOf(customfield.NewMapMust[types.String](ctx, map[string]types.String{
			"key1": types.StringValue("value1"),
			"key2": types.StringValue("value2"),
		})),
		`{"key1":"value1","key2":"value2"}`,
		`{"key1":"value1","key2":"value2"}`,
	},

	"map non-null to null": {
		reflect.ValueOf(customfield.NewMapMust[types.String](ctx, map[string]types.String{
			"key1": types.StringValue("value1"),
			"key2": types.StringValue("value2"),
		})),
		reflect.ValueOf(customfield.NullMap[types.String](ctx)),
		`null`,
		`null`,
	},

	"map with null state preserves null": {
		reflect.ValueOf(customfield.NullMap[types.String](ctx)),
		reflect.ValueOf(customfield.NullMap[types.String](ctx)),
		``,
		``,
	},

	"set null to non-null": {
		reflect.ValueOf(customfield.NullSet[types.String](ctx)),
		reflect.ValueOf(customfield.NewSetMust[types.String](ctx, []attr.Value{
			types.StringValue("test1"),
			types.StringValue("test2"),
		})),
		`["test1","test2"]`,
		`["test1","test2"]`,
	},

	"set non-null to null": {
		reflect.ValueOf(customfield.NewSetMust[types.String](ctx, []attr.Value{
			types.StringValue("test1"),
			types.StringValue("test2"),
		})),
		reflect.ValueOf(customfield.NullSet[types.String](ctx)),
		`null`,
		`null`,
	},

	"set with null state preserves null": {
		reflect.ValueOf(customfield.NullSet[types.String](ctx)),
		reflect.ValueOf(customfield.NullSet[types.String](ctx)),
		``,
		``,
	},

	"set empty array": {
		reflect.ValueOf(TfsdkStructs{
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
		}),
		reflect.ValueOf(TfsdkStructs{
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{},
		}),
		`{"float_value":3.14,"optional_array":[]}`,
		`{"optional_array":[]}`,
	},

	"set nested map": {
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{})),
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		})),
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
	},

	"unchanged nested struct": {
		reflect.ValueOf(TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		}),
		reflect.ValueOf(TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		}),
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
		`{}`,
	},

	"nested value changed in nested struct": {
		reflect.ValueOf(TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		}),
		reflect.ValueOf(TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("changed_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		}),
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"changed_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
		`{"data":{"embedded_string":"changed_string_value"}}`,
	},

	"set array element": {
		reflect.ValueOf(TfsdkStructs{
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("one"), types.StringValue("two")},
			ListObject:    customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("three"), types.StringValue("four")}),
		}),
		reflect.ValueOf(TfsdkStructs{
			FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
			OptionalArray: &[]types.String{types.StringValue("five"), types.StringValue("two")},
			ListObject:    customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("six"), types.StringValue("four")}),
		}),
		`{"float_value":3.14,"list_object":["six","four"],"optional_array":["five","two"]}`,
		`{"list_object":["six","four"],"optional_array":["five","two"]}`,
	},

	"set nested array value": {
		reflect.ValueOf(customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{
			{
				EmbeddedString: types.StringValue("string value"),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
			{
				EmbeddedString: types.StringValue("string value2"),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(19),
				}),
			},
		})),
		reflect.ValueOf(customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{
			{
				EmbeddedString: types.StringValue("string value"),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
			{
				EmbeddedString: types.StringValue("string value2"),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(20),
				}),
			},
		})),
		`[{"embedded_string":"string value"},{"data_object":{"nested_int":20},"embedded_string":"string value2"}]`,
		`[{"embedded_string":"string value"},{"data_object":{"nested_int":20},"embedded_string":"string value2"}]`,
	},

	"remove array value encodes": {
		reflect.ValueOf(customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{
			{
				EmbeddedString: types.StringValue("string value"),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
			{
				EmbeddedString: types.StringValue("string value2"),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(20),
				}),
			},
		})),
		reflect.ValueOf(customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{
			{
				EmbeddedString: types.StringValue("string value"),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
		})),
		`[{"embedded_string":"string value"}]`,
		`[{"embedded_string":"string value"}]`,
	},

	"set custom map list": {
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{})),
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		})),
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
	},

	"set built-in map list": {
		reflect.ValueOf(map[string][]*string{}),
		reflect.ValueOf(map[string][]*string{
			"Key1": {P("Value1")},
			"Key2": {P("Value2")},
		}),
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
	},

	"remove all keys from a custom map": {
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		})),
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{})),
		`{}`,
		`{"Key1":null,"Key2":null}`,
	},

	"update to add a key to a custom map": {
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
		})),
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		})),
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
		`{"Key2":["Value2"]}`,
	},

	"update a nested array in a custom map": {
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		})),
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value3"), basetypes.NewStringValue("Value2")})),
		})),
		`{"Key1":["Value1"],"Key2":["Value3","Value2"]}`,
		`{"Key2":["Value3","Value2"]}`,
	},

	"unset custom map": {
		reflect.ValueOf(customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
		})),
		reflect.ValueOf(customfield.NullMap[customfield.List[types.String]](ctx)),
		`null`,
		`null`,
	},

	"unset key in built-in map": {
		reflect.ValueOf(map[string]*string{
			"Key1": P("Value1"),
			"Key2": P("Value2"),
		}),
		reflect.ValueOf(map[string]*string{
			"Key1": P("Value1"),
		}),
		`{"Key1":"Value1"}`,
		`{"Key2":null}`,
	},

	"set custom object map": {
		reflect.ValueOf(customfield.NullObjectMap[TfsdkStructs](ctx)),
		reflect.ValueOf(customfield.NewObjectMapMust(ctx, map[string]TfsdkStructs{
			"Key1": {
				BoolValue:     types.BoolValue(true),
				StringValue:   types.StringValue("string_value"),
				FloatValue:    type_helpers.NewFloat64ValueFromStringUnsafe("3.14"),
				OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
				Data: &EmbeddedTfsdkStruct{
					EmbeddedString: types.StringValue("embedded_string_value"),
					EmbeddedInt:    types.Int64Value(17),
				},
			},
		})),
		`{"Key1":{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}}`,
		`{"Key1":{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}}`,
	},

	"set nested value on custom object map": {
		reflect.ValueOf(customfield.NewObjectMapMust(ctx, map[string]TfsdkStructs{
			"OuterKey": {
				NestedObjectMap: customfield.NewObjectMapMust(ctx, map[string]EmbeddedTfsdkStruct{
					"NestedKey": {
						EmbeddedInt:    types.Int64Value(16),
						EmbeddedString: types.StringValue("nested_string_value"),
					},
				}),
			},
		})),
		reflect.ValueOf(customfield.NewObjectMapMust(ctx, map[string]TfsdkStructs{
			"OuterKey": {
				NestedObjectMap: customfield.NewObjectMapMust(ctx, map[string]EmbeddedTfsdkStruct{
					"NestedKey": {
						EmbeddedInt:    types.Int64Value(17),
						EmbeddedString: types.StringValue("nested_string_value"),
					},
				}),
			},
		})),
		`{"OuterKey":{"nested_object_map":{"NestedKey":{"embedded_int":17,"embedded_string":"nested_string_value"}}}}`,
		`{"OuterKey":{"nested_object_map":{"NestedKey":{"embedded_int":17}}}}`,
	},

	"encode_state_for_unknown with unknown plan": {
		reflect.ValueOf(EncodeStateForUnknownStruct{
			NormalField:                     types.StringValue("state_normal"),
			ComputedWithForceEncode:         types.StringValue("computed value from state"),
			ComputedWithStateEncode:         types.StringValue("computed value 2"),
			ComputedOptionalWithStateEncode: types.StringValue("computed optional from state"),
			ComputedOptionalRegular:         types.StringValue("computed optional regular"),
			ComputedRegular:                 types.StringValue("computed regular"),
		}),
		reflect.ValueOf(EncodeStateForUnknownStruct{
			NormalField:                     types.StringUnknown(),
			ComputedWithForceEncode:         types.StringUnknown(),
			ComputedWithStateEncode:         types.StringUnknown(),
			ComputedOptionalWithStateEncode: types.StringUnknown(),
			ComputedOptionalRegular:         types.StringUnknown(),
			ComputedRegular:                 types.StringUnknown(),
		}),
		`{"computed_optional_state_encode":"computed optional from state","computed_state_encode":"computed value 2"}`,
		`{}`,
	},

	"encode_state_for_unknown with known plan": {
		reflect.ValueOf(EncodeStateForUnknownStruct{
			NormalField:                     types.StringValue("state_normal"),
			ComputedWithForceEncode:         types.StringValue("computed value from state"),
			ComputedWithStateEncode:         types.StringValue("computed value 2"),
			ComputedOptionalWithStateEncode: types.StringValue("computed optional from state"),
			ComputedOptionalRegular:         types.StringValue("computed optional regular"),
			ComputedRegular:                 types.StringValue("computed regular"),
		}),
		reflect.ValueOf(EncodeStateForUnknownStruct{
			NormalField:                     types.StringValue("plan normal"),
			ComputedWithForceEncode:         types.StringValue("plan A"),
			ComputedWithStateEncode:         types.StringValue("plan B"),
			ComputedOptionalWithStateEncode: types.StringValue("plan C"),
			ComputedOptionalRegular:         types.StringValue("plan D"),
			ComputedRegular:                 types.StringValue("plan E"),
		}),
		`{"computed_force_encode":"plan A","computed_optional_regular":"plan D","computed_optional_state_encode":"plan C","computed_state_encode":"plan B","normal_field":"plan normal"}`,
		`{"computed_force_encode":"plan A","computed_optional_regular":"plan D","computed_optional_state_encode":"plan C","computed_state_encode":"plan B","normal_field":"plan normal"}`},

	"encode_state_for_unknown with null state": {
		reflect.ValueOf(EncodeStateForUnknownStruct{
			NormalField:                     types.StringNull(),
			ComputedWithForceEncode:         types.StringNull(),
			ComputedWithStateEncode:         types.StringNull(),
			ComputedOptionalWithStateEncode: types.StringNull(),
			ComputedOptionalRegular:         types.StringNull(),
			ComputedRegular:                 types.StringNull(),
		}),
		reflect.ValueOf(EncodeStateForUnknownStruct{
			NormalField:                     types.StringUnknown(),
			ComputedWithForceEncode:         types.StringUnknown(),
			ComputedWithStateEncode:         types.StringUnknown(),
			ComputedOptionalWithStateEncode: types.StringUnknown(),
			ComputedOptionalRegular:         types.StringUnknown(),
			ComputedRegular:                 types.StringUnknown(),
		}),
		`{}`,
		`{}`,
	},
}

func TestUpdateEncoding(t *testing.T) {
	for name, test := range updateTests {
		t.Run(name, func(t *testing.T) {
			t.Run("MarshalForUpdate", func(t *testing.T) {
				raw, err := marshalValues(encoder{root: true}, test.plan, test.state)
				if err != nil {
					t.Fatalf("serialization of %v, %v failed with error %v", test.plan, test.state, err)
				}
				if string(raw) != test.expected {
					t.Fatalf("expected %+#v, %+#v to serialize to \n%s\n but got \n%s\n", test.state, test.plan, test.expected, string(raw))
				}
			})
			t.Run("MarshalForPatch", func(t *testing.T) {
				raw, err := marshalValues(encoder{root: true, patch: true}, test.plan, test.state)
				if err != nil {
					t.Fatalf("serialization of %v, %v failed with error %v", test.plan, test.state, err)
				}
				if string(raw) != test.expectedPatch {
					t.Fatalf("expected %+#v, %+#v to serialize to \n%s\n but got \n%s\n", test.state, test.plan, test.expectedPatch, string(raw))
				}
			})
		})
	}
}

func TestMarshalForPatch_AllFieldsUnknown(t *testing.T) {
	plan := EncodeStateForUnknownStruct{
		NormalField:                     types.StringUnknown(),
		ComputedWithForceEncode:         types.StringUnknown(),
		ComputedWithStateEncode:         types.StringUnknown(),
		ComputedOptionalWithStateEncode: types.StringUnknown(),
		ComputedOptionalRegular:         types.StringUnknown(),
		ComputedRegular:                 types.StringUnknown(),
	}
	state := EncodeStateForUnknownStruct{
		NormalField:                     types.StringValue("old normal"),
		ComputedWithForceEncode:         types.StringValue("old force"),
		ComputedWithStateEncode:         types.StringValue("old state encode"),
		ComputedOptionalWithStateEncode: types.StringValue("old opt state encode"),
		ComputedOptionalRegular:         types.StringValue("old opt regular"),
		ComputedRegular:                 types.StringValue("old computed"),
	}

	raw, err := MarshalForPatch(plan, state)
	if err != nil {
		t.Fatalf("MarshalForPatch failed: %v", err)
	}
	if raw == nil {
		t.Fatal("MarshalForPatch returned nil; expected '{}' so the caller has a valid request body")
	}
	if string(raw) != "{}" {
		t.Fatalf("expected '{}' but got %s", string(raw))
	}
}

var decodeFromValueTests = map[string]struct {
	buf      string
	starting reflect.Value
	expected reflect.Value
}{

	"tfsdk_dynamic_null": {
		`null`,
		reflect.ValueOf(types.DynamicNull()),
		reflect.ValueOf(types.DynamicNull()),
	},

	"tfsdk_dynamic_string_from_null": {
		`"hey"`,
		reflect.ValueOf(types.DynamicNull()),
		reflect.ValueOf(types.DynamicValue(types.StringValue("hey"))),
	},

	"tfsdk_dynamic_string_from_unknown": {
		`"hey"`,
		reflect.ValueOf(types.DynamicUnknown()),
		reflect.ValueOf(types.DynamicValue(types.StringValue("hey"))),
	},

	"tfsdk_dynamic_string_from_value": {
		`"hey"`,
		reflect.ValueOf(types.DynamicValue(types.StringValue("before_value"))),
		reflect.ValueOf(types.DynamicValue(types.StringValue("hey"))),
	},

	"tfsdk_dynamic_number": {
		"5",
		reflect.ValueOf(types.DynamicValue(basetypes.NewNumberNull())),
		reflect.ValueOf(types.DynamicValue(type_helpers.NewNumberValueFromStringUnsafe("5"))),
	},

	"tfsdk_dynamic_int_from_null": {
		`14`,
		reflect.ValueOf(types.DynamicNull()),
		reflect.ValueOf(types.DynamicValue(types.Int64Value(14))),
	},

	"tfsdk_dynamic_int_from_unknown": {
		`14`,
		reflect.ValueOf(types.DynamicUnknown()),
		reflect.ValueOf(types.DynamicValue(types.Int64Value(14))),
	},

	"tfsdk_dynamic_int_from_value": {
		`14`,
		reflect.ValueOf(types.DynamicValue(types.Int64Value(5))),
		reflect.ValueOf(types.DynamicValue(types.Int64Value(14))),
	},

	"tfsdk_dynamic_tuple": {
		`[5,"hi"]`,
		reflect.ValueOf(types.DynamicValue(types.TupleNull(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
		))),
		reflect.ValueOf(types.DynamicValue(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		))),
	},

	"tfsdk_map_value": {
		`{"foo":1,"bar":4}`,
		reflect.ValueOf(types.MapNull(types.Int64Type)),
		reflect.ValueOf(types.MapValueMust(types.Int64Type, map[string]attr.Value{"foo": types.Int64Value(1), "bar": types.Int64Value(4)})),
	},

	"tfsdk_tuple": {
		`[5,"hi"]`,
		reflect.ValueOf(types.TupleNull(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
		)),
		reflect.ValueOf(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		)),
	},

	"tfsdk_tuple_existing": {
		`[10,"hello there"]`,
		reflect.ValueOf(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		)),
		reflect.ValueOf(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(10), basetypes.NewStringValue("hello there")},
		)),
	},

	"tfsdk_tuple_missing_values": {
		`[]`,
		reflect.ValueOf(types.TupleNull(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
		)),
		reflect.ValueOf(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Null(), basetypes.NewStringNull()},
		)),
	},

	"tfsdk_tuple_single_object": {
		`[{"non":"array"}]`,
		reflect.ValueOf(types.TupleNull(
			[]attr.Type{basetypes.ObjectType{AttrTypes: map[string]attr.Type{"non": basetypes.StringType{}}}, basetypes.StringType{}},
		)),
		reflect.ValueOf(types.TupleValueMust(
			[]attr.Type{basetypes.ObjectType{AttrTypes: map[string]attr.Type{"non": basetypes.StringType{}}}, basetypes.StringType{}},
			[]attr.Value{
				basetypes.NewObjectValueMust(
					map[string]attr.Type{"non": basetypes.StringType{}},
					map[string]attr.Value{"non": basetypes.NewStringValue("array")},
				),
				basetypes.NewStringNull(),
			},
		)),
	},

	"tfsdk_tuple_non_array_num_error": {
		`5`,
		reflect.ValueOf(types.TupleNull(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
		)),
		reflect.ValueOf(fmt.Errorf("apijson: cannot deserialize unexpected type Number to types.TupleValue")),
	},

	"tfsdk_tuple_non_array_object_error": {
		`{"non":"array"}`,
		reflect.ValueOf(types.TupleNull(
			[]attr.Type{basetypes.ObjectType{AttrTypes: map[string]attr.Type{"non": basetypes.StringType{}}}, basetypes.StringType{}},
		)),
		reflect.ValueOf(fmt.Errorf("apijson: cannot deserialize unexpected type JSON to types.TupleValue")),
	},

	"tfsdk_map_value_existing_data": {
		`{"foo":1,"bar":4}`,
		reflect.ValueOf(types.MapValueMust(types.Int64Type, map[string]attr.Value{"baz": types.Int64Value(2)})),
		reflect.ValueOf(types.MapValueMust(types.Int64Type, map[string]attr.Value{"foo": types.Int64Value(1), "bar": types.Int64Value(4)})),
	},

	"tfsdk_object_with_attributes": {
		`{"baz":4,"foo":["bar","baz"]}`,
		reflect.ValueOf(types.ObjectNull(
			map[string]attr.Type{"baz": types.Int64Type, "foo": types.SetType{ElemType: types.StringType}},
		)),
		reflect.ValueOf(types.ObjectValueMust(
			map[string]attr.Type{"baz": types.Int64Type, "foo": types.SetType{ElemType: types.StringType}},
			map[string]attr.Value{"baz": types.Int64Value(4), "foo": types.SetValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("baz")})},
		)),
	},

	"tfsdk_dynamic_object_with_attributes": {
		`{"baz":4,"foo":["bar","baz"]}`,
		reflect.ValueOf(types.DynamicValue(
			types.ObjectNull(
				map[string]attr.Type{"baz": types.Int64Type, "foo": types.SetType{ElemType: types.StringType}},
			),
		)),
		reflect.ValueOf(types.DynamicValue(
			types.ObjectValueMust(
				map[string]attr.Type{"baz": types.Int64Type, "foo": types.SetType{ElemType: types.StringType}},
				map[string]attr.Value{"baz": types.Int64Value(4), "foo": types.SetValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("baz")})},
			),
		)),
	},

	"tfsdk_dynamic_object_without_attributes": {
		`{"baz":4,"foo":["bar","baz"]}`,
		reflect.ValueOf(types.DynamicNull()),
		reflect.ValueOf(types.DynamicValue(
			types.ObjectValueMust(
				map[string]attr.Type{"baz": types.Int64Type, "foo": types.ListType{ElemType: types.StringType}},
				map[string]attr.Value{"baz": types.Int64Value(4), "foo": types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("baz")})},
			),
		)),
	},

	"tfsdk_dynamic_heterogeneous_array_inference": {
		`["hello",42]`,
		reflect.ValueOf(types.DynamicNull()),
		reflect.ValueOf(types.DynamicValue(types.TupleValueMust(
			[]attr.Type{types.StringType, types.Int64Type},
			[]attr.Value{types.StringValue("hello"), types.Int64Value(42)},
		))),
	},

	"tfsdk_dynamic_empty_array_inference": {
		`[]`,
		reflect.ValueOf(types.DynamicNull()),
		reflect.ValueOf(types.DynamicValue(types.TupleValueMust([]attr.Type{}, []attr.Value{}))),
	},

	"tfsdk_list_empty_array": {
		`[]`,
		reflect.ValueOf(types.ListNull(types.StringType)),
		reflect.ValueOf(types.ListValueMust(types.StringType, []attr.Value{})),
	},

	"tfsdk_untyped_list_empty_array": {
		`[]`,
		reflect.ValueOf(types.List{}),
		reflect.ValueOf(fmt.Errorf("apijson: cannot deserialize empty array to types.ListValue without an element type")),
	},

	"tfsdk_normalized_dynamic_existing_list_empty_array": {
		`[]`,
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("url")}))),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(types.StringType, []attr.Value{}))),
	},

	"tfsdk_normalized_dynamic_empty_array_inference": {
		`[]`,
		reflect.ValueOf(customfield.RawNormalizedDynamicValue(basetypes.NewDynamicNull())),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust([]attr.Type{}, []attr.Value{}))),
	},

	"tfsdk_normalized_dynamic_nested_empty_array_inference": {
		`{"urls":[]}`,
		reflect.ValueOf(customfield.RawNormalizedDynamicValue(basetypes.NewDynamicNull())),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.ObjectValueMust(
			map[string]attr.Type{"urls": types.TupleType{ElemTypes: []attr.Type{}}},
			map[string]attr.Value{"urls": types.TupleValueMust([]attr.Type{}, []attr.Value{})},
		))),
	},

	"tfsdk_normalized_dynamic_array_of_empty_arrays_inference": {
		`[[],[]]`,
		reflect.ValueOf(customfield.RawNormalizedDynamicValue(basetypes.NewDynamicNull())),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(
			types.TupleType{ElemTypes: []attr.Type{}},
			[]attr.Value{
				types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				types.TupleValueMust([]attr.Type{}, []attr.Value{}),
			},
		))),
	},

	"tfsdk_normalized_dynamic_heterogeneous_array_inference": {
		`["hello",42]`,
		reflect.ValueOf(customfield.RawNormalizedDynamicValue(basetypes.NewDynamicNull())),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust(
			[]attr.Type{types.StringType, types.Int64Type},
			[]attr.Value{types.StringValue("hello"), types.Int64Value(42)},
		))),
	},

	"tfsdk_dynamic_homogeneous_array_inference": {
		`["hello","world"]`,
		reflect.ValueOf(types.DynamicNull()),
		reflect.ValueOf(types.DynamicValue(types.ListValueMust(
			types.StringType,
			[]attr.Value{types.StringValue("hello"), types.StringValue("world")},
		))),
	},

	"tfsdk_normalized_dynamic_homogeneous_array_inference": {
		`["hello","world"]`,
		reflect.ValueOf(customfield.RawNormalizedDynamicValue(basetypes.NewDynamicNull())),
		reflect.ValueOf(customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(
			types.StringType,
			[]attr.Value{types.StringValue("hello"), types.StringValue("world")},
		))),
	},

	"tfsdk_struct_populates_unknown_to_null_if_missing": {
		`{"embedded_string":"some_string","data_object":{}}`,
		reflect.ValueOf(EmbeddedTfsdkStruct{
			EmbeddedString: types.StringUnknown(),
			EmbeddedInt:    types.Int64Unknown(),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		}),
		reflect.ValueOf(EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("some_string"),
			EmbeddedInt:    types.Int64Null(),
			DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
				NestedInt: types.Int64Null(),
			}),
		}),
	},

	"tfsdk_struct_overwrites_from_json": {
		`{"embedded_string":"new_value"}`,
		reflect.ValueOf(EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		}),
		reflect.ValueOf(EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("new_value"),
			EmbeddedInt:    types.Int64Null(),
			DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
		}),
	},

	"tfsdk_date_time_populates_unknown_to_null_if_missing": {
		`{"date":"2006-01-02"}`,
		reflect.ValueOf(DateTimeCustom{
			DateCustom:     timetypes.NewRFC3339Unknown(),
			DateTimeCustom: timetypes.NewRFC3339Unknown(),
		}),
		reflect.ValueOf(DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: timetypes.NewRFC3339Null(),
		}),
	},
}

func TestDecodeFromValue(t *testing.T) {
	spew.Config.ContinueOnMethod = true
	for name, test := range decodeFromValueTests {
		t.Run(name, func(t *testing.T) {
			starting := reflect.New(test.starting.Type())
			starting.Elem().Set(test.starting)

			expectedErr, errorIsExpected := reflect.TypeAssert[error](test.expected)

			err := unmarshalValue([]byte(test.buf), starting)
			if errorIsExpected {
				if err == nil {
					t.Fatalf(`expected error "%s" but did not error`, expectedErr.Error())
				}
				if err.Error() != expectedErr.Error() {
					t.Fatalf(`expected error "%s" but got "%s"`, expectedErr.Error(), err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("deserialization of %v failed with error %v", test.buf, err)
				}
				if !reflect.DeepEqual(starting.Elem().Interface(), test.expected.Interface()) {
					t.Fatalf("expected '%s' to deserialize to \n%s\nbut got\n%s", test.buf, spew.Sdump(test.expected.Interface()), spew.Sdump(starting.Elem().Interface()))
				}
			}
		})
	}
}

var decodeUnsetTests = map[string]struct {
	buf string
	val reflect.Value
}{
	"nested_object_list_is_omitted_null": {
		`{}`,
		reflect.ValueOf(ListWithNestedObj{
			A: customfield.NullObjectList[Embedded2](ctx),
		}),
	},
	"nested_object_list_is_explicit_null": {
		`{"a": null}`,
		reflect.ValueOf(ListWithNestedObj{
			A: customfield.NullObjectList[Embedded2](ctx),
		}),
	},
	"nested_object_list_is_empty": {
		`{"a": []}`,
		reflect.ValueOf(ListWithNestedObj{
			A: customfield.NewObjectListMust(ctx, []Embedded2{}),
		}),
	},
}

func TestDecodeUnsetBehaviour(t *testing.T) {
	spew.Config.SortKeys = true
	for name, test := range merge(decodeUnsetTests) {
		t.Run(name, func(t *testing.T) {
			resultValue := reflect.New(test.val.Type())
			d := &decoderBuilder{
				dateFormat:            time.RFC3339,
				unmarshalComputedOnly: false,
				updateBehavior:        IfUnset,
			}
			if err := d.unmarshal([]byte(test.buf), resultValue); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", resultValue, err)
			}
			if !reflect.DeepEqual(resultValue.Elem().Interface(), test.val.Interface()) {
				t.Fatalf("incorrect deserialization for '%s':\nexpected:\n%s\nactual:\n%s\n", test.buf, spew.Sdump(test.val.Interface()), spew.Sdump(resultValue.Elem().Interface()))
			}
		})
	}
}

type StructWithComputedFields struct {
	RegStr            types.String                                             `tfsdk:"str" json:"str" api:"optional"`
	DerivedStr        types.String                                             `tfsdk:"str" json:"derived_str" api:"optional"`
	CompStr           types.String                                             `tfsdk:"comp_str" json:"comp_str" api:"computed"`
	CompOptStr        types.String                                             `tfsdk:"opt_str" json:"opt_str" api:"computed_optional"`
	CompTime          timetypes.RFC3339                                        `tfsdk:"time" json:"time" api:"computed"`
	CompOptTime       timetypes.RFC3339                                        `tfsdk:"opt_time" json:"opt_time" api:"computed_optional"`
	Nested            NestedStructWithComputedFields                           `tfsdk:"nested" json:"nested" api:"optional"`
	NestedCust        customfield.NestedObject[NestedStructWithComputedFields] `tfsdk:"nested_obj" json:"nested_obj" api:"optional"`
	CompOptNestedCust customfield.NestedObject[NestedStructWithComputedFields] `tfsdk:"opt_nested_obj" json:"opt_nested_obj" api:"computed_optional"`
	NestedList        *[]*NestedStructWithComputedFields                       `tfsdk:"nested_list" json:"nested_list" api:"optional"`
	MapCust           customfield.Map[customfield.List[types.String]]          `tfsdk:"map_cust" json:"map_cust" api:"optional"`
	MapRegular        *map[string][]*NestedStructWithComputedFields            `tfsdk:"map_regular" json:"map_regular" api:"optional"`
	CompMap           *map[string]*NestedStructWithComputedFields              `tfsdk:"comp_map" json:"comp_map" api:"computed"`
	CompMapList       *map[string][]*NestedStructWithComputedFields            `tfsdk:"comp_map_list" json:"comp_map_list" api:"computed"`
}

type NestedStructWithComputedFields struct {
	RegStr     types.String `tfsdk:"nested_str" json:"nested_str" api:"required"`
	CompStr    types.String `tfsdk:"nested_comp_str" json:"nested_comp_str" api:"computed"`
	CompOptInt types.Int64  `tfsdk:"nested_comp_opt_int" json:"nested_comp_opt_int" api:"computed_optional"`
}

var exampleNestedJSON = `{
	"str":"str",
	"derived_str":"str",
	"comp_str":"comp_str",
	"opt_str":"opt_str",
	"time":"2006-01-02T15:04:05Z",
	"opt_time":"2006-01-02T15:04:05Z",
	"nested":{"nested_str":"nested_str","nested_comp_str":"nested_comp_str","nested_comp_opt_int":42},
	"nested_obj":{"nested_str":"nested_str","nested_comp_str":"nested_comp_str","nested_comp_opt_int":42},
	"opt_nested_obj":{"nested_str":"nested_str","nested_comp_str":"nested_comp_str","nested_comp_opt_int":42},
	"nested_list":[{"nested_str":"nested_str","nested_comp_str":"list_nested_comp_str_1","nested_comp_opt_int":43},{"nested_str":"nested_str","nested_comp_str":"list_nested_comp_str_2","nested_comp_opt_int":44}],
	"map_cust":{"key":["val1","val2"]},
	"map_regular":{"key":[{"nested_str":"nested_str","nested_comp_str":"nested_comp_str","nested_comp_opt_int":42}]},
	"comp_map":{"comp_key":{"nested_comp_str":"nested_comp_str"}},
	"comp_map_list":{"comp_list_key":[{"nested_comp_str":"nested_comp_str"}]}
}`

type nestedMapExample struct {
	SomeStruct customfield.NestedObject[nestedMapStruct] `tfsdk:"some_struct" json:"some_struct" api:"computed_optional"`
}

type nestedMapStruct struct {
	NestedMap map[string]types.Float64 `tfsdk:"nested_map" json:"nested_map" api:"optional"`
}

type computedMapStruct struct {
	ComputedMap map[string]types.Float64 `tfsdk:"computed_map" json:"computed_map" api:"computed_optional"`
}

type primitiveListExample struct {
	StrList customfield.List[types.String] `tfsdk:"str_list" json:"str_list" api:"computed_optional"`
}

type primitiveSetExample struct {
	StrSet customfield.Set[types.String] `tfsdk:"str_set" json:"str_set" api:"computed_optional"`
}

type primitiveMapExample struct {
	StrMap customfield.Map[types.String] `tfsdk:"str_map" json:"str_map" api:"computed_optional"`
}

type nestedObjectExample struct {
	NestedObj customfield.NestedObject[computedMapStruct] `tfsdk:"nested_obj" json:"nested_obj" api:"computed_optional"`
}

type nestedObjectListExample struct {
	ObjList customfield.NestedObjectList[computedMapStruct] `tfsdk:"obj_list" json:"obj_list" api:"computed_optional"`
}

type nestedObjectSetExample struct {
	ObjSet customfield.NestedObjectSet[computedMapStruct] `tfsdk:"obj_set" json:"obj_set" api:"computed_optional"`
}

type listOfListsExample struct {
	ListOfLists customfield.List[customfield.List[types.String]] `tfsdk:"list_of_lists" json:"list_of_lists" api:"computed_optional"`
}

type setOfSetsExample struct {
	SetOfSets customfield.Set[customfield.Set[types.Int64]] `tfsdk:"set_of_sets" json:"set_of_sets" api:"computed_optional"`
}

type mapWithListExample struct {
	MapWithList customfield.Map[customfield.List[types.String]] `tfsdk:"map_with_list" json:"map_with_list" api:"computed_optional"`
}

type tupleExample struct {
	TupleVal types.Tuple `tfsdk:"tuple_val" json:"tuple_val" api:"computed_optional"`
}

type ComputedOptionalCollectionsExample struct {
	Name        types.String                              `tfsdk:"name" json:"name" api:"required"`
	Tags        customfield.List[types.String]            `tfsdk:"tags" json:"tags" api:"computed_optional"`
	Metadata    customfield.Map[types.String]             `tfsdk:"metadata" json:"metadata" api:"computed_optional"`
	Ports       customfield.Set[types.Int64]              `tfsdk:"ports" json:"ports" api:"computed_optional"`
	Coordinates types.Tuple                               `tfsdk:"coordinates" json:"coordinates" api:"computed_optional"`
	Rules       customfield.NestedObjectList[RuleExample] `tfsdk:"rules" json:"rules" api:"computed_optional"`
	Status      types.String                              `tfsdk:"status" json:"status" api:"computed_optional"`
}

type RuleExample struct {
	Priority types.Int64  `tfsdk:"priority" json:"priority" api:"required"`
	Action   types.String `tfsdk:"action" json:"action" api:"computed"`
}

var decodeComputedOnlyTests = map[string]struct {
	buf      string
	starting reflect.Value
	expected reflect.Value
}{
	"primitive_list_unchanged": {
		`{}`,
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("a"), types.StringValue("b"), types.StringValue("c")}),
		}),
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("a"), types.StringValue("b"), types.StringValue("c")}),
		}),
	},
	"computed_optional_list_from_null": {
		`{"str_list":["hello","world","test"]}`,
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NullList[types.String](ctx),
		}),
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("hello"), types.StringValue("world"), types.StringValue("test")}),
		}),
	},
	"computed_optional_set_from_null": {
		`{"str_set":["apple","banana","cherry"]}`,
		reflect.ValueOf(primitiveSetExample{
			StrSet: customfield.NullSet[types.String](ctx),
		}),
		reflect.ValueOf(primitiveSetExample{
			StrSet: customfield.NewSetMust[types.String](ctx, []attr.Value{types.StringValue("apple"), types.StringValue("banana"), types.StringValue("cherry")}),
		}),
	},
	"computed_optional_nested_object_from_null": {
		`{"nested_obj":{"computed_map":{"key1":1.5,"key2":2.5}}}`,
		reflect.ValueOf(nestedObjectExample{
			NestedObj: customfield.NullObject[computedMapStruct](ctx),
		}),
		reflect.ValueOf(nestedObjectExample{
			NestedObj: customfield.NewObjectMust(ctx, &computedMapStruct{
				ComputedMap: map[string]types.Float64{"key1": type_helpers.NewFloat64ValueFromStringUnsafe("1.5"), "key2": type_helpers.NewFloat64ValueFromStringUnsafe("2.5")},
			}),
		}),
	},
	"computed_optional_nested_object_list_from_null": {
		`{"obj_list":[{"computed_map":{"a":1.0}},{"computed_map":{"b":2.0}}]}`,
		reflect.ValueOf(nestedObjectListExample{
			ObjList: customfield.NullObjectList[computedMapStruct](ctx),
		}),
		reflect.ValueOf(nestedObjectListExample{
			ObjList: customfield.NewObjectListMust(ctx, []computedMapStruct{
				{ComputedMap: map[string]types.Float64{"a": type_helpers.NewFloat64ValueFromStringUnsafe("1.0")}},
				{ComputedMap: map[string]types.Float64{"b": type_helpers.NewFloat64ValueFromStringUnsafe("2.0")}},
			}),
		}),
	},
	"computed_optional_nested_object_set_from_null": {
		`{"obj_set":[{"computed_map":{"x":10.0}},{"computed_map":{"y":20.0}}]}`,
		reflect.ValueOf(nestedObjectSetExample{
			ObjSet: customfield.NullObjectSet[computedMapStruct](ctx),
		}),
		reflect.ValueOf(nestedObjectSetExample{
			ObjSet: customfield.NewObjectSetMust(ctx, []computedMapStruct{
				{ComputedMap: map[string]types.Float64{"x": type_helpers.NewFloat64ValueFromStringUnsafe("10.0")}},
				{ComputedMap: map[string]types.Float64{"y": type_helpers.NewFloat64ValueFromStringUnsafe("20.0")}},
			}),
		}),
	},
	"computed_optional_tuple_from_null": {
		`{"tuple_val":[42,"hello",true]}`,
		reflect.ValueOf(tupleExample{
			TupleVal: types.TupleNull([]attr.Type{types.Int64Type, types.StringType, types.BoolType}),
		}),
		reflect.ValueOf(tupleExample{
			TupleVal: types.TupleValueMust(
				[]attr.Type{types.Int64Type, types.StringType, types.BoolType},
				[]attr.Value{types.Int64Value(42), types.StringValue("hello"), types.BoolValue(true)},
			),
		}),
	},
	"computed_optional_list_preserves_existing": {
		`{"str_list":["new1","new2"]}`,
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("existing1"), types.StringValue("existing2")}),
		}),
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("existing1"), types.StringValue("existing2")}),
		}),
	},
	"computed_optional_set_preserves_existing": {
		`{"str_set":["new1","new2"]}`,
		reflect.ValueOf(primitiveSetExample{
			StrSet: customfield.NewSetMust[types.String](ctx, []attr.Value{types.StringValue("existing1"), types.StringValue("existing2")}),
		}),
		reflect.ValueOf(primitiveSetExample{
			StrSet: customfield.NewSetMust[types.String](ctx, []attr.Value{types.StringValue("existing1"), types.StringValue("existing2")}),
		}),
	},
	"computed_optional_nested_object_preserves_existing": {
		`{"nested_obj":{"computed_map":{"new":999.0}}}`,
		reflect.ValueOf(nestedObjectExample{
			NestedObj: customfield.NewObjectMust(ctx, &computedMapStruct{
				ComputedMap: map[string]types.Float64{"existing": type_helpers.NewFloat64ValueFromStringUnsafe("123.0")},
			}),
		}),
		reflect.ValueOf(nestedObjectExample{
			NestedObj: customfield.NewObjectMust(ctx, &computedMapStruct{
				ComputedMap: map[string]types.Float64{"existing": type_helpers.NewFloat64ValueFromStringUnsafe("123.0")},
			}),
		}),
	},
	"computed_optional_nested_object_list_preserves_existing": {
		`{"obj_list":[{"computed_map":{"new":999.0}}]}`,
		reflect.ValueOf(nestedObjectListExample{
			ObjList: customfield.NewObjectListMust(ctx, []computedMapStruct{
				{ComputedMap: map[string]types.Float64{"existing": type_helpers.NewFloat64ValueFromStringUnsafe("456.0")}},
			}),
		}),
		reflect.ValueOf(nestedObjectListExample{
			ObjList: customfield.NewObjectListMust(ctx, []computedMapStruct{
				{ComputedMap: map[string]types.Float64{"existing": type_helpers.NewFloat64ValueFromStringUnsafe("456.0")}},
			}),
		}),
	},
	"computed_optional_nested_object_set_preserves_existing": {
		`{"obj_set":[{"computed_map":{"new":999.0}}]}`,
		reflect.ValueOf(nestedObjectSetExample{
			ObjSet: customfield.NewObjectSetMust(ctx, []computedMapStruct{
				{ComputedMap: map[string]types.Float64{"existing": type_helpers.NewFloat64ValueFromStringUnsafe("789.0")}},
			}),
		}),
		reflect.ValueOf(nestedObjectSetExample{
			ObjSet: customfield.NewObjectSetMust(ctx, []computedMapStruct{
				{ComputedMap: map[string]types.Float64{"existing": type_helpers.NewFloat64ValueFromStringUnsafe("789.0")}},
			}),
		}),
	},
	"computed_optional_tuple_preserves_existing": {
		`{"tuple_val":[999,"new",false]}`,
		reflect.ValueOf(tupleExample{
			TupleVal: types.TupleValueMust(
				[]attr.Type{types.Int64Type, types.StringType, types.BoolType},
				[]attr.Value{types.Int64Value(123), types.StringValue("existing"), types.BoolValue(true)},
			),
		}),
		reflect.ValueOf(tupleExample{
			TupleVal: types.TupleValueMust(
				[]attr.Type{types.Int64Type, types.StringType, types.BoolType},
				[]attr.Value{types.Int64Value(123), types.StringValue("existing"), types.BoolValue(true)},
			),
		}),
	},
	"computed_optional_list_from_unknown": {
		`{"str_list":["alpha","beta","gamma"]}`,
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.UnknownList[types.String](ctx),
		}),
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("alpha"), types.StringValue("beta"), types.StringValue("gamma")}),
		}),
	},
	"computed_optional_set_from_unknown": {
		`{"str_set":["one","two","three"]}`,
		reflect.ValueOf(primitiveSetExample{
			StrSet: customfield.UnknownSet[types.String](ctx),
		}),
		reflect.ValueOf(primitiveSetExample{
			StrSet: customfield.NewSetMust[types.String](ctx, []attr.Value{types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
		}),
	},
	"computed_optional_map_from_unknown": {
		`{"str_map":{"key1":"value1","key2":"value2","key3":"value3"}}`,
		reflect.ValueOf(primitiveMapExample{
			StrMap: customfield.UnknownMap[types.String](ctx),
		}),
		reflect.ValueOf(primitiveMapExample{
			StrMap: customfield.NewMapMust[types.String](ctx, map[string]types.String{
				"key1": types.StringValue("value1"),
				"key2": types.StringValue("value2"),
				"key3": types.StringValue("value3"),
			}),
		}),
	},
	"computed_optional_nested_object_from_unknown": {
		`{"nested_obj":{"computed_map":{"unknown1":3.5,"unknown2":4.5}}}`,
		reflect.ValueOf(nestedObjectExample{
			NestedObj: customfield.UnknownObject[computedMapStruct](ctx),
		}),
		reflect.ValueOf(nestedObjectExample{
			NestedObj: customfield.NewObjectMust(ctx, &computedMapStruct{
				ComputedMap: map[string]types.Float64{"unknown1": type_helpers.NewFloat64ValueFromStringUnsafe("3.5"), "unknown2": type_helpers.NewFloat64ValueFromStringUnsafe("4.5")},
			}),
		}),
	},
	"computed_optional_nested_object_list_from_unknown": {
		`{"obj_list":[{"computed_map":{"u1":5.0}},{"computed_map":{"u2":6.0}}]}`,
		reflect.ValueOf(nestedObjectListExample{
			ObjList: customfield.UnknownObjectList[computedMapStruct](ctx),
		}),
		reflect.ValueOf(nestedObjectListExample{
			ObjList: customfield.NewObjectListMust(ctx, []computedMapStruct{
				{ComputedMap: map[string]types.Float64{"u1": type_helpers.NewFloat64ValueFromStringUnsafe("5.0")}},
				{ComputedMap: map[string]types.Float64{"u2": type_helpers.NewFloat64ValueFromStringUnsafe("6.0")}},
			}),
		}),
	},
	"computed_optional_nested_object_set_from_unknown": {
		`{"obj_set":[{"computed_map":{"ux":30.0}},{"computed_map":{"uy":40.0}}]}`,
		reflect.ValueOf(nestedObjectSetExample{
			ObjSet: customfield.UnknownObjectSet[computedMapStruct](ctx),
		}),
		reflect.ValueOf(nestedObjectSetExample{
			ObjSet: customfield.NewObjectSetMust(ctx, []computedMapStruct{
				{ComputedMap: map[string]types.Float64{"ux": type_helpers.NewFloat64ValueFromStringUnsafe("30.0")}},
				{ComputedMap: map[string]types.Float64{"uy": type_helpers.NewFloat64ValueFromStringUnsafe("40.0")}},
			}),
		}),
	},
	"computed_optional_tuple_from_unknown": {
		`{"tuple_val":[99,"unknown",false]}`,
		reflect.ValueOf(tupleExample{
			TupleVal: types.TupleUnknown([]attr.Type{types.Int64Type, types.StringType, types.BoolType}),
		}),
		reflect.ValueOf(tupleExample{
			TupleVal: types.TupleValueMust(
				[]attr.Type{types.Int64Type, types.StringType, types.BoolType},
				[]attr.Value{types.Int64Value(99), types.StringValue("unknown"), types.BoolValue(false)},
			),
		}),
	},
	"computed_optional_list_of_lists_from_unknown": {
		`{"list_of_lists":[["a","b"],["c","d","e"],["f"]]}`,
		reflect.ValueOf(listOfListsExample{
			ListOfLists: customfield.UnknownList[customfield.List[types.String]](ctx),
		}),
		reflect.ValueOf(listOfListsExample{
			ListOfLists: customfield.NewListMust[customfield.List[types.String]](ctx, []attr.Value{
				customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("a"), types.StringValue("b")}),
				customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("c"), types.StringValue("d"), types.StringValue("e")}),
				customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("f")}),
			}),
		}),
	},
	"computed_optional_set_of_sets_from_unknown": {
		`{"set_of_sets":[[1,2],[3,4,5],[6]]}`,
		reflect.ValueOf(setOfSetsExample{
			SetOfSets: customfield.UnknownSet[customfield.Set[types.Int64]](ctx),
		}),
		reflect.ValueOf(setOfSetsExample{
			SetOfSets: customfield.NewSetMust[customfield.Set[types.Int64]](ctx, []attr.Value{
				customfield.NewSetMust[types.Int64](ctx, []attr.Value{types.Int64Value(1), types.Int64Value(2)}),
				customfield.NewSetMust[types.Int64](ctx, []attr.Value{types.Int64Value(3), types.Int64Value(4), types.Int64Value(5)}),
				customfield.NewSetMust[types.Int64](ctx, []attr.Value{types.Int64Value(6)}),
			}),
		}),
	},
	"computed_optional_map_with_list_from_unknown": {
		`{"map_with_list":{"key1":["val1","val2"],"key2":["val3","val4","val5"],"key3":["val6"]}}`,
		reflect.ValueOf(mapWithListExample{
			MapWithList: customfield.UnknownMap[customfield.List[types.String]](ctx),
		}),
		reflect.ValueOf(mapWithListExample{
			MapWithList: customfield.NewMapMust[customfield.List[types.String]](ctx, map[string]customfield.List[types.String]{
				"key1": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("val1"), types.StringValue("val2")}),
				"key2": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("val3"), types.StringValue("val4"), types.StringValue("val5")}),
				"key3": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("val6")}),
			}),
		}),
	},
	"computed_optional_all_collections_from_unknown": {
		`{"name":"test-unknown","tags":["tag-u1","tag-u2","tag-u3"],"metadata":{"env":"prod","tier":"premium"},"ports":[3000,3001,3002],"coordinates":[51.5074,-0.1278],"rules":[{"priority":10,"action":"deny"},{"priority":20,"action":"allow"}],"status":"pending"}`,
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-unknown"),
			Tags:        customfield.UnknownList[types.String](ctx),
			Metadata:    customfield.UnknownMap[types.String](ctx),
			Ports:       customfield.UnknownSet[types.Int64](ctx),
			Coordinates: types.TupleUnknown([]attr.Type{types.Float64Type, types.Float64Type}),
			Rules:       customfield.UnknownObjectList[RuleExample](ctx),
			Status:      types.StringUnknown(),
		}),
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-unknown"),
			Tags:        customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("tag-u1"), types.StringValue("tag-u2"), types.StringValue("tag-u3")}),
			Metadata:    customfield.NewMapMust[types.String](ctx, map[string]types.String{"env": types.StringValue("prod"), "tier": types.StringValue("premium")}),
			Ports:       customfield.NewSetMust[types.Int64](ctx, []attr.Value{types.Int64Value(3000), types.Int64Value(3001), types.Int64Value(3002)}),
			Coordinates: types.TupleValueMust([]attr.Type{types.Float64Type, types.Float64Type}, []attr.Value{type_helpers.NewFloat64ValueFromStringUnsafe("51.5074"), type_helpers.NewFloat64ValueFromStringUnsafe("-0.1278")}),
			Rules: customfield.NewObjectListMust(ctx, []RuleExample{
				{
					Priority: types.Int64Value(10),
					Action:   types.StringValue("deny"),
				},
				{
					Priority: types.Int64Value(20),
					Action:   types.StringValue("allow"),
				},
			}),
			Status: types.StringValue("pending"),
		}),
	},
	"computed_optional_mixed_unknown_configured_null": {
		`{"name":"test-mixed","tags":["api-tag1","api-tag2"],"metadata":{"source":"api"},"ports":[9000,9001],"coordinates":[35.6762,139.6503],"rules":[{"priority":5,"action":"log"}],"status":"processing"}`,
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-mixed"),
			Tags:        customfield.UnknownList[types.String](ctx),
			Metadata:    customfield.NewMapMust[types.String](ctx, map[string]types.String{"user": types.StringValue("admin")}),
			Ports:       customfield.UnknownSet[types.Int64](ctx),
			Coordinates: types.TupleValueMust([]attr.Type{types.Float64Type, types.Float64Type}, []attr.Value{type_helpers.NewFloat64Value(0), type_helpers.NewFloat64Value(0)}),
			Rules:       customfield.NullObjectList[RuleExample](ctx),
			Status:      types.StringUnknown(),
		}),
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-mixed"),
			Tags:        customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("api-tag1"), types.StringValue("api-tag2")}),
			Metadata:    customfield.NewMapMust[types.String](ctx, map[string]types.String{"user": types.StringValue("admin")}),
			Ports:       customfield.NewSetMust[types.Int64](ctx, []attr.Value{types.Int64Value(9000), types.Int64Value(9001)}),
			Coordinates: types.TupleValueMust([]attr.Type{types.Float64Type, types.Float64Type}, []attr.Value{type_helpers.NewFloat64Value(0), type_helpers.NewFloat64Value(0)}),
			Rules: customfield.NewObjectListMust(ctx, []RuleExample{
				{
					Priority: types.Int64Value(5),
					Action:   types.StringValue("log"),
				},
			}),
			Status: types.StringValue("processing"),
		}),
	},
	"computed_optional_collections_bug_list": {
		`{"name":"test-instance","tags":["auto-tag1","auto-tag2"],"metadata":{"created_by":"system","version":"1.0"},"ports":[8080,8443],"coordinates":[40.7128,-74.006],"rules":[{"priority":1,"action":"allow"}],"status":"active"}`,
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-instance"),
			Tags:        customfield.NullList[types.String](ctx),
			Metadata:    customfield.NullMap[types.String](ctx),
			Ports:       customfield.NullSet[types.Int64](ctx),
			Coordinates: types.TupleNull([]attr.Type{types.Float64Type, types.Float64Type}),
			Rules:       customfield.NullObjectList[RuleExample](ctx),
			Status:      types.StringNull(),
		}),
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-instance"),
			Tags:        customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("auto-tag1"), types.StringValue("auto-tag2")}),
			Metadata:    customfield.NewMapMust[types.String](ctx, map[string]types.String{"created_by": types.StringValue("system"), "version": types.StringValue("1.0")}),
			Ports:       customfield.NewSetMust[types.Int64](ctx, []attr.Value{types.Int64Value(8080), types.Int64Value(8443)}),
			Coordinates: types.TupleValueMust([]attr.Type{types.Float64Type, types.Float64Type}, []attr.Value{type_helpers.NewFloat64ValueFromStringUnsafe("40.7128"), type_helpers.NewFloat64ValueFromStringUnsafe("-74.006")}),
			Rules: customfield.NewObjectListMust(ctx, []RuleExample{
				{
					Priority: types.Int64Value(1),
					Action:   types.StringValue("allow"),
				},
			}),
			Status: types.StringValue("active"),
		}),
	},
	"computed_optional_collections_mixed": {
		`{"name":"test-instance","tags":["user-tag"],"metadata":{"created_by":"system","version":"1.0"},"ports":[8080,8443],"coordinates":[40.7128,-74.006],"rules":[{"priority":1,"action":"allow"}],"status":"active"}`,
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-instance"),
			Tags:        customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("user-tag")}),
			Metadata:    customfield.NullMap[types.String](ctx),
			Ports:       customfield.NullSet[types.Int64](ctx),
			Coordinates: types.TupleNull([]attr.Type{types.Float64Type, types.Float64Type}),
			Rules:       customfield.NullObjectList[RuleExample](ctx),
			Status:      types.StringNull(),
		}),
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-instance"),
			Tags:        customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("user-tag")}),
			Metadata:    customfield.NewMapMust[types.String](ctx, map[string]types.String{"created_by": types.StringValue("system"), "version": types.StringValue("1.0")}),
			Ports:       customfield.NewSetMust[types.Int64](ctx, []attr.Value{types.Int64Value(8080), types.Int64Value(8443)}),
			Coordinates: types.TupleValueMust([]attr.Type{types.Float64Type, types.Float64Type}, []attr.Value{type_helpers.NewFloat64ValueFromStringUnsafe("40.7128"), type_helpers.NewFloat64ValueFromStringUnsafe("-74.006")}),
			Rules: customfield.NewObjectListMust(ctx, []RuleExample{
				{
					Priority: types.Int64Value(1),
					Action:   types.StringValue("allow"),
				},
			}),
			Status: types.StringValue("active"),
		}),
	},
	"computed_optional_empty_list_preserved": {
		`{"str_list":["from","api","response"]}`,
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{}),
		}),
		reflect.ValueOf(primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{}),
		}),
	},
	"computed_optional_empty_set_preserved": {
		`{"str_set":["api","values","here"]}`,
		reflect.ValueOf(primitiveSetExample{
			StrSet: customfield.NewSetMust[types.String](ctx, []attr.Value{}),
		}),
		reflect.ValueOf(primitiveSetExample{
			StrSet: customfield.NewSetMust[types.String](ctx, []attr.Value{}),
		}),
	},
	"computed_optional_empty_map_preserved": {
		`{"str_map":{"api_key":"api_value","another":"value"}}`,
		reflect.ValueOf(primitiveMapExample{
			StrMap: customfield.NewMapMust[types.String](ctx, map[string]types.String{}),
		}),
		reflect.ValueOf(primitiveMapExample{
			StrMap: customfield.NewMapMust[types.String](ctx, map[string]types.String{}),
		}),
	},
	"computed_optional_empty_nested_object_list_preserved": {
		`{"obj_list":[{"computed_map":{"from_api":1.0}},{"computed_map":{"also_api":2.0}}]}`,
		reflect.ValueOf(nestedObjectListExample{
			ObjList: customfield.NewObjectListMust(ctx, []computedMapStruct{}),
		}),
		reflect.ValueOf(nestedObjectListExample{
			ObjList: customfield.NewObjectListMust(ctx, []computedMapStruct{}),
		}),
	},
	"computed_optional_empty_nested_object_set_preserved": {
		`{"obj_set":[{"computed_map":{"api_x":10.0}},{"computed_map":{"api_y":20.0}}]}`,
		reflect.ValueOf(nestedObjectSetExample{
			ObjSet: customfield.NewObjectSetMust(ctx, []computedMapStruct{}),
		}),
		reflect.ValueOf(nestedObjectSetExample{
			ObjSet: customfield.NewObjectSetMust(ctx, []computedMapStruct{}),
		}),
	},
	"computed_optional_empty_nested_list_of_lists_preserved": {
		`{"list_of_lists":[["api1","api2"],["api3","api4","api5"],["api6"]]}`,
		reflect.ValueOf(listOfListsExample{
			ListOfLists: customfield.NewListMust[customfield.List[types.String]](ctx, []attr.Value{}),
		}),
		reflect.ValueOf(listOfListsExample{
			ListOfLists: customfield.NewListMust[customfield.List[types.String]](ctx, []attr.Value{}),
		}),
	},
	"computed_optional_empty_nested_set_of_sets_preserved": {
		`{"set_of_sets":[[10,20],[30,40,50],[60]]}`,
		reflect.ValueOf(setOfSetsExample{
			SetOfSets: customfield.NewSetMust[customfield.Set[types.Int64]](ctx, []attr.Value{}),
		}),
		reflect.ValueOf(setOfSetsExample{
			SetOfSets: customfield.NewSetMust[customfield.Set[types.Int64]](ctx, []attr.Value{}),
		}),
	},
	"computed_optional_empty_map_with_list_preserved": {
		`{"map_with_list":{"api_key1":["api_val1","api_val2"],"api_key2":["api_val3","api_val4","api_val5"]}}`,
		reflect.ValueOf(mapWithListExample{
			MapWithList: customfield.NewMapMust[customfield.List[types.String]](ctx, map[string]customfield.List[types.String]{}),
		}),
		reflect.ValueOf(mapWithListExample{
			MapWithList: customfield.NewMapMust[customfield.List[types.String]](ctx, map[string]customfield.List[types.String]{}),
		}),
	},
	"computed_optional_all_empty_collections_preserved": {
		`{"name":"test-empty","tags":["api-tag1","api-tag2","api-tag3"],"metadata":{"api":"data","more":"stuff"},"ports":[4000,4001,4002],"coordinates":[12.34,56.78],"rules":[{"priority":100,"action":"block"},{"priority":200,"action":"permit"}],"status":"running"}`,
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-empty"),
			Tags:        customfield.NewListMust[types.String](ctx, []attr.Value{}),
			Metadata:    customfield.NewMapMust[types.String](ctx, map[string]types.String{}),
			Ports:       customfield.NewSetMust[types.Int64](ctx, []attr.Value{}),
			Coordinates: types.TupleValueMust([]attr.Type{types.Float64Type, types.Float64Type}, []attr.Value{type_helpers.NewFloat64Value(0), type_helpers.NewFloat64Value(0)}),
			Rules:       customfield.NewObjectListMust(ctx, []RuleExample{}),
			Status:      types.StringValue("configured"),
		}),
		reflect.ValueOf(ComputedOptionalCollectionsExample{
			Name:        types.StringValue("test-empty"),
			Tags:        customfield.NewListMust[types.String](ctx, []attr.Value{}),
			Metadata:    customfield.NewMapMust[types.String](ctx, map[string]types.String{}),
			Ports:       customfield.NewSetMust[types.Int64](ctx, []attr.Value{}),
			Coordinates: types.TupleValueMust([]attr.Type{types.Float64Type, types.Float64Type}, []attr.Value{type_helpers.NewFloat64Value(0), type_helpers.NewFloat64Value(0)}),
			Rules:       customfield.NewObjectListMust(ctx, []RuleExample{}),
			Status:      types.StringValue("configured"),
		}),
	},
	"nested_map_unchanged": {
		`{"some_struct": {"nested_map":{"example_key":3.14}}}`,
		reflect.ValueOf(nestedMapExample{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStruct{
				NestedMap: map[string]types.Float64{"example_key": type_helpers.NewFloat64ValueFromStringUnsafe("3.14")},
			}),
		}),
		reflect.ValueOf(nestedMapExample{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStruct{
				NestedMap: map[string]types.Float64{"example_key": type_helpers.NewFloat64ValueFromStringUnsafe("3.14")},
			}),
		}),
	},
	"nested_optional_map_avoids_updates": {
		`{"some_struct": {"nested_map":{"example_key":0.123,"new_key":456.7}}}`,
		reflect.ValueOf(nestedMapExample{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStruct{
				NestedMap: map[string]types.Float64{"example_key": type_helpers.NewFloat64ValueFromStringUnsafe("3.14")},
			}),
		}),
		reflect.ValueOf(nestedMapExample{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStruct{
				NestedMap: map[string]types.Float64{"example_key": type_helpers.NewFloat64ValueFromStringUnsafe("3.14")},
			}),
		}),
	},
	"only_updates_computed_props": {
		exampleNestedJSON,
		reflect.ValueOf(StructWithComputedFields{
			RegStr:      types.StringNull(),
			DerivedStr:  types.StringNull(),
			CompStr:     types.StringNull(),
			CompOptStr:  types.StringNull(),
			CompTime:    timetypes.NewRFC3339Null(),
			CompOptTime: timetypes.NewRFC3339Null(),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringNull(),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Null(),
			},
			NestedCust:        customfield.NullObject[NestedStructWithComputedFields](ctx),
			CompOptNestedCust: customfield.NullObject[NestedStructWithComputedFields](ctx),
			MapRegular: P(map[string][]*NestedStructWithComputedFields{"key": {
				&NestedStructWithComputedFields{
					RegStr:     types.StringNull(),
					CompStr:    types.StringNull(),
					CompOptInt: types.Int64Null(),
				},
			}}),
		}),
		reflect.ValueOf(StructWithComputedFields{
			RegStr:      types.StringNull(),
			DerivedStr:  types.StringNull(),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			},
			NestedCust: customfield.NullObject[NestedStructWithComputedFields](ctx),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			}),
			MapRegular: P(map[string][]*NestedStructWithComputedFields{"key": {
				&NestedStructWithComputedFields{
					RegStr:     types.StringNull(),
					CompStr:    types.StringValue("nested_comp_str"),
					CompOptInt: types.Int64Value(42),
				},
			}}),
			CompMap: P(map[string]*NestedStructWithComputedFields{"comp_key": {
				CompStr: types.StringValue("nested_comp_str"),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFields{"comp_list_key": {
				&NestedStructWithComputedFields{
					CompStr: types.StringValue("nested_comp_str"),
				},
			}}),
		}),
	},
	"only_updates_computed_props_from_unknown": {
		exampleNestedJSON,
		reflect.ValueOf(StructWithComputedFields{
			RegStr:      types.StringUnknown(),
			DerivedStr:  types.StringUnknown(),
			CompStr:     types.StringUnknown(),
			CompOptStr:  types.StringUnknown(),
			CompTime:    timetypes.NewRFC3339Unknown(),
			CompOptTime: timetypes.NewRFC3339Unknown(),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringUnknown(),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Unknown(),
			},
			NestedCust:        customfield.NullObject[NestedStructWithComputedFields](ctx),
			CompOptNestedCust: customfield.UnknownObject[NestedStructWithComputedFields](ctx),
		}),
		reflect.ValueOf(StructWithComputedFields{
			RegStr:      types.StringUnknown(),
			DerivedStr:  types.StringUnknown(),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringUnknown(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			},
			NestedCust: customfield.NullObject[NestedStructWithComputedFields](ctx),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			}),
			CompMap: P(map[string]*NestedStructWithComputedFields{"comp_key": {
				CompStr: types.StringValue("nested_comp_str"),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFields{"comp_list_key": {
				&NestedStructWithComputedFields{
					CompStr: types.StringValue("nested_comp_str"),
				},
			}}),
		}),
	},

	"doesnt_update_computed_optional_if_set": {
		exampleNestedJSON,
		reflect.ValueOf(StructWithComputedFields{
			RegStr:      types.StringValue("existing_str"),
			DerivedStr:  types.StringValue("existing_str"),
			CompStr:     types.StringValue("existing_comp_str"),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFields{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_1"),
				CompOptInt: types.Int64Value(11),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
		}),
		reflect.ValueOf(StructWithComputedFields{
			RegStr:      types.StringValue("existing_str"),
			DerivedStr:  types.StringValue("existing_str"),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFields{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringValue("list_nested_comp_str_1"),
				CompOptInt: types.Int64Value(11),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
			CompMap: P(map[string]*NestedStructWithComputedFields{"comp_key": {
				CompStr: types.StringValue("nested_comp_str"),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFields{"comp_list_key": {
				&NestedStructWithComputedFields{
					CompStr: types.StringValue("nested_comp_str"),
				},
			}}),
		}),
	},

	"updates_computed_if_JSON_properties_are_missing": {
		`{}`,
		reflect.ValueOf(StructWithComputedFields{
			RegStr:      types.StringValue("existing_str"),
			DerivedStr:  types.StringValue("existing_str"),
			CompStr:     types.StringValue("existing_comp_str"),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFields{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Unknown(),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
			MapCust: customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
				"key": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringUnknown(), types.StringValue("val2")}),
			}),
		}),
		reflect.ValueOf(StructWithComputedFields{
			RegStr:      types.StringValue("existing_str"),
			DerivedStr:  types.StringValue("existing_str"),
			CompStr:     types.StringNull(),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339Null(),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFields{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Null(),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Value(12),
			}},
			MapCust: customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
				"key": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringNull(), types.StringValue("val2")}),
			}),
		}),
	},

	"tfsdk_struct_only_overwrites_computed_from_json": {
		`{"embedded_string":"new_value"}`,
		reflect.ValueOf(EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		}),
		reflect.ValueOf(EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
		}),
	},
}

var testSemanticEquivalence = map[string][]attr.Value{
	"nulls": {
		basetypes.NewBoolNull(),
		basetypes.NewBoolNull(),
		basetypes.NewInt32Null(),
		basetypes.NewMapNull(basetypes.BoolType{}),
		basetypes.NewSetNull(basetypes.StringType{}),
		basetypes.NewListNull(basetypes.NumberType{}),
		basetypes.NewTupleNull([]attr.Type{}),
		basetypes.NewObjectNull(map[string]attr.Type{"hi": basetypes.StringType{}}),
	},
	"unknowns": {
		basetypes.NewBoolUnknown(),
		basetypes.NewBoolUnknown(),
		basetypes.NewInt32Unknown(),
		basetypes.NewMapUnknown(basetypes.BoolType{}),
		basetypes.NewSetUnknown(basetypes.StringType{}),
		basetypes.NewListUnknown(basetypes.NumberType{}),
		basetypes.NewTupleUnknown([]attr.Type{}),
		basetypes.NewObjectUnknown(map[string]attr.Type{"hi": basetypes.StringType{}}),
	},
	"floats": {
		basetypes.NewFloat32Value(12.0),
		basetypes.NewFloat64Value(12.0),
		type_helpers.NewFloat64ValueFromStringUnsafe("12.0"),
		basetypes.NewNumberValue(big.NewFloat(12.0)),
		type_helpers.NewNumberValueFromStringUnsafe("12.0"),
	},
	"ints": {
		basetypes.NewInt32Value(12),
		basetypes.NewInt64Value(12),
		basetypes.NewNumberValue(big.NewFloat(12)),
	},
	"sequences": {
		basetypes.NewSetValueMust(basetypes.DynamicType{}, []attr.Value{
			basetypes.NewDynamicValue(basetypes.NewInt64Value(12)),
		}),
		basetypes.NewListValueMust(basetypes.DynamicType{}, []attr.Value{
			basetypes.NewDynamicValue(basetypes.NewInt32Value(12)),
		}),
		basetypes.NewTupleValueMust([]attr.Type{customfield.NormalizedDynamicType{}}, []attr.Value{
			customfield.RawNormalizedDynamicValueFrom(basetypes.NewInt64Value(12)),
		}),
	},
	"maps": {
		basetypes.NewMapValueMust(basetypes.DynamicType{}, map[string]attr.Value{
			"12": basetypes.NewDynamicValue(basetypes.NewNumberValue(big.NewFloat(12.0))),
			"14": basetypes.NewDynamicValue(basetypes.NewNumberValue(big.NewFloat(14.0))),
		}),
		basetypes.NewObjectValueMust(map[string]attr.Type{"12": basetypes.DynamicType{}, "14": basetypes.DynamicType{}}, map[string]attr.Value{
			"12": basetypes.NewDynamicValue(basetypes.NewInt32Value(12)),
			"14": basetypes.NewDynamicValue(basetypes.NewInt64Value(14)),
		}),
	},
	"nested": {
		basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"inner": basetypes.DynamicType{},
			},
			map[string]attr.Value{
				"inner": basetypes.NewDynamicValue(basetypes.NewListValueMust(basetypes.DynamicType{}, []attr.Value{
					basetypes.NewDynamicValue(basetypes.NewStringValue("hi")),
					basetypes.NewDynamicValue(basetypes.NewStringValue("mom")),
				})),
			},
		),
		basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"inner": basetypes.DynamicType{},
			},
			map[string]attr.Value{
				"inner": basetypes.NewDynamicValue(basetypes.NewListValueMust(customfield.NormalizedDynamicType{}, []attr.Value{
					customfield.RawNormalizedDynamicValueFrom(basetypes.NewStringValue("hi")),
					customfield.RawNormalizedDynamicValueFrom(basetypes.NewStringValue("mom")),
				})),
			},
		),
		basetypes.NewMapValueMust(basetypes.DynamicType{}, map[string]attr.Value{
			"inner": basetypes.NewDynamicValue(basetypes.NewListValueMust(basetypes.DynamicType{}, []attr.Value{
				basetypes.NewDynamicValue(basetypes.NewStringValue("hi")),
				basetypes.NewDynamicValue(basetypes.NewStringValue("mom")),
			})),
		}),
	},
}

func TestDecodeComputedOnly(t *testing.T) {
	spew.Config.ContinueOnMethod = false
	for name, test := range decodeComputedOnlyTests {
		t.Run(name, func(t *testing.T) {
			starting := reflect.New(test.starting.Type())
			starting.Elem().Set(test.starting)

			computed := &decoderBuilder{dateFormat: time.RFC3339, unmarshalComputedOnly: true}
			if err := computed.unmarshal([]byte(test.buf), starting); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", test.buf, err)
			}
			if !reflect.DeepEqual(starting.Elem().Interface(), test.expected.Interface()) {
				t.Fatalf("expected '%s' to deserialize to \n%s\nbut got\n%s", test.buf, spew.Sdump(test.expected.Interface()), spew.Sdump(starting.Elem().Interface()))
			}
		})
	}
}

func TestNoStateBetweenDecoders(t *testing.T) {
	TestDecodeComputedOnly(t)
	TestDecodeFromValue(t)
}

func TestSemanticEquivalence(t *testing.T) {
	ctx := context.TODO()
	for name, values := range testSemanticEquivalence {
		t.Run(name, func(t *testing.T) {
			for i, pair := range pairwise(values) {
				lhs := customfield.RawNormalizedDynamicValueFrom(pair[0])
				rhs := customfield.RawNormalizedDynamicValueFrom(pair[1])

				eq, d := lhs.DynamicSemanticEquals(ctx, rhs)
				if len(d) != 0 {
					t.Fatalf("unexpected Diagnostics: %v", d)
				}
				if !eq {
					t.Fatalf("unexpected inequality index: %d, %v <> %v", i, lhs, rhs)

				}
			}
		})
	}
}

func pairwise[T any](input []T) [][]T {
	pairs := [][]T{}
	if len(input) < 2 {
		return [][]T{input}
	}
	a := input[0]
	for _, b := range input[1:] {
		pairs = append(pairs, []T{a, b})
		a = b
	}
	return pairs
}

func merge[T any](testArray ...map[string]T) map[string]T {
	out := make(map[string]T)
	for _, tests := range testArray {
		for name, t := range tests {
			if _, existing := out[name]; existing {
				fmt.Printf("duplicate test name: %s", name)
			}
			out[name] = t
		}
	}
	return out
}

func formatJSON(jsonString string, out *string) error {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(jsonString), "", "    ")
	if err != nil {
		return err
	}

	*out = prettyJSON.String()
	return nil
}
