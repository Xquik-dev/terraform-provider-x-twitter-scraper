package logging

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

func TestLogRequestOmitsSensitiveData(t *testing.T) {
	t.Parallel()

	const body = `{"password":"credential-value"}`
	req := httptest.NewRequest(http.MethodPost, "https://example.com/v1/wallets?token=query-value", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer credential-value")
	req.Header.Set("Idempotency-Key", "idempotency-value")

	var output bytes.Buffer
	ctx := tflogtest.RootLogger(context.Background(), &output)
	LogRequest(ctx, req)

	assertBodyUnchanged(t, req.Body, body)
	assertLogExcludes(t, output.String(), "credential-value", "idempotency-value", "query-value", "authorization", "password")
	assertLogFields(t, &output, map[string]any{
		"@message": "HTTP request",
		"method":   http.MethodPost,
		"path":     "/v1/wallets",
		"protocol": "HTTP/1.1",
	})
}

func TestLogResponseOmitsSensitiveData(t *testing.T) {
	t.Parallel()

	const body = `{"api_key":"credential-value"}`
	req := httptest.NewRequest(http.MethodPost, "https://example.com/v1/wallets", nil)
	resp := &http.Response{
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     http.Header{"Set-Cookie": []string{"session=credential-value"}},
		Proto:      "HTTP/1.1",
		Request:    req,
		StatusCode: http.StatusCreated,
	}

	var output bytes.Buffer
	ctx := tflogtest.RootLogger(context.Background(), &output)
	LogResponse(ctx, resp)

	assertBodyUnchanged(t, resp.Body, body)
	assertLogExcludes(t, output.String(), "credential-value", "set-cookie", "api_key")
	assertLogFields(t, &output, map[string]any{
		"@message":    "HTTP response",
		"method":      http.MethodPost,
		"path":        "/v1/wallets",
		"protocol":    "HTTP/1.1",
		"status_code": float64(http.StatusCreated),
	})
}

func assertBodyUnchanged(t *testing.T, body io.Reader, expected string) {
	t.Helper()

	actual, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(actual) != expected {
		t.Fatalf("body = %q, want %q", actual, expected)
	}
}

func assertLogExcludes(t *testing.T, output string, excluded ...string) {
	t.Helper()

	for _, value := range excluded {
		if strings.Contains(strings.ToLower(output), strings.ToLower(value)) {
			t.Errorf("log output contains excluded value %q", value)
		}
	}
}

func assertLogFields(t *testing.T, output *bytes.Buffer, expected map[string]any) {
	t.Helper()

	entries, err := tflogtest.MultilineJSONDecode(output)
	if err != nil {
		t.Fatalf("MultilineJSONDecode() error = %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("log entry count = %d, want 1", len(entries))
	}
	for key, expectedValue := range expected {
		if actual := entries[0][key]; actual != expectedValue {
			t.Errorf("log field %q = %#v, want %#v", key, actual, expectedValue)
		}
	}
}
