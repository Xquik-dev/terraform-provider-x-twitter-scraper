// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/account"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/compose"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/draft"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/draw"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/event"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/guest_wallet"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor_keyword"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/style"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/support_ticket"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/webhook"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_account"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_tweet"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_user"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_write"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/x_write_action"
	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.ProviderWithConfigValidators = (*XTwitterScraperProvider)(nil)

// XTwitterScraperProvider defines the provider implementation.
type XTwitterScraperProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// XTwitterScraperProviderModel describes the provider data model.
type XTwitterScraperProviderModel struct {
	BaseURL     types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey      types.String `tfsdk:"api_key" json:"api_key,optional"`
	BearerToken types.String `tfsdk:"bearer_token" json:"bearer_token,optional"`
}

func (p *XTwitterScraperProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "x-twitter-scraper"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: `Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.`,
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"bearer_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *XTwitterScraperProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *XTwitterScraperProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data XTwitterScraperProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("X_TWITTER_SCRAPER_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	} else if o, ok := os.LookupEnv("X_TWITTER_SCRAPER_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	}

	if !data.BearerToken.IsNull() && !data.BearerToken.IsUnknown() {
		opts = append(opts, option.WithBearerToken(data.BearerToken.ValueString()))
	} else if o, ok := os.LookupEnv("X_TWITTER_SCRAPER_BEARER_TOKEN"); ok {
		opts = append(opts, option.WithBearerToken(o))
	}

	client := xtwitterscraper.NewClient(
		opts...,
	)

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *XTwitterScraperProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *XTwitterScraperProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		compose.NewResource,
		draft.NewResource,
		style.NewResource,
		monitor.NewResource,
		monitor_keyword.NewResource,
		webhook.NewResource,
		x_write.NewCreateTweetResource,
		x_write.NewDeleteTweetResource,
		x_write.NewLikeResource,
		x_write.NewUnlikeResource,
		x_write.NewRetweetResource,
		x_write.NewUnretweetResource,
		x_write.NewFollowResource,
		x_write.NewUnfollowResource,
		x_write.NewRemoveFollowerResource,
		x_write.NewSendDMResource,
		x_write.NewUploadMediaResource,
		x_write.NewUpdateProfileResource,
		x_write.NewUpdateAvatarResource,
		x_write.NewUpdateBannerResource,
		x_write.NewCreateCommunityResource,
		x_write.NewDeleteCommunityResource,
		x_write.NewJoinCommunityResource,
		x_write.NewLeaveCommunityResource,
		support_ticket.NewResource,
		guest_wallet.NewResource,
	}
}

func (p *XTwitterScraperProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		account.NewAccountDataSource,
		draft.NewDraftDataSource,
		style.NewStyleDataSource,
		monitor.NewMonitorDataSource,
		event.NewEventDataSource,
		monitor_keyword.NewMonitorKeywordDataSource,
		draw.NewDrawDataSource,
		x_write_action.NewXWriteActionDataSource,
		x_tweet.NewXTweetDataSource,
		x_user.NewXUserDataSource,
		x_account.NewXAccountDataSource,
		support_ticket.NewSupportTicketDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &XTwitterScraperProvider{
			version: version,
		}
	}
}
