// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apiform

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"path"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
)

var encoders apijson.Cache[encoderEntry, encoderFunc]

func MarshalRoot[T any](value T, writer *multipart.Writer) error {
	e := &encoder{root: true, dateFormat: time.RFC3339}
	return e.marshal(reflect.ValueOf(value), writer)
}

type encoder struct {
	dateFormat string
	root       bool
}

type encoderFunc func(key string, value reflect.Value, writer *multipart.Writer) error

type encoderField struct {
	tag parsedStructTag
	fn  encoderFunc
	idx []int
}

type encoderEntry struct {
	reflect.Type
	dateFormat string
	root       bool
}

func (e *encoder) marshal(val reflect.Value, writer *multipart.Writer) error {
	if !val.IsValid() {
		return nil
	}
	typ := val.Type()
	enc := e.typeEncoder(typ)
	return enc("", val, writer)
}

func (e *encoder) typeEncoder(t reflect.Type) encoderFunc {
	entry := encoderEntry{
		Type:       t,
		dateFormat: e.dateFormat,
		root:       e.root,
	}

	return encoders.Load(entry, func() encoderFunc { return e.newTypeEncoder(t) }, func(wait func() encoderFunc) encoderFunc {
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			return wait()(key, v, writer)
		}
	})
}

func (e *encoder) newTypeEncoder(t reflect.Type) encoderFunc {
	if t.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return e.newTimeTypeEncoder()
	}
	if t.ConvertibleTo(reflect.TypeOf((*io.Reader)(nil)).Elem()) {
		return e.newReaderTypeEncoder()
	}

	isRoot := e.root
	e.root = false

	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()

		innerEncoder := e.typeEncoder(inner)
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			if !v.IsValid() || v.IsNil() {
				return nil
			}
			return innerEncoder(key, v.Elem(), writer)
		}
	case reflect.Struct:
		attrType := reflect.TypeOf((*attr.Value)(nil)).Elem()
		if t.Implements(attrType) {
			return e.newTerraformTypeEncoder(t)
		}
		if isRoot {
			return e.newStructTypeEncoder(t)
		} else {
			return encodePartAsJSON
		}
	case reflect.Slice, reflect.Array:
		return e.newArrayTypeEncoder(t)
	case reflect.Map:
		return e.newMapEncoder(t)
	case reflect.Interface:
		return e.newInterfaceEncoder()
	default:
		return e.newPrimitiveTypeEncoder(t)
	}
}

type terraformUnwrappingFunc func(val attr.Value) (reflect.Value, diag.Diagnostics)

func (e *encoder) terraformUnwrappedEncoder(underlyingType reflect.Type, unwrap terraformUnwrappingFunc) encoderFunc {
	enc := e.typeEncoder(underlyingType)
	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		attrValue, err := valueAs[attr.Value](value)
		if err != nil {
			return err
		}
		if attrValue.IsNull() || attrValue.IsUnknown() {
			return nil
		}
		unwrapped, diags := unwrap(attrValue)
		if diags.HasError() {
			return errorFromDiagnostics(diags)
		}
		return enc(key, unwrapped, writer)
	}
}

func (e *encoder) terraformUnwrappedDynamicEncoder(unwrap terraformUnwrappingFunc) encoderFunc {
	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		attrValue, err := valueAs[attr.Value](value)
		if err != nil {
			return err
		}
		if attrValue.IsNull() || attrValue.IsUnknown() {
			return nil
		}
		unwrapped, diags := unwrap(attrValue)
		if diags.HasError() {
			return errorFromDiagnostics(diags)
		}
		if !unwrapped.IsValid() {
			return nil
		}
		enc := e.typeEncoder(unwrapped.Type())
		return enc(key, unwrapped, writer)
	}
}

