// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/option"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func Middleware(ctx context.Context) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		if req != nil {
			if err := LogRequest(ctx, req); err != nil {
				return nil, err
			}
		}

		resp, err := next(req)

		if resp != nil {
			if err := LogResponse(ctx, resp); err != nil {
				return nil, err
			}
		}

		return resp, err
	}
}

func LogRequest(ctx context.Context, req *http.Request) error {
	lines := []string{fmt.Sprintf("\n%s %s %s", req.Method, req.URL.Path, req.Proto)}

	// Log headers
	for name, values := range req.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("> %s: %s", strings.ToLower(name), redactHeader(name, value)))
		}
	}

	if req.Body != nil {
		// Read the body without mutating the original response
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}

		// Restore the original body to the response so it can be read again
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Log the body
		lines = append(lines, ">\n", redactJSONBody(bodyBytes), "\n")
	}

	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}

func LogResponse(ctx context.Context, resp *http.Response) error {
	// Log the status code
	lines := []string{fmt.Sprintf("\n< %s %s", resp.Proto, resp.Status)}

	// Log headers
	for name, values := range resp.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("< %s: %s", strings.ToLower(name), redactHeader(name, value)))
		}
	}

	// Read the body without mutating the original response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Restore the original body to the response so it can be read again
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	lines = append(lines, "<\n", redactJSONBody(bodyBytes), "\n")

	// Log the body
	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}

var sensitiveHeaders = map[string]bool{
	"authorization":       true,
	"cookie":              true,
	"idempotency-key":     true,
	"proxy-authorization": true,
	"set-cookie":          true,
	"x-api-key":           true,
	"x-auth-token":        true,
	"xquik-api-key":       true,
}

var sensitiveJSONKeys = map[string]bool{
	"accesstoken":    true,
	"apikey":         true,
	"authorization":  true,
	"clientsecret":   true,
	"credential":     true,
	"idempotencykey": true,
	"password":       true,
	"privatekey":     true,
	"refreshtoken":   true,
	"secret":         true,
	"sessiontoken":   true,
	"token":          true,
}

const omittedBody = "[OMITTED]"

func redactHeader(name string, value string) string {
	if sensitiveHeaders[strings.ToLower(name)] {
		return "[REDACTED]"
	}
	return value
}

func redactJSONBody(body []byte) string {
	var decoded any
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	if err := decoder.Decode(&decoded); err != nil {
		return omittedBody
	}
	redactJSONValue(decoded)
	redacted, err := json.Marshal(decoded)
	if err != nil {
		return omittedBody
	}
	return string(redacted)
}

func redactJSONValue(value any) {
	switch typed := value.(type) {
	case map[string]any:
		for key, child := range typed {
			normalized := strings.NewReplacer("_", "", "-", "").Replace(strings.ToLower(key))
			if sensitiveJSONKeys[normalized] {
				typed[key] = "[REDACTED]"
				continue
			}
			redactJSONValue(child)
		}
	case []any:
		for _, child := range typed {
			redactJSONValue(child)
		}
	}
}
