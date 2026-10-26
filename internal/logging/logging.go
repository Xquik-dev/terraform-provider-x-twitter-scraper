// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package logging logs API requests and responses through the Terraform plugin
// log.
package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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

	for name, values := range req.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("> %s: %s", strings.ToLower(name), redactHeader(name, value)))
		}
	}

	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}

		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		lines = append(lines, ">\n", redactJSONBody(bodyBytes), "\n")
	}

	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}

func LogResponse(ctx context.Context, resp *http.Response) error {
	lines := []string{fmt.Sprintf("\n< %s %s", resp.Proto, resp.Status)}

	for name, values := range resp.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("< %s: %s", strings.ToLower(name), redactHeader(name, value)))
		}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	lines = append(lines, "<\n", redactJSONBody(bodyBytes), "\n")

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

var errEmptyJSON = errors.New("empty JSON body")

func redactHeader(name string, value string) string {
	if sensitiveHeaders[strings.ToLower(name)] {
		return "[REDACTED]"
	}
	return value
}

func redactJSONBody(body []byte) string {
	redacted, err := redactJSON(body)
	if err != nil {
		return omittedBody
	}
	return string(redacted)
}

func redactJSON(raw json.RawMessage) (json.RawMessage, error) {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 {
		return nil, errEmptyJSON
	}
	if trimmed[0] == '{' {
		var object map[string]json.RawMessage
		if err := json.Unmarshal(trimmed, &object); err != nil {
			return nil, err
		}
		for key, child := range object {
			normalized := strings.NewReplacer("_", "", "-", "").Replace(strings.ToLower(key))
			if sensitiveJSONKeys[normalized] {
				object[key] = json.RawMessage(`"[REDACTED]"`)
				continue
			}
			redactedChild, err := redactJSON(child)
			if err != nil {
				return nil, err
			}
			object[key] = redactedChild
		}
		return json.Marshal(object)
	}
	if trimmed[0] == '[' {
		var items []json.RawMessage
		if err := json.Unmarshal(trimmed, &items); err != nil {
			return nil, err
		}
		for index, child := range items {
			redactedChild, err := redactJSON(child)
			if err != nil {
				return nil, err
			}
			items[index] = redactedChild
		}
		return json.Marshal(items)
	}
	return json.Marshal(json.RawMessage(trimmed))
}
