// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package style

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*StyleResource)(nil)

func (r *StyleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
