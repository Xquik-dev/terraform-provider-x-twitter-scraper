# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_x_profile" "example_x_profile" {
  account         = "@example"
  idempotency_key = "terraform-update-profile-v1"
  payload_json = jsonencode({
    description = "Profile managed through Terraform"
    location    = "Internet"
    name        = "Example"
    url         = "https://example.com"
  })
}
