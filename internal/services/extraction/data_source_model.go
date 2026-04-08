// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package extraction

import (
	"context"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

type ExtractionDataSourceModel struct {
	ID         types.String                                            `tfsdk:"id" path:"id,required"`
	After      types.String                                            `tfsdk:"after" query:"after,optional"`
	Limit      types.Int64                                             `tfsdk:"limit" query:"limit,computed_optional"`
	HasMore    types.Bool                                              `tfsdk:"has_more" json:"hasMore,computed"`
	NextCursor types.String                                            `tfsdk:"next_cursor" json:"nextCursor,computed"`
	Job        customfield.Map[jsontypes.Normalized]                   `tfsdk:"job" json:"job,computed"`
	Results    customfield.List[customfield.Map[jsontypes.Normalized]] `tfsdk:"results" json:"results,computed"`
}

func (m *ExtractionDataSourceModel) toReadParams(_ context.Context) (params xtwitterscraper.ExtractionGetParams, diags diag.Diagnostics) {
	params = xtwitterscraper.ExtractionGetParams{}

	if !m.After.IsNull() {
		params.After = param.NewOpt(m.After.ValueString())
	}
	if !m.Limit.IsNull() {
		params.Limit = param.NewOpt(m.Limit.ValueInt64())
	}

	return
}
