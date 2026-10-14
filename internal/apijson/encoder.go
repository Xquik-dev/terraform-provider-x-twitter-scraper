// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"bytes"
	"context"
	stdjson "encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tidwall/sjson"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
)

var explicitJsonNull = []byte("null")
var encoders sync.Map

var EscapeSJSONKey = strings.NewReplacer("\\", "\\\\", "|", "\\|", "#", "\\#", "@", "\\@", "*", "\\*", ".", "\\.", ":", "\\:", "?", "\\?").Replace

type CustomMarshaler interface {
	MarshalJSONWithState(plan any, state any) ([]byte, error)
}

func Marshal(value any) ([]byte, error) {
	e := &encoder{dateFormat: time.RFC3339}
	return e.marshal(value, value)
}

func MarshalForUpdate(plan any, state any) ([]byte, error) {
	e := &encoder{root: true, dateFormat: time.RFC3339}
	return e.marshal(plan, state)
}

func MarshalForPatch(plan any, state any) ([]byte, error) {
	e := &encoder{root: true, dateFormat: time.RFC3339, patch: true}
	return e.marshal(plan, state)
}

func MarshalRoot(value any) ([]byte, error) {
	e := &encoder{root: true, dateFormat: time.RFC3339}
	return e.marshal(value, value)
}

type encoder struct {
	dateFormat string
	root       bool
	patch      bool
}

type encoderFunc func(plan reflect.Value, state reflect.Value) ([]byte, error)

type encoderField struct {
	tag parsedStructTag
	fn  encoderFunc
	idx []int
}

type encoderEntry struct {
	reflect.Type
	dateFormat string
	root       bool
	patch      bool
}

func errorFromDiagnostics(diags diag.Diagnostics) error {
	if diags == nil {
		return nil
	}
	messages := []string{}
	for _, err := range diags {
		messages = append(messages, err.Summary())
		messages = append(messages, err.Detail())
	}
	return errors.New(strings.Join(messages, " "))
}

func (e *encoder) marshal(plan any, state any) ([]byte, error) {
	planVal := reflect.ValueOf(plan)
	stateVal := reflect.ValueOf(state)
	if !planVal.IsValid() {
		return nil, nil
	}
	if !stateVal.IsValid() {
		return nil, nil
	}
	typ := planVal.Type()
	enc := e.typeEncoder(typ)
	return enc(planVal, stateVal)
}

func (e *encoder) typeEncoder(t reflect.Type) encoderFunc {
	entry := encoderEntry{
		Type:       t,
		dateFormat: e.dateFormat,
		root:       e.root,
		patch:      e.patch,
	}

	if fi, ok := encoders.Load(entry); ok {
		return fi.(encoderFunc)
	}

	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoders.LoadOrStore(entry, encoderFunc(func(state reflect.Value, plan reflect.Value) ([]byte, error) {
		wg.Wait()
		return f(state, plan)
	}))
	if loaded {
		return fi.(encoderFunc)
	}

	f = e.newTypeEncoder(t)
	wg.Done()
	encoders.Store(entry, f)
	return f
}

