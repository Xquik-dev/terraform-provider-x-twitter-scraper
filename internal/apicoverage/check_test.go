// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestCoveredSpecPasses(t *testing.T) {
	root := writeProvider(t, docsExclusion)
	for _, spec := range []string{
		coveredSpec,
		strings.Replace(coveredSpec, `"servers": [{"url": "https://xquik.com"}],`, "", 1),
	} {
		code, out, failures := runCheck(t, root, writeSpec(t, spec))
		assertRun(t, code, out, failures, 0, coveredSummary, "")
	}
}

func TestServerPathPrefixesSpecPaths(t *testing.T) {
	spec := `{"servers": [{"url": "https://xquik.com/api/v1/"}], "paths": {"/x/media": {"post": {}}, "/guest-wallets": {"post": {}}}}`
	code, out, failures := runCheck(t, writeProvider(t, "[]"), writeSpec(t, spec))
	assertRun(t, code, out, failures, 0, "The provider covers 2 of 2 API operations, and excluded-operations.json lists the other 0.\n", "")
}

func TestMissingOperationFails(t *testing.T) {
	spec := strings.Replace(coveredSpec, `"paths": {`, `"paths": {
    "/api/v1/x/lists/{id}": {"get": {"operationId": "getXList"}},
    "/api/v1/x/jobs/{id}": {"get": {}},`, 1)
	code, out, failures := runCheck(t, writeProvider(t, docsExclusion), writeSpec(t, spec))
	assertRun(t, code, out, failures, 1, "", "GET /api/v1/x/jobs/{id} is in the API but not in the provider.\n"+
		"GET /api/v1/x/lists/{id} (getXList) is in the API but not in the provider.\n"+
		"Add a resource or data source for each operation above, or list it in excluded-operations.json with a reason.\n")
}

func TestStaleExclusionsFail(t *testing.T) {
	exclusions := `[
  {"reason": "Docs search lives outside the SDK base URL.", "operations": ["GET /ask", "POST /ask"]},
  {"reason": "Waits for the Go SDK.", "operations": ["GET /api/v1/old", "DELETE /api/v1/x/tweets/{tweetId}/like"]}
]`
	code, out, failures := runCheck(t, writeProvider(t, exclusions), writeSpec(t, coveredSpec))
	assertRun(t, code, out, failures, 1, "", "excluded-operations.json lists DELETE /api/v1/x/tweets/{tweetId}/like, which the provider already implements. Remove it.\n"+
		"excluded-operations.json lists GET /api/v1/old, which the API no longer has. Remove it.\n")
}

func TestMissingAndStaleFailTogether(t *testing.T) {
	exclusions := `[{"reason": "Docs search lives outside the SDK base URL.", "operations": ["GET /ask", "POST /ask", "GET /api/v1/old"]}]`
	spec := strings.Replace(coveredSpec, `"paths": {`, `"paths": {"/api/v1/x/trends": {"get": {}},`, 1)
	code, out, failures := runCheck(t, writeProvider(t, exclusions), writeSpec(t, spec))
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
	code, out, failures := runCheck(t, writeProvider(t, exclusions), writeSpec(t, spec))
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
			code, out, failures := runCheck(t, writeProvider(t, test.content), writeSpec(t, coveredSpec))
			assertRun(t, code, out, failures, 1, "", "excluded-operations.json is not valid. "+test.want+"\n")
		})
	}
}

func TestMissingExclusionFileFails(t *testing.T) {
	files := providerTree("[]")
	delete(files, exclusionFile)
	code, out, failures := runCheck(t, writeModules(t, sdkTree(), files), writeSpec(t, coveredSpec))
	assertRun(t, code, out, failures, 1, "", "excluded-operations.json is not valid. The repository root has no readable copy. Create it, even as an empty list [].\n")
}

func TestUnreadableSpecFails(t *testing.T) {
	failing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer failing.Close()
	root := writeProvider(t, docsExclusion)
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
			code, out, failures := runCheck(t, root, source)
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
	code, out, failures := runCheck(t, writeProvider(t, docsExclusion), server.URL)
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
	code := Run("", writeProvider(t, docsExclusion), client, &out, &failures)
	if code != 1 || requested != "https://xquik.com/openapi.json" || !strings.HasPrefix(failures.String(), "Could not read the API spec from https://xquik.com/openapi.json.") {
		t.Fatalf("code = %d, requested = %q, failures = %q", code, requested, failures.String())
	}
}

func TestProviderThatCannotLoadFails(t *testing.T) {
	withoutSDK := providerTree(docsExclusion)
	withoutSDK["go.mod"] = "module example.com/provider\n\ngo 1.22\n"
	broken := providerTree(docsExclusion)
	broken["internal/services/raw/broken.go"] = "package raw\n\nfunc broken() { undefined() }\n"
	noServices := map[string]string{"go.mod": providerModule, exclusionFile: docsExclusion}
	emptyServices := map[string]string{"go.mod": providerModule, exclusionFile: docsExclusion, "internal/services/README.md": "No Go here.\n"}
	for _, test := range []struct {
		name     string
		provider map[string]string
		want     string
	}{
		{"go.mod without the Go SDK", withoutSDK, "The coverage check could not find the source of github.com/Xquik-dev/x-twitter-scraper-go. Run go mod download, then rerun.\n"},
		{"no services folder", noServices, "The coverage check found no Go packages under internal/services. Run it from the repository root.\n"},
		{"no Go packages", emptyServices, "The coverage check found no Go packages under internal/services. Run it from the repository root.\n"},
		{"compile error", broken, "internal/services/raw/broken.go:3:17: undefined: undefined\nThe coverage check could not compile the Go packages under internal/services. Fix the errors above, then rerun.\n"},
	} {
		t.Run(test.name, func(t *testing.T) {
			code, out, failures := runCheck(t, writeModules(t, sdkTree(), test.provider), writeSpec(t, coveredSpec))
			if code != 1 || out != "" || !strings.HasSuffix(failures, test.want) {
				t.Fatalf("code = %d, out = %q, failures = %q\nwant failures ending with %q", code, out, failures, test.want)
			}
		})
	}
}

func TestSDKWithoutReadableBaseURLFails(t *testing.T) {
	files := sdkTree()
	files["option/requestoption.go"] = strings.Replace(sdkOption, `requestconfig.WithDefaultBaseURL("https://xquik.com/api/v1/")`, "requestconfig.WithDefaultBaseURL(base)", 1) + "\nvar base = \"https://xquik.com/api/v1/\"\n"
	code, out, failures := runCheck(t, writeModules(t, files, providerTree(docsExclusion)), writeSpec(t, coveredSpec))
	assertRun(t, code, out, failures, 1, "", "The coverage check could not read the default base URL from option.WithEnvironmentProduction in the Go SDK. Check the Go SDK version in go.mod.\n")
}
