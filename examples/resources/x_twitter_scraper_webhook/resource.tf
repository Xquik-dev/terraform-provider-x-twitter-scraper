# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_webhook" "example_webhook" {
  event_types = ["tweet.new", "tweet.reply"]
  url         = "https://example.com/webhook"
}
