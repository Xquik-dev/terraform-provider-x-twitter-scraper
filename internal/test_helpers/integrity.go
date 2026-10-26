// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package testhelpers checks that provider models match their Terraform
// schemas.
package testhelpers

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	ds "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rs "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
)

var _ codingerror = (*mismatch)(nil)
var _ codingerror = (*missing)(nil)
var _ codingerror = (*bug)(nil)
var _ error = (*codingerrors)(nil)
var _ attrlike = (*att)(nil)
var _ attrlike = (*attri)(nil)

type codingerror interface {
	error
	id() string
}

type path []string

type attrlike interface {
	val() reflect.Value
	tag() string
	custom(path) attr.Type
}

type bug struct {
	path path
	msg  string
}

type missing struct {
	path path
	kind string
	what string
}

type mismatch struct {
	path     path
	expected string
	tagging  bool
	tags     string
	actual   reflect.Type
}

func (path path) pathName() string {
	return "." + strings.Join(path, ".")
}

func (e *bug) Error() string {
	return fmt.Sprintf("bug at %q: %q", e.id(), e.msg)
}
func (e *missing) Error() string {
	return fmt.Sprintf("missing at %q on %s: %q", e.id(), e.kind, e.what)
}
func (e *mismatch) Error() string {
	found := "<MISSING>"
	if e.tagging {
		found = e.tags
	} else if e.actual != nil {
		found = e.actual.Name()
		if found == "" {
			found = e.actual.Kind().String()
		}
	}
	return fmt.Sprintf("mismatch at %q: expected %q, received %q", e.id(), e.expected, found)
}

func (e *bug) id() string {
	return e.path.pathName()
}
func (e *missing) id() string {
	return e.path.pathName()
}
func (e *mismatch) id() string {
	return e.path.pathName()
}

type codingerrors []codingerror

func (errs *codingerrors) Ignore(t *testing.T, paths ...string) {
	acc := map[string]bool{}
	for _, path := range paths {
		acc[path] = true
	}
	*errs = slices.DeleteFunc(*errs, func(err codingerror) bool {
		id := err.id()
		_, ok := acc[id]
		if ok {
			delete(acc, id)
		}
		return ok
	})
	if len(acc) > 0 {
		t.Errorf("Superfluous ignore paths: %v", acc)
	}
}

func (errs *codingerrors) Error() string {
	e := []string{}
	for _, err := range *errs {
		e = append(e, err.Error())
	}
	return strings.Join(e, "\n")
}

func (errs *codingerrors) Report(t *testing.T) {
	for _, err := range *errs {
		t.Error(err.Error())
	}
	if len(*errs) > 0 {
		t.Logf("If you are implementing a custom solution, you can suppress these warning(s) by adding the following to the test, as an example:\n + errs.Ignore(%q)\n", (*errs)[0].id())
	}
}

type att struct{ v attr.Type }
type attri struct{ v ds.Attribute }

func (a *att) val() reflect.Value {
	return reflect.ValueOf(a.v)
}
func (a *attri) val() reflect.Value {
	return reflect.ValueOf(a.v)
}

func (a *att) tag() string {
	return ""
}
func (a *attri) tag() string {
	if a.v.IsRequired() {
		return "required"
	}

	computed, optional := a.v.IsComputed(), a.v.IsOptional()
	if computed && optional {
		return "computed_optional"
	} else if computed {
		return "computed"
	} else if optional {
		return "optional"
	} else {
		return ""
	}
}

