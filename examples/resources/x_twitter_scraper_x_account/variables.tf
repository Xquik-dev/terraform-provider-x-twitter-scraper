variable "x_account_email" {
  description = "Email address for the connected X account"
  type        = string
  sensitive   = true
}

variable "x_account_password" {
  description = "Password for the connected X account"
  type        = string
  sensitive   = true
}

variable "x_account_username" {
  description = "Username for the connected X account"
  type        = string
}

variable "x_account_totp_secret" {
  description = "TOTP secret for the connected X account"
  type        = string
  sensitive   = true
}
