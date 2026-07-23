// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TestAccessorsRejectNilAndNullNumbers(t *testing.T) {
	t.Parallel()

	if ok, value := floatValue(nil); ok || value != nil {
		t.Fatalf("floatValue(nil) = %t, %v", ok, value)
	}
	if ok, value := floatValue(basetypes.NewNumberNull()); ok || value != nil {
		t.Fatalf("floatValue(null) = %t, %v", ok, value)
	}
	if ok, value := intValue(nil); ok || value != nil {
		t.Fatalf("intValue(nil) = %t, %v", ok, value)
	}
	if ok, value := intValue(basetypes.NewNumberUnknown()); ok || value != nil {
		t.Fatalf("intValue(unknown) = %t, %v", ok, value)
	}
	if ok, values := ChildItems(nil); ok || values != nil {
		t.Fatalf("ChildItems(nil) = %t, %v", ok, values)
	}
	if ok, values := ChildAttributes(nil); ok || values != nil {
		t.Fatalf("ChildAttributes(nil) = %t, %v", ok, values)
	}
}

func TestStringConvertersRejectInvalidNumbers(t *testing.T) {
	t.Parallel()

	if _, err := NewFloat64ValueFromString("not-a-number"); err == nil {
		t.Fatal("invalid float string produced no error")
	}
	if _, err := NewNumberValueFromString("not-a-number"); err == nil {
		t.Fatal("invalid number string produced no error")
	}
	assertConversionPanics(t, func() {
		_ = NewFloat64ValueFromStringUnsafe("not-a-number")
	})
	assertConversionPanics(t, func() {
		_ = NewNumberValueFromStringUnsafe("not-a-number")
	})
}

func assertConversionPanics(t *testing.T, action func()) {
	t.Helper()
	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("conversion did not panic")
		}
	}()
	action()
}
