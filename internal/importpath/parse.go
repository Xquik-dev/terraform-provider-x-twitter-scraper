// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package importpath decodes a Terraform import ID.
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
		diags.AddError("Invalid import ID", fmt.Sprintf("The import ID must be 1 URL-encoded segment without a slash, got %q.", id))
		return "", diags
	}
	segment, err := url.PathUnescape(id)
	if err != nil {
		diags.AddError("Invalid import ID", fmt.Sprintf("The import ID %q has an %s.", id, err))
	}
	return segment, diags
}
