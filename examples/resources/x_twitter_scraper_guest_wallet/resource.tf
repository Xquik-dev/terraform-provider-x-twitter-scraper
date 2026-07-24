# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_guest_wallet" "example_guest_wallet" {
  amount_minor    = 1000
  currency        = "usd"
  idempotency_key = var.guest_wallet_idempotency_key
}
