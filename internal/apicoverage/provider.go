// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
)

const (
	clientType = "Client"
	loadMode   = packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo
)

var rawMethods = map[string]string{"Get": "GET", "Post": "POST", "Put": "PUT", "Patch": "PATCH", "Delete": "DELETE"}

type providerPackage struct {
	root      string
	pkg       *packages.Package
	sdk       *sdk
	functions map[*types.Func]*ast.FuncDecl
}

func loadConfig(root string, mode packages.LoadMode) *packages.Config {
	return &packages.Config{Mode: mode, Dir: root, Env: append(os.Environ(), "GOWORK=off")}
}

func sdkDir(root string) (string, error) {
	found, err := packages.Load(loadConfig(root, packages.NeedName|packages.NeedModule), SDKModule)
	if err != nil || len(found) != 1 || found[0].Module == nil || found[0].Module.Dir == "" {
		return "", sentence(fmt.Sprintf("The coverage check could not find the source of %s. Run go mod download, then rerun.", SDKModule))
	}
	return found[0].Module.Dir, nil
}

func loadProvider(root string) ([]*packages.Package, []string, error) {
	noPackages := sentence(fmt.Sprintf("The coverage check found no Go packages under %s. Run it from the repository root.", servicesRoot))
	if info, err := os.Stat(filepath.Join(root, servicesRoot)); err != nil || !info.IsDir() {
		return nil, nil, noPackages
	}
	loaded, err := packages.Load(loadConfig(root, loadMode), "./"+servicesRoot+"/...")
	if err != nil || len(loaded) == 0 {
		return nil, nil, noPackages
	}
	var problems []string
	for _, pkg := range loaded {
		for _, loadError := range pkg.Errors {
			problems = append(problems, loadError.Error())
		}
	}
	if len(problems) > 0 {
		return nil, problems, sentence(fmt.Sprintf("The coverage check could not compile the Go packages under %s. Fix the errors above, then rerun.", servicesRoot))
	}
	return loaded, nil, nil
}

func providerRequests(root string, index *sdk) ([]request, []string, error) {
	loaded, problems, err := loadProvider(root)
	if err != nil {
		return nil, problems, err
	}
	var requests []request
	for _, pkg := range loaded {
		reader := &providerPackage{root: root, pkg: pkg, sdk: index, functions: map[*types.Func]*ast.FuncDecl{}}
		for _, file := range pkg.Syntax {
			for _, declaration := range file.Decls {
				if function, ok := declaration.(*ast.FuncDecl); ok {
					if object, ok := pkg.TypesInfo.Defs[function.Name].(*types.Func); ok {
						reader.functions[object] = function
					}
				}
			}
		}
		found, unreadable := reader.requests()
		requests = append(requests, found...)
		problems = append(problems, unreadable...)
	}
	return requests, problems, nil
}

func (p *providerPackage) position(node ast.Node) string {
	position := p.pkg.Fset.Position(node.Pos())
	filename := position.Filename
	if relative, err := filepath.Rel(p.root, filename); err == nil {
		filename = filepath.ToSlash(relative)
	}
	return fmt.Sprintf("%s:%d", filename, position.Line)
}

func (p *providerPackage) unreadableRequest(call *ast.CallExpr) []string {
	return []string{p.position(call) + " sends a request whose method or path the coverage check cannot read. Use an http.Method constant and a constant path, a fmt.Sprintf format, a variable assigned once to 1 of those or a 1 line function that returns 1 of those."}
}

func (p *providerPackage) objectOf(expression ast.Expr) types.Object {
	ident, ok := ast.Unparen(expression).(*ast.Ident)
	if !ok {
		return nil
	}
	if object := p.pkg.TypesInfo.Defs[ident]; object != nil {
		return object
	}
	return p.pkg.TypesInfo.Uses[ident]
}

func (p *providerPackage) constantString(expression ast.Expr) (string, bool) {
	value := p.pkg.TypesInfo.Types[expression].Value
	if value == nil || value.Kind() != constant.String {
		return "", false
	}
	return constant.StringVal(value), true
}

func (p *providerPackage) isSprintf(call *ast.CallExpr) bool {
	callee := typeutil.StaticCallee(p.pkg.TypesInfo, call)
	return callee != nil && callee.FullName() == "fmt.Sprintf"
}

func valueAt(values []ast.Expr, count, index int) ast.Expr {
	if len(values) == count {
		return values[index]
	}
	return nil
}

func (p *providerPackage) assignedValues(variable *types.Var) []ast.Expr {
	var values []ast.Expr
	record := func(target ast.Expr, all []ast.Expr, count, index int) {
		if target != nil && p.objectOf(target) == variable {
			values = append(values, valueAt(all, count, index))
		}
	}
	for _, file := range p.pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.AssignStmt:
				for index, target := range node.Lhs {
					record(target, node.Rhs, len(node.Lhs), index)
				}
			case *ast.ValueSpec:
				for index, name := range node.Names {
					record(name, node.Values, len(node.Names), index)
				}
			case *ast.RangeStmt:
				record(node.Key, nil, 1, 0)
				record(node.Value, nil, 1, 0)
			case *ast.UnaryExpr:
				if node.Op == token.AND {
					record(node.X, nil, 1, 0)
				}
			}
			return true
		})
	}
	return values
}