func (e *encoder) newTypeEncoder(t reflect.Type) encoderFunc {
	isRoot := e.root

	customMarshalerType := reflect.TypeOf((*CustomMarshaler)(nil)).Elem()
	if t.Implements(customMarshalerType) {
		return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
			if !plan.IsValid() || (plan.Kind() == reflect.Ptr && plan.IsNil()) {
				return nil, nil
			}
			marshaler := plan.Interface().(CustomMarshaler)
			var stateVal any
			if state.IsValid() {
				stateVal = state.Interface()
			}
			return marshaler.MarshalJSONWithState(plan.Interface(), stateVal)
		}
	}

	if t.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return e.newTimeTypeEncoder()
	}
	if t != reflect.TypeOf(jsontypes.Normalized{}) && t.ConvertibleTo(reflect.TypeOf(timetypes.RFC3339{})) {
		return e.newCustomTimeTypeEncoder()
	}
	if t == reflect.TypeOf((*big.Float)(nil)) {
		return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
			return []byte(plan.Interface().(*big.Float).Text('g', 10)), nil
		}
	}

	nestedEncoder := *e
	nestedEncoder.root = false
	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()
		innerEncoder := nestedEncoder.typeEncoder(inner)
		fullInnerEncoder := innerEncoder
		if nestedEncoder.patch {
			fullEncoder := nestedEncoder
			fullEncoder.patch = false
			fullInnerEncoder = fullEncoder.typeEncoder(inner)
		}

		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			if !s.IsValid() {
				s = reflect.Zero(reflect.PointerTo(inner))
			}
			if !p.IsValid() {
				p = reflect.Zero(reflect.PointerTo(inner))
			}
			if s.IsNil() && p.IsNil() {
				return nil, nil
			}

			if !s.IsNil() && p.IsNil() {
				return explicitJsonNull, nil
			}

			if s.IsNil() {
				s = reflect.New(p.Type().Elem())
				return fullInnerEncoder(p.Elem(), s.Elem())
			}

			return innerEncoder(p.Elem(), s.Elem())
		}
	case reflect.Struct:
		attrType := reflect.TypeOf((*attr.Value)(nil)).Elem()
		if t.Implements(attrType) {
			return nestedEncoder.newTerraformTypeEncoder(t)
		}
		return nestedEncoder.newStructTypeEncoder(t, isRoot)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return nestedEncoder.newArrayTypeEncoder(t)
	case reflect.Map:
		return nestedEncoder.newMapEncoder(t)
	case reflect.Interface:
		return nestedEncoder.newInterfaceEncoder()
	default:
		return nestedEncoder.newPrimitiveTypeEncoder(t)
	}
}

func (e *encoder) newPrimitiveTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.String:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			if e.patch && s.IsValid() && p.String() == s.String() {
				return nil, nil
			}
			return stdjson.Marshal(p.String())
		}
	case reflect.Bool:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			if e.patch && s.IsValid() && p.Bool() == s.Bool() {
				return nil, nil
			}
			if p.Bool() {
				return []byte("true"), nil
			}
			return []byte("false"), nil
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			if e.patch && s.IsValid() && p.Int() == s.Int() {
				return nil, nil
			}
			return []byte(strconv.FormatInt(p.Int(), 10)), nil
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			if e.patch && s.IsValid() && p.Uint() == s.Uint() {
				return nil, nil
			}
			return []byte(strconv.FormatUint(p.Uint(), 10)), nil
		}
	case reflect.Float32:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			if e.patch && s.IsValid() && p.Float() == s.Float() {
				return nil, nil
			}
			return []byte(strconv.FormatFloat(p.Float(), 'f', -1, 32)), nil
		}
	case reflect.Float64:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			if e.patch && s.IsValid() && p.Float() == s.Float() {
				return nil, nil
			}
			return []byte(strconv.FormatFloat(p.Float(), 'f', -1, 64)), nil
		}
	default:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			return nil, fmt.Errorf("unknown type received at primitive encoder: %s", t.String())
		}
	}
}

func (e *encoder) newArrayTypeEncoder(t reflect.Type) encoderFunc {
	arrayPatch := e.patch
	itemEncoderBuilder := *e
	itemEncoderBuilder.patch = false
	itemEncoder := itemEncoderBuilder.typeEncoder(t.Elem())

	return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
		stateNil := !state.IsValid() || state.IsNil()
		planNil := !plan.IsValid() || plan.IsNil()
		if stateNil && planNil {
			return nil, nil
		} else if planNil {
			return explicitJsonNull, nil
		} else if !stateNil && arrayPatch && reflect.DeepEqual(plan.Interface(), state.Interface()) {
			return nil, nil
		}

		json := []byte("[]")
		for i := 0; i < plan.Len(); i++ {
			planItem := plan.Index(i)

			var value, err = itemEncoder(planItem, planItem)

			if err != nil {
				return nil, err
			}
			if value == nil {
				value = explicitJsonNull
			}
			json, err = sjson.SetRawBytes(json, "-1", value)
			if err != nil {
				return nil, err
			}
		}

		return json, nil
	}
}

