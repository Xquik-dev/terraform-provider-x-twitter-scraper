// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor_post_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	resourceType   = "x-twitter-scraper_monitor_post"
	monitorID      = "31"
	postID         = "2104984505677430978"
	postLink       = "https://x.com/xquik/status/2104984505677430978"
	collectionPath = "/api/v1/monitors/posts"
	monitorPath    = "/api/v1/monitors/posts/31"
)

func TestCreateWithoutEventTypesOrIsActiveKeepsTheAPIDefaults(t *testing.T) {
	t.Parallel()

	fake, baseURL := newFakeXquik(t, nil)
	tf := newTerraform(t, baseURL)
	config := tf.config(map[string]tftypes.Value{"post": text(postLink)})

	created := tf.apply(tf.absent(), config, tf.plan(tf.absent(), config))

	assertXquikWrites(t, fake, `POST /api/v1/monitors/posts {"post":"https://x.com/xquik/status/2104984505677430978"}`)
	assertAttribute(t, created, "id", text(monitorID))
	assertAttribute(t, created, "post_id", text(postID))
	assertAttribute(t, created, "event_types", texts(
		"post.reply.new",
		"post.quote.new",
		"post.likes.changed",
		"post.retweets.changed",
		"post.edited",
		"post.deleted",
	))
	assertAttribute(t, created, "is_active", flag(true))

	refreshed := tf.refresh(created)
	tf.assertNoChanges(refreshed, tf.plan(refreshed, config))
}

func TestCreateWithIsActiveFalsePausesTheNewMonitor(t *testing.T) {
	t.Parallel()

	fake, baseURL := newFakeXquik(t, nil)
	tf := newTerraform(t, baseURL)
	config := tf.config(map[string]tftypes.Value{
		"post":        text(postLink),
		"event_types": texts("post.reply.new"),
		"is_active":   flag(false),
	})

	created := tf.apply(tf.absent(), config, tf.plan(tf.absent(), config))

	assertXquikWrites(t, fake,
		`POST /api/v1/monitors/posts {"eventTypes":["post.reply.new"],"post":"https://x.com/xquik/status/2104984505677430978"}`,
		`PATCH /api/v1/monitors/posts/31 {"isActive":false}`,
	)
	assertAttribute(t, created, "is_active", flag(false))
	assertAttribute(t, tf.refresh(created), "is_active", flag(false))
}

func TestImportWithThePostIDPlansNoChanges(t *testing.T) {
	t.Parallel()

	_, baseURL := newFakeXquik(t, storedMonitor())
	tf := newTerraform(t, baseURL)
	imported := tf.refresh(tf.importState(monitorID))

	planned := tf.plan(imported, tf.config(map[string]tftypes.Value{"post": text(postID)}))

	tf.assertNoChanges(imported, planned)
}

func TestImportWithThePostLinkUpdatesInPlaceWithoutWrites(t *testing.T) {
	t.Parallel()

	fake, baseURL := newFakeXquik(t, storedMonitor())
	tf := newTerraform(t, baseURL)
	config := tf.config(map[string]tftypes.Value{"post": text(postLink)})
	imported := tf.refresh(tf.importState(monitorID))

	planned := tf.plan(imported, config)
	if len(planned.RequiresReplace) != 0 {
		t.Fatalf("the plan replaces the monitor for %v", planned.RequiresReplace)
	}
	updated := tf.apply(imported, config, planned)

	assertXquikWrites(t, fake)
	assertAttribute(t, updated, "id", text(monitorID))
	assertAttribute(t, updated, "post", text(postLink))
}

