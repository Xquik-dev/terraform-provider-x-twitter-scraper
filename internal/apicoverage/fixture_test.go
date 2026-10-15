// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

const sdkModule = `module github.com/Xquik-dev/x-twitter-scraper-go

go 1.22
`

const sdkRequestConfig = `package requestconfig

import "context"

type RequestConfig struct {
	BaseURL string
}

type RequestOption func(*RequestConfig)

func WithDefaultBaseURL(base string) RequestOption {
	return func(config *RequestConfig) { config.BaseURL = base }
}

func ExecuteNewRequest(ctx context.Context, method string, path string, params any, res any) error {
	return nil
}
`

const sdkOption = `package option

import "github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"

type RequestOption = requestconfig.RequestOption

func WithEnvironmentProduction() RequestOption {
	return requestconfig.WithDefaultBaseURL("https://xquik.com/api/v1/")
}
`

const sdkClient = `package xtwitterscraper

import (
	"context"
	"net/http"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
)

type Client struct {
	X      XService
	Drafts *DraftService
}

func NewClient() *Client {
	return &Client{Drafts: &DraftService{}}
}

func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any) error {
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res)
}

func (r *Client) Get(ctx context.Context, path string, params any, res any) error {
	return r.Execute(ctx, http.MethodGet, path, params, res)
}

func (r *Client) Post(ctx context.Context, path string, params any, res any) error {
	return r.Execute(ctx, http.MethodPost, path, params, res)
}

func (r *Client) Put(ctx context.Context, path string, params any, res any) error {
	return r.Execute(ctx, http.MethodPut, path, params, res)
}

func (r *Client) Patch(ctx context.Context, path string, params any, res any) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res)
}

func (r *Client) Delete(ctx context.Context, path string, params any, res any) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res)
}

func (r *Client) Close() error {
	return nil
}
`

const sdkServices = `package xtwitterscraper

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
)

type XService struct {
	Tweets XTweetService
}

type XTweetService struct{}

type DraftService struct{}

type Tweet struct {
	ID string
}

func (t Tweet) RawJSON() string {
	return t.ID
}

type Pager struct {
	Tweet *Tweet
}

func NewPager(tweet *Tweet, err error) *Pager {
	return &Pager{Tweet: tweet}
}

func (r *XTweetService) Get(ctx context.Context, id string) (res *Tweet, err error) {
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res)
	return res, err
}

func (r *XTweetService) New(ctx context.Context) (res *Tweet, err error) {
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, "x/tweets", nil, &res)
	return res, err
}

func (r *XTweetService) GetAutoPaging(ctx context.Context, id string) *Pager {
	return NewPager(r.Get(ctx, id))
}

func (r *XTweetService) Translate(ctx context.Context, method string) error {
	if method == "" {
		return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, strings.ToLower("X"), nil, nil)
	}
	return requestconfig.ExecuteNewRequest(ctx, method, "x/tweets/translation", nil, nil)
}

func (r DraftService) Delete(ctx context.Context, id string) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, fmt.Sprintf("drafts/%s", id), nil, nil)
}
`

const providerModule = `module example.com/provider

go 1.22

require github.com/Xquik-dev/x-twitter-scraper-go v0.0.0

replace github.com/Xquik-dev/x-twitter-scraper-go => ../sdk
`

const tweetDataSource = `package x_tweet

import (
	"context"

	xtwitterscraper "github.com/Xquik-dev/x-twitter-scraper-go"
)

type XTweetDataSource struct {
	client *xtwitterscraper.Client
}

func (d *XTweetDataSource) Read(ctx context.Context) (string, error) {
	tweet, err := d.client.X.Tweets.Get(ctx, "1")
	if err != nil {
		return "", err
	}
	return tweet.RawJSON(), nil
}
`

const draftResource = `package draft

import (
	"context"

	sdk "github.com/Xquik-dev/x-twitter-scraper-go"
)

type DraftResource struct {
	*sdk.Client
}

func NewDraftResource() *DraftResource {
	return &DraftResource{Client: sdk.NewClient()}
}

func (r *DraftResource) Create(ctx context.Context) error {
	tweets := r.X.Tweets
	if _, err := tweets.New(ctx); err != nil {
		return err
	}
	return r.Drafts.Delete(ctx, "1")
}
`

const rawResource = `package raw

import (
	"context"
	"fmt"
	"net/http"

	xtwitterscraper "github.com/Xquik-dev/x-twitter-scraper-go"
)

type operation struct {
	action       string
	method       string
	pathTemplate string
}

var operations = map[string]operation{
	"like":   {action: "like", method: http.MethodPost, pathTemplate: "x/tweets/%s/like"},
	"unlike": {"unlike", http.MethodDelete, "x/tweets/%s/like"},
}

func (o operation) path(id string) string {
	escape := func(value string) string { return value }
	if o.action == "" {
		return o.pathTemplate
	}
	return fmt.Sprintf(o.pathTemplate, escape(id))
}

func (o *operation) send(ctx context.Context, client *xtwitterscraper.Client, id string) error {
	return client.Execute(ctx, o.method, o.path(id), nil, nil)
}

var sendMedia = func(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Execute(ctx, http.MethodPost, "x/media", nil, nil)
}

func send(ctx context.Context, client *xtwitterscraper.Client, id string) error {
	alias := client
	if err := alias.Post(ctx, "guest-wallets?source=terraform", nil, nil); err != nil {
		return err
	}
	if err := client.Get(ctx, fmt.Sprintf("x/write-actions/%s", id), nil, nil); err != nil {
		return err
	}
	path := fmt.Sprintf(keywordsPath+"/%s", id)
	if err := client.Patch(ctx, path, nil, nil); err != nil {
		return err
	}
	if err := client.Delete(ctx, keywordsPath, nil, nil); err != nil {
		return err
	}
	if err := client.Put(ctx, stylePath(id), nil, nil); err != nil {
		return err
	}
	return sendMedia(ctx, client)
}
`

