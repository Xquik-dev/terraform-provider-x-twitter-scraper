// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"fmt"
	"go/ast"
	"go/token"
	"io/fs"
	"strconv"
	"strings"
)

const (
	tableMethodField = "method"
	tablePathField   = "pathTemplate"
)

var rawMethods = map[string]string{"Get": "GET", "Post": "POST", "Put": "PUT", "Patch": "PATCH", "Delete": "DELETE"}

type providerPackage struct {
	sdk         *sdk
	positions   *token.FileSet
	files       []*ast.File
	clientNames map[string]bool
	constants   map[string]string
	functions   map[string]*ast.FuncDecl
}

func (p *providerPackage) position(node ast.Node) string {
	position := p.positions.Position(node.Pos())
	return fmt.Sprintf("%s:%d", position.Filename, position.Line)
}

func sdkAlias(file *ast.File, packageName string) string {
	for _, imported := range file.Imports {
		if importPath, _ := strconv.Unquote(imported.Path.Value); importPath != SDKModule {
			continue
		}
		if imported.Name != nil {
			return imported.Name.Name
		}
		return packageName
	}
	return ""
}

func (p *providerPackage) collectDeclarations(file *ast.File) {
	alias := sdkAlias(file, p.sdk.packageName)
	for _, declaration := range file.Decls {
		switch declaration := declaration.(type) {
		case *ast.FuncDecl:
			if declaration.Recv == nil {
				p.functions[declaration.Name.Name] = declaration
			}
		case *ast.GenDecl:
			if declaration.Tok == token.CONST {
				p.collectConstants(declaration)
			}
		}
	}
	if alias == "" {
		return
	}
	isClient := func(expression ast.Expr) bool {
		if star, ok := expression.(*ast.StarExpr); ok {
			expression = star.X
		}
		name, ok := selectorOf(expression, alias)
		return ok && name == clientType
	}
	ast.Inspect(file, func(node ast.Node) bool {
		var names []*ast.Ident
		switch node := node.(type) {
		case *ast.Field:
			if isClient(node.Type) {
				names = node.Names
			}
		case *ast.ValueSpec:
			if isClient(node.Type) {
				names = node.Names
			}
		}
		for _, name := range names {
			p.clientNames[name.Name] = true
		}
		return true
	})
}

func (p *providerPackage) collectConstants(declaration *ast.GenDecl) {
	for _, specification := range declaration.Specs {
		values, ok := specification.(*ast.ValueSpec)
		if !ok || len(values.Names) != len(values.Values) {
			continue
		}
		for index, name := range values.Names {
			if value, ok := stringValue(values.Values[index]); ok {
				p.constants[name.Name] = value
			}
		}
	}
}

func selectorChain(expression ast.Expr) []string {
	var names []string
	for {
		switch current := expression.(type) {
		case *ast.SelectorExpr:
			names = append([]string{current.Sel.Name}, names...)
			expression = current.X
		case *ast.Ident:
			return append([]string{current.Name}, names...)
		default:
			return nil
		}
	}
}

func (p *providerPackage) readPath(function *ast.FuncDecl, argument ast.Expr, helperAllowed bool) (string, bool) {
	if found, ok := requestPath(function, argument); ok {
		return found, true
	}
	switch expression := argument.(type) {
	case *ast.Ident:
		if len(assignments(function.Body, expression.Name)) > 0 {
			return "", false
		}
		found, ok := p.constants[expression.Name]
		return found, ok
	case *ast.CallExpr:
		name, ok := expression.Fun.(*ast.Ident)
		if !ok || !helperAllowed || p.functions[name.Name] == nil {
			return "", false
		}
		helper := p.functions[name.Name]
		if helper.Body == nil || len(helper.Body.List) != 1 {
			return "", false
		}
		result, ok := helper.Body.List[0].(*ast.ReturnStmt)
		if !ok || len(result.Results) != 1 {
			return "", false
		}
		return p.readPath(helper, result.Results[0], false)
	}
	return "", false
}

func parameterType(function *ast.FuncDecl, name string) string {
	lists := []*ast.FieldList{function.Recv, function.Type.Params}
	for _, list := range lists {
		if list == nil {
			continue
		}
		for _, field := range list.List {
			for _, ident := range field.Names {
				if ident.Name == name {
					return typeName(field.Type)
				}
			}
		}
	}
	return ""
}

func (p *providerPackage) literalsOf(name string) []*ast.CompositeLit {
	var literals []*ast.CompositeLit
	if name == "" {
		return nil
	}
	for _, file := range p.files {
		ast.Inspect(file, func(node ast.Node) bool {
			literal, ok := node.(*ast.CompositeLit)
			if !ok {
				return true
			}
			if typeName(literal.Type) == name {
				literals = append(literals, literal)
			}
			if array, ok := literal.Type.(*ast.ArrayType); ok && typeName(array.Elt) == name {
				for _, element := range literal.Elts {
					if inner, ok := element.(*ast.CompositeLit); ok && inner.Type == nil {
						literals = append(literals, inner)
					}
				}
			}
			return true
		})
	}
	return literals
}

func tableRequest(literal *ast.CompositeLit) (request, bool) {
	var found request
	var methodOK, pathOK bool
	for _, element := range literal.Elts {
		field, ok := element.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		switch key, _ := field.Key.(*ast.Ident); key.String() {
		case tableMethodField:
			found.method, methodOK = requestMethod(field.Value)
		case tablePathField:
			if format, ok := stringValue(field.Value); ok {
				found.path, pathOK = formatPath(format, strings.Count(format, "%"))
			}
		}
	}
	return found, methodOK && pathOK
}

