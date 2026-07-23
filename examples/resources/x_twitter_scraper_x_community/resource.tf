resource "x-twitter-scraper_x_community" "example_x_community" {
  account         = "@example"
  idempotency_key = "terraform-create-community-v1"
  payload_json = jsonencode({
    name        = "Example Community"
    description = "A community managed through Terraform"
  })
}