type terraformUnwrappingFunc func(val attr.Value) (any, diag.Diagnostics)

func (e *encoder) terraformUnwrappedEncoder(underlyingType reflect.Type, unwrap terraformUnwrappingFunc) encoderFunc {
	enc := e.typeEncoder(underlyingType)
	return e.handleNullAndUndefined(func(plan attr.Value, state attr.Value) ([]byte, error) {
		var unwrappedPlan, unwrappedState any
		var diags diag.Diagnostics
		if plan != nil {
			unwrappedPlan, diags = unwrap(plan)
			if diags.HasError() {
				return nil, errorFromDiagnostics(diags)
			}
		}

		if state != nil {
			unwrappedState, diags = unwrap(state)
			if diags.HasError() {
				return nil, errorFromDiagnostics(diags)
			}
		}
		return enc(reflect.ValueOf(unwrappedPlan), reflect.ValueOf(unwrappedState))
	})
}

func (e *encoder) terraformUnwrappedDynamicEncoder(unwrap terraformUnwrappingFunc) encoderFunc {
	return e.handleNullAndUndefined(func(plan attr.Value, state attr.Value) ([]byte, error) {
		var unwrappedPlan, unwrappedState any
		var diags diag.Diagnostics
		if plan != nil {
			unwrappedPlan, diags = unwrap(plan)
			if diags.HasError() {
				return nil, errorFromDiagnostics(diags)
			}
		}
		if state != nil {
			unwrappedState, diags = unwrap(state)
			if diags.HasError() {
				return nil, errorFromDiagnostics(diags)
			}
		}
		enc := e.typeEncoder(reflect.TypeOf(unwrappedPlan))
		return enc(reflect.ValueOf(unwrappedPlan), reflect.ValueOf(unwrappedState))
	})
}

func (e encoder) handleNullAndUndefined(innerFunc func(attr.Value, attr.Value) ([]byte, error)) encoderFunc {
	return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
		var tfPlan attr.Value
		var tfState attr.Value
		if plan.IsValid() {
			tfPlan = plan.Interface().(attr.Value)
		}
		if state.IsValid() {
			tfState = state.Interface().(attr.Value)
		}
		planNull := !plan.IsValid() || tfPlan.IsNull()
		stateNull := !state.IsValid() || tfState.IsNull()
		planUnknown := plan.IsValid() && tfPlan.IsUnknown()
		stateUnknown := state.IsValid() && tfState.IsUnknown()

		if stateNull && planNull {
			return nil, nil
		} else if planNull {
			return explicitJsonNull, nil
		} else if planUnknown {
			return nil, nil
		} else if e.patch && !stateNull && !stateUnknown && tfPlan.Equal(tfState) {
			return nil, nil
		} else {
			return innerFunc(tfPlan, tfState)
		}
	}
}