func (p *providerPackage) unreadableRequest(call *ast.CallExpr) []string {
	return []string{p.position(call) + " sends a request whose method or path the coverage check cannot read. Use an http.Method constant and a path string, a fmt.Sprintf format, a string constant or a 1 line function that returns 1 of those."}
}

func (p *providerPackage) tableRequests(function *ast.FuncDecl, call *ast.CallExpr) ([]request, []string) {
	unreadable := p.unreadableRequest(call)
	selector, ok := call.Args[1].(*ast.SelectorExpr)
	if !ok || selector.Sel.Name != tableMethodField {
		return nil, unreadable
	}
	value, ok := selector.X.(*ast.Ident)
	pathArgument := call.Args[2]
	if pathCall, isCall := pathArgument.(*ast.CallExpr); isCall {
		pathArgument = pathCall.Fun
	}
	pathRoot := selectorChain(pathArgument)
	if !ok || len(pathRoot) != 2 || pathRoot[0] != value.Name {
		return nil, unreadable
	}
	literals := p.literalsOf(parameterType(function, value.Name))
	if len(literals) == 0 {
		return nil, unreadable
	}
	var requests []request
	var problems []string
	for _, literal := range literals {
		found, ok := tableRequest(literal)
		if !ok {
			problems = append(problems, fmt.Sprintf("%s lists a request without a readable %s and %s. Set %s to an http.Method constant and %s to a path string or fmt.Sprintf format.", p.position(literal), tableMethodField, tablePathField, tableMethodField, tablePathField))
			continue
		}
		requests = append(requests, found)
	}
	return requests, problems
}

func (p *providerPackage) rawRequests(function *ast.FuncDecl, call *ast.CallExpr, name string) ([]request, []string) {
	method, pathIndex := rawMethods[name], 1
	if name == "Execute" && len(call.Args) >= 3 {
		var ok bool
		if method, ok = requestMethod(call.Args[1]); !ok {
			return p.tableRequests(function, call)
		}
		pathIndex = 2
	}
	if method == "" {
		return nil, []string{fmt.Sprintf("%s calls %s.%s, which sends no request the coverage check can read. Call a Go SDK service method or Get, Post, Put, Patch, Delete or Execute.", p.position(call), clientType, name)}
	}
	if len(call.Args) > pathIndex {
		if target, ok := p.readPath(function, call.Args[pathIndex], true); ok {
			return []request{{method, target}}, nil
		}
	}
	return nil, p.unreadableRequest(call)
}

func (p *providerPackage) serviceRequests(call *ast.CallExpr, chain []string) ([]request, []string) {
	current := clientType
	for _, field := range chain[:len(chain)-1] {
		current = p.sdk.fields[current][field]
	}
	method, ok := p.sdk.methods[current][chain[len(chain)-1]]
	label := clientType + "." + strings.Join(chain, ".")
	switch {
	case !ok:
		return nil, []string{fmt.Sprintf("%s calls %s, which is no service method of the Go SDK in go.mod. Call a Go SDK service method, or update the Go SDK in go.mod.", p.position(call), label)}
	case len(method.unreadable) > 0:
		return nil, []string{fmt.Sprintf("%s calls %s, whose request the coverage check cannot read in the Go SDK at %s. Fix that request in the Go SDK first.", p.position(call), label, strings.Join(method.unreadable, " and "))}
	case len(method.requests) == 0:
		return nil, []string{fmt.Sprintf("%s calls %s, which sends no request through %s. Call the Go SDK method that sends the request.", p.position(call), label, requestFunction)}
	}
	return method.requests, nil
}

func (p *providerPackage) callRequests(function *ast.FuncDecl, call *ast.CallExpr) ([]request, []string) {
	chain := selectorChain(call.Fun)
	for index := 0; index < len(chain)-1; index++ {
		if !p.clientNames[chain[index]] {
			continue
		}
		rest := chain[index+1:]
		if len(rest) == 1 {
			return p.rawRequests(function, call, rest[0])
		}
		return p.serviceRequests(call, rest)
	}
	return nil, nil
}

func (p *providerPackage) requests() ([]request, []string) {
	var requests []request
	var problems []string
	for _, file := range p.files {
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if !ok || function.Body == nil {
				continue
			}
			ast.Inspect(function.Body, func(node ast.Node) bool {
				if call, ok := node.(*ast.CallExpr); ok {
					found, unreadable := p.callRequests(function, call)
					requests = append(requests, found...)
					problems = append(problems, unreadable...)
				}
				return true
			})
		}
	}
	return requests, problems
}

func providerRequests(root fs.FS, index *sdk) ([]request, []string, error) {
	positions := token.NewFileSet()
	var requests []request
	var problems []string
	err := fs.WalkDir(root, servicesRoot, func(dir string, entry fs.DirEntry, err error) error {
		if err != nil || !entry.IsDir() {
			return err
		}
		files, err := parseDir(root, dir, positions)
		if err != nil {
			return err
		}
		pkg := &providerPackage{sdk: index, positions: positions, files: files, clientNames: map[string]bool{}, constants: map[string]string{}, functions: map[string]*ast.FuncDecl{}}
		for _, file := range files {
			pkg.collectDeclarations(file)
		}
		found, unreadable := pkg.requests()
		requests = append(requests, found...)
		problems = append(problems, unreadable...)
		return nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("The coverage check could not read the provider source under %s. Run go vet ./... and fix the error, then rerun from the repository root.", servicesRoot)
	}
	return requests, problems, nil
}