func TestConfigRejectsValuesTheAPIRewritesOrRefuses(t *testing.T) {
	t.Parallel()

	tf := newTerraform(t, "https://example.invalid/api/v1/")
	accepted := tf.config(map[string]tftypes.Value{
		"post":        text(postLink),
		"name":        text(strings.Repeat("ş", 80)),
		"event_types": texts("post.reply.new", "post.edited"),
	})
	requireNoErrors(t, "validate an accepted configuration", nil, tf.validate(accepted))

	for name, attributes := range map[string]map[string]tftypes.Value{
		"repeated event type":    {"event_types": texts("post.reply.new", "post.reply.new")},
		"uppercase event type":   {"event_types": texts("POST.REPLY.NEW")},
		"empty name":             {"name": text("")},
		"81 character name":      {"name": text(strings.Repeat("a", 81))},
		"double space in name":   {"name": text("Launch  post")},
		"outer space in name":    {"name": text(" Launch post")},
		"tab in name":            {"name": text("Launch\tpost")},
		"no-break space in name": {"name": text("Launch post")},
	} {
		attributes["post"] = text(postLink)
		if !hasError(tf.validate(tf.config(attributes))) {
			t.Errorf("%s passed validation", name)
		}
	}
}

type fakeXquik struct {
	mutex    sync.Mutex
	monitor  map[string]any
	requests []string
}

func newFakeXquik(t *testing.T, monitor map[string]any) (*fakeXquik, string) {
	t.Helper()
	fake := &fakeXquik{monitor: monitor}
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fake.serve(t, writer, request)
	}))
	t.Cleanup(server.Close)
	return fake, server.URL + "/api/v1/"
}

func (fake *fakeXquik) serve(t *testing.T, writer http.ResponseWriter, request *http.Request) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()

	entry := request.Method + " " + request.URL.Path
	body, err := io.ReadAll(request.Body)
	if err != nil {
		t.Errorf("read %s: %v", entry, err)
	}
	var fields map[string]any
	if len(body) > 0 {
		if err := json.Unmarshal(body, &fields); err != nil {
			t.Errorf("decode %s: %v", entry, err)
		}
		canonical, err := json.Marshal(fields)
		if err != nil {
			t.Errorf("encode %s: %v", entry, err)
		}
		entry += " " + string(canonical)
	}
	fake.requests = append(fake.requests, entry)

	writer.Header().Set("Content-Type", "application/json")
	switch {
	case request.Method == http.MethodPost && request.URL.Path == collectionPath:
		fake.monitor = storedMonitor()
		for _, key := range []string{"name", "eventTypes"} {
			if value, ok := fields[key]; ok {
				fake.monitor[key] = value
			}
		}
		writer.WriteHeader(http.StatusCreated)
	case request.Method == http.MethodGet && request.URL.Path == monitorPath && fake.monitor != nil:
	case request.Method == http.MethodPatch && request.URL.Path == monitorPath && fake.monitor != nil:
		for _, key := range []string{"name", "eventTypes", "isActive"} {
			if value, ok := fields[key]; ok {
				fake.monitor[key] = value
			}
		}
	default:
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte(`{"error":"not_found"}`))
		return
	}
	if err := json.NewEncoder(writer).Encode(fake.monitor); err != nil {
		t.Errorf("encode the monitor for %s: %v", entry, err)
	}
}

func (fake *fakeXquik) writes() []string {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	return slices.DeleteFunc(slices.Clone(fake.requests), func(entry string) bool {
		return strings.HasPrefix(entry, http.MethodGet+" ")
	})
}

func storedMonitor() map[string]any {
	return map[string]any{
		"id":             monitorID,
		"postId":         postID,
		"postUrl":        postLink,
		"authorId":       "1661816432185823232",
		"authorUsername": "xquik",
		"authorName":     "Xquik",
		"text":           "We shipped post monitors.",
		"postedAt":       "2026-09-29T17:19:49.000Z",
		"eventTypes": []string{
			"post.reply.new",
			"post.quote.new",
			"post.likes.changed",
			"post.retweets.changed",
			"post.edited",
			"post.deleted",
		},
		"isActive":      true,
		"createdAt":     "2026-10-22T11:00:00Z",
		"nextBillingAt": "2026-10-22T12:00:00Z",
	}
}

type terraform struct {
	t      *testing.T
	server tfprotov6.ProviderServer
	schema *tfprotov6.Schema
}