func (a *att) custom(path) attr.Type {
	return nil
}
func (a *attri) custom(path path) attr.Type {
	switch v := (a.v).(type) {
	case ds.BoolAttribute:
		return v.CustomType
	case rs.BoolAttribute:
		return v.CustomType

	case ds.Int64Attribute:
		return v.CustomType
	case rs.Int64Attribute:
		return v.CustomType

	case ds.Float64Attribute:
		return v.CustomType
	case rs.Float64Attribute:
		return v.CustomType

	case ds.NumberAttribute:
		return v.CustomType
	case rs.NumberAttribute:
		return v.CustomType

	case ds.StringAttribute:
		return v.CustomType
	case rs.StringAttribute:
		return v.CustomType

	case ds.ListAttribute:
		return v.CustomType
	case rs.ListAttribute:
		return v.CustomType

	case ds.SetAttribute:
		return v.CustomType
	case rs.SetAttribute:
		return v.CustomType

	case ds.MapAttribute:
		return v.CustomType
	case rs.MapAttribute:
		return v.CustomType

	case ds.ObjectAttribute:
		return v.CustomType
	case rs.ObjectAttribute:
		return v.CustomType

	case ds.SingleNestedAttribute:
		return v.CustomType
	case rs.SingleNestedAttribute:
		return v.CustomType

	case ds.ListNestedAttribute:
		return v.CustomType
	case rs.ListNestedAttribute:
		return v.CustomType

	case ds.SetNestedAttribute:
		return v.CustomType
	case rs.SetNestedAttribute:
		return v.CustomType

	case ds.MapNestedAttribute:
		return v.CustomType
	case rs.MapNestedAttribute:
		return v.CustomType

	default:
		log.Printf("custom: Unexpected attribute type at %q: %T", path.pathName(), v)
		return nil
	}
}

func ValidateDataSourceModelSchemaIntegrity[T any](model T, schema ds.Schema) codingerrors {
	return ValidateDataSourceModelTypeIntegrity(reflect.TypeOf(model), schema)
}

func ValidateDataSourceModelTypeIntegrity(model reflect.Type, schema ds.Schema) codingerrors {
	path := []string{}
	attributes := map[string]attrlike{}
	for name, attr := range schema.Attributes {
		attributes[name] = &attri{attr}
	}
	return walkAttributes(path, model, attributes)
}

func ValidateResourceModelSchemaIntegrity[T any](model T, schema rs.Schema) codingerrors {
	return ValidateResourceModelTypeIntegrity(reflect.TypeOf(model), schema)
}

func ValidateResourceModelTypeIntegrity(model reflect.Type, schema rs.Schema) codingerrors {
	path := []string{}
	attributes := map[string]attrlike{}
	for name, attr := range schema.Attributes {
		attributes[name] = &attri{attr}
	}
	return walkAttributes(path, model, attributes)
}

func deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return deref(t.Elem())
	}
	return t
}

func checkTag(path []string, attr attrlike, field reflect.StructField) (errs codingerrors) {
	tag := attr.tag()
	if tag == "" {
		return
	}
	tagged, tags := false, []string{}

	for _, label := range []string{"path", "query", "json"} {
		if _, ok := field.Tag.Lookup(label); ok {
			tagged = true
		}
	}
	if path, ok := field.Tag.Lookup("path"); ok {
		if _, options, found := strings.Cut(path, ","); found {
			tags = append(tags, strings.Split(options, ",")...)
		}
	}
	if options, ok := field.Tag.Lookup("api"); ok {
		tags = append(tags, strings.Split(options, ",")...)
	}

	if !tagged {
		return
	}
	if slices.Contains(tags, tag) {
		return
	}
	return codingerrors{&mismatch{path: path, expected: tag, tagging: true, tags: strings.Join(tags, ",")}}
}

func checkCustom(path []string, attribute attrlike, model reflect.Type) (errs codingerrors) {
	custom := attribute.custom(path)
	if custom == nil {
		return
	}

	t := deref(reflect.TypeOf(custom.ValueType(context.TODO())))
	if model.ConvertibleTo(t) {
		return
	}

	return codingerrors{&mismatch{path: path, expected: t.Name(), actual: model}}
}

func walkAttributes(path []string, model reflect.Type, attributes map[string]attrlike) (errs codingerrors) {
	model = deref(model)
	if model.Kind() != reflect.Struct {
		return append(errs, &mismatch{path: path, expected: "Struct", actual: model})
	}

	path = append(path, "@"+model.Name())
	fields := map[string]reflect.StructField{}
	for i := 0; i < model.NumField(); i++ {
		field := model.Field(i)
		if tag, ok := field.Tag.Lookup("tfsdk"); ok {
			fields[tag] = field
		}
	}
	for name, attr := range attributes {
		index := append(path, name)
		if field, ok := fields[name]; ok {
			errs = append(append(errs, checkTag(index, attr, field)...), walk(index, attr, field.Type)...)
		} else {
			errs = append(errs, &mismatch{path: index, expected: attr.val().Type().String(), actual: nil})
		}
	}
	return
}

