// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apijson

import "sync"

type Cache[K comparable, F any] struct {
	mu      sync.RWMutex
	entries map[K]F
}

func (c *Cache[K, F]) Load(key K, build func() F, forward func(wait func() F) F) F {
	c.mu.RLock()
	cached, ok := c.entries[key]
	c.mu.RUnlock()
	if ok {
		return cached
	}

	c.mu.Lock()
	if cached, ok := c.entries[key]; ok {
		c.mu.Unlock()
		return cached
	}
	if c.entries == nil {
		c.entries = map[K]F{}
	}
	var (
		ready sync.WaitGroup
		built F
	)
	ready.Add(1)
	c.entries[key] = forward(func() F {
		ready.Wait()
		return built
	})
	c.mu.Unlock()

	built = build()
	ready.Done()
	c.mu.Lock()
	c.entries[key] = built
	c.mu.Unlock()
	return built
}