func newTerraform(t *testing.T, baseURL string) terraform {
	t.Helper()
	ctx := context.Background()
	server := providerserver.NewProtocol6(internal.NewProvider("test")())()
	schemas, err := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	requireNoErrors(t, "read the provider schema", err, schemas.Diagnostics)
	providerConfig := objectOf(t, schemas.Provider, map[string]tftypes.Value{
		"api_key":  text("test-api-key"),
		"base_url": text(baseURL),
	})
	configured, err := server.ConfigureProvider(ctx, &tfprotov6.ConfigureProviderRequest{
		Config: dynamic(t, providerConfig),
	})
	requireNoErrors(t, "configure the provider", err, configured.Diagnostics)
	schema, ok := schemas.ResourceSchemas[resourceType]
	if !ok {
		t.Fatalf("the provider has no %s resource", resourceType)
	}
	return terraform{t: t, server: server, schema: schema}
}

func (tf terraform) config(values map[string]tftypes.Value) tftypes.Value {
	tf.t.Helper()
	return objectOf(tf.t, tf.schema, values)
}

func (tf terraform) absent() tftypes.Value {
	return tftypes.NewValue(tf.schema.ValueType(), nil)
}

func (tf terraform) validate(config tftypes.Value) []*tfprotov6.Diagnostic {
	tf.t.Helper()
	validated, err := tf.server.ValidateResourceConfig(context.Background(), &tfprotov6.ValidateResourceConfigRequest{
		TypeName: resourceType,
		Config:   dynamic(tf.t, config),
	})
	if err != nil {
		tf.t.Fatalf("validate the configuration: %v", err)
	}
	return validated.Diagnostics
}

func (tf terraform) plan(prior tftypes.Value, config tftypes.Value) *tfprotov6.PlanResourceChangeResponse {
	tf.t.Helper()
	requireNoErrors(tf.t, "validate the configuration", nil, tf.validate(config))
	planned, err := tf.server.PlanResourceChange(context.Background(), &tfprotov6.PlanResourceChangeRequest{
		TypeName:         resourceType,
		PriorState:       dynamic(tf.t, prior),
		ProposedNewState: dynamic(tf.t, tf.proposedNewState(prior, config)),
		Config:           dynamic(tf.t, config),
	})
	requireNoErrors(tf.t, "plan", err, planned.Diagnostics)
	return planned
}

func (tf terraform) proposedNewState(prior tftypes.Value, config tftypes.Value) tftypes.Value {
	tf.t.Helper()
	if prior.IsNull() {
		prior = tf.config(nil)
	}
	proposed := attributesOf(tf.t, prior)
	configured := attributesOf(tf.t, config)
	for _, attribute := range tf.schema.Block.Attributes {
		if !attribute.Computed || !configured[attribute.Name].IsNull() {
			proposed[attribute.Name] = configured[attribute.Name]
		}
	}
	return tftypes.NewValue(prior.Type(), proposed)
}

func (tf terraform) apply(prior tftypes.Value, config tftypes.Value, planned *tfprotov6.PlanResourceChangeResponse) tftypes.Value {
	tf.t.Helper()
	applied, err := tf.server.ApplyResourceChange(context.Background(), &tfprotov6.ApplyResourceChangeRequest{
		TypeName:       resourceType,
		PriorState:     dynamic(tf.t, prior),
		PlannedState:   planned.PlannedState,
		Config:         dynamic(tf.t, config),
		PlannedPrivate: planned.PlannedPrivate,
	})
	requireNoErrors(tf.t, "apply", err, applied.Diagnostics)
	newState := tf.decode(applied.NewState)
	newAttributes := attributesOf(tf.t, newState)
	for name, plannedValue := range attributesOf(tf.t, tf.decode(planned.PlannedState)) {
		if plannedValue.IsFullyKnown() && !plannedValue.Equal(newAttributes[name]) {
			tf.t.Errorf("apply changed the planned %s from %v to %v", name, plannedValue, newAttributes[name])
		}
	}
	return newState
}

func (tf terraform) importState(id string) tftypes.Value {
	tf.t.Helper()
	imported, err := tf.server.ImportResourceState(context.Background(), &tfprotov6.ImportResourceStateRequest{
		TypeName: resourceType,
		ID:       id,
	})
	requireNoErrors(tf.t, "import", err, imported.Diagnostics)
	if len(imported.ImportedResources) != 1 {
		tf.t.Fatalf("import returned %d resources, want 1", len(imported.ImportedResources))
	}
	return tf.decode(imported.ImportedResources[0].State)
}

