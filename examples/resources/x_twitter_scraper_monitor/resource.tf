# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_monitor" "example_monitor" {
  event_types = ["tweet.new", "tweet.reply"]
  username    = "elonmusk"
}
