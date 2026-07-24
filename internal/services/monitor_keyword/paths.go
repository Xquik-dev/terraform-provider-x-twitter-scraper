// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitor_keyword

import "net/url"

const monitorKeywordCollectionPath = "monitors/keywords"

func monitorKeywordItemPath(id string) string {
	return monitorKeywordCollectionPath + "/" + url.PathEscape(id)
}
