// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package apienum holds the API spec enums that provider schemas validate
// against, and describes their values for the Terraform Registry docs.
package apienum

//go:generate go run ../../cmd/api-enums EventType PostEventType

import (
	"strconv"
	"strings"
)

func Describe(values []string) string {
	quoted := make([]string, len(values))
	for index, value := range values {
		quoted[index] = strconv.Quote(value)
	}
	last := len(quoted) - 1
	if last == 0 {
		return "The only available value is " + quoted[0] + "."
	}
	return "Available values are " + strings.Join(quoted[:last], ", ") + " and " + quoted[last] + "."
}
