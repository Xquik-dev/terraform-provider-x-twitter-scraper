// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/customfield"
)

type concurrentEnvelope struct {
	Value fmt.Stringer `json:"value"`
}

type concurrentValues []concurrentValue

func (values concurrentValues) String() string {
	return fmt.Sprint(len(values))
}

type concurrentValue struct {
	Count int `json:"count"`
}

func TestMarshalRootIsConcurrencySafe(t *testing.T) {
	t.Parallel()

	if _, err := MarshalRoot(concurrentEnvelope{}); err != nil {
		t.Fatalf("warm encoder cache: %v", err)
	}

	const workerCount = 128
	start := make(chan struct{})
	errors := make(chan error, workerCount)
	var workers sync.WaitGroup
	workers.Add(workerCount)

	for worker := range workerCount {
		go func() {
			defer workers.Done()
			<-start

			encoded, err := MarshalRoot(concurrentEnvelope{
				Value: concurrentValues{{Count: worker}},
			})
			if err != nil {
				errors <- err
				return
			}
			if len(encoded) == 0 {
				errors <- fmt.Errorf("worker %d produced no JSON", worker)
			}
		}()
	}

	close(start)
	workers.Wait()
	close(errors)

	for err := range errors {
		t.Error(err)
	}
}

type concurrentItem struct {
	Name types.String `tfsdk:"name" json:"name" api:"computed"`
}

type concurrentItems struct {
	Items customfield.NestedObjectList[concurrentItem] `tfsdk:"items" json:"items" api:"computed_optional"`
}

func TestUnmarshalComputedDecodesUnsetAndSetListsConcurrently(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	const workerCount = 64
	start := make(chan struct{})
	failures := make(chan error, workerCount)
	var workers sync.WaitGroup
	workers.Add(workerCount)

	for worker := range workerCount {
		go func() {
			defer workers.Done()
			model := concurrentItems{Items: customfield.NullObjectList[concurrentItem](ctx)}
			if worker%2 == 1 {
				model.Items = customfield.NewObjectListMust(ctx, []concurrentItem{{Name: types.StringValue("kept")}})
			}
			<-start

			if err := UnmarshalComputed([]byte(`{"items":[{"name":"decoded"}]}`), &model); err != nil {
				failures <- err
				return
			}
			items, diags := model.Items.AsStructSliceT(ctx)
			if worker%2 == 0 && (diags.HasError() || len(items) != 1 || items[0].Name.ValueString() != "decoded") {
				failures <- fmt.Errorf("worker %d decoded %v with %v", worker, items, diags)
			}
		}()
	}

	close(start)
	workers.Wait()
	close(failures)

	for err := range failures {
		t.Error(err)
	}
}

type unreadableItem struct {
	Name types.String `tfsdk:"name" json:"name" api:"computed"`
}

type unreadableHolder struct {
	Item customfield.NestedObject[unreadableItem] `tfsdk:"item" json:"item" api:"computed"`
}

func TestUnmarshalComputedDecodesAnUnreadableObjectFromEmpty(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	model := unreadableHolder{Item: customfield.NestedObject[unreadableItem]{ObjectValue: basetypes.NewObjectValueMust(
		map[string]attr.Type{"name": types.Int64Type},
		map[string]attr.Value{"name": types.Int64Value(1)},
	)}}
	if err := UnmarshalComputed([]byte(`{"item":{"name":"decoded"}}`), &model); err != nil {
		t.Fatal(err)
	}
	item, diags := model.Item.Value(ctx)
	if diags.HasError() || item == nil || item.Name.ValueString() != "decoded" {
		t.Fatalf("decoded item %v with %v", item, diags)
	}
}
