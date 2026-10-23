// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor_post

import (
	"fmt"
	"net/url"
)

const monitorPostCollectionPath = "monitors/posts"

func monitorPostItemPath(id string) string {
	return fmt.Sprintf(monitorPostCollectionPath+"/%s", url.PathEscape(id))
}
