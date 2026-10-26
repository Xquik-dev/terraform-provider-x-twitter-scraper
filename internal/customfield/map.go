// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package customfield

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.MapTypable  = (*MapType[basetypes.StringValue])(nil)
	_ basetypes.MapValuable = (*Map[basetypes.StringValue])(nil)
)

type MapLike interface {
	basetypes.MapValuable
	ValueAttr(ctx context.Context) (reflect.Value, diag.Diagnostics)
	NullValue(ctx context.Context) MapLike
	UnknownValue(ctx context.Context) MapLike
	KnownValue(ctx context.Context, values reflect.Value) MapLike
	ModelType() reflect.Type
}

type MapType[T attr.Value] struct {
	basetypes.MapType
}

func newMapType[T attr.Value](ctx context.Context) (MapType[T], diag.Diagnostics) {
	return MapType[T]{MapType: basetypes.MapType{ElemType: elemType[T](ctx)}}, nil
}

func NewMapType[T attr.Value](ctx context.Context) MapType[T] {
	t, _ := newMapType[T](ctx)
	return t
}

func (t MapType[T]) Equal(o attr.Type) bool {
	other, ok := o.(MapType[T])
	if !ok {
		return false
	}

	return t.MapType.Equal(other.MapType)
}

func (t MapType[T]) String() string {
	var zero T
	return fmt.Sprintf("MapType[%T]", zero)
}

func (t MapType[T]) ValueFromMap(ctx context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullMap[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownMap[T](ctx), diags
	}

	v, d := basetypes.NewMapValue(elemType[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownMap[T](ctx), diags
	}

	value := Map[T]{
		MapValue: v,
	}

	return value, diags
}

func (t MapType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.MapType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	mapValue, ok := attrValue.(basetypes.MapValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	mapValuable, diags := t.ValueFromMap(ctx, mapValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting MapValue to MapValuable: %v", diags)
	}

	return mapValuable, nil
}

func (t MapType[T]) ValueType(ctx context.Context) attr.Value {
	return UnknownMap[T](ctx)
}

func (t MapType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullMap[T](ctx), diags
}

var _ MapLike = (*Map[basetypes.StringValue])(nil)

type Map[T attr.Value] struct {
	basetypes.MapValue
}

func (v Map[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.MapValue
	if tv.ElementType(ctx) == nil {
		tv = NullMap[T](ctx).MapValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v Map[T]) NullValue(ctx context.Context) MapLike {
	return NullMap[T](ctx)
}

func (v Map[T]) UnknownValue(ctx context.Context) MapLike {
	return UnknownMap[T](ctx)
}

func (v Map[T]) KnownValue(ctx context.Context, values reflect.Value) MapLike {
	if ts, ok := reflect.TypeAssert[map[string]T](values); ok {
		r, _ := NewMap(ctx, ts)
		return r
	}
	if attrs, ok := reflect.TypeAssert[map[string]attr.Value](values); ok {
		r, _ := NewMapFromAttributes[T](ctx, attrs)
		return r
	}
	return UnknownMap[T](ctx)
}

func (v Map[T]) ModelType() reflect.Type {
	return reflect.TypeFor[T]()
}

func (v Map[T]) Equal(o attr.Value) bool {
	other, ok := o.(Map[T])
	if !ok {
		return false
	}

	return v.MapValue.Equal(other.MapValue)
}

func (v Map[T]) Type(ctx context.Context) attr.Type {
	return NewMapType[T](ctx)
}

func (v Map[T]) ValueAttr(ctx context.Context) (reflect.Value, diag.Diagnostics) {
	ts, diags := v.Value(ctx)
	return reflect.ValueOf(ts), diags
}

func (v Map[T]) Value(ctx context.Context) (map[string]T, diag.Diagnostics) {
	var diags diag.Diagnostics
	ts := map[string]T{}
	for key, elem := range v.MapValue.Elements() {
		t, ok := elem.(T)
		if !ok {
			diags.AddError("unexpected element type", fmt.Sprintf("expected %T, got %T", t, elem))
			return nil, diags
		}
		ts[key] = t
	}
	return ts, diags
}

func NullMap[T attr.Value](ctx context.Context) Map[T] {
	return Map[T]{MapValue: basetypes.NewMapNull(elemType[T](ctx))}
}

func UnknownMap[T attr.Value](ctx context.Context) Map[T] {
	return Map[T]{MapValue: basetypes.NewMapUnknown(elemType[T](ctx))}
}

func NewMap[T attr.Value](ctx context.Context, values map[string]T) (Map[T], diag.Diagnostics) {
	attrs := make(map[string]attr.Value, len(values))
	for i, v := range values {
		attrs[i] = v
	}
	return NewMapFromAttributes[T](ctx, attrs)
}

func NewMapFromAttributes[T attr.Value](ctx context.Context, attrs map[string]attr.Value) (Map[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	setValue, d := basetypes.NewMapValue(elemType[T](ctx), attrs)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownMap[T](ctx), diags
	}

	return Map[T]{MapValue: setValue}, nil
}

func NewMapMust[T attr.Value](ctx context.Context, values map[string]T) Map[T] {
	o, diags := NewMap[T](ctx, values)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating a customfield.Map: %v", diags))
	}
	return o
}
