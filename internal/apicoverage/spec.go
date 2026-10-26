// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	placeholderPattern = regexp.MustCompile(`\{[^{}]*\}`)
	operationPattern   = regexp.MustCompile(`^([A-Z]+) (/\S*)$`)
)

type operations map[string]string

type server struct {
	URL string `json:"url"`
}

type specOperation struct {
	OperationID string   `json:"operationId"`
	Servers     []server `json:"servers"`
}

type pathItem struct {
	Servers []server       `json:"servers"`
	Get     *specOperation `json:"get"`
	Put     *specOperation `json:"put"`
	Post    *specOperation `json:"post"`
	Delete  *specOperation `json:"delete"`
	Options *specOperation `json:"options"`
	Head    *specOperation `json:"head"`
	Patch   *specOperation `json:"patch"`
	Trace   *specOperation `json:"trace"`
}

func (p pathItem) methods() map[string]*specOperation {
	return map[string]*specOperation{
		"GET": p.Get, "PUT": p.Put, "POST": p.Post, "DELETE": p.Delete,
		"OPTIONS": p.Options, "HEAD": p.Head, "PATCH": p.Patch, "TRACE": p.Trace,
	}
}

type spec struct {
	Servers []server            `json:"servers"`
	Paths   map[string]pathItem `json:"paths"`
}

type exclusionGroup struct {
	Reason     string   `json:"reason"`
	Operations []string `json:"operations"`
}

func operationKey(method, operationPath string) string {
	return method + " " + placeholderPattern.ReplaceAllString(operationPath, "{}")
}

func joinPath(base, operationPath string) string {
	return strings.TrimSuffix(base, "/") + "/" + strings.TrimPrefix(operationPath, "/")
}

func serverPath(servers ...[]server) (string, error) {
	for _, candidates := range servers {
		if len(candidates) == 0 {
			continue
		}
		parsed, err := url.Parse(candidates[0].URL)
		if err != nil {
			return "", err
		}
		return parsed.Path, nil
	}
	return "", nil
}

func unreadableSpec(source string) error {
	return fmt.Errorf("could not read the API spec from %s. Check network access, or set %s to a local copy of %s", source, SpecVariable, defaultSpecSource)
}

func fetchSpec(source string, client *http.Client) ([]byte, error) {
	if !strings.HasPrefix(source, "https://") && !strings.HasPrefix(source, "http://") {
		return os.ReadFile(source)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, source, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	return io.ReadAll(response.Body)
}

func specOperations(source string, client *http.Client) (operations, error) {
	data, err := fetchSpec(source, client)
	if err != nil {
		return nil, unreadableSpec(source)
	}
	var document spec
	if err := json.Unmarshal(data, &document); err != nil || len(document.Paths) == 0 {
		return nil, unreadableSpec(source)
	}
	found := operations{}
	for key, item := range document.Paths {
		for method, op := range item.methods() {
			if op == nil {
				continue
			}
			base, err := serverPath(op.Servers, item.Servers, document.Servers)
			if err != nil {
				return nil, unreadableSpec(source)
			}
			label := method + " " + joinPath(base, key)
			if op.OperationID != "" {
				label += " (" + op.OperationID + ")"
			}
			found[operationKey(method, joinPath(base, key))] = label
		}
	}
	return found, nil
}

func invalidExclusions(detail string) error {
	return fmt.Errorf("the file %s is not valid. %s", exclusionFile, detail)
}

func readExclusions(root fs.FS) (operations, error) {
	data, err := fs.ReadFile(root, exclusionFile)
	if err != nil {
		return nil, invalidExclusions("The repository root has no readable copy. Create it, even as an empty list []")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var groups []exclusionGroup
	if err := decoder.Decode(&groups); err != nil {
		var syntax *json.SyntaxError
		if errors.As(err, &syntax) {
			return nil, invalidExclusions(fmt.Sprintf("It is not valid JSON near byte %d", syntax.Offset))
		}
		return nil, invalidExclusions("It must be a JSON list of groups, each with only a reason string and an operations list of strings")
	}
	var extra json.RawMessage
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, invalidExclusions("It holds more than 1 JSON value")
	}
	excluded := operations{}
	for index, group := range groups {
		number := index + 1
		if strings.TrimSpace(group.Reason) == "" {
			return nil, invalidExclusions(fmt.Sprintf("Group %d needs a non-empty reason", number))
		}
		if len(group.Operations) == 0 {
			return nil, invalidExclusions(fmt.Sprintf("Group %d needs at least 1 operation", number))
		}
		for _, text := range group.Operations {
			match := operationPattern.FindStringSubmatch(text)
			if match == nil || !knownMethod(match[1]) {
				return nil, invalidExclusions(fmt.Sprintf("Group %d lists %q, which is not an upper case HTTP method, a space and a path starting with /", number, text))
			}
			key := operationKey(match[1], match[2])
			if earlier, duplicate := excluded[key]; duplicate {
				return nil, invalidExclusions(fmt.Sprintf("It lists %s and %s, which name the same operation. List it once", earlier, text))
			}
			excluded[key] = text
		}
	}
	return excluded, nil
}
