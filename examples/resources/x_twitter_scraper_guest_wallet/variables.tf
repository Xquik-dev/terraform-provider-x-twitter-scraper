variable "guest_wallet_idempotency_key" {
  description = "Cryptographically random UUID v4 for this wallet purchase"
  type        = string
  sensitive   = true
}
