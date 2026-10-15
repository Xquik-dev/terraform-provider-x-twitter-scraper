// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path"
	"regexp"
	"strconv"
	"strings"
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

func knownMethod(method string) bool {
	for _, known := range httpMethods {
		if known == method {
			return true
		}
	}
	return false
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

func receiverName(function *ast.FuncDecl) string {
	if function.Recv == nil {
		return ""
	}
	expression := function.Recv.List[0].Type
	if star, ok := expression.(*ast.StarExpr); ok {
		expression = star.X
	}
	if ident, ok := expression.(*ast.Ident); ok {
		return ident.Name
	}
	return ""
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
