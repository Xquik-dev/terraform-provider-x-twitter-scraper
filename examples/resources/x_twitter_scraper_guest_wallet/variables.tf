# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

variable "guest_wallet_idempotency_key" {
  description = "Cryptographically random UUID v4 for this wallet purchase"
  type        = string
  sensitive   = true
}
