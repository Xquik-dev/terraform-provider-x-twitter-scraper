// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package extraction

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stainless-sdks/x-twitter-scraper-go"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/logging"
)

type ExtractionDataSource struct {
	client *xtwitterscraper.Client
}

var _ datasource.DataSourceWithConfigure = (*ExtractionDataSource)(nil)

func NewExtractionDataSource() datasource.DataSource {
	return &ExtractionDataSource{}
}

func (d *ExtractionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_extraction"
}

func (d *ExtractionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*xtwitterscraper.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *xtwitterscraper.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ExtractionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ExtractionDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toReadParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := d.client.Extractions.Get(
		ctx,
		data.ID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