func (e *encoder) newTerraformTypeEncoder(t reflect.Type) encoderFunc {
	ctx := context.TODO()

	if t == reflect.TypeOf(basetypes.BoolValue{}) {
		return e.terraformUnwrappedEncoder(reflect.PointerTo(reflect.TypeOf(true)), func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			return apijson.UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.Int64Value{}) {
		return e.terraformUnwrappedEncoder(reflect.PointerTo(reflect.TypeOf(int64(0))), func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			return apijson.UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.Float64Value{}) {
		return e.terraformUnwrappedEncoder(reflect.PointerTo(reflect.TypeOf(float64(0))), func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			return apijson.UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(basetypes.StringValue{}) {
		return e.terraformUnwrappedEncoder(reflect.PointerTo(reflect.TypeOf("")), func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			return apijson.UnwrapTerraformAttrValue(ctx, value)
		})
	} else if t == reflect.TypeOf(timetypes.RFC3339{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(time.Time{}), func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			rfc3339, ok := value.(timetypes.RFC3339)
			if !ok {
				return reflect.Value{}, unexpectedValue(value, "timetypes.RFC3339")
			}
			timeValue, diags := rfc3339.ValueRFC3339Time()
			if diags.HasError() {
				return reflect.Value{}, diags
			}
			return reflect.ValueOf(timeValue), nil
		})
	} else if t == reflect.TypeOf(basetypes.ListValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			list, ok := value.(basetypes.ListValue)
			if !ok {
				return reflect.Value{}, unexpectedValue(value, "basetypes.ListValue")
			}
			return reflect.ValueOf(list.Elements()), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.SetValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			set, ok := value.(basetypes.SetValue)
			if !ok {
				return reflect.Value{}, unexpectedValue(value, "basetypes.SetValue")
			}
			return reflect.ValueOf(set.Elements()), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.MapValue{}) {
		return encodePartAsJSON
	} else if t == reflect.TypeOf(basetypes.ObjectValue{}) {
		return encodePartAsJSON
	} else if t.Implements(reflect.TypeOf((*basetypes.DynamicValuable)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			dynamic, ok := value.(basetypes.DynamicValuable)
			if !ok {
				return reflect.Value{}, unexpectedValue(value, "a dynamic value")
			}
			val, d := dynamic.ToDynamicValue(ctx)
			underlying := val.UnderlyingValue()
			if underlying == nil {
				return reflect.Value{}, d
			}
			return reflect.ValueOf(underlying), d
		})
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectLike)(nil)).Elem()) {
		return encodePartAsJSON
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectListLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			list, ok := value.(customfield.NestedObjectListLike)
			if !ok {
				return reflect.Value{}, unexpectedValue(value, "a nested object list")
			}
			return list.AsStructSlice(context.TODO())
		})
	} else if t.Implements(reflect.TypeOf((*customfield.ListLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (reflect.Value, diag.Diagnostics) {
			list, ok := value.(customfield.ListLike)
			if !ok {
				return reflect.Value{}, unexpectedValue(value, "a list")
			}
			return list.ValueAttr(context.TODO())
		})
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectMapLike)(nil)).Elem()) {
		return encodePartAsJSON
	} else if t.Implements(reflect.TypeOf((*customfield.MapLike)(nil)).Elem()) {
		return encodePartAsJSON
	} else if t == reflect.TypeOf(jsontypes.Normalized{}) {
		return encodePartAsJSON
	}

	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		return fmt.Errorf("unknown type received at terraform encoder: %s", t.String())
	}
}

func (e *encoder) newPrimitiveTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.String:
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			return writer.WriteField(key, v.String())
		}
	case reflect.Bool:
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			if v.Bool() {
				return writer.WriteField(key, "true")
			}
			return writer.WriteField(key, "false")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			return writer.WriteField(key, strconv.FormatInt(v.Int(), 10))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			return writer.WriteField(key, strconv.FormatUint(v.Uint(), 10))
		}
	case reflect.Float32:
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			return writer.WriteField(key, strconv.FormatFloat(v.Float(), 'f', -1, 32))
		}
	case reflect.Float64:
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			return writer.WriteField(key, strconv.FormatFloat(v.Float(), 'f', -1, 64))
		}
	default:
		return func(key string, v reflect.Value, writer *multipart.Writer) error {
			return fmt.Errorf("unknown type received at primitive encoder: %s", t.String())
		}
	}
}

