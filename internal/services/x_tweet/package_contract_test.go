// Copyright 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

package x_tweet

import "testing"

func TestNewXTweetDataSourceReturnsConcreteType(t *testing.T) {
	t.Parallel()

	if _, ok := NewXTweetDataSource().(*XTweetDataSource); !ok {
		t.Fatal("NewXTweetDataSource did not return *XTweetDataSource")
	}
}
