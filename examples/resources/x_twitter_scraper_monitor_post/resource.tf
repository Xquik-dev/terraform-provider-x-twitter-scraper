# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_monitor_post" "example_monitor_post" {
  post        = "2112339130483151117"
  event_types = ["post.reply.new", "post.likes.changed"]
}
