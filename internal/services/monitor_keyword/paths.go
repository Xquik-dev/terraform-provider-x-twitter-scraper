// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor_keyword

import (
	"fmt"
	"net/url"
)

const monitorKeywordCollectionPath = "monitors/keywords"

func monitorKeywordItemPath(id string) string {
	return fmt.Sprintf("monitors/keywords/%s", url.PathEscape(id))
}
