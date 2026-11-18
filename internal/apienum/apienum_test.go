// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apienum

import "testing"

func TestDescribeListsEveryValue(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		values []string
		want   string
	}{
		{[]string{"x_user_not_found"}, `The only available value is "x_user_not_found".`},
		{[]string{"post.edited", "post.deleted"}, `Available values are "post.edited" and "post.deleted".`},
		{[]string{"active", "paused", "needs_attention"}, `Available values are "active", "paused" and "needs_attention".`},
	} {
		if got := Describe(test.values); got != test.want {
			t.Errorf("Describe(%q) = %q, want %q", test.values, got, test.want)
		}
	}
}
