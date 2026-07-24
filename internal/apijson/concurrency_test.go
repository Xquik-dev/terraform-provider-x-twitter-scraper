// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"fmt"
	"sync"
	"testing"
)

type concurrentEnvelope struct {
	Value any `json:"value"`
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
				Value: []concurrentValue{{Count: worker}},
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
