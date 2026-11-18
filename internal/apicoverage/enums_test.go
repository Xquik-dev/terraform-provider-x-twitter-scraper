// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"errors"
	"net/http"
	"path/filepath"
	"strings"
	"testing"
)

const enumSchemas = `"components": {"schemas": {
    "EventType": {"type": "string", "description": "Account events.", "enum": ["tweet.new", "follower.added", "follower.removed"]},
    "PostEventType": {"type": "string", "enum": ["post.edited", "post.deleted"]},
    "EventTypeArray": {"type": "array", "items": {"$ref": "#/components/schemas/EventType"}},
    "Limit": {"type": "integer", "enum": [10, 100]},
    "Anything": true
  }},
  "paths": {`

func enumSpec() string {
	return strings.Replace(coveredSpec, `"paths": {`, enumSchemas, 1)
}

func runEnumCheck(t *testing.T, root, source string, enums map[string][]string) (int, string, string) {
	t.Helper()
	var out, failures strings.Builder
	code := Run(source, root, enums, http.DefaultClient, &out, &failures)
	return code, out.String(), failures.String()
}

func TestMatchingEnumsPass(t *testing.T) {
	enums := map[string][]string{
		"PostEventType": {"post.deleted", "post.edited"},
		"EventType":     {"tweet.new", "follower.added", "follower.removed"},
	}
	code, out, failures := runEnumCheck(t, writeProvider(t, docsExclusion), writeSpec(t, enumSpec()), enums)
	assertRun(t, code, out, failures, 0, coveredSummary+
		"The provider takes all 3 EventType values.\n"+
		"The provider takes all 2 PostEventType values.\n", "")
}

func TestEnumDifferencesFail(t *testing.T) {
	enums := map[string][]string{
		"PostEventType": {"post.edited", "post.deleted"},
		"Limit":         {"10", "100"},
		"EventType":     {"tweet.new", "follower.gained"},
		"Missing":       {"x"},
	}
	code, out, failures := runEnumCheck(t, writeProvider(t, docsExclusion), writeSpec(t, enumSpec()), enums)
	assertRun(t, code, out, failures, 1, "", `EventType "follower.added" is in the API but not in the provider.
EventType "follower.removed" is in the API but not in the provider.
EventType "follower.gained" is in the provider but not in the API.
The API spec has no Limit enum of strings in components.schemas.
The API spec has no Missing enum of strings in components.schemas.
Run go generate ./internal/apienum and ./scripts/generate-docs, then rerun.
`)
}

func TestOperationAndEnumDifferencesFailTogether(t *testing.T) {
	spec := strings.Replace(enumSpec(), `"paths": {`, `"paths": {"/api/v1/x/trends": {"get": {}},`, 1)
	code, out, failures := runEnumCheck(t, writeProvider(t, docsExclusion), writeSpec(t, spec), map[string][]string{"PostEventType": {"post.edited"}})
	assertRun(t, code, out, failures, 1, "", `GET /api/v1/x/trends is in the API but not in the provider.
Add a resource or data source for each operation above, or list it in excluded-operations.json with a reason.
PostEventType "post.deleted" is in the API but not in the provider.
Run go generate ./internal/apienum and ./scripts/generate-docs, then rerun.
`)
}

func TestGenerateEnumsWritesSpecValues(t *testing.T) {
	source, err := GenerateEnums(writeSpec(t, enumSpec()), http.DefaultClient, []string{"EventType", "PostEventType"})
	want := `// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// go generate ./internal/apienum writes this file from the enums in https://xquik.com/openapi.json. Change cmd/api-enums, not this file.

package apienum

var EventType = []string{
	"tweet.new",
	"follower.added",
	"follower.removed",
}

var PostEventType = []string{
	"post.edited",
	"post.deleted",
}

var Spec = map[string][]string{
	"EventType":     EventType,
	"PostEventType": PostEventType,
}
`
	if err != nil || string(source) != want {
		t.Fatalf("GenerateEnums() = %q, %v\nwant %q", source, err, want)
	}
}

func TestGenerateEnumsFails(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing.json")
	for _, test := range []struct {
		name, source string
		names        []string
		want         string
	}{
		{"unreadable spec", missing, []string{"EventType"}, "Could not read the API spec from " + missing + ". Check network access, or set XQUIK_OPENAPI_SPEC to a local copy of https://xquik.com/openapi.json."},
		{"missing enum", writeSpec(t, enumSpec()), []string{"EventType", "Missing"}, "The API spec has no Missing enum of strings in components.schemas."},
		{"numeric enum", writeSpec(t, enumSpec()), []string{"Limit"}, "The API spec has no Limit enum of strings in components.schemas."},
		{"schema without enum", writeSpec(t, enumSpec()), []string{"EventTypeArray"}, "The API spec has no EventTypeArray enum of strings in components.schemas."},
		{"boolean schema", writeSpec(t, enumSpec()), []string{"Anything"}, "The API spec has no Anything enum of strings in components.schemas."},
	} {
		t.Run(test.name, func(t *testing.T) {
			source, err := GenerateEnums(test.source, http.DefaultClient, test.names)
			if source != nil || err == nil || Sentence(err) != test.want {
				t.Fatalf("GenerateEnums() = %q, %v\nwant error %q", source, err, test.want)
			}
		})
	}
}

func TestGenerateEnumsReadsPublicSpecByDefault(t *testing.T) {
	var requested string
	client := &http.Client{Transport: roundTripper(func(request *http.Request) (*http.Response, error) {
		requested = request.URL.String()
		return nil, errors.New("offline")
	})}
	if _, err := GenerateEnums("", client, []string{"EventType"}); err == nil || requested != "https://xquik.com/openapi.json" {
		t.Fatalf("requested = %q, err = %v", requested, err)
	}
}
