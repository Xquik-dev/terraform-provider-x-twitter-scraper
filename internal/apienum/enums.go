// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// go generate ./internal/apienum writes this file from the enums in https://xquik.com/openapi.json. Change cmd/api-enums, not this file.

package apienum

var EventType = []string{
	"tweet.new",
	"tweet.reply",
	"tweet.retweet",
	"tweet.quote",
	"tweet.media",
	"tweet.link",
	"tweet.poll",
	"tweet.mention",
	"tweet.hashtag",
	"tweet.longform",
	"profile.avatar.changed",
	"profile.banner.changed",
	"profile.name.changed",
	"profile.username.changed",
	"profile.bio.changed",
	"profile.location.changed",
	"profile.url.changed",
	"profile.verified.changed",
	"profile.protected.changed",
	"profile.pinned_tweet.changed",
	"profile.unavailable.changed",
	"tweet.deleted",
	"broadcast.started",
	"space.started",
	"following.added",
	"following.removed",
	"follower.added",
	"follower.removed",
	"profile.followers.changed",
	"profile.following.changed",
}

var PostEventType = []string{
	"post.reply.new",
	"post.quote.new",
	"post.likes.changed",
	"post.retweets.changed",
	"post.replies.changed",
	"post.quotes.changed",
	"post.bookmarks.changed",
	"post.views.changed",
	"post.edited",
	"post.deleted",
}

var Spec = map[string][]string{
	"EventType":     EventType,
	"PostEventType": PostEventType,
}
