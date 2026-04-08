// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-go"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/account"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/api_key"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/draft"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/draw"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/event"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/extraction"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/integration"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/monitor"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/style"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/support_ticket"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/webhook"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_account"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_bookmark"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_community"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_community_join"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_community_tweet"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_profile"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_tweet"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_tweet_like"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_tweet_retweet"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_user"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/services/x_user_follow"
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
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
			"bearer_token": schema.StringAttribute{
				Optional: true,
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
		api_key.NewResource,
		draft.NewResource,
		style.NewResource,
		monitor.NewResource,
		webhook.NewResource,
		integration.NewResource,
		x_tweet.NewResource,
		x_tweet_like.NewResource,
		x_tweet_retweet.NewResource,
		x_user_follow.NewResource,
		x_profile.NewResource,
		x_community.NewResource,
		x_community_join.NewResource,
		x_account.NewResource,
		support_ticket.NewResource,
	}
}

func (p *XTwitterScraperProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		account.NewAccountDataSource,
		draft.NewDraftDataSource,
		style.NewStyleDataSource,
		monitor.NewMonitorDataSource,
		event.NewEventDataSource,
		extraction.NewExtractionDataSource,
		draw.NewDrawDataSource,
		integration.NewIntegrationDataSource,
		x_tweet.NewXTweetDataSource,
		x_user.NewXUserDataSource,
		x_community_tweet.NewXCommunityTweetsDataSource,
		x_account.NewXAccountDataSource,
		x_bookmark.NewXBookmarksDataSource,
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
