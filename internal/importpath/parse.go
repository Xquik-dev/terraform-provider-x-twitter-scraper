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

func ParseImportID(id string) (string, diag.Diagnostics) {
	var diags diag.Diagnostics
	if strings.Contains(id, "/") {
		diags.AddError("invalid ID", fmt.Sprintf("expected urlencoded segments %q, got %q", "<id>", id))
		return "", diags
	}
	segment, err := url.PathUnescape(id)
	if err != nil {
		diags.AddError("invalid urlencoded segment", fmt.Sprintf("for %q : %s -> %q", "<id>", err.Error(), id))
	}
	return segment, diags
}
