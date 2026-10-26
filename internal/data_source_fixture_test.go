// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"context"
	"embed"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"path"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

//go:embed testdata/data_sources
var dataSourceFixtures embed.FS

func TestDataSourcesReadTheSpecExamples(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		typeName string
		id       string
		path     string
		fixture  string
		want     map[string]stateWant
	}{
		{"x-twitter-scraper_x_list", "4143408", "/api/v1/x/lists/4143408", "x_list.json", map[string]stateWant{
			"id":                    wantText("4143408"),
			"list.id":               wantText("4143408"),
			"list.name":             wantText("NASA Centers"),
			"list.description":      wantText("NASA has 10 centers"),
			"list.member_count":     wantNumber(11),
			"list.subscriber_count": wantNumber(859),
			"list.owner.id":         wantText("11348282"),
			"list.owner.username":   wantText("NASA"),
			"list.url":              wantText("https://x.com/i/lists/4143408"),
			"list.created_at":       wantText("2009-12-02T22:41:02.000Z"),
			"list.banner_url":       wantText("https://pbs.twimg.com/media/example.png"),
			"list.members_context":  wantText("11 members"),
		}},
		{"x-twitter-scraper_x_job", "1899202395210649861", "/api/v1/x/jobs/1899202395210649861", "x_job.json", map[string]stateWant{
			"job.id":                    wantText("1899202395210649861"),
			"job.title":                 wantText("Engineer"),
			"job.location_type":         wantText("remote"),
			"job.employment_type":       wantText("full_time"),
			"job.seniority_level":       wantText("lead"),
			"job.salary_min":            wantNumber(225000),
			"job.salary_currency_code":  wantText("USD"),
			"job.featured":              wantBool(true),
			"job.description_markdown":  wantText("# What is Commonware?"),
			"job.account.username":      wantText("commonwarexyz"),
			"job.account.verified":      wantBool(true),
			"job.account.verified_type": wantText("Business"),
		}},
		{"x-twitter-scraper_x_space", "1OGwblLjnQMKB", "/api/v1/x/spaces/1OGwblLjnQMKB", "x_space.json", map[string]stateWant{
			"space.id":                   wantText("1OGwblLjnQMKB"),
			"space.title":                wantText("Extending the Network to the Edge"),
			"space.state":                wantText("Ended"),
			"space.started_at":           wantText("2026-03-11T17:00:26.227Z"),
			"space.total_live_listeners": wantNumber(841),
			"space.total_replay_watched": wantNumber(609),
			"space.creator.username":     wantText("solana"),
			"space.admins.0.id":          wantText("951329744804392960"),
			"space.admins.0.username":    wantText("solana"),
			"space.tweet.text":           wantText("Extending the Network to the Edge"),
			"space.sharings.0.tweet":     wantTweetJSON{text: "Join the Space", urls: []string{}},
			"space.sharings.1.tweet":     wantTweetJSON{text: "Replay the Space", urls: []string{"https://x.com/i/spaces/1OGwblLjnQMKB"}},
		}},
		{"x-twitter-scraper_x_user", "NASA", "/api/v1/x/users/NASA", "x_user.json", map[string]stateWant{
			"username":                                wantText("NASA"),
			"followers":                               wantNumber(92376187),
			"verified_type":                           wantText("Government"),
			"profile_picture_original":                wantText("https://pbs.twimg.com/profile_images/1321163587679784960/0ZxKlEKB.jpg"),
			"birthdate.year":                          wantNumber(1958),
			"birthdate.month":                         wantNumber(10),
			"contact_info.links.0.url":                wantText("http://www.nasa.gov/"),
			"contact_info.links.0.network":            wantText("website"),
			"verification_reason.entities.0.url_type": wantText("ExternalUrl"),
		}},
		{"x-twitter-scraper_x_broadcast", "1nGeLMWbkpaKX", "/api/v1/x/broadcasts/1nGeLMWbkpaKX", "x_broadcast.json", map[string]stateWant{
			"broadcast.id":             wantText("1nGeLMWbkpaKX"),
			"broadcast.title":          wantText("USSF-259 Mission"),
			"broadcast.state":          wantText("Ended"),
			"broadcast.ended_at":       wantText("2026-09-17T01:17:35.036Z"),
			"broadcast.total_watched":  wantNumber(161758),
			"broadcast.total_watching": wantNumber(21639),
			"broadcast.width":          wantNumber(3840),
			"broadcast.creator.name":   wantText("SpaceX"),
			"broadcast.tweet.text":     wantText("USSF-259 Mission"),
		}},
		{"x-twitter-scraper_x_tweet", "2031735300000000000", "/api/v1/x/tweets/2031735300000000000", "x_tweet.json", map[string]stateWant{
			"tweet.id":                                   wantText("2031735300000000000"),
			"tweet.quoted_tweet.quoted_tweet.text":       wantText("Quoted quote"),
			"tweet.quoted_tweet.retweeted_tweet.text":    wantText("Quoted retweet"),
			"tweet.retweeted_tweet.quoted_tweet.text":    wantText("Retweeted quote"),
			"tweet.retweeted_tweet.retweeted_tweet.text": wantText("Retweeted retweet"),
		}},
	} {
		t.Run(test.typeName, func(t *testing.T) {
			t.Parallel()

			fixture, err := dataSourceFixtures.ReadFile(path.Join("testdata", "data_sources", test.fixture))
			if err != nil {
				t.Fatal(err)
			}
			var requested string
			client := xtwitterscraper.NewClient(
				option.WithAPIKey("test-api-key"),
				option.WithBaseURL("https://example.invalid/api/v1/"),
				option.WithMaxRetries(0),
				option.WithMiddleware(func(request *http.Request, _ option.MiddlewareNext) (*http.Response, error) {
					requested = request.Method + " " + request.URL.Path
					response := httptest.NewRecorder()
					response.Header().Set("Content-Type", "application/json")
					_, _ = response.Write(fixture)
					return response.Result(), nil
				}),
			)

			state := readDataSource(t, dataSourceNamed(t, test.typeName), &client, test.id)

			assertEqual(t, "request", requested, "GET "+test.path)
			for attributePath, want := range test.want {
				want.check(t, attributePath, stateAt(t, state, attributePath))
			}
		})
	}
}

