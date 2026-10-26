// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

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

func sdkFiles() fstest.MapFS {
	files := fstest.MapFS{}
	for name, content := range sdkTree() {
		files[name] = &fstest.MapFile{Data: []byte(content)}
	}
	return files
}

func TestReadSDKFollowsTheGoSDKRules(t *testing.T) {
	files := sdkFiles()
	files["option/other.go"] = &fstest.MapFile{Data: []byte(`package option

type Settings struct{}

func (s Settings) WithEnvironmentProduction() RequestOption {
	return requestconfig.WithDefaultBaseURL("https://example.com/other/")
}

func WithMaxRetries(retries int) RequestOption {
	return requestconfig.WithMaxRetries(retries)
}
`)}
	files["short.go"] = &fstest.MapFile{Data: []byte(`package xtwitterscraper

func (r *XShortService) Get(ctx context.Context, id string, format string) error {
	requestconfig.ExecuteNewRequest(ctx)
	path := "x"
	path = "y"
	requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil)
	requestconfig.ExecuteNewRequest(ctx, http.MethodGet, id, nil, nil)
	requestconfig.ExecuteNewRequest(ctx, http.MethodGet, build(), nil, nil)
	requestconfig.ExecuteNewRequest(ctx, http.MethodGet, fmt.Sprintf(format, id), nil, nil)
	return requestconfig.ExecuteNewRequest(ctx, "GET", "x", nil, nil)
}

func (r *PageService[T]) Next(ctx context.Context) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "pages/next", nil, nil)
}
`)}
	index, err := readSDK(files)
	if err != nil {
		t.Fatal(err)
	}
	if index.basePath != "/api/v1/" {
		t.Fatalf("basePath = %q, want /api/v1/", index.basePath)
	}
	want := map[string]map[string]sdkMethod{
		"XShortService": {"Get": {unreadable: []string{"short.go:4", "short.go:7", "short.go:8", "short.go:9", "short.go:10", "short.go:11"}}},
		"XTweetService": {
			"Get":           {requests: []request{{"GET", "x/tweets/{}"}}},
			"New":           {requests: []request{{"POST", "x/tweets"}}},
			"GetAutoPaging": {},
			"Translate":     {unreadable: []string{"x.go:54", "x.go:56"}},
		},
		"DraftService": {"Delete": {requests: []request{{"DELETE", "drafts/{}"}}}},
	}
	if !reflect.DeepEqual(index.methods, want) {
		t.Fatalf("methods = %#v\nwant %#v", index.methods, want)
	}
}

func TestReadSDKRejectsUnreadableSources(t *testing.T) {
	sdkProblem := "the coverage check could not read the Go SDK source. Run go mod download, then rerun"
	baseURLProblem := "the coverage check could not read the default base URL from option.WithEnvironmentProduction in the Go SDK. Check the Go SDK version in go.mod"
	with := func(name, content string) fstest.MapFS {
		files := sdkFiles()
		files[name] = &fstest.MapFile{Data: []byte(content)}
		return files
	}
	without := func(name string) fstest.MapFS {
		files := sdkFiles()
		delete(files, name)
		return files
	}
	for _, test := range []struct {
		name  string
		files fs.FS
		want  string
	}{
		{"empty SDK", fstest.MapFS{}, sdkProblem},
		{"broken SDK file", with("client.go", "package sdk\nfunc ("), sdkProblem},
		{"unreadable SDK file", brokenFS{files: sdkFiles(), broken: "x.go"}, sdkProblem},
		{"SDK without options", without("option/requestoption.go"), baseURLProblem},
		{"broken option file", with("option/requestoption.go", "package option\nfunc ("), baseURLProblem},
		{"broken base URL", with("option/requestoption.go", "package option\n\nfunc WithEnvironmentProduction() RequestOption {\n\treturn requestconfig.WithDefaultBaseURL(\"://xquik.com\")\n}\n"), baseURLProblem},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := readSDK(test.files)
			if err == nil || err.Error() != test.want {
				t.Fatalf("err = %v, want %q", err, test.want)
			}
		})
	}
}
