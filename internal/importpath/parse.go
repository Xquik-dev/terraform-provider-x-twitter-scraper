// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package importpath parses Terraform import IDs into their path segments.
package importpath

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ParseImportID(str string, format string, targets ...*string) (diags diag.Diagnostics) {
	pathSpec, path := strings.Split(format, "/"), strings.Split(str, "/")

	if len(targets) != len(pathSpec) {
		panic(fmt.Sprintf("argument count does not match format: %q", format))
	}
	if len(path) != len(pathSpec) {
		diags.AddError("invalid ID", fmt.Sprintf("expected urlencoded segments %q, got %q", format, str))
		return
	}

	var err error
	for i, target := range targets {
		*target, err = url.PathUnescape(path[i])
		if err != nil {
			diags.AddError("invalid urlencoded segment", fmt.Sprintf("for %q : %s -> %q", pathSpec[i], err.Error(), path[i]))
		}
	}
	return
}