func UnwrapTerraformAttrValue(ctx context.Context, value attr.Value) (out any, diags diag.Diagnostics) {
	if value == nil {
		return nil, diags
	}

	switch val := value.(type) {
	case basetypes.DynamicValuable:
		v, d := val.ToDynamicValue(ctx)
		o, ds := UnwrapTerraformAttrValue(ctx, v.UnderlyingValue())
		d.Append(ds...)
		return o, d
	case basetypes.BoolValuable:
		v, d := val.ToBoolValue(ctx)
		return v.ValueBoolPointer(), d
	case basetypes.Int32Valuable:
		v, d := val.ToInt32Value(ctx)
		return v.ValueInt32Pointer(), d
	case basetypes.Int64Valuable:
		v, d := val.ToInt64Value(ctx)
		return v.ValueInt64Pointer(), d
	case basetypes.Float32Valuable:
		v, d := val.ToFloat32Value(ctx)
		return v.ValueFloat32Pointer(), d
	case basetypes.Float64Valuable:
		v, d := val.ToFloat64Value(ctx)
		return v.ValueFloat64Pointer(), d
	case basetypes.NumberValue:
		v, d := val.ToNumberValue(ctx)
		return v.ValueBigFloat(), d
	case basetypes.StringValue:
		v, d := val.ToStringValue(ctx)
		return v.ValueStringPointer(), d
	case basetypes.TupleValue:
		return val.Elements(), nil
	case basetypes.ListValuable:
		v, d := val.ToListValue(ctx)
		return v.Elements(), d
	case basetypes.SetValuable:
		v, d := val.ToSetValue(ctx)
		return v.Elements(), d
	case basetypes.MapValuable:
		v, d := val.ToMapValue(ctx)
		return v.Elements(), d
	case basetypes.ObjectValuable:
		v, d := val.ToObjectValue(ctx)
		return v.Attributes(), d
	default:
		diags.AddError("unknown type received at unwrap terraform encoder", fmt.Sprintf("received: %s", value.Type(context.TODO())))
		return nil, diags
	}
}

func (e encoder) newTerraformTypeEncoder(t reflect.Type) encoderFunc {
	ctx := context.TODO()

	if t == reflect.TypeOf(basetypes.BoolValue{}) {
		return e.terraformUnwrappedEncoder(reflect.PointerTo(reflect.TypeOf(true)), func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.Int64Value{}) {
		return e.terraformUnwrappedEncoder(reflect.PointerTo(reflect.TypeOf(int64(0))), func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.Float64Value{}) {
		return e.terraformUnwrappedEncoder(reflect.PointerTo(reflect.TypeOf(float64(0))), func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.NumberValue{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(big.NewFloat(0)), func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.StringValue{}) {
		return e.terraformUnwrappedEncoder(reflect.PointerTo(reflect.TypeOf("")), func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(timetypes.RFC3339{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(time.Time{}), func(value attr.Value) (any, diag.Diagnostics) {
			return value.(timetypes.RFC3339).ValueRFC3339Time()
		})
	} else if t == reflect.TypeOf(basetypes.ListValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.TupleValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.SetValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.MapValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.ObjectValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t.Implements(reflect.TypeOf((*basetypes.DynamicValuable)(nil)).Elem()) {
		return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
			tfPlan, _ := plan.Interface().(basetypes.DynamicValuable).ToDynamicValue(ctx)
			tfState, _ := state.Interface().(basetypes.DynamicValuable).ToDynamicValue(ctx)

			planNull := tfPlan.IsNull() || tfPlan.IsUnderlyingValueNull()
			stateMissing := tfState.IsNull() || tfState.IsUnderlyingValueNull() || tfState.IsUnderlyingValueNull() || tfState.IsUnderlyingValueUnknown()
			if stateMissing && planNull {
				return nil, nil
			} else if planNull {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() || tfPlan.IsUnderlyingValueUnknown() {
				return nil, nil
			} else {
				unwrappedPlan := tfPlan.UnderlyingValue()
				unwrappedState := tfState.UnderlyingValue()
				enc := e.typeEncoder(reflect.TypeOf(unwrappedPlan))
				return enc(reflect.ValueOf(unwrappedPlan), reflect.ValueOf(unwrappedState))
			}
		}
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectLike)(nil)).Elem()) {
		structType := reflect.PointerTo(t.Field(0).Type)
		return e.terraformUnwrappedEncoder(structType, func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.NestedObjectLike).ValueAny(ctx)
		})
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectListLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.NestedObjectListLike).AsStructSlice(ctx)
		})
	} else if t.Implements(reflect.TypeOf((*customfield.ListLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.ListLike).ValueAttr(ctx)
		})
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectMapLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.NestedObjectMapLike).AsStructMap(ctx)
		})
	} else if t.Implements(reflect.TypeOf((*customfield.MapLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.MapLike).ValueAttr(ctx)
		})
	} else if t == reflect.TypeOf(jsontypes.Normalized{}) {
		return e.handleNullAndUndefined(func(plan attr.Value, state attr.Value) ([]byte, error) {
			return []byte(plan.(jsontypes.Normalized).ValueString()), nil
		})
	}

	return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
		return nil, fmt.Errorf("unknown type received at terraform encoder: %s", t.String())
	}
}