func (e *encoder) newArrayTypeEncoder(t reflect.Type) encoderFunc {
	itemEncoder := e.typeEncoder(t.Elem())

	return func(key string, v reflect.Value, writer *multipart.Writer) error {
		for i := 0; i < v.Len(); i++ {
			err := itemEncoder(key, v.Index(i), writer)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (e *encoder) newStructTypeEncoder(t reflect.Type) encoderFunc {
	encoderFields := []encoderField{}
	extraEncoder := (*encoderField)(nil)

	var collectEncoderFields func(r reflect.Type, index []int)
	collectEncoderFields = func(r reflect.Type, index []int) {
		for i := 0; i < r.NumField(); i++ {
			idx := append(index, i)
			field := t.FieldByIndex(idx)
			if !field.IsExported() {
				continue
			}
			if field.Anonymous {
				collectEncoderFields(field.Type, idx)
				continue
			}
			ptag, ok := parseFormStructTag(field)
			if !ok {
				continue
			}
			if ptag.extras && len(index) == 0 {
				extraEncoder = &encoderField{ptag, e.typeEncoder(field.Type.Elem()), idx}
				continue
			}
			if ptag.name == "-" {
				continue
			}
			if ptag.computed && !ptag.forceEncode {
				continue
			}

			dateFormat, ok := parseFormatStructTag(field)
			oldFormat := e.dateFormat
			if ok {
				switch dateFormat {
				case "date-time":
					e.dateFormat = time.RFC3339
				case "date":
					e.dateFormat = "2006-01-02"
				}
			}
			encoderFields = append(encoderFields, encoderField{ptag, e.typeEncoder(field.Type), idx})
			e.dateFormat = oldFormat
		}
	}
	collectEncoderFields(t, []int{})

	sort.Slice(encoderFields, func(i, j int) bool {
		return encoderFields[i].tag.name < encoderFields[j].tag.name
	})

	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		if key != "" {
			key = key + "."
		}
		for _, ef := range encoderFields {
			field := value.FieldByIndex(ef.idx)
			err := ef.fn(key+ef.tag.name, field, writer)
			if err != nil {
				return err
			}
		}

		if extraEncoder != nil {
			err := e.encodeMapEntries(key, value.FieldByIndex(extraEncoder.idx), writer)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func (e *encoder) newTimeTypeEncoder() encoderFunc {
	format := e.dateFormat
	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		converted, err := valueAs[time.Time](value.Convert(reflect.TypeFor[time.Time]()))
		if err != nil {
			return err
		}
		return writer.WriteField(key, converted.Format(format))
	}
}

func (e encoder) newInterfaceEncoder() encoderFunc {
	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		value = value.Elem()
		if !value.IsValid() {
			return nil
		}
		return e.typeEncoder(value.Type())(key, value, writer)
	}
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func (e *encoder) newReaderTypeEncoder() encoderFunc {
	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		reader, err := valueAs[io.Reader](value.Convert(reflect.TypeFor[io.Reader]()))
		if err != nil {
			return err
		}
		filename := "anonymous_file"
		contentType := "application/octet-stream"
		if named, ok := reader.(interface{ Name() string }); ok {
			filename = path.Base(named.Name())
		}
		if typed, ok := reader.(interface{ ContentType() string }); ok {
			contentType = typed.ContentType()
		}

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(key), escapeQuotes(filename)))
		h.Set("Content-Type", contentType)
		filewriter, err := writer.CreatePart(h)
		if err != nil {
			return err
		}
		_, err = io.Copy(filewriter, reader)
		return err
	}
}

func (e *encoder) encodeMapEntries(key string, v reflect.Value, writer *multipart.Writer) error {
	type mapPair struct {
		key   string
		value reflect.Value
	}

	if key != "" {
		key = key + "."
	}

	pairs := []mapPair{}

	iter := v.MapRange()
	for iter.Next() {
		if iter.Key().Type().Kind() == reflect.String {
			pairs = append(pairs, mapPair{key: iter.Key().String(), value: iter.Value()})
		} else {
			return fmt.Errorf("cannot encode a map with a non string key")
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].key < pairs[j].key
	})

	elementEncoder := e.typeEncoder(v.Type().Elem())
	for _, p := range pairs {
		err := elementEncoder(key+string(p.key), p.value, writer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *encoder) newMapEncoder(t reflect.Type) encoderFunc {
	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		return e.encodeMapEntries(key, value, writer)
	}
}

func encodePartAsJSON(key string, value reflect.Value, writer *multipart.Writer) error {
	var mimeHeader = textproto.MIMEHeader{}
	mimeHeader.Add("Content-Disposition", "form-data; name=\""+key+"\"")
	mimeHeader.Add("Content-Type", "application/json")

	w, err := writer.CreatePart(mimeHeader)
	if err != nil {
		return err
	}
	bytes, err := apijson.MarshalRoot(value.Interface())
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func errorFromDiagnostics(diags diag.Diagnostics) error {
	if diags == nil {
		return nil
	}
	messages := []string{}
	for _, err := range diags {
		messages = append(messages, err.Summary())
		messages = append(messages, err.Detail())
	}
	return errors.New(strings.Join(messages, " "))
}

func valueAs[T any](value reflect.Value) (T, error) {
	typed, ok := reflect.TypeAssert[T](value)
	if !ok {
		return typed, fmt.Errorf("apiform: %s does not hold a %s", value.Type(), reflect.TypeFor[T]())
	}
	return typed, nil
}

func unexpectedValue(value attr.Value, want string) diag.Diagnostics {
	var diags diag.Diagnostics
	diags.AddError("unexpected value type", fmt.Sprintf("expected %s, got %T", want, value))
	return diags
}
