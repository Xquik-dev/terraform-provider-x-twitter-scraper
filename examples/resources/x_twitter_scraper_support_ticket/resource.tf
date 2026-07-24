# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_support_ticket" "example_support_ticket" {
  body    = "I am unable to connect my X account. Please help."
  subject = "Cannot connect X account"
}
