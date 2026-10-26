// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package importpath parses Terraform import IDs into their path segments.
package importpath

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type Segment interface {
	parse(spec string, segment string) diag.Diagnostics
}

type boolSegment struct{ target *bool }
type int64Segment struct{ target *int64 }
type float64Segment struct{ target *float64 }
type stringSegment struct{ target *string }

func Bool(target *bool) Segment       { return boolSegment{target} }
func Int64(target *int64) Segment     { return int64Segment{target} }
func Float64(target *float64) Segment { return float64Segment{target} }
func String(target *string) Segment   { return stringSegment{target} }

func (s boolSegment) parse(spec string, segment string) (diags diag.Diagnostics) {
	var err error
	*s.target, err = strconv.ParseBool(segment)
	if err != nil {
		diags.AddError("invalid bool segment", fmt.Sprintf("for %q : %s -> %q", spec, err.Error(), segment))
	}
	return
}

func (s int64Segment) parse(spec string, segment string) (diags diag.Diagnostics) {
	var err error
	*s.target, err = strconv.ParseInt(segment, 10, 64)
	if err != nil {
		diags.AddError("invalid int segment", fmt.Sprintf("for %q : %s -> %q", spec, err.Error(), segment))
	}
	return
}

func (s float64Segment) parse(spec string, segment string) (diags diag.Diagnostics) {
	var err error
	*s.target, err = strconv.ParseFloat(segment, 64)
	if err != nil {
		diags.AddError("invalid float segment", fmt.Sprintf("for %q : %s -> %q", spec, err.Error(), segment))
	}
	return
}

func (s stringSegment) parse(spec string, segment string) (diags diag.Diagnostics) {
	var err error
	*s.target, err = url.PathUnescape(segment)
	if err != nil {
		diags.AddError("invalid urlencoded segment", fmt.Sprintf("for %q : %s -> %q", spec, err.Error(), segment))
	}
	return
}

func ParseImportID(str string, format string, segments ...Segment) (diags diag.Diagnostics) {
	pathSpec, path := strings.Split(format, "/"), strings.Split(str, "/")

	if len(segments) != len(pathSpec) {
		panic(fmt.Sprintf("argument count does not match format: %q", format))
	}
	if len(path) != len(pathSpec) {
		diags.AddError("invalid ID", fmt.Sprintf("expected urlencoded segments %q, got %q", format, str))
		return
	}

	for i, segment := range segments {
		diags.Append(segment.parse(pathSpec[i], path[i])...)
	}
	return
}
