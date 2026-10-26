// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package xlist implements the x-twitter-scraper_x_list data source.
package xlist

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/logging"
	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type XListDataSource struct {
	client *xtwitterscraper.Client
}

var _ datasource.DataSourceWithConfigure = (*XListDataSource)(nil)

func NewXListDataSource() datasource.DataSource {
	return &XListDataSource{}
}

func (d *XListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_x_list"
}

func (d *XListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *XListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *XListDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := d.client.X.Lists.Get(
		ctx,
		data.ID.ValueString(),
		xtwitterscraper.XListGetParams{},
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
