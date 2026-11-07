# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_x_user_block" "example_x_user_block" {
  account         = "@example"
  idempotency_key = "terraform-block-user-v1"
  target_id       = "1234567890"
}
