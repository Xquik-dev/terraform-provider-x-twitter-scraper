# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_x_tweet_bookmark" "example_x_tweet_bookmark" {
  account         = "@example"
  idempotency_key = "terraform-bookmark-tweet-v1"
  target_id       = "1234567890"
}
