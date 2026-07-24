# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_x_user_follow" "example_x_user_follow" {
  account         = "@example"
  idempotency_key = "terraform-follow-user-v1"
  target_id       = "2244994945"
}
