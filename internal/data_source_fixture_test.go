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
	"reflect"
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
		want     map[string]any
	}{
		{"x-twitter-scraper_x_list", "4143408", "/api/v1/x/lists/4143408", "x_list.json", map[string]any{
			"id":                    "4143408",
			"list.id":               "4143408",
			"list.name":             "NASA Centers",
			"list.description":      "NASA has 10 centers",
			"list.member_count":     int64(11),
			"list.subscriber_count": int64(859),
			"list.owner.id":         "11348282",
			"list.owner.username":   "NASA",
			"list.url":              "https://x.com/i/lists/4143408",
			"list.created_at":       "2009-12-02T22:41:02.000Z",
			"list.banner_url":       "https://pbs.twimg.com/media/example.png",
			"list.members_context":  "11 members",
		}},
		{"x-twitter-scraper_x_job", "1899202395210649861", "/api/v1/x/jobs/1899202395210649861", "x_job.json", map[string]any{
			"job.id":                    "1899202395210649861",
			"job.title":                 "Engineer",
			"job.location_type":         "remote",
			"job.employment_type":       "full_time",
			"job.seniority_level":       "lead",
			"job.salary_min":            float64(225000),
			"job.salary_currency_code":  "USD",
			"job.featured":              true,
			"job.description_markdown":  "# What is Commonware?",
			"job.account.username":      "commonwarexyz",
			"job.account.verified":      true,
			"job.account.verified_type": "Business",
		}},
		{"x-twitter-scraper_x_space", "1OGwblLjnQMKB", "/api/v1/x/spaces/1OGwblLjnQMKB", "x_space.json", map[string]any{
			"space.id":                   "1OGwblLjnQMKB",
			"space.title":                "Extending the Network to the Edge",
			"space.state":                "Ended",
			"space.started_at":           "2026-03-11T17:00:26.227Z",
			"space.total_live_listeners": int64(841),
			"space.total_replay_watched": int64(609),
			"space.creator.username":     "solana",
			"space.admins.0.id":          "951329744804392960",
			"space.admins.0.username":    "solana",
			"space.tweet.text":           "Extending the Network to the Edge",
			"space.sharings.0.tweet": map[string]any{
				"text":     "Join the Space",
				"entities": map[string]any{"urls": []any{}},
			},
			"space.sharings.1.tweet": map[string]any{
				"text":     "Replay the Space",
				"entities": map[string]any{"urls": []any{map[string]any{"url": "https://x.com/i/spaces/1OGwblLjnQMKB"}}},
			},
		}},
		{"x-twitter-scraper_x_user", "NASA", "/api/v1/x/users/NASA", "x_user.json", map[string]any{
			"username":                                "NASA",
			"followers":                               int64(92376187),
			"verified_type":                           "Government",
			"profile_picture_original":                "https://pbs.twimg.com/profile_images/1321163587679784960/0ZxKlEKB.jpg",
			"birthdate.year":                          int64(1958),
			"birthdate.month":                         int64(10),
			"contact_info.links.0.url":                "http://www.nasa.gov/",
			"contact_info.links.0.network":            "website",
			"verification_reason.entities.0.url_type": "ExternalUrl",
		}},
		{"x-twitter-scraper_x_broadcast", "1nGeLMWbkpaKX", "/api/v1/x/broadcasts/1nGeLMWbkpaKX", "x_broadcast.json", map[string]any{
			"broadcast.id":             "1nGeLMWbkpaKX",
			"broadcast.title":          "USSF-259 Mission",
			"broadcast.state":          "Ended",
			"broadcast.ended_at":       "2026-09-17T01:17:35.036Z",
			"broadcast.total_watched":  int64(161758),
			"broadcast.total_watching": int64(21639),
			"broadcast.width":          int64(3840),
			"broadcast.creator.name":   "SpaceX",
			"broadcast.tweet.text":     "USSF-259 Mission",
		}},
		{"x-twitter-scraper_x_tweet", "2031735300000000000", "/api/v1/x/tweets/2031735300000000000", "x_tweet.json", map[string]any{
			"tweet.id":                                   "2031735300000000000",
			"tweet.quoted_tweet.quoted_tweet.text":       "Quoted quote",
			"tweet.quoted_tweet.retweeted_tweet.text":    "Quoted retweet",
			"tweet.retweeted_tweet.quoted_tweet.text":    "Retweeted quote",
			"tweet.retweeted_tweet.retweeted_tweet.text": "Retweeted retweet",
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
				assertState(t, state, attributePath, want)
			}
		})
	}
}

func TestDataSourceSchemasHaveNoDynamicAttributesInCollections(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	providerUnderTest := NewProvider("test")().(*XTwitterScraperProvider)
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

func dataSourceNamed(t *testing.T, typeName string) datasource.DataSource {
	t.Helper()
	providerUnderTest := NewProvider("test")().(*XTwitterScraperProvider)
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
	dataSourceUnderTest.(datasource.DataSourceWithConfigure).Configure(ctx, datasource.ConfigureRequest{ProviderData: client}, &configureResponse)
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

func assertState(t *testing.T, state tftypes.Value, attributePath string, want any) {
	t.Helper()
	steps := tftypes.NewAttributePath()
	for _, step := range strings.Split(attributePath, ".") {
		if index, err := strconv.Atoi(step); err == nil {
			steps = steps.WithElementKeyInt(index)
		} else {
			steps = steps.WithAttributeName(step)
		}
	}
	found, _, err := tftypes.WalkAttributePath(state, steps)
	if err != nil {
		t.Fatalf("%s is not in the state: %v", attributePath, err)
	}
	value := found.(tftypes.Value)
	var got, expected any
	switch want := want.(type) {
	case string:
		var text string
		err = value.As(&text)
		got, expected = text, want
	case bool:
		var flag bool
		err = value.As(&flag)
		got, expected = flag, want
	case int64:
		number := new(big.Float)
		err = value.As(&number)
		got, expected = number.String(), big.NewFloat(float64(want)).String()
	case float64:
		number := new(big.Float)
		err = value.As(&number)
		got, expected = number.String(), big.NewFloat(want).String()
	case map[string]any:
		var text string
		err = value.As(&text)
		var object map[string]any
		if err == nil {
			err = json.Unmarshal([]byte(text), &object)
		}
		got, expected = object, want
	default:
		t.Fatalf("%s has an unhandled expected value type %T", attributePath, want)
	}
	if err != nil || !reflect.DeepEqual(got, expected) {
		t.Fatalf("%s = %v (%v), want %v", attributePath, got, err, expected)
	}
}
