// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package apiform encodes Terraform models as multipart form bodies for API
// requests.
package apiform

type Marshaler interface {
	MarshalMultipart() ([]byte, string, error)
}
