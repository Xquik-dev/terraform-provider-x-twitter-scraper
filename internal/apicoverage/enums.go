// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"go/format"
	"maps"
	"net/http"
	"slices"
)

const (
	enumNextStep = "Run go generate ./internal/apienum and ./scripts/generate-docs, then rerun."
	enumHeader   = `// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// go generate ./internal/apienum writes this file from the enums in https://xquik.com/openapi.json. Change cmd/api-enums, not this file.

package apienum
`
)

func missingEnum(name string) error {
	return fmt.Errorf("the API spec has no %s enum of strings in components.schemas", name)
}

func specEnum(document spec, name string) ([]string, bool) {
	var schema struct {
		Enum []string `json:"enum"`
	}
	raw, found := document.Components.Schemas[name]
	if !found || json.Unmarshal(raw, &schema) != nil || len(schema.Enum) == 0 {
		return nil, false
	}
	return schema.Enum, true
}

func compareEnums(document spec, enums map[string][]string) ([]string, []string) {
	var failures, summaries []string
	names := slices.Sorted(maps.Keys(enums))
	for _, name := range names {
		api, found := specEnum(document, name)
		if !found {
			failures = append(failures, Sentence(missingEnum(name)))
			continue
		}
		provider := enums[name]
		before := len(failures)
		for _, value := range api {
			if !slices.Contains(provider, value) {
				failures = append(failures, fmt.Sprintf("%s %q is in the API but not in the provider.", name, value))
			}
		}
		for _, value := range provider {
			if !slices.Contains(api, value) {
				failures = append(failures, fmt.Sprintf("%s %q is in the provider but not in the API.", name, value))
			}
		}
		if len(failures) == before {
			summaries = append(summaries, fmt.Sprintf("The provider takes all %d %s values.", len(api), name))
		}
	}
	if len(failures) > 0 {
		failures = append(failures, enumNextStep)
	}
	return failures, summaries
}

func GenerateEnums(source string, client *http.Client, names []string) ([]byte, error) {
	document, err := readSpec(cmp.Or(source, defaultSpecSource), client)
	if err != nil {
		return nil, err
	}
	var code bytes.Buffer
	code.WriteString(enumHeader)
	for _, name := range names {
		values, found := specEnum(document, name)
		if !found {
			return nil, missingEnum(name)
		}
		fmt.Fprintf(&code, "\nvar %s = []string{\n", name)
		for _, value := range values {
			fmt.Fprintf(&code, "%q,\n", value)
		}
		code.WriteString("}\n")
	}
	code.WriteString("\nvar Spec = map[string][]string{\n")
	for _, name := range names {
		fmt.Fprintf(&code, "%q: %s,\n", name, name)
	}
	code.WriteString("}\n")
	return format.Source(code.Bytes())
}
