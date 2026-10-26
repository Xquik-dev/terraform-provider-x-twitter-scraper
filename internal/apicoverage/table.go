// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicoverage

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
)

func (p *providerPackage) fieldSelection(expression ast.Expr) (*types.Var, types.Object) {
	selector, ok := ast.Unparen(expression).(*ast.SelectorExpr)
	if !ok {
		return nil, nil
	}
	selection := p.pkg.TypesInfo.Selections[selector]
	if selection == nil || selection.Kind() != types.FieldVal || len(selection.Index()) != 1 {
		return nil, nil
	}
	owner := p.objectOf(selector.X)
	if owner == nil {
		return nil, nil
	}
	field, ok := selection.Obj().(*types.Var)
	if !ok {
		return nil, nil
	}
	return field, owner
}

func (p *providerPackage) returnedField(result *ast.ReturnStmt, receiver types.Object) *types.Var {
	if len(result.Results) != 1 {
		return nil
	}
	expression := result.Results[0]
	if call, ok := ast.Unparen(expression).(*ast.CallExpr); ok && p.isSprintf(call) {
		expression = call.Args[0]
	}
	field, owner := p.fieldSelection(expression)
	if owner != receiver {
		return nil
	}
	return field
}

func (p *providerPackage) formattedField(method *ast.FuncDecl) *types.Var {
	names := method.Recv.List[0].Names
	if len(names) != 1 || method.Body == nil {
		return nil
	}
	receiver := p.pkg.TypesInfo.Defs[names[0]]
	var field *types.Var
	consistent := true
	ast.Inspect(method.Body, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.FuncLit:
			return false
		case *ast.ReturnStmt:
			found := p.returnedField(node, receiver)
			consistent = consistent && found != nil && (field == nil || found == field)
			field = found
		}
		return true
	})
	if !consistent {
		return nil
	}
	return field
}

func (p *providerPackage) pathField(expression ast.Expr) (*types.Var, types.Object) {
	if field, owner := p.fieldSelection(expression); field != nil {
		return field, owner
	}
	call, ok := ast.Unparen(expression).(*ast.CallExpr)
	if !ok {
		return nil, nil
	}
	selector, ok := ast.Unparen(call.Fun).(*ast.SelectorExpr)
	if !ok {
		return nil, nil
	}
	selection := p.pkg.TypesInfo.Selections[selector]
	if selection == nil || selection.Kind() != types.MethodVal {
		return nil, nil
	}
	function, ok := selection.Obj().(*types.Func)
	if !ok {
		return nil, nil
	}
	method := p.functions[function]
	owner := p.objectOf(selector.X)
	if method == nil || owner == nil {
		return nil, nil
	}
	field := p.formattedField(method)
	if field == nil {
		return nil, nil
	}
	return field, owner
}

func (p *providerPackage) literalsOf(table types.Type) []*ast.CompositeLit {
	if pointer, ok := table.(*types.Pointer); ok {
		table = pointer.Elem()
	}
	var literals []*ast.CompositeLit
	for _, file := range p.pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			if literal, ok := node.(*ast.CompositeLit); ok && types.Identical(p.pkg.TypesInfo.TypeOf(literal), table) {
				literals = append(literals, literal)
			}
			return true
		})
	}
	return literals
}

func (p *providerPackage) tableRequest(literal *ast.CompositeLit, methodField, pathField *types.Var) (request, bool) {
	fields, ok := p.pkg.TypesInfo.TypeOf(literal).Underlying().(*types.Struct)
	if !ok {
		return request{}, false
	}
	values := map[*types.Var]ast.Expr{}
	for index, element := range literal.Elts {
		pair, keyed := element.(*ast.KeyValueExpr)
		if !keyed {
			values[fields.Field(index)] = element
			continue
		}
		if field, ok := p.objectOf(pair.Key).(*types.Var); ok {
			values[field] = pair.Value
		}
	}
	method, methodOK := p.constantString(values[methodField])
	template, pathOK := p.constantString(values[pathField])
	if !methodOK || !pathOK || !knownMethod(method) {
		return request{}, false
	}
	target, ok := formatPath(template, strings.Count(template, "%"))
	return request{method, target}, ok
}

func (p *providerPackage) tableRequests(call *ast.CallExpr) ([]request, []string) {
	methodField, owner := p.fieldSelection(call.Args[1])
	pathField, pathOwner := p.pathField(call.Args[2])
	if methodField == nil || pathField == nil || owner != pathOwner {
		return nil, p.unreadableRequest(call)
	}
	literals := p.literalsOf(owner.Type())
	if len(literals) == 0 {
		return nil, p.unreadableRequest(call)
	}
	var requests []request
	var problems []string
	for _, literal := range literals {
		found, ok := p.tableRequest(literal, methodField, pathField)
		if !ok {
			problems = append(problems, fmt.Sprintf("%s lists a request without a constant %s and %s. Set %s to an http.Method constant and %s to a path string or format.", p.position(literal), methodField.Name(), pathField.Name(), methodField.Name(), pathField.Name()))
			continue
		}
		requests = append(requests, found)
	}
	return requests, problems
}
