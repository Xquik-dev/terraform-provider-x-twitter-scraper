// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package monitorkeyword

import (
	"fmt"
	"net/url"
)

const monitorKeywordCollectionPath = "monitors/keywords"

func monitorKeywordItemPath(id string) string {
	return fmt.Sprintf(monitorKeywordCollectionPath+"/%s", url.PathEscape(id))
}
