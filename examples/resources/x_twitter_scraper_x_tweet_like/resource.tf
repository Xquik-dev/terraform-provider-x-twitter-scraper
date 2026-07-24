# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_x_tweet_like" "example_x_tweet_like" {
  account         = "@example"
  idempotency_key = "terraform-like-tweet-v1"
  target_id       = "1234567890"
}