func TestDataSourceSchemasHaveNoDynamicAttributesInCollections(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	providerUnderTest := testProvider(t)
	for _, factory := range providerUnderTest.DataSources(ctx) {
		dataSourceUnderTest := factory()
		var metadata datasource.MetadataResponse
		dataSourceUnderTest.Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: "x-twitter-scraper"}, &metadata)
		t.Run(metadata.TypeName, func(t *testing.T) {
			t.Parallel()

			var schemaResponse datasource.SchemaResponse
			dataSourceUnderTest.Schema(ctx, datasource.SchemaRequest{}, &schemaResponse)
			if schemaResponse.Diagnostics.HasError() {
				t.Fatalf("data source schema failed with %v", schemaResponse.Diagnostics)
			}

			var walk func(tftypes.Type, string, bool)
			walk = func(attributeType tftypes.Type, attributePath string, inCollection bool) {
				if inCollection && attributeType.Is(tftypes.DynamicPseudoType) {
					t.Errorf("%s has a dynamic attribute inside a collection", attributePath)
				}
				switch attributeType := attributeType.(type) {
				case tftypes.Object:
					for name, childType := range attributeType.AttributeTypes {
						walk(childType, attributePath+"."+name, inCollection)
					}
				case tftypes.List:
					walk(attributeType.ElementType, attributePath+"[]", true)
				case tftypes.Set:
					walk(attributeType.ElementType, attributePath+"[]", true)
				case tftypes.Map:
					walk(attributeType.ElementType, attributePath+"[]", true)
				}
			}
			walk(schemaResponse.Schema.Type().TerraformType(ctx), metadata.TypeName, false)
		})
	}
}

func testProvider(t *testing.T) *XTwitterScraperProvider {
	t.Helper()
	providerUnderTest, ok := NewProvider("test")().(*XTwitterScraperProvider)
	if !ok {
		t.Fatal("NewProvider did not return *XTwitterScraperProvider")
	}
	return providerUnderTest
}

func dataSourceNamed(t *testing.T, typeName string) datasource.DataSource {
	t.Helper()
	providerUnderTest := testProvider(t)
	for _, factory := range providerUnderTest.DataSources(context.Background()) {
		dataSourceUnderTest := factory()
		var metadata datasource.MetadataResponse
		dataSourceUnderTest.Metadata(context.Background(), datasource.MetadataRequest{ProviderTypeName: "x-twitter-scraper"}, &metadata)
		if metadata.TypeName == typeName {
			return dataSourceUnderTest
		}
	}
	t.Fatalf("the provider has no data source %s", typeName)
	return nil
}

