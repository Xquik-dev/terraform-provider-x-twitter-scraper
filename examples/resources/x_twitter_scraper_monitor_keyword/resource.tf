# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_monitor_keyword" "example_monitor_keyword" {
  event_types = ["tweet.new"]
  query       = "xquik OR \"x api\""
}
