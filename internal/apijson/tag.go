// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"reflect"
	"strings"
)

const jsonStructTag = "json"
const optionsStructTag = "api"
const formatStructTag = "format"

type parsedStructTag struct {
	name                            string
	extras                          bool
	metadata                        bool
	inline                          bool
	required                        bool
	optional                        bool
	computed                        bool
	computedOptional                bool
	noRefresh                       bool
	forceEncode                     bool
	encodeStateValueWhenPlanUnknown bool
}

func parseJSONStructTag(field reflect.StructField) (tag parsedStructTag, ok bool) {
	raw, ok := field.Tag.Lookup(jsonStructTag)
	if !ok {
		return
	}
	tag.name, _, _ = strings.Cut(raw, ",")
	options, _ := field.Tag.Lookup(optionsStructTag)
	for part := range strings.SplitSeq(options, ",") {
		switch part {
		case "extras":
			tag.extras = true
		case "metadata":
			tag.metadata = true
		case "inline":
			tag.inline = true
		case "required":
			tag.required = true
		case "optional":
			tag.optional = true
		case "computed":
			tag.computed = true
		case "computed_optional":
			tag.computedOptional = true
		case "no_refresh":
			tag.noRefresh = true
		case "encode_state_for_unknown":
			tag.encodeStateValueWhenPlanUnknown = true
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