func (tf terraform) refresh(state tftypes.Value) tftypes.Value {
	tf.t.Helper()
	read, err := tf.server.ReadResource(context.Background(), &tfprotov6.ReadResourceRequest{
		TypeName:     resourceType,
		CurrentState: dynamic(tf.t, state),
	})
	requireNoErrors(tf.t, "refresh", err, read.Diagnostics)
	return tf.decode(read.NewState)
}

func (tf terraform) assertNoChanges(prior tftypes.Value, planned *tfprotov6.PlanResourceChangeResponse) {
	tf.t.Helper()
	plannedState := tf.decode(planned.PlannedState)
	if !plannedState.Equal(prior) {
		changes, err := prior.Diff(plannedState)
		tf.t.Errorf("the plan changes %v (%v)", changes, err)
	}
}

func (tf terraform) decode(value *tfprotov6.DynamicValue) tftypes.Value {
	tf.t.Helper()
	decoded, err := value.Unmarshal(tf.schema.ValueType())
	if err != nil {
		tf.t.Fatalf("decode the state: %v", err)
	}
	return decoded
}

func objectOf(t *testing.T, schema *tfprotov6.Schema, values map[string]tftypes.Value) tftypes.Value {
	t.Helper()
	objectType, ok := schema.ValueType().(tftypes.Object)
	if !ok {
		t.Fatalf("schema type %v is not an object", schema.ValueType())
	}
	attributes := make(map[string]tftypes.Value, len(objectType.AttributeTypes))
	for name, attributeType := range objectType.AttributeTypes {
		attributes[name] = tftypes.NewValue(attributeType, nil)
	}
	for name, value := range values {
		if _, ok := attributes[name]; !ok {
			t.Fatalf("the schema has no %s attribute", name)
		}
		attributes[name] = value
	}
	return tftypes.NewValue(objectType, attributes)
}

func attributesOf(t *testing.T, object tftypes.Value) map[string]tftypes.Value {
	t.Helper()
	var attributes map[string]tftypes.Value
	if err := object.As(&attributes); err != nil {
		t.Fatalf("read the attributes of %v: %v", object, err)
	}
	return attributes
}

func dynamic(t *testing.T, value tftypes.Value) *tfprotov6.DynamicValue {
	t.Helper()
	encoded, err := tfprotov6.NewDynamicValue(value.Type(), value)
	if err != nil {
		t.Fatalf("encode %v: %v", value, err)
	}
	return &encoded
}

func text(value string) tftypes.Value {
	return tftypes.NewValue(tftypes.String, value)
}

func flag(value bool) tftypes.Value {
	return tftypes.NewValue(tftypes.Bool, value)
}

func texts(values ...string) tftypes.Value {
	elements := make([]tftypes.Value, len(values))
	for index, value := range values {
		elements[index] = text(value)
	}
	return tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, elements)
}

func assertAttribute(t *testing.T, state tftypes.Value, name string, want tftypes.Value) {
	t.Helper()
	if got := attributesOf(t, state)[name]; !got.Equal(want) {
		t.Errorf("%s = %v, want %v", name, got, want)
	}
}

func assertXquikWrites(t *testing.T, fake *fakeXquik, want ...string) {
	t.Helper()
	if got := fake.writes(); !slices.Equal(got, want) {
		t.Errorf("Xquik received writes\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

func hasError(diagnostics []*tfprotov6.Diagnostic) bool {
	return slices.ContainsFunc(diagnostics, func(diagnostic *tfprotov6.Diagnostic) bool {
		return diagnostic.Severity == tfprotov6.DiagnosticSeverityError
	})
}

func requireNoErrors(t *testing.T, step string, err error, diagnostics []*tfprotov6.Diagnostic) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s: %v", step, err)
	}
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("%s: %s. %s", step, diagnostic.Summary, diagnostic.Detail)
		}
	}
}