func (p *providerPackage) readPath(expression ast.Expr, helperAllowed bool) (string, bool) {
	if value, ok := p.constantString(expression); ok {
		return value, true
	}
	switch expression := ast.Unparen(expression).(type) {
	case *ast.Ident:
		variable, _ := p.pkg.TypesInfo.Uses[expression].(*types.Var)
		values := p.assignedValues(variable)
		if len(values) != 1 {
			return "", false
		}
		return p.readPath(values[0], helperAllowed)
	case *ast.CallExpr:
		if p.isSprintf(expression) {
			format, ok := p.constantString(expression.Args[0])
			if !ok || expression.Ellipsis.IsValid() {
				return "", false
			}
			return formatPath(format, len(expression.Args)-1)
		}
		helper := p.functions[typeutil.StaticCallee(p.pkg.TypesInfo, expression)]
		if !helperAllowed || helper == nil || helper.Recv != nil || helper.Body == nil || len(helper.Body.List) != 1 {
			return "", false
		}
		result, ok := helper.Body.List[0].(*ast.ReturnStmt)
		if !ok || len(result.Results) != 1 {
			return "", false
		}
		return p.readPath(result.Results[0], false)
	}
	return "", false
}

func (p *providerPackage) rawRequests(call *ast.CallExpr, name string) ([]request, []string) {
	method, pathIndex := rawMethods[name], 1
	if name == "Execute" {
		value, isConstant := p.constantString(call.Args[1])
		if !isConstant {
			return p.tableRequests(call)
		}
		if !knownMethod(value) {
			return nil, p.unreadableRequest(call)
		}
		method, pathIndex = value, 2
	} else if method == "" {
		return nil, []string{fmt.Sprintf("%s calls %s.%s, which sends no request the coverage check can read. Call a Go SDK service method or Get, Post, Put, Patch, Delete or Execute.", p.position(call), clientType, name)}
	}
	if target, ok := p.readPath(call.Args[pathIndex], true); ok {
		return []request{{method, target}}, nil
	}
	return nil, p.unreadableRequest(call)
}

func namedType(value types.Type) string {
	if pointer, ok := value.(*types.Pointer); ok {
		value = pointer.Elem()
	}
	named, ok := types.Unalias(value).(*types.Named)
	if !ok {
		return ""
	}
	return named.Obj().Name()
}

func (p *providerPackage) sdkMethod(selector *ast.SelectorExpr) (string, bool) {
	function, ok := p.pkg.TypesInfo.Uses[selector.Sel].(*types.Func)
	if !ok || function.Pkg() == nil || function.Pkg().Path() != SDKModule || function.Signature().Recv() == nil {
		return "", false
	}
	receiver := namedType(function.Signature().Recv().Type())
	return receiver, receiver == clientType || strings.HasSuffix(receiver, "Service")
}

func (p *providerPackage) methodRequests(selector *ast.SelectorExpr, call *ast.CallExpr, receiver string) ([]request, []string) {
	name := selector.Sel.Name
	label := receiver + "." + name
	if call == nil {
		return nil, []string{fmt.Sprintf("%s uses %s without calling it. Call it directly so the coverage check can read its request.", p.position(selector), label)}
	}
	if receiver == clientType {
		return p.rawRequests(call, name)
	}
	method := p.sdk.methods[receiver][name]
	switch {
	case len(method.unreadable) > 0:
		return nil, []string{fmt.Sprintf("%s calls %s, whose request the coverage check cannot read in the Go SDK at %s. Fix that request in the Go SDK first.", p.position(call), label, strings.Join(method.unreadable, " and "))}
	case len(method.requests) == 0:
		return nil, []string{fmt.Sprintf("%s calls %s, which sends no request through %s. Call the Go SDK method that sends the request.", p.position(call), label, requestFunction)}
	}
	return method.requests, nil
}

func (p *providerPackage) requests() ([]request, []string) {
	called := map[*ast.SelectorExpr]*ast.CallExpr{}
	var selectors []*ast.SelectorExpr
	for _, file := range p.pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.CallExpr:
				if selector, ok := ast.Unparen(node.Fun).(*ast.SelectorExpr); ok {
					called[selector] = node
				}
			case *ast.SelectorExpr:
				selectors = append(selectors, node)
			}
			return true
		})
	}
	var requests []request
	var problems []string
	for _, selector := range selectors {
		if receiver, ok := p.sdkMethod(selector); ok {
			found, unreadable := p.methodRequests(selector, called[selector], receiver)
			requests = append(requests, found...)
			problems = append(problems, unreadable...)
		}
	}
	return requests, problems
}
