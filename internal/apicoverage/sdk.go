// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"io/fs"
	"net/url"
	"strings"
)

const requestFunction = "ExecuteNewRequest"

type sdkMethod struct {
	requests   []request
	unreadable []string
}

type sdk struct {
	basePath string
	methods  map[string]map[string]sdkMethod
}

func readRequests(files *token.FileSet, function *ast.FuncDecl) sdkMethod {
	var found sdkMethod
	ast.Inspect(function.Body, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || selector.Sel.Name != requestFunction {
			return true
		}
		if len(call.Args) >= 3 {
			method, methodOK := requestMethod(call.Args[1])
			target, pathOK := requestPath(function, call.Args[2])
			if methodOK && pathOK {
				found.requests = append(found.requests, request{method, target})
				return true
			}
		}
		position := files.Position(call.Pos())
		found.unreadable = append(found.unreadable, fmt.Sprintf("%s:%d", position.Filename, position.Line))
		return true
	})
	return found
}

func sdkBasePath(root fs.FS, files *token.FileSet) (string, error) {
	unreadable := errors.New("the coverage check could not read the default base URL from option.WithEnvironmentProduction in the Go SDK. Check the Go SDK version in go.mod")
	sources, err := parseDir(root, "option", files)
	if err != nil {
		return "", unreadable
	}
	for _, file := range sources {
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if !ok || function.Recv != nil || function.Name.Name != "WithEnvironmentProduction" {
				continue
			}
			var base string
			ast.Inspect(function.Body, func(node ast.Node) bool {
				call, ok := node.(*ast.CallExpr)
				if ok && len(call.Args) == 1 {
					if value, ok := stringValue(call.Args[0]); ok {
						base = value
					}
				}
				return true
			})
			if parsed, err := url.Parse(base); base != "" && err == nil {
				return parsed.Path, nil
			}
		}
	}
	return "", unreadable
}

func readSDK(root fs.FS) (*sdk, error) {
	files := token.NewFileSet()
	sources, err := parseDir(root, ".", files)
	if err != nil || len(sources) == 0 {
		return nil, errors.New("the coverage check could not read the Go SDK source. Run go mod download, then rerun")
	}
	base, err := sdkBasePath(root, files)
	if err != nil {
		return nil, err
	}
	index := &sdk{basePath: base, methods: map[string]map[string]sdkMethod{}}
	for _, file := range sources {
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if !ok || !strings.HasSuffix(receiverName(function), "Service") {
				continue
			}
			receiver := receiverName(function)
			if index.methods[receiver] == nil {
				index.methods[receiver] = map[string]sdkMethod{}
			}
			index.methods[receiver][function.Name.Name] = readRequests(files, function)
		}
	}
	return index, nil
}
