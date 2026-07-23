// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package test_helpers

import (
	"testing"

	ds "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rs "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type emptyContractModel struct{}

func TestIntegrityValidatorsAcceptMatchingEmptyModels(t *testing.T) {
	t.Parallel()

	if errors := ValidateDataSourceModelSchemaIntegrity((*emptyContractModel)(nil), ds.Schema{}); len(errors) != 0 {
		t.Fatalf("unexpected data source errors: %v", errors)
	}
	if errors := ValidateResourceModelSchemaIntegrity((*emptyContractModel)(nil), rs.Schema{}); len(errors) != 0 {
		t.Fatalf("unexpected resource errors: %v", errors)
	}
}
