// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"strings"
	"testing"
)

const badHeader = `package bad

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	xtwitterscraper "github.com/Xquik-dev/x-twitter-scraper-go"
)

`

const badFooter = `
var (
	_ = context.Background
	_ = fmt.Sprintf
	_ = http.MethodGet
	_ = strings.ToLower
	_ = xtwitterscraper.NewClient
)
`

const unreadableRequest = " sends a request whose method or path the coverage check cannot read. Use an http.Method constant and a constant path, a fmt.Sprintf format, a variable assigned once to 1 of those or a 1 line function that returns 1 of those."

const badLiteral = " lists a request without a constant method and pathTemplate. Set method to an http.Method constant and pathTemplate to a path string or format."

var unreadableCalls = []struct {
	name  string
	code  string
	wants []string
}{
	{"method_value", `func methodValue(ctx context.Context, client *xtwitterscraper.Client) {
	get := client.X.Tweets.Get
	_, _ = get(ctx, "1")
}
`, []string{":13 uses XTweetService.Get without calling it. Call it directly so the coverage check can read its request."}},
	{"method_expression", `func methodExpression() {
	newTweet := (*xtwitterscraper.XTweetService).New
	_ = newTweet
}
`, []string{":13 uses XTweetService.New without calling it. Call it directly so the coverage check can read its request."}},
	{"no_request", `func noRequest(ctx context.Context, client *xtwitterscraper.Client) {
	client.X.Tweets.GetAutoPaging(ctx, "1")
}
`, []string{":13 calls XTweetService.GetAutoPaging, which sends no request through ExecuteNewRequest. Call the Go SDK method that sends the request."}},
	{"unreadable_sdk_request", `func unreadableSDK(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.X.Tweets.Translate(ctx, "GET")
}
`, []string{":13 calls XTweetService.Translate, whose request the coverage check cannot read in the Go SDK at x.go:54 and x.go:56. Fix that request in the Go SDK first."}},
	{"other_client_method", `func otherClientMethod(client *xtwitterscraper.Client) error {
	return client.Close()
}
`, []string{":13 calls Client.Close, which sends no request the coverage check can read. Call a Go SDK service method or Get, Post, Put, Patch, Delete or Execute."}},
	{"parameter_path", `func parameterPath(ctx context.Context, client *xtwitterscraper.Client, path string) error {
	return client.Get(ctx, path, nil, nil)
}
`, []string{":13" + unreadableRequest}},
	{"assigned_twice", `func assignedTwice(ctx context.Context, client *xtwitterscraper.Client) error {
	path := "x"
	path = "y"
	return client.Get(ctx, path, nil, nil)
}
`, []string{":15" + unreadableRequest}},
	{"multiple_values", `func pathAndError() (string, error) { return "x", nil }

func multipleValues(ctx context.Context, client *xtwitterscraper.Client) error {
	path, err := pathAndError()
	if err != nil {
		return err
	}
	return client.Get(ctx, path, nil, nil)
}
`, []string{":19" + unreadableRequest}},
	{"declared_then_assigned", `func declaredThenAssigned(ctx context.Context, client *xtwitterscraper.Client) error {
	var path string
	path = "x"
	return client.Get(ctx, path, nil, nil)
}
`, []string{":15" + unreadableRequest}},
	{"address_taken", `func fill(path *string) { *path = "y" }

func addressTaken(ctx context.Context, client *xtwitterscraper.Client) error {
	path := "x"
	fill(&path)
	return client.Get(ctx, path, nil, nil)
}
`, []string{":17" + unreadableRequest}},
	{"range_value", `func rangeValue(ctx context.Context, client *xtwitterscraper.Client) error {
	for _, path := range []string{"x"} {
		if path == "" {
			path = "y"
		}
		return client.Get(ctx, path, nil, nil)
	}
	return nil
}
`, []string{":17" + unreadableRequest}},
	{"range_key", `func rangeKey(ctx context.Context, client *xtwitterscraper.Client) error {
	for path := range map[string]bool{"x": true} {
		if path == "" {
			path = "y"
		}
		return client.Get(ctx, path, nil, nil)
	}
	return nil
}
`, []string{":17" + unreadableRequest}},
	{"two_line_helper", `func twoLines(id string) string {
	id = strings.TrimSpace(id)
	return fmt.Sprintf("x/%s", id)
}

func twoLineHelper(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Get(ctx, twoLines("1"), nil, nil)
}
`, []string{":18" + unreadableRequest}},
	{"nested_helper", `func innerPath(id string) string { return fmt.Sprintf("x/%s", id) }

func outerPath(id string) string { return innerPath(id) }

func nestedHelper(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Get(ctx, outerPath("1"), nil, nil)
}
`, []string{":17" + unreadableRequest}},
	{"helper_without_return", `func panicPath(id string) string { panic(id) }

func helperWithoutReturn(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Get(ctx, panicPath("1"), nil, nil)
}
`, []string{":15" + unreadableRequest}},
	{"helper_bare_return", `func barePath() (path string) { return }

func helperBareReturn(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Get(ctx, barePath(), nil, nil)
}
`, []string{":15" + unreadableRequest}},
	{"method_helper", `type pathHolder struct{}

func (pathHolder) path() string { return "x" }

func methodHelper(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Get(ctx, pathHolder{}.path(), nil, nil)
}
`, []string{":17" + unreadableRequest}},
	{"other_package", `func otherPackage(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Get(ctx, strings.ToLower("X"), nil, nil)
}
`, []string{":13" + unreadableRequest}},
	{"function_variable", `func functionVariable(ctx context.Context, client *xtwitterscraper.Client) error {
	build := func() string { return "x" }
	return client.Get(ctx, build(), nil, nil)
}
`, []string{":14" + unreadableRequest}},
	{"variable_format", `func variableFormat(ctx context.Context, client *xtwitterscraper.Client, format string) error {
	return client.Get(ctx, fmt.Sprintf(format, "1"), nil, nil)
}
`, []string{":13" + unreadableRequest}},
	{"short_format", `func shortFormat(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Get(ctx, fmt.Sprintf("x/%s/%s", "1"), nil, nil)
}
`, []string{":13" + unreadableRequest}},
	{"spread_format", `func spreadFormat(ctx context.Context, client *xtwitterscraper.Client, values []any) error {
	return client.Get(ctx, fmt.Sprintf("x/%s", values...), nil, nil)
}
`, []string{":13" + unreadableRequest}},
	{"variable_method", `func variableMethod(ctx context.Context, client *xtwitterscraper.Client, method string) error {
	return client.Execute(ctx, method, "x", nil, nil)
}
`, []string{":13" + unreadableRequest}},
	{"unknown_method", `func unknownMethod(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Execute(ctx, "FETCH", "x", nil, nil)
}
`, []string{":13" + unreadableRequest}},
	{"execute_path", `func executePath(ctx context.Context, client *xtwitterscraper.Client, path string) error {
	return client.Execute(ctx, http.MethodGet, path, nil, nil)
}
`, []string{":13" + unreadableRequest}},
	{"table_method_call", `type verbTable struct{ pathTemplate string }

func (t verbTable) verb() string { return http.MethodGet }

var verbs = []verbTable{{pathTemplate: "x"}}

func tableMethodCall(ctx context.Context, client *xtwitterscraper.Client, t verbTable) error {
	return client.Execute(ctx, t.verb(), t.pathTemplate, nil, nil)
}
`, []string{":19" + unreadableRequest}},
	{"table_other_value", `type pairTable struct{ method, pathTemplate string }

var pairs = []pairTable{{http.MethodGet, "x"}}

func tableOtherValue(ctx context.Context, client *xtwitterscraper.Client, first, second pairTable) error {
	return client.Execute(ctx, first.method, second.pathTemplate, nil, nil)
}
`, []string{":17" + unreadableRequest}},
	{"table_mixed_fields", `type mixedTable struct{ method, pathTemplate, fallback string }

var mixed = []mixedTable{{http.MethodGet, "x", "y"}}

func (t mixedTable) path() string {
	if t.fallback != "" {
		return t.fallback
	}
	return t.pathTemplate
}

func tableMixedFields(ctx context.Context, client *xtwitterscraper.Client, t mixedTable) error {
	return client.Execute(ctx, t.method, t.path(), nil, nil)
}
`, []string{":24" + unreadableRequest}},
	{"table_built_path", `type builtTable struct{ method, pathTemplate string }

var built = []builtTable{{http.MethodGet, "x"}}

func (t builtTable) path(id string) string { return "x/" + id }

func tableBuiltPath(ctx context.Context, client *xtwitterscraper.Client, t builtTable) error {
	return client.Execute(ctx, t.method, t.path("1"), nil, nil)
}
`, []string{":19" + unreadableRequest}},
	{"table_bare_return", `type bareTable struct{ method, pathTemplate string }

var bareRows = []bareTable{{http.MethodGet, "x"}}

func (t bareTable) path() (path string) {
	path = t.pathTemplate
	return
}

func tableBareReturn(ctx context.Context, client *xtwitterscraper.Client, t bareTable) error {
	return client.Execute(ctx, t.method, t.path(), nil, nil)
}
`, []string{":22" + unreadableRequest}},
	{"table_unnamed_receiver", `type unnamedTable struct{ method, pathTemplate string }

var unnamedRows = []unnamedTable{{http.MethodGet, "x"}}

func (unnamedTable) path() string { return "x" }

func tableUnnamedReceiver(ctx context.Context, client *xtwitterscraper.Client, t unnamedTable) error {
	return client.Execute(ctx, t.method, t.path(), nil, nil)
}
`, []string{":19" + unreadableRequest}},
	{"table_foreign_method", `type builderTable struct {
	method       string
	pathTemplate *strings.Builder
}

var builderRows = []builderTable{{method: http.MethodGet}}

func tableForeignMethod(ctx context.Context, client *xtwitterscraper.Client, t builderTable) error {
	return client.Execute(ctx, t.method, t.pathTemplate.String(), nil, nil)
}
`, []string{":20" + unreadableRequest}},
	{"table_function_path", `type funcTable struct{ method, pathTemplate string }

var funcRows = []funcTable{{http.MethodGet, "x"}}

func pathFor(t funcTable) string { return t.pathTemplate }

func tableFunctionPath(ctx context.Context, client *xtwitterscraper.Client, t funcTable) error {
	return client.Execute(ctx, t.method, pathFor(t), nil, nil)
}
`, []string{":19" + unreadableRequest}},
	{"table_literal_path", `type literalTable struct{ method, pathTemplate string }

var literalRows = []literalTable{{http.MethodGet, "x"}}

func tableLiteralPath(ctx context.Context, client *xtwitterscraper.Client, t literalTable) error {
	return client.Execute(ctx, t.method, "x", nil, nil)
}
`, []string{":17" + unreadableRequest}},
	{"table_closure_field", `type closureTable struct {
	method string
	build  func() string
}

var closureRows = []closureTable{{method: http.MethodGet}}

func tableClosureField(ctx context.Context, client *xtwitterscraper.Client, t closureTable) error {
	return client.Execute(ctx, t.method, t.build(), nil, nil)
}
`, []string{":20" + unreadableRequest}},
	{"table_promoted_field", `type innerRow struct{ method string }

type outerRow struct {
	innerRow
	pathTemplate string
}

var outerRows = []outerRow{{innerRow{http.MethodGet}, "x"}}

func tablePromotedField(ctx context.Context, client *xtwitterscraper.Client, t outerRow) error {
	return client.Execute(ctx, t.method, t.pathTemplate, nil, nil)
}
`, []string{":22" + unreadableRequest}},
	{"table_index_owner", `type indexTable struct{ method, pathTemplate string }

var indexRows = []indexTable{{http.MethodGet, "x"}}

func tableIndexOwner(ctx context.Context, client *xtwitterscraper.Client) error {
	return client.Execute(ctx, indexRows[0].method, indexRows[0].pathTemplate, nil, nil)
}
`, []string{":17" + unreadableRequest}},
	{"table_without_literals", `type emptyTable struct{ method, pathTemplate string }

func tableWithoutLiterals(ctx context.Context, client *xtwitterscraper.Client, t emptyTable) error {
	return client.Execute(ctx, t.method, t.pathTemplate, nil, nil)
}
`, []string{":15" + unreadableRequest}},
	{"table_bad_literals", `type badTable struct{ method, pathTemplate string }

var verb = http.MethodGet

var badRows = map[string]badTable{
	"variable": {method: verb, pathTemplate: "x"},
	"unknown":  {method: "FETCH", pathTemplate: "x"},
	"percent":  {method: http.MethodGet, pathTemplate: "x/%d%%"},
	"missing":  {method: http.MethodGet},
	"good":     {method: http.MethodGet, pathTemplate: "x/%s/good"},
}

func tableBadLiterals(ctx context.Context, client *xtwitterscraper.Client, t badTable) error {
	return client.Execute(ctx, t.method, t.pathTemplate, nil, nil)
}
`, []string{":17" + badLiteral, ":18" + badLiteral, ":19" + badLiteral, ":20" + badLiteral}},
}

func TestUnreadableCallsFail(t *testing.T) {
	files := providerTree(docsExclusion)
	var want []string
	for _, test := range unreadableCalls {
		name := "internal/services/bad/" + test.name + ".go"
		files[name] = badHeader + test.code + badFooter
		for _, line := range test.wants {
			want = append(want, name+line)
		}
	}
	code, out, failures := runCheck(t, writeModules(t, sdkTree(), files), writeSpec(t, coveredSpec))
	got := map[string]bool{}
	for _, line := range strings.Split(strings.TrimSuffix(failures, "\n"), "\n") {
		got[line] = true
	}
	for _, line := range want {
		if !got[line] {
			t.Errorf("the failures lack %q", line)
		}
	}
	if code != 1 || out != "" || len(got) != len(want) {
		t.Fatalf("code = %d, out = %q, %d failure lines, want %d\n%s", code, out, len(got), len(want), failures)
	}
}
