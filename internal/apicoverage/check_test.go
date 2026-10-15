// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"bytes"
	"errors"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
)

const sdkClient = `package xtwitterscraper

import "github.com/Xquik-dev/x-twitter-scraper-go/option"

type Client struct {
	options []option.RequestOption
	X       XService
	Drafts  *DraftService
	Raw     RawClient
}

type Params = map[string]any

func NewClient() (r Client) {
	return
}

func (r *Client) Get(ctx context.Context, path string, params any, res any) error {
	return r.Execute(ctx, http.MethodGet, path, params, res)
}
`

const sdkServices = `package xtwitterscraper

type XService struct {
	Tweets XTweetService
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
	requestconfig.ExecuteNewRequest(ctx)
	return requestconfig.ExecuteNewRequest(ctx, method, "x/tweets/translation", nil, nil)
}

func (r DraftService) Delete(ctx context.Context, id string) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, fmt.Sprintf("drafts/%s", id), nil, nil)
}

func (r RawClient) Fetch(ctx context.Context) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "raw", nil, nil)
}
`

const sdkOption = `package option

import "github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"

func (o RequestOption) WithEnvironmentProduction() RequestOption {
	return requestconfig.WithDefaultBaseURL("https://example.com/other/")
}

func WithMaxRetries(retries int) RequestOption {
	return requestconfig.WithMaxRetries(retries)
}

func WithEnvironmentProduction() RequestOption {
	return requestconfig.WithDefaultBaseURL("https://xquik.com/api/v1/")
}
`

const tweetDataSource = `package x_tweet

import (
	"context"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

type XTweetDataSource struct {
	client *xtwitterscraper.Client
}

func (d *XTweetDataSource) Configure(ctx context.Context, client func() *xtwitterscraper.Client) {
	d.client = client()
}

func (d *XTweetDataSource) Read(ctx context.Context) {
	defer func() {}()
	_, err := d.client.X.Tweets.Get(ctx, "1")
	resp.Diagnostics.AddError("failed", err.Error())
}
`

const draftResource = `package draft

import (
	"context"

	sdk "github.com/Xquik-dev/x-twitter-scraper-go"
)

func create(ctx context.Context, client *sdk.Client) {
	client.X.Tweets.New(ctx)
	client.Drafts.Delete(ctx, "1")
}
`

const rawResource = `package raw

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

type operation struct {
	action       string
	method       string
	pathTemplate string
}

var operations = []operation{
	{action: "like", method: http.MethodPost, pathTemplate: "x/tweets/%s/like"},
}

var unlike = &operation{action: "unlike", method: http.MethodDelete, pathTemplate: "x/tweets/%s/like"}

var client xtwitterscraper.Client

func (o operation) path(id string) string {
	return fmt.Sprintf(o.pathTemplate, id)
}

func (op operation) send(ctx context.Context, id string) {
	client.Execute(ctx, op.method, op.path(id), nil, nil)
}

func send(ctx context.Context, id string) {
	client.Post(ctx, "guest-wallets", nil, nil)
	client.Get(ctx, fmt.Sprintf("x/write-actions/%s", id), nil, nil)
	path := fmt.Sprintf("monitors/keywords/%s", id)
	client.Patch(ctx, path, nil, nil)
	client.Delete(ctx, keywordsPath, nil, nil)
	client.Put(ctx, stylePath(id), nil, nil)
	client.Execute(ctx, http.MethodPost, "x/media", nil, nil)
}
`