func genericParam(t reflect.Type) reflect.Type {
	t = deref(t)
	if typed, ok := reflect.TypeAssert[interface{ ModelType() reflect.Type }](reflect.Zero(t)); ok {
		return typed.ModelType()
	}
	if t.Kind() != reflect.Struct || t.NumField() == 0 {
		return nil
	}
	return t.Field(0).Type
}

func checkMapKeys(path []string, model reflect.Type) (errs codingerrors) {
	if model.Kind() != reflect.Map {
		return
	}
	if deref(model.Key()).Kind() != reflect.String {
		return append(errs, &bug{path: path, msg: "checkMapKeys: Expected string keys"})
	}
	return
}

func walkCollection(path path, collection attr.TypeWithElementType, kind reflect.Kind, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	idx, exp, eltype := "[]", "List[T]", collection.ElementType()
	bstype, reflectype := reflect.TypeOf((*basetypes.ListValuable)(nil)), reflect.Slice

	if eltype == nil {
		return append(errs, &bug{path: path, msg: "walkCollection: Unexpected nil ElementType"})
	}
	switch t := collection.(type) {
	case basetypes.ListTypable:
	case basetypes.SetTypable:
		idx, exp = "<>", "Set<T>"
		bstype = reflect.TypeOf((*basetypes.SetValuable)(nil))
	case basetypes.MapTypable:
		idx, exp, reflectype = "{}", "Map{T}", reflect.Map
		bstype = reflect.TypeOf((*basetypes.MapValuable)(nil))
	default:
		log.Printf("walkCollection: Unexpected attribute type at %q: %T", path.pathName(), t)
		return
	}

	if kind == reflectype {
		errs = append(errs, checkMapKeys(path, model)...)
		model = deref(model.Elem())
	} else if model.Implements(bstype.Elem()) {
		model = genericParam(model)
	} else {
		return append(errs, &mismatch{path: path, expected: exp, actual: model})
	}
	return append(errs, walk(append(path, idx), &att{eltype}, model)...)
}

func walkAttriCollection(path path, attribute ds.Attribute, kind reflect.Kind, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	idx, exp, eltype := "[]", "List[T]", (attr.Type)(nil)
	bstype, reflectype := reflect.TypeOf((*basetypes.ListValuable)(nil)), reflect.Slice

	switch t := attribute.(type) {
	case ds.ListAttribute:
		eltype = t.ElementType
	case rs.ListAttribute:
		eltype = t.ElementType
	case ds.SetAttribute:
		idx, exp, eltype = "<>", "Set<T>", t.ElementType
		bstype = reflect.TypeOf((*basetypes.SetValuable)(nil))
	case rs.SetAttribute:
		idx, exp, eltype = "<>", "Set<T>", t.ElementType
		bstype = reflect.TypeOf((*basetypes.SetValuable)(nil))
	case ds.MapAttribute:
		idx, exp, eltype = "{}", "Map{T}", t.ElementType
		bstype, reflectype = reflect.TypeOf((*basetypes.MapValuable)(nil)), reflect.Map
	case rs.MapAttribute:
		idx, exp, eltype = "{}", "Map{T}", t.ElementType
		bstype, reflectype = reflect.TypeOf((*basetypes.MapValuable)(nil)), reflect.Map
	default:
		log.Printf("walkAttriCollection: Unexpected attribute type at %q: %T", path.pathName(), t)
		return
	}

	if eltype == nil {
		return append(errs, &bug{path: path, msg: "walkAttriCollection: Unexpected nil ElementType"})
	}

	if kind == reflectype {
		errs = append(errs, checkMapKeys(path, model)...)
		model = deref(model.Elem())
	} else if model.Implements(bstype.Elem()) {
		for ty, t := range map[string]reflect.Type{
			"List": reflect.TypeOf(basetypes.ListValue{}),
			"Set":  reflect.TypeOf(basetypes.SetValue{}),
			"Map":  reflect.TypeOf(basetypes.MapValue{}),
		} {
			if model.AssignableTo(t) {
				return append(errs, &bug{path: path, msg: fmt.Sprintf("walkAttriCollection: Should use customfield.%s instead of %s", ty, t)})
			}
		}
		model = genericParam(model)
	} else {
		return append(errs, &mismatch{path: path, expected: exp, actual: model})
	}
	return append(errs, walk(append(path, idx), &att{eltype}, model)...)
}

