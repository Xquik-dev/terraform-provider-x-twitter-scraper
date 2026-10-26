// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xwrite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type operation struct {
	action         string
	resourceSuffix string
	method         string
	pathTemplate   string
	targetRequired bool
	requiredFields []string
	allowedFields  []string
	anyOfFields    []string
	testPayload    map[string]json.RawMessage
}

var operations = []operation{
	{action: "create_tweet", resourceSuffix: "x_tweet", method: http.MethodPost, pathTemplate: "x/tweets", allowedFields: []string{"text", "reply_to_tweet_id", "community_id", "is_note_tweet", "media"}, anyOfFields: []string{"text", "media"}, testPayload: map[string]json.RawMessage{"text": json.RawMessage(`"hello"`)}},
	{action: "delete_tweet", resourceSuffix: "x_tweet_delete", method: http.MethodDelete, pathTemplate: "x/tweets/%s", targetRequired: true},
	{action: "like", resourceSuffix: "x_tweet_like", method: http.MethodPost, pathTemplate: "x/tweets/%s/like", targetRequired: true},
	{action: "unlike", resourceSuffix: "x_tweet_unlike", method: http.MethodDelete, pathTemplate: "x/tweets/%s/like", targetRequired: true},
	{action: "retweet", resourceSuffix: "x_tweet_retweet", method: http.MethodPost, pathTemplate: "x/tweets/%s/retweet", targetRequired: true},
	{action: "unretweet", resourceSuffix: "x_tweet_unretweet", method: http.MethodDelete, pathTemplate: "x/tweets/%s/retweet", targetRequired: true},
	{action: "follow", resourceSuffix: "x_user_follow", method: http.MethodPost, pathTemplate: "x/users/%s/follow", targetRequired: true},
	{action: "unfollow", resourceSuffix: "x_user_unfollow", method: http.MethodDelete, pathTemplate: "x/users/%s/follow", targetRequired: true},
	{action: "remove_follower", resourceSuffix: "x_user_remove_follower", method: http.MethodPost, pathTemplate: "x/users/%s/remove-follower", targetRequired: true},
	{action: "send_dm", resourceSuffix: "x_dm", method: http.MethodPost, pathTemplate: "x/dm/%s", targetRequired: true, requiredFields: []string{"text"}, allowedFields: []string{"text", "media_ids"}, testPayload: map[string]json.RawMessage{"text": json.RawMessage(`"hello"`)}},
	{action: "upload_media", resourceSuffix: "x_media", method: http.MethodPost, pathTemplate: "x/media", requiredFields: []string{"url"}, allowedFields: []string{"url"}, testPayload: map[string]json.RawMessage{"url": json.RawMessage(`"https://example.com/image.png"`)}},
	{action: "update_profile", resourceSuffix: "x_profile", method: http.MethodPatch, pathTemplate: "x/profile", allowedFields: []string{"name", "description", "location", "url"}, anyOfFields: []string{"name", "description", "location", "url"}, testPayload: map[string]json.RawMessage{"name": json.RawMessage(`"Example"`)}},
	{action: "update_avatar", resourceSuffix: "x_profile_avatar", method: http.MethodPatch, pathTemplate: "x/profile/avatar", requiredFields: []string{"url"}, allowedFields: []string{"url"}, testPayload: map[string]json.RawMessage{"url": json.RawMessage(`"https://example.com/avatar.png"`)}},
	{action: "update_banner", resourceSuffix: "x_profile_banner", method: http.MethodPatch, pathTemplate: "x/profile/banner", requiredFields: []string{"url"}, allowedFields: []string{"url"}, testPayload: map[string]json.RawMessage{"url": json.RawMessage(`"https://example.com/banner.png"`)}},
	{action: "create_community", resourceSuffix: "x_community", method: http.MethodPost, pathTemplate: "x/communities", requiredFields: []string{"name"}, allowedFields: []string{"name", "description"}, testPayload: map[string]json.RawMessage{"name": json.RawMessage(`"Example"`)}},
	{action: "delete_community", resourceSuffix: "x_community_delete", method: http.MethodDelete, pathTemplate: "x/communities/%s", targetRequired: true, requiredFields: []string{"community_name"}, allowedFields: []string{"community_name"}, testPayload: map[string]json.RawMessage{"community_name": json.RawMessage(`"Example"`)}},
	{action: "join_community", resourceSuffix: "x_community_join", method: http.MethodPost, pathTemplate: "x/communities/%s/join", targetRequired: true},
	{action: "leave_community", resourceSuffix: "x_community_leave", method: http.MethodDelete, pathTemplate: "x/communities/%s/join", targetRequired: true},
}

func (o operation) path(targetID string) string {
	if !o.targetRequired {
		return o.pathTemplate
	}
	return fmt.Sprintf(o.pathTemplate, url.PathEscape(targetID))
}

func (o operation) requestBody(account string, payloadJSON string) (map[string]json.RawMessage, error) {
	payload := map[string]json.RawMessage{}
	if strings.TrimSpace(payloadJSON) != "" {
		if err := json.Unmarshal([]byte(payloadJSON), &payload); err != nil {
			return nil, fmt.Errorf("payload_json is invalid. Enter a JSON object: %w", err)
		}
		if payload == nil {
			payload = map[string]json.RawMessage{}
		}
	}

	if _, exists := payload["account"]; exists {
		return nil, fmt.Errorf("payload_json contains account. Use the account attribute instead")
	}

	allowed := make(map[string]struct{}, len(o.allowedFields))
	for _, field := range o.allowedFields {
		allowed[field] = struct{}{}
	}
	for field := range payload {
		if _, ok := allowed[field]; !ok {
			return nil, fmt.Errorf("payload_json contains unsupported field %q for %s. Remove it", field, o.action)
		}
	}
	for _, field := range o.requiredFields {
		if !present(payload[field]) {
			return nil, fmt.Errorf("payload_json is missing field %q for %s. Add it", field, o.action)
		}
	}
	if len(o.anyOfFields) > 0 {
		found := false
		for _, field := range o.anyOfFields {
			found = found || present(payload[field])
		}
		if !found {
			fields := append([]string(nil), o.anyOfFields...)
			sort.Strings(fields)
			return nil, fmt.Errorf("payload_json needs at least one of %s for %s. Add one", strings.Join(fields, ", "), o.action)
		}
	}

	encodedAccount, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}
	payload["account"] = encodedAccount
	return payload, nil
}

func present(value json.RawMessage) bool {
	trimmed := bytes.TrimSpace(value)
	switch {
	case len(trimmed) == 0 || string(trimmed) == "null":
		return false
	case trimmed[0] == '"':
		var text string
		_ = json.Unmarshal(trimmed, &text)
		return strings.TrimSpace(text) != ""
	case trimmed[0] == '[':
		var items []json.RawMessage
		_ = json.Unmarshal(trimmed, &items)
		return len(items) > 0
	default:
		return true
	}
}