const rawPaths = `package raw

import "fmt"

const keywordsPath = "monitors/" + "keywords"

func stylePath(id string) string {
	return fmt.Sprintf("styles/%s", id)
}
`

const skippedCall = `package skipped

import (
	"context"

	xtwitterscraper "github.com/Xquik-dev/x-twitter-scraper-go"
)

func read(ctx context.Context, client *xtwitterscraper.Client, path string) error {
	return client.Get(ctx, path, nil, nil)
}
`

const coveredSpec = `{
  "servers": [{"url": "https://xquik.com"}],
  "paths": {
    "/api/v1/x/tweets/{tweetId}": {"get": {"operationId": "getTweet"}},
    "/api/v1/x/tweets": {"post": {"operationId": "createTweet"}},
    "/api/v1/drafts/{id}": {"delete": {}},
    "/api/v1/guest-wallets": {"post": {}},
    "/api/v1/x/write-actions/{id}": {"get": {}},
    "/api/v1/monitors/keywords/{id}": {"patch": {}},
    "/api/v1/monitors/keywords": {"delete": {}},
    "/api/v1/styles/{id}": {"put": {}},
    "/api/v1/x/media": {"servers": [], "post": {"servers": []}},
    "/api/v1/x/tweets/{id}/like": {"post": {}, "delete": {}},
    "/ask": {"servers": [{"url": "https://xquik.com/"}], "get": {"operationId": "askDocs"}, "post": {"servers": [{"url": "https://docs.xquik.com"}]}}
  }
}`

const docsExclusion = `[{"reason": "Docs search lives outside the SDK base URL.", "operations": ["GET /ask", "POST /ask"]}]`

const coveredSummary = "The provider covers 11 of 13 API operations, and excluded-operations.json lists the other 2.\n"

func sdkTree() map[string]string {
	return map[string]string{
		"go.mod": sdkModule,
		"internal/requestconfig/requestconfig.go": sdkRequestConfig,
		"option/requestoption.go":                 sdkOption,
		"client.go":                               sdkClient,
		"x.go":                                    sdkServices,
		"x_test.go":                               "package xtwitterscraper\n\nfunc (r *XTweetService) Search() {}\n",
		"README.md":                               "# SDK\n",
	}
}

func providerTree(exclusions string) map[string]string {
	return map[string]string{
		"go.mod":               providerModule,
		exclusionFile:          exclusions,
		"internal/provider.go": skippedCall,
		"internal/services/x_tweet/data_source.go": tweetDataSource,
		"internal/services/draft/resource.go":      draftResource,
		"internal/services/raw/resource.go":        rawResource,
		"internal/services/raw/paths.go":           rawPaths,
		"internal/services/raw/resource_test.go":   "package raw\n\nfunc broken() { undefined() }\n",
		"internal/services/raw/testdata/skip.go":   skippedCall,
		"internal/services/raw/notes.txt":          "Not Go.\n",
	}
}

func writeTree(t *testing.T, root string, files map[string]string) {
	t.Helper()
	for name, content := range files {
		file := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
	}
}

func writeModules(t *testing.T, sdkFiles, providerFiles map[string]string) string {
	t.Helper()
	root := t.TempDir()
	writeTree(t, filepath.Join(root, "sdk"), sdkFiles)
	writeTree(t, filepath.Join(root, "provider"), providerFiles)
	return filepath.Join(root, "provider")
}

func writeProvider(t *testing.T, exclusions string) string {
	t.Helper()
	return writeModules(t, sdkTree(), providerTree(exclusions))
}

func writeSpec(t *testing.T, content string) string {
	t.Helper()
	file := filepath.Join(t.TempDir(), "openapi.json")
	if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return file
}

func runCheck(t *testing.T, root, source string) (int, string, string) {
	t.Helper()
	var out, failures bytes.Buffer
	code := Run(source, root, http.DefaultClient, &out, &failures)
	return code, out.String(), failures.String()
}

func assertRun(t *testing.T, code int, out, failures string, wantCode int, wantOut, wantFailures string) {
	t.Helper()
	if code != wantCode || out != wantOut || failures != wantFailures {
		t.Fatalf("code = %d, out = %q, failures = %q\nwant code = %d, out = %q, failures = %q", code, out, failures, wantCode, wantOut, wantFailures)
	}
}