func readDataSource(t *testing.T, dataSourceUnderTest datasource.DataSource, client *xtwitterscraper.Client, id string) tftypes.Value {
	t.Helper()
	ctx := context.Background()
	var schemaResponse datasource.SchemaResponse
	dataSourceUnderTest.Schema(ctx, datasource.SchemaRequest{}, &schemaResponse)
	var configureResponse datasource.ConfigureResponse
	configurable, ok := dataSourceUnderTest.(datasource.DataSourceWithConfigure)
	if !ok {
		t.Fatal("data source does not accept a configured client")
	}
	configurable.Configure(ctx, datasource.ConfigureRequest{ProviderData: client}, &configureResponse)
	if configureResponse.Diagnostics.HasError() {
		t.Fatalf("configure data source: %v", configureResponse.Diagnostics)
	}
	var readResponse datasource.ReadResponse
	readResponse.State.Schema = schemaResponse.Schema
	dataSourceUnderTest.Read(ctx, datasource.ReadRequest{
		Config: tfsdk.Config{Raw: rawObject(schemaResponse.Schema.Type(), map[string]string{"id": id}), Schema: schemaResponse.Schema},
	}, &readResponse)
	if readResponse.Diagnostics.HasError() {
		t.Fatalf("read data source: %v", readResponse.Diagnostics)
	}
	stateType := schemaResponse.Schema.Type().TerraformType(ctx)
	encoded, err := tfprotov6.NewDynamicValue(stateType, readResponse.State.Raw)
	if err != nil {
		t.Fatalf("encode data source state failed with %v", err)
	}
	state, err := encoded.Unmarshal(stateType)
	if err != nil {
		t.Fatalf("decode data source state failed with %v", err)
	}
	return state
}

type stateWant interface {
	check(t *testing.T, attributePath string, value tftypes.Value)
}

type wantText string

func (want wantText) check(t *testing.T, attributePath string, value tftypes.Value) {
	t.Helper()
	var text string
	if err := value.As(&text); err != nil || text != string(want) {
		t.Fatalf("%s = %q (%v), want %q", attributePath, text, err, string(want))
	}
}

type wantBool bool

func (want wantBool) check(t *testing.T, attributePath string, value tftypes.Value) {
	t.Helper()
	var flag bool
	if err := value.As(&flag); err != nil || flag != bool(want) {
		t.Fatalf("%s = %t (%v), want %t", attributePath, flag, err, bool(want))
	}
}

type wantNumber float64

func (want wantNumber) check(t *testing.T, attributePath string, value tftypes.Value) {
	t.Helper()
	number := new(big.Float)
	err := value.As(&number)
	if expected := big.NewFloat(float64(want)).String(); err != nil || number.String() != expected {
		t.Fatalf("%s = %s (%v), want %s", attributePath, number.String(), err, expected)
	}
}

type wantTweetJSON struct {
	text string
	urls []string
}

func (want wantTweetJSON) check(t *testing.T, attributePath string, value tftypes.Value) {
	t.Helper()
	var text string
	err := value.As(&text)
	var tweet struct {
		Text     string `json:"text"`
		Entities struct {
			URLs []struct {
				URL string `json:"url"`
			} `json:"urls"`
		} `json:"entities"`
	}
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(text))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&tweet)
	}
	urls := []string{}
	for _, url := range tweet.Entities.URLs {
		urls = append(urls, url.URL)
	}
	if err != nil || tweet.Text != want.text || !slices.Equal(urls, want.urls) {
		t.Fatalf("%s = %s (%v), want text %q and URLs %v", attributePath, text, err, want.text, want.urls)
	}
}

func stateAt(t *testing.T, state tftypes.Value, attributePath string) tftypes.Value {
	t.Helper()
	value := state
	for _, step := range strings.Split(attributePath, ".") {
		if index, err := strconv.Atoi(step); err == nil {
			var elements []tftypes.Value
			if err := value.As(&elements); err != nil || index >= len(elements) {
				t.Fatalf("%s is not in the state: %v", attributePath, err)
			}
			value = elements[index]
			continue
		}
		var attributes map[string]tftypes.Value
		if err := value.As(&attributes); err != nil {
			t.Fatalf("%s is not in the state: %v", attributePath, err)
		}
		child, ok := attributes[step]
		if !ok {
			t.Fatalf("%s is not in the state", attributePath)
		}
		value = child
	}
	return value
}
