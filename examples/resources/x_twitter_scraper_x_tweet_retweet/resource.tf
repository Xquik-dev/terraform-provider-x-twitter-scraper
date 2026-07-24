# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_x_tweet_retweet" "example_x_tweet_retweet" {
  account         = "@example"
  idempotency_key = "terraform-retweet-v1"
  target_id       = "1234567890"
}