func walkNested(path path, attribute ds.NestedAttribute, kind reflect.Kind, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	idx, exp := []string{}, "{...}"
	basetype, reflectype := reflect.TypeOf((*basetypes.ObjectValuable)(nil)), reflect.Struct

	switch a := attribute.(type) {
	case ds.SingleNestedAttribute, rs.SingleNestedAttribute:
	case ds.ListNestedAttribute, rs.ListNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.ListValuable)(nil)), reflect.Slice
		idx, exp = []string{"[]"}, "List[...]"
	case ds.SetNestedAttribute, rs.SetNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.SetValuable)(nil)), reflect.Slice
		idx, exp = []string{"<>"}, "Set<...>"
	case ds.MapNestedAttribute, rs.MapNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.MapValuable)(nil)), reflect.Map
		idx, exp = []string{"{}"}, "Map{...}"
	default:
		log.Printf("walkNested: Unexpected attribute type %q: %T", path.pathName(), a)
		return
	}

	if kind == reflectype && reflectype != reflect.Struct {
		errs = append(errs, checkMapKeys(path, model)...)
		model = deref(model.Elem())
	} else if model.Implements(basetype.Elem()) {
		errs = append(errs, checkCustom(append(path, idx...), &attri{attribute}, model)...)
		model = genericParam(model)
	}

	if model.Kind() != reflect.Struct {
		return append(errs, &mismatch{path: path, expected: exp, actual: model})
	}

	attributes := map[string]attrlike{}
	for name, nested := range attribute.GetNestedObject().GetAttributes() {
		attributes[name] = &attri{nested}
	}
	return append(errs, walkAttributes(append(path, idx...), model, attributes)...)
}

