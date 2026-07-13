package monitor_keyword

import "net/url"

const monitorKeywordCollectionPath = "monitors/keywords"

func monitorKeywordItemPath(id string) string {
	return monitorKeywordCollectionPath + "/" + url.PathEscape(id)
}
