// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func NewFloat64ValueFromBigFloat(value *big.Float) (basetypes.Float64Value, error) {
	tfVal := tftypes.NewValue(tftypes.Number, value)
	out, err := basetypes.Float64Type{}.ValueFromTerraform(context.Background(), tfVal)
	if out != nil && err == nil {
		return out.(basetypes.Float64Value), nil
	}
	return basetypes.Float64Value{}, err
}

func NewFloat64Value(value float64) basetypes.Float64Value {
	out, _ := NewFloat64ValueFromBigFloat(new(big.Float).SetPrec(512).SetMode(big.ToNearestEven).SetFloat64(value))
	return out
}

func NewFloat64ValueFromString(value string) (basetypes.Float64Value, error) {
	bf, err := parseBigFloatingPoint(value)
	if err != nil {
		return basetypes.Float64Value{}, fmt.Errorf("failed to parse as basetypes.Float64Value: %w", err)
	}
	return NewFloat64ValueFromBigFloat(bf)
}

func NewFloat64ValueFromStringUnsafe(value string) basetypes.Float64Value {
	out, err := NewFloat64ValueFromString(value)
	if err != nil {
		panic(err)
	}
	return out
}

func NewNumberValueFromString(value string) (basetypes.NumberValue, error) {
	bf, err := parseBigFloatingPoint(value)
	if err != nil {
		return basetypes.NumberValue{}, fmt.Errorf("failed to parse as basetypes.NumberValue: %w", err)
	}
	return basetypes.NewNumberValue(bf), nil
}

func NewNumberValueFromStringUnsafe(value string) basetypes.NumberValue {
	out, err := NewNumberValueFromString(value)
	if err != nil {
		panic(err)
	}
	return out
}

func parseBigFloatingPoint(value string) (*big.Float, error) {
	bf, _, err := big.ParseFloat(value, 10, 512, big.ToNearestEven)
	return bf, err
}
