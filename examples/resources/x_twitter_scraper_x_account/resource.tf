resource "x-twitter-scraper_x_account" "example_x_account" {
  email       = var.x_account_email
  password    = var.x_account_password
  username    = var.x_account_username
  totp_secret = var.x_account_totp_secret
}
