// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type failingBody struct{}

func (failingBody) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func (failingBody) Close() error {
	return nil
}

func TestRedactHeader(t *testing.T) {
	tests := map[string]string{
		"Authorization":   "[REDACTED]",
		"Cookie":          "[REDACTED]",
		"Idempotency-Key": "[REDACTED]",
		"Set-Cookie":      "[REDACTED]",
		"X-API-Key":       "[REDACTED]",
		"Content-Type":    "application/json",
	}
	for name, want := range tests {
		value := "secret"
		if name == "Content-Type" {
			value = want
		}
		if got := redactHeader(name, value); got != want {
			t.Errorf("redactHeader(%q) = %q, want %q", name, got, want)
		}
	}
}

func TestRedactJSONBodyRecursively(t *testing.T) {
	body := []byte(`{"api_key":"key","nested":{"accessToken":"token","safe":"value"},"items":[{"password":"password"}]}`)
	want := `{"api_key":"[REDACTED]","items":[{"password":"[REDACTED]"}],"nested":{"accessToken":"[REDACTED]","safe":"value"}}`
	if got := redactJSONBody(body); got != want {
		t.Errorf("redactJSONBody() = %q, want %q", got, want)
	}
}

func TestRedactJSONBodyOmitsNonJSON(t *testing.T) {
	const body = "plain text"
	if got := redactJSONBody([]byte(body)); got != omittedBody {
		t.Errorf("redactJSONBody() = %q, want %q", got, omittedBody)
	}
}

func TestRedactJSONBodyOmitsEmptyBody(t *testing.T) {
	if got := redactJSONBody(nil); got != omittedBody {
		t.Errorf("redactJSONBody() = %q, want %q", got, omittedBody)
	}
}

func TestLogRequestAndResponsePreserveBodies(t *testing.T) {
	t.Parallel()

	request, err := http.NewRequest(http.MethodPost, "https://example.test/search", strings.NewReader(`{"query":"tweets"}`))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Authorization", "secret")
	if err := LogRequest(context.Background(), request); err != nil {
		t.Fatal(err)
	}
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(requestBody) != `{"query":"tweets"}` {
		t.Fatalf("request body changed: %q", requestBody)
	}

	response := &http.Response{
		Proto:  "HTTP/1.1",
		Status: "200 OK",
		Header: http.Header{"Set-Cookie": []string{"secret"}},
		Body:   io.NopCloser(strings.NewReader(`{"token":"secret"}`)),
	}
	if err := LogResponse(context.Background(), response); err != nil {
		t.Fatal(err)
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(responseBody) != `{"token":"secret"}` {
		t.Fatalf("response body changed: %q", responseBody)
	}
}

func TestLoggingPropagatesBodyReadFailures(t *testing.T) {
	t.Parallel()

	request, err := http.NewRequest(http.MethodPost, "https://example.test/search", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Body = failingBody{}
	if err := LogRequest(context.Background(), request); err == nil {
		t.Fatal("LogRequest accepted an unreadable body")
	}

	response := &http.Response{Body: failingBody{}}
	if err := LogResponse(context.Background(), response); err == nil {
		t.Fatal("LogResponse accepted an unreadable body")
	}
}

func TestMiddlewarePropagatesLoggingFailures(t *testing.T) {
	t.Parallel()

	request, err := http.NewRequest(http.MethodPost, "https://example.test/search", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Body = failingBody{}
	nextCalled := false
	_, err = Middleware(context.Background())(request, func(*http.Request) (*http.Response, error) {
		nextCalled = true
		return nil, nil
	})
	if err == nil {
		t.Fatal("middleware accepted an unreadable request")
	}
	if nextCalled {
		t.Fatal("middleware called next after request logging failed")
	}

	cleanRequest, err := http.NewRequest(http.MethodGet, "https://example.test/search", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = Middleware(context.Background())(cleanRequest, func(*http.Request) (*http.Response, error) {
		return &http.Response{Body: failingBody{}}, nil
	})
	if err == nil {
		t.Fatal("middleware accepted an unreadable response")
	}
}
