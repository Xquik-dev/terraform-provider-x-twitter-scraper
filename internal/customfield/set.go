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
	_ basetypes.SetTypable  = (*SetType[basetypes.StringValue])(nil)
	_ basetypes.SetValuable = (*Set[basetypes.StringValue])(nil)
)

type SetType[T attr.Value] struct {
	basetypes.SetType
}

func NewSetType[T attr.Value](ctx context.Context) SetType[T] {
	return SetType[T]{SetType: basetypes.SetType{ElemType: elemType[T](ctx)}}
}

func (t SetType[T]) Equal(o attr.Type) bool {
	other, ok := o.(SetType[T])
	if !ok {
		return false
	}

	return t.SetType.Equal(other.SetType)
}

func (t SetType[T]) String() string {
	var zero T
	return fmt.Sprintf("SetType[%T]", zero)
}

func (t SetType[T]) ValueFromSet(ctx context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullSet[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownSet[T](ctx), diags
	}

	v, d := basetypes.NewSetValue(elemType[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownSet[T](ctx), diags
	}

	value := Set[T]{
		SetValue: v,
	}

	return value, diags
}

func (t SetType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.SetType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	setValue, ok := attrValue.(basetypes.SetValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	setValuable, diags := t.ValueFromSet(ctx, setValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return setValuable, nil
}

func (t SetType[T]) ValueType(ctx context.Context) attr.Value {
	return UnknownSet[T](ctx)
}

func (t SetType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullSet[T](ctx), diags
}

var _ ListLike = (*Set[basetypes.StringValue])(nil)

type Set[T attr.Value] struct {
	basetypes.SetValue
}

func (v Set[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.SetValue
	if tv.ElementType(ctx) == nil {
		tv = NullSet[T](ctx).SetValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v Set[T]) NullValue(ctx context.Context) ListLike {
	return NullSet[T](ctx)
}

func (v Set[T]) UnknownValue(ctx context.Context) ListLike {
	return UnknownSet[T](ctx)
}

func (v Set[T]) DecodeTarget(ctx context.Context) (reflect.Value, func() ListLike) {
	values, _ := v.Value(ctx)
	return reflect.ValueOf(&values).Elem(), func() ListLike {
		r, _ := NewSet(ctx, values)
		return r
	}
}

func (v Set[T]) ModelType() reflect.Type {
	return reflect.TypeFor[T]()
}

func (v Set[T]) IsNullOrUnknown() bool {
	return v.IsNull() || v.IsUnknown()
}

func (v Set[T]) Equal(o attr.Value) bool {
	other, ok := o.(Set[T])
	if !ok {
		return false
	}

	return v.SetValue.Equal(other.SetValue)
}

func (v Set[T]) Type(ctx context.Context) attr.Type {
	return NewSetType[T](ctx)
}

func (v Set[T]) ValueAttr(ctx context.Context) (reflect.Value, diag.Diagnostics) {
	ts, diags := v.Value(ctx)
	return reflect.ValueOf(ts), diags
}

func (v Set[T]) Value(ctx context.Context) ([]T, diag.Diagnostics) {
	var diags diag.Diagnostics
	ts := []T{}
	for _, elem := range v.SetValue.Elements() {
		t, ok := elem.(T)
		if !ok {
			diags.AddError("unexpected element type", fmt.Sprintf("expected %T, got %T", t, elem))
			return nil, diags
		}
		ts = append(ts, t)
	}
	return ts, diags
}

func NullSet[T attr.Value](ctx context.Context) Set[T] {
	return Set[T]{SetValue: basetypes.NewSetNull(elemType[T](ctx))}
}

func UnknownSet[T attr.Value](ctx context.Context) Set[T] {
	return Set[T]{SetValue: basetypes.NewSetUnknown(elemType[T](ctx))}
}

func NewSet[T attr.Value](ctx context.Context, values []T) (Set[T], diag.Diagnostics) {
	attrs := make([]attr.Value, len(values))
	for i, v := range values {
		attrs[i] = v
	}
	return NewSetFromAttributes[T](ctx, attrs)
}

func NewSetFromAttributes[T attr.Value](ctx context.Context, attrs []attr.Value) (Set[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	setValue, d := basetypes.NewSetValue(elemType[T](ctx), attrs)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownSet[T](ctx), diags
	}

	return Set[T]{SetValue: setValue}, nil
}

func NewSetMust[T attr.Value](ctx context.Context, values []attr.Value) Set[T] {
	o, diags := NewSetFromAttributes[T](ctx, values)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating a customfield.Set: %v", diags))
	}
	return o
}
