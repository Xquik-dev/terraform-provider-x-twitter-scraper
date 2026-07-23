resource "x-twitter-scraper_x_tweet" "example_x_tweet" {
  account         = "@example"
  idempotency_key = "terraform-create-tweet-v1"
  payload_json = jsonencode({
    text = "Published through the Xquik Terraform provider."
  })
}