func (e *encoder) newStructTypeEncoder(t reflect.Type, isRoot bool) encoderFunc {
	encoderFields := []encoderField{}
	extraEncoder := (*encoderField)(nil)

	var collectEncoderFields func(r reflect.Type, index []int)
	collectEncoderFields = func(r reflect.Type, index []int) {
		for i := 0; i < r.NumField(); i++ {
			idx := append(index, i)
			field := t.FieldByIndex(idx)
			if !field.IsExported() {
				continue
			}
			if field.Anonymous {
				collectEncoderFields(field.Type, idx)
				continue
			}
			ptag, ok := parseJSONStructTag(field)
			if !ok {
				continue
			}
			if ptag.extras && len(index) == 0 {
				extraEncoder = &encoderField{ptag, e.typeEncoder(field.Type.Elem()), idx}
				continue
			}
			if ptag.name == "-" {
				continue
			}
			if ptag.computed && !ptag.forceEncode {
				continue
			}

			fieldEncoder := *e
			dateFormat, ok := parseFormatStructTag(field)
			if ok {
				switch dateFormat {
				case "date-time":
					fieldEncoder.dateFormat = time.RFC3339
				case "date":
					fieldEncoder.dateFormat = "2006-01-02"
				}
			}
			encoderFields = append(
				encoderFields,
				encoderField{ptag, fieldEncoder.typeEncoder(field.Type), idx},
			)
		}
	}
	collectEncoderFields(t, []int{})

	sort.Slice(encoderFields, func(i, j int) bool {
		return encoderFields[i].tag.name < encoderFields[j].tag.name
	})

	return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
		json = []byte("{}")

		someFieldsSet := false
		for _, ef := range encoderFields {
			planField := plan.FieldByIndex(ef.idx)
			stateField, err := state.FieldByIndexErr(ef.idx)
			if err != nil {
				stateField = planField
			}

			planFieldUnknown := false
			if planField.IsValid() {
				attrType := reflect.TypeOf((*attr.Value)(nil)).Elem()
				if planField.Type().Implements(attrType) {
					planFieldUnknown = planField.Interface().(attr.Value).IsUnknown()
				}
			}

			if planFieldUnknown && ef.tag.encodeStateValueWhenPlanUnknown && stateField.IsValid() {
				planField = stateField
			}
			encoded, err := ef.fn(planField, stateField)
			if err != nil {
				return nil, err
			}
			if encoded == nil {
				continue
			}
			someFieldsSet = true
			json, err = sjson.SetRawBytes(json, EscapeSJSONKey(ef.tag.name), encoded)
			if err != nil {
				return nil, err
			}
		}

		if !someFieldsSet && e.patch {
			if isRoot {
				return []byte("{}"), nil
			}
			return nil, nil
		}

		if extraEncoder != nil {
			json, err = e.encodeMapEntries(json, plan.FieldByIndex(extraEncoder.idx), state.FieldByIndex(extraEncoder.idx))
			if err != nil {
				return nil, err
			}
		}
		return
	}
}

func (e *encoder) newTimeTypeEncoder() encoderFunc {
	format := e.dateFormat
	return func(value reflect.Value, state reflect.Value) (json []byte, err error) {
		return []byte(`"` + value.Convert(reflect.TypeOf(time.Time{})).Interface().(time.Time).Format(format) + `"`), nil
	}
}

func (e *encoder) newCustomTimeTypeEncoder() encoderFunc {
	format := e.dateFormat
	return e.handleNullAndUndefined(func(value attr.Value, state attr.Value) (json []byte, err error) {
		val, errs := value.(timetypes.RFC3339).ValueRFC3339Time()
		if errs != nil {
			return nil, errorFromDiagnostics(errs)
		}
		return stdjson.Marshal(val.Format(format))
	})
}