const rawPaths = `package raw

import "fmt"

const keywordsPath = "monitors/keywords"

const (
	first = iota
	second
)

const limit = 3

func stylePath(id string) string {
	return fmt.Sprintf("styles/%s", id)
}

func external(id string) string
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

const unreadableRequest = " sends a request whose method or path the coverage check cannot read. Use an http.Method constant and a path string, a fmt.Sprintf format, a string constant or a 1 line function that returns 1 of those.\n"

func sdkFiles() fstest.MapFS {
	return fstest.MapFS{
		"client.go":               {Data: []byte(sdkClient)},
		"x.go":                    {Data: []byte(sdkServices)},
		"x_test.go":               {Data: []byte("package xtwitterscraper\n\nfunc (r *XTweetService) Get() {}\n")},
		"README.md":               {Data: []byte("# SDK\n")},
		"internal/skip.go":        {Data: []byte("package internal\n")},
		"option/requestoption.go": {Data: []byte(sdkOption)},
	}
}

func providerFiles(exclusions string) fstest.MapFS {
	return fstest.MapFS{
		exclusionFile:                                {Data: []byte(exclusions)},
		"internal/provider.go":                       {Data: []byte("package internal\n\nimport \"github.com/Xquik-dev/x-twitter-scraper-go\"\n\nfunc configure(client *xtwitterscraper.Client) { client.Get(ctx, path, nil, nil) }\n")},
		"internal/services/x_tweet/data_source.go":   {Data: []byte(tweetDataSource)},
		"internal/services/draft/resource.go":        {Data: []byte(draftResource)},
		"internal/services/raw/resource.go":          {Data: []byte(rawResource)},
		"internal/services/raw/paths.go":             {Data: []byte(rawPaths)},
		"internal/services/raw/resource_test.go":     {Data: []byte("package raw\n\nimport \"github.com/Xquik-dev/x-twitter-scraper-go\"\n\nfunc TestSend(client *xtwitterscraper.Client) { client.Get(ctx, path, nil, nil) }\n")},
		"internal/services/raw/testdata/fixture.txt": {Data: []byte("not Go\n")},
	}
}

func writeSpec(t *testing.T, content string) string {
	t.Helper()
	file := filepath.Join(t.TempDir(), "openapi.json")
	if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return file
}

func runWith(t *testing.T, provider, sdkSource fs.FS, source string) (int, string, string) {
	t.Helper()
	var out, failures bytes.Buffer
	code := Run(source, provider, sdkSource, http.DefaultClient, &out, &failures)
	return code, out.String(), failures.String()
}

func runCheck(t *testing.T, provider fs.FS, source string) (int, string, string) {
	t.Helper()
	return runWith(t, provider, sdkFiles(), source)
}

func assertRun(t *testing.T, code int, out, failures string, wantCode int, wantOut, wantFailures string) {
	t.Helper()
	if code != wantCode || out != wantOut || failures != wantFailures {
		t.Fatalf("code = %d, out = %q, failures = %q\nwant code = %d, out = %q, failures = %q", code, out, failures, wantCode, wantOut, wantFailures)
	}
}

func TestCoveredSpecPasses(t *testing.T) {
	for _, spec := range []string{
		coveredSpec,
		strings.Replace(coveredSpec, `"servers": [{"url": "https://xquik.com"}],`, "", 1),
	} {
		code, out, failures := runCheck(t, providerFiles(docsExclusion), writeSpec(t, spec))
		assertRun(t, code, out, failures, 0, coveredSummary, "")
	}
}

func TestServerPathPrefixesSpecPaths(t *testing.T) {
	spec := `{"servers": [{"url": "https://xquik.com/api/v1/"}], "paths": {"/x/media": {"post": {}}, "/guest-wallets": {"post": {}}}}`
	code, out, failures := runCheck(t, providerFiles("[]"), writeSpec(t, spec))
	assertRun(t, code, out, failures, 0, "The provider covers 2 of 2 API operations, and excluded-operations.json lists the other 0.\n", "")
}

func TestMissingOperationFails(t *testing.T) {
	spec := strings.Replace(coveredSpec, `"paths": {`, `"paths": {
    "/api/v1/x/lists/{id}": {"get": {"operationId": "getXList"}},
    "/api/v1/x/jobs/{id}": {"get": {}},`, 1)
	code, out, failures := runCheck(t, providerFiles(docsExclusion), writeSpec(t, spec))
	assertRun(t, code, out, failures, 1, "", "GET /api/v1/x/jobs/{id} is in the API but not in the provider.\n"+
		"GET /api/v1/x/lists/{id} (getXList) is in the API but not in the provider.\n"+
		"Add a resource or data source for each operation above, or list it in excluded-operations.json with a reason.\n")
}

func TestStaleExclusionsFail(t *testing.T) {
	exclusions := `[
  {"reason": "Docs search lives outside the SDK base URL.", "operations": ["GET /ask", "POST /ask"]},
  {"reason": "Waits for the Go SDK.", "operations": ["GET /api/v1/old", "DELETE /api/v1/x/tweets/{tweetId}/like"]}
]`
	code, out, failures := runCheck(t, providerFiles(exclusions), writeSpec(t, coveredSpec))
	assertRun(t, code, out, failures, 1, "", "excluded-operations.json lists DELETE /api/v1/x/tweets/{tweetId}/like, which the provider already implements. Remove it.\n"+
		"excluded-operations.json lists GET /api/v1/old, which the API no longer has. Remove it.\n")
}

func TestMissingAndStaleFailTogether(t *testing.T) {
	exclusions := `[{"reason": "Docs search lives outside the SDK base URL.", "operations": ["GET /ask", "POST /ask", "GET /api/v1/old"]}]`
	spec := strings.Replace(coveredSpec, `"paths": {`, `"paths": {"/api/v1/x/trends": {"get": {}},`, 1)
	code, out, failures := runCheck(t, providerFiles(exclusions), writeSpec(t, spec))
	assertRun(t, code, out, failures, 1, "", "GET /api/v1/x/trends is in the API but not in the provider.\n"+
		"Add a resource or data source for each operation above, or list it in excluded-operations.json with a reason.\n"+
		"excluded-operations.json lists GET /api/v1/old, which the API no longer has. Remove it.\n")
}

func TestPathParameterNamesDoNotMatter(t *testing.T) {
	spec := `{"servers": [{"url": "https://xquik.com"}], "paths": {
    "/api/v1/x/tweets/{a}": {"get": {}},
    "/api/v1/drafts/{draftId}": {"delete": {}},
    "/api/v1/x/tweets/{b}/like": {"post": {}},
    "/api/v1/x/users/{userId}": {"get": {}}
  }}`
	exclusions := `[{"reason": "Waits for the Go SDK.", "operations": ["GET /api/v1/x/users/{id}"]}]`
	code, out, failures := runCheck(t, providerFiles(exclusions), writeSpec(t, spec))
	assertRun(t, code, out, failures, 0, "The provider covers 3 of 4 API operations, and excluded-operations.json lists the other 1.\n", "")
}

func TestInvalidExclusionFilesFail(t *testing.T) {
	shape := "It must be a JSON list of groups, each with only a reason string and an operations list of strings."
	notOperation := ", which is not an upper case HTTP method, a space and a path starting with /."
	for _, test := range []struct{ name, content, want string }{
		{"not JSON", "[}", "It is not valid JSON near byte 2."},
		{"cut short", "[{", shape},
		{"not a list", `{"reason": "x"}`, shape},
		{"unknown field", `[{"reason": "x", "operations": ["GET /x"], "note": "y"}]`, shape},
		{"two values", `[] []`, "It holds more than 1 JSON value."},
		{"missing reason", `[{"operations": ["GET /x"]}]`, "Group 1 needs a non-empty reason."},
		{"blank reason", `[{"reason": " ", "operations": ["GET /x"]}]`, "Group 1 needs a non-empty reason."},
		{"missing operations", `[{"reason": "x"}]`, "Group 1 needs at least 1 operation."},
		{"empty operations", `[{"reason": "x", "operations": ["GET /x"]}, {"reason": "y", "operations": []}]`, "Group 2 needs at least 1 operation."},
		{"lower case method", `[{"reason": "x", "operations": ["get /x"]}]`, `Group 1 lists "get /x"` + notOperation},
		{"unknown method", `[{"reason": "x", "operations": ["FETCH /x"]}]`, `Group 1 lists "FETCH /x"` + notOperation},
		{"relative path", `[{"reason": "x", "operations": ["GET x"]}]`, `Group 1 lists "GET x"` + notOperation},
		{"duplicate", `[{"reason": "x", "operations": ["GET /x/{id}"]}, {"reason": "y", "operations": ["GET /x/{alias}"]}]`, "It lists GET /x/{id} and GET /x/{alias}, which name the same operation. List it once."},
	} {
		t.Run(test.name, func(t *testing.T) {
			code, out, failures := runCheck(t, providerFiles(test.content), writeSpec(t, coveredSpec))
			assertRun(t, code, out, failures, 1, "", "excluded-operations.json is not valid. "+test.want+"\n")
		})
	}
}

func TestMissingExclusionFileFails(t *testing.T) {
	files := providerFiles("[]")
	delete(files, exclusionFile)
	code, out, failures := runCheck(t, files, writeSpec(t, coveredSpec))
	assertRun(t, code, out, failures, 1, "", "excluded-operations.json is not valid. The repository root has no readable copy. Create it, even as an empty list [].\n")
}

func TestUnreadableSpecFails(t *testing.T) {
	failing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer failing.Close()
	for _, source := range []string{
		filepath.Join(t.TempDir(), "missing.json"),
		writeSpec(t, "openapi: 3.1.0\n"),
		writeSpec(t, `{"paths": {}}`),
		writeSpec(t, `{"paths": {"/x": []}}`),
		writeSpec(t, `{"paths": {"/x": {"get": {"servers": [{"url": "https://[::1"}]}}}}`),
		failing.URL,
		"http://127.0.0.1:0/openapi.json",
		"https://exa mple.com/openapi.json",
	} {
		t.Run(source, func(t *testing.T) {
			code, out, failures := runCheck(t, providerFiles(docsExclusion), source)
			assertRun(t, code, out, failures, 1, "", "Could not read the API spec from "+source+". Check network access, or set XQUIK_OPENAPI_SPEC to a local copy of https://xquik.com/openapi.json.\n")
		})
	}
}

func TestSpecFromURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write([]byte(coveredSpec)); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	code, out, failures := runCheck(t, providerFiles(docsExclusion), server.URL)
	assertRun(t, code, out, failures, 0, coveredSummary, "")
}

type roundTripper func(*http.Request) (*http.Response, error)

func (f roundTripper) RoundTrip(request *http.Request) (*http.Response, error) { return f(request) }

func TestEmptySourceUsesPublicSpec(t *testing.T) {
	var requested string
	client := &http.Client{Transport: roundTripper(func(request *http.Request) (*http.Response, error) {
		requested = request.URL.String()
		return nil, errors.New("offline")
	})}
	var out, failures bytes.Buffer
	code := Run("", providerFiles(docsExclusion), sdkFiles(), client, &out, &failures)
	if code != 1 || requested != "https://xquik.com/openapi.json" || !strings.HasPrefix(failures.String(), "Could not read the API spec from https://xquik.com/openapi.json.") {
		t.Fatalf("code = %d, requested = %q, failures = %q", code, requested, failures.String())
	}
}

const badCallHeader = `package bad

import "github.com/Xquik-dev/x-twitter-scraper-go"

type operation struct{ method, pathTemplate string }

func helper(id string) string {
	id = strings.TrimSpace(id)
	return fmt.Sprintf("x/%s", id)
}

func stop(id string) string { panic(id) }

func pair(id string) (string, error) { return "x", nil }

func call(ctx context.Context, client *xtwitterscraper.Client, op operation, other operation, method, path string) {
	`

const badPaths = `package bad

const shadowed = "x"

func inner(id string) string { return fmt.Sprintf("x/%s", id) }

func outer(id string) string { return inner(id) }

func twice(id string) string {
	return fmt.Sprintf("x/%s", id)
	return "y"
}

func external(id string) string
`

func TestUnreadableCallsFail(t *testing.T) {
	const file = "internal/services/bad/bad.go"
	noSDKMethod := ", which is no service method of the Go SDK in go.mod. Call a Go SDK service method, or update the Go SDK in go.mod.\n"
	noRequest := ", which sends no request the coverage check can read. Call a Go SDK service method or Get, Post, Put, Patch, Delete or Execute.\n"
	badLiteral := " lists a request without a readable method and pathTemplate. Set method to an http.Method constant and pathTemplate to a path string or fmt.Sprintf format.\n"
	for _, test := range []struct{ name, body, want string }{
		{"unknown service", `client.Missing.Get(ctx)`, ":17 calls Client.Missing.Get" + noSDKMethod},
		{"unknown service method", `client.X.Tweets.Search(ctx)`, ":17 calls Client.X.Tweets.Search" + noSDKMethod},
		{"SDK method outside a service", `client.Raw.Fetch(ctx)`, ":17 calls Client.Raw.Fetch" + noSDKMethod},
		{"SDK method without a request", `client.X.Tweets.GetAutoPaging(ctx, "1")`, ":17 calls Client.X.Tweets.GetAutoPaging, which sends no request through ExecuteNewRequest. Call the Go SDK method that sends the request.\n"},
		{"unreadable SDK request", `client.X.Tweets.Translate(ctx, "GET")`, ":17 calls Client.X.Tweets.Translate, whose request the coverage check cannot read in the Go SDK at x.go:23 and x.go:24. Fix that request in the Go SDK first.\n"},
		{"unknown client method", `client.Options(ctx)`, ":17 calls Client.Options" + noRequest},
		{"execute without a path", `client.Execute(ctx, http.MethodGet)`, ":17 calls Client.Execute" + noRequest},
		{"raw call without a path", `client.Get(ctx)`, ":17" + unreadableRequest},
		{"path from a parameter", `client.Get(ctx, path, nil, nil)`, ":17" + unreadableRequest},
		{"path assigned twice", "p := \"x\"\n\tp = \"y\"\n\tclient.Get(ctx, p, nil, nil)", ":19" + unreadableRequest},
		{"path assigned from a helper", "p := helper(\"1\")\n\tclient.Get(ctx, p, nil, nil)", ":18" + unreadableRequest},
		{"path from a 2 line helper", `client.Get(ctx, helper("1"), nil, nil)`, ":17" + unreadableRequest},
		{"path from a nested helper", `client.Get(ctx, outer("1"), nil, nil)`, ":17" + unreadableRequest},
		{"path from a helper with 2 returns", `client.Get(ctx, twice("1"), nil, nil)`, ":17" + unreadableRequest},
		{"path from a function without a body", `client.Get(ctx, external("1"), nil, nil)`, ":17" + unreadableRequest},
		{"path from a local that shadows a constant", "shadowed := \"a\"\n\tshadowed = \"b\"\n\tclient.Get(ctx, shadowed, nil, nil)", ":19" + unreadableRequest},
		{"path from a helper without a return", `client.Get(ctx, stop("1"), nil, nil)`, ":17" + unreadableRequest},
		{"path from a helper with 2 results", `client.Get(ctx, pair("1"), nil, nil)`, ":17" + unreadableRequest},
		{"path from an unknown function", `client.Get(ctx, build(), nil, nil)`, ":17" + unreadableRequest},
		{"path from another package", `client.Get(ctx, paths.Item("1"), nil, nil)`, ":17" + unreadableRequest},
		{"path from strings", `client.Get(ctx, strings.Join(parts, "/"), nil, nil)`, ":17" + unreadableRequest},
		{"format from a variable", `client.Get(ctx, fmt.Sprintf(format, id), nil, nil)`, ":17" + unreadableRequest},
		{"format with too few values", `client.Get(ctx, fmt.Sprintf("x/%s/%s", id), nil, nil)`, ":17" + unreadableRequest},
		{"path not a string", `client.Get(ctx, 42, nil, nil)`, ":17" + unreadableRequest},
		{"execute with a variable method", `client.Execute(ctx, method, "x", nil, nil)`, ":17" + unreadableRequest},
		{"execute with a method name", `client.Execute(ctx, "GET", "x", nil, nil)`, ":17" + unreadableRequest},
		{"execute with an unreadable path", `client.Execute(ctx, http.MethodGet, path, nil, nil)`, ":17" + unreadableRequest},
		{"table method from another field", `client.Execute(ctx, op.verb, op.path("1"), nil, nil)`, ":17" + unreadableRequest},
		{"table method from a call", `client.Execute(ctx, ops().method, op.path("1"), nil, nil)`, ":17" + unreadableRequest},
		{"table path from another value", `client.Execute(ctx, op.method, other.path("1"), nil, nil)`, ":17" + unreadableRequest},
		{"table path from a nested field", `client.Execute(ctx, op.method, op.paths.item, nil, nil)`, ":17" + unreadableRequest},
		{"table without literals", `client.Execute(ctx, op.method, op.pathTemplate, nil, nil)`, ":17" + unreadableRequest},
		{"table value from a local", "local := op\n\tclient.Execute(ctx, local.method, local.path(\"1\"), nil, nil)", ":18" + unreadableRequest},
		{"table literal with a method name", "table := []operation{{method: \"GET\", pathTemplate: \"x/%s\"}, {method: http.MethodGet, pathTemplate: \"y\"}}\n\tclient.Execute(ctx, op.method, op.path(\"1\"), nil, nil)", ":17" + badLiteral},
		{"table literal with a path variable", "table := []operation{{method: http.MethodGet, pathTemplate: name}}\n\tclient.Execute(ctx, op.method, op.path(\"1\"), nil, nil)", ":17" + badLiteral},
		{"table literal with an escaped percent", "table := []operation{{method: http.MethodGet, pathTemplate: \"x/%d%%\"}}\n\tclient.Execute(ctx, op.method, op.path(\"1\"), nil, nil)", ":17" + badLiteral},
		{"table literal with an explicit type", "table := []operation{operation{method: \"GET\", pathTemplate: \"x\"}}\n\tclient.Execute(ctx, op.method, op.path(\"1\"), nil, nil)", ":17" + badLiteral},
		{"table literal without keys", "table := []operation{{http.MethodGet, \"x\"}}\n\tclient.Execute(ctx, op.method, op.path(\"1\"), nil, nil)", ":17" + badLiteral},
	} {
		t.Run(test.name, func(t *testing.T) {
			files := providerFiles(docsExclusion)
			files[file] = &fstest.MapFile{Data: []byte(badCallHeader + test.body + "\n}\n")}
			files["internal/services/bad/paths.go"] = &fstest.MapFile{Data: []byte(badPaths)}
			code, out, failures := runCheck(t, files, writeSpec(t, coveredSpec))
			assertRun(t, code, out, failures, 1, "", file+test.want)
		})
	}
}

func TestEveryUnreadableCallIsReported(t *testing.T) {
	files := providerFiles(docsExclusion)
	files["internal/services/bad/bad.go"] = &fstest.MapFile{Data: []byte(badCallHeader + "client.Get(ctx, path, nil, nil)\n\tclient.Missing.Get(ctx)\n}\n")}
	code, out, failures := runCheck(t, files, writeSpec(t, coveredSpec))
	assertRun(t, code, out, failures, 1, "", "internal/services/bad/bad.go:17"+unreadableRequest+
		"internal/services/bad/bad.go:18 calls Client.Missing.Get, which is no service method of the Go SDK in go.mod. Call a Go SDK service method, or update the Go SDK in go.mod.\n")
}

type brokenFS struct {
	files  fs.FS
	broken string
}

func (b brokenFS) Open(name string) (fs.File, error) {
	if name == b.broken {
		return nil, fs.ErrPermission
	}
	return b.files.Open(name)
}

func TestUnreadableSourcesFail(t *testing.T) {
	sdkProblem := "The coverage check could not read the Go SDK source. Run go mod download, then rerun."
	baseURLProblem := "The coverage check could not read the default base URL from option.WithEnvironmentProduction in the Go SDK. Check the Go SDK version in go.mod."
	providerProblem := "The coverage check could not read the provider source under internal/services. Run go vet ./... and fix the error, then rerun from the repository root."
	withSDKFile := func(name, content string) fstest.MapFS {
		files := sdkFiles()
		files[name] = &fstest.MapFile{Data: []byte(content)}
		return files
	}
	withoutSDKFile := func(name string) fstest.MapFS {
		files := sdkFiles()
		delete(files, name)
		return files
	}
	withProviderFile := func(name, content string) fstest.MapFS {
		files := providerFiles(docsExclusion)
		files[name] = &fstest.MapFile{Data: []byte(content)}
		return files
	}
	noServices := fstest.MapFS{exclusionFile: {Data: []byte(docsExclusion)}}
	for _, test := range []struct {
		name      string
		provider  fs.FS
		sdkSource fs.FS
		want      string
	}{
		{"empty SDK", providerFiles(docsExclusion), fstest.MapFS{}, sdkProblem},
		{"broken SDK file", providerFiles(docsExclusion), withSDKFile("client.go", "package sdk\nfunc ("), sdkProblem},
		{"unreadable SDK file", providerFiles(docsExclusion), brokenFS{files: sdkFiles(), broken: "x.go"}, sdkProblem},
		{"SDK without options", providerFiles(docsExclusion), withoutSDKFile("option/requestoption.go"), baseURLProblem},
		{"broken option file", providerFiles(docsExclusion), withSDKFile("option/requestoption.go", "package option\nfunc ("), baseURLProblem},
		{"base URL from a variable", providerFiles(docsExclusion), withSDKFile("option/requestoption.go", "package option\n\nfunc WithEnvironmentProduction() RequestOption {\n\treturn requestconfig.WithDefaultBaseURL(base)\n}\n"), baseURLProblem},
		{"broken base URL", providerFiles(docsExclusion), withSDKFile("option/requestoption.go", "package option\n\nfunc WithEnvironmentProduction() RequestOption {\n\treturn requestconfig.WithDefaultBaseURL(\"://xquik.com\")\n}\n"), baseURLProblem},
		{"provider without services", noServices, sdkFiles(), providerProblem},
		{"broken provider file", withProviderFile("internal/services/raw/broken.go", "package raw\nfunc ("), sdkFiles(), providerProblem},
		{"unreadable provider file", brokenFS{files: providerFiles(docsExclusion), broken: "internal/services/raw/paths.go"}, sdkFiles(), providerProblem},
	} {
		t.Run(test.name, func(t *testing.T) {
			code, out, failures := runWith(t, test.provider, test.sdkSource, writeSpec(t, coveredSpec))
			assertRun(t, code, out, failures, 1, "", test.want+"\n")
		})
	}
}
