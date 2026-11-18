// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"context"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/event"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/monitor_post"
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/services/webhook"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type describedAttribute interface {
	GetDescription() string
}

func validates(ctx context.Context, t *testing.T, attribute describedAttribute) func(string) bool {
	t.Helper()
	switch attribute := attribute.(type) {
	case interface{ ListValidators() []validator.List }:
		return func(value string) bool {
			request := validator.ListRequest{
				Path:        path.Root("event_types"),
				ConfigValue: types.ListValueMust(types.StringType, []attr.Value{types.StringValue(value)}),
			}
			for _, check := range attribute.ListValidators() {
				var response validator.ListResponse
				check.ValidateList(ctx, request, &response)
				if response.Diagnostics.HasError() {
					return false
				}
			}
			return true
		}
	case interface{ StringValidators() []validator.String }:
		return func(value string) bool {
			request := validator.StringRequest{Path: path.Root("type"), ConfigValue: types.StringValue(value)}
			for _, check := range attribute.StringValidators() {
				var response validator.StringResponse
				check.ValidateString(ctx, request, &response)
				if response.Diagnostics.HasError() {
					return false
				}
			}
			return true
		}
	}
	t.Fatalf("%T has neither list nor string validators", attribute)
	return nil
}

func TestEventTypeAttributesTakeEveryAPIEventType(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	account := []string{
		"tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote", "tweet.media",
		"tweet.link", "tweet.poll", "tweet.mention", "tweet.hashtag", "tweet.longform",
		"profile.avatar.changed", "profile.banner.changed", "profile.name.changed",
		"profile.username.changed", "profile.bio.changed", "profile.location.changed",
		"profile.url.changed", "profile.verified.changed", "profile.protected.changed",
		"profile.pinned_tweet.changed", "profile.unavailable.changed",
		"tweet.deleted", "broadcast.started", "space.started", "following.added",
		"following.removed", "follower.added", "follower.removed",
		"profile.followers.changed", "profile.following.changed",
	}
	post := []string{
		"post.reply.new", "post.quote.new", "post.likes.changed", "post.retweets.changed",
		"post.replies.changed", "post.quotes.changed", "post.bookmarks.changed",
		"post.views.changed", "post.edited", "post.deleted",
	}
	for _, test := range []struct {
		name      string
		attribute describedAttribute
		takes     []string
		refuses   []string
	}{
		{"monitor resource", monitor.ResourceSchema(ctx).Attributes["event_types"], account, []string{"follower.gained", "post.reply.new"}},
		{"monitor data source", monitor.DataSourceSchema(ctx).Attributes["event_types"], account, []string{"follower.gained", "post.reply.new"}},
		{"webhook resource", webhook.ResourceSchema(ctx).Attributes["event_types"], slices.Concat(account, post), []string{"follower.gained", "follower.lost"}},
		{"event data source", event.DataSourceSchema(ctx).Attributes["type"], slices.Concat(account, post), []string{"follower.gained", "follower.lost"}},
		{"monitor_post resource", monitorpost.ResourceSchema(ctx).Attributes["event_types"], post, []string{"tweet.new", "follower.added"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			takes := validates(ctx, t, test.attribute)
			description := test.attribute.GetDescription()
			var refused, unlisted []string
			for _, value := range test.takes {
				if !takes(value) {
					refused = append(refused, value)
				}
				if !strings.Contains(description, strconv.Quote(value)) {
					unlisted = append(unlisted, value)
				}
			}
			if len(refused) > 0 || len(unlisted) > 0 {
				t.Errorf("The validators refuse %q, and the description leaves out %q.", refused, unlisted)
			}
			for _, value := range test.refuses {
				if takes(value) || strings.Contains(description, strconv.Quote(value)) {
					t.Errorf("The attribute takes or lists %q, which the API does not.", value)
				}
			}
		})
	}
}
