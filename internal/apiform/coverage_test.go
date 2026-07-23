// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package apiform

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type coverageReader struct {
	io.Reader
}

func (coverageReader) Name() string {
	return `/tmp/file"name.txt`
}

func (coverageReader) ContentType() string {
	return "text/plain"
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func TestEncoderHandlesAllPrimitiveWidthsAndInterfaces(t *testing.T) {
	t.Parallel()

	input := struct {
		Int8      int8      `form:"int8"`
		Uint8     uint8     `form:"uint8"`
		Interface any       `form:"interface"`
		Nil       any       `form:"nil"`
		Date      time.Time `form:"date" format:"date"`
	}{
		Int8:      -8,
		Uint8:     8,
		Interface: "value",
		Date:      time.Date(2026, time.July, 23, 12, 0, 0, 0, time.UTC),
	}

	body, err := marshalForm(input)
	if err != nil {
		t.Fatalf("marshal form: %v", err)
	}
	for _, expected := range []string{
		`name="int8"`,
		"-8",
		`name="uint8"`,
		"8",
		`name="interface"`,
		"value",
		"2026-07-23",
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("multipart body does not contain %q", expected)
		}
	}

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	if err := MarshalRoot(nil, writer); err != nil {
		t.Fatalf("marshal nil: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}

	unsupported := struct {
		Value complex64 `form:"value"`
	}{Value: 1}
	if _, err := marshalForm(unsupported); err == nil {
		t.Fatal("unsupported primitive type was accepted")
	}
}

func TestReaderAndMapEncodingContracts(t *testing.T) {
	t.Parallel()

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	reader := coverageReader{Reader: strings.NewReader("contents")}
	encoder := (&encoder{}).newReaderTypeEncoder()
	if err := encoder(`file"name`, reflect.ValueOf(reader), writer); err != nil {
		t.Fatalf("encode reader: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}
	body := buffer.String()
	for _, expected := range []string{
		`name="file\"name"`,
		`filename="file\"name.txt"`,
		"Content-Type: text/plain",
		"contents",
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("reader body does not contain %q", expected)
		}
	}

	var failingBuffer bytes.Buffer
	failingWriter := multipart.NewWriter(&failingBuffer)
	if err := encoder("file", reflect.ValueOf(failingReader{}), failingWriter); err == nil {
		t.Fatal("reader failure was ignored")
	}

	if _, err := marshalForm(map[int]string{1: "value"}); err == nil {
		t.Fatal("map with non-string keys was accepted")
	}
}

func TestTagParsingAndDiagnosticFormatting(t *testing.T) {
	t.Parallel()

	type tagged struct {
		Value string `form:"value,required,extras,metadata,computed,force_encode" format:"date-time"`
		Other string
	}
	valueField, ok := reflect.TypeOf(tagged{}).FieldByName("Value")
	if !ok {
		t.Fatal("tagged field is missing")
	}
	tag, ok := parseFormStructTag(valueField)
	if !ok || tag.name != "value" || !tag.required || !tag.extras || !tag.metadata || !tag.computed || !tag.forceEncode {
		t.Fatalf("parsed tag = %#v", tag)
	}
	format, ok := parseFormatStructTag(valueField)
	if !ok || format != "date-time" {
		t.Fatalf("format = %q, present = %t", format, ok)
	}
	otherField, ok := reflect.TypeOf(tagged{}).FieldByName("Other")
	if !ok {
		t.Fatal("untagged field is missing")
	}
	if _, ok := parseFormStructTag(otherField); ok {
		t.Fatal("untagged field was accepted")
	}
	if _, ok := parseFormatStructTag(otherField); ok {
		t.Fatal("missing format was accepted")
	}

	if escaped := escapeQuotes(`a\"b`); escaped != `a\\\"b` {
		t.Fatalf("escaped string = %q", escaped)
	}
	if err := errorFromDiagnostics(nil); err != nil {
		t.Fatalf("nil diagnostics returned an error: %v", err)
	}
	diagnostics := diag.Diagnostics{diag.NewErrorDiagnostic("summary", "detail")}
	err := errorFromDiagnostics(diagnostics)
	if err == nil || !strings.Contains(err.Error(), "summary detail") {
		t.Fatalf("diagnostic error = %v", err)
	}
}

func TestEncoderHandlesNullPointersEmbeddedFieldsAndReaders(t *testing.T) {
	t.Parallel()

	type EmbeddedFields struct {
		Embedded string `form:"embedded"`
	}
	type edgeForm struct {
		EmbeddedFields
		ignored string              `form:"ignored"`
		Pointer *string             `form:"pointer"`
		Reader  coverageReader      `form:"reader"`
		List    basetypes.ListValue `form:"list"`
		Skip    string              `form:"-"`
	}
	input := edgeForm{
		EmbeddedFields: EmbeddedFields{Embedded: "embedded-value"},
		Reader:         coverageReader{Reader: strings.NewReader("reader-value")},
		List:           basetypes.NewListNull(basetypes.StringType{}),
		Skip:           "must-not-appear",
	}
	body, err := marshalForm(input)
	if err != nil {
		t.Fatalf("marshal edge form: %v", err)
	}
	if !strings.Contains(body, "embedded-value") || !strings.Contains(body, "reader-value") {
		t.Fatalf("edge form omitted expected fields: %s", body)
	}
	if strings.Contains(body, "must-not-appear") {
		t.Fatal("dash-tagged field was encoded")
	}

	var prefixed bytes.Buffer
	writer := multipart.NewWriter(&prefixed)
	formEncoder := (&encoder{root: true, dateFormat: time.RFC3339}).newStructTypeEncoder(reflect.TypeOf(input))
	if err := formEncoder("prefix", reflect.ValueOf(input), writer); err != nil {
		t.Fatalf("encode prefixed form: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close prefixed form: %v", err)
	}
	if !strings.Contains(prefixed.String(), `name="prefix.embedded"`) {
		t.Fatalf("prefixed form omitted its prefix: %s", prefixed.String())
	}
}

func TestEncoderRejectsInvalidTerraformAndNestedValues(t *testing.T) {
	t.Parallel()

	invalidTime := struct {
		Value timetypes.RFC3339 `form:"value"`
	}{
		Value: timetypes.RFC3339{StringValue: basetypes.NewStringValue("invalid")},
	}
	if _, err := marshalForm(invalidTime); err == nil {
		t.Fatal("invalid RFC3339 value was accepted")
	}

	unsupportedTuple := struct {
		Value basetypes.TupleValue `form:"value"`
	}{
		Value: basetypes.NewTupleValueMust(
			[]attr.Type{basetypes.StringType{}},
			[]attr.Value{basetypes.NewStringValue("value")},
		),
	}
	if _, err := marshalForm(unsupportedTuple); err == nil {
		t.Fatal("unsupported Terraform tuple was accepted")
	}
	if _, err := marshalForm([]complex64{1}); err == nil {
		t.Fatal("unsupported array element was accepted")
	}
	if _, err := marshalForm(map[string]complex64{"value": 1}); err == nil {
		t.Fatal("unsupported map value was accepted")
	}
}

func marshalForm(value any) (string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	if err := MarshalRoot(value, writer); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
