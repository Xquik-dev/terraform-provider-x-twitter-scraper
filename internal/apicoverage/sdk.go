// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
)

const (
	clientType      = "Client"
	requestFunction = "ExecuteNewRequest"
)

var (
	verbPattern = regexp.MustCompile(`%[sdv]`)
	httpMethods = map[string]string{
		"MethodGet": "GET", "MethodPut": "PUT", "MethodPost": "POST", "MethodDelete": "DELETE",
		"MethodOptions": "OPTIONS", "MethodHead": "HEAD", "MethodPatch": "PATCH", "MethodTrace": "TRACE",
	}
)

type request struct {
	method string
	path   string
}

type sdkMethod struct {
	requests   []request
	unreadable []string
}

type sdk struct {
	packageName string
	basePath    string
	fields      map[string]map[string]string
	methods     map[string]map[string]sdkMethod
}

func parseDir(root fs.FS, dir string, files *token.FileSet) ([]*ast.File, error) {
	entries, err := fs.ReadDir(root, dir)
	if err != nil {
		return nil, err
	}
	var parsed []*ast.File
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		filename := path.Join(dir, name)
		source, err := fs.ReadFile(root, filename)
		if err != nil {
			return nil, err
		}
		file, err := parser.ParseFile(files, filename, source, parser.SkipObjectResolution)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, file)
	}
	return parsed, nil
}

func stringValue(expression ast.Expr) (string, bool) {
	literal, ok := expression.(*ast.BasicLit)
	if !ok || literal.Kind != token.STRING {
		return "", false
	}
	value, err := strconv.Unquote(literal.Value)
	return value, err == nil
}

func selectorOf(expression ast.Expr, pkg string) (string, bool) {
	selector, ok := expression.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}
	ident, ok := selector.X.(*ast.Ident)
	return selector.Sel.Name, ok && ident.Name == pkg
}

func typeName(expression ast.Expr) string {
	if star, ok := expression.(*ast.StarExpr); ok {
		expression = star.X
	}
	if ident, ok := expression.(*ast.Ident); ok {
		return ident.Name
	}
	return ""
}

func receiverName(function *ast.FuncDecl) string {
	if function.Recv == nil {
		return ""
	}
	return typeName(function.Recv.List[0].Type)
}

func assignments(body *ast.BlockStmt, name string) []ast.Expr {
	var values []ast.Expr
	ast.Inspect(body, func(node ast.Node) bool {
		assignment, ok := node.(*ast.AssignStmt)
		if !ok || len(assignment.Lhs) != len(assignment.Rhs) {
			return true
		}
		for index, target := range assignment.Lhs {
			if ident, ok := target.(*ast.Ident); ok && ident.Name == name {
				values = append(values, assignment.Rhs[index])
			}
		}
		return true
	})
	return values
}

func formatPath(format string, arguments int) (string, bool) {
	if len(verbPattern.FindAllString(format, -1)) != arguments || strings.Count(format, "%") != arguments {
		return "", false
	}
	return verbPattern.ReplaceAllString(format, "{}"), true
}

func requestPath(function *ast.FuncDecl, argument ast.Expr) (string, bool) {
	expression := argument
	if ident, ok := argument.(*ast.Ident); ok {
		values := assignments(function.Body, ident.Name)
		if len(values) != 1 {
			return "", false
		}
		expression = values[0]
	}
	if value, ok := stringValue(expression); ok {
		return value, true
	}
	call, ok := expression.(*ast.CallExpr)
	if !ok || len(call.Args) == 0 {
		return "", false
	}
	if name, ok := selectorOf(call.Fun, "fmt"); !ok || name != "Sprintf" {
		return "", false
	}
	format, ok := stringValue(call.Args[0])
	if !ok {
		return "", false
	}
	return formatPath(format, len(call.Args)-1)
}

func requestMethod(argument ast.Expr) (string, bool) {
	name, ok := selectorOf(argument, "http")
	method, known := httpMethods[name]
	return method, ok && known
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
	unreadable := errors.New("The coverage check could not read the default base URL from option.WithEnvironmentProduction in the Go SDK. Check the Go SDK version in go.mod.")
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

func (s *sdk) addStructs(declaration *ast.GenDecl) {
	for _, specification := range declaration.Specs {
		typeSpec, ok := specification.(*ast.TypeSpec)
		if !ok {
			continue
		}
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		fields := map[string]string{}
		for _, field := range structType.Fields.List {
			for _, name := range field.Names {
				fields[name.Name] = typeName(field.Type)
			}
		}
		s.fields[typeSpec.Name.Name] = fields
	}
}

func readSDK(root fs.FS) (*sdk, error) {
	files := token.NewFileSet()
	sources, err := parseDir(root, ".", files)
	if err != nil || len(sources) == 0 {
		return nil, errors.New("The coverage check could not read the Go SDK source. Run go mod download, then rerun.")
	}
	base, err := sdkBasePath(root, files)
	if err != nil {
		return nil, err
	}
	index := &sdk{packageName: sources[0].Name.Name, basePath: base, fields: map[string]map[string]string{}, methods: map[string]map[string]sdkMethod{}}
	for _, file := range sources {
		for _, declaration := range file.Decls {
			switch declaration := declaration.(type) {
			case *ast.GenDecl:
				index.addStructs(declaration)
			case *ast.FuncDecl:
				receiver := receiverName(declaration)
				if !strings.HasSuffix(receiver, "Service") {
					continue
				}
				if index.methods[receiver] == nil {
					index.methods[receiver] = map[string]sdkMethod{}
				}
				index.methods[receiver][declaration.Name.Name] = readRequests(files, declaration)
			}
		}
	}
	return index, nil
}
