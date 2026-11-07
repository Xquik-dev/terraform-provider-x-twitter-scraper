# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_x_user_mute" "example_x_user_mute" {
  account         = "@example"
  idempotency_key = "terraform-mute-user-v1"
  target_id       = "1234567890"
}