func (e encoder) newInterfaceEncoder() encoderFunc {
	return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
		plan = plan.Elem()
		state = state.Elem()
		if !plan.IsValid() {
			return nil, nil
		}
		if !state.IsValid() {
			return nil, nil
		}
		return e.typeEncoder(plan.Type())(plan, state)
	}
}

func encodeKey(key reflect.Value, keyEncoder encoderFunc) (string, error) {
	if key.Type().Kind() == reflect.String {
		return key.String(), nil
	}

	encodedKeyBytes, err := keyEncoder(key, key)
	if err != nil {
		return "", err
	}

	return string(encodedKeyBytes), nil
}

func (e *encoder) encodeMapEntries(json []byte, plan reflect.Value, state reflect.Value) ([]byte, error) {
	type mapPair struct {
		key   []byte
		plan  reflect.Value
		state reflect.Value
	}

	pairKeys := map[string]bool{}
	pairs := []mapPair{}
	keyEncoder := e.typeEncoder(plan.Type().Key())
	elementEncoder := e.typeEncoder(plan.Type().Elem())
	fullElementEncoder := elementEncoder
	if e.patch {
		fullEncoder := *e
		fullEncoder.patch = false
		fullElementEncoder = fullEncoder.typeEncoder(plan.Type().Elem())
	}

	iter := plan.MapRange()
	for iter.Next() {
		encodedKeyString, err := encodeKey(iter.Key(), keyEncoder)
		if err != nil {
			return nil, err
		}

		pairKeys[encodedKeyString] = true
		encodedKey := []byte(encodedKeyString)
		pairs = append(pairs, mapPair{key: encodedKey, plan: iter.Value(), state: state.MapIndex(iter.Key())})
	}

	if e.patch {
		iter = state.MapRange()
		for iter.Next() {
			encodedKeyString, err := encodeKey(iter.Key(), keyEncoder)
			if err != nil {
				return nil, err
			}

			if _, ok := pairKeys[encodedKeyString]; ok {
				continue
			}

			pairKeys[encodedKeyString] = true
			encodedKey := []byte(encodedKeyString)
			pairs = append(pairs, mapPair{key: encodedKey, plan: plan.MapIndex(iter.Key()), state: iter.Value()})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return bytes.Compare(pairs[i].key, pairs[j].key) < 0
	})

	for _, pair := range pairs {
		var encodedValue []byte
		var err error

		if pair.plan.IsValid() && pair.state.IsValid() {
			if e.patch && reflect.DeepEqual(pair.plan.Interface(), pair.state.Interface()) {
				continue
			}
			encodedValue, err = elementEncoder(pair.plan, pair.state)
		} else if pair.plan.IsValid() {
			encodedValue, err = fullElementEncoder(pair.plan, pair.plan)
		} else {
			encodedValue = nil
		}
		if err != nil {
			return nil, err
		}
		if encodedValue == nil {
			encodedValue = explicitJsonNull
		}
		json, err = sjson.SetRawBytes(json, EscapeSJSONKey(string(pair.key)), encodedValue)
		if err != nil {
			return nil, err
		}
	}

	return json, nil
}

func (e *encoder) newMapEncoder(_ reflect.Type) encoderFunc {
	patch := e.patch
	return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
		stateNil := !state.IsValid() || state.IsNil()
		planNil := !plan.IsValid() || plan.IsNil()
		if stateNil && planNil {
			return nil, nil
		} else if planNil {
			return explicitJsonNull, nil
		} else if patch && !stateNil && reflect.DeepEqual(plan.Interface(), state.Interface()) {
			return nil, nil
		} else if state.Kind() != plan.Kind() {
			fullEncoder := *e
			fullEncoder.patch = false
			return fullEncoder.encodeMapEntries([]byte("{}"), plan, plan)
		}

		return e.encodeMapEntries([]byte("{}"), plan, state)
	}
}
