// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apiform

import (
	"reflect"
	"strings"
)

const jsonStructTag = "json"
const formStructTag = "form"
const optionsStructTag = "api"
const formatStructTag = "format"

type parsedStructTag struct {
	name        string
	required    bool
	extras      bool
	metadata    bool
	computed    bool
	forceEncode bool
}

func parseFormStructTag(field reflect.StructField) (tag parsedStructTag, ok bool) {
	raw, ok := field.Tag.Lookup(formStructTag)
	if !ok {
		raw, ok = field.Tag.Lookup(jsonStructTag)
	}
	if !ok {
		return
	}
	tag.name, _, _ = strings.Cut(raw, ",")
	options, _ := field.Tag.Lookup(optionsStructTag)
	for part := range strings.SplitSeq(options, ",") {
		switch part {
		case "required":
			tag.required = true
		case "extras":
			tag.extras = true
		case "metadata":
			tag.metadata = true
		case "computed":
			tag.computed = true
		case "force_encode":
			tag.forceEncode = true
		}
	}
	return
}

func parseFormatStructTag(field reflect.StructField) (format string, ok bool) {
	format, ok = field.Tag.Lookup(formatStructTag)
	return
}
