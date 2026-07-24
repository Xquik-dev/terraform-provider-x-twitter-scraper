// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor_keyword

import "testing"

func TestMonitorKeywordItemPath(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id   string
		want string
	}{
		"plain ID": {
			id:   "monitor-123",
			want: "monitors/keywords/monitor-123",
		},
		"path separator": {
			id:   "monitor/123",
			want: "monitors/keywords/monitor%2F123",
		},
		"space": {
			id:   "monitor 123",
			want: "monitors/keywords/monitor%20123",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := monitorKeywordItemPath(test.id); got != test.want {
				t.Fatalf("monitorKeywordItemPath(%q) = %q, want %q", test.id, got, test.want)
			}
		})
	}
}
