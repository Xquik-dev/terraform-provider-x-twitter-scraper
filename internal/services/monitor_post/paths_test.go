// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitorpost

import "testing"

func TestMonitorPostItemPath(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id   string
		want string
	}{
		"plain ID": {
			id:   "31",
			want: "monitors/posts/31",
		},
		"path separator": {
			id:   "monitor/31",
			want: "monitors/posts/monitor%2F31",
		},
		"space": {
			id:   "monitor 31",
			want: "monitors/posts/monitor%2031",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := monitorPostItemPath(test.id); got != test.want {
				t.Fatalf("monitorPostItemPath(%q) = %q, want %q", test.id, got, test.want)
			}
		})
	}
}