func walk(path path, attribute attrlike, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	kind := model.Kind()

	value := attribute.val()
	if !value.IsValid() {
		log.Printf("walk: Unexpected nil attribute at %q", path.pathName())
		return
	}
	switch value.Type() {
	case reflect.TypeFor[ds.DynamicAttribute](), reflect.TypeFor[rs.DynamicAttribute]():
		if !model.Implements(reflect.TypeOf((*basetypes.DynamicValuable)(nil)).Elem()) {
			return append(errs, &mismatch{path: path, expected: "Dynamic", actual: model})
		}
		if at, ok := reflect.TypeAssert[rs.DynamicAttribute](value); ok {
			if !slices.ContainsFunc(at.PlanModifiers, func(m planmodifier.Dynamic) bool {
				_, ok := m.(customfield.NormalizingDynamicPlanModifier)
				return ok
			}) {
				return append(errs, &missing{path: path, kind: "schema", what: "PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()}"})
			}
		}
	case reflect.TypeFor[ds.BoolAttribute](), reflect.TypeFor[rs.BoolAttribute](), reflect.TypeFor[basetypes.BoolType]():
		if model.Implements(reflect.TypeOf((*basetypes.BoolValuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if kind != reflect.Bool {
			return append(errs, &mismatch{path: path, expected: "Bool", actual: model})
		}
	case reflect.TypeFor[ds.Int64Attribute](), reflect.TypeFor[rs.Int64Attribute](), reflect.TypeFor[basetypes.Int64Type]():
		if model.Implements(reflect.TypeOf((*basetypes.Int64Valuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if kind != reflect.Int && kind != reflect.Int64 {
			return append(errs, &mismatch{path: path, expected: "Int", actual: model})
		}
	case reflect.TypeFor[ds.Float64Attribute](), reflect.TypeFor[rs.Float64Attribute](), reflect.TypeFor[basetypes.Float64Type]():
		if model.Implements(reflect.TypeOf((*basetypes.Float64Valuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if kind != reflect.Float64 {
			return append(errs, &mismatch{path: path, expected: "Float", actual: model})
		}
	case reflect.TypeFor[ds.NumberAttribute](), reflect.TypeFor[rs.NumberAttribute](), reflect.TypeFor[basetypes.NumberType]():
		if model.Implements(reflect.TypeOf((*basetypes.NumberValuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if model.ConvertibleTo(reflect.TypeOf(big.Int{})) || model.ConvertibleTo(reflect.TypeOf(big.Float{})) {
			return append(errs, &mismatch{path: path, expected: "Number", actual: model})
		}
	case reflect.TypeFor[timetypes.RFC3339Type]():
		if !model.ConvertibleTo(reflect.TypeOf(timetypes.RFC3339{})) {
			return append(errs, &mismatch{path: path, expected: "Time", actual: model})
		}
	case reflect.TypeFor[jsontypes.NormalizedType]():
		if !model.ConvertibleTo(reflect.TypeOf(jsontypes.Normalized{})) {
			return append(errs, &mismatch{path: path, expected: "JSON", actual: model})
		}
	case reflect.TypeFor[ds.StringAttribute](), reflect.TypeFor[rs.StringAttribute](), reflect.TypeFor[basetypes.StringType]():
		if model.Implements(reflect.TypeOf((*basetypes.StringValuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if kind != reflect.String {
			return append(errs, &mismatch{path: path, expected: "String", actual: model})
		}
	case reflect.TypeFor[basetypes.ListType](), reflect.TypeFor[basetypes.SetType](), reflect.TypeFor[basetypes.MapType]():
		collection, ok := reflect.TypeAssert[attr.TypeWithElementType](value)
		if !ok {
			return append(errs, &bug{path: path, msg: "walk: collection type has no element type"})
		}
		return append(errs, walkCollection(path, collection, kind, model)...)
	case reflect.TypeFor[ds.ListAttribute](), reflect.TypeFor[ds.SetAttribute](), reflect.TypeFor[ds.MapAttribute](), reflect.TypeFor[rs.ListAttribute](), reflect.TypeFor[rs.SetAttribute](), reflect.TypeFor[rs.MapAttribute]():
		collection, ok := reflect.TypeAssert[ds.Attribute](value)
		if !ok {
			return append(errs, &bug{path: path, msg: "walk: collection attribute is not a schema attribute"})
		}
		return append(errs, walkAttriCollection(path, collection, kind, model)...)
	case reflect.TypeFor[ds.ObjectAttribute](), reflect.TypeFor[rs.ObjectAttribute](), reflect.TypeFor[basetypes.ObjectType]():
		if model.Implements(reflect.TypeOf((*basetypes.ObjectValuable)(nil)).Elem()) {
			errs = append(errs, checkCustom(path, attribute, model)...)
			model = genericParam(model)
		} else if kind != reflect.Struct {
			return append(errs, &mismatch{path: path, expected: "Object", actual: model})
		}
		attributes := map[string]attrlike{}
		var nestedTypes map[string]attr.Type
		if objt, ok := reflect.TypeAssert[ds.ObjectAttribute](value); ok {
			nestedTypes = objt.AttributeTypes
		} else if objt, ok := reflect.TypeAssert[rs.ObjectAttribute](value); ok {
			nestedTypes = objt.AttributeTypes
		} else if objt, ok := reflect.TypeAssert[basetypes.ObjectType](value); ok {
			nestedTypes = objt.AttrTypes
		}
		for name, nested := range nestedTypes {
			attributes[name] = &att{nested}
		}
		return append(errs, walkAttributes(path, model, attributes)...)
	case reflect.TypeFor[ds.SingleNestedAttribute](), reflect.TypeFor[ds.ListNestedAttribute](), reflect.TypeFor[ds.MapNestedAttribute](), reflect.TypeFor[ds.SetNestedAttribute](), reflect.TypeFor[rs.SingleNestedAttribute](), reflect.TypeFor[rs.SetNestedAttribute](), reflect.TypeFor[rs.ListNestedAttribute](), reflect.TypeFor[rs.MapNestedAttribute]():
		nested, ok := reflect.TypeAssert[ds.NestedAttribute](value)
		if !ok {
			return append(errs, &bug{path: path, msg: "walk: nested attribute has no nested object"})
		}
		return append(errs, walkNested(path, nested, kind, model)...)
	default:
		log.Printf("walk: Unexpected attribute type at %q: %s", path.pathName(), value.Type())
	}

	return
}
